// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetDataLakeFunctionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataLakeFunctionResponseBody
	GetErrorMessage() *string
	SetFunction(v *DLFunction) *GetDataLakeFunctionResponseBody
	GetFunction() *DLFunction
	SetRequestId(v string) *GetDataLakeFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataLakeFunctionResponseBody
	GetSuccess() *bool
}

type GetDataLakeFunctionResponseBody struct {
	// The error code that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The details about the function.
	Function *DLFunction `json:"Function,omitempty" xml:"Function,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// D911009F-3E95-5AFD-8CF1-73F7B4F15D6E
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

func (s GetDataLakeFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataLakeFunctionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataLakeFunctionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataLakeFunctionResponseBody) GetFunction() *DLFunction {
	return s.Function
}

func (s *GetDataLakeFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataLakeFunctionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataLakeFunctionResponseBody) SetErrorCode(v string) *GetDataLakeFunctionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataLakeFunctionResponseBody) SetErrorMessage(v string) *GetDataLakeFunctionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataLakeFunctionResponseBody) SetFunction(v *DLFunction) *GetDataLakeFunctionResponseBody {
	s.Function = v
	return s
}

func (s *GetDataLakeFunctionResponseBody) SetRequestId(v string) *GetDataLakeFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataLakeFunctionResponseBody) SetSuccess(v bool) *GetDataLakeFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataLakeFunctionResponseBody) Validate() error {
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}
