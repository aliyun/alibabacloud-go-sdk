// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChargeModuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayType(v string) *DescribeChargeModuleRequest
	GetPayType() *string
	SetRegionId(v string) *DescribeChargeModuleRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeChargeModuleRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeChargeModuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeChargeModuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChargeModuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeChargeModuleRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeChargeModuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeChargeModuleRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeChargeModuleRequest) SetPayType(v string) *DescribeChargeModuleRequest {
	s.PayType = &v
	return s
}

func (s *DescribeChargeModuleRequest) SetRegionId(v string) *DescribeChargeModuleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeChargeModuleRequest) SetResourceManagerResourceGroupId(v string) *DescribeChargeModuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeChargeModuleRequest) Validate() error {
	return dara.Validate(s)
}
