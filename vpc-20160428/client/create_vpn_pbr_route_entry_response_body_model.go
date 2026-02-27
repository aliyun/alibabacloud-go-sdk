// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnPbrRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *CreateVpnPbrRouteEntryResponseBody
	GetCreateTime() *int64
	SetDescription(v string) *CreateVpnPbrRouteEntryResponseBody
	GetDescription() *string
	SetNextHop(v string) *CreateVpnPbrRouteEntryResponseBody
	GetNextHop() *string
	SetOverlayMode(v string) *CreateVpnPbrRouteEntryResponseBody
	GetOverlayMode() *string
	SetPriority(v int32) *CreateVpnPbrRouteEntryResponseBody
	GetPriority() *int32
	SetRequestId(v string) *CreateVpnPbrRouteEntryResponseBody
	GetRequestId() *string
	SetRouteDest(v string) *CreateVpnPbrRouteEntryResponseBody
	GetRouteDest() *string
	SetRouteSource(v string) *CreateVpnPbrRouteEntryResponseBody
	GetRouteSource() *string
	SetState(v string) *CreateVpnPbrRouteEntryResponseBody
	GetState() *string
	SetVpnInstanceId(v string) *CreateVpnPbrRouteEntryResponseBody
	GetVpnInstanceId() *string
	SetWeight(v int32) *CreateVpnPbrRouteEntryResponseBody
	GetWeight() *int32
}

type CreateVpnPbrRouteEntryResponseBody struct {
	// The timestamp generated when the policy-based route was created. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1492747187000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the route.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tunneling protocol. The value is **Ipsec**.
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
	// The priority of the policy-based route.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The priority of the policy-based route.
	//
	// example:
	//
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The VPN gateway ID.
	//
	// example:
	//
	// 10.0.0.0/24
	RouteDest *string `json:"RouteDest,omitempty" xml:"RouteDest,omitempty"`
	// The source CIDR block of the policy-based route.
	//
	// example:
	//
	// 192.168.1.0/24
	RouteSource *string `json:"RouteSource,omitempty" xml:"RouteSource,omitempty"`
	// The status of the policy-based route.
	//
	// 	- **published**: advertised to the VPC route table.
	//
	// 	- **normal**: not advertised to the VPC route table.
	//
	// example:
	//
	// normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1cmw7jh1nfe43m9****
	VpnInstanceId *string `json:"VpnInstanceId,omitempty" xml:"VpnInstanceId,omitempty"`
	// The source CIDR block of the policy-based route.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateVpnPbrRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnPbrRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpnPbrRouteEntryResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateVpnPbrRouteEntryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateVpnPbrRouteEntryResponseBody) GetNextHop() *string {
	return s.NextHop
}

func (s *CreateVpnPbrRouteEntryResponseBody) GetOverlayMode() *string {
	return s.OverlayMode
}

func (s *CreateVpnPbrRouteEntryResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateVpnPbrRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpnPbrRouteEntryResponseBody) GetRouteDest() *string {
	return s.RouteDest
}

func (s *CreateVpnPbrRouteEntryResponseBody) GetRouteSource() *string {
	return s.RouteSource
}

func (s *CreateVpnPbrRouteEntryResponseBody) GetState() *string {
	return s.State
}

func (s *CreateVpnPbrRouteEntryResponseBody) GetVpnInstanceId() *string {
	return s.VpnInstanceId
}

func (s *CreateVpnPbrRouteEntryResponseBody) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateVpnPbrRouteEntryResponseBody) SetCreateTime(v int64) *CreateVpnPbrRouteEntryResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateVpnPbrRouteEntryResponseBody) SetDescription(v string) *CreateVpnPbrRouteEntryResponseBody {
	s.Description = &v
	return s
}

func (s *CreateVpnPbrRouteEntryResponseBody) SetNextHop(v string) *CreateVpnPbrRouteEntryResponseBody {
	s.NextHop = &v
	return s
}

func (s *CreateVpnPbrRouteEntryResponseBody) SetOverlayMode(v string) *CreateVpnPbrRouteEntryResponseBody {
	s.OverlayMode = &v
	return s
}

func (s *CreateVpnPbrRouteEntryResponseBody) SetPriority(v int32) *CreateVpnPbrRouteEntryResponseBody {
	s.Priority = &v
	return s
}

func (s *CreateVpnPbrRouteEntryResponseBody) SetRequestId(v string) *CreateVpnPbrRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpnPbrRouteEntryResponseBody) SetRouteDest(v string) *CreateVpnPbrRouteEntryResponseBody {
	s.RouteDest = &v
	return s
}

func (s *CreateVpnPbrRouteEntryResponseBody) SetRouteSource(v string) *CreateVpnPbrRouteEntryResponseBody {
	s.RouteSource = &v
	return s
}

func (s *CreateVpnPbrRouteEntryResponseBody) SetState(v string) *CreateVpnPbrRouteEntryResponseBody {
	s.State = &v
	return s
}

func (s *CreateVpnPbrRouteEntryResponseBody) SetVpnInstanceId(v string) *CreateVpnPbrRouteEntryResponseBody {
	s.VpnInstanceId = &v
	return s
}

func (s *CreateVpnPbrRouteEntryResponseBody) SetWeight(v int32) *CreateVpnPbrRouteEntryResponseBody {
	s.Weight = &v
	return s
}

func (s *CreateVpnPbrRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
