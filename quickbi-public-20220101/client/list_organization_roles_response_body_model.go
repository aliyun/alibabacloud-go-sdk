// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListOrganizationRolesResponseBody
	GetRequestId() *string
	SetResult(v []*ListOrganizationRolesResponseBodyResult) *ListOrganizationRolesResponseBody
	GetResult() []*ListOrganizationRolesResponseBodyResult
	SetSuccess(v bool) *ListOrganizationRolesResponseBody
	GetSuccess() *bool
}

type ListOrganizationRolesResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 7AAB95D7-2E11-4FE2-94BC-858E4FC0C976
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the list of organization roles.
	Result []*ListOrganizationRolesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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

func (s ListOrganizationRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOrganizationRolesResponseBody) GetResult() []*ListOrganizationRolesResponseBodyResult {
	return s.Result
}

func (s *ListOrganizationRolesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOrganizationRolesResponseBody) SetRequestId(v string) *ListOrganizationRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationRolesResponseBody) SetResult(v []*ListOrganizationRolesResponseBodyResult) *ListOrganizationRolesResponseBody {
	s.Result = v
	return s
}

func (s *ListOrganizationRolesResponseBody) SetSuccess(v bool) *ListOrganizationRolesResponseBody {
	s.Success = &v
	return s
}

func (s *ListOrganizationRolesResponseBody) Validate() error {
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

type ListOrganizationRolesResponseBodyResult struct {
	// List of role permission configurations.
	AuthConfigList []*ListOrganizationRolesResponseBodyResultAuthConfigList `json:"AuthConfigList,omitempty" xml:"AuthConfigList,omitempty" type:"Repeated"`
	// Whether it is a predefined role. Possible values:
	//
	// - true: Yes
	//
	// - false: No
	//
	// example:
	//
	// true
	IsSystemRole *bool `json:"IsSystemRole,omitempty" xml:"IsSystemRole,omitempty"`
	// Role ID.
	//
	// example:
	//
	// 111111111
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// Role name.
	//
	// example:
	//
	// Organization Admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListOrganizationRolesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationRolesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOrganizationRolesResponseBodyResult) GetAuthConfigList() []*ListOrganizationRolesResponseBodyResultAuthConfigList {
	return s.AuthConfigList
}

func (s *ListOrganizationRolesResponseBodyResult) GetIsSystemRole() *bool {
	return s.IsSystemRole
}

func (s *ListOrganizationRolesResponseBodyResult) GetRoleId() *int64 {
	return s.RoleId
}

func (s *ListOrganizationRolesResponseBodyResult) GetRoleName() *string {
	return s.RoleName
}

func (s *ListOrganizationRolesResponseBodyResult) SetAuthConfigList(v []*ListOrganizationRolesResponseBodyResultAuthConfigList) *ListOrganizationRolesResponseBodyResult {
	s.AuthConfigList = v
	return s
}

func (s *ListOrganizationRolesResponseBodyResult) SetIsSystemRole(v bool) *ListOrganizationRolesResponseBodyResult {
	s.IsSystemRole = &v
	return s
}

func (s *ListOrganizationRolesResponseBodyResult) SetRoleId(v int64) *ListOrganizationRolesResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *ListOrganizationRolesResponseBodyResult) SetRoleName(v string) *ListOrganizationRolesResponseBodyResult {
	s.RoleName = &v
	return s
}

func (s *ListOrganizationRolesResponseBodyResult) Validate() error {
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

type ListOrganizationRolesResponseBodyResultAuthConfigList struct {
	// Permission type:
	//
	// - quick_monitor: Metric monitoring
	//
	// - subscription: Subscription management
	//
	// - offline_download: Self-service data retrieval
	//
	// - resource_package: Resource package management
	//
	// - organization_ask: Organization identification code (AK/SK)
	//
	// - developer_openapi: OpenAPI
	//
	// - data_service: Data service
	//
	// - admin_authorize3rd: Embedded analysis
	//
	// - component_manage: Custom component
	//
	// - template_open: Custom template
	//
	// - custom_driver: Custom driver (supported only in standalone deployment)
	//
	// - open_platform_custom_plugin: Custom plugin (supported only in standalone deployment)
	//
	// - enterprise_safety: Enterprise security
	//
	// example:
	//
	// enterprise_safety
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
}

func (s ListOrganizationRolesResponseBodyResultAuthConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationRolesResponseBodyResultAuthConfigList) GoString() string {
	return s.String()
}

func (s *ListOrganizationRolesResponseBodyResultAuthConfigList) GetAuthKey() *string {
	return s.AuthKey
}

func (s *ListOrganizationRolesResponseBodyResultAuthConfigList) SetAuthKey(v string) *ListOrganizationRolesResponseBodyResultAuthConfigList {
	s.AuthKey = &v
	return s
}

func (s *ListOrganizationRolesResponseBodyResultAuthConfigList) Validate() error {
	return dara.Validate(s)
}
