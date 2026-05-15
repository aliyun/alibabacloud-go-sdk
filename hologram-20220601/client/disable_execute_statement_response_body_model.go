// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableExecuteStatementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DisableExecuteStatementResponseBody
	GetData() *bool
	SetErrorCode(v string) *DisableExecuteStatementResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DisableExecuteStatementResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *DisableExecuteStatementResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *DisableExecuteStatementResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableExecuteStatementResponseBody
	GetSuccess() *bool
}

type DisableExecuteStatementResponseBody struct {
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// InvalidParameterValue
	ErrorCode      *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage   *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// E3F4B2A7-1234-5678-9ABC-DEF012345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DisableExecuteStatementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableExecuteStatementResponseBody) GoString() string {
	return s.String()
}

func (s *DisableExecuteStatementResponseBody) GetData() *bool {
	return s.Data
}

func (s *DisableExecuteStatementResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DisableExecuteStatementResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DisableExecuteStatementResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DisableExecuteStatementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableExecuteStatementResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableExecuteStatementResponseBody) SetData(v bool) *DisableExecuteStatementResponseBody {
	s.Data = &v
	return s
}

func (s *DisableExecuteStatementResponseBody) SetErrorCode(v string) *DisableExecuteStatementResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DisableExecuteStatementResponseBody) SetErrorMessage(v string) *DisableExecuteStatementResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DisableExecuteStatementResponseBody) SetHttpStatusCode(v string) *DisableExecuteStatementResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableExecuteStatementResponseBody) SetRequestId(v string) *DisableExecuteStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableExecuteStatementResponseBody) SetSuccess(v bool) *DisableExecuteStatementResponseBody {
	s.Success = &v
	return s
}

func (s *DisableExecuteStatementResponseBody) Validate() error {
	return dara.Validate(s)
}
