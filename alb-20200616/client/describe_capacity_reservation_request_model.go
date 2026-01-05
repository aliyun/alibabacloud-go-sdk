// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapacityReservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *DescribeCapacityReservationRequest
	GetLoadBalancerId() *string
}

type DescribeCapacityReservationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// alb-o9ulmq5hgn68jk****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DescribeCapacityReservationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationRequest) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeCapacityReservationRequest) SetLoadBalancerId(v string) *DescribeCapacityReservationRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeCapacityReservationRequest) Validate() error {
	return dara.Validate(s)
}
