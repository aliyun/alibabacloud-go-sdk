// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmptyAutoCcBlacklistRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EmptyAutoCcBlacklistRequest
  GetInstanceId() *string 
}

type EmptyAutoCcBlacklistRequest struct {
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

func (s EmptyAutoCcBlacklistRequest) String() string {
  return dara.Prettify(s)
}

func (s EmptyAutoCcBlacklistRequest) GoString() string {
  return s.String()
}

func (s *EmptyAutoCcBlacklistRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EmptyAutoCcBlacklistRequest) SetInstanceId(v string) *EmptyAutoCcBlacklistRequest {
  s.InstanceId = &v
  return s
}

func (s *EmptyAutoCcBlacklistRequest) Validate() error {
  return dara.Validate(s)
}

