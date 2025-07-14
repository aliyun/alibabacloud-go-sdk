// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserGroupResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateUserGroupResponseBody
	GetResult() *bool
	SetSuccess(v bool) *UpdateUserGroupResponseBody
	GetSuccess() *bool
}

type UpdateUserGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4AEF8C5C-D5D2-55D3-BB2F-9D3AA1B6F4FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the interface is successfully executed. Valid values:
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

func (s UpdateUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserGroupResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateUserGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUserGroupResponseBody) SetRequestId(v string) *UpdateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetResult(v bool) *UpdateUserGroupResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetSuccess(v bool) *UpdateUserGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
