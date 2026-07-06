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
	// Consumer Group name.
	//
	// - Can only contain letters, numbers, hyphens (-), and underscores (_).
	//
	// - Length must be **3-64*	- characters. If more than **64*	- characters are provided, they will be automatically truncated.
	//
	// - Cannot be modified once created.
	//
	// This parameter is required.
	//
	// example:
	//
	// kafka-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp91inkw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// When resetType is offset, this parameter is used to set the consumer offset for each partition of a topic for the consumer group.
	//
	// if can be null:
	// true
	Offsets []*UpdateConsumerOffsetRequestOffsets `json:"Offsets,omitempty" xml:"Offsets,omitempty" type:"Repeated"`
	// Region ID of the instance to which the Group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Type of consumer group offset reset, supporting the following two types:
	//
	// - **timestamp*	- (default)
	//
	// - **offset**
	//
	// example:
	//
	// timestamp
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// Time parameter in Unix timestamp format, in milliseconds.
	//
	// The parameter range should be **less than 0*	- or **within the retention period of the consumer offset**. This parameter only takes effect when resetType is timestamp.
	//
	// - To reset to the latest consumer offset, pass -1.
	//
	// - To reset to the earliest consumer offset, pass -2.
	//
	// example:
	//
	// -1
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// Topic name.
	//
	// - Can only contain letters, numbers, underscores (_), and hyphens (-).
	//
	// - Length must be **3-64*	- characters. If more than **64*	- characters are provided, they will be automatically truncated.
	//
	// - Cannot be modified once created.
	//
	// **To set the consumer offset for all topics subscribed by the current consumer, pass an empty string.**
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
	if s.Offsets != nil {
		for _, item := range s.Offsets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateConsumerOffsetRequestOffsets struct {
	// Partition offset.
	//
	// example:
	//
	// 1
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// Partition ID.
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
