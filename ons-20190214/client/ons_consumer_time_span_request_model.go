// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerTimeSpanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *OnsConsumerTimeSpanRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsConsumerTimeSpanRequest
	GetInstanceId() *string
	SetTopic(v string) *OnsConsumerTimeSpanRequest
	GetTopic() *string
}

type OnsConsumerTimeSpanRequest struct {
	// The ID of the consumer group whose reset time range you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The topic to which the consumer group subscribes.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsConsumerTimeSpanRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerTimeSpanRequest) GoString() string {
	return s.String()
}

func (s *OnsConsumerTimeSpanRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsConsumerTimeSpanRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsConsumerTimeSpanRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsConsumerTimeSpanRequest) SetGroupId(v string) *OnsConsumerTimeSpanRequest {
	s.GroupId = &v
	return s
}

func (s *OnsConsumerTimeSpanRequest) SetInstanceId(v string) *OnsConsumerTimeSpanRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerTimeSpanRequest) SetTopic(v string) *OnsConsumerTimeSpanRequest {
	s.Topic = &v
	return s
}

func (s *OnsConsumerTimeSpanRequest) Validate() error {
	return dara.Validate(s)
}
