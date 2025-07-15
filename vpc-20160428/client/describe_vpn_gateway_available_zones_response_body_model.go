// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnGatewayAvailableZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZoneIdList(v []*DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList) *DescribeVpnGatewayAvailableZonesResponseBody
	GetAvailableZoneIdList() []*DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList
	SetRegionId(v string) *DescribeVpnGatewayAvailableZonesResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeVpnGatewayAvailableZonesResponseBody
	GetRequestId() *string
}

type DescribeVpnGatewayAvailableZonesResponseBody struct {
	// The zones.
	AvailableZoneIdList []*DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList `json:"AvailableZoneIdList,omitempty" xml:"AvailableZoneIdList,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 29784052-931F-543D-A612-36B3838163FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVpnGatewayAvailableZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewayAvailableZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewayAvailableZonesResponseBody) GetAvailableZoneIdList() []*DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList {
	return s.AvailableZoneIdList
}

func (s *DescribeVpnGatewayAvailableZonesResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpnGatewayAvailableZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpnGatewayAvailableZonesResponseBody) SetAvailableZoneIdList(v []*DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList) *DescribeVpnGatewayAvailableZonesResponseBody {
	s.AvailableZoneIdList = v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesResponseBody) SetRegionId(v string) *DescribeVpnGatewayAvailableZonesResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesResponseBody) SetRequestId(v string) *DescribeVpnGatewayAvailableZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList struct {
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone name.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList) SetZoneId(v string) *DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList {
	s.ZoneId = &v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList) SetZoneName(v string) *DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList {
	s.ZoneName = &v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesResponseBodyAvailableZoneIdList) Validate() error {
	return dara.Validate(s)
}
