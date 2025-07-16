// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRowsVisibilityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRow(v int64) *SetRowsVisibilityShrinkRequest
	GetRow() *int64
	SetRowCount(v int64) *SetRowsVisibilityShrinkRequest
	GetRowCount() *int64
	SetSheetId(v string) *SetRowsVisibilityShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *SetRowsVisibilityShrinkRequest
	GetTenantContextShrink() *string
	SetVisibility(v string) *SetRowsVisibilityShrinkRequest
	GetVisibility() *string
	SetWorkbookId(v string) *SetRowsVisibilityShrinkRequest
	GetWorkbookId() *string
}

type SetRowsVisibilityShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Row *int64 `json:"Row,omitempty" xml:"Row,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
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
	// hidden
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s SetRowsVisibilityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRowsVisibilityShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetRowsVisibilityShrinkRequest) GetRow() *int64 {
	return s.Row
}

func (s *SetRowsVisibilityShrinkRequest) GetRowCount() *int64 {
	return s.RowCount
}

func (s *SetRowsVisibilityShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *SetRowsVisibilityShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *SetRowsVisibilityShrinkRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *SetRowsVisibilityShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *SetRowsVisibilityShrinkRequest) SetRow(v int64) *SetRowsVisibilityShrinkRequest {
	s.Row = &v
	return s
}

func (s *SetRowsVisibilityShrinkRequest) SetRowCount(v int64) *SetRowsVisibilityShrinkRequest {
	s.RowCount = &v
	return s
}

func (s *SetRowsVisibilityShrinkRequest) SetSheetId(v string) *SetRowsVisibilityShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *SetRowsVisibilityShrinkRequest) SetTenantContextShrink(v string) *SetRowsVisibilityShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *SetRowsVisibilityShrinkRequest) SetVisibility(v string) *SetRowsVisibilityShrinkRequest {
	s.Visibility = &v
	return s
}

func (s *SetRowsVisibilityShrinkRequest) SetWorkbookId(v string) *SetRowsVisibilityShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *SetRowsVisibilityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
