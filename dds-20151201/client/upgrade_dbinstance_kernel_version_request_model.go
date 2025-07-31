// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceKernelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBInstanceKernelVersionRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *UpgradeDBInstanceKernelVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpgradeDBInstanceKernelVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest
	GetResourceOwnerId() *int64
	SetSwitchMode(v string) *UpgradeDBInstanceKernelVersionRequest
	GetSwitchMode() *string
}

type UpgradeDBInstanceKernelVersionRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp2235****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s UpgradeDBInstanceKernelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetDBInstanceId(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetOwnerAccount(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetResourceOwnerAccount(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBInstanceKernelVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetSwitchMode(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.SwitchMode = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) Validate() error {
	return dara.Validate(s)
}
