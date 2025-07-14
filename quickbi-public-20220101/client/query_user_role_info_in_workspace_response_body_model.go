// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserRoleInfoInWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryUserRoleInfoInWorkspaceResponseBody
	GetRequestId() *string
	SetResult(v *QueryUserRoleInfoInWorkspaceResponseBodyResult) *QueryUserRoleInfoInWorkspaceResponseBody
	GetResult() *QueryUserRoleInfoInWorkspaceResponseBodyResult
	SetSuccess(v bool) *QueryUserRoleInfoInWorkspaceResponseBody
	GetSuccess() *bool
}

type QueryUserRoleInfoInWorkspaceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Preset space role information of the user.
	Result *QueryUserRoleInfoInWorkspaceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Possible values:
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

func (s QueryUserRoleInfoInWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryUserRoleInfoInWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserRoleInfoInWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryUserRoleInfoInWorkspaceResponseBody) GetResult() *QueryUserRoleInfoInWorkspaceResponseBodyResult {
	return s.Result
}

func (s *QueryUserRoleInfoInWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryUserRoleInfoInWorkspaceResponseBody) SetRequestId(v string) *QueryUserRoleInfoInWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponseBody) SetResult(v *QueryUserRoleInfoInWorkspaceResponseBodyResult) *QueryUserRoleInfoInWorkspaceResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponseBody) SetSuccess(v bool) *QueryUserRoleInfoInWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryUserRoleInfoInWorkspaceResponseBodyResult struct {
	// Preset role code.
	//
	// example:
	//
	// role_workspace_admin
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	// Preset role ID. Possible values:
	//
	// - 25: Space Administrator
	//
	// - 26: Space Developer
	//
	// - 27: Space Analyst
	//
	// - 30: Space Viewer
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// Preset role name.
	//
	// example:
	//
	// test
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s QueryUserRoleInfoInWorkspaceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryUserRoleInfoInWorkspaceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserRoleInfoInWorkspaceResponseBodyResult) GetRoleCode() *string {
	return s.RoleCode
}

func (s *QueryUserRoleInfoInWorkspaceResponseBodyResult) GetRoleId() *int64 {
	return s.RoleId
}

func (s *QueryUserRoleInfoInWorkspaceResponseBodyResult) GetRoleName() *string {
	return s.RoleName
}

func (s *QueryUserRoleInfoInWorkspaceResponseBodyResult) SetRoleCode(v string) *QueryUserRoleInfoInWorkspaceResponseBodyResult {
	s.RoleCode = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponseBodyResult) SetRoleId(v int64) *QueryUserRoleInfoInWorkspaceResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponseBodyResult) SetRoleName(v string) *QueryUserRoleInfoInWorkspaceResponseBodyResult {
	s.RoleName = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
