// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerOffsetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsumerId(v string) *UpdateConsumerOffsetShrinkRequest
	GetConsumerId() *string
	SetInstanceId(v string) *UpdateConsumerOffsetShrinkRequest
	GetInstanceId() *string
	SetOffsetsShrink(v string) *UpdateConsumerOffsetShrinkRequest
	GetOffsetsShrink() *string
	SetRegionId(v string) *UpdateConsumerOffsetShrinkRequest
	GetRegionId() *string
	SetResetType(v string) *UpdateConsumerOffsetShrinkRequest
	GetResetType() *string
	SetTime(v string) *UpdateConsumerOffsetShrinkRequest
	GetTime() *string
	SetTopic(v string) *UpdateConsumerOffsetShrinkRequest
	GetTopic() *string
}

type UpdateConsumerOffsetShrinkRequest struct {
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
	OffsetsShrink *string `json:"Offsets,omitempty" xml:"Offsets,omitempty"`
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

func (s UpdateConsumerOffsetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerOffsetShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateConsumerOffsetShrinkRequest) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *UpdateConsumerOffsetShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateConsumerOffsetShrinkRequest) GetOffsetsShrink() *string {
	return s.OffsetsShrink
}

func (s *UpdateConsumerOffsetShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateConsumerOffsetShrinkRequest) GetResetType() *string {
	return s.ResetType
}

func (s *UpdateConsumerOffsetShrinkRequest) GetTime() *string {
	return s.Time
}

func (s *UpdateConsumerOffsetShrinkRequest) GetTopic() *string {
	return s.Topic
}

func (s *UpdateConsumerOffsetShrinkRequest) SetConsumerId(v string) *UpdateConsumerOffsetShrinkRequest {
	s.ConsumerId = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) SetInstanceId(v string) *UpdateConsumerOffsetShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) SetOffsetsShrink(v string) *UpdateConsumerOffsetShrinkRequest {
	s.OffsetsShrink = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) SetRegionId(v string) *UpdateConsumerOffsetShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) SetResetType(v string) *UpdateConsumerOffsetShrinkRequest {
	s.ResetType = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) SetTime(v string) *UpdateConsumerOffsetShrinkRequest {
	s.Time = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) SetTopic(v string) *UpdateConsumerOffsetShrinkRequest {
	s.Topic = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
