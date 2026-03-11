// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iView interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *View
	GetCreatedAt() *int64
	SetCreatedBy(v string) *View
	GetCreatedBy() *string
	SetId(v string) *View
	GetId() *string
	SetName(v string) *View
	GetName() *string
	SetOwner(v string) *View
	GetOwner() *string
	SetSchema(v *ViewSchema) *View
	GetSchema() *ViewSchema
	SetUpdatedAt(v int64) *View
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *View
	GetUpdatedBy() *string
}

type View struct {
	// example:
	//
	// 1744970111419
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// acs:ram::[accountId]:root
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// 1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// view_test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// acs:ram::[accountId]:root
	Owner  *string     `json:"owner,omitempty" xml:"owner,omitempty"`
	Schema *ViewSchema `json:"schema,omitempty" xml:"schema,omitempty"`
	// example:
	//
	// 1744970111419
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// acs:ram::[accountId]:root
	UpdatedBy *string `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s View) String() string {
	return dara.Prettify(s)
}

func (s View) GoString() string {
	return s.String()
}

func (s *View) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *View) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *View) GetId() *string {
	return s.Id
}

func (s *View) GetName() *string {
	return s.Name
}

func (s *View) GetOwner() *string {
	return s.Owner
}

func (s *View) GetSchema() *ViewSchema {
	return s.Schema
}

func (s *View) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *View) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *View) SetCreatedAt(v int64) *View {
	s.CreatedAt = &v
	return s
}

func (s *View) SetCreatedBy(v string) *View {
	s.CreatedBy = &v
	return s
}

func (s *View) SetId(v string) *View {
	s.Id = &v
	return s
}

func (s *View) SetName(v string) *View {
	s.Name = &v
	return s
}

func (s *View) SetOwner(v string) *View {
	s.Owner = &v
	return s
}

func (s *View) SetSchema(v *ViewSchema) *View {
	s.Schema = v
	return s
}

func (s *View) SetUpdatedAt(v int64) *View {
	s.UpdatedAt = &v
	return s
}

func (s *View) SetUpdatedBy(v string) *View {
	s.UpdatedBy = &v
	return s
}

func (s *View) Validate() error {
	if s.Schema != nil {
		if err := s.Schema.Validate(); err != nil {
			return err
		}
	}
	return nil
}
