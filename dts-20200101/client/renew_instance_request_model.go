// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuyCount(v string) *RenewInstanceRequest
	GetBuyCount() *string
	SetChargeType(v string) *RenewInstanceRequest
	GetChargeType() *string
	SetDtsJobId(v string) *RenewInstanceRequest
	GetDtsJobId() *string
	SetPeriod(v string) *RenewInstanceRequest
	GetPeriod() *string
	SetRegionId(v string) *RenewInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *RenewInstanceRequest
	GetResourceGroupId() *string
}

type RenewInstanceRequest struct {
	// The subscription duration of the DTS instance after renewal. Default value: 1.
	//
	// 	- If **Period*	- is set to **Year**, the valid values are **1 to 5**.
	//
	// 	- If **Period*	- is set to **Month**, the valid values are **1 to 60**.
	//
	// example:
	//
	// 1
	BuyCount *string `json:"BuyCount,omitempty" xml:"BuyCount,omitempty"`
	// The billing method of the DTS instance. Set the value to **PREPAY**, which specifies the subscription billing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the data synchronization or change tracking task. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// qi0r643lc31****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The billing cycle of the DTS instance after renewal. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month*	- (default)
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the DTS instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) GetBuyCount() *string {
	return s.BuyCount
}

func (s *RenewInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *RenewInstanceRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *RenewInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *RenewInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RenewInstanceRequest) SetBuyCount(v string) *RenewInstanceRequest {
	s.BuyCount = &v
	return s
}

func (s *RenewInstanceRequest) SetChargeType(v string) *RenewInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *RenewInstanceRequest) SetDtsJobId(v string) *RenewInstanceRequest {
	s.DtsJobId = &v
	return s
}

func (s *RenewInstanceRequest) SetPeriod(v string) *RenewInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewInstanceRequest) SetRegionId(v string) *RenewInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RenewInstanceRequest) SetResourceGroupId(v string) *RenewInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RenewInstanceRequest) Validate() error {
	return dara.Validate(s)
}
