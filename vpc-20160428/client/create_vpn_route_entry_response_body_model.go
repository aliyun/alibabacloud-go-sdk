// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *CreateVpnRouteEntryResponseBody
	GetCreateTime() *int64
	SetDescription(v string) *CreateVpnRouteEntryResponseBody
	GetDescription() *string
	SetNextHop(v string) *CreateVpnRouteEntryResponseBody
	GetNextHop() *string
	SetOverlayMode(v string) *CreateVpnRouteEntryResponseBody
	GetOverlayMode() *string
	SetRequestId(v string) *CreateVpnRouteEntryResponseBody
	GetRequestId() *string
	SetRouteDest(v string) *CreateVpnRouteEntryResponseBody
	GetRouteDest() *string
	SetState(v string) *CreateVpnRouteEntryResponseBody
	GetState() *string
	SetVpnInstanceId(v string) *CreateVpnRouteEntryResponseBody
	GetVpnInstanceId() *string
	SetWeight(v int32) *CreateVpnRouteEntryResponseBody
	GetWeight() *int32
}

type CreateVpnRouteEntryResponseBody struct {
	// The timestamp when the destination-based route was created. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1492747187000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the destination-based route.
	//
	// example:
	//
	// mytest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The next hop of the destination-based route.
	//
	// example:
	//
	// vco-bp15oes1py4i66rmd****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The tunneling protocol. The value is set to **Ipsec**, which indicates the IPsec tunneling protocol.
	//
	// example:
	//
	// Ipsec
	OverlayMode *string `json:"OverlayMode,omitempty" xml:"OverlayMode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The destination CIDR block of the destination-based route.
	//
	// example:
	//
	// 10.0.0.0/24
	RouteDest *string `json:"RouteDest,omitempty" xml:"RouteDest,omitempty"`
	// The status of the destination-based route.
	//
	// 	- **published**: advertised to the VPC route table.
	//
	// 	- **normal**: not advertised to the VPC route table.
	//
	// example:
	//
	// published
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1a3kqjiiq9legfx****
	VpnInstanceId *string `json:"VpnInstanceId,omitempty" xml:"VpnInstanceId,omitempty"`
	// The weight of the destination-based route. Valid values:
	//
	// 	- **100**: a high priority
	//
	// 	- **0**: a low priority
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateVpnRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpnRouteEntryResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateVpnRouteEntryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateVpnRouteEntryResponseBody) GetNextHop() *string {
	return s.NextHop
}

func (s *CreateVpnRouteEntryResponseBody) GetOverlayMode() *string {
	return s.OverlayMode
}

func (s *CreateVpnRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpnRouteEntryResponseBody) GetRouteDest() *string {
	return s.RouteDest
}

func (s *CreateVpnRouteEntryResponseBody) GetState() *string {
	return s.State
}

func (s *CreateVpnRouteEntryResponseBody) GetVpnInstanceId() *string {
	return s.VpnInstanceId
}

func (s *CreateVpnRouteEntryResponseBody) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateVpnRouteEntryResponseBody) SetCreateTime(v int64) *CreateVpnRouteEntryResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateVpnRouteEntryResponseBody) SetDescription(v string) *CreateVpnRouteEntryResponseBody {
	s.Description = &v
	return s
}

func (s *CreateVpnRouteEntryResponseBody) SetNextHop(v string) *CreateVpnRouteEntryResponseBody {
	s.NextHop = &v
	return s
}

func (s *CreateVpnRouteEntryResponseBody) SetOverlayMode(v string) *CreateVpnRouteEntryResponseBody {
	s.OverlayMode = &v
	return s
}

func (s *CreateVpnRouteEntryResponseBody) SetRequestId(v string) *CreateVpnRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpnRouteEntryResponseBody) SetRouteDest(v string) *CreateVpnRouteEntryResponseBody {
	s.RouteDest = &v
	return s
}

func (s *CreateVpnRouteEntryResponseBody) SetState(v string) *CreateVpnRouteEntryResponseBody {
	s.State = &v
	return s
}

func (s *CreateVpnRouteEntryResponseBody) SetVpnInstanceId(v string) *CreateVpnRouteEntryResponseBody {
	s.VpnInstanceId = &v
	return s
}

func (s *CreateVpnRouteEntryResponseBody) SetWeight(v int32) *CreateVpnRouteEntryResponseBody {
	s.Weight = &v
	return s
}

func (s *CreateVpnRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
