// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStatisticSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeStatisticSummaryResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeStatisticSummaryResponseBody
	GetTotalCount() *int64
	SetVpcRequestTops(v *DescribeStatisticSummaryResponseBodyVpcRequestTops) *DescribeStatisticSummaryResponseBody
	GetVpcRequestTops() *DescribeStatisticSummaryResponseBodyVpcRequestTops
	SetZoneRequestTops(v *DescribeStatisticSummaryResponseBodyZoneRequestTops) *DescribeStatisticSummaryResponseBody
	GetZoneRequestTops() *DescribeStatisticSummaryResponseBodyZoneRequestTops
}

type DescribeStatisticSummaryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A73F3BD0-B1A8-42A9-A9B6-689BBABC4891
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2254
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The top three VPCs with the largest number of DNS requests.
	VpcRequestTops *DescribeStatisticSummaryResponseBodyVpcRequestTops `json:"VpcRequestTops,omitempty" xml:"VpcRequestTops,omitempty" type:"Struct"`
	// The top three zones with the largest number of DNS requests.
	ZoneRequestTops *DescribeStatisticSummaryResponseBodyZoneRequestTops `json:"ZoneRequestTops,omitempty" xml:"ZoneRequestTops,omitempty" type:"Struct"`
}

func (s DescribeStatisticSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStatisticSummaryResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeStatisticSummaryResponseBody) GetVpcRequestTops() *DescribeStatisticSummaryResponseBodyVpcRequestTops {
	return s.VpcRequestTops
}

func (s *DescribeStatisticSummaryResponseBody) GetZoneRequestTops() *DescribeStatisticSummaryResponseBodyZoneRequestTops {
	return s.ZoneRequestTops
}

func (s *DescribeStatisticSummaryResponseBody) SetRequestId(v string) *DescribeStatisticSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBody) SetTotalCount(v int64) *DescribeStatisticSummaryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBody) SetVpcRequestTops(v *DescribeStatisticSummaryResponseBodyVpcRequestTops) *DescribeStatisticSummaryResponseBody {
	s.VpcRequestTops = v
	return s
}

func (s *DescribeStatisticSummaryResponseBody) SetZoneRequestTops(v *DescribeStatisticSummaryResponseBodyZoneRequestTops) *DescribeStatisticSummaryResponseBody {
	s.ZoneRequestTops = v
	return s
}

func (s *DescribeStatisticSummaryResponseBody) Validate() error {
	if s.VpcRequestTops != nil {
		if err := s.VpcRequestTops.Validate(); err != nil {
			return err
		}
	}
	if s.ZoneRequestTops != nil {
		if err := s.ZoneRequestTops.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStatisticSummaryResponseBodyVpcRequestTops struct {
	VpcRequestTop []*DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop `json:"VpcRequestTop,omitempty" xml:"VpcRequestTop,omitempty" type:"Repeated"`
}

func (s DescribeStatisticSummaryResponseBodyVpcRequestTops) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBodyVpcRequestTops) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTops) GetVpcRequestTop() []*DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	return s.VpcRequestTop
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTops) SetVpcRequestTop(v []*DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) *DescribeStatisticSummaryResponseBodyVpcRequestTops {
	s.VpcRequestTop = v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTops) Validate() error {
	if s.VpcRequestTop != nil {
		for _, item := range s.VpcRequestTop {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// 华北 2
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The number of DNS requests on the previous day.
	//
	// example:
	//
	// 2254
	RequestCount *int64 `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	// The tunnel ID.
	//
	// example:
	//
	// tun-7h33lkqfuhgnyy****
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-f8zvrvr1payllgz38****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC type. Valid values:
	//
	// 	- STANDARD: standard VPC
	//
	// 	- EDS: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) GetRequestCount() *int64 {
	return s.RequestCount
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) GetTunnelId() *string {
	return s.TunnelId
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) GetVpcType() *string {
	return s.VpcType
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetRegionId(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.RegionId = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetRegionName(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.RegionName = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetRequestCount(v int64) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.RequestCount = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetTunnelId(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.TunnelId = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetVpcId(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.VpcId = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) SetVpcType(v string) *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop {
	s.VpcType = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyVpcRequestTopsVpcRequestTop) Validate() error {
	return dara.Validate(s)
}

type DescribeStatisticSummaryResponseBodyZoneRequestTops struct {
	ZoneRequestTop []*DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop `json:"ZoneRequestTop,omitempty" xml:"ZoneRequestTop,omitempty" type:"Repeated"`
}

func (s DescribeStatisticSummaryResponseBodyZoneRequestTops) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBodyZoneRequestTops) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTops) GetZoneRequestTop() []*DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop {
	return s.ZoneRequestTop
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTops) SetZoneRequestTop(v []*DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) *DescribeStatisticSummaryResponseBodyZoneRequestTops {
	s.ZoneRequestTop = v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTops) Validate() error {
	if s.ZoneRequestTop != nil {
		for _, item := range s.ZoneRequestTop {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop struct {
	// The business type. Valid values:
	//
	// 	- AUTH_ZONE: authoritative zone
	//
	// 	- RESOLVER_RULE: forwarding rule
	//
	// 	- INBOUND: inbound endpoint
	//
	// example:
	//
	// AUTH_ZONE
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The number of DNS requests on the previous day.
	//
	// example:
	//
	// 2251
	RequestCount *int64 `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	// The zone name.
	//
	// example:
	//
	// test.com
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) GoString() string {
	return s.String()
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) GetBizType() *string {
	return s.BizType
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) GetRequestCount() *int64 {
	return s.RequestCount
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) SetBizType(v string) *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop {
	s.BizType = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) SetRequestCount(v int64) *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop {
	s.RequestCount = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) SetZoneName(v string) *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop {
	s.ZoneName = &v
	return s
}

func (s *DescribeStatisticSummaryResponseBodyZoneRequestTopsZoneRequestTop) Validate() error {
	return dara.Validate(s)
}
