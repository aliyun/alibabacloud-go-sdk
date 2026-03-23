// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigDBProxyService(v string) *ModifyDBProxyShrinkRequest
	GetConfigDBProxyService() *string
	SetDBInstanceId(v string) *ModifyDBProxyShrinkRequest
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *ModifyDBProxyShrinkRequest
	GetDBProxyEngineType() *string
	SetDBProxyInstanceNum(v string) *ModifyDBProxyShrinkRequest
	GetDBProxyInstanceNum() *string
	SetDBProxyInstanceType(v string) *ModifyDBProxyShrinkRequest
	GetDBProxyInstanceType() *string
	SetDBProxyNodesShrink(v string) *ModifyDBProxyShrinkRequest
	GetDBProxyNodesShrink() *string
	SetInstanceNetworkType(v string) *ModifyDBProxyShrinkRequest
	GetInstanceNetworkType() *string
	SetOwnerId(v int64) *ModifyDBProxyShrinkRequest
	GetOwnerId() *int64
	SetPersistentConnectionStatus(v string) *ModifyDBProxyShrinkRequest
	GetPersistentConnectionStatus() *string
	SetRegionId(v string) *ModifyDBProxyShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDBProxyShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBProxyShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBProxyShrinkRequest
	GetResourceOwnerId() *int64
	SetVPCId(v string) *ModifyDBProxyShrinkRequest
	GetVPCId() *string
	SetVSwitchId(v string) *ModifyDBProxyShrinkRequest
	GetVSwitchId() *string
}

type ModifyDBProxyShrinkRequest struct {
	// This parameter is required.
	ConfigDBProxyService *string `json:"ConfigDBProxyService,omitempty" xml:"ConfigDBProxyService,omitempty"`
	// This parameter is required.
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBProxyEngineType  *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	DBProxyInstanceNum *string `json:"DBProxyInstanceNum,omitempty" xml:"DBProxyInstanceNum,omitempty"`
	// example:
	//
	// common
	DBProxyInstanceType *string `json:"DBProxyInstanceType,omitempty" xml:"DBProxyInstanceType,omitempty"`
	DBProxyNodesShrink  *string `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty"`
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// Enabled
	PersistentConnectionStatus *string `json:"PersistentConnectionStatus,omitempty" xml:"PersistentConnectionStatus,omitempty"`
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId            *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount       *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId            *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VPCId                      *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId                  *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ModifyDBProxyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyShrinkRequest) GetConfigDBProxyService() *string {
	return s.ConfigDBProxyService
}

func (s *ModifyDBProxyShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBProxyShrinkRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *ModifyDBProxyShrinkRequest) GetDBProxyInstanceNum() *string {
	return s.DBProxyInstanceNum
}

func (s *ModifyDBProxyShrinkRequest) GetDBProxyInstanceType() *string {
	return s.DBProxyInstanceType
}

func (s *ModifyDBProxyShrinkRequest) GetDBProxyNodesShrink() *string {
	return s.DBProxyNodesShrink
}

func (s *ModifyDBProxyShrinkRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *ModifyDBProxyShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBProxyShrinkRequest) GetPersistentConnectionStatus() *string {
	return s.PersistentConnectionStatus
}

func (s *ModifyDBProxyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBProxyShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBProxyShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBProxyShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBProxyShrinkRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *ModifyDBProxyShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBProxyShrinkRequest) SetConfigDBProxyService(v string) *ModifyDBProxyShrinkRequest {
	s.ConfigDBProxyService = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetDBInstanceId(v string) *ModifyDBProxyShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetDBProxyEngineType(v string) *ModifyDBProxyShrinkRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetDBProxyInstanceNum(v string) *ModifyDBProxyShrinkRequest {
	s.DBProxyInstanceNum = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetDBProxyInstanceType(v string) *ModifyDBProxyShrinkRequest {
	s.DBProxyInstanceType = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetDBProxyNodesShrink(v string) *ModifyDBProxyShrinkRequest {
	s.DBProxyNodesShrink = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetInstanceNetworkType(v string) *ModifyDBProxyShrinkRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetOwnerId(v int64) *ModifyDBProxyShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetPersistentConnectionStatus(v string) *ModifyDBProxyShrinkRequest {
	s.PersistentConnectionStatus = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetRegionId(v string) *ModifyDBProxyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetResourceGroupId(v string) *ModifyDBProxyShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetResourceOwnerAccount(v string) *ModifyDBProxyShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetResourceOwnerId(v int64) *ModifyDBProxyShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetVPCId(v string) *ModifyDBProxyShrinkRequest {
	s.VPCId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) SetVSwitchId(v string) *ModifyDBProxyShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBProxyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
