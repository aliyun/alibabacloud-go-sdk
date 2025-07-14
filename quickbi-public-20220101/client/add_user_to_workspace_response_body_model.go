// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserToWorkspaceResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddUserToWorkspaceResponseBody
	GetResult() *bool
	SetSuccess(v bool) *AddUserToWorkspaceResponseBody
	GetSuccess() *bool
}

type AddUserToWorkspaceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the result of the interface execution. Value range:
	//
	// - true: Execution successful
	//
	// - false: Execution failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Value range:
	//
	// - true: Request successful
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddUserToWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserToWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserToWorkspaceResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddUserToWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddUserToWorkspaceResponseBody) SetRequestId(v string) *AddUserToWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserToWorkspaceResponseBody) SetResult(v bool) *AddUserToWorkspaceResponseBody {
	s.Result = &v
	return s
}

func (s *AddUserToWorkspaceResponseBody) SetSuccess(v bool) *AddUserToWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *AddUserToWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
