// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInstancePublicAccessResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableInstancePublicAccessResponseBody
  GetInstanceId() *string 
  SetRequestId(v string) *EnableInstancePublicAccessResponseBody
  GetRequestId() *string 
}

type EnableInstancePublicAccessResponseBody struct {
  // The ID of the bastion host whose Internet access is enabled.
  // 
  // example:
  // 
  // bastionhost-cn-78v1gh****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The ID of the request, which is used to locate and troubleshoot issues.
  // 
  // example:
  // 
  // D47B5043-FDD6-4FBE-976E-5FC67A23578F
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableInstancePublicAccessResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableInstancePublicAccessResponseBody) GoString() string {
  return s.String()
}

func (s *EnableInstancePublicAccessResponseBody) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableInstancePublicAccessResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableInstancePublicAccessResponseBody) SetInstanceId(v string) *EnableInstancePublicAccessResponseBody {
  s.InstanceId = &v
  return s
}

func (s *EnableInstancePublicAccessResponseBody) SetRequestId(v string) *EnableInstancePublicAccessResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableInstancePublicAccessResponseBody) Validate() error {
  return dara.Validate(s)
}

