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
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The subscription duration of the instance.
	//
	// - If Period is set to **Year**, valid values are **1*	- to **5**.
	//
	// - If Period is set to **Month**, valid values are **1*	- to **60**.
	//
	// > This parameter is valid and required only when ChargeType is set to **Prepaid**.
	//
	// example:
	//
	// 5
	BuyCount *string `json:"BuyCount,omitempty" xml:"BuyCount,omitempty"`
	// The billing method after conversion. Valid values:
	//
	// - **PrePaid**: subscription.
	//
	// - **PostPaid**: pay-as-you-go.
	//
	// <props="china">
	//
	// - **sync_serverless**: pay-as-you-go Serverless..
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the data synchronization or change tracking task. You can call [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) to query the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// o4nh3g7jg56****
	DtsJobId      *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The maximum number of DUs for the Serverless instance. Valid values: 2, 4, 8, and 16.
	//
	// <props="intl">
	//
	// > This feature is currently not supported. Do not specify this parameter.
	//
	// <props="china">
	//
	// > This parameter is valid and required only when ChargeType is set to **sync_serverless**..
	//
	// example:
	//
	// 16
	MaxDu *int32 `json:"MaxDu,omitempty" xml:"MaxDu,omitempty"`
	// The minimum number of DTS Units (DUs) for the Serverless instance. Valid values: 1, 2, 4, 8, and 16.
	//
	// <props="intl">
	//
	// > This feature is currently not supported. Do not specify this parameter.
	//
	// <props="china">
	//
	// > This parameter is valid and required only when ChargeType is set to **sync_serverless**..
	//
	// example:
	//
	// 1
	MinDu *int32 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	// The billing method of the subscription instance. Valid values:
	//
	// - **Year**: annual subscription.
	//
	// - **Month**: monthly subscription.
	//
	// > This parameter is valid and required only when ChargeType is set to **PrePaid*	- (subscription).
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
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
