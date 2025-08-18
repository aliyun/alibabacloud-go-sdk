// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *RenewAppInstanceGroupRequest
	GetAppInstanceGroupId() *string
	SetAutoPay(v bool) *RenewAppInstanceGroupRequest
	GetAutoPay() *bool
	SetPeriod(v int32) *RenewAppInstanceGroupRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewAppInstanceGroupRequest
	GetPeriodUnit() *string
	SetProductType(v string) *RenewAppInstanceGroupRequest
	GetProductType() *string
	SetPromotionId(v string) *RenewAppInstanceGroupRequest
	GetPromotionId() *string
	SetRenewAmount(v int32) *RenewAppInstanceGroupRequest
	GetRenewAmount() *int32
	SetRenewMode(v string) *RenewAppInstanceGroupRequest
	GetRenewMode() *string
	SetRenewNodes(v []*string) *RenewAppInstanceGroupRequest
	GetRenewNodes() []*string
}

type RenewAppInstanceGroupRequest struct {
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
	// Valid values:
	//
	// 	- true
	//
	// 	- false: manual payment. This is the default value.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The subscription duration of resources. This parameter must be configured together with `PeriodUnit`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. This parameter must be configured together with `Period`. The following items describe valid values for the combinations of `Period` and `PeriodUnit`:
	//
	// 	- 1 Week
	//
	// 	- 1 Month
	//
	// 	- 2 Month
	//
	// 	- 3 Month
	//
	// 	- 6 Month
	//
	// 	- 1 Year
	//
	// 	- 2 Year
	//
	// 	- 3 Year
	//
	// >  The value of this parameter is case-insensitive. For example, `Week` is valid and `week` is invalid. If you specify a value combination other than the preceding combinations, such as `2 Week`, the operation can still be called. However, an error occurs when you place the order.
	//
	// This parameter is required.
	//
	// example:
	//
	// Week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The promotion ID. You can call the [GetResourcePrice](https://help.aliyun.com/document_detail/428503.html) operation to obtain the ID.
	//
	// example:
	//
	// 17440009****
	PromotionId *string   `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RenewAmount *int32    `json:"RenewAmount,omitempty" xml:"RenewAmount,omitempty"`
	RenewMode   *string   `json:"RenewMode,omitempty" xml:"RenewMode,omitempty"`
	RenewNodes  []*string `json:"RenewNodes,omitempty" xml:"RenewNodes,omitempty" type:"Repeated"`
}

func (s RenewAppInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceGroupRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *RenewAppInstanceGroupRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewAppInstanceGroupRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewAppInstanceGroupRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewAppInstanceGroupRequest) GetProductType() *string {
	return s.ProductType
}

func (s *RenewAppInstanceGroupRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *RenewAppInstanceGroupRequest) GetRenewAmount() *int32 {
	return s.RenewAmount
}

func (s *RenewAppInstanceGroupRequest) GetRenewMode() *string {
	return s.RenewMode
}

func (s *RenewAppInstanceGroupRequest) GetRenewNodes() []*string {
	return s.RenewNodes
}

func (s *RenewAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *RenewAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetAutoPay(v bool) *RenewAppInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetPeriod(v int32) *RenewAppInstanceGroupRequest {
	s.Period = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetPeriodUnit(v string) *RenewAppInstanceGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetProductType(v string) *RenewAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetPromotionId(v string) *RenewAppInstanceGroupRequest {
	s.PromotionId = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetRenewAmount(v int32) *RenewAppInstanceGroupRequest {
	s.RenewAmount = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetRenewMode(v string) *RenewAppInstanceGroupRequest {
	s.RenewMode = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetRenewNodes(v []*string) *RenewAppInstanceGroupRequest {
	s.RenewNodes = v
	return s
}

func (s *RenewAppInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}
