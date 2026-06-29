// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateUserResponseBody
	GetCode() *int32
	SetDetails(v string) *UpdateUserResponseBody
	GetDetails() *string
	SetErrorCode(v string) *UpdateUserResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateUserResponseBody
	GetSuccess() *bool
	SetUserId(v string) *UpdateUserResponseBody
	GetUserId() *string
}

type UpdateUserResponseBody struct {
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
	// User ID.
	//
	// example:
	//
	// 166***980757311
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateUserResponseBody) GetDetails() *string {
	return s.Details
}

func (s *UpdateUserResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUserResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserResponseBody) SetCode(v int32) *UpdateUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserResponseBody) SetDetails(v string) *UpdateUserResponseBody {
	s.Details = &v
	return s
}

func (s *UpdateUserResponseBody) SetErrorCode(v string) *UpdateUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateUserResponseBody) SetMessage(v string) *UpdateUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetSuccess(v bool) *UpdateUserResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUserResponseBody) SetUserId(v string) *UpdateUserResponseBody {
	s.UserId = &v
	return s
}

func (s *UpdateUserResponseBody) Validate() error {
	return dara.Validate(s)
}
