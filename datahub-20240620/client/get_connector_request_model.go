// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorId(v string) *GetConnectorRequest
	GetConnectorId() *string
	SetProjectName(v string) *GetConnectorRequest
	GetProjectName() *string
	SetTopicName(v string) *GetConnectorRequest
	GetTopicName() *string
}

type GetConnectorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c5e07a96-5069-4486-87c3-0d281951f772
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
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

func (s GetConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorRequest) GoString() string {
	return s.String()
}

func (s *GetConnectorRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *GetConnectorRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetConnectorRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *GetConnectorRequest) SetConnectorId(v string) *GetConnectorRequest {
	s.ConnectorId = &v
	return s
}

func (s *GetConnectorRequest) SetProjectName(v string) *GetConnectorRequest {
	s.ProjectName = &v
	return s
}

func (s *GetConnectorRequest) SetTopicName(v string) *GetConnectorRequest {
	s.TopicName = &v
	return s
}

func (s *GetConnectorRequest) Validate() error {
	return dara.Validate(s)
}
