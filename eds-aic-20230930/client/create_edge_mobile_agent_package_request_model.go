// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeMobileAgentPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateEdgeMobileAgentPackageRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateEdgeMobileAgentPackageRequest
	GetAutoRenew() *bool
	SetBizRegionId(v string) *CreateEdgeMobileAgentPackageRequest
	GetBizRegionId() *string
	SetClientToken(v string) *CreateEdgeMobileAgentPackageRequest
	GetClientToken() *string
	SetDeviceClass(v string) *CreateEdgeMobileAgentPackageRequest
	GetDeviceClass() *string
	SetPeriod(v int32) *CreateEdgeMobileAgentPackageRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateEdgeMobileAgentPackageRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *CreateEdgeMobileAgentPackageRequest
	GetPromotionId() *string
	SetQuantity(v int32) *CreateEdgeMobileAgentPackageRequest
	GetQuantity() *int32
}

type CreateEdgeMobileAgentPackageRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true**: enables automatic payment. Make sure that your account balance is sufficient.
	//
	// - **false*	- (default): generates the order without charging the account.
	//
	//
	//
	//
	// > If your payment method has an insufficient balance, set this parameter to false. An unpaid order is generated. You can then log on to the Cloud Phone console to complete the payment.
	//
	// >
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false*	- (default): disables auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The region where the agent is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The idempotency token.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The device form factor.
	//
	// example:
	//
	// BOX
	DeviceClass *string `json:"DeviceClass,omitempty" xml:"DeviceClass,omitempty"`
	// The subscription duration of the resource. The unit is specified by PeriodUnit.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// - **Month**: month.
	//
	// - **Year**: year.
	//
	// This parameter is required.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the promotional activity.
	//
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The number of packages.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s CreateEdgeMobileAgentPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeMobileAgentPackageRequest) GoString() string {
	return s.String()
}

func (s *CreateEdgeMobileAgentPackageRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateEdgeMobileAgentPackageRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateEdgeMobileAgentPackageRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateEdgeMobileAgentPackageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEdgeMobileAgentPackageRequest) GetDeviceClass() *string {
	return s.DeviceClass
}

func (s *CreateEdgeMobileAgentPackageRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateEdgeMobileAgentPackageRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateEdgeMobileAgentPackageRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateEdgeMobileAgentPackageRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *CreateEdgeMobileAgentPackageRequest) SetAutoPay(v bool) *CreateEdgeMobileAgentPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateEdgeMobileAgentPackageRequest) SetAutoRenew(v bool) *CreateEdgeMobileAgentPackageRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateEdgeMobileAgentPackageRequest) SetBizRegionId(v string) *CreateEdgeMobileAgentPackageRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateEdgeMobileAgentPackageRequest) SetClientToken(v string) *CreateEdgeMobileAgentPackageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEdgeMobileAgentPackageRequest) SetDeviceClass(v string) *CreateEdgeMobileAgentPackageRequest {
	s.DeviceClass = &v
	return s
}

func (s *CreateEdgeMobileAgentPackageRequest) SetPeriod(v int32) *CreateEdgeMobileAgentPackageRequest {
	s.Period = &v
	return s
}

func (s *CreateEdgeMobileAgentPackageRequest) SetPeriodUnit(v string) *CreateEdgeMobileAgentPackageRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateEdgeMobileAgentPackageRequest) SetPromotionId(v string) *CreateEdgeMobileAgentPackageRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateEdgeMobileAgentPackageRequest) SetQuantity(v int32) *CreateEdgeMobileAgentPackageRequest {
	s.Quantity = &v
	return s
}

func (s *CreateEdgeMobileAgentPackageRequest) Validate() error {
	return dara.Validate(s)
}
