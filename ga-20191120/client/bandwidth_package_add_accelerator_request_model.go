// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBandwidthPackageAddAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *BandwidthPackageAddAcceleratorRequest
	GetAcceleratorId() *string
	SetBandwidthPackageId(v string) *BandwidthPackageAddAcceleratorRequest
	GetBandwidthPackageId() *string
	SetRegionId(v string) *BandwidthPackageAddAcceleratorRequest
	GetRegionId() *string
}

type BandwidthPackageAddAcceleratorRequest struct {
	// The GA instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1qe94o52ot4pkfn****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The bandwidth plan ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BandwidthPackageAddAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s BandwidthPackageAddAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *BandwidthPackageAddAcceleratorRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *BandwidthPackageAddAcceleratorRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *BandwidthPackageAddAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BandwidthPackageAddAcceleratorRequest) SetAcceleratorId(v string) *BandwidthPackageAddAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorRequest) SetBandwidthPackageId(v string) *BandwidthPackageAddAcceleratorRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorRequest) SetRegionId(v string) *BandwidthPackageAddAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *BandwidthPackageAddAcceleratorRequest) Validate() error {
	return dara.Validate(s)
}
