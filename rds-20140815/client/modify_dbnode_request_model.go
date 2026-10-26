// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyDBNodeRequest
	GetAutoPay() *bool
	SetClientToken(v string) *ModifyDBNodeRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyDBNodeRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v string) *ModifyDBNodeRequest
	GetDBInstanceStorage() *string
	SetDBInstanceStorageType(v string) *ModifyDBNodeRequest
	GetDBInstanceStorageType() *string
	SetDBNode(v []*ModifyDBNodeRequestDBNode) *ModifyDBNodeRequest
	GetDBNode() []*ModifyDBNodeRequestDBNode
	SetDryRun(v bool) *ModifyDBNodeRequest
	GetDryRun() *bool
	SetEffectiveTime(v string) *ModifyDBNodeRequest
	GetEffectiveTime() *string
	SetOwnerAccount(v string) *ModifyDBNodeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBNodeRequest
	GetOwnerId() *int64
	SetProduceAsync(v bool) *ModifyDBNodeRequest
	GetProduceAsync() *bool
	SetResourceOwnerAccount(v string) *ModifyDBNodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodeRequest
	GetResourceOwnerId() *int64
}

type ModifyDBNodeRequest struct {
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
	DBNode []*ModifyDBNodeRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
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

func (s ModifyDBNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyDBNodeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBNodeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBNodeRequest) GetDBInstanceStorage() *string {
	return s.DBInstanceStorage
}

func (s *ModifyDBNodeRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *ModifyDBNodeRequest) GetDBNode() []*ModifyDBNodeRequestDBNode {
	return s.DBNode
}

func (s *ModifyDBNodeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDBNodeRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBNodeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBNodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBNodeRequest) GetProduceAsync() *bool {
	return s.ProduceAsync
}

func (s *ModifyDBNodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodeRequest) SetAutoPay(v bool) *ModifyDBNodeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDBNodeRequest) SetClientToken(v string) *ModifyDBNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBNodeRequest) SetDBInstanceId(v string) *ModifyDBNodeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBNodeRequest) SetDBInstanceStorage(v string) *ModifyDBNodeRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *ModifyDBNodeRequest) SetDBInstanceStorageType(v string) *ModifyDBNodeRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *ModifyDBNodeRequest) SetDBNode(v []*ModifyDBNodeRequestDBNode) *ModifyDBNodeRequest {
	s.DBNode = v
	return s
}

func (s *ModifyDBNodeRequest) SetDryRun(v bool) *ModifyDBNodeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDBNodeRequest) SetEffectiveTime(v string) *ModifyDBNodeRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBNodeRequest) SetOwnerAccount(v string) *ModifyDBNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodeRequest) SetOwnerId(v int64) *ModifyDBNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodeRequest) SetProduceAsync(v bool) *ModifyDBNodeRequest {
	s.ProduceAsync = &v
	return s
}

func (s *ModifyDBNodeRequest) SetResourceOwnerAccount(v string) *ModifyDBNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodeRequest) SetResourceOwnerId(v int64) *ModifyDBNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodeRequest) Validate() error {
	if s.DBNode != nil {
		for _, item := range s.DBNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBNodeRequestDBNode struct {
	// The node specifications.
	//
	// example:
	//
	// mysql.n2.medium.xc
	ClassCode *string `json:"classCode,omitempty" xml:"classCode,omitempty"`
	// The node ID.
	//
	// example:
	//
	// rn-6256r4a87xvv7****
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
}

func (s ModifyDBNodeRequestDBNode) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeRequestDBNode) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeRequestDBNode) GetClassCode() *string {
	return s.ClassCode
}

func (s *ModifyDBNodeRequestDBNode) GetNodeId() *string {
	return s.NodeId
}

func (s *ModifyDBNodeRequestDBNode) SetClassCode(v string) *ModifyDBNodeRequestDBNode {
	s.ClassCode = &v
	return s
}

func (s *ModifyDBNodeRequestDBNode) SetNodeId(v string) *ModifyDBNodeRequestDBNode {
	s.NodeId = &v
	return s
}

func (s *ModifyDBNodeRequestDBNode) Validate() error {
	return dara.Validate(s)
}
