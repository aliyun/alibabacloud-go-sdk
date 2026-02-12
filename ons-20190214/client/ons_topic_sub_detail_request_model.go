// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTopicSubDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnsTopicSubDetailRequest
	GetInstanceId() *string
	SetTopic(v string) *OnsTopicSubDetailRequest
	GetTopic() *string
}

type OnsTopicSubDetailRequest struct {
	// The ID of the instance that contains the topic you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
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

func (s OnsTopicSubDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsTopicSubDetailRequest) GoString() string {
	return s.String()
}

func (s *OnsTopicSubDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsTopicSubDetailRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsTopicSubDetailRequest) SetInstanceId(v string) *OnsTopicSubDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsTopicSubDetailRequest) SetTopic(v string) *OnsTopicSubDetailRequest {
	s.Topic = &v
	return s
}

func (s *OnsTopicSubDetailRequest) Validate() error {
	return dara.Validate(s)
}
