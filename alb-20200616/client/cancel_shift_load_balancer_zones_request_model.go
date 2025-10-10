// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelShiftLoadBalancerZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelShiftLoadBalancerZonesRequest
	GetClientToken() *string
	SetDryRun(v bool) *CancelShiftLoadBalancerZonesRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *CancelShiftLoadBalancerZonesRequest
	GetLoadBalancerId() *string
	SetZoneMappings(v []*CancelShiftLoadBalancerZonesRequestZoneMappings) *CancelShiftLoadBalancerZonesRequest
	GetZoneMappings() []*CancelShiftLoadBalancerZonesRequestZoneMappings
}

type CancelShiftLoadBalancerZonesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
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
	// alb-o9ulmq5hgn68jk****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The mappings between zones and vSwitches.
	//
	// >  You can add only one zone in each call.
	//
	// This parameter is required.
	ZoneMappings []*CancelShiftLoadBalancerZonesRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s CancelShiftLoadBalancerZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelShiftLoadBalancerZonesRequest) GoString() string {
	return s.String()
}

func (s *CancelShiftLoadBalancerZonesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelShiftLoadBalancerZonesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CancelShiftLoadBalancerZonesRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CancelShiftLoadBalancerZonesRequest) GetZoneMappings() []*CancelShiftLoadBalancerZonesRequestZoneMappings {
	return s.ZoneMappings
}

func (s *CancelShiftLoadBalancerZonesRequest) SetClientToken(v string) *CancelShiftLoadBalancerZonesRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelShiftLoadBalancerZonesRequest) SetDryRun(v bool) *CancelShiftLoadBalancerZonesRequest {
	s.DryRun = &v
	return s
}

func (s *CancelShiftLoadBalancerZonesRequest) SetLoadBalancerId(v string) *CancelShiftLoadBalancerZonesRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CancelShiftLoadBalancerZonesRequest) SetZoneMappings(v []*CancelShiftLoadBalancerZonesRequestZoneMappings) *CancelShiftLoadBalancerZonesRequest {
	s.ZoneMappings = v
	return s
}

func (s *CancelShiftLoadBalancerZonesRequest) Validate() error {
	return dara.Validate(s)
}

type CancelShiftLoadBalancerZonesRequestZoneMappings struct {
	// The ID of the vSwitch in the zone. By default, each zone uses one vSwitch and one subnet.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1rmcrwg3erh1fh8****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID. You can call the [DescribeZones](https://help.aliyun.com/document_detail/189196.html) operation to query the most recent zone list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CancelShiftLoadBalancerZonesRequestZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s CancelShiftLoadBalancerZonesRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *CancelShiftLoadBalancerZonesRequestZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CancelShiftLoadBalancerZonesRequestZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *CancelShiftLoadBalancerZonesRequestZoneMappings) SetVSwitchId(v string) *CancelShiftLoadBalancerZonesRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *CancelShiftLoadBalancerZonesRequestZoneMappings) SetZoneId(v string) *CancelShiftLoadBalancerZonesRequestZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *CancelShiftLoadBalancerZonesRequestZoneMappings) Validate() error {
	return dara.Validate(s)
}
