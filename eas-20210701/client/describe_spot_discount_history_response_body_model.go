// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpotDiscountHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSpotDiscountHistoryResponseBody
	GetRequestId() *string
	SetSpotDiscounts(v []*DescribeSpotDiscountHistoryResponseBodySpotDiscounts) *DescribeSpotDiscountHistoryResponseBody
	GetSpotDiscounts() []*DescribeSpotDiscountHistoryResponseBodySpotDiscounts
}

type DescribeSpotDiscountHistoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The discount for the preemptible instance.
	SpotDiscounts []*DescribeSpotDiscountHistoryResponseBodySpotDiscounts `json:"SpotDiscounts,omitempty" xml:"SpotDiscounts,omitempty" type:"Repeated"`
}

func (s DescribeSpotDiscountHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotDiscountHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSpotDiscountHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSpotDiscountHistoryResponseBody) GetSpotDiscounts() []*DescribeSpotDiscountHistoryResponseBodySpotDiscounts {
	return s.SpotDiscounts
}

func (s *DescribeSpotDiscountHistoryResponseBody) SetRequestId(v string) *DescribeSpotDiscountHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSpotDiscountHistoryResponseBody) SetSpotDiscounts(v []*DescribeSpotDiscountHistoryResponseBodySpotDiscounts) *DescribeSpotDiscountHistoryResponseBody {
	s.SpotDiscounts = v
	return s
}

func (s *DescribeSpotDiscountHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSpotDiscountHistoryResponseBodySpotDiscounts struct {
	// The type of the ECS instance.
	//
	// example:
	//
	// ecs.c7.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The discount for the preemptible instance. For example, 0.1 represents a 90% discount.
	//
	// example:
	//
	// 0.1
	SpotDiscount *string `json:"SpotDiscount,omitempty" xml:"SpotDiscount,omitempty"`
	// The time when the discount is available. The time must be in UTC.
	//
	// example:
	//
	// 2024-04-10T10:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSpotDiscountHistoryResponseBodySpotDiscounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotDiscountHistoryResponseBodySpotDiscounts) GoString() string {
	return s.String()
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) GetSpotDiscount() *string {
	return s.SpotDiscount
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) SetInstanceType(v string) *DescribeSpotDiscountHistoryResponseBodySpotDiscounts {
	s.InstanceType = &v
	return s
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) SetSpotDiscount(v string) *DescribeSpotDiscountHistoryResponseBodySpotDiscounts {
	s.SpotDiscount = &v
	return s
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) SetTimestamp(v string) *DescribeSpotDiscountHistoryResponseBodySpotDiscounts {
	s.Timestamp = &v
	return s
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) SetZoneId(v string) *DescribeSpotDiscountHistoryResponseBodySpotDiscounts {
	s.ZoneId = &v
	return s
}

func (s *DescribeSpotDiscountHistoryResponseBodySpotDiscounts) Validate() error {
	return dara.Validate(s)
}
