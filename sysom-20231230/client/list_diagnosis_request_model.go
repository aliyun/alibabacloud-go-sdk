// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *ListDiagnosisRequest
	GetXDebugId() *string
	SetCurrent(v int64) *ListDiagnosisRequest
	GetCurrent() *int64
	SetPageSize(v int64) *ListDiagnosisRequest
	GetPageSize() *int64
	SetParams(v string) *ListDiagnosisRequest
	GetParams() *string
	SetServiceName(v string) *ListDiagnosisRequest
	GetServiceName() *string
	SetStatus(v string) *ListDiagnosisRequest
	GetStatus() *string
	SetXSysomInvokeSource(v string) *ListDiagnosisRequest
	GetXSysomInvokeSource() *string
}

type ListDiagnosisRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The diagnostic parameters. Different diagnostic types require different diagnostic parameters. You can use this field to filter records whose parameters match the specified values.
	//
	// example:
	//
	// [{\\"key\\":\\"region\\",\\"value\\":\\"cn-beijing\\"}]
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// The diagnostic type.
	//
	// example:
	//
	// memgraph
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	// The execution status of the diagnostic task.
	//
	// Valid values:
	//
	// - **Ready**: Ready.
	//
	// - **Running**: Running.
	//
	// - **Success**: Succeeded.
	//
	// - **Fail**: Failed.
	//
	// example:
	//
	// Running
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s ListDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnosisRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *ListDiagnosisRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListDiagnosisRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDiagnosisRequest) GetParams() *string {
	return s.Params
}

func (s *ListDiagnosisRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListDiagnosisRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDiagnosisRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *ListDiagnosisRequest) SetXDebugId(v string) *ListDiagnosisRequest {
	s.XDebugId = &v
	return s
}

func (s *ListDiagnosisRequest) SetCurrent(v int64) *ListDiagnosisRequest {
	s.Current = &v
	return s
}

func (s *ListDiagnosisRequest) SetPageSize(v int64) *ListDiagnosisRequest {
	s.PageSize = &v
	return s
}

func (s *ListDiagnosisRequest) SetParams(v string) *ListDiagnosisRequest {
	s.Params = &v
	return s
}

func (s *ListDiagnosisRequest) SetServiceName(v string) *ListDiagnosisRequest {
	s.ServiceName = &v
	return s
}

func (s *ListDiagnosisRequest) SetStatus(v string) *ListDiagnosisRequest {
	s.Status = &v
	return s
}

func (s *ListDiagnosisRequest) SetXSysomInvokeSource(v string) *ListDiagnosisRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *ListDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}
