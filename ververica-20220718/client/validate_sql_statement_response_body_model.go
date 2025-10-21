// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateSqlStatementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SqlStatementValidationResult) *ValidateSqlStatementResponseBody
	GetData() *SqlStatementValidationResult
	SetErrorCode(v string) *ValidateSqlStatementResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ValidateSqlStatementResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ValidateSqlStatementResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ValidateSqlStatementResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ValidateSqlStatementResponseBody
	GetSuccess() *bool
}

type ValidateSqlStatementResponseBody struct {
	// The returned data, which represents the details of SQL validation results.
	Data *SqlStatementValidationResult `json:"data,omitempty" xml:"data,omitempty"`
	// If the value of success was false, an error code was returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// If the value of success was false, an error message was returned. If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The status code returned. The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-****-1D30-8A4F-882ED4DD5E02
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ValidateSqlStatementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateSqlStatementResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateSqlStatementResponseBody) GetData() *SqlStatementValidationResult {
	return s.Data
}

func (s *ValidateSqlStatementResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ValidateSqlStatementResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ValidateSqlStatementResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ValidateSqlStatementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateSqlStatementResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ValidateSqlStatementResponseBody) SetData(v *SqlStatementValidationResult) *ValidateSqlStatementResponseBody {
	s.Data = v
	return s
}

func (s *ValidateSqlStatementResponseBody) SetErrorCode(v string) *ValidateSqlStatementResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ValidateSqlStatementResponseBody) SetErrorMessage(v string) *ValidateSqlStatementResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ValidateSqlStatementResponseBody) SetHttpCode(v int32) *ValidateSqlStatementResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ValidateSqlStatementResponseBody) SetRequestId(v string) *ValidateSqlStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateSqlStatementResponseBody) SetSuccess(v bool) *ValidateSqlStatementResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateSqlStatementResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
