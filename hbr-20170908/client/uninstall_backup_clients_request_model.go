// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallBackupClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIds(v map[string]interface{}) *UninstallBackupClientsRequest
	GetClientIds() map[string]interface{}
	SetCrossAccountRoleName(v string) *UninstallBackupClientsRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *UninstallBackupClientsRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *UninstallBackupClientsRequest
	GetCrossAccountUserId() *int64
	SetInstanceIds(v map[string]interface{}) *UninstallBackupClientsRequest
	GetInstanceIds() map[string]interface{}
}

type UninstallBackupClientsRequest struct {
	// The IDs of Cloud Backup clients. The sum of the number of Cloud Backup client IDs and the number of ECS instance IDs cannot exceed 20. Otherwise, an error occurs.
	//
	// example:
	//
	// ["c-*********************"]
	ClientIds map[string]interface{} `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
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
	InstanceIds map[string]interface{} `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s UninstallBackupClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallBackupClientsRequest) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientsRequest) GetClientIds() map[string]interface{} {
	return s.ClientIds
}

func (s *UninstallBackupClientsRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *UninstallBackupClientsRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *UninstallBackupClientsRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *UninstallBackupClientsRequest) GetInstanceIds() map[string]interface{} {
	return s.InstanceIds
}

func (s *UninstallBackupClientsRequest) SetClientIds(v map[string]interface{}) *UninstallBackupClientsRequest {
	s.ClientIds = v
	return s
}

func (s *UninstallBackupClientsRequest) SetCrossAccountRoleName(v string) *UninstallBackupClientsRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *UninstallBackupClientsRequest) SetCrossAccountType(v string) *UninstallBackupClientsRequest {
	s.CrossAccountType = &v
	return s
}

func (s *UninstallBackupClientsRequest) SetCrossAccountUserId(v int64) *UninstallBackupClientsRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *UninstallBackupClientsRequest) SetInstanceIds(v map[string]interface{}) *UninstallBackupClientsRequest {
	s.InstanceIds = v
	return s
}

func (s *UninstallBackupClientsRequest) Validate() error {
	return dara.Validate(s)
}
