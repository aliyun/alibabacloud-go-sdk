// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRatePlanInstanceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeRatePlanInstanceStatusResponseBody
	GetInstanceId() *string
	SetInstanceStatus(v string) *DescribeRatePlanInstanceStatusResponseBody
	GetInstanceStatus() *string
	SetRequestId(v string) *DescribeRatePlanInstanceStatusResponseBody
	GetRequestId() *string
}

type DescribeRatePlanInstanceStatusResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// xcdn-91fknmb80f0g***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance status. Valid values:
	//
	// - running: Running.
	//
	// - renewing: Being renewed.
	//
	// - upgrading: Being upgraded.
	//
	// - releasePrepaidService: Released due to subscription expiration.
	//
	// - creating: Being created.
	//
	// - downgrading: Being downgraded.
	//
	// - ceasePrepaidService: Suspended due to subscription expiration.
	//
	// example:
	//
	// running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 60423A7F-A83D-1E24-B80E-86DD25790759
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRatePlanInstanceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanInstanceStatusResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRatePlanInstanceStatusResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeRatePlanInstanceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRatePlanInstanceStatusResponseBody) SetInstanceId(v string) *DescribeRatePlanInstanceStatusResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeRatePlanInstanceStatusResponseBody) SetInstanceStatus(v string) *DescribeRatePlanInstanceStatusResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeRatePlanInstanceStatusResponseBody) SetRequestId(v string) *DescribeRatePlanInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRatePlanInstanceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
