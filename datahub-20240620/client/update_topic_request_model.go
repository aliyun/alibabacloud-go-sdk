// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *UpdateTopicRequest
	GetComment() *string
	SetProjectName(v string) *UpdateTopicRequest
	GetProjectName() *string
	SetTopicName(v string) *UpdateTopicRequest
	GetTopicName() *string
}

type UpdateTopicRequest struct {
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
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
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s UpdateTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicRequest) GoString() string {
	return s.String()
}

func (s *UpdateTopicRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateTopicRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateTopicRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *UpdateTopicRequest) SetComment(v string) *UpdateTopicRequest {
	s.Comment = &v
	return s
}

func (s *UpdateTopicRequest) SetProjectName(v string) *UpdateTopicRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateTopicRequest) SetTopicName(v string) *UpdateTopicRequest {
	s.TopicName = &v
	return s
}

func (s *UpdateTopicRequest) Validate() error {
	return dara.Validate(s)
}
