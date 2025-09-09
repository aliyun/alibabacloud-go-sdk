// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDatabaseAccountsToUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*AttachDatabaseAccountsToUserGroupRequestDatabases) *AttachDatabaseAccountsToUserGroupRequest
	GetDatabases() []*AttachDatabaseAccountsToUserGroupRequestDatabases
	SetInstanceId(v string) *AttachDatabaseAccountsToUserGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *AttachDatabaseAccountsToUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *AttachDatabaseAccountsToUserGroupRequest
	GetUserGroupId() *string
}

type AttachDatabaseAccountsToUserGroupRequest struct {
	// An array that consists of the database objects.
	//
	// >  You can specify up to 10 databases and 10 database accounts. The database accounts are not required. If you do not specify a database account, the user group is authorized to manage only the databases.
	Databases []*AttachDatabaseAccountsToUserGroupRequestDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-zvp282aly06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AttachDatabaseAccountsToUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachDatabaseAccountsToUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AttachDatabaseAccountsToUserGroupRequest) GetDatabases() []*AttachDatabaseAccountsToUserGroupRequestDatabases {
	return s.Databases
}

func (s *AttachDatabaseAccountsToUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachDatabaseAccountsToUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachDatabaseAccountsToUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AttachDatabaseAccountsToUserGroupRequest) SetDatabases(v []*AttachDatabaseAccountsToUserGroupRequestDatabases) *AttachDatabaseAccountsToUserGroupRequest {
	s.Databases = v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupRequest) SetInstanceId(v string) *AttachDatabaseAccountsToUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupRequest) SetRegionId(v string) *AttachDatabaseAccountsToUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupRequest) SetUserGroupId(v string) *AttachDatabaseAccountsToUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupRequest) Validate() error {
	return dara.Validate(s)
}

type AttachDatabaseAccountsToUserGroupRequestDatabases struct {
	// An array that consists of database account IDs.
	DatabaseAccountIds []*string `json:"DatabaseAccountIds,omitempty" xml:"DatabaseAccountIds,omitempty" type:"Repeated"`
	// The ID of the database that you want to authorize the user group to manage.
	//
	// example:
	//
	// 58
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
}

func (s AttachDatabaseAccountsToUserGroupRequestDatabases) String() string {
	return dara.Prettify(s)
}

func (s AttachDatabaseAccountsToUserGroupRequestDatabases) GoString() string {
	return s.String()
}

func (s *AttachDatabaseAccountsToUserGroupRequestDatabases) GetDatabaseAccountIds() []*string {
	return s.DatabaseAccountIds
}

func (s *AttachDatabaseAccountsToUserGroupRequestDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *AttachDatabaseAccountsToUserGroupRequestDatabases) SetDatabaseAccountIds(v []*string) *AttachDatabaseAccountsToUserGroupRequestDatabases {
	s.DatabaseAccountIds = v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupRequestDatabases) SetDatabaseId(v string) *AttachDatabaseAccountsToUserGroupRequestDatabases {
	s.DatabaseId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupRequestDatabases) Validate() error {
	return dara.Validate(s)
}
