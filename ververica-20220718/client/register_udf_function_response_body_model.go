// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterUdfFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UdfFunction) *RegisterUdfFunctionResponseBody
	GetData() *UdfFunction
	SetErrorCode(v string) *RegisterUdfFunctionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RegisterUdfFunctionResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *RegisterUdfFunctionResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *RegisterUdfFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RegisterUdfFunctionResponseBody
	GetSuccess() *bool
}

type RegisterUdfFunctionResponseBody struct {
	// The information about the UDF.
	Data *UdfFunction `json:"data,omitempty" xml:"data,omitempty"`
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
	// The status code returned. The value was fixed to 200. The status code 200 indicates that the request was successful.
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

func (s RegisterUdfFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterUdfFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterUdfFunctionResponseBody) GetData() *UdfFunction {
	return s.Data
}

func (s *RegisterUdfFunctionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RegisterUdfFunctionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RegisterUdfFunctionResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *RegisterUdfFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterUdfFunctionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RegisterUdfFunctionResponseBody) SetData(v *UdfFunction) *RegisterUdfFunctionResponseBody {
	s.Data = v
	return s
}

func (s *RegisterUdfFunctionResponseBody) SetErrorCode(v string) *RegisterUdfFunctionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RegisterUdfFunctionResponseBody) SetErrorMessage(v string) *RegisterUdfFunctionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RegisterUdfFunctionResponseBody) SetHttpCode(v int32) *RegisterUdfFunctionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *RegisterUdfFunctionResponseBody) SetRequestId(v string) *RegisterUdfFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterUdfFunctionResponseBody) SetSuccess(v bool) *RegisterUdfFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *RegisterUdfFunctionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
