// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerOffsetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *UpdateConsumerOffsetRequest
	GetConsumerId() *string
	SetInstanceId(v string) *UpdateConsumerOffsetRequest
	GetInstanceId() *string
	SetOffsets(v []*UpdateConsumerOffsetRequestOffsets) *UpdateConsumerOffsetRequest
	GetOffsets() []*UpdateConsumerOffsetRequestOffsets
	SetRegionId(v string) *UpdateConsumerOffsetRequest
	GetRegionId() *string
	SetResetType(v string) *UpdateConsumerOffsetRequest
	GetResetType() *string
	SetTime(v string) *UpdateConsumerOffsetRequest
	GetTime() *string
	SetTopic(v string) *UpdateConsumerOffsetRequest
	GetTopic() *string
}

type UpdateConsumerOffsetRequest struct {
	// The name of the consumer group.
	//
	// 	- The name can contain letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must be **3 to 64*	- characters in length. If a name contains more than **64*	- characters, the name is automatically truncated.
	//
	// 	- The name of a consumer group cannot be changed after the consumer group is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// kafka-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp91inkw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// If you set resetType to offset, you can use this parameter to reset the consumer offset of each partition of a specific topic in the consumer group.
	//
	// if can be null:
	// true
	Offsets []*UpdateConsumerOffsetRequestOffsets `json:"Offsets,omitempty" xml:"Offsets,omitempty" type:"Repeated"`
	// The region ID of the instance to which the consumer group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The method that is used to reset the consumer offsets of the subscribed topics of a consumer group. Valid values:
	//
	// 	- **timestamp*	- (default)
	//
	// 	- **offset**
	//
	// example:
	//
	// timestamp
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The point in time when message consumption starts. The value of this parameter is a UNIX timestamp in milliseconds. The value of this parameter must be **less than 0*	- or **within the retention period of the consumer offset**. This parameter takes effect only if you set resetType to timestamp.
	//
	// 	- If you want to reset the consumer offset to the latest offset, set this parameter to -1.
	//
	// 	- If you want to reset the consumer offset to the earliest offset, set this parameter to -2.
	//
	// example:
	//
	// -1
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The topic name.
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must be **3 to 64*	- characters in length. If a name contains more than **64*	- characters, the name is automatically truncated.
	//
	// 	- The name of a topic cannot be changed after the topic is created.
	//
	// **If you want to reset the consumer offsets of all topics to which the consumer subscribes, specify an empty string.**
	//
	// This parameter is required.
	//
	// example:
	//
	// topic_name
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateConsumerOffsetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerOffsetRequest) GoString() string {
	return s.String()
}

func (s *UpdateConsumerOffsetRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *UpdateConsumerOffsetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateConsumerOffsetRequest) GetOffsets() []*UpdateConsumerOffsetRequestOffsets {
	return s.Offsets
}

func (s *UpdateConsumerOffsetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateConsumerOffsetRequest) GetResetType() *string {
	return s.ResetType
}

func (s *UpdateConsumerOffsetRequest) GetTime() *string {
	return s.Time
}

func (s *UpdateConsumerOffsetRequest) GetTopic() *string {
	return s.Topic
}

func (s *UpdateConsumerOffsetRequest) SetConsumerId(v string) *UpdateConsumerOffsetRequest {
	s.ConsumerId = &v
	return s
}

func (s *UpdateConsumerOffsetRequest) SetInstanceId(v string) *UpdateConsumerOffsetRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConsumerOffsetRequest) SetOffsets(v []*UpdateConsumerOffsetRequestOffsets) *UpdateConsumerOffsetRequest {
	s.Offsets = v
	return s
}

func (s *UpdateConsumerOffsetRequest) SetRegionId(v string) *UpdateConsumerOffsetRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateConsumerOffsetRequest) SetResetType(v string) *UpdateConsumerOffsetRequest {
	s.ResetType = &v
	return s
}

func (s *UpdateConsumerOffsetRequest) SetTime(v string) *UpdateConsumerOffsetRequest {
	s.Time = &v
	return s
}

func (s *UpdateConsumerOffsetRequest) SetTopic(v string) *UpdateConsumerOffsetRequest {
	s.Topic = &v
	return s
}

func (s *UpdateConsumerOffsetRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateConsumerOffsetRequestOffsets struct {
	// The consumer offset of the partition.
	//
	// example:
	//
	// 1
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The partition ID.
	//
	// example:
	//
	// 0
	Partition *int32 `json:"Partition,omitempty" xml:"Partition,omitempty"`
}

func (s UpdateConsumerOffsetRequestOffsets) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerOffsetRequestOffsets) GoString() string {
	return s.String()
}

func (s *UpdateConsumerOffsetRequestOffsets) GetOffset() *int64 {
	return s.Offset
}

func (s *UpdateConsumerOffsetRequestOffsets) GetPartition() *int32 {
	return s.Partition
}

func (s *UpdateConsumerOffsetRequestOffsets) SetOffset(v int64) *UpdateConsumerOffsetRequestOffsets {
	s.Offset = &v
	return s
}

func (s *UpdateConsumerOffsetRequestOffsets) SetPartition(v int32) *UpdateConsumerOffsetRequestOffsets {
	s.Partition = &v
	return s
}

func (s *UpdateConsumerOffsetRequestOffsets) Validate() error {
	return dara.Validate(s)
}
