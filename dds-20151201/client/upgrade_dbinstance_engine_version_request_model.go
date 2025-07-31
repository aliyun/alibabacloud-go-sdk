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
	SetOwnerAccount(v string) *UpgradeDBInstanceEngineVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpgradeDBInstanceEngineVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest
	GetResourceOwnerId() *int64
	SetSwitchMode(v int32) *UpgradeDBInstanceEngineVersionRequest
	GetSwitchMode() *int32
}

type UpgradeDBInstanceEngineVersionRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database version to which you want to upgrade. Valid values: **3.4**, **4.0**, and **4.2**.
	//
	// >  This database version must be later than the current database version of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4.0
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The time when to perform the upgrade. Valid values:
	//
	// 	- **0**: immediately performs the upgrade.
	//
	// 	- **1**: performs the upgrade during the maintenance window.
	//
	// example:
	//
	// 1
	SwitchMode *int32 `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
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

func (s *UpgradeDBInstanceEngineVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBInstanceEngineVersionRequest) GetSwitchMode() *int32 {
	return s.SwitchMode
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetDBInstanceId(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetEngineVersion(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.EngineVersion = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetOwnerAccount(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetResourceOwnerAccount(v string) *UpgradeDBInstanceEngineVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBInstanceEngineVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) SetSwitchMode(v int32) *UpgradeDBInstanceEngineVersionRequest {
	s.SwitchMode = &v
	return s
}

func (s *UpgradeDBInstanceEngineVersionRequest) Validate() error {
	return dara.Validate(s)
}
