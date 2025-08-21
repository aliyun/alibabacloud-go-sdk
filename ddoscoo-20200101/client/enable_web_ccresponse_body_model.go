// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWebCCResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableWebCCResponseBody
  GetRequestId() *string 
}

type EnableWebCCResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableWebCCResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableWebCCResponseBody) GoString() string {
  return s.String()
}

func (s *EnableWebCCResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableWebCCResponseBody) SetRequestId(v string) *EnableWebCCResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableWebCCResponseBody) Validate() error {
  return dara.Validate(s)
}

