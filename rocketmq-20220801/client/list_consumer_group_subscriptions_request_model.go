// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupSubscriptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTopicName(v string) *ListConsumerGroupSubscriptionsRequest
	GetTopicName() *string
}

type ListConsumerGroupSubscriptionsRequest struct {
	// The topic name. If you do not specify this parameter, all subscriptions of the consumer group are queried.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListConsumerGroupSubscriptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupSubscriptionsRequest) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupSubscriptionsRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *ListConsumerGroupSubscriptionsRequest) SetTopicName(v string) *ListConsumerGroupSubscriptionsRequest {
	s.TopicName = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsRequest) Validate() error {
	return dara.Validate(s)
}
