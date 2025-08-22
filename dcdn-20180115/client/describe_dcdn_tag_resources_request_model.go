// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v []*string) *DescribeDcdnTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *DescribeDcdnTagResourcesRequest
	GetResourceType() *string
	SetTag(v []*DescribeDcdnTagResourcesRequestTag) *DescribeDcdnTagResourcesRequest
	GetTag() []*DescribeDcdnTagResourcesRequestTag
}

type DescribeDcdnTagResourcesRequest struct {
	// The list of resource IDs. You can specify a maximum of 50 resource IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Set the value to **DOMAIN**.
	//
	// This parameter is required.
	//
	// example:
	//
	// DOMAIN
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. You can specify a maximum of 20 tag values.
	Tag []*DescribeDcdnTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDcdnTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeDcdnTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDcdnTagResourcesRequest) GetTag() []*DescribeDcdnTagResourcesRequestTag {
	return s.Tag
}

func (s *DescribeDcdnTagResourcesRequest) SetResourceId(v []*string) *DescribeDcdnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeDcdnTagResourcesRequest) SetResourceType(v string) *DescribeDcdnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDcdnTagResourcesRequest) SetTag(v []*DescribeDcdnTagResourcesRequestTag) *DescribeDcdnTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *DescribeDcdnTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnTagResourcesRequestTag struct {
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

func (s DescribeDcdnTagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDcdnTagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnTagResourcesRequestTag) SetKey(v string) *DescribeDcdnTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDcdnTagResourcesRequestTag) SetValue(v string) *DescribeDcdnTagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDcdnTagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
