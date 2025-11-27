// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGadInstanceName(v string) *DescribeGadInstancesRequest
	GetGadInstanceName() *string
	SetRegionId(v string) *DescribeGadInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeGadInstancesRequest
	GetResourceGroupId() *string
}

type DescribeGadInstancesRequest struct {
	// The ID of the global active database cluster.
	//
	// 	- If you leave this parameter empty, this operation returns the details about all global active database clusters that are created within your Alibaba Cloud account.
	//
	// 	- If you specify this parameter, this operation returns the details about the global active database cluster that you specify.
	//
	// >  If you do not specify this parameter when you call this operation for the first time, the IDs of all clusters that are created by using the current account are returned. Then, you can specify the cluster ID to view the cluster details.
	//
	// example:
	//
	// gad-rm-bp1npi2j8********
	GadInstanceName *string `json:"GadInstanceName,omitempty" xml:"GadInstanceName,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeGadInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGadInstancesRequest) GetGadInstanceName() *string {
	return s.GadInstanceName
}

func (s *DescribeGadInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGadInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGadInstancesRequest) SetGadInstanceName(v string) *DescribeGadInstancesRequest {
	s.GadInstanceName = &v
	return s
}

func (s *DescribeGadInstancesRequest) SetRegionId(v string) *DescribeGadInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGadInstancesRequest) SetResourceGroupId(v string) *DescribeGadInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGadInstancesRequest) Validate() error {
	return dara.Validate(s)
}
