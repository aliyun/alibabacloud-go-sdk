// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectRollbackFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetDataCorrectRollbackFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataCorrectRollbackFileResponseBody
	GetErrorMessage() *string
	SetFileUrl(v string) *GetDataCorrectRollbackFileResponseBody
	GetFileUrl() *string
	SetRequestId(v string) *GetDataCorrectRollbackFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCorrectRollbackFileResponseBody
	GetSuccess() *bool
}

type GetDataCorrectRollbackFileResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The download URL of the attachment.
	//
	// example:
	//
	// https://dmsxxx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3CDB8601-AD74-4A47-8114-08E08CD6****
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

func (s GetDataCorrectRollbackFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectRollbackFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCorrectRollbackFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataCorrectRollbackFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataCorrectRollbackFileResponseBody) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetDataCorrectRollbackFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCorrectRollbackFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCorrectRollbackFileResponseBody) SetErrorCode(v string) *GetDataCorrectRollbackFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataCorrectRollbackFileResponseBody) SetErrorMessage(v string) *GetDataCorrectRollbackFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataCorrectRollbackFileResponseBody) SetFileUrl(v string) *GetDataCorrectRollbackFileResponseBody {
	s.FileUrl = &v
	return s
}

func (s *GetDataCorrectRollbackFileResponseBody) SetRequestId(v string) *GetDataCorrectRollbackFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCorrectRollbackFileResponseBody) SetSuccess(v bool) *GetDataCorrectRollbackFileResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCorrectRollbackFileResponseBody) Validate() error {
	return dara.Validate(s)
}
