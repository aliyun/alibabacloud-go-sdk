// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCatalog interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Catalog
	GetCreatedAt() *int64
	SetCreatedBy(v string) *Catalog
	GetCreatedBy() *string
	SetId(v string) *Catalog
	GetId() *string
	SetIsShared(v bool) *Catalog
	GetIsShared() *bool
	SetName(v string) *Catalog
	GetName() *string
	SetOptions(v map[string]*string) *Catalog
	GetOptions() map[string]*string
	SetOwner(v string) *Catalog
	GetOwner() *string
	SetShareId(v string) *Catalog
	GetShareId() *string
	SetStatus(v string) *Catalog
	GetStatus() *string
	SetType(v string) *Catalog
	GetType() *string
	SetUpdatedAt(v int64) *Catalog
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *Catalog
	GetUpdatedBy() *string
}

type Catalog struct {
	CreatedAt *int64             `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy *string            `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Id        *string            `json:"id,omitempty" xml:"id,omitempty"`
	IsShared  *bool              `json:"isShared,omitempty" xml:"isShared,omitempty"`
	Name      *string            `json:"name,omitempty" xml:"name,omitempty"`
	Options   map[string]*string `json:"options,omitempty" xml:"options,omitempty"`
	Owner     *string            `json:"owner,omitempty" xml:"owner,omitempty"`
	ShareId   *string            `json:"shareId,omitempty" xml:"shareId,omitempty"`
	Status    *string            `json:"status,omitempty" xml:"status,omitempty"`
	Type      *string            `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt *int64             `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy *string            `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s Catalog) String() string {
	return dara.Prettify(s)
}

func (s Catalog) GoString() string {
	return s.String()
}

func (s *Catalog) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Catalog) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *Catalog) GetId() *string {
	return s.Id
}

func (s *Catalog) GetIsShared() *bool {
	return s.IsShared
}

func (s *Catalog) GetName() *string {
	return s.Name
}

func (s *Catalog) GetOptions() map[string]*string {
	return s.Options
}

func (s *Catalog) GetOwner() *string {
	return s.Owner
}

func (s *Catalog) GetShareId() *string {
	return s.ShareId
}

func (s *Catalog) GetStatus() *string {
	return s.Status
}

func (s *Catalog) GetType() *string {
	return s.Type
}

func (s *Catalog) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Catalog) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *Catalog) SetCreatedAt(v int64) *Catalog {
	s.CreatedAt = &v
	return s
}

func (s *Catalog) SetCreatedBy(v string) *Catalog {
	s.CreatedBy = &v
	return s
}

func (s *Catalog) SetId(v string) *Catalog {
	s.Id = &v
	return s
}

func (s *Catalog) SetIsShared(v bool) *Catalog {
	s.IsShared = &v
	return s
}

func (s *Catalog) SetName(v string) *Catalog {
	s.Name = &v
	return s
}

func (s *Catalog) SetOptions(v map[string]*string) *Catalog {
	s.Options = v
	return s
}

func (s *Catalog) SetOwner(v string) *Catalog {
	s.Owner = &v
	return s
}

func (s *Catalog) SetShareId(v string) *Catalog {
	s.ShareId = &v
	return s
}

func (s *Catalog) SetStatus(v string) *Catalog {
	s.Status = &v
	return s
}

func (s *Catalog) SetType(v string) *Catalog {
	s.Type = &v
	return s
}

func (s *Catalog) SetUpdatedAt(v int64) *Catalog {
	s.UpdatedAt = &v
	return s
}

func (s *Catalog) SetUpdatedBy(v string) *Catalog {
	s.UpdatedBy = &v
	return s
}

func (s *Catalog) Validate() error {
	return dara.Validate(s)
}
