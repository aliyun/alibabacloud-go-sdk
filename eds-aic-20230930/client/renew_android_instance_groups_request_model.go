// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAndroidInstanceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewAndroidInstanceGroupsRequest
	GetAutoPay() *bool
	SetInstanceGroupIds(v []*string) *RenewAndroidInstanceGroupsRequest
	GetInstanceGroupIds() []*string
	SetPaidCallBackUrl(v string) *RenewAndroidInstanceGroupsRequest
	GetPaidCallBackUrl() *string
	SetPeriod(v int32) *RenewAndroidInstanceGroupsRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewAndroidInstanceGroupsRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *RenewAndroidInstanceGroupsRequest
	GetPromotionId() *string
}

type RenewAndroidInstanceGroupsRequest struct {
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The list of instance group IDs.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The redirect URL to the payment success page.
	//
	// example:
	//
	// wya.wuying.aliyun.com/instanceGroup
	PaidCallBackUrl *string `json:"PaidCallBackUrl,omitempty" xml:"PaidCallBackUrl,omitempty"`
	// Specifies the renewal duration. The unit is specified by the PeriodUnit parameter.
	//
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration, which is the unit of the Period parameter. Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s RenewAndroidInstanceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewAndroidInstanceGroupsRequest) GoString() string {
	return s.String()
}

func (s *RenewAndroidInstanceGroupsRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewAndroidInstanceGroupsRequest) GetInstanceGroupIds() []*string {
	return s.InstanceGroupIds
}

func (s *RenewAndroidInstanceGroupsRequest) GetPaidCallBackUrl() *string {
	return s.PaidCallBackUrl
}

func (s *RenewAndroidInstanceGroupsRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewAndroidInstanceGroupsRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewAndroidInstanceGroupsRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *RenewAndroidInstanceGroupsRequest) SetAutoPay(v bool) *RenewAndroidInstanceGroupsRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) SetInstanceGroupIds(v []*string) *RenewAndroidInstanceGroupsRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) SetPaidCallBackUrl(v string) *RenewAndroidInstanceGroupsRequest {
	s.PaidCallBackUrl = &v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) SetPeriod(v int32) *RenewAndroidInstanceGroupsRequest {
	s.Period = &v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) SetPeriodUnit(v string) *RenewAndroidInstanceGroupsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) SetPromotionId(v string) *RenewAndroidInstanceGroupsRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
