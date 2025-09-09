// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAddress(v string) *ListDatabasesForUserRequest
	GetDatabaseAddress() *string
	SetDatabaseName(v string) *ListDatabasesForUserRequest
	GetDatabaseName() *string
	SetDatabaseType(v string) *ListDatabasesForUserRequest
	GetDatabaseType() *string
	SetInstanceId(v string) *ListDatabasesForUserRequest
	GetInstanceId() *string
	SetNetworkDomainId(v string) *ListDatabasesForUserRequest
	GetNetworkDomainId() *string
	SetPageNumber(v string) *ListDatabasesForUserRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDatabasesForUserRequest
	GetPageSize() *string
	SetRegionId(v string) *ListDatabasesForUserRequest
	GetRegionId() *string
	SetUserId(v string) *ListDatabasesForUserRequest
	GetUserId() *string
}

type ListDatabasesForUserRequest struct {
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
	// MySQL-8.0
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
	// bastionhost-cn-tl32swayw7o
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network domain where the database to query resides.
	//
	// example:
	//
	// 5
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The page number. Default value: **1**. Pages start from page 1.
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
	// The ID of the user to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDatabasesForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesForUserRequest) GoString() string {
	return s.String()
}

func (s *ListDatabasesForUserRequest) GetDatabaseAddress() *string {
	return s.DatabaseAddress
}

func (s *ListDatabasesForUserRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListDatabasesForUserRequest) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *ListDatabasesForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDatabasesForUserRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *ListDatabasesForUserRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDatabasesForUserRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDatabasesForUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDatabasesForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListDatabasesForUserRequest) SetDatabaseAddress(v string) *ListDatabasesForUserRequest {
	s.DatabaseAddress = &v
	return s
}

func (s *ListDatabasesForUserRequest) SetDatabaseName(v string) *ListDatabasesForUserRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListDatabasesForUserRequest) SetDatabaseType(v string) *ListDatabasesForUserRequest {
	s.DatabaseType = &v
	return s
}

func (s *ListDatabasesForUserRequest) SetInstanceId(v string) *ListDatabasesForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDatabasesForUserRequest) SetNetworkDomainId(v string) *ListDatabasesForUserRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *ListDatabasesForUserRequest) SetPageNumber(v string) *ListDatabasesForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabasesForUserRequest) SetPageSize(v string) *ListDatabasesForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatabasesForUserRequest) SetRegionId(v string) *ListDatabasesForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListDatabasesForUserRequest) SetUserId(v string) *ListDatabasesForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListDatabasesForUserRequest) Validate() error {
	return dara.Validate(s)
}
