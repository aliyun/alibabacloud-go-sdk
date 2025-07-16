// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableAllFieldsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *GetMultiDimTableAllFieldsShrinkRequest
	GetBaseId() *string
	SetSheetIdOrName(v string) *GetMultiDimTableAllFieldsShrinkRequest
	GetSheetIdOrName() *string
	SetTenantContextShrink(v string) *GetMultiDimTableAllFieldsShrinkRequest
	GetTenantContextShrink() *string
}

type GetMultiDimTableAllFieldsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 338534
	BaseId *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	SheetIdOrName       *string `json:"SheetIdOrName,omitempty" xml:"SheetIdOrName,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetMultiDimTableAllFieldsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllFieldsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllFieldsShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *GetMultiDimTableAllFieldsShrinkRequest) GetSheetIdOrName() *string {
	return s.SheetIdOrName
}

func (s *GetMultiDimTableAllFieldsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetMultiDimTableAllFieldsShrinkRequest) SetBaseId(v string) *GetMultiDimTableAllFieldsShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *GetMultiDimTableAllFieldsShrinkRequest) SetSheetIdOrName(v string) *GetMultiDimTableAllFieldsShrinkRequest {
	s.SheetIdOrName = &v
	return s
}

func (s *GetMultiDimTableAllFieldsShrinkRequest) SetTenantContextShrink(v string) *GetMultiDimTableAllFieldsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetMultiDimTableAllFieldsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
