// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHttp2Request interface {
  dara.Model
  String() string
  GoString() string
  SetAcceptLanguage(v string) *EnableHttp2Request
  GetAcceptLanguage() *string 
  SetEnableHttp2(v bool) *EnableHttp2Request
  GetEnableHttp2() *bool 
  SetGatewayUniqueId(v string) *EnableHttp2Request
  GetGatewayUniqueId() *string 
}

type EnableHttp2Request struct {
  // The language in which you want to display the results. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
  // 
  // example:
  // 
  // zh
  AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
  // Specifies whether to enable HTTP/2 for negotiation between the server and client. This parameter applies to requests. Valid values:
  // 
  // 	- **true**
  // 
  // 	- **false**
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // true
  EnableHttp2 *bool `json:"EnableHttp2,omitempty" xml:"EnableHttp2,omitempty"`
  // The unique ID of the gateway.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // gw-0adf3ad751284cc69fcf9669fba*****
  GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s EnableHttp2Request) String() string {
  return dara.Prettify(s)
}

func (s EnableHttp2Request) GoString() string {
  return s.String()
}

func (s *EnableHttp2Request) GetAcceptLanguage() *string  {
  return s.AcceptLanguage
}

func (s *EnableHttp2Request) GetEnableHttp2() *bool  {
  return s.EnableHttp2
}

func (s *EnableHttp2Request) GetGatewayUniqueId() *string  {
  return s.GatewayUniqueId
}

func (s *EnableHttp2Request) SetAcceptLanguage(v string) *EnableHttp2Request {
  s.AcceptLanguage = &v
  return s
}

func (s *EnableHttp2Request) SetEnableHttp2(v bool) *EnableHttp2Request {
  s.EnableHttp2 = &v
  return s
}

func (s *EnableHttp2Request) SetGatewayUniqueId(v string) *EnableHttp2Request {
  s.GatewayUniqueId = &v
  return s
}

func (s *EnableHttp2Request) Validate() error {
  return dara.Validate(s)
}

