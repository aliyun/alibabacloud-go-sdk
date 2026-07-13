// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCheckResourceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableCheckResourceResponseBody
  GetRequestId() *string 
}

type EnableCheckResourceResponseBody struct {
  // Unique identifier of the request.
  // 
  // example:
  // 
  // E583A0FF-803C-51C4-9AC9-E029471ACD6A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCheckResourceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCheckResourceResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCheckResourceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCheckResourceResponseBody) SetRequestId(v string) *EnableCheckResourceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCheckResourceResponseBody) Validate() error {
  return dara.Validate(s)
}

