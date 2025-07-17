// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDataExportJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RestartDataExportJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RestartDataExportJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RestartDataExportJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartDataExportJobResponseBody
	GetSuccess() *bool
}

type RestartDataExportJobResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartDataExportJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDataExportJobResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDataExportJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RestartDataExportJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RestartDataExportJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDataExportJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartDataExportJobResponseBody) SetErrorCode(v string) *RestartDataExportJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RestartDataExportJobResponseBody) SetErrorMessage(v string) *RestartDataExportJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RestartDataExportJobResponseBody) SetRequestId(v string) *RestartDataExportJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDataExportJobResponseBody) SetSuccess(v bool) *RestartDataExportJobResponseBody {
	s.Success = &v
	return s
}

func (s *RestartDataExportJobResponseBody) Validate() error {
	return dara.Validate(s)
}
