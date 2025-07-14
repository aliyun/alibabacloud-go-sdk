// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserGroupResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteUserGroupResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DeleteUserGroupResponseBody
	GetSuccess() *bool
}

type DeleteUserGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F2775AB6-DE99-5FA6-86A4-72EA0A8AFEE3
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

func (s DeleteUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserGroupResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteUserGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserGroupResponseBody) SetRequestId(v string) *DeleteUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetResult(v bool) *DeleteUserGroupResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetSuccess(v bool) *DeleteUserGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
