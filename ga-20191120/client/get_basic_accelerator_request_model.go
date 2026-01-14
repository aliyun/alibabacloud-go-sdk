// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *GetBasicAcceleratorRequest
	GetAcceleratorId() *string
	SetRegionId(v string) *GetBasicAcceleratorRequest
	GetRegionId() *string
}

type GetBasicAcceleratorRequest struct {
	// The ID of the basic GA instance that you want to query.
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

func (s GetBasicAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *GetBasicAcceleratorRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetBasicAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBasicAcceleratorRequest) SetAcceleratorId(v string) *GetBasicAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicAcceleratorRequest) SetRegionId(v string) *GetBasicAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *GetBasicAcceleratorRequest) Validate() error {
	return dara.Validate(s)
}
