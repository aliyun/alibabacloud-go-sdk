// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationApiInvokeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableApplicationApiInvokeResponseBody
  GetRequestId() *string 
}

type EnableApplicationApiInvokeResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationApiInvokeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationApiInvokeResponseBody) GoString() string {
  return s.String()
}

func (s *EnableApplicationApiInvokeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableApplicationApiInvokeResponseBody) SetRequestId(v string) *EnableApplicationApiInvokeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableApplicationApiInvokeResponseBody) Validate() error {
  return dara.Validate(s)
}

