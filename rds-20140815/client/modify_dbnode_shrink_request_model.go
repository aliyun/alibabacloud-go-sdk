// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyDBNodeShrinkRequest
	GetAutoPay() *bool
	SetClientToken(v string) *ModifyDBNodeShrinkRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyDBNodeShrinkRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v string) *ModifyDBNodeShrinkRequest
	GetDBInstanceStorage() *string
	SetDBInstanceStorageType(v string) *ModifyDBNodeShrinkRequest
	GetDBInstanceStorageType() *string
	SetDBNodeShrink(v string) *ModifyDBNodeShrinkRequest
	GetDBNodeShrink() *string
	SetDryRun(v bool) *ModifyDBNodeShrinkRequest
	GetDryRun() *bool
	SetEffectiveTime(v string) *ModifyDBNodeShrinkRequest
	GetEffectiveTime() *string
	SetOwnerAccount(v string) *ModifyDBNodeShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBNodeShrinkRequest
	GetOwnerId() *int64
	SetProduceAsync(v bool) *ModifyDBNodeShrinkRequest
	GetProduceAsync() *bool
	SetResourceOwnerAccount(v string) *ModifyDBNodeShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodeShrinkRequest
	GetResourceOwnerId() *int64
}

type ModifyDBNodeShrinkRequest struct {
	// Specifies whether to automatically pay for the order.
	//
	// Valid values:
	//
	// 1. **true**: Automatic payment is enabled. Make sure that your account balance is sufficient.
	//
	// 1. **false**: Only an order is generated and no payment is made.
	//
	//
	//
	//
	// > Default value: true. If your payment method has an insufficient balance, you can set this parameter to false. In this case, an unpaid order is generated, and you can log on to the ApsaraDB RDS console to complete the payment.
	//
	// >
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1k8s41l2o52****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The new instance storage capacity. Unit: GB. For details, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// example:
	//
	// 20
	DBInstanceStorage *string `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **cloud_essd**: PL1 ESSD.
	//
	// 	- **cloud_essd2**: PL2 ESSD.
	//
	// 	- **cloud_essd3**: PL3 ESSD.
	//
	// example:
	//
	// cloud_essd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The node information.
	//
	// > This parameter applies to ApsaraDB RDS for MySQL Cluster Edition instances.
	DBNodeShrink *string `json:"DBNode,omitempty" xml:"DBNode,omitempty"`
	// Specifies whether to perform a precheck before the node modifications. Valid values:
	//
	// 	- **true**: A precheck is performed, but the modifications are not applied. The precheck includes checks on request parameters, request format, service limits, and inventory.
	//
	// 	- **false**: A normal request is sent, and the modifications are directly applied after the precheck is passed (default).
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The effective period. Valid values:
	//
	// 	- **Immediate*	- (default): The modifications immediately take effect.
	//
	// 	- **MaintainTime**: The modifications take effect within the maintenance window. For more information, see ModifyDBInstanceMaintainTime.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to asynchronously apply the modifications. Valid values:
	//
	// 	- **true**: Only an order is generated, and the modifications are asynchronously applied (default).
	//
	// 	- **false**: The modifications are synchronously applied after the precheck is passed.
	//
	// > Default value: true. The modifications are asynchronously applied. If you set this parameter to false, the modifications are synchronously applied, and the API response time is relatively long.
	//
	// example:
	//
	// true
	ProduceAsync         *bool   `json:"ProduceAsync,omitempty" xml:"ProduceAsync,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBNodeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyDBNodeShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBNodeShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBNodeShrinkRequest) GetDBInstanceStorage() *string {
	return s.DBInstanceStorage
}

func (s *ModifyDBNodeShrinkRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *ModifyDBNodeShrinkRequest) GetDBNodeShrink() *string {
	return s.DBNodeShrink
}

func (s *ModifyDBNodeShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDBNodeShrinkRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBNodeShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBNodeShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBNodeShrinkRequest) GetProduceAsync() *bool {
	return s.ProduceAsync
}

func (s *ModifyDBNodeShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodeShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodeShrinkRequest) SetAutoPay(v bool) *ModifyDBNodeShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetClientToken(v string) *ModifyDBNodeShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetDBInstanceId(v string) *ModifyDBNodeShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetDBInstanceStorage(v string) *ModifyDBNodeShrinkRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetDBInstanceStorageType(v string) *ModifyDBNodeShrinkRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetDBNodeShrink(v string) *ModifyDBNodeShrinkRequest {
	s.DBNodeShrink = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetDryRun(v bool) *ModifyDBNodeShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetEffectiveTime(v string) *ModifyDBNodeShrinkRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetOwnerAccount(v string) *ModifyDBNodeShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetOwnerId(v int64) *ModifyDBNodeShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetProduceAsync(v bool) *ModifyDBNodeShrinkRequest {
	s.ProduceAsync = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetResourceOwnerAccount(v string) *ModifyDBNodeShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) SetResourceOwnerId(v int64) *ModifyDBNodeShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
