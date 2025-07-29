// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnterpriseVocAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetEnterpriseVocAnalysisTaskRequest
	GetTaskId() *string
}

type GetEnterpriseVocAnalysisTaskRequest struct {
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetEnterpriseVocAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEnterpriseVocAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetEnterpriseVocAnalysisTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetEnterpriseVocAnalysisTaskRequest) SetTaskId(v string) *GetEnterpriseVocAnalysisTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetEnterpriseVocAnalysisTaskRequest) Validate() error {
	return dara.Validate(s)
}
