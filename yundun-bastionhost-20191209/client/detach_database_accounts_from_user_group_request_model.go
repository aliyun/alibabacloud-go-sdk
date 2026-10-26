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
	SetProjectId(v int64) *DetachDatabaseAccountsFromUserGroupRequest
	GetProjectId() *int64
	SetRegionId(v string) *DetachDatabaseAccountsFromUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *DetachDatabaseAccountsFromUserGroupRequest
	GetUserGroupId() *string
}

type DetachDatabaseAccountsFromUserGroupRequest struct {
	// The information about the databases from which you want to remove authorization.
	Databases []*DetachDatabaseAccountsFromUserGroupRequestDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The ID of the bastion host instance.
	//
	// > You can call [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-7mz2v120f0y
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
	// The ID of the user group from which you want to remove authorized databases and database accounts.
	//
	// > You can call [ListUserGroups](https://help.aliyun.com/document_detail/204509.html) to obtain this parameter.
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

func (s *DetachDatabaseAccountsFromUserGroupRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *DetachDatabaseAccountsFromUserGroupRequest) SetProjectId(v int64) *DetachDatabaseAccountsFromUserGroupRequest {
	s.ProjectId = &v
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

type DetachDatabaseAccountsFromUserGroupRequestDatabases struct {
	// The IDs of the database accounts from which you want to remove authorization.
	DatabaseAccountIds []*string `json:"DatabaseAccountIds,omitempty" xml:"DatabaseAccountIds,omitempty" type:"Repeated"`
	// The instance ID of the database from which you want to remove authorization.
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
