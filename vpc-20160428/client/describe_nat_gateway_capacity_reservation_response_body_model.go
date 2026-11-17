// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewayCapacityReservationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityReservationStatus(v *DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) *DescribeNatGatewayCapacityReservationResponseBody
	GetCapacityReservationStatus() *DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus
	SetDecreaseRequestsRemaining(v string) *DescribeNatGatewayCapacityReservationResponseBody
	GetDecreaseRequestsRemaining() *string
	SetLastModifiedTime(v int64) *DescribeNatGatewayCapacityReservationResponseBody
	GetLastModifiedTime() *int64
	SetNatGatewayCapacity(v *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) *DescribeNatGatewayCapacityReservationResponseBody
	GetNatGatewayCapacity() *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayCapacity
	SetNatGatewayId(v string) *DescribeNatGatewayCapacityReservationResponseBody
	GetNatGatewayId() *string
	SetNatGatewayPerformanceValues(v *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) *DescribeNatGatewayCapacityReservationResponseBody
	GetNatGatewayPerformanceValues() *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues
	SetRequestId(v string) *DescribeNatGatewayCapacityReservationResponseBody
	GetRequestId() *string
}

type DescribeNatGatewayCapacityReservationResponseBody struct {
	// The resource reservation status.
	CapacityReservationStatus *DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus `json:"CapacityReservationStatus,omitempty" xml:"CapacityReservationStatus,omitempty" type:"Struct"`
	// The remaining number of times that you can downgrade or cancel the resource reservation.
	//
	// example:
	//
	// 2
	DecreaseRequestsRemaining *string `json:"DecreaseRequestsRemaining,omitempty" xml:"DecreaseRequestsRemaining,omitempty"`
	// The time when the resource reservation was last modified.
	//
	// example:
	//
	// 2026-07-02 10:07:51
	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The number of capacity units (CUs) reserved for the NAT gateway.
	NatGatewayCapacity *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayCapacity `json:"NatGatewayCapacity,omitempty" xml:"NatGatewayCapacity,omitempty" type:"Struct"`
	// The ID of the NAT gateway to query.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The performance metrics of the NAT gateway resource reservation.
	NatGatewayPerformanceValues *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues `json:"NatGatewayPerformanceValues,omitempty" xml:"NatGatewayPerformanceValues,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 2315DEB7-5E92-423A-91F7-4C1EC9AD97C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNatGatewayCapacityReservationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayCapacityReservationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) GetCapacityReservationStatus() *DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus {
	return s.CapacityReservationStatus
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) GetDecreaseRequestsRemaining() *string {
	return s.DecreaseRequestsRemaining
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) GetNatGatewayCapacity() *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayCapacity {
	return s.NatGatewayCapacity
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) GetNatGatewayPerformanceValues() *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues {
	return s.NatGatewayPerformanceValues
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) SetCapacityReservationStatus(v *DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) *DescribeNatGatewayCapacityReservationResponseBody {
	s.CapacityReservationStatus = v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) SetDecreaseRequestsRemaining(v string) *DescribeNatGatewayCapacityReservationResponseBody {
	s.DecreaseRequestsRemaining = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) SetLastModifiedTime(v int64) *DescribeNatGatewayCapacityReservationResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) SetNatGatewayCapacity(v *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) *DescribeNatGatewayCapacityReservationResponseBody {
	s.NatGatewayCapacity = v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) SetNatGatewayId(v string) *DescribeNatGatewayCapacityReservationResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) SetNatGatewayPerformanceValues(v *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) *DescribeNatGatewayCapacityReservationResponseBody {
	s.NatGatewayPerformanceValues = v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) SetRequestId(v string) *DescribeNatGatewayCapacityReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBody) Validate() error {
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

type DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus struct {
	// The reason for the processing failure.
	//
	// example:
	//
	// No Available Resource
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The resource reservation status. Valid values:
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

func (s DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) GetReason() *string {
	return s.Reason
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) SetReason(v string) *DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus {
	s.Reason = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) SetStatus(v string) *DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus {
	s.Status = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyCapacityReservationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewayCapacityReservationResponseBodyNatGatewayCapacity struct {
	// The number of load balancer capacity units (LCUs) reserved per hour.
	//
	// example:
	//
	// 900
	CapacityUnits *float64 `json:"CapacityUnits,omitempty" xml:"CapacityUnits,omitempty"`
}

func (s DescribeNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) GetCapacityUnits() *float64 {
	return s.CapacityUnits
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) SetCapacityUnits(v float64) *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayCapacity {
	s.CapacityUnits = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayCapacity) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues struct {
	// The bandwidth value. Unit: Gbit/s.
	//
	// example:
	//
	// 20
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of new connections per second.
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

func (s DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) GetCps() *int32 {
	return s.Cps
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) GetPps() *int32 {
	return s.Pps
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) GetSessionCount() *int32 {
	return s.SessionCount
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) SetBandwidth(v int32) *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues {
	s.Bandwidth = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) SetCps(v int32) *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues {
	s.Cps = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) SetPps(v int32) *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues {
	s.Pps = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) SetSessionCount(v int32) *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues {
	s.SessionCount = &v
	return s
}

func (s *DescribeNatGatewayCapacityReservationResponseBodyNatGatewayPerformanceValues) Validate() error {
	return dara.Validate(s)
}
