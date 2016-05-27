package db

import (
	"fmt"

	. "github.com/alligrader/gradebook-backend/models"
	"github.com/alligrader/gradebook-backend/util"
	"github.com/jmoiron/sqlx"
	sq "gopkg.in/Masterminds/squirrel.v1"

	_ "github.com/Sirupsen/logrus"
)

func (maker *personMaker) Create(person *Person) error {

	query := fmt.Sprintf(queries["create_person"], person.InsertColumns())

	result, err := util.PrepAndExec(query, maker, person.FirstName, person.LastName, person.Username, string(person.Password))
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	person.ID = id

	return nil
}

func (maker *personMaker) GetByID(id int64) (*Person, error) {

	var (
		person *Person = &Person{}
		query  string  = fmt.Sprintf(queries["get_person"], person.GetColumns())
		err    error   = util.GetAndMarshal(query, maker, person, id)
	)

	if err != nil {
		return nil, err
	}

	return person, nil
}

func (maker *studentMaker) Create(student *Student) error {

	// TODO make PersonStore.Create private.
	PersonStore.Create(&student.Person)

	query := queries["create_student"]

	result, err := util.PrepAndExec(query, maker, student.Person.ID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	student.ID = id

	return nil

}
func (maker *studentMaker) Update(student *Student) error {
	return nil

}
func (maker *studentMaker) GetByID(id int64) (*Student, error) {
	var (
		student *Student = &Student{}
		query   string   = queries["get_student"]
		err     error    = util.GetAndMarshal(query, maker, student, id)
	)

	if err != nil {
		return nil, err
	}

	return student, nil
}

func (maker *studentMaker) Destroy(student *Student) error {
	return nil
}

func (maker *teacherMaker) Create(teacher *Teacher) error {
	PersonStore.Create(&teacher.Person)

	query, _, err := sq.
		Insert("teacher").Columns("person_id").Values("person_id").
		ToSql()
	if err != nil {
		return err
	}

	result, err := util.PrepAndExec(query, maker, teacher.Person.ID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	teacher.ID = id

	return nil
}
func (maker *teacherMaker) Update(teacher *Teacher) error {
	return nil

}
func (maker *teacherMaker) GetByID(id int64) (*Teacher, error) {
	query, _, err := sq.
		Select("teacher.id", "person.first_name", "person.last_name", "person.username", "person.created_at", "person.last_updated").
		From("teacher").
		Join("person on teacher.person_id=person.id").
		Where(sq.Eq{"teacher.id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}
	var teacher = &Teacher{}
	err = util.GetAndMarshal(query, maker, teacher, id)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (maker *teacherMaker) Destroy(t *Teacher) error {
	return nil
}

type AssignmentMaker struct {
	*sqlx.DB
}

func (maker *AssignmentMaker) CreateAssignment(assig *Assignment) error {
	query, _, err := sq.
		Insert("assignment").Columns("student_id", "teacher_id").
		Values(assig.StudentID, assig.TeacherID).
		ToSql()
	if err != nil {
		return err
	}

	result, err := util.PrepAndExec(query, maker, assig.StudentID, assig.TeacherID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	assig.ID = id

	return nil
}

func (maker *AssignmentMaker) UpdateAssignment(assig *Assignment) error {
	return nil

}

func (maker *AssignmentMaker) GetAssignmentByID(id int) (*Assignment, error) {
	return nil, nil
}

func (maker *AssignmentMaker) DestroyAssignment(assig *Assignment) error {
	return nil
}
