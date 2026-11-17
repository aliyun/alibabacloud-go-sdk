// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryKnowledgeBaseBuildTaskInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RetryKnowledgeBaseBuildTaskInstanceRequest
	GetClientToken() *string
	SetKnowledgeBaseId(v string) *RetryKnowledgeBaseBuildTaskInstanceRequest
	GetKnowledgeBaseId() *string
	SetTaskInstanceId(v string) *RetryKnowledgeBaseBuildTaskInstanceRequest
	GetTaskInstanceId() *string
}

type RetryKnowledgeBaseBuildTaskInstanceRequest struct {
	// The idempotency key for the retry operation. Keep the original value when you resend the same request, and use a new token when you initiate another retry.
	//
	// example:
	//
	// retry-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the instance to be retried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20001
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
}

func (s RetryKnowledgeBaseBuildTaskInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseBuildTaskInstanceRequest) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseBuildTaskInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RetryKnowledgeBaseBuildTaskInstanceRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RetryKnowledgeBaseBuildTaskInstanceRequest) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *RetryKnowledgeBaseBuildTaskInstanceRequest) SetClientToken(v string) *RetryKnowledgeBaseBuildTaskInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceRequest) SetKnowledgeBaseId(v string) *RetryKnowledgeBaseBuildTaskInstanceRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceRequest) SetTaskInstanceId(v string) *RetryKnowledgeBaseBuildTaskInstanceRequest {
	s.TaskInstanceId = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceRequest) Validate() error {
	return dara.Validate(s)
}
