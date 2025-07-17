// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingActivityDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v string) *DescribeScalingActivityDetailResponseBody
	GetDetail() *string
	SetRequestId(v string) *DescribeScalingActivityDetailResponseBody
	GetRequestId() *string
	SetScalingActivityId(v string) *DescribeScalingActivityDetailResponseBody
	GetScalingActivityId() *string
}

type DescribeScalingActivityDetailResponseBody struct {
	// The details of the scaling activity. The result of a scaling activity is either successful or failed. If the scaling activity is rejected, no scaling activity details are returned.
	//
	// example:
	//
	// new ECS instances \\"i-bp16t2cgmiiymeqv****\\" are created.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B13527BF-1FBD-4334-A512-20F5E9D3FB4D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling activity.
	//
	// example:
	//
	// asa-bp1c9djwrgxjyk31****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DescribeScalingActivityDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivityDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityDetailResponseBody) GetDetail() *string {
	return s.Detail
}

func (s *DescribeScalingActivityDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScalingActivityDetailResponseBody) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *DescribeScalingActivityDetailResponseBody) SetDetail(v string) *DescribeScalingActivityDetailResponseBody {
	s.Detail = &v
	return s
}

func (s *DescribeScalingActivityDetailResponseBody) SetRequestId(v string) *DescribeScalingActivityDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingActivityDetailResponseBody) SetScalingActivityId(v string) *DescribeScalingActivityDetailResponseBody {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeScalingActivityDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
