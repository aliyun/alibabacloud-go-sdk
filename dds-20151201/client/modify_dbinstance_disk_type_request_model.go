// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDiskTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyDBInstanceDiskTypeRequest
	GetAutoPay() *bool
	SetAutoRenew(v string) *ModifyDBInstanceDiskTypeRequest
	GetAutoRenew() *string
	SetBusinessInfo(v string) *ModifyDBInstanceDiskTypeRequest
	GetBusinessInfo() *string
	SetCouponNo(v string) *ModifyDBInstanceDiskTypeRequest
	GetCouponNo() *string
	SetDBInstanceId(v string) *ModifyDBInstanceDiskTypeRequest
	GetDBInstanceId() *string
	SetDbInstanceStorageType(v string) *ModifyDBInstanceDiskTypeRequest
	GetDbInstanceStorageType() *string
	SetExtraParam(v string) *ModifyDBInstanceDiskTypeRequest
	GetExtraParam() *string
	SetOrderType(v string) *ModifyDBInstanceDiskTypeRequest
	GetOrderType() *string
	SetProvisionedIops(v int64) *ModifyDBInstanceDiskTypeRequest
	GetProvisionedIops() *int64
	SetResourceOwnerId(v int64) *ModifyDBInstanceDiskTypeRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceDiskTypeRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**: enables automatic payment. Make sure that you have sufficient balance within your account.
	//
	// 	- **false**: disables automatic payment. You must perform the following operations to pay for the instance: Log on to the ApsaraDB for MongoDB console. In the upper-right corner, click **Expenses*	- and select **User Center*	- from the drop-down list. The User Center page appears. In the left-side navigation pane, choose **Order Management*	- > Renew. On the Renewal tab, find the bill that you want to pay and then click Renew in the Actions column.
	//
	// Default value: **true**.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// Default value: **false**.
	//
	// example:
	//
	// false
	AutoRenew *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The business information. This is an additional parameter.
	//
	// example:
	//
	// {“ActivityId":"000000000"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The coupon code. Default value: `youhuiquan_promotion_option_id_for_blank`.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1fa5efaa93****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The new disk type. Valid values:
	//
	// 	- **cloud_auto**: ESSD AutoPL disk
	//
	// 	- **cloud_essd1**: PL1 ESSD
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// example:
	//
	// cloud_auto
	DbInstanceStorageType *string `json:"DbInstanceStorageType,omitempty" xml:"DbInstanceStorageType,omitempty"`
	// An additional parameter.
	//
	// example:
	//
	// async
	ExtraParam *string `json:"ExtraParam,omitempty" xml:"ExtraParam,omitempty"`
	// The type of configuration changes. Valid values:
	//
	// 	- **UPGRADE**
	//
	// 	- **DOWNGRADE**
	//
	// >  This parameter is valid only when the billing method of the instance is subscription.
	//
	// example:
	//
	// UPGRADE
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The provisioned IOPS. Valid values: 0 to 50000.
	//
	// example:
	//
	// 1960
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceDiskTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDiskTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDiskTypeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyDBInstanceDiskTypeRequest) GetAutoRenew() *string {
	return s.AutoRenew
}

func (s *ModifyDBInstanceDiskTypeRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *ModifyDBInstanceDiskTypeRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *ModifyDBInstanceDiskTypeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceDiskTypeRequest) GetDbInstanceStorageType() *string {
	return s.DbInstanceStorageType
}

func (s *ModifyDBInstanceDiskTypeRequest) GetExtraParam() *string {
	return s.ExtraParam
}

func (s *ModifyDBInstanceDiskTypeRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ModifyDBInstanceDiskTypeRequest) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *ModifyDBInstanceDiskTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceDiskTypeRequest) SetAutoPay(v bool) *ModifyDBInstanceDiskTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeRequest) SetAutoRenew(v string) *ModifyDBInstanceDiskTypeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeRequest) SetBusinessInfo(v string) *ModifyDBInstanceDiskTypeRequest {
	s.BusinessInfo = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeRequest) SetCouponNo(v string) *ModifyDBInstanceDiskTypeRequest {
	s.CouponNo = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeRequest) SetDBInstanceId(v string) *ModifyDBInstanceDiskTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeRequest) SetDbInstanceStorageType(v string) *ModifyDBInstanceDiskTypeRequest {
	s.DbInstanceStorageType = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeRequest) SetExtraParam(v string) *ModifyDBInstanceDiskTypeRequest {
	s.ExtraParam = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeRequest) SetOrderType(v string) *ModifyDBInstanceDiskTypeRequest {
	s.OrderType = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeRequest) SetProvisionedIops(v int64) *ModifyDBInstanceDiskTypeRequest {
	s.ProvisionedIops = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceDiskTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceDiskTypeRequest) Validate() error {
	return dara.Validate(s)
}
