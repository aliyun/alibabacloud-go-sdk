// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *GetSubscriptionRequest
	GetProjectName() *string
	SetSubscriptionId(v string) *GetSubscriptionRequest
	GetSubscriptionId() *string
	SetTopicName(v string) *GetSubscriptionRequest
	GetTopicName() *string
}

type GetSubscriptionRequest struct {
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
	// 1764122860063VIIZ2
	SubscriptionId *string `json:"SubscriptionId,omitempty" xml:"SubscriptionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetSubscriptionRequest) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *GetSubscriptionRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *GetSubscriptionRequest) SetProjectName(v string) *GetSubscriptionRequest {
	s.ProjectName = &v
	return s
}

func (s *GetSubscriptionRequest) SetSubscriptionId(v string) *GetSubscriptionRequest {
	s.SubscriptionId = &v
	return s
}

func (s *GetSubscriptionRequest) SetTopicName(v string) *GetSubscriptionRequest {
	s.TopicName = &v
	return s
}

func (s *GetSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
