// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserGroupMembersResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddUserGroupMembersResponseBody
	GetResult() *bool
	SetSuccess(v bool) *AddUserGroupMembersResponseBody
	GetSuccess() *bool
}

type AddUserGroupMembersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
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

func (s AddUserGroupMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserGroupMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserGroupMembersResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddUserGroupMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddUserGroupMembersResponseBody) SetRequestId(v string) *AddUserGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserGroupMembersResponseBody) SetResult(v bool) *AddUserGroupMembersResponseBody {
	s.Result = &v
	return s
}

func (s *AddUserGroupMembersResponseBody) SetSuccess(v bool) *AddUserGroupMembersResponseBody {
	s.Success = &v
	return s
}

func (s *AddUserGroupMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
