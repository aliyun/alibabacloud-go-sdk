// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorkspaceRoleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryWorkspaceRoleConfigResponseBody
	GetRequestId() *string
	SetResult(v *QueryWorkspaceRoleConfigResponseBodyResult) *QueryWorkspaceRoleConfigResponseBody
	GetResult() *QueryWorkspaceRoleConfigResponseBodyResult
	SetSuccess(v bool) *QueryWorkspaceRoleConfigResponseBody
	GetSuccess() *bool
}

type QueryWorkspaceRoleConfigResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D8749D65-E80A-433C-AF1B-CE9C180FF3B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the query result of the interface.
	Result *QueryWorkspaceRoleConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s QueryWorkspaceRoleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkspaceRoleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceRoleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryWorkspaceRoleConfigResponseBody) GetResult() *QueryWorkspaceRoleConfigResponseBodyResult {
	return s.Result
}

func (s *QueryWorkspaceRoleConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryWorkspaceRoleConfigResponseBody) SetRequestId(v string) *QueryWorkspaceRoleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBody) SetResult(v *QueryWorkspaceRoleConfigResponseBodyResult) *QueryWorkspaceRoleConfigResponseBody {
	s.Result = v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBody) SetSuccess(v bool) *QueryWorkspaceRoleConfigResponseBody {
	s.Success = &v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryWorkspaceRoleConfigResponseBodyResult struct {
	// List of role permission configurations.
	AuthConfigList []*QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList `json:"AuthConfigList,omitempty" xml:"AuthConfigList,omitempty" type:"Repeated"`
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
	// Workspace role ID, including predefined roles and custom roles:
	//
	// - 25: Workspace Administrator (predefined role)
	//
	// - 26: Developer (predefined role)
	//
	// - 27: Analyst (predefined role)
	//
	// - 30: Viewer (predefined role)
	//
	// - Custom role: The corresponding role ID for the custom role
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// Role name.
	//
	// example:
	//
	// pace administrator
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s QueryWorkspaceRoleConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkspaceRoleConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) GetAuthConfigList() []*QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList {
	return s.AuthConfigList
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) GetIsSystemRole() *bool {
	return s.IsSystemRole
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) GetRoleId() *int64 {
	return s.RoleId
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) GetRoleName() *string {
	return s.RoleName
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) SetAuthConfigList(v []*QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) *QueryWorkspaceRoleConfigResponseBodyResult {
	s.AuthConfigList = v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) SetIsSystemRole(v bool) *QueryWorkspaceRoleConfigResponseBodyResult {
	s.IsSystemRole = &v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) SetRoleId(v int64) *QueryWorkspaceRoleConfigResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) SetRoleName(v string) *QueryWorkspaceRoleConfigResponseBodyResult {
	s.RoleName = &v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBodyResult) Validate() error {
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

type QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList struct {
	// Permission scope.
	ActionAuthKeys []*string `json:"ActionAuthKeys,omitempty" xml:"ActionAuthKeys,omitempty" type:"Repeated"`
	// Permission type:
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

func (s QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) GetActionAuthKeys() []*string {
	return s.ActionAuthKeys
}

func (s *QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) GetAuthKey() *string {
	return s.AuthKey
}

func (s *QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) SetActionAuthKeys(v []*string) *QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList {
	s.ActionAuthKeys = v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) SetAuthKey(v string) *QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList {
	s.AuthKey = &v
	return s
}

func (s *QueryWorkspaceRoleConfigResponseBodyResultAuthConfigList) Validate() error {
	return dara.Validate(s)
}
