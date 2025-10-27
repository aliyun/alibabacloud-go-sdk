// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantOperatorPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GrantOperatorPermissionRequest
	GetDBClusterId() *string
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
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1565u55p32****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The expiration time of the service account permissions. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-30T00:00:00Z
	ExpiredTime  *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the permissions. Valid values:
	//
	// 	- **Control**: configuration permissions. The service account is granted permissions to query and modify cluster configurations.
	//
	// 	- **Data**: data permissions. The service account is granted permissions to query schemas, indexes, and SQL statements.
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

func (s *GrantOperatorPermissionRequest) GetDBClusterId() *string {
	return s.DBClusterId
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

func (s *GrantOperatorPermissionRequest) SetDBClusterId(v string) *GrantOperatorPermissionRequest {
	s.DBClusterId = &v
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
