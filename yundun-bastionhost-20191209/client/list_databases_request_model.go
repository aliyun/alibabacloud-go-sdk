// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseType(v string) *ListDatabasesRequest
	GetDatabaseType() *string
	SetHostGroupId(v string) *ListDatabasesRequest
	GetHostGroupId() *string
	SetInstanceId(v string) *ListDatabasesRequest
	GetInstanceId() *string
	SetNetworkDomainId(v string) *ListDatabasesRequest
	GetNetworkDomainId() *string
	SetPageNumber(v string) *ListDatabasesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDatabasesRequest
	GetPageSize() *string
	SetProjectId(v int64) *ListDatabasesRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListDatabasesRequest
	GetRegionId() *string
	SetSource(v string) *ListDatabasesRequest
	GetSource() *string
}

type ListDatabasesRequest struct {
	// The type of the database. Valid values:
	//
	// - **MySQL**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// - **SQLServer**
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The ID of the host group.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query the value of this parameter.
	//
	// example:
	//
	// 20
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The instance ID of the Bastionhost instance.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-7mz28f5tk0o
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network domain to which the database belongs.
	//
	// example:
	//
	// 2
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The page number of the current page for paging in a paged query. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page for paging in a paged query.
	//
	// The maximum value of the PageSize parameter is 100. Default value: 20. If this parameter is left empty, 20 entries are returned by default.
	//
	// > Do not leave the PageSize parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the database. Valid values:
	//
	// - **Local**: local database instance
	//
	// - **Rds**: ApsaraDB RDS instance
	//
	// - **PolarDB**: PolarDB instance
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListDatabasesRequest) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *ListDatabasesRequest) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *ListDatabasesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDatabasesRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *ListDatabasesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDatabasesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDatabasesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDatabasesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDatabasesRequest) GetSource() *string {
	return s.Source
}

func (s *ListDatabasesRequest) SetDatabaseType(v string) *ListDatabasesRequest {
	s.DatabaseType = &v
	return s
}

func (s *ListDatabasesRequest) SetHostGroupId(v string) *ListDatabasesRequest {
	s.HostGroupId = &v
	return s
}

func (s *ListDatabasesRequest) SetInstanceId(v string) *ListDatabasesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDatabasesRequest) SetNetworkDomainId(v string) *ListDatabasesRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *ListDatabasesRequest) SetPageNumber(v string) *ListDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatabasesRequest) SetPageSize(v string) *ListDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatabasesRequest) SetProjectId(v int64) *ListDatabasesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDatabasesRequest) SetRegionId(v string) *ListDatabasesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDatabasesRequest) SetSource(v string) *ListDatabasesRequest {
	s.Source = &v
	return s
}

func (s *ListDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
