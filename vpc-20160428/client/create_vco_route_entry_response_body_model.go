// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVcoRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *CreateVcoRouteEntryResponseBody
	GetCreateTime() *int64
	SetDescription(v string) *CreateVcoRouteEntryResponseBody
	GetDescription() *string
	SetNextHop(v string) *CreateVcoRouteEntryResponseBody
	GetNextHop() *string
	SetOverlayMode(v string) *CreateVcoRouteEntryResponseBody
	GetOverlayMode() *string
	SetRequestId(v string) *CreateVcoRouteEntryResponseBody
	GetRequestId() *string
	SetRouteDest(v string) *CreateVcoRouteEntryResponseBody
	GetRouteDest() *string
	SetState(v string) *CreateVcoRouteEntryResponseBody
	GetState() *string
	SetVpnConnectionId(v string) *CreateVcoRouteEntryResponseBody
	GetVpnConnectionId() *string
	SetWeight(v int32) *CreateVcoRouteEntryResponseBody
	GetWeight() *int32
}

type CreateVcoRouteEntryResponseBody struct {
	// The timestamp when the destination-based route was created. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1658387202664
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the destination-based route.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The next hop of the destination-based route.
	//
	// example:
	//
	// vco-p0w2jpkhi2eeop6q6****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The tunneling protocol.
	//
	// The value is set to **Ipsec**, which indicates the IPsec tunneling protocol.
	//
	// example:
	//
	// Ipsec
	OverlayMode *string `json:"OverlayMode,omitempty" xml:"OverlayMode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CFC4D13B-E680-3985-95B1-87AA155481DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The destination CIDR block of the destination-based route.
	//
	// example:
	//
	// 192.168.10.0/24
	RouteDest *string `json:"RouteDest,omitempty" xml:"RouteDest,omitempty"`
	// The status of the destination-based route.
	//
	// Only **published*	- is returned, which indicates that the current route is published to the transit router.
	//
	// example:
	//
	// published
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// example:
	//
	// vco-p0w2jpkhi2eeop6q6****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
	// The weight of the destination-based route. Valid values:
	//
	// 	- **0**: a low priority.
	//
	// 	- **100**: a high priority.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateVcoRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVcoRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVcoRouteEntryResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateVcoRouteEntryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateVcoRouteEntryResponseBody) GetNextHop() *string {
	return s.NextHop
}

func (s *CreateVcoRouteEntryResponseBody) GetOverlayMode() *string {
	return s.OverlayMode
}

func (s *CreateVcoRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVcoRouteEntryResponseBody) GetRouteDest() *string {
	return s.RouteDest
}

func (s *CreateVcoRouteEntryResponseBody) GetState() *string {
	return s.State
}

func (s *CreateVcoRouteEntryResponseBody) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *CreateVcoRouteEntryResponseBody) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateVcoRouteEntryResponseBody) SetCreateTime(v int64) *CreateVcoRouteEntryResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateVcoRouteEntryResponseBody) SetDescription(v string) *CreateVcoRouteEntryResponseBody {
	s.Description = &v
	return s
}

func (s *CreateVcoRouteEntryResponseBody) SetNextHop(v string) *CreateVcoRouteEntryResponseBody {
	s.NextHop = &v
	return s
}

func (s *CreateVcoRouteEntryResponseBody) SetOverlayMode(v string) *CreateVcoRouteEntryResponseBody {
	s.OverlayMode = &v
	return s
}

func (s *CreateVcoRouteEntryResponseBody) SetRequestId(v string) *CreateVcoRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVcoRouteEntryResponseBody) SetRouteDest(v string) *CreateVcoRouteEntryResponseBody {
	s.RouteDest = &v
	return s
}

func (s *CreateVcoRouteEntryResponseBody) SetState(v string) *CreateVcoRouteEntryResponseBody {
	s.State = &v
	return s
}

func (s *CreateVcoRouteEntryResponseBody) SetVpnConnectionId(v string) *CreateVcoRouteEntryResponseBody {
	s.VpnConnectionId = &v
	return s
}

func (s *CreateVcoRouteEntryResponseBody) SetWeight(v int32) *CreateVcoRouteEntryResponseBody {
	s.Weight = &v
	return s
}

func (s *CreateVcoRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
