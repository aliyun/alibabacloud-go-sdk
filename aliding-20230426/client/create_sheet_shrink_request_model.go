// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSheetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateSheetShrinkRequest
	GetName() *string
	SetTenantContextShrink(v string) *CreateSheetShrinkRequest
	GetTenantContextShrink() *string
	SetWorkbookId(v string) *CreateSheetShrinkRequest
	GetWorkbookId() *string
}

type CreateSheetShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Sheet1
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// stxxxx
	WorkbookId *string `json:"WorkbookId,omitempty" xml:"WorkbookId,omitempty"`
}

func (s CreateSheetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSheetShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSheetShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateSheetShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateSheetShrinkRequest) GetWorkbookId() *string {
	return s.WorkbookId
}

func (s *CreateSheetShrinkRequest) SetName(v string) *CreateSheetShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateSheetShrinkRequest) SetTenantContextShrink(v string) *CreateSheetShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateSheetShrinkRequest) SetWorkbookId(v string) *CreateSheetShrinkRequest {
	s.WorkbookId = &v
	return s
}

func (s *CreateSheetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
