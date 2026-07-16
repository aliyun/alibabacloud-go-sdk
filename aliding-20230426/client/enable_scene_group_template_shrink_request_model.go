// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSceneGroupTemplateShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOpenConversationId(v string) *EnableSceneGroupTemplateShrinkRequest
  GetOpenConversationId() *string 
  SetTemplateId(v string) *EnableSceneGroupTemplateShrinkRequest
  GetTemplateId() *string 
  SetTenantContextShrink(v string) *EnableSceneGroupTemplateShrinkRequest
  GetTenantContextShrink() *string 
}

type EnableSceneGroupTemplateShrinkRequest struct {
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
  TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
  TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s EnableSceneGroupTemplateShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSceneGroupTemplateShrinkRequest) GoString() string {
  return s.String()
}

func (s *EnableSceneGroupTemplateShrinkRequest) GetOpenConversationId() *string  {
  return s.OpenConversationId
}

func (s *EnableSceneGroupTemplateShrinkRequest) GetTemplateId() *string  {
  return s.TemplateId
}

func (s *EnableSceneGroupTemplateShrinkRequest) GetTenantContextShrink() *string  {
  return s.TenantContextShrink
}

func (s *EnableSceneGroupTemplateShrinkRequest) SetOpenConversationId(v string) *EnableSceneGroupTemplateShrinkRequest {
  s.OpenConversationId = &v
  return s
}

func (s *EnableSceneGroupTemplateShrinkRequest) SetTemplateId(v string) *EnableSceneGroupTemplateShrinkRequest {
  s.TemplateId = &v
  return s
}

func (s *EnableSceneGroupTemplateShrinkRequest) SetTenantContextShrink(v string) *EnableSceneGroupTemplateShrinkRequest {
  s.TenantContextShrink = &v
  return s
}

func (s *EnableSceneGroupTemplateShrinkRequest) Validate() error {
  return dara.Validate(s)
}

