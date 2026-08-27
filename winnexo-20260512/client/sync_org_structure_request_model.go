// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncOrgStructureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpId(v string) *SyncOrgStructureRequest
	GetCorpId() *string
	SetDepartments(v []*SyncOrgStructureRequestDepartments) *SyncOrgStructureRequest
	GetDepartments() []*SyncOrgStructureRequestDepartments
	SetMembers(v []*SyncOrgStructureRequestMembers) *SyncOrgStructureRequest
	GetMembers() []*SyncOrgStructureRequestMembers
	SetPlatformType(v string) *SyncOrgStructureRequest
	GetPlatformType() *string
	SetSsoSettingsId(v string) *SyncOrgStructureRequest
	GetSsoSettingsId() *string
	SetSyncMembers(v bool) *SyncOrgStructureRequest
	GetSyncMembers() *bool
	SetTenantId(v string) *SyncOrgStructureRequest
	GetTenantId() *string
}

type SyncOrgStructureRequest struct {
	// The enterprise identifier. This value must match the corpId returned by listAvailableConfigs.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleCorpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// The department list. At least one root department must be included.
	//
	// This parameter is required.
	Departments []*SyncOrgStructureRequestDepartments `json:"departments,omitempty" xml:"departments,omitempty" type:"Repeated"`
	// The member list. This parameter is required when syncMembers is set to true.
	Members []*SyncOrgStructureRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// The platform type. Valid values: saml, oauth2, or custom.
	//
	// This parameter is required.
	//
	// example:
	//
	// saml
	PlatformType *string `json:"platformType,omitempty" xml:"platformType,omitempty"`
	// The SSO configuration ID. For SAML/OAuth2, this parameter is optional. If not specified, the value is automatically derived based on corpId. If multiple IdPs use the same corpId, you must explicitly specify this parameter. Otherwise, an AMBIGUOUS error is returned. This parameter is not required for custom.
	//
	// example:
	//
	// exampleSsoSettingsId
	SsoSettingsId *string `json:"ssoSettingsId,omitempty" xml:"ssoSettingsId,omitempty"`
	// Specifies whether to synchronize member relationships. In custom mode, this parameter is forced to false.
	//
	// example:
	//
	// false
	SyncMembers *bool `json:"syncMembers,omitempty" xml:"syncMembers,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SyncOrgStructureRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncOrgStructureRequest) GoString() string {
	return s.String()
}

func (s *SyncOrgStructureRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *SyncOrgStructureRequest) GetDepartments() []*SyncOrgStructureRequestDepartments {
	return s.Departments
}

func (s *SyncOrgStructureRequest) GetMembers() []*SyncOrgStructureRequestMembers {
	return s.Members
}

func (s *SyncOrgStructureRequest) GetPlatformType() *string {
	return s.PlatformType
}

func (s *SyncOrgStructureRequest) GetSsoSettingsId() *string {
	return s.SsoSettingsId
}

func (s *SyncOrgStructureRequest) GetSyncMembers() *bool {
	return s.SyncMembers
}

func (s *SyncOrgStructureRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SyncOrgStructureRequest) SetCorpId(v string) *SyncOrgStructureRequest {
	s.CorpId = &v
	return s
}

func (s *SyncOrgStructureRequest) SetDepartments(v []*SyncOrgStructureRequestDepartments) *SyncOrgStructureRequest {
	s.Departments = v
	return s
}

func (s *SyncOrgStructureRequest) SetMembers(v []*SyncOrgStructureRequestMembers) *SyncOrgStructureRequest {
	s.Members = v
	return s
}

func (s *SyncOrgStructureRequest) SetPlatformType(v string) *SyncOrgStructureRequest {
	s.PlatformType = &v
	return s
}

func (s *SyncOrgStructureRequest) SetSsoSettingsId(v string) *SyncOrgStructureRequest {
	s.SsoSettingsId = &v
	return s
}

func (s *SyncOrgStructureRequest) SetSyncMembers(v bool) *SyncOrgStructureRequest {
	s.SyncMembers = &v
	return s
}

func (s *SyncOrgStructureRequest) SetTenantId(v string) *SyncOrgStructureRequest {
	s.TenantId = &v
	return s
}

func (s *SyncOrgStructureRequest) Validate() error {
	if s.Departments != nil {
		for _, item := range s.Departments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SyncOrgStructureRequestDepartments struct {
	// The department ID. This is an external identifier. The client is responsible for ensuring uniqueness.
	//
	// example:
	//
	// exampleDeptId
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// The department name.
	//
	// example:
	//
	// string_value
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// The sort order. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Order *int64 `json:"order,omitempty" xml:"order,omitempty"`
	// The parent department ID. A value of null indicates a top-level department or root department.
	//
	// example:
	//
	// exampleParentDeptId
	ParentDeptId *string `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
}

func (s SyncOrgStructureRequestDepartments) String() string {
	return dara.Prettify(s)
}

func (s SyncOrgStructureRequestDepartments) GoString() string {
	return s.String()
}

func (s *SyncOrgStructureRequestDepartments) GetDeptId() *string {
	return s.DeptId
}

func (s *SyncOrgStructureRequestDepartments) GetDeptName() *string {
	return s.DeptName
}

func (s *SyncOrgStructureRequestDepartments) GetOrder() *int64 {
	return s.Order
}

func (s *SyncOrgStructureRequestDepartments) GetParentDeptId() *string {
	return s.ParentDeptId
}

func (s *SyncOrgStructureRequestDepartments) SetDeptId(v string) *SyncOrgStructureRequestDepartments {
	s.DeptId = &v
	return s
}

func (s *SyncOrgStructureRequestDepartments) SetDeptName(v string) *SyncOrgStructureRequestDepartments {
	s.DeptName = &v
	return s
}

func (s *SyncOrgStructureRequestDepartments) SetOrder(v int64) *SyncOrgStructureRequestDepartments {
	s.Order = &v
	return s
}

func (s *SyncOrgStructureRequestDepartments) SetParentDeptId(v string) *SyncOrgStructureRequestDepartments {
	s.ParentDeptId = &v
	return s
}

func (s *SyncOrgStructureRequestDepartments) Validate() error {
	return dara.Validate(s)
}

type SyncOrgStructureRequestMembers struct {
	// The user identifier. In the SAML scenario, this is an email address or UPN, which must match rbj_user_account.account_id.
	//
	// example:
	//
	// exampleAccountId
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The department ID to which the member belongs. This value must correspond to a deptId in the departments list.
	//
	// example:
	//
	// exampleDeptId
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// The username for display purposes. This parameter is optional.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s SyncOrgStructureRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s SyncOrgStructureRequestMembers) GoString() string {
	return s.String()
}

func (s *SyncOrgStructureRequestMembers) GetAccountId() *string {
	return s.AccountId
}

func (s *SyncOrgStructureRequestMembers) GetDeptId() *string {
	return s.DeptId
}

func (s *SyncOrgStructureRequestMembers) GetName() *string {
	return s.Name
}

func (s *SyncOrgStructureRequestMembers) SetAccountId(v string) *SyncOrgStructureRequestMembers {
	s.AccountId = &v
	return s
}

func (s *SyncOrgStructureRequestMembers) SetDeptId(v string) *SyncOrgStructureRequestMembers {
	s.DeptId = &v
	return s
}

func (s *SyncOrgStructureRequestMembers) SetName(v string) *SyncOrgStructureRequestMembers {
	s.Name = &v
	return s
}

func (s *SyncOrgStructureRequestMembers) Validate() error {
	return dara.Validate(s)
}
