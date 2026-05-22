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
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
