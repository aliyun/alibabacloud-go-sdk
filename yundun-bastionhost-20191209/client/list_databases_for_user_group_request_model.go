// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesForUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAddress(v string) *ListDatabasesForUserGroupRequest
	GetDatabaseAddress() *string
	SetDatabaseName(v string) *ListDatabasesForUserGroupRequest
	GetDatabaseName() *string
	SetDatabaseType(v string) *ListDatabasesForUserGroupRequest
	GetDatabaseType() *string
	SetInstanceId(v string) *ListDatabasesForUserGroupRequest
	GetInstanceId() *string
	SetNetworkDomainId(v string) *ListDatabasesForUserGroupRequest
	GetNetworkDomainId() *string
	SetPageNumber(v string) *ListDatabasesForUserGroupRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDatabasesForUserGroupRequest
	GetPageSize() *string
	SetRegionId(v string) *ListDatabasesForUserGroupRequest
	GetRegionId() *string
	SetUserGroupId(v string) *ListDatabasesForUserGroupRequest
	GetUserGroupId() *string
}

type ListDatabasesForUserGroupRequest struct {
	// The address of the database to query. Only exact match is supported.
	//
	// example:
	//
	// ``47.101.**.**``
	DatabaseAddress *string `json:"DatabaseAddress,omitempty" xml:"DatabaseAddress,omitempty"`
	// The name of the database to query.
	//
	// example:
	//
	// test
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The engine of the database to query. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **Oracle**
	//
	// 	- **PostgreSQL**
	//
	// 	- **SQLServer**
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-7mz2ve7h00a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Valid values: 1 to 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListDatabasesForUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListDatabasesForUserGroupRequest) GetDatabaseAddress() *string {
	return s.DatabaseAddress
}

func (s *ListDatabasesForUserGroupRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListDatabasesForUserGroupRequest) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *ListDatabasesForUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDatabasesForUserGroupRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *ListDatabasesForUserGroupRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDatabasesForUserGroupRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDatabasesForUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDatabasesForUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListDatabasesForUserGroupRequest) SetDatabaseAddress(v string) *ListDatabasesForUserGroupRequest {
	s.DatabaseAddress = &v
	return s
}

func (s *ListDatabasesForUserGroupRequest) SetDatabaseName(v string) *ListDatabasesForUserGroupRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListDatabasesForUserGroupRequest) SetDatabaseType(v string) *ListDatabasesForUserGroupRequest {
	s.DatabaseType = &v
	return s
}

func (s *ListDatabasesForUserGroupRequest) SetInstanceId(v string) *ListDatabasesForUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDatabasesForUserGroupRequest) SetNetworkDomainId(v string) *ListDatabasesForUserGroupRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *ListDatabasesForUserGroupRequest) SetPageNumber(v string) *ListDatabasesForUserGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabasesForUserGroupRequest) SetPageSize(v string) *ListDatabasesForUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatabasesForUserGroupRequest) SetRegionId(v string) *ListDatabasesForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListDatabasesForUserGroupRequest) SetUserGroupId(v string) *ListDatabasesForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListDatabasesForUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
