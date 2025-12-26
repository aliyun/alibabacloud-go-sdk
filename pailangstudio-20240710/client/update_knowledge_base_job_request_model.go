// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateKnowledgeBaseJobRequest
	GetDescription() *string
	SetWorkspaceId(v string) *UpdateKnowledgeBaseJobRequest
	GetWorkspaceId() *string
}

type UpdateKnowledgeBaseJobRequest struct {
	// example:
	//
	// This is a description of the knowledge base job.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 285***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateKnowledgeBaseJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseJobRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeBaseJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateKnowledgeBaseJobRequest) SetDescription(v string) *UpdateKnowledgeBaseJobRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseJobRequest) SetWorkspaceId(v string) *UpdateKnowledgeBaseJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateKnowledgeBaseJobRequest) Validate() error {
	return dara.Validate(s)
}
