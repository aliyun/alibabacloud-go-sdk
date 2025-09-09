// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDatabaseAccountsFromUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*DetachDatabaseAccountsFromUserGroupRequestDatabases) *DetachDatabaseAccountsFromUserGroupRequest
	GetDatabases() []*DetachDatabaseAccountsFromUserGroupRequestDatabases
	SetInstanceId(v string) *DetachDatabaseAccountsFromUserGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *DetachDatabaseAccountsFromUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *DetachDatabaseAccountsFromUserGroupRequest
	GetUserGroupId() *string
}

type DetachDatabaseAccountsFromUserGroupRequest struct {
	// The information about the database.
	Databases []*DetachDatabaseAccountsFromUserGroupRequestDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-7mz2v120f0y
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group from which you want to revoke permissions on databases and database accounts.
	//
	// > You can call the [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) operation to query the ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DetachDatabaseAccountsFromUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachDatabaseAccountsFromUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DetachDatabaseAccountsFromUserGroupRequest) GetDatabases() []*DetachDatabaseAccountsFromUserGroupRequestDatabases {
	return s.Databases
}

func (s *DetachDatabaseAccountsFromUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachDatabaseAccountsFromUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachDatabaseAccountsFromUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DetachDatabaseAccountsFromUserGroupRequest) SetDatabases(v []*DetachDatabaseAccountsFromUserGroupRequestDatabases) *DetachDatabaseAccountsFromUserGroupRequest {
	s.Databases = v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupRequest) SetInstanceId(v string) *DetachDatabaseAccountsFromUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupRequest) SetRegionId(v string) *DetachDatabaseAccountsFromUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupRequest) SetUserGroupId(v string) *DetachDatabaseAccountsFromUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupRequest) Validate() error {
	return dara.Validate(s)
}

type DetachDatabaseAccountsFromUserGroupRequestDatabases struct {
	// An array that consists of database account IDs.
	DatabaseAccountIds []*string `json:"DatabaseAccountIds,omitempty" xml:"DatabaseAccountIds,omitempty" type:"Repeated"`
	// The ID of the database on which the permissions are to be revoked.
	//
	// example:
	//
	// 4
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
}

func (s DetachDatabaseAccountsFromUserGroupRequestDatabases) String() string {
	return dara.Prettify(s)
}

func (s DetachDatabaseAccountsFromUserGroupRequestDatabases) GoString() string {
	return s.String()
}

func (s *DetachDatabaseAccountsFromUserGroupRequestDatabases) GetDatabaseAccountIds() []*string {
	return s.DatabaseAccountIds
}

func (s *DetachDatabaseAccountsFromUserGroupRequestDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *DetachDatabaseAccountsFromUserGroupRequestDatabases) SetDatabaseAccountIds(v []*string) *DetachDatabaseAccountsFromUserGroupRequestDatabases {
	s.DatabaseAccountIds = v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupRequestDatabases) SetDatabaseId(v string) *DetachDatabaseAccountsFromUserGroupRequestDatabases {
	s.DatabaseId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupRequestDatabases) Validate() error {
	return dara.Validate(s)
}
