// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyDesktopChargeTypeRequest
	GetAutoPay() *bool
	SetChargeType(v string) *ModifyDesktopChargeTypeRequest
	GetChargeType() *string
	SetDesktopId(v []*string) *ModifyDesktopChargeTypeRequest
	GetDesktopId() []*string
	SetPeriod(v int32) *ModifyDesktopChargeTypeRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *ModifyDesktopChargeTypeRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *ModifyDesktopChargeTypeRequest
	GetPromotionId() *string
	SetRegionId(v string) *ModifyDesktopChargeTypeRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *ModifyDesktopChargeTypeRequest
	GetResellerOwnerUid() *int64
	SetUseDuration(v int32) *ModifyDesktopChargeTypeRequest
	GetUseDuration() *int32
}

type ModifyDesktopChargeTypeRequest struct {
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The new billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The IDs of the cloud desktops. You can specify 1 to 20 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The subscription duration. This parameter is required only when you set the `ChargeType` parameter to `PrePaid`. The unit of the duration is specified by the `PeriodUnit` parameter.
	//
	// - If you set the `PeriodUnit` parameter to `Week`, you can set this parameter only to 1.
	//
	// - If you set the `PeriodUnit` parameter to `Month`, you can set this parameter to 1, 2, 3, or 6.
	//
	// - If you set the `PeriodUnit` parameter to `Year`, you can set this parameter to 1, 2, 3, 4, or 5.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 50003836003****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// sample
	UseDuration *int32 `json:"UseDuration,omitempty" xml:"UseDuration,omitempty"`
}

func (s ModifyDesktopChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopChargeTypeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyDesktopChargeTypeRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ModifyDesktopChargeTypeRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *ModifyDesktopChargeTypeRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyDesktopChargeTypeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyDesktopChargeTypeRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *ModifyDesktopChargeTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDesktopChargeTypeRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *ModifyDesktopChargeTypeRequest) GetUseDuration() *int32 {
	return s.UseDuration
}

func (s *ModifyDesktopChargeTypeRequest) SetAutoPay(v bool) *ModifyDesktopChargeTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetChargeType(v string) *ModifyDesktopChargeTypeRequest {
	s.ChargeType = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetDesktopId(v []*string) *ModifyDesktopChargeTypeRequest {
	s.DesktopId = v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetPeriod(v int32) *ModifyDesktopChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetPeriodUnit(v string) *ModifyDesktopChargeTypeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetPromotionId(v string) *ModifyDesktopChargeTypeRequest {
	s.PromotionId = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetRegionId(v string) *ModifyDesktopChargeTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetResellerOwnerUid(v int64) *ModifyDesktopChargeTypeRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) SetUseDuration(v int32) *ModifyDesktopChargeTypeRequest {
	s.UseDuration = &v
	return s
}

func (s *ModifyDesktopChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
