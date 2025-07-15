// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicSelectionPerspectiveAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetTopicSelectionPerspectiveAnalysisTaskRequest
	GetAgentKey() *string
	SetTaskId(v string) *GetTopicSelectionPerspectiveAnalysisTaskRequest
	GetTaskId() *string
}

type GetTopicSelectionPerspectiveAnalysisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// c9f226b02cca4f42a84c5e955c39dfd2
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTopicSelectionPerspectiveAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTopicSelectionPerspectiveAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskRequest) SetAgentKey(v string) *GetTopicSelectionPerspectiveAnalysisTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskRequest) SetTaskId(v string) *GetTopicSelectionPerspectiveAnalysisTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetTopicSelectionPerspectiveAnalysisTaskRequest) Validate() error {
	return dara.Validate(s)
}
