// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListControlPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetControlPolicies(v *ListControlPoliciesResponseBodyControlPolicies) *ListControlPoliciesResponseBody
	GetControlPolicies() *ListControlPoliciesResponseBodyControlPolicies
	SetPageNumber(v int32) *ListControlPoliciesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListControlPoliciesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListControlPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListControlPoliciesResponseBody
	GetTotalCount() *int32
}

type ListControlPoliciesResponseBody struct {
	// The access control policies.
	ControlPolicies *ListControlPoliciesResponseBodyControlPolicies `json:"ControlPolicies,omitempty" xml:"ControlPolicies,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9424A34C-3471-45AD-B6AB-924BBDFE42F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of access control policies.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListControlPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListControlPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseBody) GetControlPolicies() *ListControlPoliciesResponseBodyControlPolicies {
	return s.ControlPolicies
}

func (s *ListControlPoliciesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListControlPoliciesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListControlPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListControlPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListControlPoliciesResponseBody) SetControlPolicies(v *ListControlPoliciesResponseBodyControlPolicies) *ListControlPoliciesResponseBody {
	s.ControlPolicies = v
	return s
}

func (s *ListControlPoliciesResponseBody) SetPageNumber(v int32) *ListControlPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListControlPoliciesResponseBody) SetPageSize(v int32) *ListControlPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListControlPoliciesResponseBody) SetRequestId(v string) *ListControlPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListControlPoliciesResponseBody) SetTotalCount(v int32) *ListControlPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListControlPoliciesResponseBody) Validate() error {
	if s.ControlPolicies != nil {
		if err := s.ControlPolicies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListControlPoliciesResponseBodyControlPolicies struct {
	ControlPolicy []*ListControlPoliciesResponseBodyControlPoliciesControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Repeated"`
}

func (s ListControlPoliciesResponseBodyControlPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListControlPoliciesResponseBodyControlPolicies) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseBodyControlPolicies) GetControlPolicy() []*ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	return s.ControlPolicy
}

func (s *ListControlPoliciesResponseBodyControlPolicies) SetControlPolicy(v []*ListControlPoliciesResponseBodyControlPoliciesControlPolicy) *ListControlPoliciesResponseBodyControlPolicies {
	s.ControlPolicy = v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPolicies) Validate() error {
	if s.ControlPolicy != nil {
		for _, item := range s.ControlPolicy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListControlPoliciesResponseBodyControlPoliciesControlPolicy struct {
	// The number of times that the access control policy is referenced.
	//
	// example:
	//
	// 44
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the access control policy was created.
	//
	// example:
	//
	// 2020-08-05T06:32:24Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the access control policy.
	//
	// example:
	//
	// System access control policy available for all operations on the cloud
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy. Valid values:
	//
	// 	- All: The access control policy is in effect for Alibaba Cloud accounts, RAM users, and RAM roles.
	//
	// 	- RAM: The access control policy is in effect only for RAM users and RAM roles.
	//
	// example:
	//
	// All
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The ID of the access control policy.
	//
	// example:
	//
	// cp-FullAliyunAccess
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	//
	// example:
	//
	// FullAliyunAccess
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// 	- System: system access control policy
	//
	// 	- Custom: custom access control policy
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The tags.
	Tags *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The time when the access control policy was updated.
	//
	// example:
	//
	// 2020-08-05T06:32:24Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListControlPoliciesResponseBodyControlPoliciesControlPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListControlPoliciesResponseBodyControlPoliciesControlPolicy) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) GetAttachmentCount() *string {
	return s.AttachmentCount
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) GetDescription() *string {
	return s.Description
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) GetEffectScope() *string {
	return s.EffectScope
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) GetTags() *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTags {
	return s.Tags
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetAttachmentCount(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetCreateDate(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetDescription(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.Description = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetEffectScope(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.EffectScope = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetPolicyId(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetPolicyName(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetPolicyType(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetTags(v *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTags) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.Tags = v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) SetUpdateDate(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicy {
	s.UpdateDate = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicy) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListControlPoliciesResponseBodyControlPoliciesControlPolicyTags struct {
	Tag []*ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListControlPoliciesResponseBodyControlPoliciesControlPolicyTags) String() string {
	return dara.Prettify(s)
}

func (s ListControlPoliciesResponseBodyControlPoliciesControlPolicyTags) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTags) GetTag() []*ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag {
	return s.Tag
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTags) SetTag(v []*ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag) *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTags {
	s.Tag = v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag) GoString() string {
	return s.String()
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag) GetKey() *string {
	return s.Key
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag) GetValue() *string {
	return s.Value
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag) SetKey(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag {
	s.Key = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag) SetValue(v string) *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag {
	s.Value = &v
	return s
}

func (s *ListControlPoliciesResponseBodyControlPoliciesControlPolicyTagsTag) Validate() error {
	return dara.Validate(s)
}
