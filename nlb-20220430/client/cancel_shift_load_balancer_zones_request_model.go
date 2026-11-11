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
	SetRegionId(v string) *CancelShiftLoadBalancerZonesRequest
	GetRegionId() *string
	SetZoneMappings(v []*CancelShiftLoadBalancerZonesRequestZoneMappings) *CancelShiftLoadBalancerZonesRequest
	GetZoneMappings() []*CancelShiftLoadBalancerZonesRequestZoneMappings
}

type CancelShiftLoadBalancerZonesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client. Make sure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: sends a precheck request without adding the zone back. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): sends a regular request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The NLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// nlb-83ckzc8d4xlp8o****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The region ID of the NLB instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of zone and vSwitch mappings.
	//
	// > You can add back at most one zone in a single call.
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

func (s *CancelShiftLoadBalancerZonesRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *CancelShiftLoadBalancerZonesRequest) SetRegionId(v string) *CancelShiftLoadBalancerZonesRequest {
	s.RegionId = &v
	return s
}

func (s *CancelShiftLoadBalancerZonesRequest) SetZoneMappings(v []*CancelShiftLoadBalancerZonesRequestZoneMappings) *CancelShiftLoadBalancerZonesRequest {
	s.ZoneMappings = v
	return s
}

func (s *CancelShiftLoadBalancerZonesRequest) Validate() error {
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

type CancelShiftLoadBalancerZonesRequestZoneMappings struct {
	// The vSwitch that corresponds to the zone. By default, each zone uses one vSwitch and one subnet.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-sersdf****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the NLB instance.
	//
	// > You can add back at most one zone in a single call.
	//
	// You can call the [DescribeZones](https://help.aliyun.com/document_detail/443890.html) operation to query zone IDs.
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
