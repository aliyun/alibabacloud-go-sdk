// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDatabaseAccountsToUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*AttachDatabaseAccountsToUserRequestDatabases) *AttachDatabaseAccountsToUserRequest
	GetDatabases() []*AttachDatabaseAccountsToUserRequestDatabases
	SetInstanceId(v string) *AttachDatabaseAccountsToUserRequest
	GetInstanceId() *string
	SetRegionId(v string) *AttachDatabaseAccountsToUserRequest
	GetRegionId() *string
	SetUserId(v string) *AttachDatabaseAccountsToUserRequest
	GetUserId() *string
}

type AttachDatabaseAccountsToUserRequest struct {
	// An array that consists of database objects.
	//
	// >  You can specify up to 10 databases and 10 database accounts. The database accounts are not required. If you do not specify a database account, the user is authorized to manage only the databases.
	Databases []*AttachDatabaseAccountsToUserRequestDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The ID of the bastion host whose user you want to grant permissions.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-wwo36qbv601
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user to be authorized.
	//
	// >  You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to query the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AttachDatabaseAccountsToUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachDatabaseAccountsToUserRequest) GoString() string {
	return s.String()
}

func (s *AttachDatabaseAccountsToUserRequest) GetDatabases() []*AttachDatabaseAccountsToUserRequestDatabases {
	return s.Databases
}

func (s *AttachDatabaseAccountsToUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachDatabaseAccountsToUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachDatabaseAccountsToUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *AttachDatabaseAccountsToUserRequest) SetDatabases(v []*AttachDatabaseAccountsToUserRequestDatabases) *AttachDatabaseAccountsToUserRequest {
	s.Databases = v
	return s
}

func (s *AttachDatabaseAccountsToUserRequest) SetInstanceId(v string) *AttachDatabaseAccountsToUserRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserRequest) SetRegionId(v string) *AttachDatabaseAccountsToUserRequest {
	s.RegionId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserRequest) SetUserId(v string) *AttachDatabaseAccountsToUserRequest {
	s.UserId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserRequest) Validate() error {
	if s.Databases != nil {
		for _, item := range s.Databases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachDatabaseAccountsToUserRequestDatabases struct {
	// An array that consists of database account IDs.
	DatabaseAccountIds []*string `json:"DatabaseAccountIds,omitempty" xml:"DatabaseAccountIds,omitempty" type:"Repeated"`
	// The ID of the database that you want to authorize the user to manage.
	//
	// example:
	//
	// 22
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
}

func (s AttachDatabaseAccountsToUserRequestDatabases) String() string {
	return dara.Prettify(s)
}

func (s AttachDatabaseAccountsToUserRequestDatabases) GoString() string {
	return s.String()
}

func (s *AttachDatabaseAccountsToUserRequestDatabases) GetDatabaseAccountIds() []*string {
	return s.DatabaseAccountIds
}

func (s *AttachDatabaseAccountsToUserRequestDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *AttachDatabaseAccountsToUserRequestDatabases) SetDatabaseAccountIds(v []*string) *AttachDatabaseAccountsToUserRequestDatabases {
	s.DatabaseAccountIds = v
	return s
}

func (s *AttachDatabaseAccountsToUserRequestDatabases) SetDatabaseId(v string) *AttachDatabaseAccountsToUserRequestDatabases {
	s.DatabaseId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserRequestDatabases) Validate() error {
	return dara.Validate(s)
}
