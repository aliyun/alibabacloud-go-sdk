// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateUserResponseBody
	GetCode() *int32
	SetDetails(v string) *CreateUserResponseBody
	GetDetails() *string
	SetErrorCode(v string) *CreateUserResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateUserResponseBody
	GetSuccess() *bool
	SetUserId(v int64) *CreateUserResponseBody
	GetUserId() *int64
}

type CreateUserResponseBody struct {
	// Return code. The default value is 0, indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// null
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
	// Indicates whether the operation succeeded. Valid values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// User ID.
	//
	// example:
	//
	// 166233998075****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateUserResponseBody) GetDetails() *string {
	return s.Details
}

func (s *CreateUserResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUserResponseBody) GetUserId() *int64 {
	return s.UserId
}

func (s *CreateUserResponseBody) SetCode(v int32) *CreateUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserResponseBody) SetDetails(v string) *CreateUserResponseBody {
	s.Details = &v
	return s
}

func (s *CreateUserResponseBody) SetErrorCode(v string) *CreateUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUserResponseBody) SetMessage(v string) *CreateUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetSuccess(v bool) *CreateUserResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUserResponseBody) SetUserId(v int64) *CreateUserResponseBody {
	s.UserId = &v
	return s
}

func (s *CreateUserResponseBody) Validate() error {
	return dara.Validate(s)
}
