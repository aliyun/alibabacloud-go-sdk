// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DeleteKnowledgeBaseJobRequest
	GetWorkspaceId() *string
}

type DeleteKnowledgeBaseJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 478***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteKnowledgeBaseJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteKnowledgeBaseJobRequest) SetWorkspaceId(v string) *DeleteKnowledgeBaseJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteKnowledgeBaseJobRequest) Validate() error {
	return dara.Validate(s)
}
