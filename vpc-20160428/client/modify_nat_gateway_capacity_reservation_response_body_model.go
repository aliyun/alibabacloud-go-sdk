// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatGatewayCapacityReservationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityReservationStatus(v *ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) *ModifyNatGatewayCapacityReservationResponseBody
	GetCapacityReservationStatus() *ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus
	SetDecreaseRequestsRemaining(v int32) *ModifyNatGatewayCapacityReservationResponseBody
	GetDecreaseRequestsRemaining() *int32
	SetNatGatewayCapacity(v *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) *ModifyNatGatewayCapacityReservationResponseBody
	GetNatGatewayCapacity() *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayCapacity
	SetNatGatewayId(v string) *ModifyNatGatewayCapacityReservationResponseBody
	GetNatGatewayId() *string
	SetNatGatewayPerformanceValues(v *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) *ModifyNatGatewayCapacityReservationResponseBody
	GetNatGatewayPerformanceValues() *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues
	SetRequestId(v string) *ModifyNatGatewayCapacityReservationResponseBody
	GetRequestId() *string
}

type ModifyNatGatewayCapacityReservationResponseBody struct {
	// The status of the resource reservation.
	CapacityReservationStatus *ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus `json:"CapacityReservationStatus,omitempty" xml:"CapacityReservationStatus,omitempty" type:"Struct"`
	// The remaining number of times that you can downgrade or cancel the resource reservation.
	//
	// example:
	//
	// 2
	DecreaseRequestsRemaining *int32 `json:"DecreaseRequestsRemaining,omitempty" xml:"DecreaseRequestsRemaining,omitempty"`
	// The number of capacity units (CUs) for the NAT resource reservation.
	NatGatewayCapacity *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayCapacity `json:"NatGatewayCapacity,omitempty" xml:"NatGatewayCapacity,omitempty" type:"Struct"`
	// The ID of the NAT gateway to be modified.
	//
	// example:
	//
	// ngw-112za33e4****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The performance metrics of the NAT resource reservation.
	NatGatewayPerformanceValues *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues `json:"NatGatewayPerformanceValues,omitempty" xml:"NatGatewayPerformanceValues,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// B400EF57-60E3-4D61-B8FB-7FA8F72DF5A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNatGatewayCapacityReservationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayCapacityReservationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) GetCapacityReservationStatus() *ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus {
	return s.CapacityReservationStatus
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) GetDecreaseRequestsRemaining() *int32 {
	return s.DecreaseRequestsRemaining
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) GetNatGatewayCapacity() *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayCapacity {
	return s.NatGatewayCapacity
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) GetNatGatewayPerformanceValues() *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues {
	return s.NatGatewayPerformanceValues
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) SetCapacityReservationStatus(v *ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) *ModifyNatGatewayCapacityReservationResponseBody {
	s.CapacityReservationStatus = v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) SetDecreaseRequestsRemaining(v int32) *ModifyNatGatewayCapacityReservationResponseBody {
	s.DecreaseRequestsRemaining = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) SetNatGatewayCapacity(v *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) *ModifyNatGatewayCapacityReservationResponseBody {
	s.NatGatewayCapacity = v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) SetNatGatewayId(v string) *ModifyNatGatewayCapacityReservationResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) SetNatGatewayPerformanceValues(v *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) *ModifyNatGatewayCapacityReservationResponseBody {
	s.NatGatewayPerformanceValues = v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) SetRequestId(v string) *ModifyNatGatewayCapacityReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBody) Validate() error {
	if s.CapacityReservationStatus != nil {
		if err := s.CapacityReservationStatus.Validate(); err != nil {
			return err
		}
	}
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

type ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus struct {
	// The reason for the processing failure.
	//
	// example:
	//
	// No Available Resource
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the resource reservation. Valid values:
	//
	// - pending: The reservation is in progress.
	//
	// - provisioned: The reservation is completed.
	//
	// - failed: The reservation failed.
	//
	// - deleting: The reservation is being deleted.
	//
	// - deleted: The reservation is deleted.
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) GetReason() *string {
	return s.Reason
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) GetStatus() *string {
	return s.Status
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) SetReason(v string) *ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus {
	s.Reason = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) SetStatus(v string) *ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus {
	s.Status = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) Validate() error {
	return dara.Validate(s)
}

type ModifyNatGatewayCapacityReservationResponseBodyNatGatewayCapacity struct {
	// The number of LCUs reserved per hour.
	//
	// example:
	//
	// 100
	CapacityUnits *float64 `json:"CapacityUnits,omitempty" xml:"CapacityUnits,omitempty"`
}

func (s ModifyNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) GetCapacityUnits() *float64 {
	return s.CapacityUnits
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) SetCapacityUnits(v float64) *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayCapacity {
	s.CapacityUnits = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) Validate() error {
	return dara.Validate(s)
}

type ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues struct {
	// The bandwidth value. Unit: Gbit/s.
	//
	// example:
	//
	// 20
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of new connections.
	//
	// example:
	//
	// 100000
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The number of packets per second.
	//
	// example:
	//
	// 3200000
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The number of sessions.
	//
	// example:
	//
	// 2000000
	SessionCount *int32 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
}

func (s ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) GetCps() *int32 {
	return s.Cps
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) GetPps() *int32 {
	return s.Pps
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) GetSessionCount() *int32 {
	return s.SessionCount
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) SetBandwidth(v int32) *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues {
	s.Bandwidth = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) SetCps(v int32) *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues {
	s.Cps = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) SetPps(v int32) *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues {
	s.Pps = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) SetSessionCount(v int32) *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues {
	s.SessionCount = &v
	return s
}

func (s *ModifyNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) Validate() error {
	return dara.Validate(s)
}
