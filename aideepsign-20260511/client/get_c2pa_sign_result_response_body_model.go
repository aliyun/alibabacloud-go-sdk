// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetC2paSignResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetC2paSignResultResponseBody
	GetCode() *string
	SetErrorCode(v string) *GetC2paSignResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetC2paSignResultResponseBody
	GetErrorMessage() *string
	SetFailedCount(v int64) *GetC2paSignResultResponseBody
	GetFailedCount() *int64
	SetHttpStatusCode(v int32) *GetC2paSignResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetC2paSignResultResponseBody
	GetMessage() *string
	SetReportFileUrl(v string) *GetC2paSignResultResponseBody
	GetReportFileUrl() *string
	SetReportObjectKey(v string) *GetC2paSignResultResponseBody
	GetReportObjectKey() *string
	SetRequestId(v string) *GetC2paSignResultResponseBody
	GetRequestId() *string
	SetSignedFileUrl(v string) *GetC2paSignResultResponseBody
	GetSignedFileUrl() *string
	SetStatus(v string) *GetC2paSignResultResponseBody
	GetStatus() *string
	SetSucceededCount(v int64) *GetC2paSignResultResponseBody
	GetSucceededCount() *int64
	SetSuccess(v bool) *GetC2paSignResultResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *GetC2paSignResultResponseBody
	GetTaskId() *string
	SetTaskType(v string) *GetC2paSignResultResponseBody
	GetTaskType() *string
	SetTotalCount(v int64) *GetC2paSignResultResponseBody
	GetTotalCount() *int64
}

type GetC2paSignResultResponseBody struct {
	// The business error code for the API call. OK is returned when the call succeeds.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code for task failure. This parameter is returned only when Status is FAILED.
	//
	// example:
	//
	// C2PA_FORMAT_UNSUPPORTED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message for task failure. This parameter is returned only when Status is FAILED.
	//
	// example:
	//
	// only valid JPEG and PNG source files are supported
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The number of objects that failed to be signed, were skipped, or were terminated in the Bucket task. This parameter is returned only when the Bucket task is in the SUCCEEDED or CANCELED state.
	//
	// example:
	//
	// 2
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The HTTP status code. 200 is returned when the API call succeeds.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The additional information. success is returned when the API call succeeds.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The temporary download URL returned when a Bucket task enters the CANCELED state and the JSON report is successfully generated. The report covers only the objects discovered before termination.
	//
	// example:
	//
	// https://example-bucket.oss-cn-hangzhou.aliyuncs.com/c2pa_report.json?Expires=1788000000&Signature=***
	ReportFileUrl *string `json:"ReportFileUrl,omitempty" xml:"ReportFileUrl,omitempty"`
	// The object key of the final standard JSON report for the Bucket task. This parameter is returned only when the Bucket task succeeds.
	//
	// example:
	//
	// ai-deepsign-signed/047b8305-9572-4eed-b7fb-8588b989b6e5/c2pa_report.json
	ReportObjectKey *string `json:"ReportObjectKey,omitempty" xml:"ReportObjectKey,omitempty"`
	// The request ID of the OpenAPI call.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The temporary download URL of the signature file. This parameter is returned only when Status is SUCCEEDED.
	//
	// example:
	//
	// https://example.oss-cn-hangzhou.aliyuncs.com/...
	SignedFileUrl *string `json:"SignedFileUrl,omitempty" xml:"SignedFileUrl,omitempty"`
	// The task status. Valid values: PROCESSING, SUCCEEDED, and FAILED.
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of successfully signed objects in the Bucket task. This parameter is returned only when the Bucket task is in the SUCCEEDED or CANCELED state.
	//
	// example:
	//
	// 1
	SucceededCount *int64 `json:"SucceededCount,omitempty" xml:"SucceededCount,omitempty"`
	// Indicates whether the query call is successful. This value may still be true even if the task itself fails.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The C2PA signing task ID.
	//
	// example:
	//
	// 7093c29e-f472-4e1e-91c6-a43df7c7382e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values: FILE and BUCKET.
	//
	// example:
	//
	// BUCKET
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The total number of objects scanned by the Bucket task. This parameter is returned only when the Bucket task is in the SUCCEEDED or CANCELED state.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetC2paSignResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetC2paSignResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetC2paSignResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetC2paSignResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetC2paSignResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetC2paSignResultResponseBody) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *GetC2paSignResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetC2paSignResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetC2paSignResultResponseBody) GetReportFileUrl() *string {
	return s.ReportFileUrl
}

func (s *GetC2paSignResultResponseBody) GetReportObjectKey() *string {
	return s.ReportObjectKey
}

func (s *GetC2paSignResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetC2paSignResultResponseBody) GetSignedFileUrl() *string {
	return s.SignedFileUrl
}

func (s *GetC2paSignResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetC2paSignResultResponseBody) GetSucceededCount() *int64 {
	return s.SucceededCount
}

func (s *GetC2paSignResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetC2paSignResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetC2paSignResultResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *GetC2paSignResultResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetC2paSignResultResponseBody) SetCode(v string) *GetC2paSignResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetErrorCode(v string) *GetC2paSignResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetErrorMessage(v string) *GetC2paSignResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetFailedCount(v int64) *GetC2paSignResultResponseBody {
	s.FailedCount = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetHttpStatusCode(v int32) *GetC2paSignResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetMessage(v string) *GetC2paSignResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetReportFileUrl(v string) *GetC2paSignResultResponseBody {
	s.ReportFileUrl = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetReportObjectKey(v string) *GetC2paSignResultResponseBody {
	s.ReportObjectKey = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetRequestId(v string) *GetC2paSignResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetSignedFileUrl(v string) *GetC2paSignResultResponseBody {
	s.SignedFileUrl = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetStatus(v string) *GetC2paSignResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetSucceededCount(v int64) *GetC2paSignResultResponseBody {
	s.SucceededCount = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetSuccess(v bool) *GetC2paSignResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetTaskId(v string) *GetC2paSignResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetTaskType(v string) *GetC2paSignResultResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetC2paSignResultResponseBody) SetTotalCount(v int64) *GetC2paSignResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetC2paSignResultResponseBody) Validate() error {
	return dara.Validate(s)
}
