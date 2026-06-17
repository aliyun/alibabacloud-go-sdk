// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *ConfirmAppInstanceRequest
	GetApplicationType() *string
	SetAutoRenew(v bool) *ConfirmAppInstanceRequest
	GetAutoRenew() *bool
	SetClientToken(v string) *ConfirmAppInstanceRequest
	GetClientToken() *string
	SetDeployArea(v string) *ConfirmAppInstanceRequest
	GetDeployArea() *string
	SetDuration(v int32) *ConfirmAppInstanceRequest
	GetDuration() *int32
	SetExtend(v string) *ConfirmAppInstanceRequest
	GetExtend() *string
	SetPaymentType(v string) *ConfirmAppInstanceRequest
	GetPaymentType() *string
	SetPricingCycle(v string) *ConfirmAppInstanceRequest
	GetPricingCycle() *string
	SetQuantity(v int32) *ConfirmAppInstanceRequest
	GetQuantity() *int32
	SetSiteVersion(v string) *ConfirmAppInstanceRequest
	GetSiteVersion() *string
	SetTrialBizId(v string) *ConfirmAppInstanceRequest
	GetTrialBizId() *string
	SetVersion(v string) *ConfirmAppInstanceRequest
	GetVersion() *string
}

type ConfirmAppInstanceRequest struct {
	// The application type.
	//
	// example:
	//
	// PC_WebSite
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// Specifies whether to enable auto-renewal upon expiration.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Generate a unique value from your client. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123456
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The deployment region.
	//
	// example:
	//
	// HongKong
	DeployArea *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	// Required. The number of subscription periods.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The extended information.
	//
	// example:
	//
	// {\\"deliveryNodeName\\":\\"视觉设计确认\\",\\"deliveryNodeStatus\\":\\"Reject\\",\\"deliveryOperatorRole\\":\\"Customer\\"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The payment type.
	//
	// example:
	//
	// AUTO_PAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// Required. The unit of the subscription period. Valid values:
	//
	// - Year: year
	//
	// - Month: month
	//
	// - Day: day
	//
	// - Hour: hour.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// Required. The number of instances to purchase.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The website version.
	//
	// example:
	//
	// Basic_Edition
	SiteVersion *string `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// The business ID of the trial instance.
	//
	// example:
	//
	// WS123456
	TrialBizId *string `json:"TrialBizId,omitempty" xml:"TrialBizId,omitempty"`
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ConfirmAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConfirmAppInstanceRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *ConfirmAppInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ConfirmAppInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ConfirmAppInstanceRequest) GetDeployArea() *string {
	return s.DeployArea
}

func (s *ConfirmAppInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ConfirmAppInstanceRequest) GetExtend() *string {
	return s.Extend
}

func (s *ConfirmAppInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ConfirmAppInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *ConfirmAppInstanceRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *ConfirmAppInstanceRequest) GetSiteVersion() *string {
	return s.SiteVersion
}

func (s *ConfirmAppInstanceRequest) GetTrialBizId() *string {
	return s.TrialBizId
}

func (s *ConfirmAppInstanceRequest) GetVersion() *string {
	return s.Version
}

func (s *ConfirmAppInstanceRequest) SetApplicationType(v string) *ConfirmAppInstanceRequest {
	s.ApplicationType = &v
	return s
}

func (s *ConfirmAppInstanceRequest) SetAutoRenew(v bool) *ConfirmAppInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *ConfirmAppInstanceRequest) SetClientToken(v string) *ConfirmAppInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ConfirmAppInstanceRequest) SetDeployArea(v string) *ConfirmAppInstanceRequest {
	s.DeployArea = &v
	return s
}

func (s *ConfirmAppInstanceRequest) SetDuration(v int32) *ConfirmAppInstanceRequest {
	s.Duration = &v
	return s
}

func (s *ConfirmAppInstanceRequest) SetExtend(v string) *ConfirmAppInstanceRequest {
	s.Extend = &v
	return s
}

func (s *ConfirmAppInstanceRequest) SetPaymentType(v string) *ConfirmAppInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *ConfirmAppInstanceRequest) SetPricingCycle(v string) *ConfirmAppInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *ConfirmAppInstanceRequest) SetQuantity(v int32) *ConfirmAppInstanceRequest {
	s.Quantity = &v
	return s
}

func (s *ConfirmAppInstanceRequest) SetSiteVersion(v string) *ConfirmAppInstanceRequest {
	s.SiteVersion = &v
	return s
}

func (s *ConfirmAppInstanceRequest) SetTrialBizId(v string) *ConfirmAppInstanceRequest {
	s.TrialBizId = &v
	return s
}

func (s *ConfirmAppInstanceRequest) SetVersion(v string) *ConfirmAppInstanceRequest {
	s.Version = &v
	return s
}

func (s *ConfirmAppInstanceRequest) Validate() error {
	return dara.Validate(s)
}
