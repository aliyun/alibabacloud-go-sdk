// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSecurityProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SwitchSecurityProxyRequest
	GetLang() *string
	SetNatRouteEntryList(v []*SwitchSecurityProxyRequestNatRouteEntryList) *SwitchSecurityProxyRequest
	GetNatRouteEntryList() []*SwitchSecurityProxyRequestNatRouteEntryList
	SetProxyId(v string) *SwitchSecurityProxyRequest
	GetProxyId() *string
	SetRouteAction(v string) *SwitchSecurityProxyRequest
	GetRouteAction() *string
	SetSwitch(v string) *SwitchSecurityProxyRequest
	GetSwitch() *string
}

type SwitchSecurityProxyRequest struct {
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The route entry list for routing to be redirected or un-redirected. Maximum number of entries: 10,000. When RouteAction (divert or undivert) is specified, a non-empty list must be provided and used according to the incremental semantics. Each entry must include the RouteTableId, DestinationCidr, NextHopId, and NextHopType fields, and the combination of RouteTableId and DestinationCidr must be unique. In the Gateway Load Balancer (GWLB) architecture, when committed together with Switch=open, this parameter is used to specify the route entry for which traffic redirection is enabled in the current request. If no list is provided, the saved list is used. If no active list exists, a parameter error is returned.
	NatRouteEntryList []*SwitchSecurityProxyRequestNatRouteEntryList `json:"NatRouteEntryList,omitempty" xml:"NatRouteEntryList,omitempty" type:"Repeated"`
	// The ID of the NAT firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// proxy-natbf******
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// Makes incremental traffic redirection changes to the route entry passed in this request. Other redirected route tables in the list are not affected. Valid values: divert (enables traffic redirection for the passed route entry) and undivert (shutdown of traffic redirection for the passed route entry). This parameter is active only when the firewall is in the traffic redirection state (normal). This parameter and Switch are mutually exclusive and cannot be both set to non-empty values at the same time. If this parameter is not specified or is set to an empty string, the incremental semantics are not used.
	//
	// example:
	//
	// divert
	RouteAction *string `json:"RouteAction,omitempty" xml:"RouteAction,omitempty"`
	// The security protection switch. Valid values:
	//
	// - open: Enabled.
	//
	// - close: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	Switch *string `json:"Switch,omitempty" xml:"Switch,omitempty"`
}

func (s SwitchSecurityProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchSecurityProxyRequest) GoString() string {
	return s.String()
}

func (s *SwitchSecurityProxyRequest) GetLang() *string {
	return s.Lang
}

func (s *SwitchSecurityProxyRequest) GetNatRouteEntryList() []*SwitchSecurityProxyRequestNatRouteEntryList {
	return s.NatRouteEntryList
}

func (s *SwitchSecurityProxyRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *SwitchSecurityProxyRequest) GetRouteAction() *string {
	return s.RouteAction
}

func (s *SwitchSecurityProxyRequest) GetSwitch() *string {
	return s.Switch
}

func (s *SwitchSecurityProxyRequest) SetLang(v string) *SwitchSecurityProxyRequest {
	s.Lang = &v
	return s
}

func (s *SwitchSecurityProxyRequest) SetNatRouteEntryList(v []*SwitchSecurityProxyRequestNatRouteEntryList) *SwitchSecurityProxyRequest {
	s.NatRouteEntryList = v
	return s
}

func (s *SwitchSecurityProxyRequest) SetProxyId(v string) *SwitchSecurityProxyRequest {
	s.ProxyId = &v
	return s
}

func (s *SwitchSecurityProxyRequest) SetRouteAction(v string) *SwitchSecurityProxyRequest {
	s.RouteAction = &v
	return s
}

func (s *SwitchSecurityProxyRequest) SetSwitch(v string) *SwitchSecurityProxyRequest {
	s.Switch = &v
	return s
}

func (s *SwitchSecurityProxyRequest) Validate() error {
	if s.NatRouteEntryList != nil {
		for _, item := range s.NatRouteEntryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SwitchSecurityProxyRequestNatRouteEntryList struct {
	// The destination CIDR block of the route entry. When RouteAction is specified, only the default route 0.0.0.0/0 is supported. If other CIDR blocks are passed, a parameter error is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The ID of the NAT gateway to which the route entry pointed before traffic redirection. You can obtain the ID by calling the DescribeSecurityProxyResources operation. When RouteAction is specified, this value must be consistent with the NAT gateway involved in the current operation. When traffic redirection is disabled, you must still pass this NAT gateway ID instead of the current next hop ID of the firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-bp1a2b3c4d5e6f7g8h9i0
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The next hop type of the route entry before traffic redirection. When RouteAction is specified, this value is fixed to NatGateway and must be passed in pairs with NextHopId. If other values are passed, a parameter error is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// NatGateway
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The ID of the route table to which the route entry belongs. The route table must exist in the user\\"s VPC. This parameter and DestinationCidr are used together to locate a route. If the route does not exist, the task fails and a message indicating that the route entry does not exist is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1a2b3c4d5e6f7g8h9i0
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s SwitchSecurityProxyRequestNatRouteEntryList) String() string {
	return dara.Prettify(s)
}

func (s SwitchSecurityProxyRequestNatRouteEntryList) GoString() string {
	return s.String()
}

func (s *SwitchSecurityProxyRequestNatRouteEntryList) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *SwitchSecurityProxyRequestNatRouteEntryList) GetNextHopId() *string {
	return s.NextHopId
}

func (s *SwitchSecurityProxyRequestNatRouteEntryList) GetNextHopType() *string {
	return s.NextHopType
}

func (s *SwitchSecurityProxyRequestNatRouteEntryList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *SwitchSecurityProxyRequestNatRouteEntryList) SetDestinationCidr(v string) *SwitchSecurityProxyRequestNatRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *SwitchSecurityProxyRequestNatRouteEntryList) SetNextHopId(v string) *SwitchSecurityProxyRequestNatRouteEntryList {
	s.NextHopId = &v
	return s
}

func (s *SwitchSecurityProxyRequestNatRouteEntryList) SetNextHopType(v string) *SwitchSecurityProxyRequestNatRouteEntryList {
	s.NextHopType = &v
	return s
}

func (s *SwitchSecurityProxyRequestNatRouteEntryList) SetRouteTableId(v string) *SwitchSecurityProxyRequestNatRouteEntryList {
	s.RouteTableId = &v
	return s
}

func (s *SwitchSecurityProxyRequestNatRouteEntryList) Validate() error {
	return dara.Validate(s)
}
