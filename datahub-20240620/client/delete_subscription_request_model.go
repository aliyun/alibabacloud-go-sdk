// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *DeleteSubscriptionRequest
	GetProjectName() *string
	SetSubscriptionId(v string) *DeleteSubscriptionRequest
	GetSubscriptionId() *string
	SetTopicName(v string) *DeleteSubscriptionRequest
	GetTopicName() *string
}

type DeleteSubscriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1713853462590KA0YP
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s DeleteSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteSubscriptionRequest) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *DeleteSubscriptionRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *DeleteSubscriptionRequest) SetProjectName(v string) *DeleteSubscriptionRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteSubscriptionRequest) SetSubscriptionId(v string) *DeleteSubscriptionRequest {
	s.SubscriptionId = &v
	return s
}

func (s *DeleteSubscriptionRequest) SetTopicName(v string) *DeleteSubscriptionRequest {
	s.TopicName = &v
	return s
}

func (s *DeleteSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
