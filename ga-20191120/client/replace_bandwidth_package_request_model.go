// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ReplaceBandwidthPackageRequest
	GetAcceleratorId() *string
	SetBandwidthPackageId(v string) *ReplaceBandwidthPackageRequest
	GetBandwidthPackageId() *string
	SetRegionId(v string) *ReplaceBandwidthPackageRequest
	GetRegionId() *string
	SetTargetBandwidthPackageId(v string) *ReplaceBandwidthPackageRequest
	GetTargetBandwidthPackageId() *string
}

type ReplaceBandwidthPackageRequest struct {
	// The GA instance ID.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the required bandwidth plan. When you specify a replacement bandwidth plan, take note of the following items:
	//
	// 	- Only a bandwidth plan that is not associated with a GA instance can be specified.
	//
	// 	- If you want to replace a basic bandwidth plan, make sure that the bandwidth provided by the replacement bandwidth plan is not less than the total bandwidth allocated to the acceleration area.
	//
	// This parameter is required.
	//
	// example:
	//
	// gbwp-bp176neb61yhcymow****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the bandwidth plan that you want to replace.
	//
	// This parameter is required.
	//
	// example:
	//
	// gbwp-o978hgeb61yhcymow****
	TargetBandwidthPackageId *string `json:"TargetBandwidthPackageId,omitempty" xml:"TargetBandwidthPackageId,omitempty"`
}

func (s ReplaceBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *ReplaceBandwidthPackageRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ReplaceBandwidthPackageRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *ReplaceBandwidthPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReplaceBandwidthPackageRequest) GetTargetBandwidthPackageId() *string {
	return s.TargetBandwidthPackageId
}

func (s *ReplaceBandwidthPackageRequest) SetAcceleratorId(v string) *ReplaceBandwidthPackageRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ReplaceBandwidthPackageRequest) SetBandwidthPackageId(v string) *ReplaceBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *ReplaceBandwidthPackageRequest) SetRegionId(v string) *ReplaceBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *ReplaceBandwidthPackageRequest) SetTargetBandwidthPackageId(v string) *ReplaceBandwidthPackageRequest {
	s.TargetBandwidthPackageId = &v
	return s
}

func (s *ReplaceBandwidthPackageRequest) Validate() error {
	return dara.Validate(s)
}
