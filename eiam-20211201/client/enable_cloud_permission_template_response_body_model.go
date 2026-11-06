// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCloudPermissionTemplateResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableCloudPermissionTemplateResponseBody
  GetRequestId() *string 
}

type EnableCloudPermissionTemplateResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 0441BD79-92F3-53AA-8657-F8CE4A2B912A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCloudPermissionTemplateResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableCloudPermissionTemplateResponseBody) GoString() string {
  return s.String()
}

func (s *EnableCloudPermissionTemplateResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableCloudPermissionTemplateResponseBody) SetRequestId(v string) *EnableCloudPermissionTemplateResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableCloudPermissionTemplateResponseBody) Validate() error {
  return dara.Validate(s)
}

