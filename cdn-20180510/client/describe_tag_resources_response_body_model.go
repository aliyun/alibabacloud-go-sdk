// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTagResourcesResponseBody
	GetRequestId() *string
	SetTagResources(v []*DescribeTagResourcesResponseBodyTagResources) *DescribeTagResourcesResponseBody
	GetTagResources() []*DescribeTagResourcesResponseBodyTagResources
}

type DescribeTagResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 34AB41F1-04A5-496F-8C8D-634BDBE6A9FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags that are attached to the specified resource.
	TagResources []*DescribeTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s DescribeTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagResourcesResponseBody) GetTagResources() []*DescribeTagResourcesResponseBodyTagResources {
	return s.TagResources
}

func (s *DescribeTagResourcesResponseBody) SetRequestId(v string) *DescribeTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagResourcesResponseBody) SetTagResources(v []*DescribeTagResourcesResponseBodyTagResources) *DescribeTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *DescribeTagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTagResourcesResponseBodyTagResources struct {
	// The ID of the resource.
	//
	// example:
	//
	// example.com
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The key-value pair of the tag.
	Tag []*DescribeTagResourcesResponseBodyTagResourcesTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponseBodyTagResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeTagResourcesResponseBodyTagResources) GetTag() []*DescribeTagResourcesResponseBodyTagResourcesTag {
	return s.Tag
}

func (s *DescribeTagResourcesResponseBodyTagResources) SetResourceId(v string) *DescribeTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResources) SetTag(v []*DescribeTagResourcesResponseBodyTagResourcesTag) *DescribeTagResourcesResponseBodyTagResources {
	s.Tag = v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
}

type DescribeTagResourcesResponseBodyTagResourcesTag struct {
	// The key of the tag.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTagResourcesResponseBodyTagResourcesTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResourcesResponseBodyTagResourcesTag) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTag) GetKey() *string {
	return s.Key
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTag) GetValue() *string {
	return s.Value
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTag) SetKey(v string) *DescribeTagResourcesResponseBodyTagResourcesTag {
	s.Key = &v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTag) SetValue(v string) *DescribeTagResourcesResponseBodyTagResourcesTag {
	s.Value = &v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTag) Validate() error {
	return dara.Validate(s)
}
