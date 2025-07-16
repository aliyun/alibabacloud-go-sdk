// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpotDiscountHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceType(v string) *DescribeSpotDiscountHistoryRequest
	GetInstanceType() *string
	SetIsProtect(v bool) *DescribeSpotDiscountHistoryRequest
	GetIsProtect() *bool
}

type DescribeSpotDiscountHistoryRequest struct {
	// The type of the Elastic Algorithm Service (EAS) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.c6.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether the preemptible instance has a protection period. During the 1-hour protection period of the preemptible instance, the preemptible instance will not be released.
	//
	// example:
	//
	// false
	IsProtect *bool `json:"IsProtect,omitempty" xml:"IsProtect,omitempty"`
}

func (s DescribeSpotDiscountHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotDiscountHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpotDiscountHistoryRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSpotDiscountHistoryRequest) GetIsProtect() *bool {
	return s.IsProtect
}

func (s *DescribeSpotDiscountHistoryRequest) SetInstanceType(v string) *DescribeSpotDiscountHistoryRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeSpotDiscountHistoryRequest) SetIsProtect(v bool) *DescribeSpotDiscountHistoryRequest {
	s.IsProtect = &v
	return s
}

func (s *DescribeSpotDiscountHistoryRequest) Validate() error {
	return dara.Validate(s)
}
