// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServiceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableServiceResponseBody
  GetRequestId() *string 
}

type EnableServiceResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 3D57EAD2-8723-1F26-B69C-F8707D8B565D
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableServiceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableServiceResponseBody) GoString() string {
  return s.String()
}

func (s *EnableServiceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableServiceResponseBody) SetRequestId(v string) *EnableServiceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableServiceResponseBody) Validate() error {
  return dara.Validate(s)
}

