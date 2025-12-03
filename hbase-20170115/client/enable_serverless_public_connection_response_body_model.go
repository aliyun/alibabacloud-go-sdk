// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableServerlessPublicConnectionResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableServerlessPublicConnectionResponseBody
  GetRequestId() *string 
}

type EnableServerlessPublicConnectionResponseBody struct {
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableServerlessPublicConnectionResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableServerlessPublicConnectionResponseBody) GoString() string {
  return s.String()
}

func (s *EnableServerlessPublicConnectionResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableServerlessPublicConnectionResponseBody) SetRequestId(v string) *EnableServerlessPublicConnectionResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableServerlessPublicConnectionResponseBody) Validate() error {
  return dara.Validate(s)
}

