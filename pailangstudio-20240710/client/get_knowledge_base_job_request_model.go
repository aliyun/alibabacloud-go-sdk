// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *GetKnowledgeBaseJobRequest
	GetWorkspaceId() *string
}

type GetKnowledgeBaseJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 478***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetKnowledgeBaseJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseJobRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetKnowledgeBaseJobRequest) SetWorkspaceId(v string) *GetKnowledgeBaseJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetKnowledgeBaseJobRequest) Validate() error {
	return dara.Validate(s)
}
