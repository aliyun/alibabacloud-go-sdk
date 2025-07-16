// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteColumnsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumn(v int64) *DeleteColumnsShrinkRequest
	GetColumn() *int64
	SetColumnCount(v int64) *DeleteColumnsShrinkRequest
	GetColumnCount() *int64
	SetSheetId(v string) *DeleteColumnsShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *DeleteColumnsShrinkRequest
	GetTenantContextShrink() *string
	SetWorkbookId(v string) *DeleteColumnsShrinkRequest
	GetWorkbookId() *string
}

type DeleteColumnsShrinkRequest struct {
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
	// 10
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

func (s DeleteColumnsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteColumnsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteColumnsShrinkRequest) GetColumn() *int64 {
	return s.Column
}

func (s *DeleteColumnsShrinkRequest) GetColumnCount() *int64 {
	return s.ColumnCount
}

func (s *DeleteColumnsShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *DeleteColumnsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteColumnsShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *DeleteColumnsShrinkRequest) SetColumn(v int64) *DeleteColumnsShrinkRequest {
	s.Column = &v
	return s
}

func (s *DeleteColumnsShrinkRequest) SetColumnCount(v int64) *DeleteColumnsShrinkRequest {
	s.ColumnCount = &v
	return s
}

func (s *DeleteColumnsShrinkRequest) SetSheetId(v string) *DeleteColumnsShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *DeleteColumnsShrinkRequest) SetTenantContextShrink(v string) *DeleteColumnsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteColumnsShrinkRequest) SetWorkbookId(v string) *DeleteColumnsShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *DeleteColumnsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
