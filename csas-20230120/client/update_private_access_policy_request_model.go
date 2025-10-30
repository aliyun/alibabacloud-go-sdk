// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *UpdatePrivateAccessPolicyRequest
	GetApplicationIds() []*string
	SetApplicationType(v string) *UpdatePrivateAccessPolicyRequest
	GetApplicationType() *string
	SetCustomUserAttributes(v []*UpdatePrivateAccessPolicyRequestCustomUserAttributes) *UpdatePrivateAccessPolicyRequest
	GetCustomUserAttributes() []*UpdatePrivateAccessPolicyRequestCustomUserAttributes
	SetDescription(v string) *UpdatePrivateAccessPolicyRequest
	GetDescription() *string
	SetDeviceAttributeAction(v string) *UpdatePrivateAccessPolicyRequest
	GetDeviceAttributeAction() *string
	SetDeviceAttributeId(v string) *UpdatePrivateAccessPolicyRequest
	GetDeviceAttributeId() *string
	SetModifyType(v string) *UpdatePrivateAccessPolicyRequest
	GetModifyType() *string
	SetPolicyAction(v string) *UpdatePrivateAccessPolicyRequest
	GetPolicyAction() *string
	SetPolicyId(v string) *UpdatePrivateAccessPolicyRequest
	GetPolicyId() *string
	SetPriority(v int32) *UpdatePrivateAccessPolicyRequest
	GetPriority() *int32
	SetStatus(v string) *UpdatePrivateAccessPolicyRequest
	GetStatus() *string
	SetTagIds(v []*string) *UpdatePrivateAccessPolicyRequest
	GetTagIds() []*string
	SetTriggerTemplateId(v string) *UpdatePrivateAccessPolicyRequest
	GetTriggerTemplateId() *string
	SetTrustedProcessGroupIds(v []*string) *UpdatePrivateAccessPolicyRequest
	GetTrustedProcessGroupIds() []*string
	SetTrustedProcessStatus(v string) *UpdatePrivateAccessPolicyRequest
	GetTrustedProcessStatus() *string
	SetTrustedSoftwareIds(v []*string) *UpdatePrivateAccessPolicyRequest
	GetTrustedSoftwareIds() []*string
	SetUserGroupIds(v []*string) *UpdatePrivateAccessPolicyRequest
	GetUserGroupIds() []*string
	SetUserGroupMode(v string) *UpdatePrivateAccessPolicyRequest
	GetUserGroupMode() *string
}

type UpdatePrivateAccessPolicyRequest struct {
	// Set of application IDs for the private access policy. A single policy supports up to 100 private access application IDs.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// Application type of the private access policy. Values:
	//
	// - **Application**: Application.
	//
	// - **Tag**: Tag.
	//
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// Set of custom user attributes for the private access policy, required when the user group type is **Custom**. Mutually exclusive with the user group ID set. The total number of custom user groups is limited to 10.
	CustomUserAttributes []*UpdatePrivateAccessPolicyRequestCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	// Description of the private access policy. Length should be 1 to 128 characters, supporting Chinese and English letters (both uppercase and lowercase), and can include numbers, periods (.), underscores (_), hyphens (-), and spaces.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution strategy for not meeting the security baseline. Values:
	//
	// - **Block**: Block.
	//
	// - **Observe**: Observe.
	//
	// example:
	//
	// Block
	DeviceAttributeAction *string `json:"DeviceAttributeAction,omitempty" xml:"DeviceAttributeAction,omitempty"`
	// The ID of the security baseline policy.
	//
	// example:
	//
	// dag-d3f64e8bdd4a****
	DeviceAttributeId *string `json:"DeviceAttributeId,omitempty" xml:"DeviceAttributeId,omitempty"`
	// The modification type of the private access policy. Values:
	//
	// - **Cover*	- (default): Use the values of **ApplicationIds**, **UserGroupIds**, and **CustomUserAttributes*	- to overwrite the original application ID set, user group ID set, and custom user attribute set, respectively.
	//
	// - **Append**: Add the values provided in **ApplicationIds**, **UserGroupIds**, and **CustomUserAttributes*	- to the original application ID set, user group ID set, and custom user attribute set, respectively.
	//
	// example:
	//
	// Cover
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// Action of the private access policy. Values:
	//
	// - **Block**: Block.
	//
	// - **Allow**: Allow.
	//
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// ID of the private access policy. Value sources:
	//
	// - [ListPrivateAccessPolicies](~~ListPrivateAccessPolicies~~): Batch query for private access policies.
	//
	// - [CreatePrivateAccessPolicy](~~CreatePrivateAccessPolicy~~): Create a private access policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The priority of the private access policy. The number 1 indicates the highest priority. Range: 1~1000, with the maximum value being the total number of private access policies minus one.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of the private access policy. Values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Set of tag IDs for the private access policy. A single policy supports up to 100 private access tag IDs.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The trigger template ID.
	//
	// example:
	//
	// dag-d3f64e8bdd4a****
	TriggerTemplateId *string `json:"TriggerTemplateId,omitempty" xml:"TriggerTemplateId,omitempty"`
	// Trusted process group ID.
	//
	// if can be null:
	// false
	TrustedProcessGroupIds []*string `json:"TrustedProcessGroupIds,omitempty" xml:"TrustedProcessGroupIds,omitempty" type:"Repeated"`
	// Trusted process switch status. Values:
	//
	// - **Enabled**: On.
	//
	// - **Disabled**: Off.
	//
	// example:
	//
	// Disabled
	TrustedProcessStatus *string `json:"TrustedProcessStatus,omitempty" xml:"TrustedProcessStatus,omitempty"`
	// Trusted Software ID.
	//
	// if can be null:
	// false
	TrustedSoftwareIds []*string `json:"TrustedSoftwareIds,omitempty" xml:"TrustedSoftwareIds,omitempty" type:"Repeated"`
	// Set of user group IDs for the private access policy, required when the user group type is **Normal**. Mutually exclusive with the custom user group set. A single policy supports up to 10,000 user groups, and a maximum of 2,000 user group IDs can be modified at once.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// User group type of the private access policy. Values:
	//
	// - **Normal**: Normal user group.
	//
	// - **Custom**: Custom user group.
	//
	// example:
	//
	// Normal
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
}

func (s UpdatePrivateAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessPolicyRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *UpdatePrivateAccessPolicyRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *UpdatePrivateAccessPolicyRequest) GetCustomUserAttributes() []*UpdatePrivateAccessPolicyRequestCustomUserAttributes {
	return s.CustomUserAttributes
}

func (s *UpdatePrivateAccessPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePrivateAccessPolicyRequest) GetDeviceAttributeAction() *string {
	return s.DeviceAttributeAction
}

func (s *UpdatePrivateAccessPolicyRequest) GetDeviceAttributeId() *string {
	return s.DeviceAttributeId
}

func (s *UpdatePrivateAccessPolicyRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *UpdatePrivateAccessPolicyRequest) GetPolicyAction() *string {
	return s.PolicyAction
}

func (s *UpdatePrivateAccessPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdatePrivateAccessPolicyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdatePrivateAccessPolicyRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdatePrivateAccessPolicyRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *UpdatePrivateAccessPolicyRequest) GetTriggerTemplateId() *string {
	return s.TriggerTemplateId
}

func (s *UpdatePrivateAccessPolicyRequest) GetTrustedProcessGroupIds() []*string {
	return s.TrustedProcessGroupIds
}

func (s *UpdatePrivateAccessPolicyRequest) GetTrustedProcessStatus() *string {
	return s.TrustedProcessStatus
}

func (s *UpdatePrivateAccessPolicyRequest) GetTrustedSoftwareIds() []*string {
	return s.TrustedSoftwareIds
}

func (s *UpdatePrivateAccessPolicyRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdatePrivateAccessPolicyRequest) GetUserGroupMode() *string {
	return s.UserGroupMode
}

func (s *UpdatePrivateAccessPolicyRequest) SetApplicationIds(v []*string) *UpdatePrivateAccessPolicyRequest {
	s.ApplicationIds = v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetApplicationType(v string) *UpdatePrivateAccessPolicyRequest {
	s.ApplicationType = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetCustomUserAttributes(v []*UpdatePrivateAccessPolicyRequestCustomUserAttributes) *UpdatePrivateAccessPolicyRequest {
	s.CustomUserAttributes = v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetDescription(v string) *UpdatePrivateAccessPolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetDeviceAttributeAction(v string) *UpdatePrivateAccessPolicyRequest {
	s.DeviceAttributeAction = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetDeviceAttributeId(v string) *UpdatePrivateAccessPolicyRequest {
	s.DeviceAttributeId = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetModifyType(v string) *UpdatePrivateAccessPolicyRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetPolicyAction(v string) *UpdatePrivateAccessPolicyRequest {
	s.PolicyAction = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetPolicyId(v string) *UpdatePrivateAccessPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetPriority(v int32) *UpdatePrivateAccessPolicyRequest {
	s.Priority = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetStatus(v string) *UpdatePrivateAccessPolicyRequest {
	s.Status = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetTagIds(v []*string) *UpdatePrivateAccessPolicyRequest {
	s.TagIds = v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetTriggerTemplateId(v string) *UpdatePrivateAccessPolicyRequest {
	s.TriggerTemplateId = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetTrustedProcessGroupIds(v []*string) *UpdatePrivateAccessPolicyRequest {
	s.TrustedProcessGroupIds = v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetTrustedProcessStatus(v string) *UpdatePrivateAccessPolicyRequest {
	s.TrustedProcessStatus = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetTrustedSoftwareIds(v []*string) *UpdatePrivateAccessPolicyRequest {
	s.TrustedSoftwareIds = v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetUserGroupIds(v []*string) *UpdatePrivateAccessPolicyRequest {
	s.UserGroupIds = v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetUserGroupMode(v string) *UpdatePrivateAccessPolicyRequest {
	s.UserGroupMode = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) Validate() error {
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

type UpdatePrivateAccessPolicyRequestCustomUserAttributes struct {
	// The identity source ID of the custom user group. Required when the custom user group type is **department**.
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// Relation of the custom user group. Values:
	//
	// - **Equal**: Equal.
	//
	// - **Unequal**: Not equal.
	//
	// This parameter is required.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// Type of the custom user group. Values:
	//
	// - **username**: Username.
	//
	// - **department**: Department.
	//
	// - **email**: Email.
	//
	// - **telephone**: Telephone.
	//
	// This parameter is required.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// Custom user group attribute values. - When the user group type is **username**, it represents the value of the username. The length should be 1 to 128 characters, supporting Chinese and case-sensitive English letters, and can include numbers, half-width periods (.), underscores (_), hyphens (-), asterisks (*), at symbols (@), and spaces. - When the user group type is **department**, it represents the value of the department. For example: OU=Department1,OU=SASE DingTalk. - When the user group type is **email**, it represents the value of the email. For example: username@example.com. - When the user group type is **telephone**, it represents the value of the mobile phone. For example: 13900001234.
	//
	// This parameter is required.
	//
	// example:
	//
	// OU=部门1,OU=SASE钉钉
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePrivateAccessPolicyRequestCustomUserAttributes) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessPolicyRequestCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) GetRelation() *string {
	return s.Relation
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) GetValue() *string {
	return s.Value
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) SetIdpId(v int32) *UpdatePrivateAccessPolicyRequestCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) SetRelation(v string) *UpdatePrivateAccessPolicyRequestCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) SetUserGroupType(v string) *UpdatePrivateAccessPolicyRequestCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) SetValue(v string) *UpdatePrivateAccessPolicyRequestCustomUserAttributes {
	s.Value = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) Validate() error {
	return dara.Validate(s)
}
