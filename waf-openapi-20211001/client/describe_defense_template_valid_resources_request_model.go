// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplateValidResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScene(v string) *DescribeDefenseTemplateValidResourcesRequest
	GetDefenseScene() *string
	SetInstanceId(v string) *DescribeDefenseTemplateValidResourcesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeDefenseTemplateValidResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDefenseTemplateValidResourcesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDefenseTemplateValidResourcesRequest
	GetRegionId() *string
	SetResource(v string) *DescribeDefenseTemplateValidResourcesRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplateValidResourcesRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateId(v int64) *DescribeDefenseTemplateValidResourcesRequest
	GetTemplateId() *int64
}

type DescribeDefenseTemplateValidResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// whitelist
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-27a3jyw0v04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// mgw.realperson.antdigital.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// rg-acfm2thcppfv6ay
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// example:
	//
	// 21202
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseTemplateValidResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplateValidResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseTemplateValidResourcesRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetDefenseScene(v string) *DescribeDefenseTemplateValidResourcesRequest {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetInstanceId(v string) *DescribeDefenseTemplateValidResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetPageNumber(v int32) *DescribeDefenseTemplateValidResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetPageSize(v int32) *DescribeDefenseTemplateValidResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetRegionId(v string) *DescribeDefenseTemplateValidResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetResource(v string) *DescribeDefenseTemplateValidResourcesRequest {
	s.Resource = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplateValidResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) SetTemplateId(v int64) *DescribeDefenseTemplateValidResourcesRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesRequest) Validate() error {
	return dara.Validate(s)
}
