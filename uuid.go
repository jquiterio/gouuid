package uuid

import (
	"bytes"
	"crypto/rand"
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

type UUID [16]byte

func randomize() (b []byte, err error) {
	b = make([]byte, 16)
	_, err = rand.Read(b)

	if err != nil {
		return nil, err
	}
	return b, nil
}
func Compare(ua, ub UUID) bool {
	return bytes.Equal(ua[:], ub[:])
}

func (uuid *UUID) ToString() string {
	b := uuid
	str := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return strings.ToLower(str)
}

func ToUUID(str string) (uuid UUID, err error) {
	if len(str) != 36 || str[8:9] != "-" || str[13:14] != "-" || str[18:19] != "-" || str[23:24] != "-" {
		return uuid, errors.New("Invalid uuid")
	}
	b, err := hex.DecodeString(str[0:8] + str[9:13] + str[14:18] + str[19:23] + str[24:])
	if err != nil {
		return uuid, errors.New("error decoding uuid")
	}
	_, err = io.ReadFull(bytes.NewBuffer(b), uuid[:])
	return uuid, err
}

func New() UUID {
	var uuid UUID
	b, _ := randomize()
	_, _ = io.ReadFull(bytes.NewBuffer(b), uuid[:])

	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid
}

func (uuid UUID) Value() (driver.Value, error) {
	str := uuid.ToString()
	return str, nil
}

func (uuid *UUID) Scan(value interface{}) error {
	switch value := value.(type) {
	case nil:
		return nil
	case UUID:
		*uuid = value
		return nil
	case []byte:
		if len(value) == 16 {
			return uuid.Scan(uuid.ToString())
		}
		copy((*uuid)[:], value)
		return nil
	case string:
		//str := uuid.ToString()
		*uuid, _ = ToUUID(value)
		return nil
	default:
		return fmt.Errorf("unable to convert to UUID")
	}
}
