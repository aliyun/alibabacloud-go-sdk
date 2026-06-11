// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceUserRolesByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWorkspaceUserRolesByUserIdResponseBody
	GetRequestId() *string
	SetResult(v []*ListWorkspaceUserRolesByUserIdResponseBodyResult) *ListWorkspaceUserRolesByUserIdResponseBody
	GetResult() []*ListWorkspaceUserRolesByUserIdResponseBodyResult
	SetSuccess(v bool) *ListWorkspaceUserRolesByUserIdResponseBody
	GetSuccess() *bool
}

type ListWorkspaceUserRolesByUserIdResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DC4E***************F67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of all workspace role information.
	//
	// This parameter is required.
	Result []*ListWorkspaceUserRolesByUserIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWorkspaceUserRolesByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUserRolesByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUserRolesByUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspaceUserRolesByUserIdResponseBody) GetResult() []*ListWorkspaceUserRolesByUserIdResponseBodyResult {
	return s.Result
}

func (s *ListWorkspaceUserRolesByUserIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkspaceUserRolesByUserIdResponseBody) SetRequestId(v string) *ListWorkspaceUserRolesByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdResponseBody) SetResult(v []*ListWorkspaceUserRolesByUserIdResponseBodyResult) *ListWorkspaceUserRolesByUserIdResponseBody {
	s.Result = v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdResponseBody) SetSuccess(v bool) *ListWorkspaceUserRolesByUserIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspaceUserRolesByUserIdResponseBodyResult struct {
	// The workspace role information.
	RoleModel *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel `json:"RoleModel,omitempty" xml:"RoleModel,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// example:
	//
	// 9337d121-a78f-4c1b-a8bc-f81de117****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// test
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListWorkspaceUserRolesByUserIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUserRolesByUserIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResult) GetRoleModel() *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel {
	return s.RoleModel
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResult) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResult) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResult) SetRoleModel(v *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel) *ListWorkspaceUserRolesByUserIdResponseBodyResult {
	s.RoleModel = v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResult) SetWorkspaceId(v string) *ListWorkspaceUserRolesByUserIdResponseBodyResult {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResult) SetWorkspaceName(v string) *ListWorkspaceUserRolesByUserIdResponseBodyResult {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResult) Validate() error {
	if s.RoleModel != nil {
		if err := s.RoleModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel struct {
	// The code that corresponds to the preset workspace role. This parameter is not returned for custom roles.
	//
	// example:
	//
	// 34637***35
	RoleCode *string `json:"RoleCode,omitempty" xml:"RoleCode,omitempty"`
	// The workspace role ID. This includes, but is not limited to, the following preset workspace roles:
	//
	// - 25: Workspace Admin (preset role)
	//
	// - 26: Workspace Developer (preset role)
	//
	// - 27: Workspace Analyst (preset role)
	//
	// - 30: Workspace Viewer (preset role)
	//
	// example:
	//
	// 111111111
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the workspace role.
	//
	// example:
	//
	// arms-admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel) GetRoleCode() *string {
	return s.RoleCode
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel) GetRoleId() *string {
	return s.RoleId
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel) GetRoleName() *string {
	return s.RoleName
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel) SetRoleCode(v string) *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel {
	s.RoleCode = &v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel) SetRoleId(v string) *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel {
	s.RoleId = &v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel) SetRoleName(v string) *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel {
	s.RoleName = &v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdResponseBodyResultRoleModel) Validate() error {
	return dara.Validate(s)
}
