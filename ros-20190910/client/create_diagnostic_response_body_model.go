// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDiagnosticResponseBody
	GetCode() *string
	SetDiagnosticType(v string) *CreateDiagnosticResponseBody
	GetDiagnosticType() *string
	SetHttpStatusCode(v int32) *CreateDiagnosticResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDiagnosticResponseBody
	GetMessage() *string
	SetReportId(v string) *CreateDiagnosticResponseBody
	GetReportId() *string
	SetRequestId(v string) *CreateDiagnosticResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateDiagnosticResponseBody
	GetSuccess() *string
}

type CreateDiagnosticResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// Forbidden
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The type of the item that is diagnosed.
	//
	// example:
	//
	// Stack
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// You are not authorized to complete this action.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the diagnostic report.
	//
	// example:
	//
	// dr-e94e39a1274d44b6****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 943B24D7-1A67-55A4-B045-818F90693D3A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDiagnosticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDiagnosticResponseBody) GetDiagnosticType() *string {
	return s.DiagnosticType
}

func (s *CreateDiagnosticResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDiagnosticResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDiagnosticResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *CreateDiagnosticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDiagnosticResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateDiagnosticResponseBody) SetCode(v string) *CreateDiagnosticResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDiagnosticResponseBody) SetDiagnosticType(v string) *CreateDiagnosticResponseBody {
	s.DiagnosticType = &v
	return s
}

func (s *CreateDiagnosticResponseBody) SetHttpStatusCode(v int32) *CreateDiagnosticResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDiagnosticResponseBody) SetMessage(v string) *CreateDiagnosticResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDiagnosticResponseBody) SetReportId(v string) *CreateDiagnosticResponseBody {
	s.ReportId = &v
	return s
}

func (s *CreateDiagnosticResponseBody) SetRequestId(v string) *CreateDiagnosticResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiagnosticResponseBody) SetSuccess(v string) *CreateDiagnosticResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDiagnosticResponseBody) Validate() error {
	return dara.Validate(s)
}
