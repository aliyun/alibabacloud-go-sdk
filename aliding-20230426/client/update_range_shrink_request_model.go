// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRangeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackgroundColorsShrink(v string) *UpdateRangeShrinkRequest
	GetBackgroundColorsShrink() *string
	SetHyperlinksShrink(v string) *UpdateRangeShrinkRequest
	GetHyperlinksShrink() *string
	SetNumberFormat(v string) *UpdateRangeShrinkRequest
	GetNumberFormat() *string
	SetRangeAddress(v string) *UpdateRangeShrinkRequest
	GetRangeAddress() *string
	SetSheetId(v string) *UpdateRangeShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *UpdateRangeShrinkRequest
	GetTenantContextShrink() *string
	SetValuesShrink(v string) *UpdateRangeShrinkRequest
	GetValuesShrink() *string
	SetWorkbookId(v string) *UpdateRangeShrinkRequest
	GetWorkbookId() *string
}

type UpdateRangeShrinkRequest struct {
	// example:
	//
	// [["#ff0000","#ff0000","#ff0000"]]
	BackgroundColorsShrink *string `json:"BackgroundColors,omitempty" xml:"BackgroundColors,omitempty"`
	// example:
	//
	// [["type": "path","link": "https://www.dingtalk.com","text": "test"]]
	HyperlinksShrink *string `json:"Hyperlinks,omitempty" xml:"Hyperlinks,omitempty"`
	// example:
	//
	// General
	NumberFormat *string `json:"NumberFormat,omitempty" xml:"NumberFormat,omitempty"`
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
	// example:
	//
	// [["1","2","3"]]
	ValuesShrink *string `json:"Values,omitempty" xml:"Values,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s UpdateRangeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRangeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRangeShrinkRequest) GetBackgroundColorsShrink() *string {
	return s.BackgroundColorsShrink
}

func (s *UpdateRangeShrinkRequest) GetHyperlinksShrink() *string {
	return s.HyperlinksShrink
}

func (s *UpdateRangeShrinkRequest) GetNumberFormat() *string {
	return s.NumberFormat
}

func (s *UpdateRangeShrinkRequest) GetRangeAddress() *string {
	return s.RangeAddress
}

func (s *UpdateRangeShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *UpdateRangeShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateRangeShrinkRequest) GetValuesShrink() *string {
	return s.ValuesShrink
}

func (s *UpdateRangeShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *UpdateRangeShrinkRequest) SetBackgroundColorsShrink(v string) *UpdateRangeShrinkRequest {
	s.BackgroundColorsShrink = &v
	return s
}

func (s *UpdateRangeShrinkRequest) SetHyperlinksShrink(v string) *UpdateRangeShrinkRequest {
	s.HyperlinksShrink = &v
	return s
}

func (s *UpdateRangeShrinkRequest) SetNumberFormat(v string) *UpdateRangeShrinkRequest {
	s.NumberFormat = &v
	return s
}

func (s *UpdateRangeShrinkRequest) SetRangeAddress(v string) *UpdateRangeShrinkRequest {
	s.RangeAddress = &v
	return s
}

func (s *UpdateRangeShrinkRequest) SetSheetId(v string) *UpdateRangeShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *UpdateRangeShrinkRequest) SetTenantContextShrink(v string) *UpdateRangeShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateRangeShrinkRequest) SetValuesShrink(v string) *UpdateRangeShrinkRequest {
	s.ValuesShrink = &v
	return s
}

func (s *UpdateRangeShrinkRequest) SetWorkbookId(v string) *UpdateRangeShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *UpdateRangeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
