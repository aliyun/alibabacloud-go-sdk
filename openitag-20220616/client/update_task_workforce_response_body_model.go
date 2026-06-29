// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskWorkforceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateTaskWorkforceResponseBody
	GetCode() *int32
	SetDetails(v string) *UpdateTaskWorkforceResponseBody
	GetDetails() *string
	SetErrorCode(v string) *UpdateTaskWorkforceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateTaskWorkforceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTaskWorkforceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskWorkforceResponseBody
	GetSuccess() *bool
}

type UpdateTaskWorkforceResponseBody struct {
	// Return encoding. The default value is 0, indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// ""
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// - When Success is false, returns a business error code.
	//
	// - When Success is true, returns an empty value.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Response message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded. Possible values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTaskWorkforceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskWorkforceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskWorkforceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateTaskWorkforceResponseBody) GetDetails() *string {
	return s.Details
}

func (s *UpdateTaskWorkforceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTaskWorkforceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTaskWorkforceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskWorkforceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskWorkforceResponseBody) SetCode(v int32) *UpdateTaskWorkforceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTaskWorkforceResponseBody) SetDetails(v string) *UpdateTaskWorkforceResponseBody {
	s.Details = &v
	return s
}

func (s *UpdateTaskWorkforceResponseBody) SetErrorCode(v string) *UpdateTaskWorkforceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskWorkforceResponseBody) SetMessage(v string) *UpdateTaskWorkforceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTaskWorkforceResponseBody) SetRequestId(v string) *UpdateTaskWorkforceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskWorkforceResponseBody) SetSuccess(v bool) *UpdateTaskWorkforceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskWorkforceResponseBody) Validate() error {
	return dara.Validate(s)
}
