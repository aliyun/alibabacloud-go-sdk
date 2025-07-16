// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableSheetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *GetMultiDimTableSheetShrinkRequest
	GetBaseId() *string
	SetSheetIdOrName(v string) *GetMultiDimTableSheetShrinkRequest
	GetSheetIdOrName() *string
	SetTenantContextShrink(v string) *GetMultiDimTableSheetShrinkRequest
	GetTenantContextShrink() *string
}

type GetMultiDimTableSheetShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 169899
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	SheetIdOrName       *string `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetMultiDimTableSheetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableSheetShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableSheetShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *GetMultiDimTableSheetShrinkRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *GetMultiDimTableSheetShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetMultiDimTableSheetShrinkRequest) SetBaseId(v string) *GetMultiDimTableSheetShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *GetMultiDimTableSheetShrinkRequest) SetSheetIdOrName(v string) *GetMultiDimTableSheetShrinkRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *GetMultiDimTableSheetShrinkRequest) SetTenantContextShrink(v string) *GetMultiDimTableSheetShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetMultiDimTableSheetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
