// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableFieldShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *UpdateMultiDimTableFieldShrinkRequest
	GetBaseId() *string
	SetFieldIdOrName(v string) *UpdateMultiDimTableFieldShrinkRequest
	GetFieldIdOrName() *string
	SetName(v string) *UpdateMultiDimTableFieldShrinkRequest
	GetName() *string
	SetPropertyShrink(v string) *UpdateMultiDimTableFieldShrinkRequest
	GetPropertyShrink() *string
	SetSheetIdOrName(v string) *UpdateMultiDimTableFieldShrinkRequest
	GetSheetIdOrName() *string
	SetTenantContextShrink(v string) *UpdateMultiDimTableFieldShrinkRequest
	GetTenantContextShrink() *string
}

type UpdateMultiDimTableFieldShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7noNyJxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	FieldIdOrName *string `json:"FieldIdOrName,omitempty" xml:"FieldIdOrName,omitempty"`
	// This parameter is required.
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PropertyShrink *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// This parameter is required.
	SheetIdOrName       *string `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s UpdateMultiDimTableFieldShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableFieldShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableFieldShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *UpdateMultiDimTableFieldShrinkRequest) GetFieldIdOrName() *string {
	return s.FieldIdOrName
}

func (s *UpdateMultiDimTableFieldShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMultiDimTableFieldShrinkRequest) GetPropertyShrink() *string {
	return s.PropertyShrink
}

func (s *UpdateMultiDimTableFieldShrinkRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *UpdateMultiDimTableFieldShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateMultiDimTableFieldShrinkRequest) SetBaseId(v string) *UpdateMultiDimTableFieldShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *UpdateMultiDimTableFieldShrinkRequest) SetFieldIdOrName(v string) *UpdateMultiDimTableFieldShrinkRequest {
	s.FieldIdOrName = &v
	return s
}

func (s *UpdateMultiDimTableFieldShrinkRequest) SetName(v string) *UpdateMultiDimTableFieldShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateMultiDimTableFieldShrinkRequest) SetPropertyShrink(v string) *UpdateMultiDimTableFieldShrinkRequest {
	s.PropertyShrink = &v
	return s
}

func (s *UpdateMultiDimTableFieldShrinkRequest) SetSheetIdOrName(v string) *UpdateMultiDimTableFieldShrinkRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *UpdateMultiDimTableFieldShrinkRequest) SetTenantContextShrink(v string) *UpdateMultiDimTableFieldShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateMultiDimTableFieldShrinkRequest) Validate() error {
	return dara.Validate(s)
}
