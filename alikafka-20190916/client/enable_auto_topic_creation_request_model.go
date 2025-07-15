// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAutoTopicCreationRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableAutoTopicCreationRequest
  GetInstanceId() *string 
  SetOperate(v string) *EnableAutoTopicCreationRequest
  GetOperate() *string 
  SetPartitionNum(v int64) *EnableAutoTopicCreationRequest
  GetPartitionNum() *int64 
  SetRegionId(v string) *EnableAutoTopicCreationRequest
  GetRegionId() *string 
  SetUpdatePartition(v bool) *EnableAutoTopicCreationRequest
  GetUpdatePartition() *bool 
}

type EnableAutoTopicCreationRequest struct {
  // The instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // alikafka_post-cn-v0h1fgs2****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The operation that you want to perform. Valid values:
  // 
  // 	- enable: enables the automatic topic creation feature.
  // 
  // 	- disable: disables the automatic topic creation feature.
  // 
  // 	- updatePartition: changes the number of partitions in topics that are automatically created.
  // 
  // example:
  // 
  // enable
  Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
  // The changed number of partitions in topics that are automatically created.
  // 
  // This parameter takes effect only if you set Operate to updatePartition.
  // 
  // example:
  // 
  // 12
  PartitionNum *int64 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
  // The region ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  UpdatePartition *bool `json:"UpdatePartition,omitempty" xml:"UpdatePartition,omitempty"`
}

func (s EnableAutoTopicCreationRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableAutoTopicCreationRequest) GoString() string {
  return s.String()
}

func (s *EnableAutoTopicCreationRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableAutoTopicCreationRequest) GetOperate() *string  {
  return s.Operate
}

func (s *EnableAutoTopicCreationRequest) GetPartitionNum() *int64  {
  return s.PartitionNum
}

func (s *EnableAutoTopicCreationRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableAutoTopicCreationRequest) GetUpdatePartition() *bool  {
  return s.UpdatePartition
}

func (s *EnableAutoTopicCreationRequest) SetInstanceId(v string) *EnableAutoTopicCreationRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableAutoTopicCreationRequest) SetOperate(v string) *EnableAutoTopicCreationRequest {
  s.Operate = &v
  return s
}

func (s *EnableAutoTopicCreationRequest) SetPartitionNum(v int64) *EnableAutoTopicCreationRequest {
  s.PartitionNum = &v
  return s
}

func (s *EnableAutoTopicCreationRequest) SetRegionId(v string) *EnableAutoTopicCreationRequest {
  s.RegionId = &v
  return s
}

func (s *EnableAutoTopicCreationRequest) SetUpdatePartition(v bool) *EnableAutoTopicCreationRequest {
  s.UpdatePartition = &v
  return s
}

func (s *EnableAutoTopicCreationRequest) Validate() error {
  return dara.Validate(s)
}

