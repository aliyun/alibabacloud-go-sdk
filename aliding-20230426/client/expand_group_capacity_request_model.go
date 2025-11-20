// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandGroupCapacityRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOpenConversationId(v string) *ExpandGroupCapacityRequest
  GetOpenConversationId() *string 
  SetTenantContext(v *ExpandGroupCapacityRequestTenantContext) *ExpandGroupCapacityRequest
  GetTenantContext() *ExpandGroupCapacityRequestTenantContext 
}

type ExpandGroupCapacityRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // cidt*****Xa4K10w==
  OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
  TenantContext *ExpandGroupCapacityRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s ExpandGroupCapacityRequest) String() string {
  return dara.Prettify(s)
}

func (s ExpandGroupCapacityRequest) GoString() string {
  return s.String()
}

func (s *ExpandGroupCapacityRequest) GetOpenConversationId() *string  {
  return s.OpenConversationId
}

func (s *ExpandGroupCapacityRequest) GetTenantContext() *ExpandGroupCapacityRequestTenantContext  {
  return s.TenantContext
}

func (s *ExpandGroupCapacityRequest) SetOpenConversationId(v string) *ExpandGroupCapacityRequest {
  s.OpenConversationId = &v
  return s
}

func (s *ExpandGroupCapacityRequest) SetTenantContext(v *ExpandGroupCapacityRequestTenantContext) *ExpandGroupCapacityRequest {
  s.TenantContext = v
  return s
}

func (s *ExpandGroupCapacityRequest) Validate() error {
  if s.TenantContext != nil {
    if err := s.TenantContext.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExpandGroupCapacityRequestTenantContext struct {
  // example:
  // 
  // xxxxxx
  TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ExpandGroupCapacityRequestTenantContext) String() string {
  return dara.Prettify(s)
}

func (s ExpandGroupCapacityRequestTenantContext) GoString() string {
  return s.String()
}

func (s *ExpandGroupCapacityRequestTenantContext) GetTenantId() *string  {
  return s.TenantId
}

func (s *ExpandGroupCapacityRequestTenantContext) SetTenantId(v string) *ExpandGroupCapacityRequestTenantContext {
  s.TenantId = &v
  return s
}

func (s *ExpandGroupCapacityRequestTenantContext) Validate() error {
  return dara.Validate(s)
}

