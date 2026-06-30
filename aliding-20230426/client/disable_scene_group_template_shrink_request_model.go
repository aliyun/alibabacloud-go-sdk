// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSceneGroupTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenConversationId(v string) *DisableSceneGroupTemplateShrinkRequest
	GetOpenConversationId() *string
	SetTemplateId(v string) *DisableSceneGroupTemplateShrinkRequest
	GetTemplateId() *string
	SetTenantContextShrink(v string) *DisableSceneGroupTemplateShrinkRequest
	GetTenantContextShrink() *string
}

type DisableSceneGroupTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidt*****Xa4K10w==
	OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2efdt*****fswe==
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s DisableSceneGroupTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSceneGroupTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *DisableSceneGroupTemplateShrinkRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *DisableSceneGroupTemplateShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DisableSceneGroupTemplateShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DisableSceneGroupTemplateShrinkRequest) SetOpenConversationId(v string) *DisableSceneGroupTemplateShrinkRequest {
	s.OpenConversationId = &v
	return s
}

func (s *DisableSceneGroupTemplateShrinkRequest) SetTemplateId(v string) *DisableSceneGroupTemplateShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *DisableSceneGroupTemplateShrinkRequest) SetTenantContextShrink(v string) *DisableSceneGroupTemplateShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DisableSceneGroupTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
