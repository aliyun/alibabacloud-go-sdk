// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPartitionNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddPartitionNum(v int32) *ModifyPartitionNumRequest
	GetAddPartitionNum() *int32
	SetInstanceId(v string) *ModifyPartitionNumRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyPartitionNumRequest
	GetRegionId() *string
	SetTopic(v string) *ModifyPartitionNumRequest
	GetTopic() *string
}

type ModifyPartitionNumRequest struct {
	// Number of partitions to add.
	//
	// - Must be greater than 0.
	//
	// - It is recommended that the number of partitions is a multiple of 6 to reduce the risk of data skew.
	//
	// - The total number of partitions is limited to 1-360.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	AddPartitionNum *int32 `json:"AddPartitionNum,omitempty" xml:"AddPartitionNum,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-0pp1l9z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Topic name.
	//
	// This parameter is required.
	//
	// example:
	//
	// TopicPartitionNum
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ModifyPartitionNumRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPartitionNumRequest) GoString() string {
	return s.String()
}

func (s *ModifyPartitionNumRequest) GetAddPartitionNum() *int32 {
	return s.AddPartitionNum
}

func (s *ModifyPartitionNumRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPartitionNumRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPartitionNumRequest) GetTopic() *string {
	return s.Topic
}

func (s *ModifyPartitionNumRequest) SetAddPartitionNum(v int32) *ModifyPartitionNumRequest {
	s.AddPartitionNum = &v
	return s
}

func (s *ModifyPartitionNumRequest) SetInstanceId(v string) *ModifyPartitionNumRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPartitionNumRequest) SetRegionId(v string) *ModifyPartitionNumRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPartitionNumRequest) SetTopic(v string) *ModifyPartitionNumRequest {
	s.Topic = &v
	return s
}

func (s *ModifyPartitionNumRequest) Validate() error {
	return dara.Validate(s)
}
