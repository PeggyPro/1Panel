package model

type Clam struct {
	BaseModel

	Name        string `gorm:"type:varchar(64);not null" json:"name"`
	Path        string `gorm:"type:varchar(64);not null" json:"path"`
	Description string `gorm:"type:varchar(64);not null" json:"description"`
}
