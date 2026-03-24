// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChargeResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeCycle(v string) *DescribeChargeResultRequest
	GetChargeCycle() *string
	SetChargeModules(v []*DescribeChargeResultRequestChargeModules) *DescribeChargeResultRequest
	GetChargeModules() []*DescribeChargeResultRequestChargeModules
	SetPayType(v string) *DescribeChargeResultRequest
	GetPayType() *string
	SetRegionId(v string) *DescribeChargeResultRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeChargeResultRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeChargeResultRequest struct {
	// The billing cycle for the WAF instance. Valid values:
	//
	// - **Year**: yearly billing cycle.
	//
	// - **Month**: monthly billing cycle.
	//
	// - **Day**: daily billing cycle.
	//
	// example:
	//
	// Day
	ChargeCycle *string `json:"ChargeCycle,omitempty" xml:"ChargeCycle,omitempty"`
	// The billing modules to calculate.
	//
	// This parameter is required.
	ChargeModules []*DescribeChargeResultRequestChargeModules `json:"ChargeModules,omitempty" xml:"ChargeModules,omitempty" type:"Repeated"`
	// The billing method of the WAF instance. Valid value:
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeChargeResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChargeResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeChargeResultRequest) GetChargeCycle() *string {
	return s.ChargeCycle
}

func (s *DescribeChargeResultRequest) GetChargeModules() []*DescribeChargeResultRequestChargeModules {
	return s.ChargeModules
}

func (s *DescribeChargeResultRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeChargeResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeChargeResultRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeChargeResultRequest) SetChargeCycle(v string) *DescribeChargeResultRequest {
	s.ChargeCycle = &v
	return s
}

func (s *DescribeChargeResultRequest) SetChargeModules(v []*DescribeChargeResultRequestChargeModules) *DescribeChargeResultRequest {
	s.ChargeModules = v
	return s
}

func (s *DescribeChargeResultRequest) SetPayType(v string) *DescribeChargeResultRequest {
	s.PayType = &v
	return s
}

func (s *DescribeChargeResultRequest) SetRegionId(v string) *DescribeChargeResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeChargeResultRequest) SetResourceManagerResourceGroupId(v string) *DescribeChargeResultRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeChargeResultRequest) Validate() error {
	if s.ChargeModules != nil {
		for _, item := range s.ChargeModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeChargeResultRequestChargeModules struct {
	// The ID of the billing module.
	//
	// example:
	//
	// domainCount
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The usage amount of the billing module.
	//
	// example:
	//
	// 10
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeChargeResultRequestChargeModules) String() string {
	return dara.Prettify(s)
}

func (s DescribeChargeResultRequestChargeModules) GoString() string {
	return s.String()
}

func (s *DescribeChargeResultRequestChargeModules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribeChargeResultRequestChargeModules) GetUsage() *int64 {
	return s.Usage
}

func (s *DescribeChargeResultRequestChargeModules) SetModuleCode(v string) *DescribeChargeResultRequestChargeModules {
	s.ModuleCode = &v
	return s
}

func (s *DescribeChargeResultRequestChargeModules) SetUsage(v int64) *DescribeChargeResultRequestChargeModules {
	s.Usage = &v
	return s
}

func (s *DescribeChargeResultRequestChargeModules) Validate() error {
	return dara.Validate(s)
}
