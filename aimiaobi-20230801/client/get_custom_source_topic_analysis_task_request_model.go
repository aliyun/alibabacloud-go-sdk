// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomSourceTopicAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetCustomSourceTopicAnalysisTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetCustomSourceTopicAnalysisTaskRequest
	GetWorkspaceId() *string
}

type GetCustomSourceTopicAnalysisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c9f226b02cca4f42a84c5e955c39dfd2
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetCustomSourceTopicAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomSourceTopicAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetCustomSourceTopicAnalysisTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCustomSourceTopicAnalysisTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetCustomSourceTopicAnalysisTaskRequest) SetTaskId(v string) *GetCustomSourceTopicAnalysisTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskRequest) SetWorkspaceId(v string) *GetCustomSourceTopicAnalysisTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetCustomSourceTopicAnalysisTaskRequest) Validate() error {
	return dara.Validate(s)
}
