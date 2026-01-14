// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableBusiRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListAvailableBusiRegionsRequest
	GetAcceleratorId() *string
	SetRegionId(v string) *ListAvailableBusiRegionsRequest
	GetRegionId() *string
}

type ListAvailableBusiRegionsRequest struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
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

func (s ListAvailableBusiRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableBusiRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableBusiRegionsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListAvailableBusiRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAvailableBusiRegionsRequest) SetAcceleratorId(v string) *ListAvailableBusiRegionsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListAvailableBusiRegionsRequest) SetRegionId(v string) *ListAvailableBusiRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAvailableBusiRegionsRequest) Validate() error {
	return dara.Validate(s)
}
