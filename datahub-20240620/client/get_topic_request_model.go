// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *GetTopicRequest
	GetProjectName() *string
	SetTopicName(v string) *GetTopicRequest
	GetTopicName() *string
}

type GetTopicRequest struct {
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

func (s GetTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTopicRequest) GoString() string {
	return s.String()
}

func (s *GetTopicRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetTopicRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *GetTopicRequest) SetProjectName(v string) *GetTopicRequest {
	s.ProjectName = &v
	return s
}

func (s *GetTopicRequest) SetTopicName(v string) *GetTopicRequest {
	s.TopicName = &v
	return s
}

func (s *GetTopicRequest) Validate() error {
	return dara.Validate(s)
}
