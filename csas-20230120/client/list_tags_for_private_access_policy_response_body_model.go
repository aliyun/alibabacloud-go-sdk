// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsForPrivateAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolices(v []*ListTagsForPrivateAccessPolicyResponseBodyPolices) *ListTagsForPrivateAccessPolicyResponseBody
	GetPolices() []*ListTagsForPrivateAccessPolicyResponseBodyPolices
	SetRequestId(v string) *ListTagsForPrivateAccessPolicyResponseBody
	GetRequestId() *string
}

type ListTagsForPrivateAccessPolicyResponseBody struct {
	Polices []*ListTagsForPrivateAccessPolicyResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// example:
	//
	// 9EE61139-A6A8-5E13-80AF-83435C21B26B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagsForPrivateAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagsForPrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessPolicyResponseBody) GetPolices() []*ListTagsForPrivateAccessPolicyResponseBodyPolices {
	return s.Polices
}

func (s *ListTagsForPrivateAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagsForPrivateAccessPolicyResponseBody) SetPolices(v []*ListTagsForPrivateAccessPolicyResponseBodyPolices) *ListTagsForPrivateAccessPolicyResponseBody {
	s.Polices = v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBody) SetRequestId(v string) *ListTagsForPrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBody) Validate() error {
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

type ListTagsForPrivateAccessPolicyResponseBodyPolices struct {
	// example:
	//
	// pa-policy-1b0d0e8b4bcf****
	PolicyId *string                                                  `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Tags     []*ListTagsForPrivateAccessPolicyResponseBodyPolicesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagsForPrivateAccessPolicyResponseBodyPolices) String() string {
	return dara.Prettify(s)
}

func (s ListTagsForPrivateAccessPolicyResponseBodyPolices) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolices) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolices) GetTags() []*ListTagsForPrivateAccessPolicyResponseBodyPolicesTags {
	return s.Tags
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolices) SetPolicyId(v string) *ListTagsForPrivateAccessPolicyResponseBodyPolices {
	s.PolicyId = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolices) SetTags(v []*ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) *ListTagsForPrivateAccessPolicyResponseBodyPolices {
	s.Tags = v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolices) Validate() error {
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

type ListTagsForPrivateAccessPolicyResponseBodyPolicesTags struct {
	// 内网访问标签创建时间。
	//
	// example:
	//
	// 2023-02-21 14:10:16
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// Default
	TagType *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
}

func (s ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) String() string {
	return dara.Prettify(s)
}

func (s ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) GetDescription() *string {
	return s.Description
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) GetName() *string {
	return s.Name
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) GetTagId() *string {
	return s.TagId
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) GetTagType() *string {
	return s.TagType
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) SetCreateTime(v string) *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags {
	s.CreateTime = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) SetDescription(v string) *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags {
	s.Description = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) SetName(v string) *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags {
	s.Name = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) SetTagId(v string) *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags {
	s.TagId = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) SetTagType(v string) *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags {
	s.TagType = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) Validate() error {
	return dara.Validate(s)
}
