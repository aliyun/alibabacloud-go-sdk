// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigDBProxyService(v string) *ModifyDBProxyRequest
	GetConfigDBProxyService() *string
	SetDBInstanceId(v string) *ModifyDBProxyRequest
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *ModifyDBProxyRequest
	GetDBProxyEngineType() *string
	SetDBProxyInstanceNum(v string) *ModifyDBProxyRequest
	GetDBProxyInstanceNum() *string
	SetDBProxyInstanceType(v string) *ModifyDBProxyRequest
	GetDBProxyInstanceType() *string
	SetDBProxyNodes(v []*ModifyDBProxyRequestDBProxyNodes) *ModifyDBProxyRequest
	GetDBProxyNodes() []*ModifyDBProxyRequestDBProxyNodes
	SetInstanceNetworkType(v string) *ModifyDBProxyRequest
	GetInstanceNetworkType() *string
	SetOwnerId(v int64) *ModifyDBProxyRequest
	GetOwnerId() *int64
	SetPersistentConnectionStatus(v string) *ModifyDBProxyRequest
	GetPersistentConnectionStatus() *string
	SetRegionId(v string) *ModifyDBProxyRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDBProxyRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBProxyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBProxyRequest
	GetResourceOwnerId() *int64
	SetVPCId(v string) *ModifyDBProxyRequest
	GetVPCId() *string
	SetVSwitchId(v string) *ModifyDBProxyRequest
	GetVSwitchId() *string
}

type ModifyDBProxyRequest struct {
	// This parameter is required.
	ConfigDBProxyService *string `json:"ConfigDBProxyService,omitempty" xml:"ConfigDBProxyService,omitempty"`
	// This parameter is required.
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBProxyEngineType  *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	DBProxyInstanceNum *string `json:"DBProxyInstanceNum,omitempty" xml:"DBProxyInstanceNum,omitempty"`
	// example:
	//
	// common
	DBProxyInstanceType *string                             `json:"DBProxyInstanceType,omitempty" xml:"DBProxyInstanceType,omitempty"`
	DBProxyNodes        []*ModifyDBProxyRequestDBProxyNodes `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty" type:"Repeated"`
	InstanceNetworkType *string                             `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerId             *int64                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s ModifyDBProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyRequest) GetConfigDBProxyService() *string {
	return s.ConfigDBProxyService
}

func (s *ModifyDBProxyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBProxyRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *ModifyDBProxyRequest) GetDBProxyInstanceNum() *string {
	return s.DBProxyInstanceNum
}

func (s *ModifyDBProxyRequest) GetDBProxyInstanceType() *string {
	return s.DBProxyInstanceType
}

func (s *ModifyDBProxyRequest) GetDBProxyNodes() []*ModifyDBProxyRequestDBProxyNodes {
	return s.DBProxyNodes
}

func (s *ModifyDBProxyRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *ModifyDBProxyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBProxyRequest) GetPersistentConnectionStatus() *string {
	return s.PersistentConnectionStatus
}

func (s *ModifyDBProxyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBProxyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBProxyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBProxyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBProxyRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *ModifyDBProxyRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBProxyRequest) SetConfigDBProxyService(v string) *ModifyDBProxyRequest {
	s.ConfigDBProxyService = &v
	return s
}

func (s *ModifyDBProxyRequest) SetDBInstanceId(v string) *ModifyDBProxyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBProxyRequest) SetDBProxyEngineType(v string) *ModifyDBProxyRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *ModifyDBProxyRequest) SetDBProxyInstanceNum(v string) *ModifyDBProxyRequest {
	s.DBProxyInstanceNum = &v
	return s
}

func (s *ModifyDBProxyRequest) SetDBProxyInstanceType(v string) *ModifyDBProxyRequest {
	s.DBProxyInstanceType = &v
	return s
}

func (s *ModifyDBProxyRequest) SetDBProxyNodes(v []*ModifyDBProxyRequestDBProxyNodes) *ModifyDBProxyRequest {
	s.DBProxyNodes = v
	return s
}

func (s *ModifyDBProxyRequest) SetInstanceNetworkType(v string) *ModifyDBProxyRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *ModifyDBProxyRequest) SetOwnerId(v int64) *ModifyDBProxyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBProxyRequest) SetPersistentConnectionStatus(v string) *ModifyDBProxyRequest {
	s.PersistentConnectionStatus = &v
	return s
}

func (s *ModifyDBProxyRequest) SetRegionId(v string) *ModifyDBProxyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBProxyRequest) SetResourceGroupId(v string) *ModifyDBProxyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBProxyRequest) SetResourceOwnerAccount(v string) *ModifyDBProxyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBProxyRequest) SetResourceOwnerId(v int64) *ModifyDBProxyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBProxyRequest) SetVPCId(v string) *ModifyDBProxyRequest {
	s.VPCId = &v
	return s
}

func (s *ModifyDBProxyRequest) SetVSwitchId(v string) *ModifyDBProxyRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBProxyRequest) Validate() error {
	if s.DBProxyNodes != nil {
		for _, item := range s.DBProxyNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBProxyRequestDBProxyNodes struct {
	// example:
	//
	// 1
	CpuCores *string `json:"cpuCores,omitempty" xml:"cpuCores,omitempty"`
	// example:
	//
	// 2
	NodeCounts *string `json:"nodeCounts,omitempty" xml:"nodeCounts,omitempty"`
	// example:
	//
	// cn-hagnzhou-c
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ModifyDBProxyRequestDBProxyNodes) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyRequestDBProxyNodes) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyRequestDBProxyNodes) GetCpuCores() *string {
	return s.CpuCores
}

func (s *ModifyDBProxyRequestDBProxyNodes) GetNodeCounts() *string {
	return s.NodeCounts
}

func (s *ModifyDBProxyRequestDBProxyNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBProxyRequestDBProxyNodes) SetCpuCores(v string) *ModifyDBProxyRequestDBProxyNodes {
	s.CpuCores = &v
	return s
}

func (s *ModifyDBProxyRequestDBProxyNodes) SetNodeCounts(v string) *ModifyDBProxyRequestDBProxyNodes {
	s.NodeCounts = &v
	return s
}

func (s *ModifyDBProxyRequestDBProxyNodes) SetZoneId(v string) *ModifyDBProxyRequestDBProxyNodes {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBProxyRequestDBProxyNodes) Validate() error {
	return dara.Validate(s)
}
