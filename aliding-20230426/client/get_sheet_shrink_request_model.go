// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSheetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSheetId(v string) *GetSheetShrinkRequest
	GetSheetId() *string
	SetTenantContextShrink(v string) *GetSheetShrinkRequest
	GetTenantContextShrink() *string
	SetWorkbookId(v string) *GetSheetShrinkRequest
	GetWorkbookId() *string
}

type GetSheetShrinkRequest struct {
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

func (s GetSheetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSheetShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSheetShrinkRequest) GetSheetId() *string {
	return s.SheetId
}

func (s *GetSheetShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetSheetShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *GetSheetShrinkRequest) SetSheetId(v string) *GetSheetShrinkRequest {
	s.SheetId = &v
	return s
}

func (s *GetSheetShrinkRequest) SetTenantContextShrink(v string) *GetSheetShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetSheetShrinkRequest) SetWorkbookId(v string) *GetSheetShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *GetSheetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
