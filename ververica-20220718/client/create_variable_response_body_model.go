// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Variable) *CreateVariableResponseBody
	GetData() *Variable
	SetErrorCode(v string) *CreateVariableResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateVariableResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CreateVariableResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateVariableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVariableResponseBody
	GetSuccess() *bool
}

type CreateVariableResponseBody struct {
	// 	- If the value of success was true, the variable that you created was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data *Variable `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVariableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVariableResponseBody) GetData() *Variable {
	return s.Data
}

func (s *CreateVariableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateVariableResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateVariableResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVariableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVariableResponseBody) SetData(v *Variable) *CreateVariableResponseBody {
	s.Data = v
	return s
}

func (s *CreateVariableResponseBody) SetErrorCode(v string) *CreateVariableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateVariableResponseBody) SetErrorMessage(v string) *CreateVariableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateVariableResponseBody) SetHttpCode(v int32) *CreateVariableResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateVariableResponseBody) SetRequestId(v string) *CreateVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVariableResponseBody) SetSuccess(v bool) *CreateVariableResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVariableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
