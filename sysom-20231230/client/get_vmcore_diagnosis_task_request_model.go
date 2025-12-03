// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVmcoreDiagnosisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetVmcoreDiagnosisTaskRequest
	GetTaskId() *string
}

type GetVmcoreDiagnosisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bbe94a98-4192-4172-b856-95777e0a55d7
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetVmcoreDiagnosisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVmcoreDiagnosisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetVmcoreDiagnosisTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVmcoreDiagnosisTaskRequest) SetTaskId(v string) *GetVmcoreDiagnosisTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskRequest) Validate() error {
	return dara.Validate(s)
}
