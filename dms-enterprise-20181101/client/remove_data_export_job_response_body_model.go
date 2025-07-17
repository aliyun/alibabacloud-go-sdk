// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataExportJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RemoveDataExportJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RemoveDataExportJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RemoveDataExportJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveDataExportJobResponseBody
	GetSuccess() *bool
}

type RemoveDataExportJobResponseBody struct {
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
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveDataExportJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataExportJobResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDataExportJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RemoveDataExportJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RemoveDataExportJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDataExportJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveDataExportJobResponseBody) SetErrorCode(v string) *RemoveDataExportJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RemoveDataExportJobResponseBody) SetErrorMessage(v string) *RemoveDataExportJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveDataExportJobResponseBody) SetRequestId(v string) *RemoveDataExportJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDataExportJobResponseBody) SetSuccess(v bool) *RemoveDataExportJobResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveDataExportJobResponseBody) Validate() error {
	return dara.Validate(s)
}
