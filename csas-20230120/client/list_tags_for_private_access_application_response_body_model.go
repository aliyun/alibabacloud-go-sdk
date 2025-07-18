// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsForPrivateAccessApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListTagsForPrivateAccessApplicationResponseBodyApplications) *ListTagsForPrivateAccessApplicationResponseBody
	GetApplications() []*ListTagsForPrivateAccessApplicationResponseBodyApplications
	SetRequestId(v string) *ListTagsForPrivateAccessApplicationResponseBody
	GetRequestId() *string
}

type ListTagsForPrivateAccessApplicationResponseBody struct {
	Applications []*ListTagsForPrivateAccessApplicationResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// 7241F45B-E8D3-5BA3-8172-8A58AC2AB0FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagsForPrivateAccessApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagsForPrivateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessApplicationResponseBody) GetApplications() []*ListTagsForPrivateAccessApplicationResponseBodyApplications {
	return s.Applications
}

func (s *ListTagsForPrivateAccessApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagsForPrivateAccessApplicationResponseBody) SetApplications(v []*ListTagsForPrivateAccessApplicationResponseBodyApplications) *ListTagsForPrivateAccessApplicationResponseBody {
	s.Applications = v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBody) SetRequestId(v string) *ListTagsForPrivateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTagsForPrivateAccessApplicationResponseBodyApplications struct {
	// example:
	//
	// pa-application-7a4445897856****
	ApplicationId *string                                                            `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Tags          []*ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagsForPrivateAccessApplicationResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListTagsForPrivateAccessApplicationResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplications) GetTags() []*ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags {
	return s.Tags
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplications) SetApplicationId(v string) *ListTagsForPrivateAccessApplicationResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplications) SetTags(v []*ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) *ListTagsForPrivateAccessApplicationResponseBodyApplications {
	s.Tags = v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}

type ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags struct {
	// example:
	//
	// 2022-07-01 16:05:26
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// tag-c0cb77857a99****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// Default
	TagType *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
}

func (s ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) String() string {
	return dara.Prettify(s)
}

func (s ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) GetDescription() *string {
	return s.Description
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) GetName() *string {
	return s.Name
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) GetTagId() *string {
	return s.TagId
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) GetTagType() *string {
	return s.TagType
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) SetCreateTime(v string) *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags {
	s.CreateTime = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) SetDescription(v string) *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags {
	s.Description = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) SetName(v string) *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags {
	s.Name = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) SetTagId(v string) *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags {
	s.TagId = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) SetTagType(v string) *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags {
	s.TagType = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) Validate() error {
	return dara.Validate(s)
}
