// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserGroupMembersResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteUserGroupMembersResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DeleteUserGroupMembersResponseBody
	GetSuccess() *bool
}

type DeleteUserGroupMembersResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// ABBAD906-****-5D18-B23D-****53AB0AA2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the result of the interface execution. Possible values:
	//
	// - true: Execution succeeded
	//
	// - false: Execution failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserGroupMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserGroupMembersResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteUserGroupMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserGroupMembersResponseBody) SetRequestId(v string) *DeleteUserGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserGroupMembersResponseBody) SetResult(v bool) *DeleteUserGroupMembersResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteUserGroupMembersResponseBody) SetSuccess(v bool) *DeleteUserGroupMembersResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserGroupMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
