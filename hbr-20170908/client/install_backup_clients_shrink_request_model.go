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
	// The name of the RAM role that is created in the source account for cross-account backup.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// The type of cross-account backup. Valid values:
	//
	// - SELF_ACCOUNT: Backs up data within the current account.
	//
	// - CROSS_ACCOUNT: Backs up data across accounts.
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source account that is used for cross-account backup.
	//
	// example:
	//
	// 16392782xxxxxx
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The IDs of the ECS instances. You can specify a maximum of 20 instance IDs.
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
