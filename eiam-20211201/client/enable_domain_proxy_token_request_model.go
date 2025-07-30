// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDomainProxyTokenRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDomainId(v string) *EnableDomainProxyTokenRequest
  GetDomainId() *string 
  SetDomainProxyTokenId(v string) *EnableDomainProxyTokenRequest
  GetDomainProxyTokenId() *string 
  SetInstanceId(v string) *EnableDomainProxyTokenRequest
  GetInstanceId() *string 
}

type EnableDomainProxyTokenRequest struct {
  // The ID of the domain name.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // dm_examplexxxxx
  DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
  // The ID of the proxy token of the domain name.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // pt_examplexxxx
  DomainProxyTokenId *string `json:"DomainProxyTokenId,omitempty" xml:"DomainProxyTokenId,omitempty"`
  // The instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableDomainProxyTokenRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDomainProxyTokenRequest) GoString() string {
  return s.String()
}

func (s *EnableDomainProxyTokenRequest) GetDomainId() *string  {
  return s.DomainId
}

func (s *EnableDomainProxyTokenRequest) GetDomainProxyTokenId() *string  {
  return s.DomainProxyTokenId
}

func (s *EnableDomainProxyTokenRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableDomainProxyTokenRequest) SetDomainId(v string) *EnableDomainProxyTokenRequest {
  s.DomainId = &v
  return s
}

func (s *EnableDomainProxyTokenRequest) SetDomainProxyTokenId(v string) *EnableDomainProxyTokenRequest {
  s.DomainProxyTokenId = &v
  return s
}

func (s *EnableDomainProxyTokenRequest) SetInstanceId(v string) *EnableDomainProxyTokenRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableDomainProxyTokenRequest) Validate() error {
  return dara.Validate(s)
}

