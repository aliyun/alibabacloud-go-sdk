// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawAllUserGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *WithdrawAllUserGroupsResponseBody
	GetRequestId() *string
	SetResult(v bool) *WithdrawAllUserGroupsResponseBody
	GetResult() *bool
	SetSuccess(v bool) *WithdrawAllUserGroupsResponseBody
	GetSuccess() *bool
}

type WithdrawAllUserGroupsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the interface is returned. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request fails.
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

func (s WithdrawAllUserGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WithdrawAllUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *WithdrawAllUserGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WithdrawAllUserGroupsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *WithdrawAllUserGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *WithdrawAllUserGroupsResponseBody) SetRequestId(v string) *WithdrawAllUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *WithdrawAllUserGroupsResponseBody) SetResult(v bool) *WithdrawAllUserGroupsResponseBody {
	s.Result = &v
	return s
}

func (s *WithdrawAllUserGroupsResponseBody) SetSuccess(v bool) *WithdrawAllUserGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *WithdrawAllUserGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
