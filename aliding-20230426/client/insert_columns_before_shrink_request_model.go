// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertColumnsBeforeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumn(v int64) *InsertColumnsBeforeShrinkRequest
	GetColumn() *int64
	SetColumnCount(v int64) *InsertColumnsBeforeShrinkRequest
	GetColumnCount() *int64
	SetSheetId(v string) *InsertColumnsBeforeShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *InsertColumnsBeforeShrinkRequest
	GetTenantContextShrink() *string
	SetWorkbookId(v string) *InsertColumnsBeforeShrinkRequest
	GetWorkbookId() *string
}

type InsertColumnsBeforeShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Column *int64 `json:"Column,omitempty" xml:"Column,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	ColumnCount *int64 `json:"ColumnCount,omitempty" xml:"ColumnCount,omitempty"`
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

func (s InsertColumnsBeforeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertColumnsBeforeShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertColumnsBeforeShrinkRequest) GetColumn() *int64 {
	return s.Column
}

func (s *InsertColumnsBeforeShrinkRequest) GetColumnCount() *int64 {
	return s.ColumnCount
}

func (s *InsertColumnsBeforeShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *InsertColumnsBeforeShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *InsertColumnsBeforeShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *InsertColumnsBeforeShrinkRequest) SetColumn(v int64) *InsertColumnsBeforeShrinkRequest {
	s.Column = &v
	return s
}

func (s *InsertColumnsBeforeShrinkRequest) SetColumnCount(v int64) *InsertColumnsBeforeShrinkRequest {
	s.ColumnCount = &v
	return s
}

func (s *InsertColumnsBeforeShrinkRequest) SetSheetId(v string) *InsertColumnsBeforeShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *InsertColumnsBeforeShrinkRequest) SetTenantContextShrink(v string) *InsertColumnsBeforeShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *InsertColumnsBeforeShrinkRequest) SetWorkbookId(v string) *InsertColumnsBeforeShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *InsertColumnsBeforeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
