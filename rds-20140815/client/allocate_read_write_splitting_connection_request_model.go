// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateReadWriteSplittingConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *AllocateReadWriteSplittingConnectionRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceId(v string) *AllocateReadWriteSplittingConnectionRequest
	GetDBInstanceId() *string
	SetDistributionType(v string) *AllocateReadWriteSplittingConnectionRequest
	GetDistributionType() *string
	SetMaxDelayTime(v string) *AllocateReadWriteSplittingConnectionRequest
	GetMaxDelayTime() *string
	SetNetType(v string) *AllocateReadWriteSplittingConnectionRequest
	GetNetType() *string
	SetOwnerAccount(v string) *AllocateReadWriteSplittingConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateReadWriteSplittingConnectionRequest
	GetOwnerId() *int64
	SetPort(v string) *AllocateReadWriteSplittingConnectionRequest
	GetPort() *string
	SetResourceOwnerAccount(v string) *AllocateReadWriteSplittingConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateReadWriteSplittingConnectionRequest
	GetResourceOwnerId() *int64
	SetWeight(v string) *AllocateReadWriteSplittingConnectionRequest
	GetWeight() *string
}

type AllocateReadWriteSplittingConnectionRequest struct {
	// The prefix of the read-only endpoint. The prefix must be unique, can contain lowercase letters and hyphens (-), must start with a letter, and cannot exceed 30 characters in length.
	//
	// > By default, the prefix is in the format of "instance name + rw".
	//
	// example:
	//
	// rr-m5e****-rw.mysql.rds.aliyuncs.com
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The ID of the primary instance. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The mode of read weight distribution. Valid values:
	//
	// - **Standard**: Read weights are automatically assigned based on instance specifications.
	//
	// - **Custom**: Read weights are manually assigned.
	//
	// example:
	//
	// Standard
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// The latency threshold. Valid values: 0 to 7200. Unit: seconds. Default value: 30.
	//
	// > When the latency of a read-only instance exceeds this threshold, read traffic is not routed to the instance.
	//
	// example:
	//
	// 30
	MaxDelayTime *string `json:"MaxDelayTime,omitempty" xml:"MaxDelayTime,omitempty"`
	// The network type of the read-only endpoint. Valid values:
	//
	// - **Internet**: public endpoint.
	//
	// - **Intranet**: internal endpoint.
	//
	// > The default value is Intranet, and the network type of the internal endpoint is the same as that of the primary instance.
	//
	// example:
	//
	// Intranet
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port of the read-only endpoint. Valid values: 1000 to 5999. Default value: 1433.
	//
	// example:
	//
	// 1433
	Port                 *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The read weight distribution, which specifies the ratio of read requests that are routed to the primary instance and read-only instances. The value is incremented in steps of 100. Maximum value: 10000.
	//
	// 	- Format for ApsaraDB RDS instances: `{"<Read-only instance ID>":<Weight>,"master":<Weight>,"slave":<Weight>}`
	//
	// 	- Format for MyBASE instances: `[{"instanceName":"<Primary instance ID>","weight":<Weight>,"role":"master"},{"instanceName":"<Primary instance ID>","weight":<Weight>,"role":"slave"},{"instanceName":"<Read-only instance ID>","weight":<Weight>,"role":"master"}]`
	//
	// > - This parameter is required when **DistributionType*	- is set to **Custom**.
	//
	// > - This parameter is invalid when **DistributionType*	- is set to **Standard**.
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

func (s AllocateReadWriteSplittingConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateReadWriteSplittingConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetDistributionType() *string {
	return s.DistributionType
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetMaxDelayTime() *string {
	return s.MaxDelayTime
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetNetType() *string {
	return s.NetType
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetPort() *string {
	return s.Port
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateReadWriteSplittingConnectionRequest) GetWeight() *string {
	return s.Weight
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetConnectionStringPrefix(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetDBInstanceId(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetDistributionType(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.DistributionType = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetMaxDelayTime(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.MaxDelayTime = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetNetType(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.NetType = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetOwnerAccount(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetOwnerId(v int64) *AllocateReadWriteSplittingConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetPort(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.Port = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetResourceOwnerAccount(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetResourceOwnerId(v int64) *AllocateReadWriteSplittingConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) SetWeight(v string) *AllocateReadWriteSplittingConnectionRequest {
	s.Weight = &v
	return s
}

func (s *AllocateReadWriteSplittingConnectionRequest) Validate() error {
	return dara.Validate(s)
}
