// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetColumnsVisibilityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumn(v int64) *SetColumnsVisibilityShrinkRequest
	GetColumn() *int64
	SetColumnCount(v int64) *SetColumnsVisibilityShrinkRequest
	GetColumnCount() *int64
	SetSheetId(v string) *SetColumnsVisibilityShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *SetColumnsVisibilityShrinkRequest
	GetTenantContextShrink() *string
	SetVisibility(v string) *SetColumnsVisibilityShrinkRequest
	GetVisibility() *string
	SetWorkbookId(v string) *SetColumnsVisibilityShrinkRequest
	GetWorkbookId() *string
}

type SetColumnsVisibilityShrinkRequest struct {
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
	// 20
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
	// hidden
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s SetColumnsVisibilityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetColumnsVisibilityShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetColumnsVisibilityShrinkRequest) GetColumn() *int64 {
	return s.Column
}

func (s *SetColumnsVisibilityShrinkRequest) GetColumnCount() *int64 {
	return s.ColumnCount
}

func (s *SetColumnsVisibilityShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *SetColumnsVisibilityShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *SetColumnsVisibilityShrinkRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *SetColumnsVisibilityShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *SetColumnsVisibilityShrinkRequest) SetColumn(v int64) *SetColumnsVisibilityShrinkRequest {
	s.Column = &v
	return s
}

func (s *SetColumnsVisibilityShrinkRequest) SetColumnCount(v int64) *SetColumnsVisibilityShrinkRequest {
	s.ColumnCount = &v
	return s
}

func (s *SetColumnsVisibilityShrinkRequest) SetSheetId(v string) *SetColumnsVisibilityShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *SetColumnsVisibilityShrinkRequest) SetTenantContextShrink(v string) *SetColumnsVisibilityShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *SetColumnsVisibilityShrinkRequest) SetVisibility(v string) *SetColumnsVisibilityShrinkRequest {
	s.Visibility = &v
	return s
}

func (s *SetColumnsVisibilityShrinkRequest) SetWorkbookId(v string) *SetColumnsVisibilityShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *SetColumnsVisibilityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
