// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseBuildTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *GetKnowledgeBaseBuildTaskRequest
	GetKnowledgeBaseId() *string
	SetTaskId(v string) *GetKnowledgeBaseBuildTaskRequest
	GetTaskId() *string
}

type GetKnowledgeBaseBuildTaskRequest struct {
	// The ID of the associated knowledge base.
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

func (s GetKnowledgeBaseBuildTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeBaseBuildTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetKnowledgeBaseBuildTaskRequest) SetKnowledgeBaseId(v string) *GetKnowledgeBaseBuildTaskRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskRequest) SetTaskId(v string) *GetKnowledgeBaseBuildTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskRequest) Validate() error {
	return dara.Validate(s)
}
