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
	// The prefix of the read/write splitting endpoint. The prefix must be unique and can contain only lowercase letters and hyphens. It must start with a letter and be up to 30 characters in length.
	//
	// > By default, the prefix consists of the instance name and the "rw" string.
	//
	// example:
	//
	// rm-m5****rw.mysql.rds.aliyuncs.com
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// Instance ID of the primary instance. You can call the DescribeDBInstances operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The read weight distribution mode. Valid values:
	//
	// 	- **Standard**: automatically distributes weights based on instance specifications.
	//
	// 	- **Custom**: uses custom weights.
	//
	// > You must specify at least one of the **MaxDelayTime*	- and **DistributionType*	- parameters.
	//
	// example:
	//
	// Standard
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// The latency threshold, in seconds. If the latency of a read-only instance exceeds this threshold, read traffic is not routed to the instance. If you do not specify this parameter, the original value is retained.
	//
	// > 	- The **MaxDelayTime*	- parameter does not apply to SQL Server 2017 Cluster Edition instances.
	//
	// > 	- You must specify at least one of the **MaxDelayTime*	- and **DistributionType*	- parameters.
	//
	// example:
	//
	// 12
	MaxDelayTime *string `json:"MaxDelayTime,omitempty" xml:"MaxDelayTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port number of the read/write splitting endpoint.
	//
	// example:
	//
	// 3306
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The read weights of the primary and read-only instances. The weight value must be a multiple of 100 and can be up to 10000.
	//
	// 	- Format for ApsaraDB RDS instances: `{"<Read-only instance ID>":<Weight>,"master":<Weight>,"slave":<Weight>}`
	//
	// 	- Format for MyBase instances: `[{"instanceName":"<Primary instance ID>","weight":<Weight>,"role":"master"},{"instanceName":"<Primary instance ID>","weight":<Weight>,"role":"slave"},{"instanceName":"<Read-only instance ID>","weight":<Weight>,"role":"master"}]`
	//
	// > 	- This parameter is required when **DistributionType*	- is set to **Custom**.
	//
	// > 	- This parameter does not take effect when **DistributionType*	- is set to **Standard**.
	//
	// example:
	//
	// {
	//
	//       "rm-bp1****": 800,
	//
	//       "master": 400,
	//
	//       "slave": 400
	//
	// }
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
