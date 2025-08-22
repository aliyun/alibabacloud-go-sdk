// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDcdnTagResourcesResponseBody
	GetRequestId() *string
	SetTagResources(v []*DescribeDcdnTagResourcesResponseBodyTagResources) *DescribeDcdnTagResourcesResponseBody
	GetTagResources() []*DescribeDcdnTagResourcesResponseBodyTagResources
}

type DescribeDcdnTagResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 34AB41F1-04A5-496F-8C8D-634BDBE6A9FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag of the instance.
	TagResources []*DescribeDcdnTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s DescribeDcdnTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnTagResourcesResponseBody) GetTagResources() []*DescribeDcdnTagResourcesResponseBodyTagResources {
	return s.TagResources
}

func (s *DescribeDcdnTagResourcesResponseBody) SetRequestId(v string) *DescribeDcdnTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnTagResourcesResponseBody) SetTagResources(v []*DescribeDcdnTagResourcesResponseBodyTagResources) *DescribeDcdnTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *DescribeDcdnTagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnTagResourcesResponseBodyTagResources struct {
	// The ID of the resource.
	//
	// example:
	//
	// example.com
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The tags of the snapshot.
	Tag []*DescribeDcdnTagResourcesResponseBodyTagResourcesTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDcdnTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResources) GetTag() []*DescribeDcdnTagResourcesResponseBodyTagResourcesTag {
	return s.Tag
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResources) SetResourceId(v string) *DescribeDcdnTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResources) SetTag(v []*DescribeDcdnTagResourcesResponseBodyTagResourcesTag) *DescribeDcdnTagResourcesResponseBodyTagResources {
	s.Tag = v
	return s
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnTagResourcesResponseBodyTagResourcesTag struct {
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

func (s DescribeDcdnTagResourcesResponseBodyTagResourcesTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnTagResourcesResponseBodyTagResourcesTag) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResourcesTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResourcesTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResourcesTag) SetKey(v string) *DescribeDcdnTagResourcesResponseBodyTagResourcesTag {
	s.Key = &v
	return s
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResourcesTag) SetValue(v string) *DescribeDcdnTagResourcesResponseBodyTagResourcesTag {
	s.Value = &v
	return s
}

func (s *DescribeDcdnTagResourcesResponseBodyTagResourcesTag) Validate() error {
	return dara.Validate(s)
}
