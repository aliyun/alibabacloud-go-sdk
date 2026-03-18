// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExcpetionCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeExcpetionCountRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeExcpetionCountRequest
	GetResourceGroupId() *string
}

type DescribeExcpetionCountRequest struct {
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeExcpetionCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcpetionCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExcpetionCountRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeExcpetionCountRequest) SetRegionId(v string) *DescribeExcpetionCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeExcpetionCountRequest) SetResourceGroupId(v string) *DescribeExcpetionCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeExcpetionCountRequest) Validate() error {
	return dara.Validate(s)
}
