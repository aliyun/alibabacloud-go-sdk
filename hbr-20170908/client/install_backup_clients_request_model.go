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
