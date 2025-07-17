// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachLoadBalancersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachLoadBalancersResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *AttachLoadBalancersResponseBody
	GetScalingActivityId() *string
}

type AttachLoadBalancersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity.
	//
	// This parameter is returned only if you set `Async` to `true`. You can call the DescribeScalingActivities operation to query the scaling activity IDs and status.
	//
	// example:
	//
	// asa-bp140qd7mak8k63f****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s AttachLoadBalancersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *AttachLoadBalancersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachLoadBalancersResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *AttachLoadBalancersResponseBody) SetRequestId(v string) *AttachLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachLoadBalancersResponseBody) SetScalingActivityId(v string) *AttachLoadBalancersResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *AttachLoadBalancersResponseBody) Validate() error {
	return dara.Validate(s)
}
