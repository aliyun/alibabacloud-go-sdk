// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatabase interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Database
	GetCreatedAt() *int64
	SetCreatedBy(v string) *Database
	GetCreatedBy() *string
	SetId(v string) *Database
	GetId() *string
	SetLocation(v string) *Database
	GetLocation() *string
	SetName(v string) *Database
	GetName() *string
	SetOptions(v map[string]*string) *Database
	GetOptions() map[string]*string
	SetOwner(v string) *Database
	GetOwner() *string
	SetUpdatedAt(v int64) *Database
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *Database
	GetUpdatedBy() *string
}

type Database struct {
	CreatedAt *int64             `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy *string            `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Id        *string            `json:"id,omitempty" xml:"id,omitempty"`
	Location  *string            `json:"location,omitempty" xml:"location,omitempty"`
	Name      *string            `json:"name,omitempty" xml:"name,omitempty"`
	Options   map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
	Owner     *string            `json:"owner,omitempty" xml:"owner,omitempty"`
	UpdatedAt *int64             `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy *string            `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s Database) String() string {
	return dara.Prettify(s)
}

func (s Database) GoString() string {
	return s.String()
}

func (s *Database) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Database) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *Database) GetId() *string {
	return s.Id
}

func (s *Database) GetLocation() *string {
	return s.Location
}

func (s *Database) GetName() *string {
	return s.Name
}

func (s *Database) GetOptions() map[string]*string {
	return s.Options
}

func (s *Database) GetOwner() *string {
	return s.Owner
}

func (s *Database) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Database) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *Database) SetCreatedAt(v int64) *Database {
	s.CreatedAt = &v
	return s
}

func (s *Database) SetCreatedBy(v string) *Database {
	s.CreatedBy = &v
	return s
}

func (s *Database) SetId(v string) *Database {
	s.Id = &v
	return s
}

func (s *Database) SetLocation(v string) *Database {
	s.Location = &v
	return s
}

func (s *Database) SetName(v string) *Database {
	s.Name = &v
	return s
}

func (s *Database) SetOptions(v map[string]*string) *Database {
	s.Options = v
	return s
}

func (s *Database) SetOwner(v string) *Database {
	s.Owner = &v
	return s
}

func (s *Database) SetUpdatedAt(v int64) *Database {
	s.UpdatedAt = &v
	return s
}

func (s *Database) SetUpdatedBy(v string) *Database {
	s.UpdatedBy = &v
	return s
}

func (s *Database) Validate() error {
	return dara.Validate(s)
}
