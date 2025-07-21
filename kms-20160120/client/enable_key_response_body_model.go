// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableKeyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableKeyResponseBody
  GetRequestId() *string 
}

type EnableKeyResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // efb1cbbd-a093-4278-bc03-639dd4fcc207
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableKeyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableKeyResponseBody) GoString() string {
  return s.String()
}

func (s *EnableKeyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableKeyResponseBody) SetRequestId(v string) *EnableKeyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableKeyResponseBody) Validate() error {
  return dara.Validate(s)
}

