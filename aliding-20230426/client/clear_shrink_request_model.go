// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRangeAddress(v string) *ClearShrinkRequest
	GetRangeAddress() *string
	SetSheetId(v string) *ClearShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *ClearShrinkRequest
	GetTenantContextShrink() *string
	SetWorkbookId(v string) *ClearShrinkRequest
	GetWorkbookId() *string
}

type ClearShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// A3:C3
	RangeAddress *string `json:"RangeAddress,omitempty" xml:"RangeAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	SheetId             *string `json:"SheetId,omitempty" xml:"SheetId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s ClearShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearShrinkRequest) GoString() string {
	return s.String()
}

func (s *ClearShrinkRequest) GetRangeAddress() *string {
	return s.RangeAddress
}

func (s *ClearShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *ClearShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ClearShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *ClearShrinkRequest) SetRangeAddress(v string) *ClearShrinkRequest {
	s.RangeAddress = &v
	return s
}

func (s *ClearShrinkRequest) SetSheetId(v string) *ClearShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *ClearShrinkRequest) SetTenantContextShrink(v string) *ClearShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ClearShrinkRequest) SetWorkbookId(v string) *ClearShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *ClearShrinkRequest) Validate() error {
	return dara.Validate(s)
}
