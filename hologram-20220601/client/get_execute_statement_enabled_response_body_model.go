// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteStatementEnabledResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *GetExecuteStatementEnabledResponseBody
	GetData() *bool
	SetErrorCode(v string) *GetExecuteStatementEnabledResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetExecuteStatementEnabledResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *GetExecuteStatementEnabledResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *GetExecuteStatementEnabledResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetExecuteStatementEnabledResponseBody
	GetSuccess() *bool
}

type GetExecuteStatementEnabledResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ok
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2C2ECDC1-FBAD-14A5-AA4A-96BC787FBDBC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetExecuteStatementEnabledResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteStatementEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecuteStatementEnabledResponseBody) GetData() *bool {
	return s.Data
}

func (s *GetExecuteStatementEnabledResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetExecuteStatementEnabledResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetExecuteStatementEnabledResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetExecuteStatementEnabledResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExecuteStatementEnabledResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetExecuteStatementEnabledResponseBody) SetData(v bool) *GetExecuteStatementEnabledResponseBody {
	s.Data = &v
	return s
}

func (s *GetExecuteStatementEnabledResponseBody) SetErrorCode(v string) *GetExecuteStatementEnabledResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetExecuteStatementEnabledResponseBody) SetErrorMessage(v string) *GetExecuteStatementEnabledResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetExecuteStatementEnabledResponseBody) SetHttpStatusCode(v string) *GetExecuteStatementEnabledResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetExecuteStatementEnabledResponseBody) SetRequestId(v string) *GetExecuteStatementEnabledResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExecuteStatementEnabledResponseBody) SetSuccess(v bool) *GetExecuteStatementEnabledResponseBody {
	s.Success = &v
	return s
}

func (s *GetExecuteStatementEnabledResponseBody) Validate() error {
	return dara.Validate(s)
}
