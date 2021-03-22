package uuid

import (
	"fmt"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type test struct{}

var _ = check.Suite(&test{})

func (s *test) TestNew(c *check.C) {
	uuid := New()
	str := uuid.ToString()
	u, err := ToUUID(str)

	fmt.Println(u)
	c.Assert(err, check.IsNil)
	c.Assert(len(str), check.Equals, 36)
	c.Assert(len(u), check.Equals, 16)
}

func BenchmarkUUID(b *testing.B) {
	uuid := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.ToString()
	}
}
