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
	SetName(v string) *UpdatePrivateAccessPolicyRequest
	GetName() *string
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
	SetValidFrom(v int64) *UpdatePrivateAccessPolicyRequest
	GetValidFrom() *int64
	SetValidTimeStatus(v string) *UpdatePrivateAccessPolicyRequest
	GetValidTimeStatus() *string
	SetValidUntil(v int64) *UpdatePrivateAccessPolicyRequest
	GetValidUntil() *int64
}

type UpdatePrivateAccessPolicyRequest struct {
	// The IDs of applications associated with the internal network access policy. A single policy supports up to 100 application IDs.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The application type for the internal network access policy. Valid values:
	//
	// - **Application**: Application.
	//
	// - **Tag**: Tag.
	//
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// This parameter specifies a collection of custom user groups for the private network access policy. It is required when the user group type is **Custom**. This collection is mutually exclusive with the user group ID collection, and you can specify a maximum of 10 custom user groups.
	CustomUserAttributes []*UpdatePrivateAccessPolicyRequestCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	// A description of the internal network access policy. The description must be 1 to 128 characters in length. It can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), hyphens (-), and spaces.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 这是一条内网访问策略
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The action taken when a device does not meet the security baseline. Valid values:
	//
	// - **Block**: Block access.
	//
	// - **Observe**: Monitor access.
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
	// The method used to update the internal network access policy. Valid values:
	//
	// - **Cover*	- (default): Replace the existing application IDs, user group IDs, and custom user attributes with the values specified in **ApplicationIds**, **UserGroupIds**, and **CustomUserAttributes**.
	//
	// - **Append**: Add the values specified in **ApplicationIds**, **UserGroupIds**, and **CustomUserAttributes*	- to the existing application IDs, user group IDs, and custom user attributes.
	//
	// example:
	//
	// Cover
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The action that the internal network access policy takes. Valid values:
	//
	// - **Block**: Block access.
	//
	// - **Allow**: Allow access.
	//
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// The ID of the internal network access policy. Get this value from one of the following operations:
	//
	// - [ListPrivateAccessPolices](~~ListPrivateAccessPolices~~): List internal network access policies in batches.
	//
	// - [CreatePrivateAccessPolicy](~~CreatePrivateAccessPolicy~~): Create an internal network access policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The priority of the internal network access policy. Priority 1 is the highest. Valid values: 1 to 1000. The maximum value is the total number of internal network access policies minus 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of the internal network access policy. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IDs of tags associated with the internal network access policy. A single policy supports up to 100 tag IDs.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The ID of the trigger template.
	//
	// example:
	//
	// dag-d3f64e8bdd4a****
	TriggerTemplateId *string `json:"TriggerTemplateId,omitempty" xml:"TriggerTemplateId,omitempty"`
	// The IDs of trusted process groups.
	//
	// if can be null:
	// false
	TrustedProcessGroupIds []*string `json:"TrustedProcessGroupIds,omitempty" xml:"TrustedProcessGroupIds,omitempty" type:"Repeated"`
	// The status of the trusted process feature. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Disabled
	TrustedProcessStatus *string `json:"TrustedProcessStatus,omitempty" xml:"TrustedProcessStatus,omitempty"`
	// The IDs of trusted software.
	//
	// if can be null:
	// false
	TrustedSoftwareIds []*string `json:"TrustedSoftwareIds,omitempty" xml:"TrustedSoftwareIds,omitempty" type:"Repeated"`
	// The IDs of user groups associated with the internal network access policy. This parameter is required when UserGroupMode is set to Normal. This parameter is mutually exclusive with **CustomUserAttributes**. A single policy supports up to 10,000 user groups. You can update up to 2,000 user group IDs at a time.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The user group type for the internal network access policy. Valid values:
	//
	// - **Normal**: Regular user group.
	//
	// - **Custom**: Custom user group.
	//
	// example:
	//
	// Normal
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
	// The start time of the zero-trust policy\\"s effective period, in seconds since the Unix epoch.
	//
	// example:
	//
	// 0
	ValidFrom *int64 `json:"ValidFrom,omitempty" xml:"ValidFrom,omitempty"`
	// The status of the effective time feature. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	ValidTimeStatus *string `json:"ValidTimeStatus,omitempty" xml:"ValidTimeStatus,omitempty"`
	// The end time of the zero-trust policy\\"s effective period, in seconds since the Unix epoch.
	//
	// example:
	//
	// 1764727544
	ValidUntil *int64 `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
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

func (s *UpdatePrivateAccessPolicyRequest) GetName() *string {
	return s.Name
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

func (s *UpdatePrivateAccessPolicyRequest) GetValidFrom() *int64 {
	return s.ValidFrom
}

func (s *UpdatePrivateAccessPolicyRequest) GetValidTimeStatus() *string {
	return s.ValidTimeStatus
}

func (s *UpdatePrivateAccessPolicyRequest) GetValidUntil() *int64 {
	return s.ValidUntil
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

func (s *UpdatePrivateAccessPolicyRequest) SetName(v string) *UpdatePrivateAccessPolicyRequest {
	s.Name = &v
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

func (s *UpdatePrivateAccessPolicyRequest) SetValidFrom(v int64) *UpdatePrivateAccessPolicyRequest {
	s.ValidFrom = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetValidTimeStatus(v string) *UpdatePrivateAccessPolicyRequest {
	s.ValidTimeStatus = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetValidUntil(v int64) *UpdatePrivateAccessPolicyRequest {
	s.ValidUntil = &v
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
	// The identity provider ID for the custom user attribute. This parameter is required when UserGroupType is **department**.
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// The relation used to match the custom user attribute. Valid values:
	//
	// - **Equal**: Equal to.
	//
	// - **Unequal**: Not equal to.
	//
	// This parameter is required.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// The type of the custom user attribute. Valid values:
	//
	// - **username**: Username.
	//
	// - **department**: Department.
	//
	// - **email**: Email address.
	//
	// - **telephone**: Phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// The value of the custom user attribute.
	//
	// - If UserGroupType is **username**, this is the username. The value must be 1 to 128 characters in length. It can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), hyphens (-), asterisks (\\*), at signs (@), and spaces.
	//
	// - If UserGroupType is **department**, this is the department name. Example: OU=Department 1,OU=SASE DingTalk.
	//
	// - If UserGroupType is **email**, this is the email address. Example: username\\@example.com.
	//
	// - If UserGroupType is **telephone**, this is the phone number. Example: 13900001234.
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
