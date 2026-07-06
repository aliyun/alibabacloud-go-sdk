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
  // Instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // alikafka_post-cn-v0h1fgs2****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // Currently only these three request parameters are supported:
  // 
  // - enable: Enable automatic topic creation.
  // 
  // - disable: Disable automatic topic creation.
  // 
  // - updatePartition: Modify the number of partitions for automatic creation.
  // 
  // example:
  // 
  // enable
  Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
  // Adjust the default number of partitions for automatically created topics.
  // 
  // > This value is passed only when the Operate value is updatePartition, or when UpdatePartition is true.
  // 
  // example:
  // 
  // 12
  PartitionNum *int64 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
  // Region ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // Modify the number of partitions for automatic creation.
  // 
  // > If this parameter is set to true, the Operate parameter must be updatePartition or left empty.
  // 
  // example:
  // 
  // true
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

