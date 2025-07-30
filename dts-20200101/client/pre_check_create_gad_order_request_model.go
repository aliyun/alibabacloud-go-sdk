// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreCheckCreateGadOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PreCheckCreateGadOrderRequest
	GetInstanceId() *string
	SetMasterDatabaseName(v string) *PreCheckCreateGadOrderRequest
	GetMasterDatabaseName() *string
	SetMasterEngineArchType(v string) *PreCheckCreateGadOrderRequest
	GetMasterEngineArchType() *string
	SetMasterShardAccountName(v string) *PreCheckCreateGadOrderRequest
	GetMasterShardAccountName() *string
	SetMasterShardAccountPassword(v string) *PreCheckCreateGadOrderRequest
	GetMasterShardAccountPassword() *string
	SetOwnerId(v string) *PreCheckCreateGadOrderRequest
	GetOwnerId() *string
	SetRegionId(v string) *PreCheckCreateGadOrderRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *PreCheckCreateGadOrderRequest
	GetResourceGroupId() *string
	SetSlaveDatabaseName(v string) *PreCheckCreateGadOrderRequest
	GetSlaveDatabaseName() *string
	SetSlaveDbInstanceId(v string) *PreCheckCreateGadOrderRequest
	GetSlaveDbInstanceId() *string
	SetSlaveDbInstanceRegion(v string) *PreCheckCreateGadOrderRequest
	GetSlaveDbInstanceRegion() *string
	SetSlaveEngineArchType(v string) *PreCheckCreateGadOrderRequest
	GetSlaveEngineArchType() *string
}

type PreCheckCreateGadOrderRequest struct {
	// example:
	//
	// gad-bp1i99e8l7913****
	InstanceId                 *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MasterDatabaseName         *string `json:"MasterDatabaseName,omitempty" xml:"MasterDatabaseName,omitempty"`
	MasterEngineArchType       *string `json:"MasterEngineArchType,omitempty" xml:"MasterEngineArchType,omitempty"`
	MasterShardAccountName     *string `json:"MasterShardAccountName,omitempty" xml:"MasterShardAccountName,omitempty"`
	MasterShardAccountPassword *string `json:"MasterShardAccountPassword,omitempty" xml:"MasterShardAccountPassword,omitempty"`
	OwnerId                    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfntftbiobqyky
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SlaveDatabaseName *string `json:"SlaveDatabaseName,omitempty" xml:"SlaveDatabaseName,omitempty"`
	// example:
	//
	// rm-bp17562h64****
	SlaveDbInstanceId *string `json:"SlaveDbInstanceId,omitempty" xml:"SlaveDbInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	SlaveDbInstanceRegion *string `json:"SlaveDbInstanceRegion,omitempty" xml:"SlaveDbInstanceRegion,omitempty"`
	SlaveEngineArchType   *string `json:"SlaveEngineArchType,omitempty" xml:"SlaveEngineArchType,omitempty"`
}

func (s PreCheckCreateGadOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s PreCheckCreateGadOrderRequest) GoString() string {
	return s.String()
}

func (s *PreCheckCreateGadOrderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PreCheckCreateGadOrderRequest) GetMasterDatabaseName() *string {
	return s.MasterDatabaseName
}

func (s *PreCheckCreateGadOrderRequest) GetMasterEngineArchType() *string {
	return s.MasterEngineArchType
}

func (s *PreCheckCreateGadOrderRequest) GetMasterShardAccountName() *string {
	return s.MasterShardAccountName
}

func (s *PreCheckCreateGadOrderRequest) GetMasterShardAccountPassword() *string {
	return s.MasterShardAccountPassword
}

func (s *PreCheckCreateGadOrderRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *PreCheckCreateGadOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PreCheckCreateGadOrderRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *PreCheckCreateGadOrderRequest) GetSlaveDatabaseName() *string {
	return s.SlaveDatabaseName
}

func (s *PreCheckCreateGadOrderRequest) GetSlaveDbInstanceId() *string {
	return s.SlaveDbInstanceId
}

func (s *PreCheckCreateGadOrderRequest) GetSlaveDbInstanceRegion() *string {
	return s.SlaveDbInstanceRegion
}

func (s *PreCheckCreateGadOrderRequest) GetSlaveEngineArchType() *string {
	return s.SlaveEngineArchType
}

func (s *PreCheckCreateGadOrderRequest) SetInstanceId(v string) *PreCheckCreateGadOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *PreCheckCreateGadOrderRequest) SetMasterDatabaseName(v string) *PreCheckCreateGadOrderRequest {
	s.MasterDatabaseName = &v
	return s
}

func (s *PreCheckCreateGadOrderRequest) SetMasterEngineArchType(v string) *PreCheckCreateGadOrderRequest {
	s.MasterEngineArchType = &v
	return s
}

func (s *PreCheckCreateGadOrderRequest) SetMasterShardAccountName(v string) *PreCheckCreateGadOrderRequest {
	s.MasterShardAccountName = &v
	return s
}

func (s *PreCheckCreateGadOrderRequest) SetMasterShardAccountPassword(v string) *PreCheckCreateGadOrderRequest {
	s.MasterShardAccountPassword = &v
	return s
}

func (s *PreCheckCreateGadOrderRequest) SetOwnerId(v string) *PreCheckCreateGadOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *PreCheckCreateGadOrderRequest) SetRegionId(v string) *PreCheckCreateGadOrderRequest {
	s.RegionId = &v
	return s
}

func (s *PreCheckCreateGadOrderRequest) SetResourceGroupId(v string) *PreCheckCreateGadOrderRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *PreCheckCreateGadOrderRequest) SetSlaveDatabaseName(v string) *PreCheckCreateGadOrderRequest {
	s.SlaveDatabaseName = &v
	return s
}

func (s *PreCheckCreateGadOrderRequest) SetSlaveDbInstanceId(v string) *PreCheckCreateGadOrderRequest {
	s.SlaveDbInstanceId = &v
	return s
}

func (s *PreCheckCreateGadOrderRequest) SetSlaveDbInstanceRegion(v string) *PreCheckCreateGadOrderRequest {
	s.SlaveDbInstanceRegion = &v
	return s
}

func (s *PreCheckCreateGadOrderRequest) SetSlaveEngineArchType(v string) *PreCheckCreateGadOrderRequest {
	s.SlaveEngineArchType = &v
	return s
}

func (s *PreCheckCreateGadOrderRequest) Validate() error {
	return dara.Validate(s)
}
