// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseDataExportJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *PauseDataExportJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *PauseDataExportJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *PauseDataExportJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PauseDataExportJobResponseBody
	GetSuccess() *bool
}

type PauseDataExportJobResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PauseDataExportJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseDataExportJobResponseBody) GoString() string {
	return s.String()
}

func (s *PauseDataExportJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PauseDataExportJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PauseDataExportJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseDataExportJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PauseDataExportJobResponseBody) SetErrorCode(v string) *PauseDataExportJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PauseDataExportJobResponseBody) SetErrorMessage(v string) *PauseDataExportJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PauseDataExportJobResponseBody) SetRequestId(v string) *PauseDataExportJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseDataExportJobResponseBody) SetSuccess(v bool) *PauseDataExportJobResponseBody {
	s.Success = &v
	return s
}

func (s *PauseDataExportJobResponseBody) Validate() error {
	return dara.Validate(s)
}
