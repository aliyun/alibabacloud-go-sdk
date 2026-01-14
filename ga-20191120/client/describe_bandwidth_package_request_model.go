// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *DescribeBandwidthPackageRequest
	GetBandwidthPackageId() *string
	SetRegionId(v string) *DescribeBandwidthPackageRequest
	GetRegionId() *string
}

type DescribeBandwidthPackageRequest struct {
	// The ID of the bandwidth plan that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeBandwidthPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBandwidthPackageRequest) SetBandwidthPackageId(v string) *DescribeBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeBandwidthPackageRequest) SetRegionId(v string) *DescribeBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBandwidthPackageRequest) Validate() error {
	return dara.Validate(s)
}
