// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPrivateAccessPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolices(v []*ListUserPrivateAccessPoliciesResponseBodyPolices) *ListUserPrivateAccessPoliciesResponseBody
	GetPolices() []*ListUserPrivateAccessPoliciesResponseBodyPolices
	SetRequestId(v string) *ListUserPrivateAccessPoliciesResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListUserPrivateAccessPoliciesResponseBody
	GetTotalNum() *int32
}

type ListUserPrivateAccessPoliciesResponseBody struct {
	// The list of authorized access policies.
	Polices []*ListUserPrivateAccessPoliciesResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// Id of the request.
	//
	// example:
	//
	// 9D852F87-AFB5-51B8-AACD-F7D0EFB8277D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of authorized access policies.
	//
	// example:
	//
	// 20
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListUserPrivateAccessPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserPrivateAccessPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserPrivateAccessPoliciesResponseBody) GetPolices() []*ListUserPrivateAccessPoliciesResponseBodyPolices {
	return s.Polices
}

func (s *ListUserPrivateAccessPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserPrivateAccessPoliciesResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListUserPrivateAccessPoliciesResponseBody) SetPolices(v []*ListUserPrivateAccessPoliciesResponseBodyPolices) *ListUserPrivateAccessPoliciesResponseBody {
	s.Polices = v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBody) SetRequestId(v string) *ListUserPrivateAccessPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBody) SetTotalNum(v int32) *ListUserPrivateAccessPoliciesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBody) Validate() error {
	if s.Polices != nil {
		for _, item := range s.Polices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserPrivateAccessPoliciesResponseBodyPolices struct {
	// The collection of custom user group attributes. Multiple custom user group attributes have an OR relationship and take effect as a union.
	CustomUserAttributes []*ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	// The associated security baseline.
	//
	// example:
	//
	// device_attribute_name
	DeviceAttributeName *string `json:"DeviceAttributeName,omitempty" xml:"DeviceAttributeName,omitempty"`
	// The user group to which the policy belongs.
	//
	// example:
	//
	// user_group_name
	MatchedUserGroup *string `json:"MatchedUserGroup,omitempty" xml:"MatchedUserGroup,omitempty"`
	// The name of the private access policy.
	//
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The action of the private access policy. Valid values:
	//
	// - **Block**: blocks the access.
	//
	// - **Allow**: allows the access.
	//
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// The ID of the private access policy.
	//
	// example:
	//
	// pa-policy-1b0d0e8b4bcf****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The priority of the private access policy. The number 1 indicates the highest priority.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The list of trusted process group IDs.
	TrustedProcessGroupIds []*string `json:"TrustedProcessGroupIds,omitempty" xml:"TrustedProcessGroupIds,omitempty" type:"Repeated"`
	// The list of trusted software IDs.
	TrustedSoftwareIds []*string `json:"TrustedSoftwareIds,omitempty" xml:"TrustedSoftwareIds,omitempty" type:"Repeated"`
	// The user group type of the private access policy. Valid values:
	//
	// - **Normal**: a normal user group.
	//
	// - **Custom**: a custom user group.
	//
	// example:
	//
	// Custom
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
}

func (s ListUserPrivateAccessPoliciesResponseBodyPolices) String() string {
	return dara.Prettify(s)
}

func (s ListUserPrivateAccessPoliciesResponseBodyPolices) GoString() string {
	return s.String()
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) GetCustomUserAttributes() []*ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes {
	return s.CustomUserAttributes
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) GetDeviceAttributeName() *string {
	return s.DeviceAttributeName
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) GetMatchedUserGroup() *string {
	return s.MatchedUserGroup
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) GetName() *string {
	return s.Name
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) GetPolicyAction() *string {
	return s.PolicyAction
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) GetPriority() *int64 {
	return s.Priority
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) GetTrustedProcessGroupIds() []*string {
	return s.TrustedProcessGroupIds
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) GetTrustedSoftwareIds() []*string {
	return s.TrustedSoftwareIds
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) GetUserGroupMode() *string {
	return s.UserGroupMode
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetCustomUserAttributes(v []*ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.CustomUserAttributes = v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetDeviceAttributeName(v string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.DeviceAttributeName = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetMatchedUserGroup(v string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.MatchedUserGroup = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetName(v string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.Name = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetPolicyAction(v string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.PolicyAction = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetPolicyId(v string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.PolicyId = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetPriority(v int64) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.Priority = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetTrustedProcessGroupIds(v []*string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.TrustedProcessGroupIds = v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetTrustedSoftwareIds(v []*string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.TrustedSoftwareIds = v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetUserGroupMode(v string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.UserGroupMode = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) Validate() error {
	if s.CustomUserAttributes != nil {
		for _, item := range s.CustomUserAttributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes struct {
	// The identity provider ID of the user group. This value exists when the custom user group type is **department**.
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// The relationship of the user group. Valid values:
	//
	// - **Equal**: equal to.
	//
	// - **Unequal**: not equal to.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// The type of the user group. Valid values:
	//
	// - **username**: the username.
	//
	// - **department**: the department.
	//
	// - **email**: the email address.
	//
	// - **telephone**: the mobile phone number.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// The value of the user group attribute.
	//
	// - When the user group type is **username**, this parameter indicates the value of the username. The value must be 1 to 128 characters in length, and can contain Chinese characters, uppercase letters, lowercase letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// - When the user group type is **department**, this parameter indicates the value of the department. For example: OU=Department1,OU=SASE DingTalk.
	//
	// - When the user group type is **email**, this parameter indicates the value of the email address. For example: username@example.com.
	//
	// - When the user group type is **telephone**, this parameter indicates the value of the mobile phone number. For example: 13900001234.
	//
	// example:
	//
	// OU=部门1,OU=SASE钉钉
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) String() string {
	return dara.Prettify(s)
}

func (s ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) GetRelation() *string {
	return s.Relation
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) GetValue() *string {
	return s.Value
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) SetIdpId(v int32) *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) SetRelation(v string) *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) SetUserGroupType(v string) *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) SetValue(v string) *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes {
	s.Value = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) Validate() error {
	return dara.Validate(s)
}
