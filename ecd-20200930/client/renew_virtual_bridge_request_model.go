// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewVirtualBridgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewVirtualBridgeRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RenewVirtualBridgeRequest
	GetAutoRenew() *bool
	SetBridgeId(v string) *RenewVirtualBridgeRequest
	GetBridgeId() *string
	SetPaidCallBackUrl(v string) *RenewVirtualBridgeRequest
	GetPaidCallBackUrl() *string
	SetPeriod(v int32) *RenewVirtualBridgeRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewVirtualBridgeRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *RenewVirtualBridgeRequest
	GetPromotionId() *string
	SetRegionId(v string) *RenewVirtualBridgeRequest
	GetRegionId() *string
}

type RenewVirtualBridgeRequest struct {
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
	// The virtual bridge ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vb-shfisahfihs***
	BridgeId *string `json:"BridgeId,omitempty" xml:"BridgeId,omitempty"`
	// The payment callback URL.
	//
	// example:
	//
	// https://edu.wuying.aliyun.com/edu/school-manage
	PaidCallBackUrl *string `json:"PaidCallBackUrl,omitempty" xml:"PaidCallBackUrl,omitempty"`
	// The renewal duration. The valid values of this parameter are determined by the value of the `PeriodUnit` parameter.
	//
	// - If `PeriodUnit` is set to `Month`, valid values are 1, 2, 3, and 6.
	//
	// - If `PeriodUnit` is set to `Year`, valid values are 1, 2, and 3.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration, which is the unit of the `Period` parameter.
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 50003836003****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RenewVirtualBridgeRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewVirtualBridgeRequest) GoString() string {
	return s.String()
}

func (s *RenewVirtualBridgeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewVirtualBridgeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewVirtualBridgeRequest) GetBridgeId() *string {
	return s.BridgeId
}

func (s *RenewVirtualBridgeRequest) GetPaidCallBackUrl() *string {
	return s.PaidCallBackUrl
}

func (s *RenewVirtualBridgeRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewVirtualBridgeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewVirtualBridgeRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *RenewVirtualBridgeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewVirtualBridgeRequest) SetAutoPay(v bool) *RenewVirtualBridgeRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewVirtualBridgeRequest) SetAutoRenew(v bool) *RenewVirtualBridgeRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewVirtualBridgeRequest) SetBridgeId(v string) *RenewVirtualBridgeRequest {
	s.BridgeId = &v
	return s
}

func (s *RenewVirtualBridgeRequest) SetPaidCallBackUrl(v string) *RenewVirtualBridgeRequest {
	s.PaidCallBackUrl = &v
	return s
}

func (s *RenewVirtualBridgeRequest) SetPeriod(v int32) *RenewVirtualBridgeRequest {
	s.Period = &v
	return s
}

func (s *RenewVirtualBridgeRequest) SetPeriodUnit(v string) *RenewVirtualBridgeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewVirtualBridgeRequest) SetPromotionId(v string) *RenewVirtualBridgeRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewVirtualBridgeRequest) SetRegionId(v string) *RenewVirtualBridgeRequest {
	s.RegionId = &v
	return s
}

func (s *RenewVirtualBridgeRequest) Validate() error {
	return dara.Validate(s)
}
