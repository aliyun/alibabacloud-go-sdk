// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeLiveTagResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveTagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *DescribeLiveTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *DescribeLiveTagResourcesRequest
	GetResourceType() *string
	SetTag(v []*DescribeLiveTagResourcesRequestTag) *DescribeLiveTagResourcesRequest
	GetTag() []*DescribeLiveTagResourcesRequestTag
}

type DescribeLiveTagResourcesRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The N resources. The resources are domain names in this operation. Valid values of N: **1 to 50**.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resources. Set the value to **DOMAIN**.
	//
	// This parameter is required.
	//
	// example:
	//
	// DOMAIN
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The N tags.
	Tag []*DescribeLiveTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeLiveTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveTagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeLiveTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeLiveTagResourcesRequest) GetTag() []*DescribeLiveTagResourcesRequestTag {
	return s.Tag
}

func (s *DescribeLiveTagResourcesRequest) SetOwnerId(v int64) *DescribeLiveTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveTagResourcesRequest) SetRegionId(v string) *DescribeLiveTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveTagResourcesRequest) SetResourceId(v []*string) *DescribeLiveTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeLiveTagResourcesRequest) SetResourceType(v string) *DescribeLiveTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeLiveTagResourcesRequest) SetTag(v []*DescribeLiveTagResourcesRequestTag) *DescribeLiveTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *DescribeLiveTagResourcesRequest) Validate() error {
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

type DescribeLiveTagResourcesRequestTag struct {
	// The key of the tag. Valid values of N: **1 to 20**.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. Valid values of N: **1 to 20**.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveTagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeLiveTagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeLiveTagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeLiveTagResourcesRequestTag) SetKey(v string) *DescribeLiveTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeLiveTagResourcesRequestTag) SetValue(v string) *DescribeLiveTagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeLiveTagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
