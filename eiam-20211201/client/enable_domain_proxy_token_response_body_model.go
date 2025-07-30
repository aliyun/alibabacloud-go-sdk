// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDomainProxyTokenResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableDomainProxyTokenResponseBody
  GetRequestId() *string 
}

type EnableDomainProxyTokenResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDomainProxyTokenResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDomainProxyTokenResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDomainProxyTokenResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDomainProxyTokenResponseBody) SetRequestId(v string) *EnableDomainProxyTokenResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDomainProxyTokenResponseBody) Validate() error {
  return dara.Validate(s)
}

