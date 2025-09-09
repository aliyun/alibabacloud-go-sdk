// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeDrdsInstancesRequest
	GetDescription() *string
	SetExpired(v bool) *DescribeDrdsInstancesRequest
	GetExpired() *bool
	SetMix(v bool) *DescribeDrdsInstancesRequest
	GetMix() *bool
	SetPageNumber(v int32) *DescribeDrdsInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDrdsInstancesRequest
	GetPageSize() *int32
	SetProductVersion(v string) *DescribeDrdsInstancesRequest
	GetProductVersion() *string
	SetRegionId(v string) *DescribeDrdsInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDrdsInstancesRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeDrdsInstancesRequestTag) *DescribeDrdsInstancesRequest
	GetTag() []*DescribeDrdsInstancesRequestTag
	SetType(v string) *DescribeDrdsInstancesRequest
	GetType() *string
}

type DescribeDrdsInstancesRequest struct {
	// The description of the instances.
	//
	// example:
	//
	// drds_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the instances that you want to query expire.
	//
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Specifies whether hybrid queries are supported.
	//
	// example:
	//
	// FALSE
	Mix *bool `json:"Mix,omitempty" xml:"Mix,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// V1
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instances you want to query belong. The value of this parameter can be NULL.
	//
	// example:
	//
	// NULL
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeDrdsInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of the instances that you want to query. Valid values:
	//
	// 	- **0**: shared instances
	//
	// 	- **1**: dedicated instances
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDrdsInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeDrdsInstancesRequest) GetExpired() *bool {
	return s.Expired
}

func (s *DescribeDrdsInstancesRequest) GetMix() *bool {
	return s.Mix
}

func (s *DescribeDrdsInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDrdsInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDrdsInstancesRequest) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *DescribeDrdsInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrdsInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDrdsInstancesRequest) GetTag() []*DescribeDrdsInstancesRequestTag {
	return s.Tag
}

func (s *DescribeDrdsInstancesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDrdsInstancesRequest) SetDescription(v string) *DescribeDrdsInstancesRequest {
	s.Description = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetExpired(v bool) *DescribeDrdsInstancesRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetMix(v bool) *DescribeDrdsInstancesRequest {
	s.Mix = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetPageNumber(v int32) *DescribeDrdsInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetPageSize(v int32) *DescribeDrdsInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetProductVersion(v string) *DescribeDrdsInstancesRequest {
	s.ProductVersion = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetRegionId(v string) *DescribeDrdsInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetResourceGroupId(v string) *DescribeDrdsInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetTag(v []*DescribeDrdsInstancesRequestTag) *DescribeDrdsInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeDrdsInstancesRequest) SetType(v string) *DescribeDrdsInstancesRequest {
	s.Type = &v
	return s
}

func (s *DescribeDrdsInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstancesRequestTag struct {
	// The key of the tag configured for the instances you want to query.
	//
	// example:
	//
	// acs:newretail:domain
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag configured for the instances you want to query.
	//
	// example:
	//
	// NEW_RETAIL
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDrdsInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDrdsInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDrdsInstancesRequestTag) SetKey(v string) *DescribeDrdsInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstancesRequestTag) SetValue(v string) *DescribeDrdsInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDrdsInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
