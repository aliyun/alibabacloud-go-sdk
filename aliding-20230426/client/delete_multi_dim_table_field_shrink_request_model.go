// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultiDimTableFieldShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *DeleteMultiDimTableFieldShrinkRequest
	GetBaseId() *string
	SetFieldIdOrName(v string) *DeleteMultiDimTableFieldShrinkRequest
	GetFieldIdOrName() *string
	SetSheetIdOrName(v string) *DeleteMultiDimTableFieldShrinkRequest
	GetSheetIdOrName() *string
	SetTenantContextShrink(v string) *DeleteMultiDimTableFieldShrinkRequest
	GetTenantContextShrink() *string
}

type DeleteMultiDimTableFieldShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7nxxxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	FieldIdOrName *string `json:"FieldIdOrName,omitempty" xml:"FieldIdOrName,omitempty"`
	// This parameter is required.
	SheetIdOrName       *string `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s DeleteMultiDimTableFieldShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultiDimTableFieldShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultiDimTableFieldShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *DeleteMultiDimTableFieldShrinkRequest) GetFieldIdOrName() *string {
	return s.FieldIdOrName
}

func (s *DeleteMultiDimTableFieldShrinkRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *DeleteMultiDimTableFieldShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteMultiDimTableFieldShrinkRequest) SetBaseId(v string) *DeleteMultiDimTableFieldShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *DeleteMultiDimTableFieldShrinkRequest) SetFieldIdOrName(v string) *DeleteMultiDimTableFieldShrinkRequest {
	s.FieldIdOrName = &v
	return s
}

func (s *DeleteMultiDimTableFieldShrinkRequest) SetSheetIdOrName(v string) *DeleteMultiDimTableFieldShrinkRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *DeleteMultiDimTableFieldShrinkRequest) SetTenantContextShrink(v string) *DeleteMultiDimTableFieldShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteMultiDimTableFieldShrinkRequest) Validate() error {
	return dara.Validate(s)
}
