// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableInstancePublicAccessRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableInstancePublicAccessRequest
  GetInstanceId() *string 
  SetRegionId(v string) *EnableInstancePublicAccessRequest
  GetRegionId() *string 
}

type EnableInstancePublicAccessRequest struct {
  // The ID of the Bastionhost instance.
  // 
  // > Call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the ID of the Bastionhost instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // bastionhost-cn-78v1gh****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The region ID of the Bastionhost instance.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableInstancePublicAccessRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableInstancePublicAccessRequest) GoString() string {
  return s.String()
}

func (s *EnableInstancePublicAccessRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableInstancePublicAccessRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableInstancePublicAccessRequest) SetInstanceId(v string) *EnableInstancePublicAccessRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableInstancePublicAccessRequest) SetRegionId(v string) *EnableInstancePublicAccessRequest {
  s.RegionId = &v
  return s
}

func (s *EnableInstancePublicAccessRequest) Validate() error {
  return dara.Validate(s)
}

