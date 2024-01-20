package models

type Student struct {
	ID        string     `json:"id" validate:"required,min=1,max=36"`
	Name      string     `json:"name" validate:"required,min=1,max=50"`
	Age       uint8      `json:"age" validate:"required,min=1,max=80"`
	Grade     uint8      `json:"grade" validate:"required,min=1,max=11"`
	Guardians []Guardian `json:"guardians" validate:"required,min=1,dive"`
	Hobbies   []string   `json:"hobbies" validate:"omitempty,dive,min=2,max=50"`
}

type Guardian struct {
	ID           string `json:"id" validate:"required,min=1,max=36"`
	Name         string `json:"name" validate:"required,min=1,max=50"`
	Relationship string `json:"relationship" validate:"required,oneof='parent' 'grandparent' 'brother' 'sister'"`
	Phone        string `json:"phone" validate:"required,is-phone"`
}
