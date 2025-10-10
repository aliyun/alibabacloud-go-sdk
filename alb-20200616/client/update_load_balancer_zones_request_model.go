// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateLoadBalancerZonesRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateLoadBalancerZonesRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *UpdateLoadBalancerZonesRequest
	GetLoadBalancerId() *string
	SetZoneMappings(v []*UpdateLoadBalancerZonesRequestZoneMappings) *UpdateLoadBalancerZonesRequest
	GetZoneMappings() []*UpdateLoadBalancerZonesRequestZoneMappings
}

type UpdateLoadBalancerZonesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the ALB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The zones and the vSwitches in the zones. You can specify a maximum of 10 zones. If the selected region supports two or more zones, select at least two zones to ensure the high availability of your service. The specified zones and vSwitches overwrite the existing configurations.
	//
	// This parameter is required.
	ZoneMappings []*UpdateLoadBalancerZonesRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s UpdateLoadBalancerZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerZonesRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateLoadBalancerZonesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateLoadBalancerZonesRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *UpdateLoadBalancerZonesRequest) GetZoneMappings() []*UpdateLoadBalancerZonesRequestZoneMappings {
	return s.ZoneMappings
}

func (s *UpdateLoadBalancerZonesRequest) SetClientToken(v string) *UpdateLoadBalancerZonesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetDryRun(v bool) *UpdateLoadBalancerZonesRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetLoadBalancerId(v string) *UpdateLoadBalancerZonesRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) SetZoneMappings(v []*UpdateLoadBalancerZonesRequestZoneMappings) *UpdateLoadBalancerZonesRequest {
	s.ZoneMappings = v
	return s
}

func (s *UpdateLoadBalancerZonesRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateLoadBalancerZonesRequestZoneMappings struct {
	// The type of EIP. Valid values:
	//
	// 	- **Common**: an EIP.
	//
	// 	- **Anycast**: an Anycast EIP.
	//
	// >  For more information about the regions in which ALB supports Anycast EIPs, see [Limits](https://help.aliyun.com/document_detail/460727.html).
	//
	// example:
	//
	// Common
	EipType *string `json:"EipType,omitempty" xml:"EipType,omitempty"`
	// The private IPv4 address. You must specify at least two zones. You can specify a maximum of 10 zones.
	//
	// example:
	//
	// 192.168.10.1
	IntranetAddress *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	// The ID of the vSwitch in the zone. By default, each zone contains one vSwitch and one subnet. You can specify at most 10 zones. If the region supports two or more zones, specify at least two zones.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1rmcrwg3erh1fh8****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone name. You can call the [DescribeZones](https://help.aliyun.com/document_detail/189196.html) operation to query the most recent zone list. You can specify at most 10 zones. If the region supports two or more zones, specify at least two zones.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateLoadBalancerZonesRequestZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerZonesRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) GetEipType() *string {
	return s.EipType
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) GetIntranetAddress() *string {
	return s.IntranetAddress
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetEipType(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.EipType = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetIntranetAddress(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.IntranetAddress = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetVSwitchId(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) SetZoneId(v string) *UpdateLoadBalancerZonesRequestZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *UpdateLoadBalancerZonesRequestZoneMappings) Validate() error {
	return dara.Validate(s)
}
