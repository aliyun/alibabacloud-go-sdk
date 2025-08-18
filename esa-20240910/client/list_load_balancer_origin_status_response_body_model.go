// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLoadBalancerOriginStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOriginStatus(v []*ListLoadBalancerOriginStatusResponseBodyOriginStatus) *ListLoadBalancerOriginStatusResponseBody
	GetOriginStatus() []*ListLoadBalancerOriginStatusResponseBodyOriginStatus
	SetRequestId(v string) *ListLoadBalancerOriginStatusResponseBody
	GetRequestId() *string
}

type ListLoadBalancerOriginStatusResponseBody struct {
	// List of origin statuses under the load balancer.
	OriginStatus []*ListLoadBalancerOriginStatusResponseBodyOriginStatus `json:"OriginStatus,omitempty" xml:"OriginStatus,omitempty" type:"Repeated"`
	// Request ID, used for tracking the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLoadBalancerOriginStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancerOriginStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerOriginStatusResponseBody) GetOriginStatus() []*ListLoadBalancerOriginStatusResponseBodyOriginStatus {
	return s.OriginStatus
}

func (s *ListLoadBalancerOriginStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLoadBalancerOriginStatusResponseBody) SetOriginStatus(v []*ListLoadBalancerOriginStatusResponseBodyOriginStatus) *ListLoadBalancerOriginStatusResponseBody {
	s.OriginStatus = v
	return s
}

func (s *ListLoadBalancerOriginStatusResponseBody) SetRequestId(v string) *ListLoadBalancerOriginStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalancerOriginStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancerOriginStatusResponseBodyOriginStatus struct {
	// ID of the load balancer.
	//
	// example:
	//
	// 99874066052****
	LoadBalancerId *int64 `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// ID of the origin.
	//
	// example:
	//
	// 99750209487****
	OriginId *int64 `json:"OriginId,omitempty" xml:"OriginId,omitempty"`
	// ID of the source address pool.
	//
	// example:
	//
	// 99750209487****
	PoolId *int64 `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	// The origin pool to which the source belongs, under this load balancer. Only \\"default_pool\\" (default address pool) will be displayed; other types will return an empty string.
	//
	// example:
	//
	// default_pool
	PoolType *string `json:"PoolType,omitempty" xml:"PoolType,omitempty"`
	// Reason for the probe failure.
	//
	// example:
	//
	// TCP connection error
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Status of the origin:
	//
	// - Healthy(healthy): The probe result is available.
	//
	// - Unhealthy(unhealthy): The probe result is unavailable.
	//
	// - Unknown(unknown): Unknown, the monitor has not yet probed.
	//
	// - Undetected(undetected): The load balancer to which the origin belongs is not bound to a monitor.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLoadBalancerOriginStatusResponseBodyOriginStatus) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancerOriginStatusResponseBodyOriginStatus) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) GetLoadBalancerId() *int64 {
	return s.LoadBalancerId
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) GetOriginId() *int64 {
	return s.OriginId
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) GetPoolId() *int64 {
	return s.PoolId
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) GetPoolType() *string {
	return s.PoolType
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) GetReason() *string {
	return s.Reason
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) GetStatus() *string {
	return s.Status
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) SetLoadBalancerId(v int64) *ListLoadBalancerOriginStatusResponseBodyOriginStatus {
	s.LoadBalancerId = &v
	return s
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) SetOriginId(v int64) *ListLoadBalancerOriginStatusResponseBodyOriginStatus {
	s.OriginId = &v
	return s
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) SetPoolId(v int64) *ListLoadBalancerOriginStatusResponseBodyOriginStatus {
	s.PoolId = &v
	return s
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) SetPoolType(v string) *ListLoadBalancerOriginStatusResponseBodyOriginStatus {
	s.PoolType = &v
	return s
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) SetReason(v string) *ListLoadBalancerOriginStatusResponseBodyOriginStatus {
	s.Reason = &v
	return s
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) SetStatus(v string) *ListLoadBalancerOriginStatusResponseBodyOriginStatus {
	s.Status = &v
	return s
}

func (s *ListLoadBalancerOriginStatusResponseBodyOriginStatus) Validate() error {
	return dara.Validate(s)
}
