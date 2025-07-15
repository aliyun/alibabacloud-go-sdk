// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndCoordinationResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EndCoordinationResponseBody
  GetRequestId() *string 
}

type EndCoordinationResponseBody struct {
  // example:
  // 
  // 5C5CEF0A-D6E1-58D3-8750-67DB4F82****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EndCoordinationResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EndCoordinationResponseBody) GoString() string {
  return s.String()
}

func (s *EndCoordinationResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EndCoordinationResponseBody) SetRequestId(v string) *EndCoordinationResponseBody {
  s.RequestId = &v
  return s
}

func (s *EndCoordinationResponseBody) Validate() error {
  return dara.Validate(s)
}

