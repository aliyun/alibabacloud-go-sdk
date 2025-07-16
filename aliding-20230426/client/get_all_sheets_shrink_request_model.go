// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllSheetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetAllSheetsShrinkRequest
	GetTenantContextShrink() *string
	SetWorkbookId(v string) *GetAllSheetsShrinkRequest
	GetWorkbookId() *string
}

type GetAllSheetsShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s GetAllSheetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAllSheetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAllSheetsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetAllSheetsShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *GetAllSheetsShrinkRequest) SetTenantContextShrink(v string) *GetAllSheetsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetAllSheetsShrinkRequest) SetWorkbookId(v string) *GetAllSheetsShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *GetAllSheetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
