// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *GetCardTemplateShrinkRequest
	GetTemplateId() *string
	SetTenantContextShrink(v string) *GetCardTemplateShrinkRequest
	GetTenantContextShrink() *string
}

type GetCardTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetCardTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCardTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCardTemplateShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetCardTemplateShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetCardTemplateShrinkRequest) SetTemplateId(v string) *GetCardTemplateShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *GetCardTemplateShrinkRequest) SetTenantContextShrink(v string) *GetCardTemplateShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetCardTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
