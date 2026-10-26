// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceUserRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWorkspaceUserRoleResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateWorkspaceUserRoleResponseBody
	GetResult() *bool
	SetSuccess(v bool) *UpdateWorkspaceUserRoleResponseBody
	GetSuccess() *bool
}

type UpdateWorkspaceUserRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the API operation. Valid values:
	//
	// - true: The execution succeeded.
	//
	// - false: The execution failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request succeeded. Valid values:
	//
	// - true: The request succeeded.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkspaceUserRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceUserRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUserRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkspaceUserRoleResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateWorkspaceUserRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkspaceUserRoleResponseBody) SetRequestId(v string) *UpdateWorkspaceUserRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceUserRoleResponseBody) SetResult(v bool) *UpdateWorkspaceUserRoleResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateWorkspaceUserRoleResponseBody) SetSuccess(v bool) *UpdateWorkspaceUserRoleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkspaceUserRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
