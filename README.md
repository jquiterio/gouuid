# gouuid

## install

go get github.com/jquiterio/uuid

## use

```
import "github.com/jquiterio/uuid"

func main() {
  // Create a new UUID
  id := uuid.New()

  // Get UUID string
  str := id.String

  fmt.Println(str)

  // From String to UUID
  id = ToUUID(str)

  // Compare uuids
  newuuid := new(UUID)
  ok := uuid.Compare(id,newuuid)

}
```

### As Database ID

```
// As Database ID

type User struct {
  ID uuid.UUID `gorm:"type:uuid"`
  ProfileID uuid.UUID `gorm:"type:uuid"`
  Name string
}

type Prodile struct {
  ID uuid.UUID `gorm:"type:uuid"`
  Name string
}
```