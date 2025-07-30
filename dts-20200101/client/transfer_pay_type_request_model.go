// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferPayTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *TransferPayTypeRequest
	GetAutoPay() *bool
	SetBuyCount(v string) *TransferPayTypeRequest
	GetBuyCount() *string
	SetChargeType(v string) *TransferPayTypeRequest
	GetChargeType() *string
	SetDtsJobId(v string) *TransferPayTypeRequest
	GetDtsJobId() *string
	SetInstanceClass(v string) *TransferPayTypeRequest
	GetInstanceClass() *string
	SetMaxDu(v int32) *TransferPayTypeRequest
	GetMaxDu() *int32
	SetMinDu(v int32) *TransferPayTypeRequest
	GetMinDu() *int32
	SetPeriod(v string) *TransferPayTypeRequest
	GetPeriod() *string
	SetRegionId(v string) *TransferPayTypeRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *TransferPayTypeRequest
	GetResourceGroupId() *string
}

type TransferPayTypeRequest struct {
	// Specifies whether to automatically renew the DTS instance when it expires. Valid values:
	//
	// 	- **false**: does not automatically renew the DTS instance when it expires. This is the default value.
	//
	// 	- **true**: automatically renews the DTS instance when it expires.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The subscription length.
	//
	// 	- If the **Period*	- parameter is set to **Year**, the value range is **1*	- to **5**.
	//
	// 	- If the **Period*	- parameter is set to **Month**, the value range is **1*	- to **60**.
	//
	// >  You must specify this parameter only if you set the **ChargeType*	- parameter to **PrePaid**.
	//
	// example:
	//
	// 5
	BuyCount *string `json:"BuyCount,omitempty" xml:"BuyCount,omitempty"`
	// The new billing method. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid**: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the data synchronization or change tracking task. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// o4nh3g7jg56****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The new instance class of the DTS instance. You can call the [DescribeDtsJobDetail](https://help.aliyun.com/document_detail/208925.html) operation to query the original instance class of the DTS instance.
	//
	// 	- DTS supports the following instance classes for a data migration instance: **xxlarge**, **xlarge**, **large**, **medium**, and **small**.
	//
	// 	- DTS supports the following instance classes for a data synchronization instance: **large**, **medium**, **small**, and **micro**.
	//
	// > For more information about the test performance of each instance class, see [Specifications of data migration instances](https://help.aliyun.com/document_detail/26606.html) and [Specifications of data synchronization channels](https://help.aliyun.com/document_detail/26605.html).
	//
	// example:
	//
	// small
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The maximum number of DUs in a serverless instance. Valid values: 2, 4, 8, and 16.
	//
	// >  This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// 16
	MaxDu *int32 `json:"MaxDu,omitempty" xml:"MaxDu,omitempty"`
	// The minimum number of DTS Units (DUs) in a serverless instance. Valid values: 1, 2, 4, 8, and 16.
	//
	// >  This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// 1
	MinDu *int32 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	// The billing cycle of the subscription instance. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month*	- (default value)
	//
	// >  You must specify this parameter only if you set the **ChargeType*	- parameter to **PrePaid**.
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the region where the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s TransferPayTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferPayTypeRequest) GoString() string {
	return s.String()
}

func (s *TransferPayTypeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *TransferPayTypeRequest) GetBuyCount() *string {
	return s.BuyCount
}

func (s *TransferPayTypeRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *TransferPayTypeRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *TransferPayTypeRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *TransferPayTypeRequest) GetMaxDu() *int32 {
	return s.MaxDu
}

func (s *TransferPayTypeRequest) GetMinDu() *int32 {
	return s.MinDu
}

func (s *TransferPayTypeRequest) GetPeriod() *string {
	return s.Period
}

func (s *TransferPayTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TransferPayTypeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *TransferPayTypeRequest) SetAutoPay(v bool) *TransferPayTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *TransferPayTypeRequest) SetBuyCount(v string) *TransferPayTypeRequest {
	s.BuyCount = &v
	return s
}

func (s *TransferPayTypeRequest) SetChargeType(v string) *TransferPayTypeRequest {
	s.ChargeType = &v
	return s
}

func (s *TransferPayTypeRequest) SetDtsJobId(v string) *TransferPayTypeRequest {
	s.DtsJobId = &v
	return s
}

func (s *TransferPayTypeRequest) SetInstanceClass(v string) *TransferPayTypeRequest {
	s.InstanceClass = &v
	return s
}

func (s *TransferPayTypeRequest) SetMaxDu(v int32) *TransferPayTypeRequest {
	s.MaxDu = &v
	return s
}

func (s *TransferPayTypeRequest) SetMinDu(v int32) *TransferPayTypeRequest {
	s.MinDu = &v
	return s
}

func (s *TransferPayTypeRequest) SetPeriod(v string) *TransferPayTypeRequest {
	s.Period = &v
	return s
}

func (s *TransferPayTypeRequest) SetRegionId(v string) *TransferPayTypeRequest {
	s.RegionId = &v
	return s
}

func (s *TransferPayTypeRequest) SetResourceGroupId(v string) *TransferPayTypeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TransferPayTypeRequest) Validate() error {
	return dara.Validate(s)
}
