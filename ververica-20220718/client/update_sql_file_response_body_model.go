// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SqlFile) *UpdateSqlFileResponseBody
	GetData() *SqlFile
	SetErrorCode(v string) *UpdateSqlFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateSqlFileResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdateSqlFileResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateSqlFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSqlFileResponseBody
	GetSuccess() *bool
}

type UpdateSqlFileResponseBody struct {
	// The complete SQL script information returned upon success. This value is valid when success is true.
	//
	// example:
	//
	// 123
	Data *SqlFile `json:"data,omitempty" xml:"data,omitempty"`
	// The business error code. This value is not empty when success is false, and is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The business error message. This value is not empty when success is false, and is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The business status code, which is always 200. Use success to determine whether the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-ABCD-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the business request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateSqlFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSqlFileResponseBody) GetData() *SqlFile {
	return s.Data
}

func (s *UpdateSqlFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateSqlFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateSqlFileResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateSqlFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSqlFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSqlFileResponseBody) SetData(v *SqlFile) *UpdateSqlFileResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSqlFileResponseBody) SetErrorCode(v string) *UpdateSqlFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateSqlFileResponseBody) SetErrorMessage(v string) *UpdateSqlFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateSqlFileResponseBody) SetHttpCode(v int32) *UpdateSqlFileResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateSqlFileResponseBody) SetRequestId(v string) *UpdateSqlFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSqlFileResponseBody) SetSuccess(v bool) *UpdateSqlFileResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSqlFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
