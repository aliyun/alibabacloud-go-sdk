// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSheetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSheetId(v string) *DeleteSheetShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *DeleteSheetShrinkRequest
	GetTenantContextShrink() *string
	SetWorkbookId(v string) *DeleteSheetShrinkRequest
	GetWorkbookId() *string
}

type DeleteSheetShrinkRequest struct {
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

func (s DeleteSheetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSheetShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteSheetShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *DeleteSheetShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteSheetShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *DeleteSheetShrinkRequest) SetSheetId(v string) *DeleteSheetShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *DeleteSheetShrinkRequest) SetTenantContextShrink(v string) *DeleteSheetShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteSheetShrinkRequest) SetWorkbookId(v string) *DeleteSheetShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *DeleteSheetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
