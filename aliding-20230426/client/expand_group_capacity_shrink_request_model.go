// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandGroupCapacityShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOpenConversationId(v string) *ExpandGroupCapacityShrinkRequest
  GetOpenConversationId() *string 
  SetTenantContextShrink(v string) *ExpandGroupCapacityShrinkRequest
  GetTenantContextShrink() *string 
}

type ExpandGroupCapacityShrinkRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // cidt*****Xa4K10w==
  OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
  TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s ExpandGroupCapacityShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExpandGroupCapacityShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExpandGroupCapacityShrinkRequest) GetOpenConversationId() *string  {
  return s.OpenConversationId
}

func (s *ExpandGroupCapacityShrinkRequest) GetTenantContextShrink() *string  {
  return s.TenantContextShrink
}

func (s *ExpandGroupCapacityShrinkRequest) SetOpenConversationId(v string) *ExpandGroupCapacityShrinkRequest {
  s.OpenConversationId = &v
  return s
}

func (s *ExpandGroupCapacityShrinkRequest) SetTenantContextShrink(v string) *ExpandGroupCapacityShrinkRequest {
  s.TenantContextShrink = &v
  return s
}

func (s *ExpandGroupCapacityShrinkRequest) Validate() error {
  return dara.Validate(s)
}

