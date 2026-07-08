// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomTopicSelectionPerspectiveAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest
	GetAgentKey() *string
	SetTaskId(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest
	GetTaskId() *string
}

type GetCustomTopicSelectionPerspectiveAnalysisTaskRequest struct {
	// Unique identifier of the workspace: [AgentKey](https://help.aliyun.com/document_detail/2587494.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// Unique ID of the task.
	//
	// > The system generates a TaskId by default. If you specify the same TaskId for multiple tasks, those tasks belong to the same conversation group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0dbf1055f8a2475d99904c3b76a0ffba
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) SetAgentKey(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) SetTaskId(v string) *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetCustomTopicSelectionPerspectiveAnalysisTaskRequest) Validate() error {
	return dara.Validate(s)
}
