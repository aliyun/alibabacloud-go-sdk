// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRowsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRow(v int64) *DeleteRowsShrinkRequest
	GetRow() *int64
	SetRowCount(v int64) *DeleteRowsShrinkRequest
	GetRowCount() *int64
	SetSheetId(v string) *DeleteRowsShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *DeleteRowsShrinkRequest
	GetTenantContextShrink() *string
	SetWorkbookId(v string) *DeleteRowsShrinkRequest
	GetWorkbookId() *string
}

type DeleteRowsShrinkRequest struct {
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
	// 10
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
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s DeleteRowsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteRowsShrinkRequest) GetRow() *int64 {
	return s.Row
}

func (s *DeleteRowsShrinkRequest) GetRowCount() *int64 {
	return s.RowCount
}

func (s *DeleteRowsShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *DeleteRowsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteRowsShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *DeleteRowsShrinkRequest) SetRow(v int64) *DeleteRowsShrinkRequest {
	s.Row = &v
	return s
}

func (s *DeleteRowsShrinkRequest) SetRowCount(v int64) *DeleteRowsShrinkRequest {
	s.RowCount = &v
	return s
}

func (s *DeleteRowsShrinkRequest) SetSheetId(v string) *DeleteRowsShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *DeleteRowsShrinkRequest) SetTenantContextShrink(v string) *DeleteRowsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteRowsShrinkRequest) SetWorkbookId(v string) *DeleteRowsShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *DeleteRowsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
