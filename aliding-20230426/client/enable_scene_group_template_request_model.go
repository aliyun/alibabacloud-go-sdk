// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSceneGroupTemplateRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOpenConversationId(v string) *EnableSceneGroupTemplateRequest
  GetOpenConversationId() *string 
  SetTemplateId(v string) *EnableSceneGroupTemplateRequest
  GetTemplateId() *string 
  SetTenantContext(v *EnableSceneGroupTemplateRequestTenantContext) *EnableSceneGroupTemplateRequest
  GetTenantContext() *EnableSceneGroupTemplateRequestTenantContext 
}

type EnableSceneGroupTemplateRequest struct {
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
  TenantContext *EnableSceneGroupTemplateRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s EnableSceneGroupTemplateRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSceneGroupTemplateRequest) GoString() string {
  return s.String()
}

func (s *EnableSceneGroupTemplateRequest) GetOpenConversationId() *string  {
  return s.OpenConversationId
}

func (s *EnableSceneGroupTemplateRequest) GetTemplateId() *string  {
  return s.TemplateId
}

func (s *EnableSceneGroupTemplateRequest) GetTenantContext() *EnableSceneGroupTemplateRequestTenantContext  {
  return s.TenantContext
}

func (s *EnableSceneGroupTemplateRequest) SetOpenConversationId(v string) *EnableSceneGroupTemplateRequest {
  s.OpenConversationId = &v
  return s
}

func (s *EnableSceneGroupTemplateRequest) SetTemplateId(v string) *EnableSceneGroupTemplateRequest {
  s.TemplateId = &v
  return s
}

func (s *EnableSceneGroupTemplateRequest) SetTenantContext(v *EnableSceneGroupTemplateRequestTenantContext) *EnableSceneGroupTemplateRequest {
  s.TenantContext = v
  return s
}

func (s *EnableSceneGroupTemplateRequest) Validate() error {
  if s.TenantContext != nil {
    if err := s.TenantContext.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableSceneGroupTemplateRequestTenantContext struct {
  // example:
  // 
  // xxxxxx
  TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s EnableSceneGroupTemplateRequestTenantContext) String() string {
  return dara.Prettify(s)
}

func (s EnableSceneGroupTemplateRequestTenantContext) GoString() string {
  return s.String()
}

func (s *EnableSceneGroupTemplateRequestTenantContext) GetTenantId() *string  {
  return s.TenantId
}

func (s *EnableSceneGroupTemplateRequestTenantContext) SetTenantId(v string) *EnableSceneGroupTemplateRequestTenantContext {
  s.TenantId = &v
  return s
}

func (s *EnableSceneGroupTemplateRequestTenantContext) Validate() error {
  return dara.Validate(s)
}

