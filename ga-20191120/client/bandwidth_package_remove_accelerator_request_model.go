// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBandwidthPackageRemoveAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *BandwidthPackageRemoveAcceleratorRequest
	GetAcceleratorId() *string
	SetBandwidthPackageId(v string) *BandwidthPackageRemoveAcceleratorRequest
	GetBandwidthPackageId() *string
	SetRegionId(v string) *BandwidthPackageRemoveAcceleratorRequest
	GetRegionId() *string
}

type BandwidthPackageRemoveAcceleratorRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1qe94o52ot4pkfn****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BandwidthPackageRemoveAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s BandwidthPackageRemoveAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *BandwidthPackageRemoveAcceleratorRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *BandwidthPackageRemoveAcceleratorRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *BandwidthPackageRemoveAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BandwidthPackageRemoveAcceleratorRequest) SetAcceleratorId(v string) *BandwidthPackageRemoveAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorRequest) SetBandwidthPackageId(v string) *BandwidthPackageRemoveAcceleratorRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorRequest) SetRegionId(v string) *BandwidthPackageRemoveAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *BandwidthPackageRemoveAcceleratorRequest) Validate() error {
	return dara.Validate(s)
}
