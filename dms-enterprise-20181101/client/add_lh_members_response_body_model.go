// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLhMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AddLhMembersResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddLhMembersResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddLhMembersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddLhMembersResponseBody
	GetSuccess() *bool
}

type AddLhMembersResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 31853A2B-DC9D-5B39-8492-D2AC8BCF550E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddLhMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLhMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddLhMembersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddLhMembersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddLhMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLhMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddLhMembersResponseBody) SetErrorCode(v string) *AddLhMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddLhMembersResponseBody) SetErrorMessage(v string) *AddLhMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddLhMembersResponseBody) SetRequestId(v string) *AddLhMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLhMembersResponseBody) SetSuccess(v bool) *AddLhMembersResponseBody {
	s.Success = &v
	return s
}

func (s *AddLhMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
