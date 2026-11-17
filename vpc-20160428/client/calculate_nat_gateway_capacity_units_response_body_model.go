// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCalculateNatGatewayCapacityUnitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNatGatewayCapacity(v *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayCapacity) *CalculateNatGatewayCapacityUnitsResponseBody
	GetNatGatewayCapacity() *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayCapacity
	SetNatGatewayPerformanceValues(v *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) *CalculateNatGatewayCapacityUnitsResponseBody
	GetNatGatewayPerformanceValues() *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues
	SetRequestId(v string) *CalculateNatGatewayCapacityUnitsResponseBody
	GetRequestId() *string
}

type CalculateNatGatewayCapacityUnitsResponseBody struct {
	// The details of the calculated equivalent reserved CUs.
	NatGatewayCapacity *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayCapacity `json:"NatGatewayCapacity,omitempty" xml:"NatGatewayCapacity,omitempty" type:"Struct"`
	// The details of the calculated equivalent resource reservation metrics.
	NatGatewayPerformanceValues *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues `json:"NatGatewayPerformanceValues,omitempty" xml:"NatGatewayPerformanceValues,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CalculateNatGatewayCapacityUnitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CalculateNatGatewayCapacityUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *CalculateNatGatewayCapacityUnitsResponseBody) GetNatGatewayCapacity() *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayCapacity {
	return s.NatGatewayCapacity
}

func (s *CalculateNatGatewayCapacityUnitsResponseBody) GetNatGatewayPerformanceValues() *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues {
	return s.NatGatewayPerformanceValues
}

func (s *CalculateNatGatewayCapacityUnitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CalculateNatGatewayCapacityUnitsResponseBody) SetNatGatewayCapacity(v *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayCapacity) *CalculateNatGatewayCapacityUnitsResponseBody {
	s.NatGatewayCapacity = v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsResponseBody) SetNatGatewayPerformanceValues(v *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) *CalculateNatGatewayCapacityUnitsResponseBody {
	s.NatGatewayPerformanceValues = v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsResponseBody) SetRequestId(v string) *CalculateNatGatewayCapacityUnitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsResponseBody) Validate() error {
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

type CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayCapacity struct {
	// The number of reserved LCUs per hour.
	//
	// example:
	//
	// 100
	CapacityUnits *float64 `json:"CapacityUnits,omitempty" xml:"CapacityUnits,omitempty"`
}

func (s CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayCapacity) String() string {
	return dara.Prettify(s)
}

func (s CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayCapacity) GoString() string {
	return s.String()
}

func (s *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayCapacity) GetCapacityUnits() *float64 {
	return s.CapacityUnits
}

func (s *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayCapacity) SetCapacityUnits(v float64) *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayCapacity {
	s.CapacityUnits = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayCapacity) Validate() error {
	return dara.Validate(s)
}

type CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues struct {
	// The bandwidth. Unit: Gbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The maximum number of new connections per second.
	//
	// example:
	//
	// 50000
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The number of packets per second. Unit: pps.
	//
	// example:
	//
	// 160000
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The number of sessions.
	//
	// example:
	//
	// 1000
	SessionCount *int32 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
}

func (s CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) String() string {
	return dara.Prettify(s)
}

func (s CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) GoString() string {
	return s.String()
}

func (s *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) GetCps() *int32 {
	return s.Cps
}

func (s *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) GetPps() *int32 {
	return s.Pps
}

func (s *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) GetSessionCount() *int32 {
	return s.SessionCount
}

func (s *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) SetBandwidth(v int32) *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues {
	s.Bandwidth = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) SetCps(v int32) *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues {
	s.Cps = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) SetPps(v int32) *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues {
	s.Pps = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) SetSessionCount(v int32) *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues {
	s.SessionCount = &v
	return s
}

func (s *CalculateNatGatewayCapacityUnitsResponseBodyNatGatewayPerformanceValues) Validate() error {
	return dara.Validate(s)
}
