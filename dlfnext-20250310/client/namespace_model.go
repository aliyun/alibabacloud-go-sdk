// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNamespace interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Namespace
	GetCreatedAt() *int64
	SetCreatedBy(v string) *Namespace
	GetCreatedBy() *string
	SetId(v string) *Namespace
	GetId() *string
	SetLocation(v string) *Namespace
	GetLocation() *string
	SetName(v string) *Namespace
	GetName() *string
	SetOptions(v map[string]*string) *Namespace
	GetOptions() map[string]*string
	SetOwner(v string) *Namespace
	GetOwner() *string
	SetUpdatedAt(v int64) *Namespace
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *Namespace
	GetUpdatedBy() *string
}

type Namespace struct {
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

func (s Namespace) String() string {
	return dara.Prettify(s)
}

func (s Namespace) GoString() string {
	return s.String()
}

func (s *Namespace) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Namespace) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *Namespace) GetId() *string {
	return s.Id
}

func (s *Namespace) GetLocation() *string {
	return s.Location
}

func (s *Namespace) GetName() *string {
	return s.Name
}

func (s *Namespace) GetOptions() map[string]*string {
	return s.Options
}

func (s *Namespace) GetOwner() *string {
	return s.Owner
}

func (s *Namespace) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Namespace) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *Namespace) SetCreatedAt(v int64) *Namespace {
	s.CreatedAt = &v
	return s
}

func (s *Namespace) SetCreatedBy(v string) *Namespace {
	s.CreatedBy = &v
	return s
}

func (s *Namespace) SetId(v string) *Namespace {
	s.Id = &v
	return s
}

func (s *Namespace) SetLocation(v string) *Namespace {
	s.Location = &v
	return s
}

func (s *Namespace) SetName(v string) *Namespace {
	s.Name = &v
	return s
}

func (s *Namespace) SetOptions(v map[string]*string) *Namespace {
	s.Options = v
	return s
}

func (s *Namespace) SetOwner(v string) *Namespace {
	s.Owner = &v
	return s
}

func (s *Namespace) SetUpdatedAt(v int64) *Namespace {
	s.UpdatedAt = &v
	return s
}

func (s *Namespace) SetUpdatedBy(v string) *Namespace {
	s.UpdatedBy = &v
	return s
}

func (s *Namespace) Validate() error {
	return dara.Validate(s)
}
