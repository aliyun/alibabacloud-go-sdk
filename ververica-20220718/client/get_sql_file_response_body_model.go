// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SqlFile) *GetSqlFileResponseBody
	GetData() *SqlFile
	SetErrorCode(v string) *GetSqlFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetSqlFileResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetSqlFileResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetSqlFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSqlFileResponseBody
	GetSuccess() *bool
}

type GetSqlFileResponseBody struct {
	// The complete SQL script information returned when the request succeeds. This field is valid when success is true.
	//
	// example:
	//
	// See the response example
	Data *SqlFile `json:"data,omitempty" xml:"data,omitempty"`
	// The error code returned when success is false. This value is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message returned when success is false. This value is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code, which is always 200. Use the success field to determine whether the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSqlFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSqlFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlFileResponseBody) GetData() *SqlFile {
	return s.Data
}

func (s *GetSqlFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSqlFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSqlFileResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetSqlFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSqlFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSqlFileResponseBody) SetData(v *SqlFile) *GetSqlFileResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlFileResponseBody) SetErrorCode(v string) *GetSqlFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSqlFileResponseBody) SetErrorMessage(v string) *GetSqlFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSqlFileResponseBody) SetHttpCode(v int32) *GetSqlFileResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetSqlFileResponseBody) SetRequestId(v string) *GetSqlFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlFileResponseBody) SetSuccess(v bool) *GetSqlFileResponseBody {
	s.Success = &v
	return s
}

func (s *GetSqlFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
