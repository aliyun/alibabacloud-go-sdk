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
	// Specifies whether to automatically complete the payment. Valid values:
	//
	// 1.  **true**: automatically completes the payment. Make sure that your account balance is sufficient.
	//
	// 2.  **false**: does not automatically complete the payment. An unpaid order is generated.
	//
	// >  The default value is true. If your account balance is insufficient, you can set the AutoPay parameter to false to generate an unpaid order. Then, you can log on to the ApsaraDB RDS console to pay for the order.
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
	// The new storage capacity of the instance. Unit: GB For more information, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// example:
	//
	// 20
	DBInstanceStorage *string `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **cloud_essd**: performance level 1 (PL1) enhanced SSD (ESSD)
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// example:
	//
	// cloud_essd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The information about the node.
	//
	// >  This parameter is used for ApsaraDB RDS for MySQL instances that run RDS Cluster Edition.
	DBNodeShrink *string `json:"DBNode,omitempty" xml:"DBNode,omitempty"`
	// Specifies whether to perform a dry run. Valid values: Valid values:
	//
	// 	- **true**: performs a dry run and does not perform the actual request. The system checks items such as the request parameters, request format, service limits, and available resources.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The time when you want the change to take effect. Valid values:
	//
	// 	- **Immediate*	- (default): The change immediately takes effect.
	//
	// 	- **MaintainTime**: The effective time is within the maintenance window. For more information, see ModifyDBInstanceMaintainTime.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to asynchronously perform the operation. Valid values:
	//
	// 	- **true*	- (default): sends only the order. The operation is asynchronously performed.
	//
	// 	- **false**: sends the request. After the request passes the check, the operation is directly performed.
	//
	// >  The default value is true, which indicates that the change operation is asynchronously performed. If you set this parameter to false, the change operation is simultaneously performed. This prolongs the response time of the operation.
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
