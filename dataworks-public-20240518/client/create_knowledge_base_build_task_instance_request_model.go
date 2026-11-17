// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseBuildTaskInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateKnowledgeBaseBuildTaskInstanceRequest
	GetClientToken() *string
	SetKnowledgeBaseId(v string) *CreateKnowledgeBaseBuildTaskInstanceRequest
	GetKnowledgeBaseId() *string
	SetTaskId(v string) *CreateKnowledgeBaseBuildTaskInstanceRequest
	GetTaskId() *string
}

type CreateKnowledgeBaseBuildTaskInstanceRequest struct {
	// The idempotency key for this creation request. We recommend providing it for automatic retries. Reuse the same value when retrying the same operation.
	//
	// example:
	//
	// run-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The source task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateKnowledgeBaseBuildTaskInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseBuildTaskInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseBuildTaskInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateKnowledgeBaseBuildTaskInstanceRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKnowledgeBaseBuildTaskInstanceRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateKnowledgeBaseBuildTaskInstanceRequest) SetClientToken(v string) *CreateKnowledgeBaseBuildTaskInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceRequest) SetKnowledgeBaseId(v string) *CreateKnowledgeBaseBuildTaskInstanceRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceRequest) SetTaskId(v string) *CreateKnowledgeBaseBuildTaskInstanceRequest {
	s.TaskId = &v
	return s
}

func (s *CreateKnowledgeBaseBuildTaskInstanceRequest) Validate() error {
	return dara.Validate(s)
}
