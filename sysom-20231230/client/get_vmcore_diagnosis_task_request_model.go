// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVmcoreDiagnosisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetVmcoreDiagnosisTaskRequest
	GetXDebugId() *string
	SetTaskId(v string) *GetVmcoreDiagnosisTaskRequest
	GetTaskId() *string
	SetXSysomInvokeSource(v string) *GetVmcoreDiagnosisTaskRequest
	GetXSysomInvokeSource() *string
}

type GetVmcoreDiagnosisTaskRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bbe94a98-4192-4172-b856-95777e0a55d7
	TaskId             *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetVmcoreDiagnosisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVmcoreDiagnosisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetVmcoreDiagnosisTaskRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetVmcoreDiagnosisTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVmcoreDiagnosisTaskRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetVmcoreDiagnosisTaskRequest) SetXDebugId(v string) *GetVmcoreDiagnosisTaskRequest {
	s.XDebugId = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskRequest) SetTaskId(v string) *GetVmcoreDiagnosisTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskRequest) SetXSysomInvokeSource(v string) *GetVmcoreDiagnosisTaskRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetVmcoreDiagnosisTaskRequest) Validate() error {
	return dara.Validate(s)
}
