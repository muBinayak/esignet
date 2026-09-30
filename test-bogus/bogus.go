package bogus

import (
	"database/sql"
	"fmt"
)

// Deliberately vulnerable file. Exists only to verify reviewdog posts inline
// PR comments for secret-scan and semgrep. Not part of the real esignet
// codebase.

const dbPassword = "Tr0ub4dor&3xk9mQp2ZvL7wRfN4hYbJ8"

func GetUser(db *sql.DB, userID string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = '%s'", userID)
	return db.Query(query)
}
