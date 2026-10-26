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
	// ETnLKlblzczshOTUbOC****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the destination instance. The name must be 2 to 256 characters in length. It must start with a Chinese or English letter and can contain digits, Chinese characters, English letters, underscores (_), and hyphens (-).
	//
	// >The name cannot start with http:// or https://.
	//
	// example:
	//
	// testdb
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network connectivity type of the instance. Valid values:
	//
	// 	- **Internet**: public network connection.
	//
	// 	- **Intranet**: internal network connection.
	//
	// example:
	//
	// Internet
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The network type of the destination instance. Valid values:
	//
	// 	- **VPC**: VPC.
	//
	// 	- **Classic**: classic network.
	//
	// By default, an instance in the classic network is created.
	//
	// >If the instance is a cloud disk instance, you must set this parameter to **VPC**. If you set this parameter to **VPC**, you must also specify the **VpcId*	- and **VSwitchId*	- parameters.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the destination instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription period of the subscription instance. Valid values:
	//
	// 	- **Year**: yearly subscription.
	//
	// 	- **Month**: monthly subscription.
	//
	// >You must specify this parameter when **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. This parameter can be left empty.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The [IP address whitelist](https://help.aliyun.com/document_detail/43185.html) of the instance. Separate multiple entries with commas (,). Duplicate entries are not allowed. A maximum of 1,000 entries are supported. The following formats are supported:
	//
	// 	- IP address format, such as 10.10.XX.XX.
	//
	// 	- CIDR format, such as 10.10.XX.XX/24. In this format, 24 indicates the length of the prefix in the CIDR block, and the length ranges from 1 to 32.
	//
	// If you do not specify this parameter, the default whitelist of the original instance is used.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The subscription duration. Valid values:
	//
	// 	- If the **Period*	- parameter is set to **Year**, the value of the **UsedTime*	- parameter ranges from **1*	- to **5**.
	//
	// 	- If the **Period*	- parameter is set to **Month**, the value of the **UsedTime*	- parameter ranges from **1*	- to **11**.
	//
	// >You must specify this parameter when **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID of the destination instance. You must specify this parameter when **InstanceNetworkType*	- is set to **VPC**.
	//
	// >If you specify this parameter, you must also specify the **ZoneId*	- parameter.
	//
	// example:
	//
	// vpc-uf6f7l4fg90****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID. The zone where the vSwitch is located must match the zone ID specified by the **ZoneId*	- parameter.
	//
	// > - The **InstanceNetworkType*	- parameter must be set to **VPC**.
	//
	// > - If you specify the ZoneSlaveId1 (secondary zone ID) parameter, you must specify two vSwitch IDs and separate them with a comma (,).
	//
	// example:
	//
	// vsw-uf6adz52c2p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the primary node. You can call the DescribeRegions operation to query the zone ID.
	//
	// >If you specify a VPC and a vSwitch, you must specify this parameter to match the zone of the vSwitch.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The secondary zone 1.
	//
	// >You must specify this parameter for instances that are not of the Basic Edition.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// The secondary zone 2.
	//
	// >You can specify this parameter only for instances of the Enterprise Edition.
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
