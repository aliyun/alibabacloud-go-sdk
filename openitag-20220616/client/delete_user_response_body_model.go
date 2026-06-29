// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteUserResponseBody
	GetCode() *int32
	SetDetails(v string) *DeleteUserResponseBody
	GetDetails() *string
	SetErrorCode(v string) *DeleteUserResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteUserResponseBody
	GetSuccess() *bool
}

type DeleteUserResponseBody struct {
	// Total amount of data under the conditions of this request. This parameter is optional and is not returned by default.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Response message of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteUserResponseBody) GetDetails() *string {
	return s.Details
}

func (s *DeleteUserResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserResponseBody) SetCode(v int32) *DeleteUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserResponseBody) SetDetails(v string) *DeleteUserResponseBody {
	s.Details = &v
	return s
}

func (s *DeleteUserResponseBody) SetErrorCode(v string) *DeleteUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUserResponseBody) SetMessage(v string) *DeleteUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetSuccess(v bool) *DeleteUserResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserResponseBody) Validate() error {
	return dara.Validate(s)
}
