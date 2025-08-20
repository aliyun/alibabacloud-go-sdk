// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiagnosticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDiagnosticResponseBody
	GetCode() *string
	SetDiagnosticKey(v string) *GetDiagnosticResponseBody
	GetDiagnosticKey() *string
	SetDiagnosticProduct(v string) *GetDiagnosticResponseBody
	GetDiagnosticProduct() *string
	SetDiagnosticResult(v *GetDiagnosticResponseBodyDiagnosticResult) *GetDiagnosticResponseBody
	GetDiagnosticResult() *GetDiagnosticResponseBodyDiagnosticResult
	SetDiagnosticTime(v string) *GetDiagnosticResponseBody
	GetDiagnosticTime() *string
	SetHttpCode(v string) *GetDiagnosticResponseBody
	GetHttpCode() *string
	SetHttpStatusCode(v int32) *GetDiagnosticResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDiagnosticResponseBody
	GetMessage() *string
	SetRecommends(v map[string]interface{}) *GetDiagnosticResponseBody
	GetRecommends() map[string]interface{}
	SetReportId(v string) *GetDiagnosticResponseBody
	GetReportId() *string
	SetRequestId(v string) *GetDiagnosticResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetDiagnosticResponseBody
	GetStatus() *string
	SetStatusReason(v string) *GetDiagnosticResponseBody
	GetStatusReason() *string
	SetSuccess(v string) *GetDiagnosticResponseBody
	GetSuccess() *string
}

type GetDiagnosticResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// Forbidden
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The keyword in the diagnosis.
	//
	// example:
	//
	// 0938F60C-A2CA-5A2E-9983-03EB1E6D7AE2
	DiagnosticKey *string `json:"DiagnosticKey,omitempty" xml:"DiagnosticKey,omitempty"`
	// The name of the diagnostic item.
	//
	// example:
	//
	// ros
	DiagnosticProduct *string `json:"DiagnosticProduct,omitempty" xml:"DiagnosticProduct,omitempty"`
	// The diagnosis result.
	DiagnosticResult *GetDiagnosticResponseBodyDiagnosticResult `json:"DiagnosticResult,omitempty" xml:"DiagnosticResult,omitempty" type:"Struct"`
	// The time when the diagnosis was performed.
	//
	// example:
	//
	// 2023-03-27T03:32:03Z
	DiagnosticTime *string `json:"DiagnosticTime,omitempty" xml:"DiagnosticTime,omitempty"`
	// The HTTP status code
	//
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
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
	// The specified parameter ReportId is invalid, Can not find diagnostic report dr-5f6135782f104b0f****.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The suggestion for the diagnosis.
	//
	// example:
	//
	// {\\"RosActionMessages\\": {\\"Reason\\": \\"Fail to delete stack (fc966920-450a-456b-983c-deeaec8e****), as deletion protection is enabled.\\", \\"Recommend\\": \\"\\"}}
	Recommends map[string]interface{} `json:"Recommends,omitempty" xml:"Recommends,omitempty"`
	// The ID of the diagnostic report.
	//
	// example:
	//
	// dr-cc80afc48c8741e9****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 96A24844-9603-5E79-BDF4-EFD412FC5D4C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The diagnosis status. Valid values:
	//
	// 	- Running: The diagnosis is in progress.
	//
	// 	- Complete: The diagnosis is complete.
	//
	// 	- Failed: The diagnosis failed.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason for the diagnosis status.
	//
	// example:
	//
	// Complete
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
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

func (s GetDiagnosticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDiagnosticResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiagnosticResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDiagnosticResponseBody) GetDiagnosticKey() *string {
	return s.DiagnosticKey
}

func (s *GetDiagnosticResponseBody) GetDiagnosticProduct() *string {
	return s.DiagnosticProduct
}

func (s *GetDiagnosticResponseBody) GetDiagnosticResult() *GetDiagnosticResponseBodyDiagnosticResult {
	return s.DiagnosticResult
}

func (s *GetDiagnosticResponseBody) GetDiagnosticTime() *string {
	return s.DiagnosticTime
}

func (s *GetDiagnosticResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *GetDiagnosticResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDiagnosticResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDiagnosticResponseBody) GetRecommends() map[string]interface{} {
	return s.Recommends
}

func (s *GetDiagnosticResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *GetDiagnosticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDiagnosticResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetDiagnosticResponseBody) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetDiagnosticResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetDiagnosticResponseBody) SetCode(v string) *GetDiagnosticResponseBody {
	s.Code = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetDiagnosticKey(v string) *GetDiagnosticResponseBody {
	s.DiagnosticKey = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetDiagnosticProduct(v string) *GetDiagnosticResponseBody {
	s.DiagnosticProduct = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetDiagnosticResult(v *GetDiagnosticResponseBodyDiagnosticResult) *GetDiagnosticResponseBody {
	s.DiagnosticResult = v
	return s
}

func (s *GetDiagnosticResponseBody) SetDiagnosticTime(v string) *GetDiagnosticResponseBody {
	s.DiagnosticTime = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetHttpCode(v string) *GetDiagnosticResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetHttpStatusCode(v int32) *GetDiagnosticResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetMessage(v string) *GetDiagnosticResponseBody {
	s.Message = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetRecommends(v map[string]interface{}) *GetDiagnosticResponseBody {
	s.Recommends = v
	return s
}

func (s *GetDiagnosticResponseBody) SetReportId(v string) *GetDiagnosticResponseBody {
	s.ReportId = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetRequestId(v string) *GetDiagnosticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetStatus(v string) *GetDiagnosticResponseBody {
	s.Status = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetStatusReason(v string) *GetDiagnosticResponseBody {
	s.StatusReason = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetSuccess(v string) *GetDiagnosticResponseBody {
	s.Success = &v
	return s
}

func (s *GetDiagnosticResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDiagnosticResponseBodyDiagnosticResult struct {
	// The resources that failed to be diagnosed.
	FailedResources map[string]interface{} `json:"FailedResources,omitempty" xml:"FailedResources,omitempty"`
	// The information about Resource Orchestration Service (ROS) calling.
	RosActionMessages map[string]interface{} `json:"RosActionMessages,omitempty" xml:"RosActionMessages,omitempty"`
	// The stack information.
	StackMessages map[string]interface{} `json:"StackMessages,omitempty" xml:"StackMessages,omitempty"`
}

func (s GetDiagnosticResponseBodyDiagnosticResult) String() string {
	return dara.Prettify(s)
}

func (s GetDiagnosticResponseBodyDiagnosticResult) GoString() string {
	return s.String()
}

func (s *GetDiagnosticResponseBodyDiagnosticResult) GetFailedResources() map[string]interface{} {
	return s.FailedResources
}

func (s *GetDiagnosticResponseBodyDiagnosticResult) GetRosActionMessages() map[string]interface{} {
	return s.RosActionMessages
}

func (s *GetDiagnosticResponseBodyDiagnosticResult) GetStackMessages() map[string]interface{} {
	return s.StackMessages
}

func (s *GetDiagnosticResponseBodyDiagnosticResult) SetFailedResources(v map[string]interface{}) *GetDiagnosticResponseBodyDiagnosticResult {
	s.FailedResources = v
	return s
}

func (s *GetDiagnosticResponseBodyDiagnosticResult) SetRosActionMessages(v map[string]interface{}) *GetDiagnosticResponseBodyDiagnosticResult {
	s.RosActionMessages = v
	return s
}

func (s *GetDiagnosticResponseBodyDiagnosticResult) SetStackMessages(v map[string]interface{}) *GetDiagnosticResponseBodyDiagnosticResult {
	s.StackMessages = v
	return s
}

func (s *GetDiagnosticResponseBodyDiagnosticResult) Validate() error {
	return dara.Validate(s)
}
