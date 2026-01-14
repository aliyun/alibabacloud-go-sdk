// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthPackageAutoRenewAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeBandwidthPackageAutoRenewAttributeRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeBandwidthPackageAutoRenewAttributeRequest
	GetRegionId() *string
}

type DescribeBandwidthPackageAutoRenewAttributeRequest struct {
	// The ID of the bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// gbwp-bp1iquvlp8khla5emb3ia
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBandwidthPackageAutoRenewAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackageAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageAutoRenewAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBandwidthPackageAutoRenewAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBandwidthPackageAutoRenewAttributeRequest) SetInstanceId(v string) *DescribeBandwidthPackageAutoRenewAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeRequest) SetRegionId(v string) *DescribeBandwidthPackageAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeRequest) Validate() error {
	return dara.Validate(s)
}
