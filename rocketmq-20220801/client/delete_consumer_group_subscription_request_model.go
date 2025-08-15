// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerGroupSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterExpression(v string) *DeleteConsumerGroupSubscriptionRequest
	GetFilterExpression() *string
	SetFilterType(v string) *DeleteConsumerGroupSubscriptionRequest
	GetFilterType() *string
	SetTopicName(v string) *DeleteConsumerGroupSubscriptionRequest
	GetTopicName() *string
}

type DeleteConsumerGroupSubscriptionRequest struct {
	// The filter expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// *
	FilterExpression *string `json:"filterExpression,omitempty" xml:"filterExpression,omitempty"`
	// The type of the filter expression. Valid values:
	//
	// 	- SQL: filters messages by using SQL expressions.
	//
	// 	- TAG: filters messages by using tags.
	//
	// This parameter is required.
	//
	// example:
	//
	// TAG
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// The topic name.
	//
	// This parameter is required.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s DeleteConsumerGroupSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerGroupSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupSubscriptionRequest) GetFilterExpression() *string {
	return s.FilterExpression
}

func (s *DeleteConsumerGroupSubscriptionRequest) GetFilterType() *string {
	return s.FilterType
}

func (s *DeleteConsumerGroupSubscriptionRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *DeleteConsumerGroupSubscriptionRequest) SetFilterExpression(v string) *DeleteConsumerGroupSubscriptionRequest {
	s.FilterExpression = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionRequest) SetFilterType(v string) *DeleteConsumerGroupSubscriptionRequest {
	s.FilterType = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionRequest) SetTopicName(v string) *DeleteConsumerGroupSubscriptionRequest {
	s.TopicName = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
