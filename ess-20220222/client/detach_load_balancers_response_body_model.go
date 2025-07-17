// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachLoadBalancersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachLoadBalancersResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *DetachLoadBalancersResponseBody
	GetScalingActivityId() *string
}

type DetachLoadBalancersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity. The value of this parameter is returned only if you set the Async parameter to true. You can call the describescalingactivities operation to query all scaling activity IDs and use the scaling activity IDs to check the status of the scaling activities.
	//
	// example:
	//
	// asa-bp140qd7mak8k63f****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DetachLoadBalancersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *DetachLoadBalancersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachLoadBalancersResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *DetachLoadBalancersResponseBody) SetRequestId(v string) *DetachLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachLoadBalancersResponseBody) SetScalingActivityId(v string) *DetachLoadBalancersResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *DetachLoadBalancersResponseBody) Validate() error {
	return dara.Validate(s)
}
