// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateSecurityIPModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *MigrateSecurityIPModeRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *MigrateSecurityIPModeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *MigrateSecurityIPModeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MigrateSecurityIPModeRequest
	GetResourceOwnerId() *int64
}

type MigrateSecurityIPModeRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s MigrateSecurityIPModeRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateSecurityIPModeRequest) GoString() string {
	return s.String()
}

func (s *MigrateSecurityIPModeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MigrateSecurityIPModeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MigrateSecurityIPModeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MigrateSecurityIPModeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MigrateSecurityIPModeRequest) SetDBInstanceId(v string) *MigrateSecurityIPModeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateSecurityIPModeRequest) SetOwnerId(v int64) *MigrateSecurityIPModeRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateSecurityIPModeRequest) SetResourceOwnerAccount(v string) *MigrateSecurityIPModeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateSecurityIPModeRequest) SetResourceOwnerId(v int64) *MigrateSecurityIPModeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateSecurityIPModeRequest) Validate() error {
	return dara.Validate(s)
}
