// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAddress(v string) *ListOperationDatabasesRequest
	GetDatabaseAddress() *string
	SetDatabaseName(v string) *ListOperationDatabasesRequest
	GetDatabaseName() *string
	SetDatabaseType(v string) *ListOperationDatabasesRequest
	GetDatabaseType() *string
	SetInstanceId(v string) *ListOperationDatabasesRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListOperationDatabasesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListOperationDatabasesRequest
	GetPageSize() *string
	SetRegionId(v string) *ListOperationDatabasesRequest
	GetRegionId() *string
	SetSource(v string) *ListOperationDatabasesRequest
	GetSource() *string
	SetSourceInstanceId(v string) *ListOperationDatabasesRequest
	GetSourceInstanceId() *string
	SetSourceInstanceState(v string) *ListOperationDatabasesRequest
	GetSourceInstanceState() *string
}

type ListOperationDatabasesRequest struct {
	// The address of the database.
	//
	// example:
	//
	// 10.167.XX.XX
	DatabaseAddress *string `json:"DatabaseAddress,omitempty" xml:"DatabaseAddress,omitempty"`
	// The name of the database. This parameter supports only exact matches.
	//
	// example:
	//
	// aaa
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database type. Valid values:
	//
	// - **MySQL**
	//
	// - **SQLServer**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The ID of the Bastionhost instance.
	//
	// > Call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-tl32wdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. The default value is **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.<br>The maximum value is 100. The default value is 20. If you do not specify this parameter, 20 entries are returned.<br>
	//
	// > Specify a value for this parameter.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the database. Valid values:
	//
	// - **Local**: a local database
	//
	// - **Rds**: an ApsaraDB RDS database
	//
	// - **PolarDB**: a PolarDB database
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the source instance. This parameter supports only exact matches.
	//
	// example:
	//
	// i-bp19ienyt0yax748****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The status of the source instance. You can use this parameter to filter the results.
	//
	// - **Normal**: The instance is running.
	//
	// - **RemoteRelease**: The instance is released.
	//
	// example:
	//
	// Normal
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
}

func (s ListOperationDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListOperationDatabasesRequest) GetDatabaseAddress() *string {
	return s.DatabaseAddress
}

func (s *ListOperationDatabasesRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListOperationDatabasesRequest) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *ListOperationDatabasesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationDatabasesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListOperationDatabasesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListOperationDatabasesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOperationDatabasesRequest) GetSource() *string {
	return s.Source
}

func (s *ListOperationDatabasesRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ListOperationDatabasesRequest) GetSourceInstanceState() *string {
	return s.SourceInstanceState
}

func (s *ListOperationDatabasesRequest) SetDatabaseAddress(v string) *ListOperationDatabasesRequest {
	s.DatabaseAddress = &v
	return s
}

func (s *ListOperationDatabasesRequest) SetDatabaseName(v string) *ListOperationDatabasesRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListOperationDatabasesRequest) SetDatabaseType(v string) *ListOperationDatabasesRequest {
	s.DatabaseType = &v
	return s
}

func (s *ListOperationDatabasesRequest) SetInstanceId(v string) *ListOperationDatabasesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOperationDatabasesRequest) SetPageNumber(v string) *ListOperationDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOperationDatabasesRequest) SetPageSize(v string) *ListOperationDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationDatabasesRequest) SetRegionId(v string) *ListOperationDatabasesRequest {
	s.RegionId = &v
	return s
}

func (s *ListOperationDatabasesRequest) SetSource(v string) *ListOperationDatabasesRequest {
	s.Source = &v
	return s
}

func (s *ListOperationDatabasesRequest) SetSourceInstanceId(v string) *ListOperationDatabasesRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *ListOperationDatabasesRequest) SetSourceInstanceState(v string) *ListOperationDatabasesRequest {
	s.SourceInstanceState = &v
	return s
}

func (s *ListOperationDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
