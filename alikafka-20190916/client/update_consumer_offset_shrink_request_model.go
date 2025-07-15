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
	OffsetsShrink *string `json:"Offsets,omitempty" xml:"Offsets,omitempty"`
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
