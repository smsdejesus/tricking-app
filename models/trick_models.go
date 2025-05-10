package models

import "time"

// === SCHEMA: trick_data ===

type Category struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement;schema:trick_data"`
	Name string `json:"name" gorm:"unique;not null"`
}

type BaseTrick struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement;schema:trick_data"`
	Name string `json:"name" gorm:"unique;not null"`
}

type Rotation struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement;schema:trick_data"`
	Name string `json:"name" gorm:"unique;not null"`
}

type Flip struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement;schema:trick_data"`
	Name string `json:"name" gorm:"unique;not null"`
}

type Stance struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement;schema:trick_data"`
	Name string `json:"name" gorm:"unique;not null"`
}

type Variation struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement;schema:trick_data"`
	Name string `json:"name" gorm:"unique;not null"`
}

// === SCHEMA: auth ===

type User struct {
	ID       string `json:"id" gorm:"primaryKey;schema:auth"`
	Username string `json:"username" gorm:"unique;not null"`
}

// === SCHEMA: trick_data ===

type Trick struct {
	ID             int       `json:"id" gorm:"primaryKey;autoIncrement;schema:trick_data"`
	Name           string    `json:"name" gorm:"not null"`
	Description    *string   `json:"description"`
	Difficulty     int       `json:"difficulty"`
	VideoURL       *string   `json:"video_url"`
	ThumbnailURL   *string   `json:"thumbnail_url"`
	Tricker        *string   `json:"tricker"`
	ExecutionNotes *string   `json:"execution_notes"`
	CreatedBy      *string   `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	// Associations
	Categories    []Category `gorm:"many2many:trick_data.trick_categories"`
	BaseTrick     *BaseTrick
	Rotation      *Rotation
	Flip          *Flip
	TakeoffStance *Stance
	LandingStance *Stance
	Variations    []Variation `gorm:"many2many:trick_data.trick_variations"`
}

type Prerequisite struct {
	TrickID        int `json:"trick_id" gorm:"schema:trick_data"`
	PrerequisiteID int `json:"prerequisite_id" gorm:"schema:trick_data"`
}
