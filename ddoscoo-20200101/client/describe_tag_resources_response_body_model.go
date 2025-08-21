// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeTagResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeTagResourcesResponseBody
	GetRequestId() *string
	SetTagResources(v *DescribeTagResourcesResponseBodyTagResources) *DescribeTagResourcesResponseBody
	GetTagResources() *DescribeTagResourcesResponseBodyTagResources
}

type DescribeTagResourcesResponseBody struct {
	// The query token that is returned in this call.
	//
	// example:
	//
	// RGuYpqDdKhzXb8C3.D1BwQgc1tMBsoxdGiEKHHUUCf****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 36E698F7-48A4-48D0-9554-0BB4BAAB99B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags that are added to the Anti-DDoS Proxy (Chinese Mainland) instance.
	TagResources *DescribeTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s DescribeTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagResourcesResponseBody) GetTagResources() *DescribeTagResourcesResponseBodyTagResources {
	return s.TagResources
}

func (s *DescribeTagResourcesResponseBody) SetNextToken(v string) *DescribeTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTagResourcesResponseBody) SetRequestId(v string) *DescribeTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagResourcesResponseBody) SetTagResources(v *DescribeTagResourcesResponseBodyTagResources) *DescribeTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *DescribeTagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTagResourcesResponseBodyTagResources struct {
	TagResource []*DescribeTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s DescribeTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponseBodyTagResources) GetTagResource() []*DescribeTagResourcesResponseBodyTagResourcesTagResource {
	return s.TagResource
}

func (s *DescribeTagResourcesResponseBodyTagResources) SetTagResource(v []*DescribeTagResourcesResponseBodyTagResourcesTagResource) *DescribeTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
}

type DescribeTagResourcesResponseBodyTagResourcesTagResource struct {
	// The ID of the Anti-DDoS Proxy (Chinese Mainland) instance.
	//
	// example:
	//
	// ddoscoo-cn-zz121ogz****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. The value is fixed as **INSTANCE**, which indicates an Anti-DDoS Proxy instance.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag that is added to the Anti-DDoS Proxy (Chinese Mainland) instance.
	//
	// example:
	//
	// testvalue
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the Anti-DDoS Proxy (Chinese Mainland) instance.
	//
	// example:
	//
	// testkey
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *DescribeTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *DescribeTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *DescribeTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *DescribeTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *DescribeTagResourcesResponseBodyTagResourcesTagResource) Validate() error {
	return dara.Validate(s)
}
