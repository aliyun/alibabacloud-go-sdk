// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveTagResourcesResponseBody
	GetRequestId() *string
	SetTagResources(v []*DescribeLiveTagResourcesResponseBodyTagResources) *DescribeLiveTagResourcesResponseBody
	GetTagResources() []*DescribeLiveTagResourcesResponseBodyTagResources
}

type DescribeLiveTagResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 34AB41F1-04A5-496F-8C8D-634BDBE6A9FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags of the resource.
	TagResources []*DescribeLiveTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s DescribeLiveTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveTagResourcesResponseBody) GetTagResources() []*DescribeLiveTagResourcesResponseBodyTagResources {
	return s.TagResources
}

func (s *DescribeLiveTagResourcesResponseBody) SetRequestId(v string) *DescribeLiveTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveTagResourcesResponseBody) SetTagResources(v []*DescribeLiveTagResourcesResponseBodyTagResources) *DescribeLiveTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *DescribeLiveTagResourcesResponseBody) Validate() error {
	if s.TagResources != nil {
		for _, item := range s.TagResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveTagResourcesResponseBodyTagResources struct {
	// The resource.
	//
	// example:
	//
	// example.com
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The key-value pair of the tag.
	Tag []*DescribeLiveTagResourcesResponseBodyTagResourcesTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLiveTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *DescribeLiveTagResourcesResponseBodyTagResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeLiveTagResourcesResponseBodyTagResources) GetTag() []*DescribeLiveTagResourcesResponseBodyTagResourcesTag {
	return s.Tag
}

func (s *DescribeLiveTagResourcesResponseBodyTagResources) SetResourceId(v string) *DescribeLiveTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeLiveTagResourcesResponseBodyTagResources) SetTag(v []*DescribeLiveTagResourcesResponseBodyTagResourcesTag) *DescribeLiveTagResourcesResponseBodyTagResources {
	s.Tag = v
	return s
}

func (s *DescribeLiveTagResourcesResponseBodyTagResources) Validate() error {
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

type DescribeLiveTagResourcesResponseBodyTagResourcesTag struct {
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

func (s DescribeLiveTagResourcesResponseBodyTagResourcesTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTagResourcesResponseBodyTagResourcesTag) GoString() string {
	return s.String()
}

func (s *DescribeLiveTagResourcesResponseBodyTagResourcesTag) GetKey() *string {
	return s.Key
}

func (s *DescribeLiveTagResourcesResponseBodyTagResourcesTag) GetValue() *string {
	return s.Value
}

func (s *DescribeLiveTagResourcesResponseBodyTagResourcesTag) SetKey(v string) *DescribeLiveTagResourcesResponseBodyTagResourcesTag {
	s.Key = &v
	return s
}

func (s *DescribeLiveTagResourcesResponseBodyTagResourcesTag) SetValue(v string) *DescribeLiveTagResourcesResponseBodyTagResourcesTag {
	s.Value = &v
	return s
}

func (s *DescribeLiveTagResourcesResponseBodyTagResourcesTag) Validate() error {
	return dara.Validate(s)
}
