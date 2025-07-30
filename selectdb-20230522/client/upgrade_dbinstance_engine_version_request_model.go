// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceEngineVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBInstanceEngineVersionRequest
	GetDBInstanceId() *string
	SetEngineVersion(v string) *UpgradeDBInstanceEngineVersionRequest
	GetEngineVersion() *string
	SetParallelOperation(v bool) *UpgradeDBInstanceEngineVersionRequest
	GetParallelOperation() *bool
	SetRegionId(v string) *UpgradeDBInstanceEngineVersionRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest
	GetResourceOwnerId() *int64
	SetSwitchTimeMode(v string) *UpgradeDBInstanceEngineVersionRequest
	GetSwitchTimeMode() *string
}

type UpgradeDBInstanceEngineVersionRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database engine version of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.4
	EngineVersion     *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ParallelOperation *bool   `json:"ParallelOperation,omitempty" xml:"ParallelOperation,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The update mode. If you do not specify this parameter, the system immediately updates the database engine version. If you set this parameter to 1, the system updates the database engine version during the maintenance window.
	//
	// example:
	//
	// 1
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s UpgradeDBInstanceEngineVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceEngineVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetParallelOperation() *bool {
	return s.ParallelOperation
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetDBInstanceId(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetEngineVersion(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.EngineVersion = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetParallelOperation(v bool) *UpgradeDBInstanceEngineVersionRequest {
	s.ParallelOperation = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetRegionId(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetSwitchTimeMode(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) Validate() error {
	return dara.Validate(s)
}
