package models

import (
	"strconv"
	"time"

	. "github.com/alligrader/gradebook-backend/models/privileges"
)

type User struct {
	ID        int64
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Username  string     `db:"username"`
	Password  []byte     `db:"password"`
	Status    UserStatus `db:"status"`

	Privileges
	CreatedAt   time.Time `db:"created_at"`
	LastUpdated time.Time `db:"last_updated"`
}

func (user *User) InsertColumns() string {
	return "first_name, last_name, username, password, status"
}

func (user *User) GetColumns() string {
	return "first_name, last_name, username, status, created_at, last_updated"
}

func (user *User) Fields() []interface{} {
	return []interface{}{
		&user.ID, &user.FirstName, &user.LastName,
		&user.Username, &user.CreatedAt, &user.LastUpdated,
	}
}

func (user *User) GetID() string {
	return strconv.FormatInt(user.ID, 10)
}

func (user *User) Equals(other *User) bool {
	if other == nil {
		return false
	}
	return user.FirstName == other.FirstName &&
		user.LastName == other.LastName &&
		user.Username == other.Username
}

// User Status
type UserStatus int

const (
	Invited UserStatus = iota
	Active
)
