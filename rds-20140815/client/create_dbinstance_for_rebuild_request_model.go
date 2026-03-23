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
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// This parameter is required.
	DBInstanceId        *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceNetType   *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period  *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityIPList       *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	UsedTime             *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VPCId                *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneIdSlave1         *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	ZoneIdSlave2         *string `json:"ZoneIdSlave2,omitempty" xml:"ZoneIdSlave2,omitempty"`
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
