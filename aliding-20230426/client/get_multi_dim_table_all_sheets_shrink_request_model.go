// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableAllSheetsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *GetMultiDimTableAllSheetsShrinkRequest
	GetBaseId() *string
	SetTenantContextShrink(v string) *GetMultiDimTableAllSheetsShrinkRequest
	GetTenantContextShrink() *string
}

type GetMultiDimTableAllSheetsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 144972
	BaseId              *string `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetMultiDimTableAllSheetsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableAllSheetsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableAllSheetsShrinkRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *GetMultiDimTableAllSheetsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetMultiDimTableAllSheetsShrinkRequest) SetBaseId(v string) *GetMultiDimTableAllSheetsShrinkRequest {
	s.BaseId = &v
	return s
}

func (s *GetMultiDimTableAllSheetsShrinkRequest) SetTenantContextShrink(v string) *GetMultiDimTableAllSheetsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetMultiDimTableAllSheetsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
