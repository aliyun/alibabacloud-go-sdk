// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmartAccessGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlreadyHaveSag(v bool) *CreateSmartAccessGatewayRequest
	GetAlreadyHaveSag() *bool
	SetAutoPay(v bool) *CreateSmartAccessGatewayRequest
	GetAutoPay() *bool
	SetBuyerMessage(v string) *CreateSmartAccessGatewayRequest
	GetBuyerMessage() *string
	SetCPEVersion(v string) *CreateSmartAccessGatewayRequest
	GetCPEVersion() *string
	SetChargeType(v string) *CreateSmartAccessGatewayRequest
	GetChargeType() *string
	SetDescription(v string) *CreateSmartAccessGatewayRequest
	GetDescription() *string
	SetHaType(v string) *CreateSmartAccessGatewayRequest
	GetHaType() *string
	SetHardWareSpec(v string) *CreateSmartAccessGatewayRequest
	GetHardWareSpec() *string
	SetMaxBandWidth(v int32) *CreateSmartAccessGatewayRequest
	GetMaxBandWidth() *int32
	SetName(v string) *CreateSmartAccessGatewayRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateSmartAccessGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSmartAccessGatewayRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreateSmartAccessGatewayRequest
	GetPeriod() *int32
	SetReceiverAddress(v string) *CreateSmartAccessGatewayRequest
	GetReceiverAddress() *string
	SetReceiverCity(v string) *CreateSmartAccessGatewayRequest
	GetReceiverCity() *string
	SetReceiverCountry(v string) *CreateSmartAccessGatewayRequest
	GetReceiverCountry() *string
	SetReceiverDistrict(v string) *CreateSmartAccessGatewayRequest
	GetReceiverDistrict() *string
	SetReceiverEmail(v string) *CreateSmartAccessGatewayRequest
	GetReceiverEmail() *string
	SetReceiverMobile(v string) *CreateSmartAccessGatewayRequest
	GetReceiverMobile() *string
	SetReceiverName(v string) *CreateSmartAccessGatewayRequest
	GetReceiverName() *string
	SetReceiverPhone(v string) *CreateSmartAccessGatewayRequest
	GetReceiverPhone() *string
	SetReceiverState(v string) *CreateSmartAccessGatewayRequest
	GetReceiverState() *string
	SetReceiverTown(v string) *CreateSmartAccessGatewayRequest
	GetReceiverTown() *string
	SetReceiverZip(v string) *CreateSmartAccessGatewayRequest
	GetReceiverZip() *string
	SetRegionId(v string) *CreateSmartAccessGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateSmartAccessGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmartAccessGatewayRequest
	GetResourceOwnerId() *int64
}

type CreateSmartAccessGatewayRequest struct {
	// Specifies whether you already have an SAG device. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// example:
	//
	// false
	AlreadyHaveSag *bool `json:"AlreadyHaveSag,omitempty" xml:"AlreadyHaveSag,omitempty"`
	// Specifies whether to enable auto-payment for the instance. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// If you set the parameter to false, go to Billing Management to complete the payment after you call this operation. After you complete the payment, the instance can be created.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The remarks left by the buyer.
	//
	// This parameter is required.
	//
	// example:
	//
	// Remarks
	BuyerMessage *string `json:"BuyerMessage,omitempty" xml:"BuyerMessage,omitempty"`
	// The edition of SAG when you create an SAG vCPE instance.
	//
	// Set the value to **basic**, which specifies Basic Edition.
	//
	// example:
	//
	// basic
	CPEVersion *string `json:"CPEVersion,omitempty" xml:"CPEVersion,omitempty"`
	// The billing method of the SAG instance.
	//
	// Set the value to **PREPAY**, which specifies the subscription billing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The description of the SAG instance.
	//
	// The description must be 2 to 256 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// testdesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The deployment mode. Valid values:
	//
	// 	- **no_backup**: You buy only one SAG device to connect private networks to Alibaba Cloud.
	//
	// 	- **cold_backup**: You buy two SAG devices in active-standby mode. One SAG device serves as an active device and the other serves as a standby device. Only the active device is connected to Alibaba Cloud. If the active device is not working as expected, you must manually perform a switchover.
	//
	// 	- **warm_backup**: You buy two SAG devices in active-active mode. Both SAG devices are connected to Alibaba Cloud. If an active device is not working as expected, a failover is automatically performed.
	//
	// >  If you want to create an SAG vCPE instance, set the value to **warm_backup**.
	//
	// This parameter is required.
	//
	// example:
	//
	// no_backup
	HaType *string `json:"HaType,omitempty" xml:"HaType,omitempty"`
	// The type of the SAG instance. Valid values:
	//
	// 	- **sag-100wm**
	//
	// 	- **sag-1000**
	//
	// 	- **sag-vcpe**
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-100wm
	HardWareSpec *string `json:"HardWareSpec,omitempty" xml:"HardWareSpec,omitempty"`
	// The bandwidth of the SAG instance.
	//
	// 	- If you want to create an SAG CPE instance and the model is **sag-100wm**, valid values of this parameter are **2*	- to **50**. Unit: Mbit/s.
	//
	// 	- If you want to create an SAG CPE instance and the model is **sag-1000**, valid values of this parameter are **10*	- to **500**. Unit: Mbit/s.
	//
	// 	- If you want to create an SAG vCPE instance, valid values of this parameter are **10*	- to **1000**. Unit: Mbit/s.
	//
	// example:
	//
	// 12
	MaxBandWidth *int32 `json:"MaxBandWidth,omitempty" xml:"MaxBandWidth,omitempty"`
	// The name of the SAG instance.
	//
	// The name must be 2 to 128 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// testname
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription period of the SAG instance. Unit: months.
	//
	// Valid values: **1*	- to **9**, **12**, **24**, and **36**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The detailed address of the recipient.
	//
	// This parameter is required.
	//
	// example:
	//
	// No.XX
	ReceiverAddress *string `json:"ReceiverAddress,omitempty" xml:"ReceiverAddress,omitempty"`
	// The city of the recipient address.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou
	ReceiverCity *string `json:"ReceiverCity,omitempty" xml:"ReceiverCity,omitempty"`
	// The country of the recipient address.
	//
	// This parameter is required.
	//
	// example:
	//
	// China
	ReceiverCountry *string `json:"ReceiverCountry,omitempty" xml:"ReceiverCountry,omitempty"`
	// The district of the recipient address.
	//
	// This parameter is required.
	//
	// example:
	//
	// West Lake
	ReceiverDistrict *string `json:"ReceiverDistrict,omitempty" xml:"ReceiverDistrict,omitempty"`
	// The email address of the recipient.
	//
	// This parameter is required.
	//
	// example:
	//
	// xx@example.com
	ReceiverEmail *string `json:"ReceiverEmail,omitempty" xml:"ReceiverEmail,omitempty"`
	// The mobile phone number of the recipient.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1884085****
	ReceiverMobile *string `json:"ReceiverMobile,omitempty" xml:"ReceiverMobile,omitempty"`
	// The name of the recipient.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alice
	ReceiverName *string `json:"ReceiverName,omitempty" xml:"ReceiverName,omitempty"`
	// The landline phone number of the recipient.
	//
	// example:
	//
	// 8585****
	ReceiverPhone *string `json:"ReceiverPhone,omitempty" xml:"ReceiverPhone,omitempty"`
	// The province of the recipient address.
	//
	// This parameter is required.
	//
	// example:
	//
	// Zhejiang
	ReceiverState *string `json:"ReceiverState,omitempty" xml:"ReceiverState,omitempty"`
	// The town of the recipient address.
	//
	// This parameter is required.
	//
	// example:
	//
	// Zhuan Tang
	ReceiverTown *string `json:"ReceiverTown,omitempty" xml:"ReceiverTown,omitempty"`
	// The postcode of the recipient address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 310000
	ReceiverZip *string `json:"ReceiverZip,omitempty" xml:"ReceiverZip,omitempty"`
	// The ID of the region where you want to deploy the SAG instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateSmartAccessGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartAccessGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateSmartAccessGatewayRequest) GetAlreadyHaveSag() *bool {
	return s.AlreadyHaveSag
}

func (s *CreateSmartAccessGatewayRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateSmartAccessGatewayRequest) GetBuyerMessage() *string {
	return s.BuyerMessage
}

func (s *CreateSmartAccessGatewayRequest) GetCPEVersion() *string {
	return s.CPEVersion
}

func (s *CreateSmartAccessGatewayRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateSmartAccessGatewayRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSmartAccessGatewayRequest) GetHaType() *string {
	return s.HaType
}

func (s *CreateSmartAccessGatewayRequest) GetHardWareSpec() *string {
	return s.HardWareSpec
}

func (s *CreateSmartAccessGatewayRequest) GetMaxBandWidth() *int32 {
	return s.MaxBandWidth
}

func (s *CreateSmartAccessGatewayRequest) GetName() *string {
	return s.Name
}

func (s *CreateSmartAccessGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSmartAccessGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmartAccessGatewayRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateSmartAccessGatewayRequest) GetReceiverAddress() *string {
	return s.ReceiverAddress
}

func (s *CreateSmartAccessGatewayRequest) GetReceiverCity() *string {
	return s.ReceiverCity
}

func (s *CreateSmartAccessGatewayRequest) GetReceiverCountry() *string {
	return s.ReceiverCountry
}

func (s *CreateSmartAccessGatewayRequest) GetReceiverDistrict() *string {
	return s.ReceiverDistrict
}

func (s *CreateSmartAccessGatewayRequest) GetReceiverEmail() *string {
	return s.ReceiverEmail
}

func (s *CreateSmartAccessGatewayRequest) GetReceiverMobile() *string {
	return s.ReceiverMobile
}

func (s *CreateSmartAccessGatewayRequest) GetReceiverName() *string {
	return s.ReceiverName
}

func (s *CreateSmartAccessGatewayRequest) GetReceiverPhone() *string {
	return s.ReceiverPhone
}

func (s *CreateSmartAccessGatewayRequest) GetReceiverState() *string {
	return s.ReceiverState
}

func (s *CreateSmartAccessGatewayRequest) GetReceiverTown() *string {
	return s.ReceiverTown
}

func (s *CreateSmartAccessGatewayRequest) GetReceiverZip() *string {
	return s.ReceiverZip
}

func (s *CreateSmartAccessGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSmartAccessGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmartAccessGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmartAccessGatewayRequest) SetAlreadyHaveSag(v bool) *CreateSmartAccessGatewayRequest {
	s.AlreadyHaveSag = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetAutoPay(v bool) *CreateSmartAccessGatewayRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetBuyerMessage(v string) *CreateSmartAccessGatewayRequest {
	s.BuyerMessage = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetCPEVersion(v string) *CreateSmartAccessGatewayRequest {
	s.CPEVersion = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetChargeType(v string) *CreateSmartAccessGatewayRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetDescription(v string) *CreateSmartAccessGatewayRequest {
	s.Description = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetHaType(v string) *CreateSmartAccessGatewayRequest {
	s.HaType = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetHardWareSpec(v string) *CreateSmartAccessGatewayRequest {
	s.HardWareSpec = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetMaxBandWidth(v int32) *CreateSmartAccessGatewayRequest {
	s.MaxBandWidth = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetName(v string) *CreateSmartAccessGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetOwnerAccount(v string) *CreateSmartAccessGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetOwnerId(v int64) *CreateSmartAccessGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetPeriod(v int32) *CreateSmartAccessGatewayRequest {
	s.Period = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetReceiverAddress(v string) *CreateSmartAccessGatewayRequest {
	s.ReceiverAddress = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetReceiverCity(v string) *CreateSmartAccessGatewayRequest {
	s.ReceiverCity = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetReceiverCountry(v string) *CreateSmartAccessGatewayRequest {
	s.ReceiverCountry = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetReceiverDistrict(v string) *CreateSmartAccessGatewayRequest {
	s.ReceiverDistrict = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetReceiverEmail(v string) *CreateSmartAccessGatewayRequest {
	s.ReceiverEmail = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetReceiverMobile(v string) *CreateSmartAccessGatewayRequest {
	s.ReceiverMobile = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetReceiverName(v string) *CreateSmartAccessGatewayRequest {
	s.ReceiverName = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetReceiverPhone(v string) *CreateSmartAccessGatewayRequest {
	s.ReceiverPhone = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetReceiverState(v string) *CreateSmartAccessGatewayRequest {
	s.ReceiverState = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetReceiverTown(v string) *CreateSmartAccessGatewayRequest {
	s.ReceiverTown = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetReceiverZip(v string) *CreateSmartAccessGatewayRequest {
	s.ReceiverZip = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetRegionId(v string) *CreateSmartAccessGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetResourceOwnerAccount(v string) *CreateSmartAccessGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) SetResourceOwnerId(v int64) *CreateSmartAccessGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmartAccessGatewayRequest) Validate() error {
	return dara.Validate(s)
}
