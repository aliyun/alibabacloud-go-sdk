// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectSQLFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetDataCorrectSQLFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataCorrectSQLFileResponseBody
	GetErrorMessage() *string
	SetFileUrl(v string) *GetDataCorrectSQLFileResponseBody
	GetFileUrl() *string
	SetRequestId(v string) *GetDataCorrectSQLFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCorrectSQLFileResponseBody
	GetSuccess() *bool
}

type GetDataCorrectSQLFileResponseBody struct {
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
	// The download URL of the SQL script.
	//
	// example:
	//
	// https://dmsxxx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 283C461F-11D8-48AA-B695-DF092DA32AF3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataCorrectSQLFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectSQLFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCorrectSQLFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataCorrectSQLFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataCorrectSQLFileResponseBody) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetDataCorrectSQLFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCorrectSQLFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCorrectSQLFileResponseBody) SetErrorCode(v string) *GetDataCorrectSQLFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataCorrectSQLFileResponseBody) SetErrorMessage(v string) *GetDataCorrectSQLFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataCorrectSQLFileResponseBody) SetFileUrl(v string) *GetDataCorrectSQLFileResponseBody {
	s.FileUrl = &v
	return s
}

func (s *GetDataCorrectSQLFileResponseBody) SetRequestId(v string) *GetDataCorrectSQLFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCorrectSQLFileResponseBody) SetSuccess(v bool) *GetDataCorrectSQLFileResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCorrectSQLFileResponseBody) Validate() error {
	return dara.Validate(s)
}
