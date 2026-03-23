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
	GadInstanceName *string `json:"GadInstanceName,omitempty" xml:"GadInstanceName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
