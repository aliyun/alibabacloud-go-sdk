// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProhibitedTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListProhibitedTagsResponseBody
	GetRequestId() *string
	SetTags(v []*ListProhibitedTagsResponseBodyTags) *ListProhibitedTagsResponseBody
	GetTags() []*ListProhibitedTagsResponseBodyTags
	SetTotalNum(v int64) *ListProhibitedTagsResponseBody
	GetTotalNum() *int64
}

type ListProhibitedTagsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 287434FF-344F-565A-8623-439005BA9287
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of prohibited software tags.
	Tags []*ListProhibitedTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The total number of prohibited software tags.
	//
	// example:
	//
	// 10
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListProhibitedTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProhibitedTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProhibitedTagsResponseBody) GetTags() []*ListProhibitedTagsResponseBodyTags {
	return s.Tags
}

func (s *ListProhibitedTagsResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListProhibitedTagsResponseBody) SetRequestId(v string) *ListProhibitedTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProhibitedTagsResponseBody) SetTags(v []*ListProhibitedTagsResponseBodyTags) *ListProhibitedTagsResponseBody {
	s.Tags = v
	return s
}

func (s *ListProhibitedTagsResponseBody) SetTotalNum(v int64) *ListProhibitedTagsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListProhibitedTagsResponseBody) Validate() error {
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

type ListProhibitedTagsResponseBodyTags struct {
	// The time when the prohibited software tag was created, in the yyyy-MM-dd HH:mm:ss format. The time is in the UTC+8 time zone.
	//
	// example:
	//
	// 2022-10-10 11:39:34
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the prohibited software tag.
	//
	// example:
	//
	// created
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the tag is a system built-in device tag. Valid values:
	//
	// - **true**: A system built-in device tag.
	//
	// - **false**: A user-defined device tag.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The name of the prohibited software tag.
	//
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The collection of software prohibition policy IDs that reference the tag.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The collection of prohibited software IDs included in the tag.
	SoftwareIds []*string `json:"SoftwareIds,omitempty" xml:"SoftwareIds,omitempty" type:"Repeated"`
	// The ID of the prohibited software tag.
	//
	// example:
	//
	// tag-7b2c9e4a1d8f****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListProhibitedTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s ListProhibitedTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListProhibitedTagsResponseBodyTags) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProhibitedTagsResponseBodyTags) GetDescription() *string {
	return s.Description
}

func (s *ListProhibitedTagsResponseBodyTags) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListProhibitedTagsResponseBodyTags) GetName() *string {
	return s.Name
}

func (s *ListProhibitedTagsResponseBodyTags) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListProhibitedTagsResponseBodyTags) GetSoftwareIds() []*string {
	return s.SoftwareIds
}

func (s *ListProhibitedTagsResponseBodyTags) GetTagId() *string {
	return s.TagId
}

func (s *ListProhibitedTagsResponseBodyTags) SetCreateTime(v string) *ListProhibitedTagsResponseBodyTags {
	s.CreateTime = &v
	return s
}

func (s *ListProhibitedTagsResponseBodyTags) SetDescription(v string) *ListProhibitedTagsResponseBodyTags {
	s.Description = &v
	return s
}

func (s *ListProhibitedTagsResponseBodyTags) SetIsDefault(v bool) *ListProhibitedTagsResponseBodyTags {
	s.IsDefault = &v
	return s
}

func (s *ListProhibitedTagsResponseBodyTags) SetName(v string) *ListProhibitedTagsResponseBodyTags {
	s.Name = &v
	return s
}

func (s *ListProhibitedTagsResponseBodyTags) SetPolicyIds(v []*string) *ListProhibitedTagsResponseBodyTags {
	s.PolicyIds = v
	return s
}

func (s *ListProhibitedTagsResponseBodyTags) SetSoftwareIds(v []*string) *ListProhibitedTagsResponseBodyTags {
	s.SoftwareIds = v
	return s
}

func (s *ListProhibitedTagsResponseBodyTags) SetTagId(v string) *ListProhibitedTagsResponseBodyTags {
	s.TagId = &v
	return s
}

func (s *ListProhibitedTagsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
