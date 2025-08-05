// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallBackupClientsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIdsShrink(v string) *UninstallBackupClientsShrinkRequest
	GetClientIdsShrink() *string
	SetCrossAccountRoleName(v string) *UninstallBackupClientsShrinkRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *UninstallBackupClientsShrinkRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *UninstallBackupClientsShrinkRequest
	GetCrossAccountUserId() *int64
	SetInstanceIdsShrink(v string) *UninstallBackupClientsShrinkRequest
	GetInstanceIdsShrink() *string
}

type UninstallBackupClientsShrinkRequest struct {
	// The IDs of Cloud Backup clients. The sum of the number of Cloud Backup client IDs and the number of ECS instance IDs cannot exceed 20. Otherwise, an error occurs.
	//
	// example:
	//
	// ["c-*********************"]
	ClientIdsShrink *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The name of the RAM role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up and restore data across Alibaba Cloud accounts.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Specifies whether data is backed up and restored within the same Alibaba Cloud account or across Alibaba Cloud accounts. Valid values:
	//
	// 	- SELF_ACCOUNT: Data is backed up and restored within the same Alibaba Cloud account.
	//
	// 	- CROSS_ACCOUNT: Data is backed up and restored across Alibaba Cloud accounts.
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up and restore data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 129349237xxxxx
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The IDs of Elastic Compute Service (ECS) instances. You can specify a maximum of 20 ECS instances.
	//
	// example:
	//
	// ["i-0xi5wj5*****v3j3bh2gj5"]
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s UninstallBackupClientsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallBackupClientsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientsShrinkRequest) GetClientIdsShrink() *string {
	return s.ClientIdsShrink
}

func (s *UninstallBackupClientsShrinkRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *UninstallBackupClientsShrinkRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *UninstallBackupClientsShrinkRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *UninstallBackupClientsShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *UninstallBackupClientsShrinkRequest) SetClientIdsShrink(v string) *UninstallBackupClientsShrinkRequest {
	s.ClientIdsShrink = &v
	return s
}

func (s *UninstallBackupClientsShrinkRequest) SetCrossAccountRoleName(v string) *UninstallBackupClientsShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *UninstallBackupClientsShrinkRequest) SetCrossAccountType(v string) *UninstallBackupClientsShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *UninstallBackupClientsShrinkRequest) SetCrossAccountUserId(v int64) *UninstallBackupClientsShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *UninstallBackupClientsShrinkRequest) SetInstanceIdsShrink(v string) *UninstallBackupClientsShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *UninstallBackupClientsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
