// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAutoGroupCreationRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEnable(v bool) *EnableAutoGroupCreationRequest
  GetEnable() *bool 
  SetInstanceId(v string) *EnableAutoGroupCreationRequest
  GetInstanceId() *string 
  SetRegionId(v string) *EnableAutoGroupCreationRequest
  GetRegionId() *string 
}

type EnableAutoGroupCreationRequest struct {
  // Specify whether to enable the flexible group creation feature. Valid values:
  // 
  // 	- **true**: enables the flexible group creation feature.
  // 
  // 	- **false**: disabled the flexible group creation feature.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // true
  Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
  // The instance ID.
  // 
  // You can call the [GetInstanceList](https://help.aliyun.com/document_detail/437663.html) operation to query instances.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // alikafka_post-cn-mp919o4v****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The region ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableAutoGroupCreationRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableAutoGroupCreationRequest) GoString() string {
  return s.String()
}

func (s *EnableAutoGroupCreationRequest) GetEnable() *bool  {
  return s.Enable
}

func (s *EnableAutoGroupCreationRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableAutoGroupCreationRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableAutoGroupCreationRequest) SetEnable(v bool) *EnableAutoGroupCreationRequest {
  s.Enable = &v
  return s
}

func (s *EnableAutoGroupCreationRequest) SetInstanceId(v string) *EnableAutoGroupCreationRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableAutoGroupCreationRequest) SetRegionId(v string) *EnableAutoGroupCreationRequest {
  s.RegionId = &v
  return s
}

func (s *EnableAutoGroupCreationRequest) Validate() error {
  return dara.Validate(s)
}

