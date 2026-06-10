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
	SetCreateAction(v string) *CreateAppInstanceShrinkRequest
	GetCreateAction() *string
	SetDeployArea(v string) *CreateAppInstanceShrinkRequest
	GetDeployArea() *string
	SetDescription(v string) *CreateAppInstanceShrinkRequest
	GetDescription() *string
	SetDuration(v int32) *CreateAppInstanceShrinkRequest
	GetDuration() *int32
	SetExtend(v string) *CreateAppInstanceShrinkRequest
	GetExtend() *string
	SetName(v string) *CreateAppInstanceShrinkRequest
	GetName() *string
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
	SetVersion(v string) *CreateAppInstanceShrinkRequest
	GetVersion() *string
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
	// example:
	//
	// OPEN_SUBSCRIPTION
	CreateAction *string `json:"CreateAction,omitempty" xml:"CreateAction,omitempty"`
	// Deployment area
	//
	// example:
	//
	// ChineseMainland
	DeployArea *string `json:"DeployArea,omitempty" xml:"DeployArea,omitempty"`
	// example:
	//
	// go-to-the-docks-for-french-fries
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// example:
	//
	// docs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// Resource group ID
	//
	// example:
	//
	// rg-aek2smovqqpvuly
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Site version
	//
	// example:
	//
	// Basic_Edition
	SiteVersion *string `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// List of tags
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 2023-09-01
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
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

func (s *CreateAppInstanceShrinkRequest) GetCreateAction() *string {
	return s.CreateAction
}

func (s *CreateAppInstanceShrinkRequest) GetDeployArea() *string {
	return s.DeployArea
}

func (s *CreateAppInstanceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppInstanceShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateAppInstanceShrinkRequest) GetExtend() *string {
	return s.Extend
}

func (s *CreateAppInstanceShrinkRequest) GetName() *string {
	return s.Name
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

func (s *CreateAppInstanceShrinkRequest) GetVersion() *string {
	return s.Version
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

func (s *CreateAppInstanceShrinkRequest) SetCreateAction(v string) *CreateAppInstanceShrinkRequest {
	s.CreateAction = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDeployArea(v string) *CreateAppInstanceShrinkRequest {
	s.DeployArea = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) SetDescription(v string) *CreateAppInstanceShrinkRequest {
	s.Description = &v
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

func (s *CreateAppInstanceShrinkRequest) SetName(v string) *CreateAppInstanceShrinkRequest {
	s.Name = &v
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

func (s *CreateAppInstanceShrinkRequest) SetVersion(v string) *CreateAppInstanceShrinkRequest {
	s.Version = &v
	return s
}

func (s *CreateAppInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
