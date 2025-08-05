// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeBackupClientsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIdsShrink(v string) *UpgradeBackupClientsShrinkRequest
	GetClientIdsShrink() *string
	SetCrossAccountRoleName(v string) *UpgradeBackupClientsShrinkRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *UpgradeBackupClientsShrinkRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *UpgradeBackupClientsShrinkRequest
	GetCrossAccountUserId() *int64
	SetInstanceIdsShrink(v string) *UpgradeBackupClientsShrinkRequest
	GetInstanceIdsShrink() *string
}

type UpgradeBackupClientsShrinkRequest struct {
	// The IDs of Cloud Backup clients. The total number of Cloud Backup client IDs and ECS instance IDs cannot exceed 100.
	//
	// example:
	//
	// ["i-0xi5wj******3j3bh2gj5"]
	ClientIdsShrink *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The name of the RAM role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// hbrcrossrole
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
	// 1283948272xxxxx
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The IDs of Elastic Compute Service (ECS) instances. The total number of ECS instance IDs and Cloud Backup client IDs cannot exceed 100.
	//
	// example:
	//
	// ["c-*********************"]
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s UpgradeBackupClientsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeBackupClientsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradeBackupClientsShrinkRequest) GetClientIdsShrink() *string {
	return s.ClientIdsShrink
}

func (s *UpgradeBackupClientsShrinkRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *UpgradeBackupClientsShrinkRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *UpgradeBackupClientsShrinkRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *UpgradeBackupClientsShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *UpgradeBackupClientsShrinkRequest) SetClientIdsShrink(v string) *UpgradeBackupClientsShrinkRequest {
	s.ClientIdsShrink = &v
	return s
}

func (s *UpgradeBackupClientsShrinkRequest) SetCrossAccountRoleName(v string) *UpgradeBackupClientsShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *UpgradeBackupClientsShrinkRequest) SetCrossAccountType(v string) *UpgradeBackupClientsShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *UpgradeBackupClientsShrinkRequest) SetCrossAccountUserId(v int64) *UpgradeBackupClientsShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *UpgradeBackupClientsShrinkRequest) SetInstanceIdsShrink(v string) *UpgradeBackupClientsShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *UpgradeBackupClientsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
