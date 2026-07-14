// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateNetworkPackageRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateNetworkPackageRequest
	GetAutoRenew() *bool
	SetBandwidth(v int32) *CreateNetworkPackageRequest
	GetBandwidth() *int32
	SetChannelCookie(v string) *CreateNetworkPackageRequest
	GetChannelCookie() *string
	SetInternetChargeType(v string) *CreateNetworkPackageRequest
	GetInternetChargeType() *string
	SetOfficeSiteId(v string) *CreateNetworkPackageRequest
	GetOfficeSiteId() *string
	SetPayType(v string) *CreateNetworkPackageRequest
	GetPayType() *string
	SetPeriod(v int32) *CreateNetworkPackageRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateNetworkPackageRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *CreateNetworkPackageRequest
	GetPromotionId() *string
	SetRegionId(v string) *CreateNetworkPackageRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *CreateNetworkPackageRequest
	GetResellerOwnerUid() *int64
	SetTag(v []*CreateNetworkPackageRequestTag) *CreateNetworkPackageRequest
	GetTag() []*CreateNetworkPackageRequestTag
}

type CreateNetworkPackageRequest struct {
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The bandwidth of the premium bandwidth plan. Unit: Mbit/s.
	//
	// - If the premium bandwidth plan uses the subscription billing method, the valid values are 2 to 1000.
	//
	// - If the premium bandwidth plan uses the pay-as-you-go billing method and the billing type is pay-by-data-transfer (PayByTraffic), the valid values are 2 to 200.
	//
	// - If the premium bandwidth plan uses the pay-as-you-go billing method and the billing type is pay-by-bandwidth (PayByBandwidth), the valid values are 2 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Bandwidth     *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ChannelCookie *string `json:"ChannelCookie,omitempty" xml:"ChannelCookie,omitempty"`
	// The billable methods of the premium bandwidth plan.
	//
	// - If the parameter `PayType` is set to `PrePaid`, valid values:
	//
	//     - PayByBandwidth: billing by fixed bandwidth.
	//
	// - If the parameter `PayType` is set to `PostPaid`, valid values:
	//
	//     - PayByTraffic: billing by data transfer.
	//
	//     - PayByBandwidth: billing by fixed bandwidth.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription duration of the premium bandwidth plan. This parameter takes effect and is required only when PayType is set to PrePaid. Valid values are determined by the PeriodUnit parameter.
	//
	// - If PeriodUnit is set to Week, the valid value is 1.
	//
	// - If PeriodUnit is set to Month, valid values are 1, 2, 3, and 6.
	//
	// - If PeriodUnit is set to Year, valid values are 1, 2, and 3.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration for the premium bandwidth plan. This parameter takes effect and is required only when PayType is set to PrePaid.
	//
	// example:
	//
	// Week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 23141
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64                            `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	Tag              []*CreateNetworkPackageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateNetworkPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateNetworkPackageRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateNetworkPackageRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateNetworkPackageRequest) GetChannelCookie() *string {
	return s.ChannelCookie
}

func (s *CreateNetworkPackageRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateNetworkPackageRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateNetworkPackageRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateNetworkPackageRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateNetworkPackageRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateNetworkPackageRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateNetworkPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNetworkPackageRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *CreateNetworkPackageRequest) GetTag() []*CreateNetworkPackageRequestTag {
	return s.Tag
}

func (s *CreateNetworkPackageRequest) SetAutoPay(v bool) *CreateNetworkPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetAutoRenew(v bool) *CreateNetworkPackageRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetBandwidth(v int32) *CreateNetworkPackageRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetChannelCookie(v string) *CreateNetworkPackageRequest {
	s.ChannelCookie = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetInternetChargeType(v string) *CreateNetworkPackageRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetOfficeSiteId(v string) *CreateNetworkPackageRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPayType(v string) *CreateNetworkPackageRequest {
	s.PayType = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPeriod(v int32) *CreateNetworkPackageRequest {
	s.Period = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPeriodUnit(v string) *CreateNetworkPackageRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetPromotionId(v string) *CreateNetworkPackageRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetRegionId(v string) *CreateNetworkPackageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetResellerOwnerUid(v int64) *CreateNetworkPackageRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *CreateNetworkPackageRequest) SetTag(v []*CreateNetworkPackageRequestTag) *CreateNetworkPackageRequest {
	s.Tag = v
	return s
}

func (s *CreateNetworkPackageRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNetworkPackageRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNetworkPackageRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkPackageRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateNetworkPackageRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateNetworkPackageRequestTag) SetKey(v string) *CreateNetworkPackageRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNetworkPackageRequestTag) SetValue(v string) *CreateNetworkPackageRequestTag {
	s.Value = &v
	return s
}

func (s *CreateNetworkPackageRequestTag) Validate() error {
	return dara.Validate(s)
}
