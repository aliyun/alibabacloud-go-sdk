// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *ReleaseInstanceRequest
	GetInstanceIds() *string
	SetOwnerId(v int64) *ReleaseInstanceRequest
	GetOwnerId() *int64
	SetProductCode(v string) *ReleaseInstanceRequest
	GetProductCode() *string
	SetProductType(v string) *ReleaseInstanceRequest
	GetProductType() *string
	SetRegion(v string) *ReleaseInstanceRequest
	GetRegion() *string
	SetRenewStatus(v string) *ReleaseInstanceRequest
	GetRenewStatus() *string
	SetSubscriptionType(v string) *ReleaseInstanceRequest
	GetSubscriptionType() *string
}

type ReleaseInstanceRequest struct {
	// The ID of the instance. Separate multiple IDs with commas (,). A maximum of 100 IDs can be specified.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-xxxxxxxxxxxx
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// This parameter is required.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The method that is used to renew the instance. Valid values:
	//
	// AutoRenewal: automatically renews the instance.
	//
	// ManualRenewal: manually renews the instance.
	//
	// NotRenewal: does not renew the instance.
	//
	// example:
	//
	// AutoRenewal
	RenewStatus *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	// The billing method. Valid values:
	//
	// Subscription: the subscription billing method.
	//
	// PayAsYouGo: the pay-as-you-go billing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ReleaseInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseInstanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ReleaseInstanceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ReleaseInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *ReleaseInstanceRequest) GetRenewStatus() *string {
	return s.RenewStatus
}

func (s *ReleaseInstanceRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *ReleaseInstanceRequest) SetInstanceIds(v string) *ReleaseInstanceRequest {
	s.InstanceIds = &v
	return s
}

func (s *ReleaseInstanceRequest) SetOwnerId(v int64) *ReleaseInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseInstanceRequest) SetProductCode(v string) *ReleaseInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *ReleaseInstanceRequest) SetProductType(v string) *ReleaseInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *ReleaseInstanceRequest) SetRegion(v string) *ReleaseInstanceRequest {
	s.Region = &v
	return s
}

func (s *ReleaseInstanceRequest) SetRenewStatus(v string) *ReleaseInstanceRequest {
	s.RenewStatus = &v
	return s
}

func (s *ReleaseInstanceRequest) SetSubscriptionType(v string) *ReleaseInstanceRequest {
	s.SubscriptionType = &v
	return s
}

func (s *ReleaseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
