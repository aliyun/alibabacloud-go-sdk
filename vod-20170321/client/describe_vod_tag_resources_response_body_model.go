// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVodTagResourcesResponseBody
	GetRequestId() *string
	SetTagResources(v []*DescribeVodTagResourcesResponseBodyTagResources) *DescribeVodTagResourcesResponseBody
	GetTagResources() []*DescribeVodTagResourcesResponseBodyTagResources
}

type DescribeVodTagResourcesResponseBody struct {
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*DescribeVodTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s DescribeVodTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodTagResourcesResponseBody) GetTagResources() []*DescribeVodTagResourcesResponseBodyTagResources {
	return s.TagResources
}

func (s *DescribeVodTagResourcesResponseBody) SetRequestId(v string) *DescribeVodTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodTagResourcesResponseBody) SetTagResources(v []*DescribeVodTagResourcesResponseBodyTagResources) *DescribeVodTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *DescribeVodTagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodTagResourcesResponseBodyTagResources struct {
	ResourceId *string                                               `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Tag        []*DescribeVodTagResourcesResponseBodyTagResourcesTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVodTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *DescribeVodTagResourcesResponseBodyTagResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeVodTagResourcesResponseBodyTagResources) GetTag() []*DescribeVodTagResourcesResponseBodyTagResourcesTag {
	return s.Tag
}

func (s *DescribeVodTagResourcesResponseBodyTagResources) SetResourceId(v string) *DescribeVodTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeVodTagResourcesResponseBodyTagResources) SetTag(v []*DescribeVodTagResourcesResponseBodyTagResourcesTag) *DescribeVodTagResourcesResponseBodyTagResources {
	s.Tag = v
	return s
}

func (s *DescribeVodTagResourcesResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
}

type DescribeVodTagResourcesResponseBodyTagResourcesTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodTagResourcesResponseBodyTagResourcesTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTagResourcesResponseBodyTagResourcesTag) GoString() string {
	return s.String()
}

func (s *DescribeVodTagResourcesResponseBodyTagResourcesTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVodTagResourcesResponseBodyTagResourcesTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVodTagResourcesResponseBodyTagResourcesTag) SetKey(v string) *DescribeVodTagResourcesResponseBodyTagResourcesTag {
	s.Key = &v
	return s
}

func (s *DescribeVodTagResourcesResponseBodyTagResourcesTag) SetValue(v string) *DescribeVodTagResourcesResponseBodyTagResourcesTag {
	s.Value = &v
	return s
}

func (s *DescribeVodTagResourcesResponseBodyTagResourcesTag) Validate() error {
	return dara.Validate(s)
}
