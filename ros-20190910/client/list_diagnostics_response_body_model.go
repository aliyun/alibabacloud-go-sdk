// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnosticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnostics(v []*ListDiagnosticsResponseBodyDiagnostics) *ListDiagnosticsResponseBody
	GetDiagnostics() []*ListDiagnosticsResponseBodyDiagnostics
	SetHttpStatusCode(v int32) *ListDiagnosticsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDiagnosticsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListDiagnosticsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDiagnosticsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListDiagnosticsResponseBody
	GetSuccess() *string
}

type ListDiagnosticsResponseBody struct {
	// The items that are diagnosed.
	Diagnostics []*ListDiagnosticsResponseBodyDiagnostics `json:"Diagnostics,omitempty" xml:"Diagnostics,omitempty" type:"Repeated"`
	// The HTTP status code returned. The value 200 indicates that the request was successful.
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
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f01****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1983D1C4-88EA-5D7D-90AB-467D01867A5D
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

func (s ListDiagnosticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnosticsResponseBody) GetDiagnostics() []*ListDiagnosticsResponseBodyDiagnostics {
	return s.Diagnostics
}

func (s *ListDiagnosticsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDiagnosticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDiagnosticsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDiagnosticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDiagnosticsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListDiagnosticsResponseBody) SetDiagnostics(v []*ListDiagnosticsResponseBodyDiagnostics) *ListDiagnosticsResponseBody {
	s.Diagnostics = v
	return s
}

func (s *ListDiagnosticsResponseBody) SetHttpStatusCode(v int32) *ListDiagnosticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDiagnosticsResponseBody) SetMessage(v string) *ListDiagnosticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDiagnosticsResponseBody) SetNextToken(v string) *ListDiagnosticsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDiagnosticsResponseBody) SetRequestId(v string) *ListDiagnosticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnosticsResponseBody) SetSuccess(v string) *ListDiagnosticsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDiagnosticsResponseBody) Validate() error {
	if s.Diagnostics != nil {
		for _, item := range s.Diagnostics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDiagnosticsResponseBodyDiagnostics struct {
	// The time when the diagnostic report was generated.
	//
	// example:
	//
	// 2022-08-01T02:23:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The keyword in the diagnosis.
	//
	// example:
	//
	// 047D84D9-D3EB-5DA8-87F1-9A7DD5598A5D
	DiagnosticKey *string `json:"DiagnosticKey,omitempty" xml:"DiagnosticKey,omitempty"`
	// The product that is diagnosed.
	//
	// example:
	//
	// ros
	DiagnosticProduct *string `json:"DiagnosticProduct,omitempty" xml:"DiagnosticProduct,omitempty"`
	// The ID of the diagnostic report.
	//
	// example:
	//
	// dr-2963bfbcac834f1a****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The diagnosis status.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDiagnosticsResponseBodyDiagnostics) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosticsResponseBodyDiagnostics) GoString() string {
	return s.String()
}

func (s *ListDiagnosticsResponseBodyDiagnostics) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDiagnosticsResponseBodyDiagnostics) GetDiagnosticKey() *string {
	return s.DiagnosticKey
}

func (s *ListDiagnosticsResponseBodyDiagnostics) GetDiagnosticProduct() *string {
	return s.DiagnosticProduct
}

func (s *ListDiagnosticsResponseBodyDiagnostics) GetReportId() *string {
	return s.ReportId
}

func (s *ListDiagnosticsResponseBodyDiagnostics) GetStatus() *string {
	return s.Status
}

func (s *ListDiagnosticsResponseBodyDiagnostics) SetCreateTime(v string) *ListDiagnosticsResponseBodyDiagnostics {
	s.CreateTime = &v
	return s
}

func (s *ListDiagnosticsResponseBodyDiagnostics) SetDiagnosticKey(v string) *ListDiagnosticsResponseBodyDiagnostics {
	s.DiagnosticKey = &v
	return s
}

func (s *ListDiagnosticsResponseBodyDiagnostics) SetDiagnosticProduct(v string) *ListDiagnosticsResponseBodyDiagnostics {
	s.DiagnosticProduct = &v
	return s
}

func (s *ListDiagnosticsResponseBodyDiagnostics) SetReportId(v string) *ListDiagnosticsResponseBodyDiagnostics {
	s.ReportId = &v
	return s
}

func (s *ListDiagnosticsResponseBodyDiagnostics) SetStatus(v string) *ListDiagnosticsResponseBodyDiagnostics {
	s.Status = &v
	return s
}

func (s *ListDiagnosticsResponseBodyDiagnostics) Validate() error {
	return dara.Validate(s)
}
