// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupIds(v string) *DeleteBackupsRequest
	GetBackupIds() *string
	SetDBClusterId(v string) *DeleteBackupsRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DeleteBackupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteBackupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteBackupsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteBackupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteBackupsRequest
	GetResourceOwnerId() *int64
}

type DeleteBackupsRequest struct {
	// The ID of the backup set that you want to delete. Separate multiple backup set IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 766,767
	BackupIds *string `json:"BackupIds,omitempty" xml:"BackupIds,omitempty"`
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupsRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupsRequest) GetBackupIds() *string {
	return s.BackupIds
}

func (s *DeleteBackupsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteBackupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteBackupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteBackupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBackupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteBackupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteBackupsRequest) SetBackupIds(v string) *DeleteBackupsRequest {
	s.BackupIds = &v
	return s
}

func (s *DeleteBackupsRequest) SetDBClusterId(v string) *DeleteBackupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteBackupsRequest) SetOwnerAccount(v string) *DeleteBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteBackupsRequest) SetOwnerId(v int64) *DeleteBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteBackupsRequest) SetRegionId(v string) *DeleteBackupsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBackupsRequest) SetResourceOwnerAccount(v string) *DeleteBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteBackupsRequest) SetResourceOwnerId(v int64) *DeleteBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteBackupsRequest) Validate() error {
	return dara.Validate(s)
}
