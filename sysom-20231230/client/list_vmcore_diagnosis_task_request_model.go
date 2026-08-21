// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVmcoreDiagnosisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *ListVmcoreDiagnosisTaskRequest
	GetXDebugId() *string
	SetDays(v int64) *ListVmcoreDiagnosisTaskRequest
	GetDays() *int64
	SetXSysomInvokeSource(v string) *ListVmcoreDiagnosisTaskRequest
	GetXSysomInvokeSource() *string
}

type ListVmcoreDiagnosisTaskRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The number of past days to query tasks for, up to a maximum of 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Days               *int64  `json:"days,omitempty" xml:"days,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s ListVmcoreDiagnosisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVmcoreDiagnosisTaskRequest) GoString() string {
	return s.String()
}

func (s *ListVmcoreDiagnosisTaskRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *ListVmcoreDiagnosisTaskRequest) GetDays() *int64 {
	return s.Days
}

func (s *ListVmcoreDiagnosisTaskRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *ListVmcoreDiagnosisTaskRequest) SetXDebugId(v string) *ListVmcoreDiagnosisTaskRequest {
	s.XDebugId = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskRequest) SetDays(v int64) *ListVmcoreDiagnosisTaskRequest {
	s.Days = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskRequest) SetXSysomInvokeSource(v string) *ListVmcoreDiagnosisTaskRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *ListVmcoreDiagnosisTaskRequest) Validate() error {
	return dara.Validate(s)
}
