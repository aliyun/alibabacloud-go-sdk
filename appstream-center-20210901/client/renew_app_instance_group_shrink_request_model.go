// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppInstanceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *RenewAppInstanceGroupShrinkRequest
	GetAppInstanceGroupId() *string
	SetAutoPay(v bool) *RenewAppInstanceGroupShrinkRequest
	GetAutoPay() *bool
	SetPeriod(v int32) *RenewAppInstanceGroupShrinkRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewAppInstanceGroupShrinkRequest
	GetPeriodUnit() *string
	SetProductType(v string) *RenewAppInstanceGroupShrinkRequest
	GetProductType() *string
	SetPromotionId(v string) *RenewAppInstanceGroupShrinkRequest
	GetPromotionId() *string
	SetRenewAmount(v int32) *RenewAppInstanceGroupShrinkRequest
	GetRenewAmount() *int32
	SetRenewMode(v string) *RenewAppInstanceGroupShrinkRequest
	GetRenewMode() *string
	SetRenewNodesShrink(v string) *RenewAppInstanceGroupShrinkRequest
	GetRenewNodesShrink() *string
}

type RenewAppInstanceGroupShrinkRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The numeric part of the resource purchase duration. This parameter is used together with PeriodUnit to specify the complete purchase duration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit part of the resource purchase duration. This parameter is used together with Period to specify the complete purchase duration. Valid combinations of Period and PeriodUnit:
	//
	// - 1 Week (1 week)
	//
	// - 1 Month (1 month)
	//
	// - 2 Month (2 months)
	//
	// - 3 Month (3 months)
	//
	// - 6 Month (6 months)
	//
	// - 1 Year (1 year)
	//
	// - 2 Year (2 years)
	//
	// - 3 Year (3 years)
	//
	// > This parameter is case-sensitive. For example, `Week` is valid, but `week` is invalid. If the request parameters do not match the combinations listed above, such as `2 Week`, the call to this operation succeeds, but an error occurs during the order placement phase.
	//
	// This parameter is required.
	//
	// example:
	//
	// Week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The product type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The promotion ID. You can obtain this value by calling the [GetResourcePrice](https://help.aliyun.com/document_detail/428503.html) operation.
	//
	// example:
	//
	// 17440009****
	PromotionId      *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RenewAmount      *int32  `json:"RenewAmount,omitempty" xml:"RenewAmount,omitempty"`
	RenewMode        *string `json:"RenewMode,omitempty" xml:"RenewMode,omitempty"`
	RenewNodesShrink *string `json:"RenewNodes,omitempty" xml:"RenewNodes,omitempty"`
}

func (s RenewAppInstanceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewAppInstanceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceGroupShrinkRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *RenewAppInstanceGroupShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewAppInstanceGroupShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewAppInstanceGroupShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewAppInstanceGroupShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *RenewAppInstanceGroupShrinkRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *RenewAppInstanceGroupShrinkRequest) GetRenewAmount() *int32 {
	return s.RenewAmount
}

func (s *RenewAppInstanceGroupShrinkRequest) GetRenewMode() *string {
	return s.RenewMode
}

func (s *RenewAppInstanceGroupShrinkRequest) GetRenewNodesShrink() *string {
	return s.RenewNodesShrink
}

func (s *RenewAppInstanceGroupShrinkRequest) SetAppInstanceGroupId(v string) *RenewAppInstanceGroupShrinkRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *RenewAppInstanceGroupShrinkRequest) SetAutoPay(v bool) *RenewAppInstanceGroupShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewAppInstanceGroupShrinkRequest) SetPeriod(v int32) *RenewAppInstanceGroupShrinkRequest {
	s.Period = &v
	return s
}

func (s *RenewAppInstanceGroupShrinkRequest) SetPeriodUnit(v string) *RenewAppInstanceGroupShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewAppInstanceGroupShrinkRequest) SetProductType(v string) *RenewAppInstanceGroupShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *RenewAppInstanceGroupShrinkRequest) SetPromotionId(v string) *RenewAppInstanceGroupShrinkRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewAppInstanceGroupShrinkRequest) SetRenewAmount(v int32) *RenewAppInstanceGroupShrinkRequest {
	s.RenewAmount = &v
	return s
}

func (s *RenewAppInstanceGroupShrinkRequest) SetRenewMode(v string) *RenewAppInstanceGroupShrinkRequest {
	s.RenewMode = &v
	return s
}

func (s *RenewAppInstanceGroupShrinkRequest) SetRenewNodesShrink(v string) *RenewAppInstanceGroupShrinkRequest {
	s.RenewNodesShrink = &v
	return s
}

func (s *RenewAppInstanceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
