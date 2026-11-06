// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCloudAccountPrivilegeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableCloudAccountPrivilegeResponseBody
  GetRequestId() *string 
}

type EnableCloudAccountPrivilegeResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCloudAccountPrivilegeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCloudAccountPrivilegeResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCloudAccountPrivilegeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCloudAccountPrivilegeResponseBody) SetRequestId(v string) *EnableCloudAccountPrivilegeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCloudAccountPrivilegeResponseBody) Validate() error {
  return dara.Validate(s)
}

