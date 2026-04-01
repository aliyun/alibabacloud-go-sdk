// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReadWriteSplittingConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *ModifyReadWriteSplittingConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceId(v string) *ModifyReadWriteSplittingConnectionRequest
	GetDBInstanceId() *string
	SetDistributionType(v string) *ModifyReadWriteSplittingConnectionRequest
	GetDistributionType() *string
	SetMaxDelayTime(v string) *ModifyReadWriteSplittingConnectionRequest
	GetMaxDelayTime() *string
	SetOwnerAccount(v string) *ModifyReadWriteSplittingConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyReadWriteSplittingConnectionRequest
	GetOwnerId() *int64
	SetPort(v string) *ModifyReadWriteSplittingConnectionRequest
	GetPort() *string
	SetResourceOwnerAccount(v string) *ModifyReadWriteSplittingConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyReadWriteSplittingConnectionRequest
	GetResourceOwnerId() *int64
	SetWeight(v string) *ModifyReadWriteSplittingConnectionRequest
	GetWeight() *string
}

type ModifyReadWriteSplittingConnectionRequest struct {
	// The prefix of the read/write splitting endpoint. The prefix must be unique. It can be up to 30 characters in length and can contain lowercase letters and hyphens (-). It must start with a lowercase letter.
	//
	// > The default prefix consists of the name of the primary instance followed by the letters rw.
	//
	// example:
	//
	// rm-m5xxxxxxxxrw.mysql.rds.aliyuncs.com
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The ID of the primary instance. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The method that is used to assign read weights. Valid values:
	//
	// 	- **Standard**: The system automatically assigns read weights to the primary and read-only instances based on the specifications of these instances.
	//
	// 	- **Custom**: You must manually assign a read weight to each instance.
	//
	// > You must specify at least one of **MaxDelayTime*	- and **DistributionType**.
	//
	// example:
	//
	// Standard
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// The latency threshold that is allowed by the read/write splitting link. Unit: seconds. If the latency on a read-only instance exceeds the specified threshold, the system no longer routes read requests to the read-only instance. If you do not specify this parameter, the default value of this parameter is retained.
	//
	// > 	- If the primary instance runs SQL Server 2017 on RDS Cluster Edition, the **MaxDelayTime*	- parameter is not supported.
	//
	// > 	- You must specify at least one of **MaxDelayTime*	- and **DistributionType**.
	//
	// example:
	//
	// 12
	MaxDelayTime *string `json:"MaxDelayTime,omitempty" xml:"MaxDelayTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port that is associated with the read/write splitting endpoint.
	//
	// example:
	//
	// 3306
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The read weights of the primary instance and its read-only instances. A read weight must be a multiple of 100 and cannot exceed 10,000.
	//
	// 	- For ApsaraDB RDS instances, the value of this parameter is in the following format: `{"<ID of the read-only instance >":<Weight>,"master":<Weight>,"slave":<Weight>}`.
	//
	// 	- For ApsaraDB MyBase instances, the value of this parameter is in the following format: `[{"instanceName":"<ID of the primary instance>","weight":<Weight>,"role":"master"},{"instanceName":"<ID of the primary instance>","weight":<Weight>,"role":"slave"},{"instanceName":"<ID of the read-only instance>","weight":<Weight>,"role":"master"}]`
	//
	// >
	//
	// 	- This parameter must be specified when **DistributionType*	- is set to **Custom**.
	//
	// 	- If **DistributionType*	- is set to **Standard**, this parameter is invalid.
	//
	// example:
	//
	// {"rm-bp1**********":800,"master":400,"slave":400}
	Weight *string `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ModifyReadWriteSplittingConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyReadWriteSplittingConnectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetDistributionType() *string {
	return s.DistributionType
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetMaxDelayTime() *string {
	return s.MaxDelayTime
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyReadWriteSplittingConnectionRequest) GetWeight() *string {
	return s.Weight
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetConnectionStringPrefix(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetDBInstanceId(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetDistributionType(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.DistributionType = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetMaxDelayTime(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.MaxDelayTime = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetOwnerAccount(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetOwnerId(v int64) *ModifyReadWriteSplittingConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetPort(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.Port = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetResourceOwnerAccount(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetResourceOwnerId(v int64) *ModifyReadWriteSplittingConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) SetWeight(v string) *ModifyReadWriteSplittingConnectionRequest {
	s.Weight = &v
	return s
}

func (s *ModifyReadWriteSplittingConnectionRequest) Validate() error {
	return dara.Validate(s)
}
