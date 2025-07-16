// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiDimTableFieldShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *CreateMultiDimTableFieldShrinkRequest
	GetBaseId() *string
	SetName(v string) *CreateMultiDimTableFieldShrinkRequest
	GetName() *string
	SetPropertyShrink(v string) *CreateMultiDimTableFieldShrinkRequest
	GetPropertyShrink() *string
	SetSheetIdOrName(v string) *CreateMultiDimTableFieldShrinkRequest
	GetSheetIdOrName() *string
	SetTenantContextShrink(v string) *CreateMultiDimTableFieldShrinkRequest
	GetTenantContextShrink() *string
	SetType(v string) *CreateMultiDimTableFieldShrinkRequest
	GetType() *string
}

type CreateMultiDimTableFieldShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7noNyJxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PropertyShrink *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// This parameter is required.
	SheetIdOrName       *string `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateMultiDimTableFieldShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiDimTableFieldShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiDimTableFieldShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *CreateMultiDimTableFieldShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateMultiDimTableFieldShrinkRequest) GetPropertyShrink() *string {
	return s.PropertyShrink
}

func (s *CreateMultiDimTableFieldShrinkRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *CreateMultiDimTableFieldShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateMultiDimTableFieldShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateMultiDimTableFieldShrinkRequest) SetBaseId(v string) *CreateMultiDimTableFieldShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *CreateMultiDimTableFieldShrinkRequest) SetName(v string) *CreateMultiDimTableFieldShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateMultiDimTableFieldShrinkRequest) SetPropertyShrink(v string) *CreateMultiDimTableFieldShrinkRequest {
	s.PropertyShrink = &v
	return s
}

func (s *CreateMultiDimTableFieldShrinkRequest) SetSheetIdOrName(v string) *CreateMultiDimTableFieldShrinkRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *CreateMultiDimTableFieldShrinkRequest) SetTenantContextShrink(v string) *CreateMultiDimTableFieldShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateMultiDimTableFieldShrinkRequest) SetType(v string) *CreateMultiDimTableFieldShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateMultiDimTableFieldShrinkRequest) Validate() error {
	return dara.Validate(s)
}
