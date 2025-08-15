// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerGroupLagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTopicName(v string) *GetConsumerGroupLagRequest
	GetTopicName() *string
}

type GetConsumerGroupLagRequest struct {
	// The topic name.
	//
	// example:
	//
	// normal-topic-1
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s GetConsumerGroupLagRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupLagRequest) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupLagRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *GetConsumerGroupLagRequest) SetTopicName(v string) *GetConsumerGroupLagRequest {
	s.TopicName = &v
	return s
}

func (s *GetConsumerGroupLagRequest) Validate() error {
	return dara.Validate(s)
}
