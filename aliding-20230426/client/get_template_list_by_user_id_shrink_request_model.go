// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateListByUserIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOffset(v int64) *GetTemplateListByUserIdShrinkRequest
	GetOffset() *int64
	SetSize(v int64) *GetTemplateListByUserIdShrinkRequest
	GetSize() *int64
	SetTenantContextShrink(v string) *GetTemplateListByUserIdShrinkRequest
	GetTenantContextShrink() *string
}

type GetTemplateListByUserIdShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Size                *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetTemplateListByUserIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateListByUserIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateListByUserIdShrinkRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *GetTemplateListByUserIdShrinkRequest) GetSize() *int64 {
	return s.Size
}

func (s *GetTemplateListByUserIdShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetTemplateListByUserIdShrinkRequest) SetOffset(v int64) *GetTemplateListByUserIdShrinkRequest {
	s.Offset = &v
	return s
}

func (s *GetTemplateListByUserIdShrinkRequest) SetSize(v int64) *GetTemplateListByUserIdShrinkRequest {
	s.Size = &v
	return s
}

func (s *GetTemplateListByUserIdShrinkRequest) SetTenantContextShrink(v string) *GetTemplateListByUserIdShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetTemplateListByUserIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
