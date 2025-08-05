// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallBackupClientsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccountRoleName(v string) *InstallBackupClientsShrinkRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *InstallBackupClientsShrinkRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *InstallBackupClientsShrinkRequest
	GetCrossAccountUserId() *int64
	SetInstanceIdsShrink(v string) *InstallBackupClientsShrinkRequest
	GetInstanceIdsShrink() *string
}

type InstallBackupClientsShrinkRequest struct {
	// The name of the RAM role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Specifies whether data is backed up within the same Alibaba Cloud account or across Alibaba Cloud accounts. Valid values:
	//
	// 	- SELF_ACCOUNT: Data is backed up within the same Alibaba Cloud account.
	//
	// 	- CROSS_ACCOUNT: Data is backed up across Alibaba Cloud accounts.
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 16392782xxxxxx
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The IDs of the ECS instances. You can specify up to 20 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["i-0xi5wj5*****v3j3bh2gj5"]
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s InstallBackupClientsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallBackupClientsShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallBackupClientsShrinkRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *InstallBackupClientsShrinkRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *InstallBackupClientsShrinkRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *InstallBackupClientsShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *InstallBackupClientsShrinkRequest) SetCrossAccountRoleName(v string) *InstallBackupClientsShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *InstallBackupClientsShrinkRequest) SetCrossAccountType(v string) *InstallBackupClientsShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *InstallBackupClientsShrinkRequest) SetCrossAccountUserId(v int64) *InstallBackupClientsShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *InstallBackupClientsShrinkRequest) SetInstanceIdsShrink(v string) *InstallBackupClientsShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *InstallBackupClientsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
