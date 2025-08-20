// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAdviceServiceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableAdviceServiceResponseBody
  GetRequestId() *string 
}

type EnableAdviceServiceResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // E1745C03-7CCE-55CF-932E-08121AAFA6AF
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAdviceServiceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAdviceServiceResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAdviceServiceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAdviceServiceResponseBody) SetRequestId(v string) *EnableAdviceServiceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAdviceServiceResponseBody) Validate() error {
  return dara.Validate(s)
}

