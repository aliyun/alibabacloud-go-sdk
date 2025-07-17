// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateDataLakeFunctionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataLakeFunctionResponseBody
	GetErrorMessage() *string
	SetFunction(v *DLFunction) *CreateDataLakeFunctionResponseBody
	GetFunction() *DLFunction
	SetRequestId(v string) *CreateDataLakeFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataLakeFunctionResponseBody
	GetSuccess() *bool
}

type CreateDataLakeFunctionResponseBody struct {
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
	// EE214ECD-4330-503A-82F0-FFB039757DC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataLakeFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataLakeFunctionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataLakeFunctionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataLakeFunctionResponseBody) GetFunction() *DLFunction {
	return s.Function
}

func (s *CreateDataLakeFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataLakeFunctionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataLakeFunctionResponseBody) SetErrorCode(v string) *CreateDataLakeFunctionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataLakeFunctionResponseBody) SetErrorMessage(v string) *CreateDataLakeFunctionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataLakeFunctionResponseBody) SetFunction(v *DLFunction) *CreateDataLakeFunctionResponseBody {
	s.Function = v
	return s
}

func (s *CreateDataLakeFunctionResponseBody) SetRequestId(v string) *CreateDataLakeFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataLakeFunctionResponseBody) SetSuccess(v bool) *CreateDataLakeFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataLakeFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
