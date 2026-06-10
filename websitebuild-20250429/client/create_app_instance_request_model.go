// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *CreateAppInstanceRequest
	GetApplicationType() *string
	SetAutoRenew(v bool) *CreateAppInstanceRequest
	GetAutoRenew() *bool
	SetClientToken(v string) *CreateAppInstanceRequest
	GetClientToken() *string
	SetCreateAction(v string) *CreateAppInstanceRequest
	GetCreateAction() *string
	SetDeployArea(v string) *CreateAppInstanceRequest
	GetDeployArea() *string
	SetDescription(v string) *CreateAppInstanceRequest
	GetDescription() *string
	SetDuration(v int32) *CreateAppInstanceRequest
	GetDuration() *int32
	SetExtend(v string) *CreateAppInstanceRequest
	GetExtend() *string
	SetName(v string) *CreateAppInstanceRequest
	GetName() *string
	SetPaymentType(v string) *CreateAppInstanceRequest
	GetPaymentType() *string
	SetPricingCycle(v string) *CreateAppInstanceRequest
	GetPricingCycle() *string
	SetQuantity(v int32) *CreateAppInstanceRequest
	GetQuantity() *int32
	SetResourceGroupId(v string) *CreateAppInstanceRequest
	GetResourceGroupId() *string
	SetSiteVersion(v string) *CreateAppInstanceRequest
	GetSiteVersion() *string
	SetTags(v []*CreateAppInstanceRequestTags) *CreateAppInstanceRequest
	GetTags() []*CreateAppInstanceRequestTags
	SetVersion(v string) *CreateAppInstanceRequest
	GetVersion() *string
}

type CreateAppInstanceRequest struct {
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
	Tags []*CreateAppInstanceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-09-01
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *CreateAppInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateAppInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppInstanceRequest) GetCreateAction() *string {
	return s.CreateAction
}

func (s *CreateAppInstanceRequest) GetDeployArea() *string {
	return s.DeployArea
}

func (s *CreateAppInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateAppInstanceRequest) GetExtend() *string {
	return s.Extend
}

func (s *CreateAppInstanceRequest) GetName() *string {
	return s.Name
}

func (s *CreateAppInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateAppInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateAppInstanceRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *CreateAppInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAppInstanceRequest) GetSiteVersion() *string {
	return s.SiteVersion
}

func (s *CreateAppInstanceRequest) GetTags() []*CreateAppInstanceRequestTags {
	return s.Tags
}

func (s *CreateAppInstanceRequest) GetVersion() *string {
	return s.Version
}

func (s *CreateAppInstanceRequest) SetApplicationType(v string) *CreateAppInstanceRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateAppInstanceRequest) SetAutoRenew(v bool) *CreateAppInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAppInstanceRequest) SetClientToken(v string) *CreateAppInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppInstanceRequest) SetCreateAction(v string) *CreateAppInstanceRequest {
	s.CreateAction = &v
	return s
}

func (s *CreateAppInstanceRequest) SetDeployArea(v string) *CreateAppInstanceRequest {
	s.DeployArea = &v
	return s
}

func (s *CreateAppInstanceRequest) SetDescription(v string) *CreateAppInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateAppInstanceRequest) SetDuration(v int32) *CreateAppInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateAppInstanceRequest) SetExtend(v string) *CreateAppInstanceRequest {
	s.Extend = &v
	return s
}

func (s *CreateAppInstanceRequest) SetName(v string) *CreateAppInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateAppInstanceRequest) SetPaymentType(v string) *CreateAppInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateAppInstanceRequest) SetPricingCycle(v string) *CreateAppInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateAppInstanceRequest) SetQuantity(v int32) *CreateAppInstanceRequest {
	s.Quantity = &v
	return s
}

func (s *CreateAppInstanceRequest) SetResourceGroupId(v string) *CreateAppInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAppInstanceRequest) SetSiteVersion(v string) *CreateAppInstanceRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateAppInstanceRequest) SetTags(v []*CreateAppInstanceRequestTags) *CreateAppInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreateAppInstanceRequest) SetVersion(v string) *CreateAppInstanceRequest {
	s.Version = &v
	return s
}

func (s *CreateAppInstanceRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppInstanceRequestTags struct {
	// tag key
	//
	// example:
	//
	// Group
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// value of tag 0
	//
	// example:
	//
	// ufo
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateAppInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceRequestTags) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateAppInstanceRequestTags) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateAppInstanceRequestTags) SetTagKey(v string) *CreateAppInstanceRequestTags {
	s.TagKey = &v
	return s
}

func (s *CreateAppInstanceRequestTags) SetTagValue(v string) *CreateAppInstanceRequestTags {
	s.TagValue = &v
	return s
}

func (s *CreateAppInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
