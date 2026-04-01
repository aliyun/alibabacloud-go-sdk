// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenewalPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessInfo(v string) *DescribeRenewalPriceRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *DescribeRenewalPriceRequest
	GetClientToken() *string
	SetDBInstanceClass(v string) *DescribeRenewalPriceRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *DescribeRenewalPriceRequest
	GetDBInstanceId() *string
	SetOrderType(v string) *DescribeRenewalPriceRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *DescribeRenewalPriceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRenewalPriceRequest
	GetOwnerId() *int64
	SetPayType(v string) *DescribeRenewalPriceRequest
	GetPayType() *string
	SetQuantity(v int32) *DescribeRenewalPriceRequest
	GetQuantity() *int32
	SetRegionId(v string) *DescribeRenewalPriceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeRenewalPriceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeRenewalPriceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRenewalPriceRequest
	GetResourceOwnerId() *int64
	SetTimeType(v string) *DescribeRenewalPriceRequest
	GetTimeType() *string
	SetUsedTime(v int32) *DescribeRenewalPriceRequest
	GetUsedTime() *int32
}

type DescribeRenewalPriceRequest struct {
	// The additional business information about the instance.
	//
	// example:
	//
	// 121436975448952
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance type of the instance. For more information, see [Primary instance types](https://help.aliyun.com/document_detail/26312.html). By default, the current instance type applies.
	//
	// example:
	//
	// mysql.n2.medium.2c
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The type of order. Set the value to **BUY**.
	//
	// example:
	//
	// BUY
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The number of the instances. Default value: **1**.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmx****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The renewal cycle of the instance. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// This parameter is required.
	//
	// example:
	//
	// Year
	TimeType *string `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The subscription duration of the instance. Valid values:
	//
	// 	- If you set the **TimeType*	- parameter to **Year**, the value of the UsedTime parameter is within the range of **1 to 3**.
	//
	// 	- If you set the **TimeType*	- parameter to **Month**, the value of the UsedTime parameter is within the range of **1 to 9**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s DescribeRenewalPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *DescribeRenewalPriceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeRenewalPriceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeRenewalPriceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeRenewalPriceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeRenewalPriceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRenewalPriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRenewalPriceRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeRenewalPriceRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *DescribeRenewalPriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRenewalPriceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRenewalPriceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRenewalPriceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRenewalPriceRequest) GetTimeType() *string {
	return s.TimeType
}

func (s *DescribeRenewalPriceRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *DescribeRenewalPriceRequest) SetBusinessInfo(v string) *DescribeRenewalPriceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetClientToken(v string) *DescribeRenewalPriceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetDBInstanceClass(v string) *DescribeRenewalPriceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetDBInstanceId(v string) *DescribeRenewalPriceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetOrderType(v string) *DescribeRenewalPriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetOwnerAccount(v string) *DescribeRenewalPriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetOwnerId(v int64) *DescribeRenewalPriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetPayType(v string) *DescribeRenewalPriceRequest {
	s.PayType = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetQuantity(v int32) *DescribeRenewalPriceRequest {
	s.Quantity = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetRegionId(v string) *DescribeRenewalPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetResourceGroupId(v string) *DescribeRenewalPriceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetResourceOwnerAccount(v string) *DescribeRenewalPriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetResourceOwnerId(v int64) *DescribeRenewalPriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetTimeType(v string) *DescribeRenewalPriceRequest {
	s.TimeType = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetUsedTime(v int32) *DescribeRenewalPriceRequest {
	s.UsedTime = &v
	return s
}

func (s *DescribeRenewalPriceRequest) Validate() error {
	return dara.Validate(s)
}
