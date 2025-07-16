// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v []*string) *DescribeTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *DescribeTagResourcesRequest
	GetResourceType() *string
	SetTag(v []*DescribeTagResourcesRequestTag) *DescribeTagResourcesRequest
	GetTag() []*DescribeTagResourcesRequestTag
}

type DescribeTagResourcesRequest struct {
	// The IDs of the resources. You can specify up to 50 IDs in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Set the value to **DOMAIN**.
	//
	// This parameter is required.
	//
	// example:
	//
	// DOMAIN
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. You can specify up to 20 tags in each request.
	Tag []*DescribeTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagResourcesRequest) GetTag() []*DescribeTagResourcesRequestTag {
	return s.Tag
}

func (s *DescribeTagResourcesRequest) SetResourceId(v []*string) *DescribeTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeTagResourcesRequest) SetResourceType(v string) *DescribeTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagResourcesRequest) SetTag(v []*DescribeTagResourcesRequestTag) *DescribeTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *DescribeTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeTagResourcesRequestTag struct {
	// The key of the tag. Valid values of N: **1*	- to **20**.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. Valid values of N: **1*	- to **20**.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeTagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeTagResourcesRequestTag) SetKey(v string) *DescribeTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeTagResourcesRequestTag) SetValue(v string) *DescribeTagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeTagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
