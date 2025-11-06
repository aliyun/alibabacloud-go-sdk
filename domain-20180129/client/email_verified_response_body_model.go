// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmailVerifiedResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EmailVerifiedResponseBody
  GetRequestId() *string 
}

type EmailVerifiedResponseBody struct {
  // example:
  // 
  // BF014B60-C708-4253-B5F2-3F9B493F398B
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EmailVerifiedResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EmailVerifiedResponseBody) GoString() string {
  return s.String()
}

func (s *EmailVerifiedResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EmailVerifiedResponseBody) SetRequestId(v string) *EmailVerifiedResponseBody {
  s.RequestId = &v
  return s
}

func (s *EmailVerifiedResponseBody) Validate() error {
  return dara.Validate(s)
}

