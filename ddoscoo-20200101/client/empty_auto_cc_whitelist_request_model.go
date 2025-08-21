// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmptyAutoCcWhitelistRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EmptyAutoCcWhitelistRequest
  GetInstanceId() *string 
}

type EmptyAutoCcWhitelistRequest struct {
  // The ID of the instance.
  // 
  // > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ddoscoo-cn-mp91j1ao****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EmptyAutoCcWhitelistRequest) String() string {
  return dara.Prettify(s)
}

func (s EmptyAutoCcWhitelistRequest) GoString() string {
  return s.String()
}

func (s *EmptyAutoCcWhitelistRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EmptyAutoCcWhitelistRequest) SetInstanceId(v string) *EmptyAutoCcWhitelistRequest {
  s.InstanceId = &v
  return s
}

func (s *EmptyAutoCcWhitelistRequest) Validate() error {
  return dara.Validate(s)
}

