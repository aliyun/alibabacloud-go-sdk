// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrivateAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *GetPrivateAccessPolicyResponseBodyPolicy) *GetPrivateAccessPolicyResponseBody
	GetPolicy() *GetPrivateAccessPolicyResponseBodyPolicy
	SetRequestId(v string) *GetPrivateAccessPolicyResponseBody
	GetRequestId() *string
}

type GetPrivateAccessPolicyResponseBody struct {
	// Intranet access policy.
	Policy *GetPrivateAccessPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The ID of the current request.
	//
	// example:
	//
	// 7E9D7ACD-53D5-56EF-A913-79D148D06299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPrivateAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessPolicyResponseBody) GetPolicy() *GetPrivateAccessPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *GetPrivateAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPrivateAccessPolicyResponseBody) SetPolicy(v *GetPrivateAccessPolicyResponseBodyPolicy) *GetPrivateAccessPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GetPrivateAccessPolicyResponseBody) SetRequestId(v string) *GetPrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPrivateAccessPolicyResponseBodyPolicy struct {
	// A collection of application IDs for the private access policy. This field has a value when the application type is Application.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The application type of the private access policy. Possible values:
	//
	// - **Application**: Application.
	//
	// - **Tag**: Tag.
	//
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// Intranet access policy creation time.
	//
	// example:
	//
	// 2021-07-29 11:26:02
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Collection of custom user group attributes. Multiple custom user group attributes are combined with an OR relationship and take effect as a set.
	CustomUserAttributes []*GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	// Intranet access policy description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The action to take if the security baseline is not met. Possible values:
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
	// Intranet access policy name.
	//
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Intranet access policy action. Values:
	//
	// - **Block**: Block.
	//
	// - **Allow**: Allow.
	//
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// Intranet access policy ID.
	//
	// example:
	//
	// pa-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Intranet access policy priority. A value of 1 indicates the highest priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Intranet access policy status. Values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// A collection of tag IDs for the private access policy. This field has a value when the application type is Tag.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The trigger template ID.
	//
	// example:
	//
	// dag-d3f64e8bdd4a****
	TriggerTemplateId *string `json:"TriggerTemplateId,omitempty" xml:"TriggerTemplateId,omitempty"`
	// A list of trusted process group IDs.
	TrustedProcessGroupIds []*string `json:"TrustedProcessGroupIds,omitempty" xml:"TrustedProcessGroupIds,omitempty" type:"Repeated"`
	// The status of the trusted process switch. Possible values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	TrustedProcessStatus *string `json:"TrustedProcessStatus,omitempty" xml:"TrustedProcessStatus,omitempty"`
	// A list of trusted software IDs.
	TrustedSoftwareIds []*string `json:"TrustedSoftwareIds,omitempty" xml:"TrustedSoftwareIds,omitempty" type:"Repeated"`
	// Collection of user group IDs for the intranet access policy. This field is populated when the user group type is Normal.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// User group type for the intranet access policy. Values:
	//
	// - **Normal**: Normal user group.
	//
	// - **Custom**: Custom user group.
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
	// Enabled
	ValidTimeStatus *string `json:"ValidTimeStatus,omitempty" xml:"ValidTimeStatus,omitempty"`
	// The expiration time of the zero trust policy, in seconds timestamp.
	//
	// example:
	//
	// 1764727544
	ValidUntil *int64 `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
}

func (s GetPrivateAccessPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetPrivateAccessPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetCustomUserAttributes() []*GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes {
	return s.CustomUserAttributes
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetDescription() *string {
	return s.Description
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetDeviceAttributeAction() *string {
	return s.DeviceAttributeAction
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetDeviceAttributeId() *string {
	return s.DeviceAttributeId
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetName() *string {
	return s.Name
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetPolicyAction() *string {
	return s.PolicyAction
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetPriority() *int32 {
	return s.Priority
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetStatus() *string {
	return s.Status
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetTagIds() []*string {
	return s.TagIds
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetTriggerTemplateId() *string {
	return s.TriggerTemplateId
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetTrustedProcessGroupIds() []*string {
	return s.TrustedProcessGroupIds
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetTrustedProcessStatus() *string {
	return s.TrustedProcessStatus
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetTrustedSoftwareIds() []*string {
	return s.TrustedSoftwareIds
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetUserGroupMode() *string {
	return s.UserGroupMode
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetValidFrom() *int64 {
	return s.ValidFrom
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetValidTimeStatus() *string {
	return s.ValidTimeStatus
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) GetValidUntil() *int64 {
	return s.ValidUntil
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetApplicationIds(v []*string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.ApplicationIds = v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetApplicationType(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.ApplicationType = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetCreateTime(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.CreateTime = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetCustomUserAttributes(v []*GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.CustomUserAttributes = v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetDescription(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetDeviceAttributeAction(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.DeviceAttributeAction = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetDeviceAttributeId(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.DeviceAttributeId = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetName(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.Name = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetPolicyAction(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.PolicyAction = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetPolicyId(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetPriority(v int32) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.Priority = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetStatus(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.Status = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetTagIds(v []*string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.TagIds = v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetTriggerTemplateId(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.TriggerTemplateId = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetTrustedProcessGroupIds(v []*string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.TrustedProcessGroupIds = v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetTrustedProcessStatus(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.TrustedProcessStatus = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetTrustedSoftwareIds(v []*string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.TrustedSoftwareIds = v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetUserGroupIds(v []*string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.UserGroupIds = v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetUserGroupMode(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.UserGroupMode = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetValidFrom(v int64) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.ValidFrom = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetValidTimeStatus(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.ValidTimeStatus = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetValidUntil(v int64) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.ValidUntil = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) Validate() error {
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

type GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes struct {
	// The identity provider ID for the custom user group. This field is required when the custom user group type is **department**.
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// Custom user group relationship. Values: - **Equal**: Equal to. - **Unequal**: Not equal to.
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
	// The value of the custom user group attribute.
	//
	// - When the user group type is **username**, it represents the value of the username. The length should be 1 to 128 characters, supporting Chinese and case-sensitive English letters, and can include numbers, periods (.), underscores (_), and hyphens (-).
	//
	// - When the user group type is **department**, it represents the value of the department. For example: OU=Department1,OU=SASE DingTalk.
	//
	// - When the user group type is **email**, it represents the value of the email. For example: username@example.com.
	//
	// - When the user group type is **telephone**, it represents the value of the phone number. For example: 13900001234.
	//
	// example:
	//
	// OU=部门1,OU=SASE钉钉
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) String() string {
	return dara.Prettify(s)
}

func (s GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) GetRelation() *string {
	return s.Relation
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) GetValue() *string {
	return s.Value
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) SetIdpId(v int32) *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) SetRelation(v string) *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) SetUserGroupType(v string) *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) SetValue(v string) *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes {
	s.Value = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) Validate() error {
	return dara.Validate(s)
}
