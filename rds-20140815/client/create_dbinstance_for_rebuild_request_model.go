// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceForRebuildRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDBInstanceForRebuildRequest
	GetClientToken() *string
	SetDBInstanceDescription(v string) *CreateDBInstanceForRebuildRequest
	GetDBInstanceDescription() *string
	SetDBInstanceId(v string) *CreateDBInstanceForRebuildRequest
	GetDBInstanceId() *string
	SetDBInstanceNetType(v string) *CreateDBInstanceForRebuildRequest
	GetDBInstanceNetType() *string
	SetInstanceNetworkType(v string) *CreateDBInstanceForRebuildRequest
	GetInstanceNetworkType() *string
	SetOwnerAccount(v string) *CreateDBInstanceForRebuildRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBInstanceForRebuildRequest
	GetOwnerId() *int64
	SetPayType(v string) *CreateDBInstanceForRebuildRequest
	GetPayType() *string
	SetPeriod(v string) *CreateDBInstanceForRebuildRequest
	GetPeriod() *string
	SetRegionId(v string) *CreateDBInstanceForRebuildRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBInstanceForRebuildRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateDBInstanceForRebuildRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBInstanceForRebuildRequest
	GetResourceOwnerId() *int64
	SetSecurityIPList(v string) *CreateDBInstanceForRebuildRequest
	GetSecurityIPList() *string
	SetSecurityToken(v string) *CreateDBInstanceForRebuildRequest
	GetSecurityToken() *string
	SetUsedTime(v string) *CreateDBInstanceForRebuildRequest
	GetUsedTime() *string
	SetVPCId(v string) *CreateDBInstanceForRebuildRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBInstanceForRebuildRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateDBInstanceForRebuildRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *CreateDBInstanceForRebuildRequest
	GetZoneIdSlave1() *string
	SetZoneIdSlave2(v string) *CreateDBInstanceForRebuildRequest
	GetZoneIdSlave2() *string
}

type CreateDBInstanceForRebuildRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the instance. The name must be 2 to 256 characters in length. The name can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// > : The name cannot start with http:// or https://.
	//
	// example:
	//
	// Test database
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/610396.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The type of the network connection to the instance. Valid values:
	//
	// 	- **Internet**
	//
	// 	- **Intranet**
	//
	// example:
	//
	// Internet
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **Classic**
	//
	// Default value: Classic.
	//
	// >  If the instance uses cloud disks, this parameter is required. Set the value to **VPC**. The **VpcId*	- and **VSwitchId*	- parameters must be specified when this parameter is set to **VPC**.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the destination instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// >  If you set the PayType parameter to **Prepaid**, you must also specify this parameter.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/610399.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. The value of this parameter can be NULL.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address whitelist of the serverless instance. For more information, see [Use a database client or the CLI to connect to an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/43185.html). If the IP address whitelist contains more than one entry, separate the entries with commas (,). Each entry must be unique. You can specify up to 1,000 entries. The entries in the IP address whitelist must be in one of the following formats:
	//
	// 	- IP addresses, such as 10.10.XX.XX.
	//
	// 	- CIDR blocks, such as 10.10.XX.XX/24. In this example, 24 indicates that the prefix of each IP address in the IP address whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
	//
	// If this parameter is not specified, the default IP address whitelist is used.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The subscription duration of the instance. Valid values:
	//
	// 	- If you set **Period*	- to **Year**, the value of **UsedTime*	- ranges from **1*	- to **5**.
	//
	// 	- If you set **Period*	- to **Month**, the value of **UsedTime*	- ranges from **1*	- to **11**.
	//
	// > If you set **PayType*	- to **Prepaid**, you must specify this parameter.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID of the instance. If you set **InstanceNetworkType*	- to **VPC**, you must specify this parameter.
	//
	// > : If you specify this parameter, you must also specify **ZoneId**.
	//
	// example:
	//
	// vpc-uf6f7l4fg90xxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch. The vSwitch must belong to the zone that is specified by **ZoneId**.
	//
	// >
	//
	// 	- If you set **InstanceNetworkType*	- to **VPC**, you must also specify this parameter.
	//
	// 	- If you specify the ZoneSlaveId1 parameter, you must specify the IDs of two vSwitches for this parameter and separate the IDs with a comma (,).
	//
	// example:
	//
	// vsw-uf6adz52c2pxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the primary instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent zone list.
	//
	// > If you specify a virtual private cloud (VPC) and a vSwitch, you must specify this parameter to identify the zone for the vSwitch.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone ID of the secondary instance.
	//
	// >  If the instance does not run RDS Basic Edition, you must specify this parameter.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// The zone ID of the logger instance.
	//
	// >  This parameter is available only when the instance runs RDS Enterprise Edition.
	//
	// example:
	//
	// cn-hangzhou-d
	ZoneIdSlave2 *string `json:"ZoneIdSlave2,omitempty" xml:"ZoneIdSlave2,omitempty"`
}

func (s CreateDBInstanceForRebuildRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceForRebuildRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceForRebuildRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBInstanceForRebuildRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateDBInstanceForRebuildRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceForRebuildRequest) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *CreateDBInstanceForRebuildRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *CreateDBInstanceForRebuildRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBInstanceForRebuildRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBInstanceForRebuildRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateDBInstanceForRebuildRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateDBInstanceForRebuildRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBInstanceForRebuildRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBInstanceForRebuildRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBInstanceForRebuildRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBInstanceForRebuildRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateDBInstanceForRebuildRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateDBInstanceForRebuildRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateDBInstanceForRebuildRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBInstanceForRebuildRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBInstanceForRebuildRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateDBInstanceForRebuildRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *CreateDBInstanceForRebuildRequest) GetZoneIdSlave2() *string {
	return s.ZoneIdSlave2
}

func (s *CreateDBInstanceForRebuildRequest) SetClientToken(v string) *CreateDBInstanceForRebuildRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetDBInstanceDescription(v string) *CreateDBInstanceForRebuildRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetDBInstanceId(v string) *CreateDBInstanceForRebuildRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetDBInstanceNetType(v string) *CreateDBInstanceForRebuildRequest {
	s.DBInstanceNetType = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetInstanceNetworkType(v string) *CreateDBInstanceForRebuildRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetOwnerAccount(v string) *CreateDBInstanceForRebuildRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetOwnerId(v int64) *CreateDBInstanceForRebuildRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetPayType(v string) *CreateDBInstanceForRebuildRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetPeriod(v string) *CreateDBInstanceForRebuildRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetRegionId(v string) *CreateDBInstanceForRebuildRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetResourceGroupId(v string) *CreateDBInstanceForRebuildRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetResourceOwnerAccount(v string) *CreateDBInstanceForRebuildRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetResourceOwnerId(v int64) *CreateDBInstanceForRebuildRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetSecurityIPList(v string) *CreateDBInstanceForRebuildRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetSecurityToken(v string) *CreateDBInstanceForRebuildRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetUsedTime(v string) *CreateDBInstanceForRebuildRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetVPCId(v string) *CreateDBInstanceForRebuildRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetVSwitchId(v string) *CreateDBInstanceForRebuildRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetZoneId(v string) *CreateDBInstanceForRebuildRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetZoneIdSlave1(v string) *CreateDBInstanceForRebuildRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) SetZoneIdSlave2(v string) *CreateDBInstanceForRebuildRequest {
	s.ZoneIdSlave2 = &v
	return s
}

func (s *CreateDBInstanceForRebuildRequest) Validate() error {
	return dara.Validate(s)
}
