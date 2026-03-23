// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBProxyInstanceRequest
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *ModifyDBProxyInstanceRequest
	GetDBProxyEngineType() *string
	SetDBProxyInstanceNum(v string) *ModifyDBProxyInstanceRequest
	GetDBProxyInstanceNum() *string
	SetDBProxyInstanceType(v string) *ModifyDBProxyInstanceRequest
	GetDBProxyInstanceType() *string
	SetDBProxyNodes(v []*ModifyDBProxyInstanceRequestDBProxyNodes) *ModifyDBProxyInstanceRequest
	GetDBProxyNodes() []*ModifyDBProxyInstanceRequestDBProxyNodes
	SetEffectiveSpecificTime(v string) *ModifyDBProxyInstanceRequest
	GetEffectiveSpecificTime() *string
	SetEffectiveTime(v string) *ModifyDBProxyInstanceRequest
	GetEffectiveTime() *string
	SetMigrateAZ(v []*ModifyDBProxyInstanceRequestMigrateAZ) *ModifyDBProxyInstanceRequest
	GetMigrateAZ() []*ModifyDBProxyInstanceRequestMigrateAZ
	SetOwnerId(v int64) *ModifyDBProxyInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDBProxyInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDBProxyInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBProxyInstanceRequest
	GetResourceOwnerId() *int64
	SetVSwitchIds(v string) *ModifyDBProxyInstanceRequest
	GetVSwitchIds() *string
}

type ModifyDBProxyInstanceRequest struct {
	// This parameter is required.
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// This parameter is required.
	DBProxyInstanceNum *string `json:"DBProxyInstanceNum,omitempty" xml:"DBProxyInstanceNum,omitempty"`
	// This parameter is required.
	DBProxyInstanceType   *string                                     `json:"DBProxyInstanceType,omitempty" xml:"DBProxyInstanceType,omitempty"`
	DBProxyNodes          []*ModifyDBProxyInstanceRequestDBProxyNodes `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty" type:"Repeated"`
	EffectiveSpecificTime *string                                     `json:"EffectiveSpecificTime,omitempty" xml:"EffectiveSpecificTime,omitempty"`
	EffectiveTime         *string                                     `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	MigrateAZ             []*ModifyDBProxyInstanceRequestMigrateAZ    `json:"MigrateAZ,omitempty" xml:"MigrateAZ,omitempty" type:"Repeated"`
	OwnerId               *int64                                      `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount  *string                                     `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64                                      `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VSwitchIds            *string                                     `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
}

func (s ModifyDBProxyInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBProxyInstanceRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *ModifyDBProxyInstanceRequest) GetDBProxyInstanceNum() *string {
	return s.DBProxyInstanceNum
}

func (s *ModifyDBProxyInstanceRequest) GetDBProxyInstanceType() *string {
	return s.DBProxyInstanceType
}

func (s *ModifyDBProxyInstanceRequest) GetDBProxyNodes() []*ModifyDBProxyInstanceRequestDBProxyNodes {
	return s.DBProxyNodes
}

func (s *ModifyDBProxyInstanceRequest) GetEffectiveSpecificTime() *string {
	return s.EffectiveSpecificTime
}

func (s *ModifyDBProxyInstanceRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBProxyInstanceRequest) GetMigrateAZ() []*ModifyDBProxyInstanceRequestMigrateAZ {
	return s.MigrateAZ
}

func (s *ModifyDBProxyInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBProxyInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBProxyInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBProxyInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBProxyInstanceRequest) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ModifyDBProxyInstanceRequest) SetDBInstanceId(v string) *ModifyDBProxyInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetDBProxyEngineType(v string) *ModifyDBProxyInstanceRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetDBProxyInstanceNum(v string) *ModifyDBProxyInstanceRequest {
	s.DBProxyInstanceNum = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetDBProxyInstanceType(v string) *ModifyDBProxyInstanceRequest {
	s.DBProxyInstanceType = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetDBProxyNodes(v []*ModifyDBProxyInstanceRequestDBProxyNodes) *ModifyDBProxyInstanceRequest {
	s.DBProxyNodes = v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetEffectiveSpecificTime(v string) *ModifyDBProxyInstanceRequest {
	s.EffectiveSpecificTime = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetEffectiveTime(v string) *ModifyDBProxyInstanceRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetMigrateAZ(v []*ModifyDBProxyInstanceRequestMigrateAZ) *ModifyDBProxyInstanceRequest {
	s.MigrateAZ = v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetOwnerId(v int64) *ModifyDBProxyInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetRegionId(v string) *ModifyDBProxyInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetResourceOwnerAccount(v string) *ModifyDBProxyInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetResourceOwnerId(v int64) *ModifyDBProxyInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) SetVSwitchIds(v string) *ModifyDBProxyInstanceRequest {
	s.VSwitchIds = &v
	return s
}

func (s *ModifyDBProxyInstanceRequest) Validate() error {
	if s.DBProxyNodes != nil {
		for _, item := range s.DBProxyNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MigrateAZ != nil {
		for _, item := range s.MigrateAZ {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBProxyInstanceRequestDBProxyNodes struct {
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

func (s ModifyDBProxyInstanceRequestDBProxyNodes) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyInstanceRequestDBProxyNodes) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) GetCpuCores() *string {
	return s.CpuCores
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) GetNodeCounts() *string {
	return s.NodeCounts
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) SetCpuCores(v string) *ModifyDBProxyInstanceRequestDBProxyNodes {
	s.CpuCores = &v
	return s
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) SetNodeCounts(v string) *ModifyDBProxyInstanceRequestDBProxyNodes {
	s.NodeCounts = &v
	return s
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) SetZoneId(v string) *ModifyDBProxyInstanceRequestDBProxyNodes {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequestDBProxyNodes) Validate() error {
	return dara.Validate(s)
}

type ModifyDBProxyInstanceRequestMigrateAZ struct {
	DbProxyEndpointId *string `json:"dbProxyEndpointId,omitempty" xml:"dbProxyEndpointId,omitempty"`
	DestVSwitchId     *string `json:"destVSwitchId,omitempty" xml:"destVSwitchId,omitempty"`
	DestVpcId         *string `json:"destVpcId,omitempty" xml:"destVpcId,omitempty"`
}

func (s ModifyDBProxyInstanceRequestMigrateAZ) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyInstanceRequestMigrateAZ) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) GetDbProxyEndpointId() *string {
	return s.DbProxyEndpointId
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) GetDestVSwitchId() *string {
	return s.DestVSwitchId
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) GetDestVpcId() *string {
	return s.DestVpcId
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) SetDbProxyEndpointId(v string) *ModifyDBProxyInstanceRequestMigrateAZ {
	s.DbProxyEndpointId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) SetDestVSwitchId(v string) *ModifyDBProxyInstanceRequestMigrateAZ {
	s.DestVSwitchId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) SetDestVpcId(v string) *ModifyDBProxyInstanceRequestMigrateAZ {
	s.DestVpcId = &v
	return s
}

func (s *ModifyDBProxyInstanceRequestMigrateAZ) Validate() error {
	return dara.Validate(s)
}
