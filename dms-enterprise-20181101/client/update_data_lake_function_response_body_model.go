// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakeFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateDataLakeFunctionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDataLakeFunctionResponseBody
	GetErrorMessage() *string
	SetFunction(v *DLFunction) *UpdateDataLakeFunctionResponseBody
	GetFunction() *DLFunction
	SetRequestId(v string) *UpdateDataLakeFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataLakeFunctionResponseBody
	GetSuccess() *bool
}

type UpdateDataLakeFunctionResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Function     *DLFunction `json:"Function,omitempty" xml:"Function,omitempty"`
	// example:
	//
	// C1D39814-9808-47F8-AFE0-AF167239AC9B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataLakeFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeFunctionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDataLakeFunctionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDataLakeFunctionResponseBody) GetFunction() *DLFunction {
	return s.Function
}

func (s *UpdateDataLakeFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataLakeFunctionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataLakeFunctionResponseBody) SetErrorCode(v string) *UpdateDataLakeFunctionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDataLakeFunctionResponseBody) SetErrorMessage(v string) *UpdateDataLakeFunctionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDataLakeFunctionResponseBody) SetFunction(v *DLFunction) *UpdateDataLakeFunctionResponseBody {
	s.Function = v
	return s
}

func (s *UpdateDataLakeFunctionResponseBody) SetRequestId(v string) *UpdateDataLakeFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataLakeFunctionResponseBody) SetSuccess(v bool) *UpdateDataLakeFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataLakeFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
