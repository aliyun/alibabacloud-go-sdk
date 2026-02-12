// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsTopicDeleteRequest
	GetInstanceId() *string
	SetTopic(v string) *OnsTopicDeleteRequest
	GetTopic() *string
}

type OnsTopicDeleteRequest struct {
	// The ID of the instance to which the topic you want to delete belongs.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the topic that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicDeleteRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicDeleteRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTopicDeleteRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsTopicDeleteRequest) SetInstanceId(v string) *OnsTopicDeleteRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicDeleteRequest) SetTopic(v string) *OnsTopicDeleteRequest {
	s.Topic = &v
	return s
}

func (s *OnsTopicDeleteRequest) Validate() error {
	return dara.Validate(s)
}
