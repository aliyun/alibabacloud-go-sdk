// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCheckProductResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableCheckProductResponseBody
  GetRequestId() *string 
}

type EnableCheckProductResponseBody struct {
  // The unique ID of the request.
  // 
  // example:
  // 
  // 8724BC18-904D-5A0D-BFF4-F0554F0037E7
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCheckProductResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCheckProductResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCheckProductResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCheckProductResponseBody) SetRequestId(v string) *EnableCheckProductResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCheckProductResponseBody) Validate() error {
  return dara.Validate(s)
}

