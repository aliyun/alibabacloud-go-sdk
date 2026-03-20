// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *DeleteTopicRequest
	GetProjectName() *string
	SetTopicName(v string) *DeleteTopicRequest
	GetTopicName() *string
}

type DeleteTopicRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// feeding
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// liuywdb01_hh1
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s DeleteTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteTopicRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeleteTopicRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *DeleteTopicRequest) SetProjectName(v string) *DeleteTopicRequest {
	s.ProjectName = &v
	return s
}

func (s *DeleteTopicRequest) SetTopicName(v string) *DeleteTopicRequest {
	s.TopicName = &v
	return s
}

func (s *DeleteTopicRequest) Validate() error {
	return dara.Validate(s)
}
