// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteDBInstanceRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DeleteDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDBInstanceRequest
	GetOwnerId() *int64
	SetReleasedKeepPolicy(v string) *DeleteDBInstanceRequest
	GetReleasedKeepPolicy() *string
	SetResourceOwnerAccount(v string) *DeleteDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBInstanceRequest
	GetResourceOwnerId() *int64
}

type DeleteDBInstanceRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The policy that is used to retain archived backup files if the instance is released. Default value: None. Valid values:
	//
	// 	- **None**: No archived backup files are retained.
	//
	// 	- **Lastest**: Only the last archived backup file is retained.
	//
	// 	- **All**: All archived backup files are retained.
	//
	// > This parameter is supported only for ApsaraDB RDS for MySQL instance with local disks.
	//
	// example:
	//
	// Lastest
	ReleasedKeepPolicy   *string `json:"ReleasedKeepPolicy,omitempty" xml:"ReleasedKeepPolicy,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBInstanceRequest) GetReleasedKeepPolicy() *string {
	return s.ReleasedKeepPolicy
}

func (s *DeleteDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBInstanceRequest) SetDBInstanceId(v string) *DeleteDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetOwnerAccount(v string) *DeleteDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetOwnerId(v int64) *DeleteDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetReleasedKeepPolicy(v string) *DeleteDBInstanceRequest {
	s.ReleasedKeepPolicy = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetResourceOwnerAccount(v string) *DeleteDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetResourceOwnerId(v int64) *DeleteDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
