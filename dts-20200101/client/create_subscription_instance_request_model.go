// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscriptionInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceEndpoint(v *CreateSubscriptionInstanceRequestSourceEndpoint) *CreateSubscriptionInstanceRequest
	GetSourceEndpoint() *CreateSubscriptionInstanceRequestSourceEndpoint
	SetAccountId(v string) *CreateSubscriptionInstanceRequest
	GetAccountId() *string
	SetClientToken(v string) *CreateSubscriptionInstanceRequest
	GetClientToken() *string
	SetOwnerId(v string) *CreateSubscriptionInstanceRequest
	GetOwnerId() *string
	SetPayType(v string) *CreateSubscriptionInstanceRequest
	GetPayType() *string
	SetPeriod(v string) *CreateSubscriptionInstanceRequest
	GetPeriod() *string
	SetRegion(v string) *CreateSubscriptionInstanceRequest
	GetRegion() *string
	SetRegionId(v string) *CreateSubscriptionInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateSubscriptionInstanceRequest
	GetResourceGroupId() *string
	SetUsedTime(v int32) *CreateSubscriptionInstanceRequest
	GetUsedTime() *int32
}

type CreateSubscriptionInstanceRequest struct {
	SourceEndpoint *CreateSubscriptionInstanceRequestSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter. This parameter will be discontinued.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Generate a value from your client to make sure that the value is unique among different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method.
	//
	// - **Postpaid**: pay-as-you-go. This is the default value.
	//
	// - **Prepaid**: subscription.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing method of the subscription instance. Valid values:
	//
	// - **Year**: annual subscription.
	//
	// - **Month**: monthly subscription.
	//
	// > This parameter is valid and required only when PayType is set to **Prepaid*	- (subscription).
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID. Set this parameter to the region where the subscription object resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region to which the change tracking instance belongs. You do not need to specify this parameter. This parameter will be discontinued.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The purchase duration of the subscription instance.
	//
	// - If the billing method is set to **Year*	- (annual subscription), the valid values are **1 to 5**.
	//
	// - If the billing method is set to **Month*	- (monthly subscription), the valid values are **1 to 60**.
	//
	// > This parameter is valid and required only when PayType is set to **Prepaid*	- (subscription).
	//
	// example:
	//
	// 12
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s CreateSubscriptionInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscriptionInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionInstanceRequest) GetSourceEndpoint() *CreateSubscriptionInstanceRequestSourceEndpoint {
	return s.SourceEndpoint
}

func (s *CreateSubscriptionInstanceRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateSubscriptionInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSubscriptionInstanceRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateSubscriptionInstanceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateSubscriptionInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateSubscriptionInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateSubscriptionInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSubscriptionInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSubscriptionInstanceRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateSubscriptionInstanceRequest) SetSourceEndpoint(v *CreateSubscriptionInstanceRequestSourceEndpoint) *CreateSubscriptionInstanceRequest {
	s.SourceEndpoint = v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetAccountId(v string) *CreateSubscriptionInstanceRequest {
	s.AccountId = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetClientToken(v string) *CreateSubscriptionInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetOwnerId(v string) *CreateSubscriptionInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetPayType(v string) *CreateSubscriptionInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetPeriod(v string) *CreateSubscriptionInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetRegion(v string) *CreateSubscriptionInstanceRequest {
	s.Region = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetRegionId(v string) *CreateSubscriptionInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetResourceGroupId(v string) *CreateSubscriptionInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetUsedTime(v int32) *CreateSubscriptionInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) Validate() error {
	if s.SourceEndpoint != nil {
		if err := s.SourceEndpoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSubscriptionInstanceRequestSourceEndpoint struct {
	// 数据订阅的实例类型，取值为：**MySQL**、**PolarDB**、**DRDS**、**Oracle**。
	//
	// > 默认取值为：**MySQL**。
	//
	// example:
	//
	// MySQL
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateSubscriptionInstanceRequestSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscriptionInstanceRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionInstanceRequestSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateSubscriptionInstanceRequestSourceEndpoint) SetInstanceType(v string) *CreateSubscriptionInstanceRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *CreateSubscriptionInstanceRequestSourceEndpoint) Validate() error {
	return dara.Validate(s)
}
