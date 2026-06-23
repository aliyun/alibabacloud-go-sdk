// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterKubeconfigStatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v *ListClusterKubeconfigStatesResponseBodyPage) *ListClusterKubeconfigStatesResponseBody
	GetPage() *ListClusterKubeconfigStatesResponseBodyPage
	SetStates(v []*ListClusterKubeconfigStatesResponseBodyStates) *ListClusterKubeconfigStatesResponseBody
	GetStates() []*ListClusterKubeconfigStatesResponseBodyStates
}

type ListClusterKubeconfigStatesResponseBody struct {
	// The pagination information.
	Page *ListClusterKubeconfigStatesResponseBodyPage `json:"page,omitempty" xml:"page,omitempty" type:"Struct"`
	// The list of KubeConfig states associated with the cluster.
	States []*ListClusterKubeconfigStatesResponseBodyStates `json:"states,omitempty" xml:"states,omitempty" type:"Repeated"`
}

func (s ListClusterKubeconfigStatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterKubeconfigStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterKubeconfigStatesResponseBody) GetPage() *ListClusterKubeconfigStatesResponseBodyPage {
	return s.Page
}

func (s *ListClusterKubeconfigStatesResponseBody) GetStates() []*ListClusterKubeconfigStatesResponseBodyStates {
	return s.States
}

func (s *ListClusterKubeconfigStatesResponseBody) SetPage(v *ListClusterKubeconfigStatesResponseBodyPage) *ListClusterKubeconfigStatesResponseBody {
	s.Page = v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBody) SetStates(v []*ListClusterKubeconfigStatesResponseBodyStates) *ListClusterKubeconfigStatesResponseBody {
	s.States = v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	if s.States != nil {
		for _, item := range s.States {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterKubeconfigStatesResponseBodyPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The maximum number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of results.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s ListClusterKubeconfigStatesResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s ListClusterKubeconfigStatesResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) SetPageNumber(v int32) *ListClusterKubeconfigStatesResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) SetPageSize(v int32) *ListClusterKubeconfigStatesResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) SetTotalCount(v int32) *ListClusterKubeconfigStatesResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyPage) Validate() error {
	return dara.Validate(s)
}

type ListClusterKubeconfigStatesResponseBodyStates struct {
	// The display name of the Resource Access Management (RAM) user or the role name.
	//
	// example:
	//
	// tom
	AccountDisplayName *string `json:"account_display_name,omitempty" xml:"account_display_name,omitempty"`
	// The Alibaba Cloud account, Resource Access Management (RAM) user, or role ID.
	//
	// example:
	//
	// 22855*****************
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// The logon name of the Resource Access Management (RAM) user or the role name.
	//
	// example:
	//
	// tom
	AccountName *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	// The account status. Valid values:
	//
	// - Active: The account is active.
	//
	// - InActive: The account is frozen.
	//
	// - Deleted: The account is deleted.
	//
	// example:
	//
	// Active
	AccountState *string `json:"account_state,omitempty" xml:"account_state,omitempty"`
	// The account type. Valid values:
	//
	// - RootAccount: Alibaba Cloud account.
	//
	// - RamUser: Resource Access Management (RAM) user.
	//
	// - RamRole: RAM role.
	//
	// example:
	//
	// RamUser
	AccountType *string `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// The expiration time of the KubeConfig client certificate.
	//
	// example:
	//
	// 2027-07-15T01:32:20Z
	CertExpireTime *string `json:"cert_expire_time,omitempty" xml:"cert_expire_time,omitempty"`
	// The status of the KubeConfig client certificate. Valid values:
	//
	// - Unexpired: The certificate has not expired.
	//
	// - Expired: The certificate has expired.
	//
	// - Unknown: The certificate status is unknown (abnormal state).
	//
	// example:
	//
	// Expired
	CertState *string `json:"cert_state,omitempty" xml:"cert_state,omitempty"`
	// The name of the cloud service.
	//
	// example:
	//
	// cs/ecs/sls
	CloudServiceName *string `json:"cloud_service_name,omitempty" xml:"cloud_service_name,omitempty"`
	// The list of cluster roles associated with the cloud service role.
	CloudServiceRoles []*ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles `json:"cloud_service_roles,omitempty" xml:"cloud_service_roles,omitempty" type:"Repeated"`
	// Indicates whether the KubeConfig client certificate can be revoked.
	//
	// example:
	//
	// true
	Revokable *bool `json:"revokable,omitempty" xml:"revokable,omitempty"`
}

func (s ListClusterKubeconfigStatesResponseBodyStates) String() string {
	return dara.Prettify(s)
}

func (s ListClusterKubeconfigStatesResponseBodyStates) GoString() string {
	return s.String()
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetAccountDisplayName() *string {
	return s.AccountDisplayName
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetAccountId() *string {
	return s.AccountId
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetAccountName() *string {
	return s.AccountName
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetAccountState() *string {
	return s.AccountState
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetAccountType() *string {
	return s.AccountType
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetCertState() *string {
	return s.CertState
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetCloudServiceName() *string {
	return s.CloudServiceName
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetCloudServiceRoles() []*ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles {
	return s.CloudServiceRoles
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) GetRevokable() *bool {
	return s.Revokable
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetAccountDisplayName(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.AccountDisplayName = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetAccountId(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.AccountId = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetAccountName(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.AccountName = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetAccountState(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.AccountState = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetAccountType(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.AccountType = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetCertExpireTime(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.CertExpireTime = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetCertState(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.CertState = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetCloudServiceName(v string) *ListClusterKubeconfigStatesResponseBodyStates {
	s.CloudServiceName = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetCloudServiceRoles(v []*ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles) *ListClusterKubeconfigStatesResponseBodyStates {
	s.CloudServiceRoles = v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) SetRevokable(v bool) *ListClusterKubeconfigStatesResponseBodyStates {
	s.Revokable = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStates) Validate() error {
	if s.CloudServiceRoles != nil {
		for _, item := range s.CloudServiceRoles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles struct {
	// Indicates whether the content matches the default cluster role template. Valid values:
	//
	// - true: The content matches the default template.
	//
	// - false: The content does not match the default template.
	IsDefaultTemplate *bool `json:"is_default_template,omitempty" xml:"is_default_template,omitempty"`
	// The name of the cluster role associated with the cloud service role.
	//
	// example:
	//
	// cluster-admin
	RoleName *string `json:"role_name,omitempty" xml:"role_name,omitempty"`
	// The namespace of the cluster role associated with the cloud service role.
	//
	// example:
	//
	// kube-system
	RoleNamespace *string `json:"role_namespace,omitempty" xml:"role_namespace,omitempty"`
	// The type of the cluster role associated with the cloud service role.
	//
	// example:
	//
	// ClusterRole
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles) String() string {
	return dara.Prettify(s)
}

func (s ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles) GoString() string {
	return s.String()
}

func (s *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles) GetIsDefaultTemplate() *bool {
	return s.IsDefaultTemplate
}

func (s *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles) GetRoleName() *string {
	return s.RoleName
}

func (s *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles) GetRoleNamespace() *string {
	return s.RoleNamespace
}

func (s *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles) GetType() *string {
	return s.Type
}

func (s *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles) SetIsDefaultTemplate(v bool) *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles {
	s.IsDefaultTemplate = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles) SetRoleName(v string) *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles {
	s.RoleName = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles) SetRoleNamespace(v string) *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles {
	s.RoleNamespace = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles) SetType(v string) *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles {
	s.Type = &v
	return s
}

func (s *ListClusterKubeconfigStatesResponseBodyStatesCloudServiceRoles) Validate() error {
	return dara.Validate(s)
}
