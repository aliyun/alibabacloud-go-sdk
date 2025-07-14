// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserGroupMemberResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteUserGroupMemberResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DeleteUserGroupMemberResponseBody
	GetSuccess() *bool
}

type DeleteUserGroupMemberResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the result of deleting a user group member. Valid values:
	//
	// 	- true: The task is deleted.
	//
	// 	- false: The deletion failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserGroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserGroupMemberResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteUserGroupMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserGroupMemberResponseBody) SetRequestId(v string) *DeleteUserGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserGroupMemberResponseBody) SetResult(v bool) *DeleteUserGroupMemberResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteUserGroupMemberResponseBody) SetSuccess(v bool) *DeleteUserGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserGroupMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
