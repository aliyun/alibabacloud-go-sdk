// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExchangeEntitlementHeaders interface {
  dara.Model
  String() string
  GoString() string
  SetCommonHeaders(v map[string]*string) *ExchangeEntitlementHeaders
  GetCommonHeaders() map[string]*string 
  SetXLoadTest(v bool) *ExchangeEntitlementHeaders
  GetXLoadTest() *bool 
}

type ExchangeEntitlementHeaders struct {
  CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
  // example:
  // 
  // false
  XLoadTest *bool `json:"X-Load-Test,omitempty" xml:"X-Load-Test,omitempty"`
}

func (s ExchangeEntitlementHeaders) String() string {
  return dara.Prettify(s)
}

func (s ExchangeEntitlementHeaders) GoString() string {
  return s.String()
}

func (s *ExchangeEntitlementHeaders) GetCommonHeaders() map[string]*string  {
  return s.CommonHeaders
}

func (s *ExchangeEntitlementHeaders) GetXLoadTest() *bool  {
  return s.XLoadTest
}

func (s *ExchangeEntitlementHeaders) SetCommonHeaders(v map[string]*string) *ExchangeEntitlementHeaders {
  s.CommonHeaders = v
  return s
}

func (s *ExchangeEntitlementHeaders) SetXLoadTest(v bool) *ExchangeEntitlementHeaders {
  s.XLoadTest = &v
  return s
}

func (s *ExchangeEntitlementHeaders) Validate() error {
  return dara.Validate(s)
}

