// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubscriptionName(v string) *UnsubscribeRequest
	GetSubscriptionName() *string
	SetTopicName(v string) *UnsubscribeRequest
	GetTopicName() *string
}

type UnsubscribeRequest struct {
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s UnsubscribeRequest) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeRequest) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *UnsubscribeRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *UnsubscribeRequest) SetSubscriptionName(v string) *UnsubscribeRequest {
	s.SubscriptionName = &v
	return s
}

func (s *UnsubscribeRequest) SetTopicName(v string) *UnsubscribeRequest {
	s.TopicName = &v
	return s
}

func (s *UnsubscribeRequest) Validate() error {
	return dara.Validate(s)
}
