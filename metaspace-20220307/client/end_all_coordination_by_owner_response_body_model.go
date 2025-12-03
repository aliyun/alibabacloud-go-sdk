// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndAllCoordinationByOwnerResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EndAllCoordinationByOwnerResponseBody
  GetRequestId() *string 
}

type EndAllCoordinationByOwnerResponseBody struct {
  // example:
  // 
  // AD2D0761-1FE5-549D-B169******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EndAllCoordinationByOwnerResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EndAllCoordinationByOwnerResponseBody) GoString() string {
  return s.String()
}

func (s *EndAllCoordinationByOwnerResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EndAllCoordinationByOwnerResponseBody) SetRequestId(v string) *EndAllCoordinationByOwnerResponseBody {
  s.RequestId = &v
  return s
}

func (s *EndAllCoordinationByOwnerResponseBody) Validate() error {
  return dara.Validate(s)
}

