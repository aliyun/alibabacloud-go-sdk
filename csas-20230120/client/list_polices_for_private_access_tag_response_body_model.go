// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicesForPrivateAccessTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPolicesForPrivateAccessTagResponseBody
	GetRequestId() *string
	SetTags(v []*ListPolicesForPrivateAccessTagResponseBodyTags) *ListPolicesForPrivateAccessTagResponseBody
	GetTags() []*ListPolicesForPrivateAccessTagResponseBodyTags
}

type ListPolicesForPrivateAccessTagResponseBody struct {
	// The ID of this request.
	//
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of private network access tags.
	Tags []*ListPolicesForPrivateAccessTagResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListPolicesForPrivateAccessTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForPrivateAccessTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolicesForPrivateAccessTagResponseBody) GetTags() []*ListPolicesForPrivateAccessTagResponseBodyTags {
	return s.Tags
}

func (s *ListPolicesForPrivateAccessTagResponseBody) SetRequestId(v string) *ListPolicesForPrivateAccessTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBody) SetTags(v []*ListPolicesForPrivateAccessTagResponseBodyTags) *ListPolicesForPrivateAccessTagResponseBody {
	s.Tags = v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBody) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPolicesForPrivateAccessTagResponseBodyTags struct {
	// Collection of private network access policies.
	Polices []*ListPolicesForPrivateAccessTagResponseBodyTagsPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// Private network access tag ID.
	//
	// example:
	//
	// tag-b927baf3e592****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListPolicesForPrivateAccessTagResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForPrivateAccessTagResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTags) GetPolices() []*ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	return s.Polices
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTags) GetTagId() *string {
	return s.TagId
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTags) SetPolices(v []*ListPolicesForPrivateAccessTagResponseBodyTagsPolices) *ListPolicesForPrivateAccessTagResponseBodyTags {
	s.Polices = v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTags) SetTagId(v string) *ListPolicesForPrivateAccessTagResponseBodyTags {
	s.TagId = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTags) Validate() error {
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

type ListPolicesForPrivateAccessTagResponseBodyTagsPolices struct {
	// The application type of the private network access policy. Values:
	//
	// - **Application**: Application.
	//
	// - **Tag**: Tag.
	//
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// Creation time of the private network access policy.
	//
	// example:
	//
	// 2023-02-21 14:10:16
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// A collection of custom user group attributes. Multiple custom user group attributes have an OR relationship and take effect as a union.
	CustomUserAttributes []*ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	// Description of the private network access policy.
	//
	// example:
	//
	// 这是一条内网访问策略
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Private network access policy name.
	//
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The action that the private access policy performs. Valid values:
	//
	// - **Block**: Blocks access.
	//
	// - **Allow**: Allows access.
	//
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// Private network access policy ID.
	//
	// example:
	//
	// pa-policy-867ef4007c8a****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The private network access policy priority. The number 1 indicates the highest priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of the private network access policy. Values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user group type of the private network access policy. Values:
	//
	// - **Normal**: Normal user group.
	//
	// - **Custom**: Custom user group.
	//
	// example:
	//
	// Normal
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
}

func (s ListPolicesForPrivateAccessTagResponseBodyTagsPolices) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForPrivateAccessTagResponseBodyTagsPolices) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) GetCustomUserAttributes() []*ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes {
	return s.CustomUserAttributes
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) GetDescription() *string {
	return s.Description
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) GetName() *string {
	return s.Name
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) GetPolicyAction() *string {
	return s.PolicyAction
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) GetPriority() *int32 {
	return s.Priority
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) GetStatus() *string {
	return s.Status
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetApplicationType(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.ApplicationType = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetCreateTime(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.CreateTime = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetCustomUserAttributes(v []*ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.CustomUserAttributes = v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetDescription(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.Description = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetName(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.Name = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetPolicyAction(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.PolicyAction = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetPolicyId(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.PolicyId = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetPriority(v int32) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.Priority = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetStatus(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.Status = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetUserGroupType(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.UserGroupType = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) Validate() error {
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

type ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes struct {
	// The identity provider ID of the user group. This value exists if the custom user group type is **department**.
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// The relationship of the user group. Values:
	//
	// - **Equal**: Equal.
	//
	// - **Unequal**: Unequal.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// The type of user group. Values:
	//
	// - **username**: Username.
	//
	// - **department**: Department.
	//
	// - **email**: Mailbox.
	//
	// - **telephone**: Mobile phone.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// The value of the user group attribute.
	//
	// - If the user group type is **username**, this indicates the username\\"s value. The length is 1 to 128 characters. It supports Chinese characters, uppercase and lowercase English letters, and can include numbers, periods (.), underscores (_), and hyphens (-).
	//
	// - If the user group type is **department**, this indicates the department\\"s value. For example: OU=Department 1,OU=SASE DingTalk.
	//
	// - If the user group type is **email**, this indicates the mailbox\\"s value. For example: username\\@example.com.
	//
	// - If the user group type is **telephone**, this indicates the mobile phone\\"s value. For example: 13900001234.
	//
	// example:
	//
	// OU=部门1,OU=SASE钉钉
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) String() string {
	return dara.Prettify(s)
}

func (s ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) GetIdpId() *int32 {
	return s.IdpId
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) GetRelation() *string {
	return s.Relation
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) GetUserGroupType() *string {
	return s.UserGroupType
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) GetValue() *string {
	return s.Value
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) SetIdpId(v int32) *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) SetRelation(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) SetUserGroupType(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) SetValue(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes {
	s.Value = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) Validate() error {
	return dara.Validate(s)
}
