// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryReportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RetryReportTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RetryReportTaskResponseBody
	GetErrorMessage() *string
	SetOutRequestNo(v string) *RetryReportTaskResponseBody
	GetOutRequestNo() *string
	SetRetryAvailable(v bool) *RetryReportTaskResponseBody
	GetRetryAvailable() *bool
	SetRetryCount(v int32) *RetryReportTaskResponseBody
	GetRetryCount() *int32
	SetTaskStatus(v string) *RetryReportTaskResponseBody
	GetTaskStatus() *string
}

type RetryReportTaskResponseBody struct {
	ErrorCode      *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage   *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	OutRequestNo   *string `json:"outRequestNo,omitempty" xml:"outRequestNo,omitempty"`
	RetryAvailable *bool   `json:"retryAvailable,omitempty" xml:"retryAvailable,omitempty"`
	RetryCount     *int32  `json:"retryCount,omitempty" xml:"retryCount,omitempty"`
	TaskStatus     *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s RetryReportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryReportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RetryReportTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RetryReportTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RetryReportTaskResponseBody) GetOutRequestNo() *string {
	return s.OutRequestNo
}

func (s *RetryReportTaskResponseBody) GetRetryAvailable() *bool {
	return s.RetryAvailable
}

func (s *RetryReportTaskResponseBody) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *RetryReportTaskResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *RetryReportTaskResponseBody) SetErrorCode(v string) *RetryReportTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RetryReportTaskResponseBody) SetErrorMessage(v string) *RetryReportTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RetryReportTaskResponseBody) SetOutRequestNo(v string) *RetryReportTaskResponseBody {
	s.OutRequestNo = &v
	return s
}

func (s *RetryReportTaskResponseBody) SetRetryAvailable(v bool) *RetryReportTaskResponseBody {
	s.RetryAvailable = &v
	return s
}

func (s *RetryReportTaskResponseBody) SetRetryCount(v int32) *RetryReportTaskResponseBody {
	s.RetryCount = &v
	return s
}

func (s *RetryReportTaskResponseBody) SetTaskStatus(v string) *RetryReportTaskResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *RetryReportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
