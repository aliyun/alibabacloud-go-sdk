// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrganizationRoleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryOrganizationRoleConfigResponseBody
	GetRequestId() *string
	SetResult(v *QueryOrganizationRoleConfigResponseBodyResult) *QueryOrganizationRoleConfigResponseBody
	GetResult() *QueryOrganizationRoleConfigResponseBodyResult
	SetSuccess(v bool) *QueryOrganizationRoleConfigResponseBody
	GetSuccess() *bool
}

type QueryOrganizationRoleConfigResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// BCE45E6D-9304-4F94-86BB-5A772B1615FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the organization role configuration.
	Result *QueryOrganizationRoleConfigResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful
	//
	// - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryOrganizationRoleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOrganizationRoleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrganizationRoleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryOrganizationRoleConfigResponseBody) GetResult() *QueryOrganizationRoleConfigResponseBodyResult {
	return s.Result
}

func (s *QueryOrganizationRoleConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryOrganizationRoleConfigResponseBody) SetRequestId(v string) *QueryOrganizationRoleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBody) SetResult(v *QueryOrganizationRoleConfigResponseBodyResult) *QueryOrganizationRoleConfigResponseBody {
	s.Result = v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBody) SetSuccess(v bool) *QueryOrganizationRoleConfigResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryOrganizationRoleConfigResponseBodyResult struct {
	// List of role permission configurations.
	AuthConfigList []*QueryOrganizationRoleConfigResponseBodyResultAuthConfigList `json:"AuthConfigList,omitempty" xml:"AuthConfigList,omitempty" type:"Repeated"`
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
	// Organization role ID, including predefined roles and custom roles:
	//
	// - Organization Administrator (predefined role): 111111111
	//
	// - Permission Administrator (predefined role): 111111112
	//
	// - Regular User (predefined role): 111111113
	//
	// - Custom Role: The corresponding role ID of the custom role
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

func (s QueryOrganizationRoleConfigResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryOrganizationRoleConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) GetAuthConfigList() []*QueryOrganizationRoleConfigResponseBodyResultAuthConfigList {
	return s.AuthConfigList
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) GetIsSystemRole() *bool {
	return s.IsSystemRole
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) GetRoleId() *int64 {
	return s.RoleId
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) GetRoleName() *string {
	return s.RoleName
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) SetAuthConfigList(v []*QueryOrganizationRoleConfigResponseBodyResultAuthConfigList) *QueryOrganizationRoleConfigResponseBodyResult {
	s.AuthConfigList = v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) SetIsSystemRole(v bool) *QueryOrganizationRoleConfigResponseBodyResult {
	s.IsSystemRole = &v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) SetRoleId(v int64) *QueryOrganizationRoleConfigResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) SetRoleName(v string) *QueryOrganizationRoleConfigResponseBodyResult {
	s.RoleName = &v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type QueryOrganizationRoleConfigResponseBodyResultAuthConfigList struct {
	// Permission type:
	//
	// - quick_monitor: Metric Monitoring
	//
	// - subscription: Subscription Management
	//
	// - offline_download: Self-service Data Retrieval
	//
	// - resource_package: Resource Package Management
	//
	// - organization_ask: Organization Access Key/Secret (AK/SK)
	//
	// - developer_openapi: OpenAPI
	//
	// - data_service: Data Service
	//
	// - admin_authorize3rd: Embedded Analysis
	//
	// - component_manage: Custom Component
	//
	// - template_open: Custom Template
	//
	// - custom_driver: Custom Driver (supported only in standalone deployment)
	//
	// - open_platform_custom_plugin: Custom Plugin (supported only in standalone deployment)
	//
	// - enterprise_safety: Enterprise Security
	//
	// example:
	//
	// quick_monitor
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
}

func (s QueryOrganizationRoleConfigResponseBodyResultAuthConfigList) String() string {
	return dara.Prettify(s)
}

func (s QueryOrganizationRoleConfigResponseBodyResultAuthConfigList) GoString() string {
	return s.String()
}

func (s *QueryOrganizationRoleConfigResponseBodyResultAuthConfigList) GetAuthKey() *string {
	return s.AuthKey
}

func (s *QueryOrganizationRoleConfigResponseBodyResultAuthConfigList) SetAuthKey(v string) *QueryOrganizationRoleConfigResponseBodyResultAuthConfigList {
	s.AuthKey = &v
	return s
}

func (s *QueryOrganizationRoleConfigResponseBodyResultAuthConfigList) Validate() error {
	return dara.Validate(s)
}
