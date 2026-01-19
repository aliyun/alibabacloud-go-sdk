// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableClientPublicKeyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableClientPublicKeyResponseBody
  GetRequestId() *string 
}

type EnableClientPublicKeyResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableClientPublicKeyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableClientPublicKeyResponseBody) GoString() string {
  return s.String()
}

func (s *EnableClientPublicKeyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableClientPublicKeyResponseBody) SetRequestId(v string) *EnableClientPublicKeyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableClientPublicKeyResponseBody) Validate() error {
  return dara.Validate(s)
}

