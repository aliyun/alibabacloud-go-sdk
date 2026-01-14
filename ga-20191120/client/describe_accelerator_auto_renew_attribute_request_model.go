// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAcceleratorAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeAcceleratorAutoRenewAttributeRequest
	GetAcceleratorId() *string
	SetRegionId(v string) *DescribeAcceleratorAutoRenewAttributeRequest
	GetRegionId() *string
}

type DescribeAcceleratorAutoRenewAttributeRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAcceleratorAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorAutoRenewAttributeRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeAcceleratorAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAcceleratorAutoRenewAttributeRequest) SetAcceleratorId(v string) *DescribeAcceleratorAutoRenewAttributeRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeRequest) SetRegionId(v string) *DescribeAcceleratorAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAcceleratorAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
