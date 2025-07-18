// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListPrivateAccessTagsResponseBody
	GetRequestId() *string
	SetTags(v []*ListPrivateAccessTagsResponseBodyTags) *ListPrivateAccessTagsResponseBody
	GetTags() []*ListPrivateAccessTagsResponseBodyTags
	SetTotalNum(v int32) *ListPrivateAccessTagsResponseBody
	GetTotalNum() *int32
}

type ListPrivateAccessTagsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54C1D236-CDB9-586C-B44D-AFDCEA195545
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The internal access tags.
	Tags []*ListPrivateAccessTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The total number of internal access tags.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListPrivateAccessTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivateAccessTagsResponseBody) GetTags() []*ListPrivateAccessTagsResponseBodyTags {
	return s.Tags
}

func (s *ListPrivateAccessTagsResponseBody) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ListPrivateAccessTagsResponseBody) SetRequestId(v string) *ListPrivateAccessTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateAccessTagsResponseBody) SetTags(v []*ListPrivateAccessTagsResponseBodyTags) *ListPrivateAccessTagsResponseBody {
	s.Tags = v
	return s
}

func (s *ListPrivateAccessTagsResponseBody) SetTotalNum(v int32) *ListPrivateAccessTagsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListPrivateAccessTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPrivateAccessTagsResponseBodyTags struct {
	// The IDs of the internal access applications.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The time when the internal access tag was created.
	//
	// example:
	//
	// 2022-10-10 11:39:34
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the internal access tag.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the internal access tag.
	//
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of the internal access policies.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The ID of the internal access tag.
	//
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The type of the internal access tag. Valid values:
	//
	// 	- **Default**
	//
	// 	- **Custom**
	//
	// example:
	//
	// Default
	TagType *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
}

func (s ListPrivateAccessTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsResponseBodyTags) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListPrivateAccessTagsResponseBodyTags) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPrivateAccessTagsResponseBodyTags) GetDescription() *string {
	return s.Description
}

func (s *ListPrivateAccessTagsResponseBodyTags) GetName() *string {
	return s.Name
}

func (s *ListPrivateAccessTagsResponseBodyTags) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *ListPrivateAccessTagsResponseBodyTags) GetTagId() *string {
	return s.TagId
}

func (s *ListPrivateAccessTagsResponseBodyTags) GetTagType() *string {
	return s.TagType
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetApplicationIds(v []*string) *ListPrivateAccessTagsResponseBodyTags {
	s.ApplicationIds = v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetCreateTime(v string) *ListPrivateAccessTagsResponseBodyTags {
	s.CreateTime = &v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetDescription(v string) *ListPrivateAccessTagsResponseBodyTags {
	s.Description = &v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetName(v string) *ListPrivateAccessTagsResponseBodyTags {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetPolicyIds(v []*string) *ListPrivateAccessTagsResponseBodyTags {
	s.PolicyIds = v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetTagId(v string) *ListPrivateAccessTagsResponseBodyTags {
	s.TagId = &v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetTagType(v string) *ListPrivateAccessTagsResponseBodyTags {
	s.TagType = &v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
