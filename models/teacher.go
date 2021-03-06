package models

import (
	"strconv"
	"time"
)

type Teacher struct {
	ID int64
	User

	CreatedAt   time.Time
	LastUpdated time.Time
}

func (teacher *Teacher) GetID() string {
	return strconv.FormatInt(teacher.ID, 10)
}

func (teacher *Teacher) Equals(other *Teacher) bool {
	if other == nil {
		return false
	}
	return teacher.User.Equals(&other.User)
}
