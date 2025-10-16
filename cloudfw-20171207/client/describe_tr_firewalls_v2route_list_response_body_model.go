// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallsV2RouteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallRouteDetailList(v []*DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) *DescribeTrFirewallsV2RouteListResponseBody
	GetFirewallRouteDetailList() []*DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList
	SetRequestId(v string) *DescribeTrFirewallsV2RouteListResponseBody
	GetRequestId() *string
}

type DescribeTrFirewallsV2RouteListResponseBody struct {
	// The route tables of Cloud Firewall.
	FirewallRouteDetailList []*DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList `json:"FirewallRouteDetailList,omitempty" xml:"FirewallRouteDetailList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// ABF190A2-B4D0-53F6-995A-5690A721F91C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTrFirewallsV2RouteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2RouteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2RouteListResponseBody) GetFirewallRouteDetailList() []*DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList {
	return s.FirewallRouteDetailList
}

func (s *DescribeTrFirewallsV2RouteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTrFirewallsV2RouteListResponseBody) SetFirewallRouteDetailList(v []*DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) *DescribeTrFirewallsV2RouteListResponseBody {
	s.FirewallRouteDetailList = v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponseBody) SetRequestId(v string) *DescribeTrFirewallsV2RouteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponseBody) Validate() error {
	if s.FirewallRouteDetailList != nil {
		for _, item := range s.FirewallRouteDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList struct {
	// The destination address of the route.
	//
	// example:
	//
	// 192.168.10.0/25
	TrFirewallRouteDestination *string `json:"TrFirewallRouteDestination,omitempty" xml:"TrFirewallRouteDestination,omitempty"`
	// The ID of the next hop for the route.
	//
	// example:
	//
	// tr-attach-hnxab1y0pxn16p****
	TrFirewallRouteNexthop *string `json:"TrFirewallRouteNexthop,omitempty" xml:"TrFirewallRouteNexthop,omitempty"`
	// The ID of the routing policy.
	//
	// example:
	//
	// policy-04ecbbc6720d4f90****
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
	// The ID of the route table to which the route entry belongs.
	//
	// example:
	//
	// vtb-2zeockxxxorv0mnhz****
	TrFirewallRouteTableId *string `json:"TrFirewallRouteTableId,omitempty" xml:"TrFirewallRouteTableId,omitempty"`
}

func (s DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) GetTrFirewallRouteDestination() *string {
	return s.TrFirewallRouteDestination
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) GetTrFirewallRouteNexthop() *string {
	return s.TrFirewallRouteNexthop
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) GetTrFirewallRoutePolicyId() *string {
	return s.TrFirewallRoutePolicyId
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) GetTrFirewallRouteTableId() *string {
	return s.TrFirewallRouteTableId
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) SetTrFirewallRouteDestination(v string) *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList {
	s.TrFirewallRouteDestination = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) SetTrFirewallRouteNexthop(v string) *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList {
	s.TrFirewallRouteNexthop = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) SetTrFirewallRouteTableId(v string) *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList {
	s.TrFirewallRouteTableId = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) Validate() error {
	return dara.Validate(s)
}
