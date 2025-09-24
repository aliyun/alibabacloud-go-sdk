// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRenewalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIDs(v string) *SetRenewalRequest
	GetInstanceIDs() *string
	SetOwnerId(v int64) *SetRenewalRequest
	GetOwnerId() *int64
	SetProductCode(v string) *SetRenewalRequest
	GetProductCode() *string
	SetProductType(v string) *SetRenewalRequest
	GetProductType() *string
	SetRenewalPeriod(v int32) *SetRenewalRequest
	GetRenewalPeriod() *int32
	SetRenewalPeriodUnit(v string) *SetRenewalRequest
	GetRenewalPeriodUnit() *string
	SetRenewalStatus(v string) *SetRenewalRequest
	GetRenewalStatus() *string
	SetSubscriptionType(v string) *SetRenewalRequest
	GetSubscriptionType() *string
}

type SetRenewalRequest struct {
	// The ID of the instance. You can enable auto-renewal for up to 100 subscription instances at a time. Separate multiple instance IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-askjdhaskfjh
	InstanceIDs *string `json:"InstanceIDs,omitempty" xml:"InstanceIDs,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The auto-renewal period. Valid values:
	//
	// 	- 1
	//
	// 	- 2
	//
	// 	- 3
	//
	// 	- 6
	//
	// 	- 12
	//
	// >  This parameter is required if the RenewalStatus parameter is set to AutoRenewal.
	//
	// example:
	//
	// 1
	RenewalPeriod *int32 `json:"RenewalPeriod,omitempty" xml:"RenewalPeriod,omitempty"`
	// The unit of the auto-renewal period. Valid values:
	//
	// 	- M: months
	//
	// 	- Y: years
	//
	// >  This parameter is required if the RenewalStatus parameter is set to AutoRenewal.
	//
	// example:
	//
	// M
	RenewalPeriodUnit *string `json:"RenewalPeriodUnit,omitempty" xml:"RenewalPeriodUnit,omitempty"`
	// The status of renewal. Valid values:
	//
	// 	- AutoRenewal: The instance is automatically renewed.
	//
	// 	- ManualRenewal: The instance is manually renewed.
	//
	// 	- NotRenewal: The instance is not renewed.
	//
	// This parameter is required.
	//
	// example:
	//
	// AutoRenewal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: subscription
	//
	// 	- PayAsYouGo: pay-as-you-go
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s SetRenewalRequest) String() string {
	return dara.Prettify(s)
}

func (s SetRenewalRequest) GoString() string {
	return s.String()
}

func (s *SetRenewalRequest) GetInstanceIDs() *string {
	return s.InstanceIDs
}

func (s *SetRenewalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetRenewalRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *SetRenewalRequest) GetProductType() *string {
	return s.ProductType
}

func (s *SetRenewalRequest) GetRenewalPeriod() *int32 {
	return s.RenewalPeriod
}

func (s *SetRenewalRequest) GetRenewalPeriodUnit() *string {
	return s.RenewalPeriodUnit
}

func (s *SetRenewalRequest) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *SetRenewalRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *SetRenewalRequest) SetInstanceIDs(v string) *SetRenewalRequest {
	s.InstanceIDs = &v
	return s
}

func (s *SetRenewalRequest) SetOwnerId(v int64) *SetRenewalRequest {
	s.OwnerId = &v
	return s
}

func (s *SetRenewalRequest) SetProductCode(v string) *SetRenewalRequest {
	s.ProductCode = &v
	return s
}

func (s *SetRenewalRequest) SetProductType(v string) *SetRenewalRequest {
	s.ProductType = &v
	return s
}

func (s *SetRenewalRequest) SetRenewalPeriod(v int32) *SetRenewalRequest {
	s.RenewalPeriod = &v
	return s
}

func (s *SetRenewalRequest) SetRenewalPeriodUnit(v string) *SetRenewalRequest {
	s.RenewalPeriodUnit = &v
	return s
}

func (s *SetRenewalRequest) SetRenewalStatus(v string) *SetRenewalRequest {
	s.RenewalStatus = &v
	return s
}

func (s *SetRenewalRequest) SetSubscriptionType(v string) *SetRenewalRequest {
	s.SubscriptionType = &v
	return s
}

func (s *SetRenewalRequest) Validate() error {
	return dara.Validate(s)
}
