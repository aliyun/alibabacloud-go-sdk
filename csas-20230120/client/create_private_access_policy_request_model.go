// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *CreatePrivateAccessPolicyRequest
	GetApplicationIds() []*string
	SetApplicationType(v string) *CreatePrivateAccessPolicyRequest
	GetApplicationType() *string
	SetCustomUserAttributes(v []*CreatePrivateAccessPolicyRequestCustomUserAttributes) *CreatePrivateAccessPolicyRequest
	GetCustomUserAttributes() []*CreatePrivateAccessPolicyRequestCustomUserAttributes
	SetDescription(v string) *CreatePrivateAccessPolicyRequest
	GetDescription() *string
	SetDeviceAttributeAction(v string) *CreatePrivateAccessPolicyRequest
	GetDeviceAttributeAction() *string
	SetDeviceAttributeId(v string) *CreatePrivateAccessPolicyRequest
	GetDeviceAttributeId() *string
	SetName(v string) *CreatePrivateAccessPolicyRequest
	GetName() *string
	SetPolicyAction(v string) *CreatePrivateAccessPolicyRequest
	GetPolicyAction() *string
	SetPriority(v int32) *CreatePrivateAccessPolicyRequest
	GetPriority() *int32
	SetStatus(v string) *CreatePrivateAccessPolicyRequest
	GetStatus() *string
	SetTagIds(v []*string) *CreatePrivateAccessPolicyRequest
	GetTagIds() []*string
	SetTriggerTemplateId(v string) *CreatePrivateAccessPolicyRequest
	GetTriggerTemplateId() *string
	SetTrustedProcessGroupIds(v []*string) *CreatePrivateAccessPolicyRequest
	GetTrustedProcessGroupIds() []*string
	SetTrustedProcessStatus(v string) *CreatePrivateAccessPolicyRequest
	GetTrustedProcessStatus() *string
	SetTrustedSoftwareIds(v []*string) *CreatePrivateAccessPolicyRequest
	GetTrustedSoftwareIds() []*string
	SetUserGroupIds(v []*string) *CreatePrivateAccessPolicyRequest
	GetUserGroupIds() []*string
	SetUserGroupMode(v string) *CreatePrivateAccessPolicyRequest
	GetUserGroupMode() *string
	SetValidFrom(v int64) *CreatePrivateAccessPolicyRequest
	GetValidFrom() *int64
	SetValidTimeStatus(v string) *CreatePrivateAccessPolicyRequest
	GetValidTimeStatus() *string
	SetValidUntil(v int64) *CreatePrivateAccessPolicyRequest
	GetValidUntil() *int64
}

type CreatePrivateAccessPolicyRequest struct {
	// Set of application IDs for the private access policy. Up to 100 application IDs can be entered. Required when **ApplicationType*	- is **Application**. Mutually exclusive with **TagIds**.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// Application type of the private access policy. Values:
	//
	// - **Application**: Application.
	//
	// - **Tag**: Tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// Set of custom user groups for the private access policy. Required when the user group type is **Custom**. Mutually exclusive with the user group ID set. Up to 10 custom user groups can be entered.
	CustomUserAttributes []*CreatePrivateAccessPolicyRequestCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	// Description of the private access policy. The length is 1 to 128 characters, supporting Chinese and uppercase and lowercase English letters, and can include numbers, periods (.), underscores (_), hyphens (-), and spaces.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution policy for not meeting the security baseline. Values:
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
	// Name of the private access policy. The length is 1 to 128 characters, supporting Chinese and uppercase and lowercase English letters, and can include numbers, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Action of the private access policy. Values:
	//
	// - **Block**: Block.
	//
	// - **Allow**: Allow.
	//
	// This parameter is required.
	//
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// The priority of the private access policy. The number 1 indicates the highest priority. Range: 1~1000, with the maximum value being the total number of private access policies.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Set of tag IDs for the private access policy. Up to 100 tag IDs can be entered. Required when **ApplicationType*	- is **Tag**. Mutually exclusive with **ApplicationIds**.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The trigger template ID.
	//
	// example:
	//
	// dag-d3f64e8bdd4a****
	TriggerTemplateId *string `json:"TriggerTemplateId,omitempty" xml:"TriggerTemplateId,omitempty"`
	// The ID of the trusted process group.
	TrustedProcessGroupIds []*string `json:"TrustedProcessGroupIds,omitempty" xml:"TrustedProcessGroupIds,omitempty" type:"Repeated"`
	// The switch status of the trusted process. Values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	TrustedProcessStatus *string `json:"TrustedProcessStatus,omitempty" xml:"TrustedProcessStatus,omitempty"`
	// The ID of the trusted software.
	TrustedSoftwareIds []*string `json:"TrustedSoftwareIds,omitempty" xml:"TrustedSoftwareIds,omitempty" type:"Repeated"`
	// Set of user group IDs for the private access policy. Required when the user group type is **Normal**. Mutually exclusive with the custom user group set. Up to 2000 user group IDs can be entered.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// User group type of the private access policy. Values:
	//
	// - **Normal**: Normal user group.
	//
	// - **Custom**: Custom user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Normal
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
	// The start time when the zero trust policy takes effect, represented as a timestamp in seconds.
	//
	// example:
	//
	// 0
	ValidFrom *int64 `json:"ValidFrom,omitempty" xml:"ValidFrom,omitempty"`
	// Switch status for effective time. Values: - **Enabled**: On. - **Disabled**: Off.
	//
	// example:
	//
	// Disabled
	ValidTimeStatus *string `json:"ValidTimeStatus,omitempty" xml:"ValidTimeStatus,omitempty"`
	// The expiration time of the zero trust policy, in seconds timestamp.
	//
	// example:
	//
	// 1764727544
	ValidUntil *int64 `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
}

func (s CreatePrivateAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessPolicyRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *CreatePrivateAccessPolicyRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *CreatePrivateAccessPolicyRequest) GetCustomUserAttributes() []*CreatePrivateAccessPolicyRequestCustomUserAttributes {
	return s.CustomUserAttributes
}

func (s *CreatePrivateAccessPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePrivateAccessPolicyRequest) GetDeviceAttributeAction() *string {
	return s.DeviceAttributeAction
}

func (s *CreatePrivateAccessPolicyRequest) GetDeviceAttributeId() *string {
	return s.DeviceAttributeId
}

func (s *CreatePrivateAccessPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreatePrivateAccessPolicyRequest) GetPolicyAction() *string {
	return s.PolicyAction
}

func (s *CreatePrivateAccessPolicyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreatePrivateAccessPolicyRequest) GetStatus() *string {
	return s.Status
}

func (s *CreatePrivateAccessPolicyRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *CreatePrivateAccessPolicyRequest) GetTriggerTemplateId() *string {
	return s.TriggerTemplateId
}

func (s *CreatePrivateAccessPolicyRequest) GetTrustedProcessGroupIds() []*string {
	return s.TrustedProcessGroupIds
}

func (s *CreatePrivateAccessPolicyRequest) GetTrustedProcessStatus() *string {
	return s.TrustedProcessStatus
}

func (s *CreatePrivateAccessPolicyRequest) GetTrustedSoftwareIds() []*string {
	return s.TrustedSoftwareIds
}

func (s *CreatePrivateAccessPolicyRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreatePrivateAccessPolicyRequest) GetUserGroupMode() *string {
	return s.UserGroupMode
}

func (s *CreatePrivateAccessPolicyRequest) GetValidFrom() *int64 {
	return s.ValidFrom
}

func (s *CreatePrivateAccessPolicyRequest) GetValidTimeStatus() *string {
	return s.ValidTimeStatus
}

func (s *CreatePrivateAccessPolicyRequest) GetValidUntil() *int64 {
	return s.ValidUntil
}

func (s *CreatePrivateAccessPolicyRequest) SetApplicationIds(v []*string) *CreatePrivateAccessPolicyRequest {
	s.ApplicationIds = v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetApplicationType(v string) *CreatePrivateAccessPolicyRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetCustomUserAttributes(v []*CreatePrivateAccessPolicyRequestCustomUserAttributes) *CreatePrivateAccessPolicyRequest {
	s.CustomUserAttributes = v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetDescription(v string) *CreatePrivateAccessPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetDeviceAttributeAction(v string) *CreatePrivateAccessPolicyRequest {
	s.DeviceAttributeAction = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetDeviceAttributeId(v string) *CreatePrivateAccessPolicyRequest {
	s.DeviceAttributeId = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetName(v string) *CreatePrivateAccessPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetPolicyAction(v string) *CreatePrivateAccessPolicyRequest {
	s.PolicyAction = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetPriority(v int32) *CreatePrivateAccessPolicyRequest {
	s.Priority = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetStatus(v string) *CreatePrivateAccessPolicyRequest {
	s.Status = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetTagIds(v []*string) *CreatePrivateAccessPolicyRequest {
	s.TagIds = v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetTriggerTemplateId(v string) *CreatePrivateAccessPolicyRequest {
	s.TriggerTemplateId = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetTrustedProcessGroupIds(v []*string) *CreatePrivateAccessPolicyRequest {
	s.TrustedProcessGroupIds = v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetTrustedProcessStatus(v string) *CreatePrivateAccessPolicyRequest {
	s.TrustedProcessStatus = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetTrustedSoftwareIds(v []*string) *CreatePrivateAccessPolicyRequest {
	s.TrustedSoftwareIds = v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetUserGroupIds(v []*string) *CreatePrivateAccessPolicyRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetUserGroupMode(v string) *CreatePrivateAccessPolicyRequest {
	s.UserGroupMode = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetValidFrom(v int64) *CreatePrivateAccessPolicyRequest {
	s.ValidFrom = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetValidTimeStatus(v string) *CreatePrivateAccessPolicyRequest {
	s.ValidTimeStatus = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetValidUntil(v int64) *CreatePrivateAccessPolicyRequest {
	s.ValidUntil = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) Validate() error {
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

type CreatePrivateAccessPolicyRequestCustomUserAttributes struct {
	// The ID of the identity source for the custom user group. Required when the custom user group type is **department**.
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
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// Value of the custom user group attribute.
	//
	// - When the user group type is **username**, it represents the value of the username. The length is 1 to 128 characters, supporting Chinese and uppercase and lowercase English letters, and can include numbers, periods (.), underscores (_), hyphens (-), asterisks (*), at (@) symbols, and spaces.
	//
	// - When the user group type is **department**, it represents the value of the department. For example: OU=Department1,OU=SASE DingTalk.
	//
	// - When the user group type is **email**, it represents the value of the email. For example: username@example.com.
	//
	// - When the user group type is **telephone**, it represents the value of the telephone. For example: 13900001234.
	//
	// example:
	//
	// OU=部门1,OU=SASE钉钉
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrivateAccessPolicyRequestCustomUserAttributes) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessPolicyRequestCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) GetRelation() *string {
	return s.Relation
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) GetValue() *string {
	return s.Value
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) SetIdpId(v int32) *CreatePrivateAccessPolicyRequestCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) SetRelation(v string) *CreatePrivateAccessPolicyRequestCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) SetUserGroupType(v string) *CreatePrivateAccessPolicyRequestCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) SetValue(v string) *CreatePrivateAccessPolicyRequestCustomUserAttributes {
	s.Value = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) Validate() error {
	return dara.Validate(s)
}
