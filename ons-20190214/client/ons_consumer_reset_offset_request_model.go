// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerResetOffsetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *OnsConsumerResetOffsetRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsConsumerResetOffsetRequest
	GetInstanceId() *string
	SetResetTimestamp(v int64) *OnsConsumerResetOffsetRequest
	GetResetTimestamp() *int64
	SetTopic(v string) *OnsConsumerResetOffsetRequest
	GetTopic() *string
	SetType(v int32) *OnsConsumerResetOffsetRequest
	GetType() *int32
}

type OnsConsumerResetOffsetRequest struct {
	// The ID of the consumer group whose dead-letter message you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_consumer_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The timestamp to which you want to reset the consumer offset. This parameter takes effect only when the **Type*	- parameter is set to **1**. Unit: milliseconds.
	//
	// example:
	//
	// 1591153871000
	ResetTimestamp *int64 `json:"ResetTimestamp,omitempty" xml:"ResetTimestamp,omitempty"`
	// The name of the topic for which you want to reset the consumer offset.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq-topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The method that you want to use to clear accumulated messages. Valid values:
	//
	// 	- **0:*	- All accumulated messages are cleared. Messages that are not consumed are ignored. Consumers in the specified consumer group consume only messages that are published to the topic after the specified point in time.
	//
	// If "reconsumeLater" is returned, the accumulated messages are not cleared because the system is retrying to resend the messages to consumers.
	//
	// 	- **1:*	- The messages that were published to the topic before the specified point in time are cleared. You must specify a point in time. Consumers in the specified consumer group consume only the messages that are published to the topic after the specified point in time.
	//
	// You can specify a point in time from the earliest point in time when a message was published to the topic to the most recent point in time when a message was published to the topic. Points in time that are not within the allowed time range are invalid.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OnsConsumerResetOffsetRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerResetOffsetRequest) GoString() string {
	return s.String()
}

func (s *OnsConsumerResetOffsetRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsConsumerResetOffsetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsConsumerResetOffsetRequest) GetResetTimestamp() *int64 {
	return s.ResetTimestamp
}

func (s *OnsConsumerResetOffsetRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsConsumerResetOffsetRequest) GetType() *int32 {
	return s.Type
}

func (s *OnsConsumerResetOffsetRequest) SetGroupId(v string) *OnsConsumerResetOffsetRequest {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerResetOffsetRequest) SetInstanceId(v string) *OnsConsumerResetOffsetRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerResetOffsetRequest) SetResetTimestamp(v int64) *OnsConsumerResetOffsetRequest {
	s.ResetTimestamp = &v
	return s
}

func (s *OnsConsumerResetOffsetRequest) SetTopic(v string) *OnsConsumerResetOffsetRequest {
	s.Topic = &v
	return s
}

func (s *OnsConsumerResetOffsetRequest) SetType(v int32) *OnsConsumerResetOffsetRequest {
	s.Type = &v
	return s
}

func (s *OnsConsumerResetOffsetRequest) Validate() error {
	return dara.Validate(s)
}
