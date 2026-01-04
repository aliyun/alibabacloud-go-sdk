// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *GetSchemaRequest
	GetProjectName() *string
	SetTopicName(v string) *GetSchemaRequest
	GetTopicName() *string
	SetVersionId(v string) *GetSchemaRequest
	GetVersionId() *string
}

type GetSchemaRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 0
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetSchemaRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetSchemaRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *GetSchemaRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetSchemaRequest) SetProjectName(v string) *GetSchemaRequest {
	s.ProjectName = &v
	return s
}

func (s *GetSchemaRequest) SetTopicName(v string) *GetSchemaRequest {
	s.TopicName = &v
	return s
}

func (s *GetSchemaRequest) SetVersionId(v string) *GetSchemaRequest {
	s.VersionId = &v
	return s
}

func (s *GetSchemaRequest) Validate() error {
	return dara.Validate(s)
}
