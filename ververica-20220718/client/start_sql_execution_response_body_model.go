// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSqlExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *StartSqlExecutionResult) *StartSqlExecutionResponseBody
	GetData() *StartSqlExecutionResult
	SetErrorCode(v string) *StartSqlExecutionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StartSqlExecutionResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *StartSqlExecutionResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *StartSqlExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartSqlExecutionResponseBody
	GetSuccess() *bool
}

type StartSqlExecutionResponseBody struct {
	// The complete information of the SQL script returned upon success. This parameter is valid when success is true.
	//
	// example:
	//
	// 如返回示例所示
	Data *StartSqlExecutionResult `json:"data,omitempty" xml:"data,omitempty"`
	// The business error code. This parameter is not empty when success is false. This parameter is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The business error message. This parameter is not empty when success is false. This parameter is empty when success is true.
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
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the business request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StartSqlExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSqlExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartSqlExecutionResponseBody) GetData() *StartSqlExecutionResult {
	return s.Data
}

func (s *StartSqlExecutionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StartSqlExecutionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StartSqlExecutionResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *StartSqlExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSqlExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartSqlExecutionResponseBody) SetData(v *StartSqlExecutionResult) *StartSqlExecutionResponseBody {
	s.Data = v
	return s
}

func (s *StartSqlExecutionResponseBody) SetErrorCode(v string) *StartSqlExecutionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartSqlExecutionResponseBody) SetErrorMessage(v string) *StartSqlExecutionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartSqlExecutionResponseBody) SetHttpCode(v int32) *StartSqlExecutionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StartSqlExecutionResponseBody) SetRequestId(v string) *StartSqlExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSqlExecutionResponseBody) SetSuccess(v bool) *StartSqlExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *StartSqlExecutionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
