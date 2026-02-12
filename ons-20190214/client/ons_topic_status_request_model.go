// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsTopicStatusRequest
	GetInstanceId() *string
	SetTopic(v string) *OnsTopicStatusRequest
	GetTopic() *string
}

type OnsTopicStatusRequest struct {
	// The ID of the instance that contains the topic you want to query.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the topic that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsTopicStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicStatusRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTopicStatusRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsTopicStatusRequest) SetInstanceId(v string) *OnsTopicStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicStatusRequest) SetTopic(v string) *OnsTopicStatusRequest {
	s.Topic = &v
	return s
}

func (s *OnsTopicStatusRequest) Validate() error {
	return dara.Validate(s)
}
