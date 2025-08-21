// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAccountSummaryResponseBody
	GetRequestId() *string
	SetSummaryMap(v *GetAccountSummaryResponseBodySummaryMap) *GetAccountSummaryResponseBody
	GetSummaryMap() *GetAccountSummaryResponseBodySummaryMap
}

type GetAccountSummaryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 81313F5E-3C85-478F-BCC9-E1B70E4556DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The overview information about the Alibaba Cloud account.
	SummaryMap *GetAccountSummaryResponseBodySummaryMap `json:"SummaryMap,omitempty" xml:"SummaryMap,omitempty" type:"Struct"`
}

func (s GetAccountSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountSummaryResponseBody) GetSummaryMap() *GetAccountSummaryResponseBodySummaryMap {
	return s.SummaryMap
}

func (s *GetAccountSummaryResponseBody) SetRequestId(v string) *GetAccountSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountSummaryResponseBody) SetSummaryMap(v *GetAccountSummaryResponseBodySummaryMap) *GetAccountSummaryResponseBody {
	s.SummaryMap = v
	return s
}

func (s *GetAccountSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAccountSummaryResponseBodySummaryMap struct {
	// The maximum number of AccessKey pairs that a Resource Access Management (RAM) user can have.
	//
	// example:
	//
	// 2
	AccessKeysPerUserQuota           *int32 `json:"AccessKeysPerUserQuota,omitempty" xml:"AccessKeysPerUserQuota,omitempty"`
	AccountAccessKeysPerAccountQuota *int32 `json:"AccountAccessKeysPerAccountQuota,omitempty" xml:"AccountAccessKeysPerAccountQuota,omitempty"`
	// The maximum number of custom policies that can be added to a RAM user group.
	//
	// example:
	//
	// 5
	AttachedPoliciesPerGroupQuota *int32 `json:"AttachedPoliciesPerGroupQuota,omitempty" xml:"AttachedPoliciesPerGroupQuota,omitempty"`
	// The maximum number of custom policies that can be added to a RAM role.
	//
	// example:
	//
	// 5
	AttachedPoliciesPerRoleQuota *int32 `json:"AttachedPoliciesPerRoleQuota,omitempty" xml:"AttachedPoliciesPerRoleQuota,omitempty"`
	// The maximum number of custom policies that can be added to a RAM user.
	//
	// example:
	//
	// 10
	AttachedPoliciesPerUserQuota *int32 `json:"AttachedPoliciesPerUserQuota,omitempty" xml:"AttachedPoliciesPerUserQuota,omitempty"`
	// The maximum number of system policies that can be added to a RAM user group.
	//
	// example:
	//
	// 20
	AttachedSystemPoliciesPerGroupQuota *int32 `json:"AttachedSystemPoliciesPerGroupQuota,omitempty" xml:"AttachedSystemPoliciesPerGroupQuota,omitempty"`
	// The maximum number of system policies that can be added to a RAM role.
	//
	// example:
	//
	// 20
	AttachedSystemPoliciesPerRoleQuota *int32 `json:"AttachedSystemPoliciesPerRoleQuota,omitempty" xml:"AttachedSystemPoliciesPerRoleQuota,omitempty"`
	// The maximum number of system policies that can be added to a RAM user.
	//
	// example:
	//
	// 20
	AttachedSystemPoliciesPerUserQuota *int32 `json:"AttachedSystemPoliciesPerUserQuota,omitempty" xml:"AttachedSystemPoliciesPerUserQuota,omitempty"`
	// The maximum number of network access control policies that can be configured for an Alibaba Cloud account or AccessKey pair.
	//
	// example:
	//
	// 8
	ConditionsPerAKPolicyQuota *int32 `json:"ConditionsPerAKPolicyQuota,omitempty" xml:"ConditionsPerAKPolicyQuota,omitempty"`
	// The number of RAM user groups.
	//
	// example:
	//
	// 7
	Groups *int32 `json:"Groups,omitempty" xml:"Groups,omitempty"`
	// The maximum number of RAM user groups to which a RAM user can be added.
	//
	// example:
	//
	// 5
	GroupsPerUserQuota *int32 `json:"GroupsPerUserQuota,omitempty" xml:"GroupsPerUserQuota,omitempty"`
	// The maximum number of RAM user groups that can be created.
	//
	// example:
	//
	// 50
	GroupsQuota *int32 `json:"GroupsQuota,omitempty" xml:"GroupsQuota,omitempty"`
	// The maximum number of IP addresses that can be specified in an account-level AccessKey pair-based or AccessKey pair-level policy for network access control.
	//
	// example:
	//
	// 50
	IPItemsPerAKPolicyQuota *int32 `json:"IPItemsPerAKPolicyQuota,omitempty" xml:"IPItemsPerAKPolicyQuota,omitempty"`
	// The number of virtual multi-factor authentication (MFA) devices.
	//
	// example:
	//
	// 13
	MFADevices *int32 `json:"MFADevices,omitempty" xml:"MFADevices,omitempty"`
	// The number of virtual MFA devices in use.
	//
	// example:
	//
	// 2
	MFADevicesInUse *int32 `json:"MFADevicesInUse,omitempty" xml:"MFADevicesInUse,omitempty"`
	// The number of custom policies.
	//
	// example:
	//
	// 13
	Policies *int32 `json:"Policies,omitempty" xml:"Policies,omitempty"`
	// The maximum number of custom policies that can be created.
	//
	// example:
	//
	// 1500
	PoliciesQuota *int32 `json:"PoliciesQuota,omitempty" xml:"PoliciesQuota,omitempty"`
	// The maximum length of the policy content.
	//
	// example:
	//
	// 2048
	PolicySizeQuota *int32 `json:"PolicySizeQuota,omitempty" xml:"PolicySizeQuota,omitempty"`
	// The number of RAM roles.
	//
	// example:
	//
	// 19
	Roles *int32 `json:"Roles,omitempty" xml:"Roles,omitempty"`
	// The maximum number of RAM roles that can be created.
	//
	// example:
	//
	// 1000
	RolesQuota *int32 `json:"RolesQuota,omitempty" xml:"RolesQuota,omitempty"`
	// The number of RAM users.
	//
	// example:
	//
	// 9
	Users *int32 `json:"Users,omitempty" xml:"Users,omitempty"`
	// The maximum number of RAM users that can be created.
	//
	// example:
	//
	// 1000
	UsersQuota *int32 `json:"UsersQuota,omitempty" xml:"UsersQuota,omitempty"`
	// The maximum number of policy versions.
	//
	// example:
	//
	// 5
	VersionsPerPolicyQuota *int32 `json:"VersionsPerPolicyQuota,omitempty" xml:"VersionsPerPolicyQuota,omitempty"`
	// The maximum number of virtual MFA devices that can be created.
	//
	// example:
	//
	// 1000
	VirtualMFADevicesQuota *int32 `json:"VirtualMFADevicesQuota,omitempty" xml:"VirtualMFADevicesQuota,omitempty"`
}

func (s GetAccountSummaryResponseBodySummaryMap) String() string {
	return dara.Prettify(s)
}

func (s GetAccountSummaryResponseBodySummaryMap) GoString() string {
	return s.String()
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetAccessKeysPerUserQuota() *int32 {
	return s.AccessKeysPerUserQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetAccountAccessKeysPerAccountQuota() *int32 {
	return s.AccountAccessKeysPerAccountQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetAttachedPoliciesPerGroupQuota() *int32 {
	return s.AttachedPoliciesPerGroupQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetAttachedPoliciesPerRoleQuota() *int32 {
	return s.AttachedPoliciesPerRoleQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetAttachedPoliciesPerUserQuota() *int32 {
	return s.AttachedPoliciesPerUserQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetAttachedSystemPoliciesPerGroupQuota() *int32 {
	return s.AttachedSystemPoliciesPerGroupQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetAttachedSystemPoliciesPerRoleQuota() *int32 {
	return s.AttachedSystemPoliciesPerRoleQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetAttachedSystemPoliciesPerUserQuota() *int32 {
	return s.AttachedSystemPoliciesPerUserQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetConditionsPerAKPolicyQuota() *int32 {
	return s.ConditionsPerAKPolicyQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetGroups() *int32 {
	return s.Groups
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetGroupsPerUserQuota() *int32 {
	return s.GroupsPerUserQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetGroupsQuota() *int32 {
	return s.GroupsQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetIPItemsPerAKPolicyQuota() *int32 {
	return s.IPItemsPerAKPolicyQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetMFADevices() *int32 {
	return s.MFADevices
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetMFADevicesInUse() *int32 {
	return s.MFADevicesInUse
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetPolicies() *int32 {
	return s.Policies
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetPoliciesQuota() *int32 {
	return s.PoliciesQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetPolicySizeQuota() *int32 {
	return s.PolicySizeQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetRoles() *int32 {
	return s.Roles
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetRolesQuota() *int32 {
	return s.RolesQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetUsers() *int32 {
	return s.Users
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetUsersQuota() *int32 {
	return s.UsersQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetVersionsPerPolicyQuota() *int32 {
	return s.VersionsPerPolicyQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) GetVirtualMFADevicesQuota() *int32 {
	return s.VirtualMFADevicesQuota
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAccessKeysPerUserQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AccessKeysPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAccountAccessKeysPerAccountQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AccountAccessKeysPerAccountQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedPoliciesPerGroupQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedPoliciesPerGroupQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedPoliciesPerRoleQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedPoliciesPerRoleQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedPoliciesPerUserQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedPoliciesPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedSystemPoliciesPerGroupQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedSystemPoliciesPerGroupQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedSystemPoliciesPerRoleQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedSystemPoliciesPerRoleQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedSystemPoliciesPerUserQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedSystemPoliciesPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetConditionsPerAKPolicyQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.ConditionsPerAKPolicyQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetGroups(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.Groups = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetGroupsPerUserQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.GroupsPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetGroupsQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.GroupsQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetIPItemsPerAKPolicyQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.IPItemsPerAKPolicyQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetMFADevices(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.MFADevices = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetMFADevicesInUse(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.MFADevicesInUse = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetPolicies(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.Policies = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetPoliciesQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.PoliciesQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetPolicySizeQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.PolicySizeQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetRoles(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.Roles = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetRolesQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.RolesQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetUsers(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.Users = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetUsersQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.UsersQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetVersionsPerPolicyQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.VersionsPerPolicyQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetVirtualMFADevicesQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.VirtualMFADevicesQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) Validate() error {
	return dara.Validate(s)
}
