// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesByTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListResourcesByTagResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResourcesByTagResponseBody
	GetRequestId() *string
	SetResources(v []*ListResourcesByTagResponseBodyResources) *ListResourcesByTagResponseBody
	GetResources() []*ListResourcesByTagResponseBodyResources
}

type ListResourcesByTagResponseBody struct {
	// Indicates whether the `next query` is required.
	//
	// 	- If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the `next query` is not required.
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the `token` used to start the next query.
	//
	// This parameter is required.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7D61FF74-61C2-5768-B01F-05FC97F24F35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resources.
	Resources []*ListResourcesByTagResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListResourcesByTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesByTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourcesByTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourcesByTagResponseBody) GetResources() []*ListResourcesByTagResponseBodyResources {
	return s.Resources
}

func (s *ListResourcesByTagResponseBody) SetNextToken(v string) *ListResourcesByTagResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourcesByTagResponseBody) SetRequestId(v string) *ListResourcesByTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesByTagResponseBody) SetResources(v []*ListResourcesByTagResponseBodyResources) *ListResourcesByTagResponseBody {
	s.Resources = v
	return s
}

func (s *ListResourcesByTagResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListResourcesByTagResponseBodyResources struct {
	// The ID of the resource.
	//
	// example:
	//
	// vpc-wz9pifyuw26esxd05****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The information of the tags.
	//
	// This parameter is returned only if the `IncludeAllTags` parameter is set to `True`.
	Tags []*ListResourcesByTagResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListResourcesByTagResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesByTagResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagResponseBodyResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourcesByTagResponseBodyResources) GetTags() []*ListResourcesByTagResponseBodyResourcesTags {
	return s.Tags
}

func (s *ListResourcesByTagResponseBodyResources) SetResourceId(v string) *ListResourcesByTagResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesByTagResponseBodyResources) SetTags(v []*ListResourcesByTagResponseBodyResourcesTags) *ListResourcesByTagResponseBodyResources {
	s.Tags = v
	return s
}

func (s *ListResourcesByTagResponseBodyResources) Validate() error {
	return dara.Validate(s)
}

type ListResourcesByTagResponseBodyResourcesTags struct {
	// The type of the tag. Valid values:
	//
	// 	- custom
	//
	// 	- system
	//
	// example:
	//
	// custom
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

func (s ListResourcesByTagResponseBodyResourcesTags) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesByTagResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagResponseBodyResourcesTags) GetCategory() *string {
	return s.Category
}

func (s *ListResourcesByTagResponseBodyResourcesTags) GetKey() *string {
	return s.Key
}

func (s *ListResourcesByTagResponseBodyResourcesTags) GetValue() *string {
	return s.Value
}

func (s *ListResourcesByTagResponseBodyResourcesTags) SetCategory(v string) *ListResourcesByTagResponseBodyResourcesTags {
	s.Category = &v
	return s
}

func (s *ListResourcesByTagResponseBodyResourcesTags) SetKey(v string) *ListResourcesByTagResponseBodyResourcesTags {
	s.Key = &v
	return s
}

func (s *ListResourcesByTagResponseBodyResourcesTags) SetValue(v string) *ListResourcesByTagResponseBodyResourcesTags {
	s.Value = &v
	return s
}

func (s *ListResourcesByTagResponseBodyResourcesTags) Validate() error {
	return dara.Validate(s)
}
