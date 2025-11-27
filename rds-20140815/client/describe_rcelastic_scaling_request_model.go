// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCElasticScalingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceChargeType(v string) *DescribeRCElasticScalingRequest
	GetInstanceChargeType() *string
	SetInstanceId(v string) *DescribeRCElasticScalingRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeRCElasticScalingRequest
	GetInstanceType() *string
	SetRegionId(v string) *DescribeRCElasticScalingRequest
	GetRegionId() *string
	SetSupportCase(v string) *DescribeRCElasticScalingRequest
	GetSupportCase() *string
}

type DescribeRCElasticScalingRequest struct {
	// example:
	//
	// Prepaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// example:
	//
	// rc-a0e466b01tif2pkrgg
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// mysql.x2.medium.9cm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// cn-hanghzou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// vnode
	SupportCase *string `json:"SupportCase,omitempty" xml:"SupportCase,omitempty"`
}

func (s DescribeRCElasticScalingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCElasticScalingRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCElasticScalingRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeRCElasticScalingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCElasticScalingRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRCElasticScalingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCElasticScalingRequest) GetSupportCase() *string {
	return s.SupportCase
}

func (s *DescribeRCElasticScalingRequest) SetInstanceChargeType(v string) *DescribeRCElasticScalingRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeRCElasticScalingRequest) SetInstanceId(v string) *DescribeRCElasticScalingRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCElasticScalingRequest) SetInstanceType(v string) *DescribeRCElasticScalingRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeRCElasticScalingRequest) SetRegionId(v string) *DescribeRCElasticScalingRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCElasticScalingRequest) SetSupportCase(v string) *DescribeRCElasticScalingRequest {
	s.SupportCase = &v
	return s
}

func (s *DescribeRCElasticScalingRequest) Validate() error {
	return dara.Validate(s)
}
