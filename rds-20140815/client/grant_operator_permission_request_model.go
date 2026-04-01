// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantOperatorPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *GrantOperatorPermissionRequest
	GetDBInstanceId() *string
	SetExpiredTime(v string) *GrantOperatorPermissionRequest
	GetExpiredTime() *string
	SetOwnerAccount(v string) *GrantOperatorPermissionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GrantOperatorPermissionRequest
	GetOwnerId() *int64
	SetPrivileges(v string) *GrantOperatorPermissionRequest
	GetPrivileges() *string
	SetResourceOwnerAccount(v string) *GrantOperatorPermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GrantOperatorPermissionRequest
	GetResourceOwnerId() *int64
}

type GrantOperatorPermissionRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The expiration time of the permissions. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-27T16:00:00Z
	ExpiredTime  *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The permissions that you want to grant to the service account. Valid values:
	//
	// 	- **Control**: the configuration permissions, which allow you to view and modify configurations of the instance.
	//
	// 	- **Data**: the data permissions, which allow you to view schemas, indexes, and SQL statements of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Control
	Privileges           *string `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GrantOperatorPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantOperatorPermissionRequest) GoString() string {
	return s.String()
}

func (s *GrantOperatorPermissionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GrantOperatorPermissionRequest) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *GrantOperatorPermissionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GrantOperatorPermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GrantOperatorPermissionRequest) GetPrivileges() *string {
	return s.Privileges
}

func (s *GrantOperatorPermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GrantOperatorPermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GrantOperatorPermissionRequest) SetDBInstanceId(v string) *GrantOperatorPermissionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetExpiredTime(v string) *GrantOperatorPermissionRequest {
	s.ExpiredTime = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetOwnerAccount(v string) *GrantOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetOwnerId(v int64) *GrantOperatorPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetPrivileges(v string) *GrantOperatorPermissionRequest {
	s.Privileges = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetResourceOwnerAccount(v string) *GrantOperatorPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantOperatorPermissionRequest) SetResourceOwnerId(v int64) *GrantOperatorPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantOperatorPermissionRequest) Validate() error {
	return dara.Validate(s)
}
