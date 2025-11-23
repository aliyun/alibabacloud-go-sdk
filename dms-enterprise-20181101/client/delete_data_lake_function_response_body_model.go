// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteDataLakeFunctionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDataLakeFunctionResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDataLakeFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataLakeFunctionResponseBody
	GetSuccess() *bool
}

type DeleteDataLakeFunctionResponseBody struct {
	// The error code.
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
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// B4B07137-F6AE-4756-8474-7F92BB6C4E04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataLakeFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeFunctionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDataLakeFunctionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDataLakeFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataLakeFunctionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataLakeFunctionResponseBody) SetErrorCode(v string) *DeleteDataLakeFunctionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDataLakeFunctionResponseBody) SetErrorMessage(v string) *DeleteDataLakeFunctionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDataLakeFunctionResponseBody) SetRequestId(v string) *DeleteDataLakeFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataLakeFunctionResponseBody) SetSuccess(v bool) *DeleteDataLakeFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataLakeFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
