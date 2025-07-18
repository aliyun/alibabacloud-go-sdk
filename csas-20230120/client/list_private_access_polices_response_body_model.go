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
	// The private access policies.
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
	return dara.Validate(s)
}

type ListPrivateAccessPolicesResponseBodyPolices struct {
	// The IDs of the applications that are specified in the private access policy. If the value of ApplicationType is **Application**, this parameter is returned.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The application type of the private access policy. Valid values:
	//
	// 	- **Application**
	//
	// 	- **Tag**
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
	// The attributes of the custom user group. The attributes of the custom user group are evaluated by using a logical OR. If an attribute is matched, the policy takes effect.
	CustomUserAttributes []*ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	// The description of the private access policy.
	//
	// example:
	//
	// a private access policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The action that is performed when the security baseline is not met. Valid values:
	//
	// 	- **Block**
	//
	// 	- **Observe**
	//
	// example:
	//
	// Block
	DeviceAttributeAction *string `json:"DeviceAttributeAction,omitempty" xml:"DeviceAttributeAction,omitempty"`
	// The ID of the security baseline.
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
	// The action in the private access policy. Valid values:
	//
	// 	- **Block**
	//
	// 	- **Allow**
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
	// The priority of the private access policy. The value 1 indicates the highest priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of the private access policy. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IDs of the tags that are specified in the private access policy. If the value of ApplicationType is **Tag**, this parameter is returned.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The ID of the trigger template.
	//
	// example:
	//
	// dag-d3f64e8bdd4a****
	TriggerTemplateId *string `json:"TriggerTemplateId,omitempty" xml:"TriggerTemplateId,omitempty"`
	// List of trusted process group IDs.
	TrustedProcessGroupIds []*string `json:"TrustedProcessGroupIds,omitempty" xml:"TrustedProcessGroupIds,omitempty" type:"Repeated"`
	// Trusted process switch status. Values:
	//
	// - **Enabled**: On.
	//
	// - **Disabled**: Off.
	//
	// example:
	//
	// Enabled
	TrustedProcessStatus *string `json:"TrustedProcessStatus,omitempty" xml:"TrustedProcessStatus,omitempty"`
	// List of trusted software IDs.
	TrustedSoftwareIds []*string `json:"TrustedSoftwareIds,omitempty" xml:"TrustedSoftwareIds,omitempty" type:"Repeated"`
	// The IDs of user groups in the private access policy. If the value of UserGroupMode is **Normal**, this parameter is returned.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The type of the user group in the private access policy. Valid values:
	//
	// 	- **Normal**: regular user group.
	//
	// 	- **Custom**: custom user group.
	//
	// example:
	//
	// Normal
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
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

func (s *ListPrivateAccessPolicesResponseBodyPolices) Validate() error {
	return dara.Validate(s)
}

type ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes struct {
	// The ID of the identity provider (IdP) for the user group. If the value of UserGroupType is **department**, this parameter is returned.
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// The logical operator for the user group. Valid values:
	//
	// 	- **Equal**
	//
	// 	- **Unequal**
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// The type of the user group, which is the key of the attribute. Valid values:
	//
	// 	- **username**
	//
	// 	- **department**
	//
	// 	- **email**
	//
	// 	- **telephone**
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// The value of the attribute.
	//
	// 	- If the value of UserGroupType is **username**, the value of this parameter is a username. The value must be 1 to 128 characters in length and can contain letters, digits, hyphens (-), underscores (_), and periods (.).
	//
	// 	- If the value of UserGroupType is **department**, the value of this parameter is a department. Examples: OU=Department 1, OU=SASE DingTalk.
	//
	// 	- If the value of UserGroupType is **email**, the value of this parameter is an email address. Example: username@example.com.
	//
	// 	- If the value of UserGroupType is **telephone**, the value of this parameter is a mobile phone number. Example: 13900001234.
	//
	// example:
	//
	// OU=Department 1, OU=SASE DingTalk
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
