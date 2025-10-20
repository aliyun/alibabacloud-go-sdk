// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBrandResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableBrandResponseBody
  GetRequestId() *string 
}

type EnableBrandResponseBody struct {
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableBrandResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableBrandResponseBody) GoString() string {
  return s.String()
}

func (s *EnableBrandResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableBrandResponseBody) SetRequestId(v string) *EnableBrandResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableBrandResponseBody) Validate() error {
  return dara.Validate(s)
}

