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
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The **ClientToken*	- parameter can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the change tracking instance.
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing cycle of the subscription instance. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// >  You must specify this parameter only if you set the PayType parameter to **Prepaid**.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the change tracking instance. The region ID is the same as that of the source instance. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region ID of the change tracking instance. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The subscription length.
	//
	// 	- If the billing cycle is **Year**, the value range is **1 to 5**.
	//
	// 	- If the billing cycle is **Month**, the value range is **1 to 60**.
	//
	// >  You must specify this parameter only if you set the PayType parameter to **Prepaid**.
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
	return dara.Validate(s)
}

type CreateSubscriptionInstanceRequestSourceEndpoint struct {
	// The type of the source instance. Valid values: **MySQL**, **PolarDB**, **DRDS**, and **Oracle**.
	//
	// >  Default value: **MySQL**.
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
