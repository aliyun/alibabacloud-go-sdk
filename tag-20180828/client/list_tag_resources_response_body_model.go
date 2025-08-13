// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagResourcesResponseBody
	GetRequestId() *string
	SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody
	GetTagResources() []*ListTagResourcesResponseBodyTagResources
}

type ListTagResourcesResponseBody struct {
	// Indicates whether the `next query` is required.
	//
	// 	- If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the `next query` is not required.
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the `token` used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 014738E0-3C7F-47D8-8FB9-469500C6F387
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the tags that are added to the resources.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResourcesResponseBody) GetTagResources() []*ListTagResourcesResponseBodyTagResources {
	return s.TagResources
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesResponseBodyTagResources struct {
	// The ARN of the resource.
	//
	// example:
	//
	// arn:acs:ecs:cn-hangzhou:123456789****:instance/i-bp15hr53jws84akg****
	ResourceARN *string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	// The information of the tags.
	Tags []*ListTagResourcesResponseBodyTagResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) GetResourceARN() *string {
	return s.ResourceARN
}

func (s *ListTagResourcesResponseBodyTagResources) GetTags() []*ListTagResourcesResponseBodyTagResourcesTags {
	return s.Tags
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceARN(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceARN = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTags(v []*ListTagResourcesResponseBodyTagResourcesTags) *ListTagResourcesResponseBodyTagResources {
	s.Tags = v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesResponseBodyTagResourcesTags struct {
	// The type of the tag. Valid values:
	//
	// 	- Custom
	//
	// 	- System
	//
	// example:
	//
	// Custom
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTags) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTags) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) GetCategory() *string {
	return s.Category
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) GetKey() *string {
	return s.Key
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) GetValue() *string {
	return s.Value
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) SetCategory(v string) *ListTagResourcesResponseBodyTagResourcesTags {
	s.Category = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) SetKey(v string) *ListTagResourcesResponseBodyTagResourcesTags {
	s.Key = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) SetValue(v string) *ListTagResourcesResponseBodyTagResourcesTags {
	s.Value = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) Validate() error {
	return dara.Validate(s)
}
