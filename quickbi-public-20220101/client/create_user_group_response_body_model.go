// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUserGroupResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateUserGroupResponseBody
	GetResult() *string
	SetSuccess(v bool) *CreateUserGroupResponseBody
	GetSuccess() *bool
}

type CreateUserGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 36829379-0C38-5BC0-830A-92665BF77D4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the added user group is returned. An empty string \\"\\" is returned if the add fails.
	//
	// example:
	//
	// f5eeb52e-d9c2-4a8b-80e3-47ab55c2****
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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

func (s CreateUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserGroupResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateUserGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUserGroupResponseBody) SetRequestId(v string) *CreateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetResult(v string) *CreateUserGroupResponseBody {
	s.Result = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetSuccess(v bool) *CreateUserGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
