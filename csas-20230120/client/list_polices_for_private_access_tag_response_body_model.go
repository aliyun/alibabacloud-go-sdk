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
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*ListPolicesForPrivateAccessTagResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	Polices []*ListPolicesForPrivateAccessTagResponseBodyTagsPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
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
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// 内网访问策略创建时间。
	//
	// example:
	//
	// 2023-02-21 14:10:16
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 自定义用户组属性集合。多个自定义用户组属性之间是或的关系，按照合集生效。
	CustomUserAttributes []*ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description          *string                                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// example:
	//
	// pa-policy-867ef4007c8a****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// 用户组的身份源ID。当自定义用户组类型为**department**时，存在该值。
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// 用户组的关系。取值：
	//
	// - **Equal**：等于。
	//
	// - **Unequal**：不等于。
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// 用户组的类型。取值：
	//
	// - **username**：用户名。
	//
	// - **department**：部门。
	//
	// - **email**：邮箱。
	//
	// - **telephone**：手机。
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// 用户组属性的值。
	//
	// - 当用户组类型为**username**时，表示用户名的值。长度为1~128个字符，支持中文和大小写英文字母，可包含数字、半角句号（.）、下划线（_）和短划线（-）。
	//
	// - 当用户组类型为**department**时，表示部门的值。如：OU=部门1,OU=SASE钉钉。
	//
	// - 当用户组类型为**email**时，表示邮箱的值。如：username@example.com。
	//
	// - 当用户组类型为**telephone**时，表示手机的值。如：13900001234。
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
