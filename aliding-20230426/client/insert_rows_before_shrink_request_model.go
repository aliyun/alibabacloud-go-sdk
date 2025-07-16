// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertRowsBeforeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRow(v int64) *InsertRowsBeforeShrinkRequest
	GetRow() *int64
	SetRowCount(v int64) *InsertRowsBeforeShrinkRequest
	GetRowCount() *int64
	SetSheetId(v string) *InsertRowsBeforeShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *InsertRowsBeforeShrinkRequest
	GetTenantContextShrink() *string
	SetWorkbookId(v string) *InsertRowsBeforeShrinkRequest
	GetWorkbookId() *string
}

type InsertRowsBeforeShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
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

func (s InsertRowsBeforeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertRowsBeforeShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeShrinkRequest) GetRow() *int64 {
	return s.Row
}

func (s *InsertRowsBeforeShrinkRequest) GetRowCount() *int64 {
	return s.RowCount
}

func (s *InsertRowsBeforeShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *InsertRowsBeforeShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *InsertRowsBeforeShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *InsertRowsBeforeShrinkRequest) SetRow(v int64) *InsertRowsBeforeShrinkRequest {
	s.Row = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) SetRowCount(v int64) *InsertRowsBeforeShrinkRequest {
	s.RowCount = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) SetSheetId(v string) *InsertRowsBeforeShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) SetTenantContextShrink(v string) *InsertRowsBeforeShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) SetWorkbookId(v string) *InsertRowsBeforeShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *InsertRowsBeforeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
