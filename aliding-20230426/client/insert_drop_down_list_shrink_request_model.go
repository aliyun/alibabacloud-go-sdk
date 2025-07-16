// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertDropDownListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOptionsShrink(v string) *InsertDropDownListShrinkRequest
	GetOptionsShrink() *string
	SetRangeAddress(v string) *InsertDropDownListShrinkRequest
	GetRangeAddress() *string
	SetSheetId(v string) *InsertDropDownListShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *InsertDropDownListShrinkRequest
	GetTenantContextShrink() *string
	SetWorkbookId(v string) *InsertDropDownListShrinkRequest
	GetWorkbookId() *string
}

type InsertDropDownListShrinkRequest struct {
	// This parameter is required.
	OptionsShrink *string `json:"Options,omitempty" xml:"Options,omitempty"`
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

func (s InsertDropDownListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertDropDownListShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertDropDownListShrinkRequest) GetOptionsShrink() *string {
	return s.OptionsShrink
}

func (s *InsertDropDownListShrinkRequest) GetRangeAddress() *string {
	return s.RangeAddress
}

func (s *InsertDropDownListShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *InsertDropDownListShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *InsertDropDownListShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *InsertDropDownListShrinkRequest) SetOptionsShrink(v string) *InsertDropDownListShrinkRequest {
	s.OptionsShrink = &v
	return s
}

func (s *InsertDropDownListShrinkRequest) SetRangeAddress(v string) *InsertDropDownListShrinkRequest {
	s.RangeAddress = &v
	return s
}

func (s *InsertDropDownListShrinkRequest) SetSheetId(v string) *InsertDropDownListShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *InsertDropDownListShrinkRequest) SetTenantContextShrink(v string) *InsertDropDownListShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *InsertDropDownListShrinkRequest) SetWorkbookId(v string) *InsertDropDownListShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *InsertDropDownListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
