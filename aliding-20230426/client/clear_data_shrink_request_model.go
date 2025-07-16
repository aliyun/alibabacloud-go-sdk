// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRangeAddress(v string) *ClearDataShrinkRequest
	GetRangeAddress() *string
	SetSheetId(v string) *ClearDataShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *ClearDataShrinkRequest
	GetTenantContextShrink() *string
	SetWorkbookId(v string) *ClearDataShrinkRequest
	GetWorkbookId() *string
}

type ClearDataShrinkRequest struct {
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

func (s ClearDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *ClearDataShrinkRequest) GetRangeAddress() *string {
	return s.RangeAddress
}

func (s *ClearDataShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *ClearDataShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ClearDataShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *ClearDataShrinkRequest) SetRangeAddress(v string) *ClearDataShrinkRequest {
	s.RangeAddress = &v
	return s
}

func (s *ClearDataShrinkRequest) SetSheetId(v string) *ClearDataShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *ClearDataShrinkRequest) SetTenantContextShrink(v string) *ClearDataShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ClearDataShrinkRequest) SetWorkbookId(v string) *ClearDataShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *ClearDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
