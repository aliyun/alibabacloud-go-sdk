// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessPolicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolices(v []*ListPrivateAccessPolicesResponseBodyPolices) *ListPrivateAccessPolicesResponseBody
	GetPolices() []*ListPrivateAccessPolicesResponseBodyPolices
	SetRequestId(v string) *ListPrivateAccessPolicesResponseBody
	GetRequestId() *string
	SetTotalNum(v int32) *ListPrivateAccessPolicesResponseBody
	GetTotalNum() *int32
}

type ListPrivateAccessPolicesResponseBody struct {
	// The list of private access policies.
	Polices []*ListPrivateAccessPolicesResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 748CFDC7-1EB6-5B8B-9405-DA76ED5BB60D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of private access policies.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListPrivateAccessPolicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessPolicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessPolicesResponseBody) GetPolices() []*ListPrivateAccessPolicesResponseBodyPolices {
	return s.Polices
}

func (s *ListPrivateAccessPolicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivateAccessPolicesResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListPrivateAccessPolicesResponseBody) SetPolices(v []*ListPrivateAccessPolicesResponseBodyPolices) *ListPrivateAccessPolicesResponseBody {
	s.Polices = v
	return s
}

func (s *ListPrivateAccessPolicesResponseBody) SetRequestId(v string) *ListPrivateAccessPolicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBody) SetTotalNum(v int32) *ListPrivateAccessPolicesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBody) Validate() error {
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

type ListPrivateAccessPolicesResponseBodyPolices struct {
	// The collection of application IDs of the private access policy. This field has a value when the application type is **Application**.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The application type of the private access policy. Valid values:
	//
	// - **Application**: Application.
	//
	// - **Tag**: Tag.
	//
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The creation time of the private access policy.
	//
	// example:
	//
	// 2022-07-10 15:50:23
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The collection of custom user group attributes. Multiple custom user group attributes have an OR relationship and take effect by union.
	CustomUserAttributes []*ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	// The description of the private access policy.
	//
	// example:
	//
	// 这是一条内网访问策略
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The action to take when the security baseline is not met. Valid values:
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
	// The name of the private access policy.
	//
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The action of the private access policy. Valid values:
	//
	// - **Block**: Block.
	//
	// - **Allow**: Allow.
	//
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// The ID of the private access policy.
	//
	// example:
	//
	// pa-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The priority of the private access policy. A value of 1 indicates the highest priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of the private access policy. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The collection of tag IDs of the private access policy. This field has a value when the application type is **Tag**.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The ID of the trigger template.
	//
	// example:
	//
	// dag-d3f64e8bdd4a****
	TriggerTemplateId *string `json:"TriggerTemplateId,omitempty" xml:"TriggerTemplateId,omitempty"`
	// The list of trusted process group IDs.
	TrustedProcessGroupIds []*string `json:"TrustedProcessGroupIds,omitempty" xml:"TrustedProcessGroupIds,omitempty" type:"Repeated"`
	// The status of the trusted process switch. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	TrustedProcessStatus *string `json:"TrustedProcessStatus,omitempty" xml:"TrustedProcessStatus,omitempty"`
	// The list of trusted software IDs.
	TrustedSoftwareIds []*string `json:"TrustedSoftwareIds,omitempty" xml:"TrustedSoftwareIds,omitempty" type:"Repeated"`
	// The collection of user group IDs for the private access policy. This field has a value when the user group type is **Normal**.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The user group type of the private access policy. Valid values:
	//
	// - **Normal**: Normal user group.
	//
	// - **Custom**: Custom user group.
	//
	// example:
	//
	// Normal
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
	// The effective start time of the zero trust policy, in second-level Unix timestamp.
	//
	// example:
	//
	// 0
	ValidFrom *int64 `json:"ValidFrom,omitempty" xml:"ValidFrom,omitempty"`
	// The status of the effective time switch. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	ValidTimeStatus *string `json:"ValidTimeStatus,omitempty" xml:"ValidTimeStatus,omitempty"`
	// The effective end time of the zero trust policy, in second-level Unix timestamp.
	//
	// example:
	//
	// 1764727544
	ValidUntil *int64 `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
}

func (s ListPrivateAccessPolicesResponseBodyPolices) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessPolicesResponseBodyPolices) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetCustomUserAttributes() []*ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes {
	return s.CustomUserAttributes
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetDescription() *string {
	return s.Description
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetDeviceAttributeAction() *string {
	return s.DeviceAttributeAction
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetDeviceAttributeId() *string {
	return s.DeviceAttributeId
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetName() *string {
	return s.Name
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetPolicyAction() *string {
	return s.PolicyAction
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetPriority() *int32 {
	return s.Priority
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetStatus() *string {
	return s.Status
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetTagIds() []*string {
	return s.TagIds
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetTriggerTemplateId() *string {
	return s.TriggerTemplateId
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetTrustedProcessGroupIds() []*string {
	return s.TrustedProcessGroupIds
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetTrustedProcessStatus() *string {
	return s.TrustedProcessStatus
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetTrustedSoftwareIds() []*string {
	return s.TrustedSoftwareIds
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetUserGroupMode() *string {
	return s.UserGroupMode
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetValidFrom() *int64 {
	return s.ValidFrom
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetValidTimeStatus() *string {
	return s.ValidTimeStatus
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) GetValidUntil() *int64 {
	return s.ValidUntil
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetApplicationIds(v []*string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.ApplicationIds = v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetApplicationType(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.ApplicationType = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetCreateTime(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.CreateTime = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetCustomUserAttributes(v []*ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) *ListPrivateAccessPolicesResponseBodyPolices {
	s.CustomUserAttributes = v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetDescription(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.Description = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetDeviceAttributeAction(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.DeviceAttributeAction = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetDeviceAttributeId(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.DeviceAttributeId = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetName(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetPolicyAction(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.PolicyAction = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetPolicyId(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.PolicyId = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetPriority(v int32) *ListPrivateAccessPolicesResponseBodyPolices {
	s.Priority = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetStatus(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.Status = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetTagIds(v []*string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.TagIds = v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetTriggerTemplateId(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.TriggerTemplateId = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetTrustedProcessGroupIds(v []*string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.TrustedProcessGroupIds = v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetTrustedProcessStatus(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.TrustedProcessStatus = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetTrustedSoftwareIds(v []*string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.TrustedSoftwareIds = v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetUserGroupIds(v []*string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.UserGroupIds = v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetUserGroupMode(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.UserGroupMode = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetValidFrom(v int64) *ListPrivateAccessPolicesResponseBodyPolices {
	s.ValidFrom = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetValidTimeStatus(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.ValidTimeStatus = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetValidUntil(v int64) *ListPrivateAccessPolicesResponseBodyPolices {
	s.ValidUntil = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) Validate() error {
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

type ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes struct {
	// The identity provider ID of the user group. This value exists when the custom user group type is **department**.
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// The relation of the user group. Valid values:
	//
	// - **Equal**: Equal.
	//
	// - **Unequal**: Not equal.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// The type of the user group. Valid values:
	//
	// - **username**: Username.
	//
	// - **department**: Department.
	//
	// - **email**: Email.
	//
	// - **telephone**: Mobile phone.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// The value of the user group attribute.
	//
	// - When the user group type is **username**, this indicates the value of the username. The value must be 1 to 128 characters in length and supports Chinese characters and uppercase and lowercase English letters. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// - When the user group type is **department**, this indicates the value of the department. For example: OU=Department1,OU=SASE DingTalk.
	//
	// - When the user group type is **email**, this indicates the value of the email. For example: username@example.com.
	//
	// - When the user group type is **telephone**, this indicates the value of the mobile phone. For example: 13900001234.
	//
	// example:
	//
	// OU=部门1,OU=SASE钉钉
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) GetRelation() *string {
	return s.Relation
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) GetValue() *string {
	return s.Value
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) SetIdpId(v int32) *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) SetRelation(v string) *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) SetUserGroupType(v string) *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) SetValue(v string) *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes {
	s.Value = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) Validate() error {
	return dara.Validate(s)
}
