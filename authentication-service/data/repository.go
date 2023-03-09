package data

type Repository interface {
	GetAll() ([]*User, error)
	GetByEmail(email string) (*User, error)
	GetOne(ID int) (*User, error)
	Update(user User) error
	DeleteByID(ID int) error
	Insert(user User) (int, error)
	ResetPassword(user User, password string) error
	PasswordMatches(user User, plaintext string) (bool, error)
}
