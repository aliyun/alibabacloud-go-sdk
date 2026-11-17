// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseBuildTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *DeleteKnowledgeBaseBuildTaskRequest
	GetKnowledgeBaseId() *string
	SetTaskId(v string) *DeleteKnowledgeBaseBuildTaskRequest
	GetTaskId() *string
}

type DeleteKnowledgeBaseBuildTaskRequest struct {
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteKnowledgeBaseBuildTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseBuildTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseBuildTaskRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DeleteKnowledgeBaseBuildTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteKnowledgeBaseBuildTaskRequest) SetKnowledgeBaseId(v string) *DeleteKnowledgeBaseBuildTaskRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DeleteKnowledgeBaseBuildTaskRequest) SetTaskId(v string) *DeleteKnowledgeBaseBuildTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteKnowledgeBaseBuildTaskRequest) Validate() error {
	return dara.Validate(s)
}
