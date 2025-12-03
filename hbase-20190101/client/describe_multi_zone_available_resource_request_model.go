// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiZoneAvailableResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *DescribeMultiZoneAvailableResourceRequest
	GetChargeType() *string
	SetRegionId(v string) *DescribeMultiZoneAvailableResourceRequest
	GetRegionId() *string
	SetZoneCombination(v string) *DescribeMultiZoneAvailableResourceRequest
	GetZoneCombination() *string
}

type DescribeMultiZoneAvailableResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-bef-aliyun
	ZoneCombination *string `json:"ZoneCombination,omitempty" xml:"ZoneCombination,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeMultiZoneAvailableResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMultiZoneAvailableResourceRequest) GetZoneCombination() *string {
	return s.ZoneCombination
}

func (s *DescribeMultiZoneAvailableResourceRequest) SetChargeType(v string) *DescribeMultiZoneAvailableResourceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceRequest) SetRegionId(v string) *DescribeMultiZoneAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceRequest) SetZoneCombination(v string) *DescribeMultiZoneAvailableResourceRequest {
	s.ZoneCombination = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceRequest) Validate() error {
	return dara.Validate(s)
}
