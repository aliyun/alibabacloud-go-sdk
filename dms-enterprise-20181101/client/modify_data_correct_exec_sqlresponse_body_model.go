// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataCorrectExecSQLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ModifyDataCorrectExecSQLResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ModifyDataCorrectExecSQLResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ModifyDataCorrectExecSQLResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDataCorrectExecSQLResponseBody
	GetSuccess() *bool
}

type ModifyDataCorrectExecSQLResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// InvalidOrderId
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Specified parameter OrderId is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52****
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

func (s ModifyDataCorrectExecSQLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataCorrectExecSQLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataCorrectExecSQLResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModifyDataCorrectExecSQLResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModifyDataCorrectExecSQLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDataCorrectExecSQLResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDataCorrectExecSQLResponseBody) SetErrorCode(v string) *ModifyDataCorrectExecSQLResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyDataCorrectExecSQLResponseBody) SetErrorMessage(v string) *ModifyDataCorrectExecSQLResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyDataCorrectExecSQLResponseBody) SetRequestId(v string) *ModifyDataCorrectExecSQLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataCorrectExecSQLResponseBody) SetSuccess(v bool) *ModifyDataCorrectExecSQLResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDataCorrectExecSQLResponseBody) Validate() error {
	return dara.Validate(s)
}
