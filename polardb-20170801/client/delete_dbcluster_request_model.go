// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRetentionPolicyOnClusterDeletion(v string) *DeleteDBClusterRequest
	GetBackupRetentionPolicyOnClusterDeletion() *string
	SetCloudProvider(v string) *DeleteDBClusterRequest
	GetCloudProvider() *string
	SetDBClusterId(v string) *DeleteDBClusterRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DeleteDBClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDBClusterRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDBClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBClusterRequest
	GetResourceOwnerId() *int64
}

type DeleteDBClusterRequest struct {
	// The retention policy applied to the backup sets when the cluster is released. Valid values:
	//
	// 	- **ALL**: permanently retains all backup sets.
	//
	// 	- **LATEST**: permanently retains the most recent backup set that is automatically created before the cluster is released.
	//
	// 	- **NONE**: does not retain backup sets.
	//
	// example:
	//
	// NONE
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	CloudProvider                          *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterRequest) GetBackupRetentionPolicyOnClusterDeletion() *string {
	return s.BackupRetentionPolicyOnClusterDeletion
}

func (s *DeleteDBClusterRequest) GetCloudProvider() *string {
	return s.CloudProvider
}

func (s *DeleteDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBClusterRequest) SetBackupRetentionPolicyOnClusterDeletion(v string) *DeleteDBClusterRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *DeleteDBClusterRequest) SetCloudProvider(v string) *DeleteDBClusterRequest {
	s.CloudProvider = &v
	return s
}

func (s *DeleteDBClusterRequest) SetDBClusterId(v string) *DeleteDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetOwnerAccount(v string) *DeleteDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBClusterRequest) SetOwnerId(v int64) *DeleteDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceOwnerAccount(v string) *DeleteDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceOwnerId(v int64) *DeleteDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
