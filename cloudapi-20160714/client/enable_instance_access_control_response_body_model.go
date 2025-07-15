// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInstanceAccessControlResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableInstanceAccessControlResponseBody
  GetRequestId() *string 
}

type EnableInstanceAccessControlResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // CE5722A6-AE78-4741-A9B0-6C817D360510
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableInstanceAccessControlResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableInstanceAccessControlResponseBody) GoString() string {
  return s.String()
}

func (s *EnableInstanceAccessControlResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableInstanceAccessControlResponseBody) SetRequestId(v string) *EnableInstanceAccessControlResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableInstanceAccessControlResponseBody) Validate() error {
  return dara.Validate(s)
}

