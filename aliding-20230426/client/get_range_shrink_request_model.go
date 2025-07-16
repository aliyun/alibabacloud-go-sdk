// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRangeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRangeAddress(v string) *GetRangeShrinkRequest
	GetRangeAddress() *string
	SetSelect(v string) *GetRangeShrinkRequest
	GetSelect() *string
	SetSheetId(v string) *GetRangeShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *GetRangeShrinkRequest
	GetTenantContextShrink() *string
	SetWorkbookId(v string) *GetRangeShrinkRequest
	GetWorkbookId() *string
}

type GetRangeShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// A3:C3
	RangeAddress *string `json:"RangeAddress,omitempty" xml:"RangeAddress,omitempty"`
	// example:
	//
	// values
	Select *string `json:"Select,omitempty" xml:"Select,omitempty"`
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

func (s GetRangeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRangeShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetRangeShrinkRequest) GetRangeAddress() *string {
	return s.RangeAddress
}

func (s *GetRangeShrinkRequest) GetSelect() *string {
	return s.Select
}

func (s *GetRangeShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *GetRangeShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetRangeShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *GetRangeShrinkRequest) SetRangeAddress(v string) *GetRangeShrinkRequest {
	s.RangeAddress = &v
	return s
}

func (s *GetRangeShrinkRequest) SetSelect(v string) *GetRangeShrinkRequest {
	s.Select = &v
	return s
}

func (s *GetRangeShrinkRequest) SetSheetId(v string) *GetRangeShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *GetRangeShrinkRequest) SetTenantContextShrink(v string) *GetRangeShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetRangeShrinkRequest) SetWorkbookId(v string) *GetRangeShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *GetRangeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
