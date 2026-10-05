// TEST ONLY: deliberately insecure code to verify PR-gate diff flagging. Not real code.
package bogus

import (
	"crypto/md5"
	"database/sql"
	"fmt"
)

const dbPassword = "x9Kq2LmPz7TvB4nW8rYc3HdFgJ6sAe1U"

func FindUser(db *sql.DB, name string) (*sql.Rows, error) {
	return db.Query(fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", name))
}

func WeakHash(in string) []byte {
	h := md5.Sum([]byte(in))
	return h[:]
}
