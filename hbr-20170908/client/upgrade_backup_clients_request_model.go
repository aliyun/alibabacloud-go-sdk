// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeBackupClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIds(v map[string]interface{}) *UpgradeBackupClientsRequest
	GetClientIds() map[string]interface{}
	SetCrossAccountRoleName(v string) *UpgradeBackupClientsRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *UpgradeBackupClientsRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *UpgradeBackupClientsRequest
	GetCrossAccountUserId() *int64
	SetInstanceIds(v map[string]interface{}) *UpgradeBackupClientsRequest
	GetInstanceIds() map[string]interface{}
}

type UpgradeBackupClientsRequest struct {
	// The IDs of Cloud Backup clients. The total number of Cloud Backup client IDs and ECS instance IDs cannot exceed 100.
	//
	// example:
	//
	// ["i-0xi5wj******3j3bh2gj5"]
	ClientIds map[string]interface{} `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
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
	InstanceIds map[string]interface{} `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s UpgradeBackupClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeBackupClientsRequest) GoString() string {
	return s.String()
}

func (s *UpgradeBackupClientsRequest) GetClientIds() map[string]interface{} {
	return s.ClientIds
}

func (s *UpgradeBackupClientsRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *UpgradeBackupClientsRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *UpgradeBackupClientsRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *UpgradeBackupClientsRequest) GetInstanceIds() map[string]interface{} {
	return s.InstanceIds
}

func (s *UpgradeBackupClientsRequest) SetClientIds(v map[string]interface{}) *UpgradeBackupClientsRequest {
	s.ClientIds = v
	return s
}

func (s *UpgradeBackupClientsRequest) SetCrossAccountRoleName(v string) *UpgradeBackupClientsRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *UpgradeBackupClientsRequest) SetCrossAccountType(v string) *UpgradeBackupClientsRequest {
	s.CrossAccountType = &v
	return s
}

func (s *UpgradeBackupClientsRequest) SetCrossAccountUserId(v int64) *UpgradeBackupClientsRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *UpgradeBackupClientsRequest) SetInstanceIds(v map[string]interface{}) *UpgradeBackupClientsRequest {
	s.InstanceIds = v
	return s
}

func (s *UpgradeBackupClientsRequest) Validate() error {
	return dara.Validate(s)
}
