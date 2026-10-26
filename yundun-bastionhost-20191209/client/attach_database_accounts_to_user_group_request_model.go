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
	SetProjectId(v int64) *AttachDatabaseAccountsToUserGroupRequest
	GetProjectId() *int64
	SetRegionId(v string) *AttachDatabaseAccountsToUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *AttachDatabaseAccountsToUserGroupRequest
	GetUserGroupId() *string
}

type AttachDatabaseAccountsToUserGroupRequest struct {
	// The array of database objects.
	//
	// > A maximum of 10 databases and 10 database accounts are supported. You can leave the database accounts unspecified. If no database accounts are specified, permissions are granted at the database level.
	Databases []*AttachDatabaseAccountsToUserGroupRequestDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-zvp282aly06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group to which you want to grant permissions.
	//
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

func (s *AttachDatabaseAccountsToUserGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *AttachDatabaseAccountsToUserGroupRequest) SetProjectId(v int64) *AttachDatabaseAccountsToUserGroupRequest {
	s.ProjectId = &v
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

type AttachDatabaseAccountsToUserGroupRequestDatabases struct {
	// The array of database account IDs.
	DatabaseAccountIds []*string `json:"DatabaseAccountIds,omitempty" xml:"DatabaseAccountIds,omitempty" type:"Repeated"`
	// The ID of the database instance to which you want to grant permissions.
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
