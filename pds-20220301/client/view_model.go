// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iView interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *View
	GetCategory() *string
	SetCreatedAt(v string) *View
	GetCreatedAt() *string
	SetDescription(v string) *View
	GetDescription() *string
	SetExFieldsInfo(v map[string]interface{}) *View
	GetExFieldsInfo() map[string]interface{}
	SetFileCount(v int64) *View
	GetFileCount() *int64
	SetName(v string) *View
	GetName() *string
	SetOwner(v string) *View
	GetOwner() *string
	SetUpdatedAt(v string) *View
	GetUpdatedAt() *string
	SetViewId(v string) *View
	GetViewId() *string
}

type View struct {
	Category     *string                `json:"category,omitempty" xml:"category,omitempty"`
	CreatedAt    *string                `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description  *string                `json:"description,omitempty" xml:"description,omitempty"`
	ExFieldsInfo map[string]interface{} `json:"ex_fields_info,omitempty" xml:"ex_fields_info,omitempty"`
	FileCount    *int64                 `json:"file_count,omitempty" xml:"file_count,omitempty"`
	Name         *string                `json:"name,omitempty" xml:"name,omitempty"`
	Owner        *string                `json:"owner,omitempty" xml:"owner,omitempty"`
	UpdatedAt    *string                `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	ViewId       *string                `json:"view_id,omitempty" xml:"view_id,omitempty"`
}

func (s View) String() string {
	return dara.Prettify(s)
}

func (s View) GoString() string {
	return s.String()
}

func (s *View) GetCategory() *string {
	return s.Category
}

func (s *View) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *View) GetDescription() *string {
	return s.Description
}

func (s *View) GetExFieldsInfo() map[string]interface{} {
	return s.ExFieldsInfo
}

func (s *View) GetFileCount() *int64 {
	return s.FileCount
}

func (s *View) GetName() *string {
	return s.Name
}

func (s *View) GetOwner() *string {
	return s.Owner
}

func (s *View) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *View) GetViewId() *string {
	return s.ViewId
}

func (s *View) SetCategory(v string) *View {
	s.Category = &v
	return s
}

func (s *View) SetCreatedAt(v string) *View {
	s.CreatedAt = &v
	return s
}

func (s *View) SetDescription(v string) *View {
	s.Description = &v
	return s
}

func (s *View) SetExFieldsInfo(v map[string]interface{}) *View {
	s.ExFieldsInfo = v
	return s
}

func (s *View) SetFileCount(v int64) *View {
	s.FileCount = &v
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

func (s *View) SetUpdatedAt(v string) *View {
	s.UpdatedAt = &v
	return s
}

func (s *View) SetViewId(v string) *View {
	s.ViewId = &v
	return s
}

func (s *View) Validate() error {
	return dara.Validate(s)
}
