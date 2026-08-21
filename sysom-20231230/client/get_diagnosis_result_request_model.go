// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiagnosisResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetDiagnosisResultRequest
	GetXDebugId() *string
	SetTaskId(v string) *GetDiagnosisResultRequest
	GetTaskId() *string
	SetXSysomInvokeSource(v string) *GetDiagnosisResultRequest
	GetXSysomInvokeSource() *string
}

type GetDiagnosisResultRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The diagnostic task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// quzuYl23
	TaskId             *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetDiagnosisResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDiagnosisResultRequest) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetDiagnosisResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDiagnosisResultRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetDiagnosisResultRequest) SetXDebugId(v string) *GetDiagnosisResultRequest {
	s.XDebugId = &v
	return s
}

func (s *GetDiagnosisResultRequest) SetTaskId(v string) *GetDiagnosisResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetDiagnosisResultRequest) SetXSysomInvokeSource(v string) *GetDiagnosisResultRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetDiagnosisResultRequest) Validate() error {
	return dara.Validate(s)
}
