// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceGroupSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeResourceGroupSpecRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeResourceGroupSpecRequest
	GetRegionId() *string
	SetResourceGroupType(v string) *DescribeResourceGroupSpecRequest
	GetResourceGroupType() *string
}

type DescribeResourceGroupSpecRequest struct {
	// The Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-8vbo40tl1dxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ai
	ResourceGroupType *string `json:"ResourceGroupType,omitempty" xml:"ResourceGroupType,omitempty"`
}

func (s DescribeResourceGroupSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupSpecRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeResourceGroupSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourceGroupSpecRequest) GetResourceGroupType() *string {
	return s.ResourceGroupType
}

func (s *DescribeResourceGroupSpecRequest) SetDBClusterId(v string) *DescribeResourceGroupSpecRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeResourceGroupSpecRequest) SetRegionId(v string) *DescribeResourceGroupSpecRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceGroupSpecRequest) SetResourceGroupType(v string) *DescribeResourceGroupSpecRequest {
	s.ResourceGroupType = &v
	return s
}

func (s *DescribeResourceGroupSpecRequest) Validate() error {
	return dara.Validate(s)
}
