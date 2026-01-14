// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DeleteBasicAcceleratorRequest
	GetAcceleratorId() *string
	SetRegionId(v string) *DeleteBasicAcceleratorRequest
	GetRegionId() *string
}

type DeleteBasicAcceleratorRequest struct {
	// The ID of the basic GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the region to which the basic GA instance belongs. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBasicAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *DeleteBasicAcceleratorRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DeleteBasicAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBasicAcceleratorRequest) SetAcceleratorId(v string) *DeleteBasicAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteBasicAcceleratorRequest) SetRegionId(v string) *DeleteBasicAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBasicAcceleratorRequest) Validate() error {
	return dara.Validate(s)
}
