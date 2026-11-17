// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseBuildTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *GetKnowledgeBaseBuildTaskResultRequest
	GetKnowledgeBaseId() *string
	SetTaskInstanceId(v string) *GetKnowledgeBaseBuildTaskResultRequest
	GetTaskInstanceId() *string
}

type GetKnowledgeBaseBuildTaskResultRequest struct {
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The ID of the successfully executed instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20001
	TaskInstanceId *string `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
}

func (s GetKnowledgeBaseBuildTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseBuildTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseBuildTaskResultRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgeBaseBuildTaskResultRequest) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *GetKnowledgeBaseBuildTaskResultRequest) SetKnowledgeBaseId(v string) *GetKnowledgeBaseBuildTaskResultRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultRequest) SetTaskInstanceId(v string) *GetKnowledgeBaseBuildTaskResultRequest {
	s.TaskInstanceId = &v
	return s
}

func (s *GetKnowledgeBaseBuildTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
