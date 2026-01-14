// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAcceleratorServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeAcceleratorServiceStatusRequest
	GetRegionId() *string
}

type DescribeAcceleratorServiceStatusRequest struct {
	// The region ID of the GA instance. Set the value to cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAcceleratorServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorServiceStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAcceleratorServiceStatusRequest) SetRegionId(v string) *DescribeAcceleratorServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAcceleratorServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
