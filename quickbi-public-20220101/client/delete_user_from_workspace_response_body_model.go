// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserFromWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserFromWorkspaceResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteUserFromWorkspaceResponseBody
	GetResult() *bool
	SetSuccess(v bool) *DeleteUserFromWorkspaceResponseBody
	GetSuccess() *bool
}

type DeleteUserFromWorkspaceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the result of the API execution. Possible values:
	//
	// - true: Execution succeeded
	//
	// - false: Execution failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Value range:
	//
	// - true: The request succeeded - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserFromWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserFromWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserFromWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserFromWorkspaceResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteUserFromWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserFromWorkspaceResponseBody) SetRequestId(v string) *DeleteUserFromWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserFromWorkspaceResponseBody) SetResult(v bool) *DeleteUserFromWorkspaceResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteUserFromWorkspaceResponseBody) SetSuccess(v bool) *DeleteUserFromWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserFromWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
