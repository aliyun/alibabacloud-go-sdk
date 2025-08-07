// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseDBClusterMigrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinueEnableBinlog(v bool) *CloseDBClusterMigrationRequest
	GetContinueEnableBinlog() *bool
	SetDBClusterId(v string) *CloseDBClusterMigrationRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CloseDBClusterMigrationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CloseDBClusterMigrationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloseDBClusterMigrationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloseDBClusterMigrationRequest
	GetResourceOwnerId() *int64
}

type CloseDBClusterMigrationRequest struct {
	// Specifies whether to continue to enable binary logging. Valid values:
	//
	// 	- **true**: continues to enable binary logging.
	//
	// 	- **false**: disables binary logging.
	//
	// Default value: **true**.
	//
	// > If binary logging is disabled, your PolarDB cluster is restarted.
	//
	// example:
	//
	// true
	ContinueEnableBinlog *bool `json:"ContinueEnableBinlog,omitempty" xml:"ContinueEnableBinlog,omitempty"`
	// The ID of the cluster.
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

func (s CloseDBClusterMigrationRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseDBClusterMigrationRequest) GoString() string {
	return s.String()
}

func (s *CloseDBClusterMigrationRequest) GetContinueEnableBinlog() *bool {
	return s.ContinueEnableBinlog
}

func (s *CloseDBClusterMigrationRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CloseDBClusterMigrationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CloseDBClusterMigrationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloseDBClusterMigrationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloseDBClusterMigrationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloseDBClusterMigrationRequest) SetContinueEnableBinlog(v bool) *CloseDBClusterMigrationRequest {
	s.ContinueEnableBinlog = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetDBClusterId(v string) *CloseDBClusterMigrationRequest {
	s.DBClusterId = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetOwnerAccount(v string) *CloseDBClusterMigrationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetOwnerId(v int64) *CloseDBClusterMigrationRequest {
	s.OwnerId = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetResourceOwnerAccount(v string) *CloseDBClusterMigrationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetResourceOwnerId(v int64) *CloseDBClusterMigrationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) Validate() error {
	return dara.Validate(s)
}
