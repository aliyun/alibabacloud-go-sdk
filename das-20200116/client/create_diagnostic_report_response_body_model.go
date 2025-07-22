// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDiagnosticReportResponseBody
	GetCode() *string
	SetData(v string) *CreateDiagnosticReportResponseBody
	GetData() *string
	SetMessage(v string) *CreateDiagnosticReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDiagnosticReportResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateDiagnosticReportResponseBody
	GetSuccess() *string
}

type CreateDiagnosticReportResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 70af71852fcdf2c5dc7b90596e2cf05b
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ac544623-f6ad-45fd-9a74-9be3db65****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDiagnosticReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDiagnosticReportResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateDiagnosticReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDiagnosticReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDiagnosticReportResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateDiagnosticReportResponseBody) SetCode(v string) *CreateDiagnosticReportResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDiagnosticReportResponseBody) SetData(v string) *CreateDiagnosticReportResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDiagnosticReportResponseBody) SetMessage(v string) *CreateDiagnosticReportResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDiagnosticReportResponseBody) SetRequestId(v string) *CreateDiagnosticReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiagnosticReportResponseBody) SetSuccess(v string) *CreateDiagnosticReportResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDiagnosticReportResponseBody) Validate() error {
	return dara.Validate(s)
}
