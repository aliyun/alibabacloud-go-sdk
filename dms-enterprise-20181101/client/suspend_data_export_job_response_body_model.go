// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendDataExportJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SuspendDataExportJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SuspendDataExportJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SuspendDataExportJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendDataExportJobResponseBody
	GetSuccess() *bool
}

type SuspendDataExportJobResponseBody struct {
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
	// FE8EE2F1-4880-46BC-A704-5CF63EAF9A04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendDataExportJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendDataExportJobResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendDataExportJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SuspendDataExportJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SuspendDataExportJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendDataExportJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendDataExportJobResponseBody) SetErrorCode(v string) *SuspendDataExportJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SuspendDataExportJobResponseBody) SetErrorMessage(v string) *SuspendDataExportJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SuspendDataExportJobResponseBody) SetRequestId(v string) *SuspendDataExportJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendDataExportJobResponseBody) SetSuccess(v bool) *SuspendDataExportJobResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendDataExportJobResponseBody) Validate() error {
	return dara.Validate(s)
}
