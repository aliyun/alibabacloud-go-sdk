// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCalculateNatGatewayCapacityUnitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CalculateNatGatewayCapacityUnitsRequest
	GetClientToken() *string
	SetDryRun(v string) *CalculateNatGatewayCapacityUnitsRequest
	GetDryRun() *string
	SetNatGatewayCapacity(v *CalculateNatGatewayCapacityUnitsRequestNatGatewayCapacity) *CalculateNatGatewayCapacityUnitsRequest
	GetNatGatewayCapacity() *CalculateNatGatewayCapacityUnitsRequestNatGatewayCapacity
	SetNatGatewayId(v string) *CalculateNatGatewayCapacityUnitsRequest
	GetNatGatewayId() *string
	SetNatGatewayPerformanceValues(v *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) *CalculateNatGatewayCapacityUnitsRequest
	GetNatGatewayPerformanceValues() *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues
	SetRegionId(v string) *CalculateNatGatewayCapacityUnitsRequest
	GetRegionId() *string
}

type CalculateNatGatewayCapacityUnitsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. Note: If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - true: sends the request without upgrading the Internet NAT gateway. The system checks whether your AccessKey pair is valid, the authorization of the Resource Access Management (RAM) user, and whether the required parameters are specified. If the request fails the check, an error message is returned. If the request passes the check, the DryRunOperation error code is returned.
	//
	// - false (default value): sends a normal request. If the request passes the check, a 2xx HTTP status code is returned and the Internet NAT gateway is directly upgraded.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The details of the reserved capacity units (CUs) to be calculated.
	NatGatewayCapacity *CalculateNatGatewayCapacityUnitsRequestNatGatewayCapacity `json:"NatGatewayCapacity,omitempty" xml:"NatGatewayCapacity,omitempty" type:"Struct"`
	// The instance ID of the NAT gateway.
	//
	// example:
	//
	// ngw-2ze0dcn4mq31qx2jc****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The details of the resource reservation metrics to be calculated.
	NatGatewayPerformanceValues *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues `json:"NatGatewayPerformanceValues,omitempty" xml:"NatGatewayPerformanceValues,omitempty" type:"Struct"`
	// The region ID of the NAT gateway.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CalculateNatGatewayCapacityUnitsRequest) String() string {
	return dara.Prettify(s)
}

func (s CalculateNatGatewayCapacityUnitsRequest) GoString() string {
	return s.String()
}

func (s *CalculateNatGatewayCapacityUnitsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CalculateNatGatewayCapacityUnitsRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *CalculateNatGatewayCapacityUnitsRequest) GetNatGatewayCapacity() *CalculateNatGatewayCapacityUnitsRequestNatGatewayCapacity {
	return s.NatGatewayCapacity
}

func (s *CalculateNatGatewayCapacityUnitsRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CalculateNatGatewayCapacityUnitsRequest) GetNatGatewayPerformanceValues() *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues {
	return s.NatGatewayPerformanceValues
}

func (s *CalculateNatGatewayCapacityUnitsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CalculateNatGatewayCapacityUnitsRequest) SetClientToken(v string) *CalculateNatGatewayCapacityUnitsRequest {
	s.ClientToken = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsRequest) SetDryRun(v string) *CalculateNatGatewayCapacityUnitsRequest {
	s.DryRun = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsRequest) SetNatGatewayCapacity(v *CalculateNatGatewayCapacityUnitsRequestNatGatewayCapacity) *CalculateNatGatewayCapacityUnitsRequest {
	s.NatGatewayCapacity = v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsRequest) SetNatGatewayId(v string) *CalculateNatGatewayCapacityUnitsRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsRequest) SetNatGatewayPerformanceValues(v *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) *CalculateNatGatewayCapacityUnitsRequest {
	s.NatGatewayPerformanceValues = v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsRequest) SetRegionId(v string) *CalculateNatGatewayCapacityUnitsRequest {
	s.RegionId = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsRequest) Validate() error {
	if s.NatGatewayCapacity != nil {
		if err := s.NatGatewayCapacity.Validate(); err != nil {
			return err
		}
	}
	if s.NatGatewayPerformanceValues != nil {
		if err := s.NatGatewayPerformanceValues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CalculateNatGatewayCapacityUnitsRequestNatGatewayCapacity struct {
	// The number of reserved LCUs per hour.
	//
	// example:
	//
	// 100
	CapacityUnits *int32 `json:"CapacityUnits,omitempty" xml:"CapacityUnits,omitempty"`
}

func (s CalculateNatGatewayCapacityUnitsRequestNatGatewayCapacity) String() string {
	return dara.Prettify(s)
}

func (s CalculateNatGatewayCapacityUnitsRequestNatGatewayCapacity) GoString() string {
	return s.String()
}

func (s *CalculateNatGatewayCapacityUnitsRequestNatGatewayCapacity) GetCapacityUnits() *int32 {
	return s.CapacityUnits
}

func (s *CalculateNatGatewayCapacityUnitsRequestNatGatewayCapacity) SetCapacityUnits(v int32) *CalculateNatGatewayCapacityUnitsRequestNatGatewayCapacity {
	s.CapacityUnits = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsRequestNatGatewayCapacity) Validate() error {
	return dara.Validate(s)
}

type CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues struct {
	// The bandwidth. Unit: Gbit/s.
	//
	// example:
	//
	// 20
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The maximum number of new connections per second.
	//
	// example:
	//
	// 500
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The number of packets per second. Unit: pps.
	//
	// example:
	//
	// 50000
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The number of sessions.
	//
	// example:
	//
	// 1000
	SessionCount *int32 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
}

func (s CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) String() string {
	return dara.Prettify(s)
}

func (s CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) GoString() string {
	return s.String()
}

func (s *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) GetCps() *int32 {
	return s.Cps
}

func (s *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) GetPps() *int32 {
	return s.Pps
}

func (s *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) GetSessionCount() *int32 {
	return s.SessionCount
}

func (s *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) SetBandwidth(v int32) *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues {
	s.Bandwidth = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) SetCps(v int32) *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues {
	s.Cps = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) SetPps(v int32) *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues {
	s.Pps = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) SetSessionCount(v int32) *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues {
	s.SessionCount = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsRequestNatGatewayPerformanceValues) Validate() error {
	return dara.Validate(s)
}
