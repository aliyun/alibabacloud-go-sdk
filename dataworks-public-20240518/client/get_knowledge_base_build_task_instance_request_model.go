// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseBuildTaskInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *GetKnowledgeBaseBuildTaskInstanceRequest
	GetKnowledgeBaseId() *string
	SetTaskInstanceId(v string) *GetKnowledgeBaseBuildTaskInstanceRequest
	GetTaskInstanceId() *string
}

type GetKnowledgeBaseBuildTaskInstanceRequest struct {
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The execution ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20001
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskInstanceRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeBaseBuildTaskInstanceRequest) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *GetKnowledgeBaseBuildTaskInstanceRequest) SetKnowledgeBaseId(v string) *GetKnowledgeBaseBuildTaskInstanceRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceRequest) SetTaskInstanceId(v string) *GetKnowledgeBaseBuildTaskInstanceRequest {
	s.TaskInstanceId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskInstanceRequest) Validate() error {
	return dara.Validate(s)
}
