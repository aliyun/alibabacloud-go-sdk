// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableTrustedOriginResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableTrustedOriginResponseBody
  GetRequestId() *string 
}

type EnableTrustedOriginResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 0441BD79-example
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableTrustedOriginResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableTrustedOriginResponseBody) GoString() string {
  return s.String()
}

func (s *EnableTrustedOriginResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableTrustedOriginResponseBody) SetRequestId(v string) *EnableTrustedOriginResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableTrustedOriginResponseBody) Validate() error {
  return dara.Validate(s)
}

