// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInitDomainAutoRedirectResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableInitDomainAutoRedirectResponseBody
  GetRequestId() *string 
}

type EnableInitDomainAutoRedirectResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableInitDomainAutoRedirectResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableInitDomainAutoRedirectResponseBody) GoString() string {
  return s.String()
}

func (s *EnableInitDomainAutoRedirectResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableInitDomainAutoRedirectResponseBody) SetRequestId(v string) *EnableInitDomainAutoRedirectResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableInitDomainAutoRedirectResponseBody) Validate() error {
  return dara.Validate(s)
}

