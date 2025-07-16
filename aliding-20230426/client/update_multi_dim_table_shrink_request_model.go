// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *UpdateMultiDimTableShrinkRequest
	GetBaseId() *string
	SetName(v string) *UpdateMultiDimTableShrinkRequest
	GetName() *string
	SetSheetIdOrName(v string) *UpdateMultiDimTableShrinkRequest
	GetSheetIdOrName() *string
	SetTenantContextShrink(v string) *UpdateMultiDimTableShrinkRequest
	GetTenantContextShrink() *string
}

type UpdateMultiDimTableShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7noNyJxxxxxx
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	SheetIdOrName       *string `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s UpdateMultiDimTableShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *UpdateMultiDimTableShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMultiDimTableShrinkRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *UpdateMultiDimTableShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateMultiDimTableShrinkRequest) SetBaseId(v string) *UpdateMultiDimTableShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *UpdateMultiDimTableShrinkRequest) SetName(v string) *UpdateMultiDimTableShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateMultiDimTableShrinkRequest) SetSheetIdOrName(v string) *UpdateMultiDimTableShrinkRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *UpdateMultiDimTableShrinkRequest) SetTenantContextShrink(v string) *UpdateMultiDimTableShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateMultiDimTableShrinkRequest) Validate() error {
	return dara.Validate(s)
}
