// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeAcceleratorRequest
	GetAcceleratorId() *string
	SetRegionId(v string) *DescribeAcceleratorRequest
	GetRegionId() *string
}

type DescribeAcceleratorRequest struct {
	// The ID of the GA instance that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// Deprecated
	//
	// The region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAcceleratorRequest) SetAcceleratorId(v string) *DescribeAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeAcceleratorRequest) SetRegionId(v string) *DescribeAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAcceleratorRequest) Validate() error {
	return dara.Validate(s)
}
