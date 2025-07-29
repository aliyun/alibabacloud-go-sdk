// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTagMiningAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetTagMiningAnalysisTaskRequest
	GetTaskId() *string
}

type GetTagMiningAnalysisTaskRequest struct {
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTagMiningAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTagMiningAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTagMiningAnalysisTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTagMiningAnalysisTaskRequest) SetTaskId(v string) *GetTagMiningAnalysisTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetTagMiningAnalysisTaskRequest) Validate() error {
	return dara.Validate(s)
}
