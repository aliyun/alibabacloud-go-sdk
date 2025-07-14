// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserGroupMemberResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddUserGroupMemberResponseBody
	GetResult() *bool
	SetSuccess(v bool) *AddUserGroupMemberResponseBody
	GetSuccess() *bool
}

type AddUserGroupMemberResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B6141A5A-A9EF-5F16-BF34-EFB9C1CCE4F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of adding members to a user group is returned. Valid values:
	//
	// 	- true: The task is added.
	//
	// 	- false: The tag failed to be added.
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

func (s AddUserGroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserGroupMemberResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddUserGroupMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddUserGroupMemberResponseBody) SetRequestId(v string) *AddUserGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetResult(v bool) *AddUserGroupMemberResponseBody {
	s.Result = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetSuccess(v bool) *AddUserGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
