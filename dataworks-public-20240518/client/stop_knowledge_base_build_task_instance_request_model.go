// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopKnowledgeBaseBuildTaskInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *StopKnowledgeBaseBuildTaskInstanceRequest
	GetKnowledgeBaseId() *string
	SetTaskInstanceId(v string) *StopKnowledgeBaseBuildTaskInstanceRequest
	GetTaskInstanceId() *string
}

type StopKnowledgeBaseBuildTaskInstanceRequest struct {
	// The knowledge base.
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

func (s StopKnowledgeBaseBuildTaskInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopKnowledgeBaseBuildTaskInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopKnowledgeBaseBuildTaskInstanceRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *StopKnowledgeBaseBuildTaskInstanceRequest) GetTaskInstanceId() *string {
	return s.TaskInstanceId
}

func (s *StopKnowledgeBaseBuildTaskInstanceRequest) SetKnowledgeBaseId(v string) *StopKnowledgeBaseBuildTaskInstanceRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceRequest) SetTaskInstanceId(v string) *StopKnowledgeBaseBuildTaskInstanceRequest {
	s.TaskInstanceId = &v
	return s
}

func (s *StopKnowledgeBaseBuildTaskInstanceRequest) Validate() error {
	return dara.Validate(s)
}
