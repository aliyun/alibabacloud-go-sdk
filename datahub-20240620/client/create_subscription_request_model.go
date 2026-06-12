// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplication(v string) *CreateSubscriptionRequest
	GetApplication() *string
	SetComment(v string) *CreateSubscriptionRequest
	GetComment() *string
	SetProjectName(v string) *CreateSubscriptionRequest
	GetProjectName() *string
	SetSubscriptionId(v string) *CreateSubscriptionRequest
	GetSubscriptionId() *string
	SetTopicName(v string) *CreateSubscriptionRequest
	GetTopicName() *string
}

type CreateSubscriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_application
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_abc
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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

func (s CreateSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionRequest) GetApplication() *string {
	return s.Application
}

func (s *CreateSubscriptionRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateSubscriptionRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateSubscriptionRequest) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *CreateSubscriptionRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *CreateSubscriptionRequest) SetApplication(v string) *CreateSubscriptionRequest {
	s.Application = &v
	return s
}

func (s *CreateSubscriptionRequest) SetComment(v string) *CreateSubscriptionRequest {
	s.Comment = &v
	return s
}

func (s *CreateSubscriptionRequest) SetProjectName(v string) *CreateSubscriptionRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSubscriptionId(v string) *CreateSubscriptionRequest {
	s.SubscriptionId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetTopicName(v string) *CreateSubscriptionRequest {
	s.TopicName = &v
	return s
}

func (s *CreateSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
