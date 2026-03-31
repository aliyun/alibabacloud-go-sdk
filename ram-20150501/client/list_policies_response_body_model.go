// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListPoliciesResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *ListPoliciesResponseBody
	GetMarker() *string
	SetPolicies(v *ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody
	GetPolicies() *ListPoliciesResponseBodyPolicies
	SetRequestId(v string) *ListPoliciesResponseBody
	GetRequestId() *string
}

type ListPoliciesResponseBody struct {
	// Indicates whether the response is truncated.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The marker. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set `Marker` to obtain the truncated part.``
	//
	// example:
	//
	// EXAMPLE
	Marker   *string                           `json:"Marker,omitempty" xml:"Marker,omitempty"`
	Policies *ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListPoliciesResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListPoliciesResponseBody) GetPolicies() *ListPoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *ListPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoliciesResponseBody) SetIsTruncated(v bool) *ListPoliciesResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListPoliciesResponseBody) SetMarker(v string) *ListPoliciesResponseBody {
	s.Marker = &v
	return s
}

func (s *ListPoliciesResponseBody) SetPolicies(v *ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesResponseBody) SetRequestId(v string) *ListPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesResponseBody) Validate() error {
	if s.Policies != nil {
		if err := s.Policies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPoliciesResponseBodyPolicies struct {
	Policy []*ListPoliciesResponseBodyPoliciesPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
}

func (s ListPoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPolicies) GetPolicy() []*ListPoliciesResponseBodyPoliciesPolicy {
	return s.Policy
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicy(v []*ListPoliciesResponseBodyPoliciesPolicy) *ListPoliciesResponseBodyPolicies {
	s.Policy = v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) Validate() error {
	if s.Policy != nil {
		for _, item := range s.Policy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPoliciesResponseBodyPoliciesPolicy struct {
	AttachmentCount *int32                                      `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	CreateDate      *string                                     `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DefaultVersion  *string                                     `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Description     *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName      *string                                     `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType      *string                                     `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	Tags            *ListPoliciesResponseBodyPoliciesPolicyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UpdateDate      *string                                     `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListPoliciesResponseBodyPoliciesPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBodyPoliciesPolicy) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetAttachmentCount() *int32 {
	return s.AttachmentCount
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetDescription() *string {
	return s.Description
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetTags() *ListPoliciesResponseBodyPoliciesPolicyTags {
	return s.Tags
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetAttachmentCount(v int32) *ListPoliciesResponseBodyPoliciesPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetCreateDate(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.CreateDate = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetDefaultVersion(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetDescription(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.Description = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetPolicyName(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetPolicyType(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetTags(v *ListPoliciesResponseBodyPoliciesPolicyTags) *ListPoliciesResponseBodyPoliciesPolicy {
	s.Tags = v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) SetUpdateDate(v string) *ListPoliciesResponseBodyPoliciesPolicy {
	s.UpdateDate = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicy) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPoliciesResponseBodyPoliciesPolicyTags struct {
	Tag []*ListPoliciesResponseBodyPoliciesPolicyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListPoliciesResponseBodyPoliciesPolicyTags) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBodyPoliciesPolicyTags) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPoliciesPolicyTags) GetTag() []*ListPoliciesResponseBodyPoliciesPolicyTagsTag {
	return s.Tag
}

func (s *ListPoliciesResponseBodyPoliciesPolicyTags) SetTag(v []*ListPoliciesResponseBodyPoliciesPolicyTagsTag) *ListPoliciesResponseBodyPoliciesPolicyTags {
	s.Tag = v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicyTags) Validate() error {
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

type ListPoliciesResponseBodyPoliciesPolicyTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListPoliciesResponseBodyPoliciesPolicyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBodyPoliciesPolicyTagsTag) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPoliciesPolicyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListPoliciesResponseBodyPoliciesPolicyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListPoliciesResponseBodyPoliciesPolicyTagsTag) SetTagKey(v string) *ListPoliciesResponseBodyPoliciesPolicyTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicyTagsTag) SetTagValue(v string) *ListPoliciesResponseBodyPoliciesPolicyTagsTag {
	s.TagValue = &v
	return s
}

func (s *ListPoliciesResponseBodyPoliciesPolicyTagsTag) Validate() error {
	return dara.Validate(s)
}
