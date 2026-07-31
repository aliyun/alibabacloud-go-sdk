// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseStorageCapacityUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *PurchaseStorageCapacityUnitRequest
	GetAmount() *int32
	SetCapacity(v int32) *PurchaseStorageCapacityUnitRequest
	GetCapacity() *int32
	SetClientToken(v string) *PurchaseStorageCapacityUnitRequest
	GetClientToken() *string
	SetDescription(v string) *PurchaseStorageCapacityUnitRequest
	GetDescription() *string
	SetFromApp(v string) *PurchaseStorageCapacityUnitRequest
	GetFromApp() *string
	SetName(v string) *PurchaseStorageCapacityUnitRequest
	GetName() *string
	SetOwnerAccount(v string) *PurchaseStorageCapacityUnitRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *PurchaseStorageCapacityUnitRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *PurchaseStorageCapacityUnitRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *PurchaseStorageCapacityUnitRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *PurchaseStorageCapacityUnitRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *PurchaseStorageCapacityUnitRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *PurchaseStorageCapacityUnitRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PurchaseStorageCapacityUnitRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *PurchaseStorageCapacityUnitRequest
	GetStartTime() *string
	SetTag(v []*PurchaseStorageCapacityUnitRequestTag) *PurchaseStorageCapacityUnitRequest
	GetTag() []*PurchaseStorageCapacityUnitRequestTag
}

type PurchaseStorageCapacityUnitRequest struct {
	// The number of SCUs to purchase. Valid values: 1 to 20.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The capacity of the SCU. Unit: GiB. Valid values: 20, 40, 100, 200, 500, 1024, 2048, 5210, 10240, 20480, and 52100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Capacity *int32 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the SCU. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ScuPurchaseDemo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The source of the request. The default value is OpenAPI. You do not need to manually set this parameter.
	//
	// example:
	//
	// OpenAPI
	FromApp *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	// The name of the SCU. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// ScuPurchaseDemo
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The validity period of the SCU. Valid values:
	//
	// - When PeriodUnit is set to Month, valid values of Period are 1, 2, 3, and 6.
	//
	// - When PeriodUnit is set to Year, valid values of Period are 1, 3, and 5.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the validity period of the SCU. Valid values:
	//
	// - Month: month.
	//
	// - Year: year.
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the SCU. After you specify a region, the SCU can only offset pay-as-you-go bills for cloud disks in that region. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the storage capacity unit (SCU) belongs. You can specify only a resource group ID to which you have permissions.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The effective period of the SCU. The effective period cannot be more than 180 days after the creation time. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHHZ format. The time must be in UTC.
	//
	// Default value: null, which indicates that the SCU takes effect immediately after it is created.
	//
	// example:
	//
	// 2020-09-09T02Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The tags. Array length: 1 to 20.
	Tag []*PurchaseStorageCapacityUnitRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s PurchaseStorageCapacityUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s PurchaseStorageCapacityUnitRequest) GoString() string {
	return s.String()
}

func (s *PurchaseStorageCapacityUnitRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *PurchaseStorageCapacityUnitRequest) GetCapacity() *int32 {
	return s.Capacity
}

func (s *PurchaseStorageCapacityUnitRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PurchaseStorageCapacityUnitRequest) GetDescription() *string {
	return s.Description
}

func (s *PurchaseStorageCapacityUnitRequest) GetFromApp() *string {
	return s.FromApp
}

func (s *PurchaseStorageCapacityUnitRequest) GetName() *string {
	return s.Name
}

func (s *PurchaseStorageCapacityUnitRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *PurchaseStorageCapacityUnitRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PurchaseStorageCapacityUnitRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *PurchaseStorageCapacityUnitRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *PurchaseStorageCapacityUnitRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PurchaseStorageCapacityUnitRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *PurchaseStorageCapacityUnitRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PurchaseStorageCapacityUnitRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PurchaseStorageCapacityUnitRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *PurchaseStorageCapacityUnitRequest) GetTag() []*PurchaseStorageCapacityUnitRequestTag {
	return s.Tag
}

func (s *PurchaseStorageCapacityUnitRequest) SetAmount(v int32) *PurchaseStorageCapacityUnitRequest {
	s.Amount = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetCapacity(v int32) *PurchaseStorageCapacityUnitRequest {
	s.Capacity = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetClientToken(v string) *PurchaseStorageCapacityUnitRequest {
	s.ClientToken = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetDescription(v string) *PurchaseStorageCapacityUnitRequest {
	s.Description = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetFromApp(v string) *PurchaseStorageCapacityUnitRequest {
	s.FromApp = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetName(v string) *PurchaseStorageCapacityUnitRequest {
	s.Name = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetOwnerAccount(v string) *PurchaseStorageCapacityUnitRequest {
	s.OwnerAccount = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetOwnerId(v int64) *PurchaseStorageCapacityUnitRequest {
	s.OwnerId = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetPeriod(v int32) *PurchaseStorageCapacityUnitRequest {
	s.Period = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetPeriodUnit(v string) *PurchaseStorageCapacityUnitRequest {
	s.PeriodUnit = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetRegionId(v string) *PurchaseStorageCapacityUnitRequest {
	s.RegionId = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetResourceGroupId(v string) *PurchaseStorageCapacityUnitRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetResourceOwnerAccount(v string) *PurchaseStorageCapacityUnitRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetResourceOwnerId(v int64) *PurchaseStorageCapacityUnitRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetStartTime(v string) *PurchaseStorageCapacityUnitRequest {
	s.StartTime = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) SetTag(v []*PurchaseStorageCapacityUnitRequestTag) *PurchaseStorageCapacityUnitRequest {
	s.Tag = v
	return s
}

func (s *PurchaseStorageCapacityUnitRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PurchaseStorageCapacityUnitRequestTag struct {
	// The tag key of the SCU. If you specify this parameter, the tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the SCU. If you specify this parameter, the tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PurchaseStorageCapacityUnitRequestTag) String() string {
	return dara.Prettify(s)
}

func (s PurchaseStorageCapacityUnitRequestTag) GoString() string {
	return s.String()
}

func (s *PurchaseStorageCapacityUnitRequestTag) GetKey() *string {
	return s.Key
}

func (s *PurchaseStorageCapacityUnitRequestTag) GetValue() *string {
	return s.Value
}

func (s *PurchaseStorageCapacityUnitRequestTag) SetKey(v string) *PurchaseStorageCapacityUnitRequestTag {
	s.Key = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequestTag) SetValue(v string) *PurchaseStorageCapacityUnitRequestTag {
	s.Value = &v
	return s
}

func (s *PurchaseStorageCapacityUnitRequestTag) Validate() error {
	return dara.Validate(s)
}
