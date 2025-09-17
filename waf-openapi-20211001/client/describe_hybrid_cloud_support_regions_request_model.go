// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudSupportRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeHybridCloudSupportRegionsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeHybridCloudSupportRegionsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudSupportRegionsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudSupportRegionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-uqm3e3s6x03
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeHybridCloudSupportRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudSupportRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudSupportRegionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudSupportRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudSupportRegionsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudSupportRegionsRequest) SetInstanceId(v string) *DescribeHybridCloudSupportRegionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudSupportRegionsRequest) SetRegionId(v string) *DescribeHybridCloudSupportRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudSupportRegionsRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudSupportRegionsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudSupportRegionsRequest) Validate() error {
	return dara.Validate(s)
}
