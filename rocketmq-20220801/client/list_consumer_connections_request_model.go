// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiteTopicName(v string) *ListConsumerConnectionsRequest
	GetLiteTopicName() *string
	SetTopicName(v string) *ListConsumerConnectionsRequest
	GetTopicName() *string
}

type ListConsumerConnectionsRequest struct {
	// example:
	//
	// abc
	LiteTopicName *string `json:"liteTopicName,omitempty" xml:"liteTopicName,omitempty"`
	// example:
	//
	// test1
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListConsumerConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListConsumerConnectionsRequest) GetLiteTopicName() *string {
	return s.LiteTopicName
}

func (s *ListConsumerConnectionsRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *ListConsumerConnectionsRequest) SetLiteTopicName(v string) *ListConsumerConnectionsRequest {
	s.LiteTopicName = &v
	return s
}

func (s *ListConsumerConnectionsRequest) SetTopicName(v string) *ListConsumerConnectionsRequest {
	s.TopicName = &v
	return s
}

func (s *ListConsumerConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
