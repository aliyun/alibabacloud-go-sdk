// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableProxyProtocolRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAcceptLanguage(v string) *EnableProxyProtocolRequest
  GetAcceptLanguage() *string 
  SetEnableProxyProtocol(v bool) *EnableProxyProtocolRequest
  GetEnableProxyProtocol() *bool 
  SetGatewayUniqueId(v string) *EnableProxyProtocolRequest
  GetGatewayUniqueId() *string 
}

type EnableProxyProtocolRequest struct {
  // The language in which you want to display the results. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
  // 
  // example:
  // 
  // zh
  AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
  // Specifies whether to use the proxy protocol to preserve client IP addresses. Valid values:
  // 
  // 	- **true**
  // 
  // 	- **false*	- (default)
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // true
  EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
  // The unique ID of the gateway.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // gw-c9bc5afd61014165bd58f621b491*****
  GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s EnableProxyProtocolRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableProxyProtocolRequest) GoString() string {
  return s.String()
}

func (s *EnableProxyProtocolRequest) GetAcceptLanguage() *string  {
  return s.AcceptLanguage
}

func (s *EnableProxyProtocolRequest) GetEnableProxyProtocol() *bool  {
  return s.EnableProxyProtocol
}

func (s *EnableProxyProtocolRequest) GetGatewayUniqueId() *string  {
  return s.GatewayUniqueId
}

func (s *EnableProxyProtocolRequest) SetAcceptLanguage(v string) *EnableProxyProtocolRequest {
  s.AcceptLanguage = &v
  return s
}

func (s *EnableProxyProtocolRequest) SetEnableProxyProtocol(v bool) *EnableProxyProtocolRequest {
  s.EnableProxyProtocol = &v
  return s
}

func (s *EnableProxyProtocolRequest) SetGatewayUniqueId(v string) *EnableProxyProtocolRequest {
  s.GatewayUniqueId = &v
  return s
}

func (s *EnableProxyProtocolRequest) Validate() error {
  return dara.Validate(s)
}

