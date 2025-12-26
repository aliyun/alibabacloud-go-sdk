// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DeleteKnowledgeBaseRequest
	GetWorkspaceId() *string
}

type DeleteKnowledgeBaseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 478***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteKnowledgeBaseRequest) SetWorkspaceId(v string) *DeleteKnowledgeBaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
