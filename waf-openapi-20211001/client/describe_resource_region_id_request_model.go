// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceRegionIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeResourceRegionIdRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeResourceRegionIdRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeResourceRegionIdRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeResourceRegionIdRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek2lrm****6pnq
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeResourceRegionIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceRegionIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceRegionIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeResourceRegionIdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourceRegionIdRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeResourceRegionIdRequest) SetInstanceId(v string) *DescribeResourceRegionIdRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceRegionIdRequest) SetRegionId(v string) *DescribeResourceRegionIdRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceRegionIdRequest) SetResourceManagerResourceGroupId(v string) *DescribeResourceRegionIdRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeResourceRegionIdRequest) Validate() error {
	return dara.Validate(s)
}
