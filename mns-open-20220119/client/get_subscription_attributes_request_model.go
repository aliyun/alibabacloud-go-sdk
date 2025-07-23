// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubscriptionName(v string) *GetSubscriptionAttributesRequest
	GetSubscriptionName() *string
	SetTopicName(v string) *GetSubscriptionAttributesRequest
	GetTopicName() *string
}

type GetSubscriptionAttributesRequest struct {
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
	// MyTopic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetSubscriptionAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionAttributesRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesRequest) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *GetSubscriptionAttributesRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *GetSubscriptionAttributesRequest) SetSubscriptionName(v string) *GetSubscriptionAttributesRequest {
	s.SubscriptionName = &v
	return s
}

func (s *GetSubscriptionAttributesRequest) SetTopicName(v string) *GetSubscriptionAttributesRequest {
	s.TopicName = &v
	return s
}

func (s *GetSubscriptionAttributesRequest) Validate() error {
	return dara.Validate(s)
}
