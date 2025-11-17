// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWorkspaceRolesResponseBody
	GetRequestId() *string
	SetResult(v []*ListWorkspaceRolesResponseBodyResult) *ListWorkspaceRolesResponseBody
	GetResult() []*ListWorkspaceRolesResponseBodyResult
	SetSuccess(v bool) *ListWorkspaceRolesResponseBody
	GetSuccess() *bool
}

type ListWorkspaceRolesResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of workspace roles.
	Result []*ListWorkspaceRolesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Possible values:
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

func (s ListWorkspaceRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspaceRolesResponseBody) GetResult() []*ListWorkspaceRolesResponseBodyResult {
	return s.Result
}

func (s *ListWorkspaceRolesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkspaceRolesResponseBody) SetRequestId(v string) *ListWorkspaceRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspaceRolesResponseBody) SetResult(v []*ListWorkspaceRolesResponseBodyResult) *ListWorkspaceRolesResponseBody {
	s.Result = v
	return s
}

func (s *ListWorkspaceRolesResponseBody) SetSuccess(v bool) *ListWorkspaceRolesResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkspaceRolesResponseBody) Validate() error {
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

type ListWorkspaceRolesResponseBodyResult struct {
	// List of role authorization configurations.
	AuthConfigList []*ListWorkspaceRolesResponseBodyResultAuthConfigList `json:"AuthConfigList,omitempty" xml:"AuthConfigList,omitempty" type:"Repeated"`
	// Whether it is a predefined role. Value range:
	//
	// - true: Yes
	//
	// - false: No
	//
	// example:
	//
	// true
	IsSystemRole *bool `json:"IsSystemRole,omitempty" xml:"IsSystemRole,omitempty"`
	// Workspace role ID, including predefined and custom roles:
	//
	// - 25: Workspace Administrator (predefined role)
	//
	// - 26: Developer (predefined role)
	//
	// - 27: Analyst (predefined role)
	//
	// - 30: Viewer (predefined role)
	//
	// - Custom role: The corresponding role ID for a custom role
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// Role name.
	//
	// example:
	//
	// Space administrator
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListWorkspaceRolesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRolesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponseBodyResult) GetAuthConfigList() []*ListWorkspaceRolesResponseBodyResultAuthConfigList {
	return s.AuthConfigList
}

func (s *ListWorkspaceRolesResponseBodyResult) GetIsSystemRole() *bool {
	return s.IsSystemRole
}

func (s *ListWorkspaceRolesResponseBodyResult) GetRoleId() *int64 {
	return s.RoleId
}

func (s *ListWorkspaceRolesResponseBodyResult) GetRoleName() *string {
	return s.RoleName
}

func (s *ListWorkspaceRolesResponseBodyResult) SetAuthConfigList(v []*ListWorkspaceRolesResponseBodyResultAuthConfigList) *ListWorkspaceRolesResponseBodyResult {
	s.AuthConfigList = v
	return s
}

func (s *ListWorkspaceRolesResponseBodyResult) SetIsSystemRole(v bool) *ListWorkspaceRolesResponseBodyResult {
	s.IsSystemRole = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyResult) SetRoleId(v int64) *ListWorkspaceRolesResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyResult) SetRoleName(v string) *ListWorkspaceRolesResponseBodyResult {
	s.RoleName = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyResult) Validate() error {
	if s.AuthConfigList != nil {
		for _, item := range s.AuthConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspaceRolesResponseBodyResultAuthConfigList struct {
	// Authorization scope.
	ActionAuthKeys []*string `json:"ActionAuthKeys,omitempty" xml:"ActionAuthKeys,omitempty" type:"Repeated"`
	// Authorization type:
	//
	// - portal_create: Data Portal
	//
	// - dashboard_create: Dashboard
	//
	// - report_create: Spreadsheet
	//
	// - screen_create: Data Screen
	//
	// - analysis: Ad-hoc Analysis
	//
	// - offline_download: Self-service Data Retrieval
	//
	// - data_form: Data Entry
	//
	// - quick_etl: Data Preparation
	//
	// - cube: Dataset
	//
	// - datasource: Data Source
	//
	// example:
	//
	// portal_create
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
}

func (s ListWorkspaceRolesResponseBodyResultAuthConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRolesResponseBodyResultAuthConfigList) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponseBodyResultAuthConfigList) GetActionAuthKeys() []*string {
	return s.ActionAuthKeys
}

func (s *ListWorkspaceRolesResponseBodyResultAuthConfigList) GetAuthKey() *string {
	return s.AuthKey
}

func (s *ListWorkspaceRolesResponseBodyResultAuthConfigList) SetActionAuthKeys(v []*string) *ListWorkspaceRolesResponseBodyResultAuthConfigList {
	s.ActionAuthKeys = v
	return s
}

func (s *ListWorkspaceRolesResponseBodyResultAuthConfigList) SetAuthKey(v string) *ListWorkspaceRolesResponseBodyResultAuthConfigList {
	s.AuthKey = &v
	return s
}

func (s *ListWorkspaceRolesResponseBodyResultAuthConfigList) Validate() error {
	return dara.Validate(s)
}
