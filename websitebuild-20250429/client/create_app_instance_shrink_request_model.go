// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *CreateAppInstanceShrinkRequest
	GetApplicationType() *string
	SetAutoRenew(v bool) *CreateAppInstanceShrinkRequest
	GetAutoRenew() *bool
	SetClientToken(v string) *CreateAppInstanceShrinkRequest
	GetClientToken() *string
	SetDeployArea(v string) *CreateAppInstanceShrinkRequest
	GetDeployArea() *string
	SetDuration(v int32) *CreateAppInstanceShrinkRequest
	GetDuration() *int32
	SetExtend(v string) *CreateAppInstanceShrinkRequest
	GetExtend() *string
	SetPaymentType(v string) *CreateAppInstanceShrinkRequest
	GetPaymentType() *string
	SetPricingCycle(v string) *CreateAppInstanceShrinkRequest
	GetPricingCycle() *string
	SetQuantity(v int32) *CreateAppInstanceShrinkRequest
	GetQuantity() *int32
	SetResourceGroupId(v string) *CreateAppInstanceShrinkRequest
	GetResourceGroupId() *string
	SetSiteVersion(v string) *CreateAppInstanceShrinkRequest
	GetSiteVersion() *string
	SetTagsShrink(v string) *CreateAppInstanceShrinkRequest
	GetTagsShrink() *string
}

type CreateAppInstanceShrinkRequest struct {
	// Application type
	//
	// example:
	//
	// PC_WebSite
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// Whether to enable auto-renewal upon expiration
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Ensures idempotence of the request. Generate a unique value from your client to ensure that it is unique across different requests. ClientToken only supports ASCII characters and cannot exceed 64 characters
	//
	// example:
	//
	// 210713a117660695309606626a
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Deployment area
	//
	// example:
	//
	// ChineseMainland
	DeployArea *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	// Required. The number of subscription periods
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Extended information
	//
	// example:
	//
	// {}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// Payment type
	//
	// example:
	//
	// AUTO_PAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// Required. The unit of the subscription period, Year: Year, Month: Month, Day: Day, Hour: Hour
	//
	// example:
	//
	// Year
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// Required. The quantity of instances to be ordered.
	//
	// example:
	//
	// 1
	Quantity        *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Site version
	//
	// example:
	//
	// Basic_Edition
	SiteVersion *string `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	TagsShrink  *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateAppInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceShrinkRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *CreateAppInstanceShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateAppInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppInstanceShrinkRequest) GetDeployArea() *string {
	return s.DeployArea
}

func (s *CreateAppInstanceShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateAppInstanceShrinkRequest) GetExtend() *string {
	return s.Extend
}

func (s *CreateAppInstanceShrinkRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateAppInstanceShrinkRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateAppInstanceShrinkRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *CreateAppInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAppInstanceShrinkRequest) GetSiteVersion() *string {
	return s.SiteVersion
}

func (s *CreateAppInstanceShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateAppInstanceShrinkRequest) SetApplicationType(v string) *CreateAppInstanceShrinkRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetAutoRenew(v bool) *CreateAppInstanceShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetClientToken(v string) *CreateAppInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDeployArea(v string) *CreateAppInstanceShrinkRequest {
	s.DeployArea = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDuration(v int32) *CreateAppInstanceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetExtend(v string) *CreateAppInstanceShrinkRequest {
	s.Extend = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetPaymentType(v string) *CreateAppInstanceShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetPricingCycle(v string) *CreateAppInstanceShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetQuantity(v int32) *CreateAppInstanceShrinkRequest {
	s.Quantity = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetResourceGroupId(v string) *CreateAppInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetSiteVersion(v string) *CreateAppInstanceShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetTagsShrink(v string) *CreateAppInstanceShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
