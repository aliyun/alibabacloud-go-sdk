// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDatabaseAccountsFromUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*DetachDatabaseAccountsFromUserRequestDatabases) *DetachDatabaseAccountsFromUserRequest
	GetDatabases() []*DetachDatabaseAccountsFromUserRequestDatabases
	SetInstanceId(v string) *DetachDatabaseAccountsFromUserRequest
	GetInstanceId() *string
	SetRegionId(v string) *DetachDatabaseAccountsFromUserRequest
	GetRegionId() *string
	SetUserId(v string) *DetachDatabaseAccountsFromUserRequest
	GetUserId() *string
}

type DetachDatabaseAccountsFromUserRequest struct {
	// The databases.
	Databases []*DetachDatabaseAccountsFromUserRequestDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The bastion host ID.
	//
	// > You can call the DescribeInstances operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-pe335ipfk01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user from whom you want to revoke the permissions on databases and database accounts.
	//
	// > You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the ID of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DetachDatabaseAccountsFromUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachDatabaseAccountsFromUserRequest) GoString() string {
	return s.String()
}

func (s *DetachDatabaseAccountsFromUserRequest) GetDatabases() []*DetachDatabaseAccountsFromUserRequestDatabases {
	return s.Databases
}

func (s *DetachDatabaseAccountsFromUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachDatabaseAccountsFromUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachDatabaseAccountsFromUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *DetachDatabaseAccountsFromUserRequest) SetDatabases(v []*DetachDatabaseAccountsFromUserRequestDatabases) *DetachDatabaseAccountsFromUserRequest {
	s.Databases = v
	return s
}

func (s *DetachDatabaseAccountsFromUserRequest) SetInstanceId(v string) *DetachDatabaseAccountsFromUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserRequest) SetRegionId(v string) *DetachDatabaseAccountsFromUserRequest {
	s.RegionId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserRequest) SetUserId(v string) *DetachDatabaseAccountsFromUserRequest {
	s.UserId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserRequest) Validate() error {
	return dara.Validate(s)
}

type DetachDatabaseAccountsFromUserRequestDatabases struct {
	// An array that consists of database account IDs.
	DatabaseAccountIds []*string `json:"DatabaseAccountIds,omitempty" xml:"DatabaseAccountIds,omitempty" type:"Repeated"`
	// The ID of the database on which you want to revoke permissions.
	//
	// example:
	//
	// 8
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
}

func (s DetachDatabaseAccountsFromUserRequestDatabases) String() string {
	return dara.Prettify(s)
}

func (s DetachDatabaseAccountsFromUserRequestDatabases) GoString() string {
	return s.String()
}

func (s *DetachDatabaseAccountsFromUserRequestDatabases) GetDatabaseAccountIds() []*string {
	return s.DatabaseAccountIds
}

func (s *DetachDatabaseAccountsFromUserRequestDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *DetachDatabaseAccountsFromUserRequestDatabases) SetDatabaseAccountIds(v []*string) *DetachDatabaseAccountsFromUserRequestDatabases {
	s.DatabaseAccountIds = v
	return s
}

func (s *DetachDatabaseAccountsFromUserRequestDatabases) SetDatabaseId(v string) *DetachDatabaseAccountsFromUserRequestDatabases {
	s.DatabaseId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserRequestDatabases) Validate() error {
	return dara.Validate(s)
}
