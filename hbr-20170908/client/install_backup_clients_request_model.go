// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallBackupClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccountRoleName(v string) *InstallBackupClientsRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *InstallBackupClientsRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *InstallBackupClientsRequest
	GetCrossAccountUserId() *int64
	SetInstanceIds(v map[string]interface{}) *InstallBackupClientsRequest
	GetInstanceIds() map[string]interface{}
}

type InstallBackupClientsRequest struct {
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
	InstanceIds map[string]interface{} `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s InstallBackupClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallBackupClientsRequest) GoString() string {
	return s.String()
}

func (s *InstallBackupClientsRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *InstallBackupClientsRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *InstallBackupClientsRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *InstallBackupClientsRequest) GetInstanceIds() map[string]interface{} {
	return s.InstanceIds
}

func (s *InstallBackupClientsRequest) SetCrossAccountRoleName(v string) *InstallBackupClientsRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *InstallBackupClientsRequest) SetCrossAccountType(v string) *InstallBackupClientsRequest {
	s.CrossAccountType = &v
	return s
}

func (s *InstallBackupClientsRequest) SetCrossAccountUserId(v int64) *InstallBackupClientsRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *InstallBackupClientsRequest) SetInstanceIds(v map[string]interface{}) *InstallBackupClientsRequest {
	s.InstanceIds = v
	return s
}

func (s *InstallBackupClientsRequest) Validate() error {
	return dara.Validate(s)
}
