// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartShiftLoadBalancerZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartShiftLoadBalancerZonesRequest
	GetClientToken() *string
	SetDryRun(v bool) *StartShiftLoadBalancerZonesRequest
	GetDryRun() *bool
	SetLoadBalancerId(v string) *StartShiftLoadBalancerZonesRequest
	GetLoadBalancerId() *string
	SetZoneMappings(v []*StartShiftLoadBalancerZonesRequestZoneMappings) *StartShiftLoadBalancerZonesRequest
	GetZoneMappings() []*StartShiftLoadBalancerZonesRequestZoneMappings
}

type StartShiftLoadBalancerZonesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
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
	// The ALB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alb-o9ulmq5hgn68jk****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The mappings between zones and vSwitches.
	//
	// >  You can remove only one zone in each call.
	//
	// This parameter is required.
	ZoneMappings []*StartShiftLoadBalancerZonesRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s StartShiftLoadBalancerZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s StartShiftLoadBalancerZonesRequest) GoString() string {
	return s.String()
}

func (s *StartShiftLoadBalancerZonesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartShiftLoadBalancerZonesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StartShiftLoadBalancerZonesRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *StartShiftLoadBalancerZonesRequest) GetZoneMappings() []*StartShiftLoadBalancerZonesRequestZoneMappings {
	return s.ZoneMappings
}

func (s *StartShiftLoadBalancerZonesRequest) SetClientToken(v string) *StartShiftLoadBalancerZonesRequest {
	s.ClientToken = &v
	return s
}

func (s *StartShiftLoadBalancerZonesRequest) SetDryRun(v bool) *StartShiftLoadBalancerZonesRequest {
	s.DryRun = &v
	return s
}

func (s *StartShiftLoadBalancerZonesRequest) SetLoadBalancerId(v string) *StartShiftLoadBalancerZonesRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *StartShiftLoadBalancerZonesRequest) SetZoneMappings(v []*StartShiftLoadBalancerZonesRequestZoneMappings) *StartShiftLoadBalancerZonesRequest {
	s.ZoneMappings = v
	return s
}

func (s *StartShiftLoadBalancerZonesRequest) Validate() error {
	if s.ZoneMappings != nil {
		for _, item := range s.ZoneMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartShiftLoadBalancerZonesRequestZoneMappings struct {
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

func (s StartShiftLoadBalancerZonesRequestZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s StartShiftLoadBalancerZonesRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *StartShiftLoadBalancerZonesRequestZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *StartShiftLoadBalancerZonesRequestZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *StartShiftLoadBalancerZonesRequestZoneMappings) SetVSwitchId(v string) *StartShiftLoadBalancerZonesRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *StartShiftLoadBalancerZonesRequestZoneMappings) SetZoneId(v string) *StartShiftLoadBalancerZonesRequestZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *StartShiftLoadBalancerZonesRequestZoneMappings) Validate() error {
	return dara.Validate(s)
}
