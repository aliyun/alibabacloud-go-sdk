// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyInstanceChargeTypeRequest
	GetAutoPay() *bool
	SetClientToken(v string) *ModifyInstanceChargeTypeRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyInstanceChargeTypeRequest
	GetDryRun() *bool
	SetIncludeDataDisks(v bool) *ModifyInstanceChargeTypeRequest
	GetIncludeDataDisks() *bool
	SetInstanceChargeType(v string) *ModifyInstanceChargeTypeRequest
	GetInstanceChargeType() *string
	SetInstanceIds(v string) *ModifyInstanceChargeTypeRequest
	GetInstanceIds() *string
	SetIsDetailFee(v bool) *ModifyInstanceChargeTypeRequest
	GetIsDetailFee() *bool
	SetOwnerAccount(v string) *ModifyInstanceChargeTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceChargeTypeRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *ModifyInstanceChargeTypeRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *ModifyInstanceChargeTypeRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *ModifyInstanceChargeTypeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceChargeTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceChargeTypeRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceChargeTypeRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - true: Automatic payment is enabled. Make sure that your account balance is sufficient. If your account balance is insufficient, an abnormal order is generated and can only be canceled.
	//
	// - false: Only an order is generated and no payment is made.
	//
	// Default value: true.
	//
	// > If the balance of your payment method is insufficient, you can set the AutoPay parameter to false. In this case, an unpaid order is generated, and you can log on to the ECS console to complete the payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, and make sure that the value is unique across different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only perform a dry run for this request. Valid values:
	//
	// - true: send a dry run request. The system does not query resources. The system checks whether your AccessKey pair is valid, whether the Resource Access Management (RAM) user has the required authorization, and whether required parameters are specified. If the dry run fails, the corresponding error code is returned. If the dry run succeeds, the `DryRunOperation` error code is returned.
	//
	// - false: send a normal request. If the dry run succeeds, the 2xx HTTP status code is returned and resources are queried.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to change the billing method of all pay-as-you-go data disks attached to the instance to subscription. Valid values:
	//
	// - true
	//
	// - false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IncludeDataDisks *bool `json:"IncludeDataDisks,omitempty" xml:"IncludeDataDisks,omitempty"`
	// The target billing method of the instance. Valid values:
	//
	// - PrePaid: switches the billing method of the instance from pay-as-you-go to subscription.
	//
	// - PostPaid: switches the billing method of the instance from subscription to pay-as-you-go.
	//
	// Default value: PrePaid.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance IDs. The value can be a JSON array that consists of up to 20 instance IDs. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["i-bp67acfmxazb4p****","i-bp67acfmxazb4d****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// Specifies whether to return the order fee details when you switch the billing method from subscription to pay-as-you-go. Valid values:
	//
	// - true
	//
	// - false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IsDetailFee  *bool   `json:"IsDetailFee,omitempty" xml:"IsDetailFee,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription period. If the host of the ECS instance is a dedicated host (DDH), the value cannot exceed the subscription period of the DDH. Valid values:
	//
	// <props="china">
	//
	// - When `PeriodUnit` is set to `Week`, valid values of `Period` are 1, 2, 3, and 4.
	//
	// - When `PeriodUnit` is set to `Month`, valid values of `Period` are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	//
	// <props="intl">When `PeriodUnit` is set to `Month`, valid values of `Period` are 1, 2, 3, 4, 5, 6, 7, 8, 9, and 12.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period, which is the unit of the `Period` parameter. Valid values:
	//
	// <props="china">
	//
	// - Week
	//
	// - Month
	//
	// - Year
	//
	// <props="intl">Month
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyInstanceChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyInstanceChargeTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceChargeTypeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyInstanceChargeTypeRequest) GetIncludeDataDisks() *bool {
	return s.IncludeDataDisks
}

func (s *ModifyInstanceChargeTypeRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ModifyInstanceChargeTypeRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ModifyInstanceChargeTypeRequest) GetIsDetailFee() *bool {
	return s.IsDetailFee
}

func (s *ModifyInstanceChargeTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceChargeTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceChargeTypeRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyInstanceChargeTypeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyInstanceChargeTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceChargeTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceChargeTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceChargeTypeRequest) SetAutoPay(v bool) *ModifyInstanceChargeTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetClientToken(v string) *ModifyInstanceChargeTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetDryRun(v bool) *ModifyInstanceChargeTypeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetIncludeDataDisks(v bool) *ModifyInstanceChargeTypeRequest {
	s.IncludeDataDisks = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetInstanceChargeType(v string) *ModifyInstanceChargeTypeRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetInstanceIds(v string) *ModifyInstanceChargeTypeRequest {
	s.InstanceIds = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetIsDetailFee(v bool) *ModifyInstanceChargeTypeRequest {
	s.IsDetailFee = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetOwnerAccount(v string) *ModifyInstanceChargeTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetOwnerId(v int64) *ModifyInstanceChargeTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetPeriod(v int32) *ModifyInstanceChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetPeriodUnit(v string) *ModifyInstanceChargeTypeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetRegionId(v string) *ModifyInstanceChargeTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceChargeTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) SetResourceOwnerId(v int64) *ModifyInstanceChargeTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
