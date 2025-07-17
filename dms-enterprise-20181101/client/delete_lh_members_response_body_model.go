// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLhMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteLhMembersResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteLhMembersResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteLhMembersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteLhMembersResponseBody
	GetSuccess() *bool
}

type DeleteLhMembersResponseBody struct {
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
	// F1C78D32-1AFD-58AD-9DD2-C8A0896969DD
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

func (s DeleteLhMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLhMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLhMembersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteLhMembersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteLhMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLhMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteLhMembersResponseBody) SetErrorCode(v string) *DeleteLhMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteLhMembersResponseBody) SetErrorMessage(v string) *DeleteLhMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteLhMembersResponseBody) SetRequestId(v string) *DeleteLhMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLhMembersResponseBody) SetSuccess(v bool) *DeleteLhMembersResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLhMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
