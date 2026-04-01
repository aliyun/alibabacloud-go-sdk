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
	DBNode []*ModifyDBNodeRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
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
	// The specification information about the node.
	//
	// example:
	//
	// mysql.n2.medium.xc
	ClassCode *string `json:"classCode,omitempty" xml:"classCode,omitempty"`
	// The node ID.
	//
	// example:
	//
	// rn-6256r4a87xvv7he5p
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
