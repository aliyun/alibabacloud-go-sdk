// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*ListOperationDatabasesResponseBodyDatabases) *ListOperationDatabasesResponseBody
	GetDatabases() []*ListOperationDatabasesResponseBodyDatabases
	SetRequestId(v string) *ListOperationDatabasesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListOperationDatabasesResponseBody
	GetTotalCount() *int64
}

type ListOperationDatabasesResponseBody struct {
	// The databases returned.
	Databases []*ListOperationDatabasesResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOperationDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationDatabasesResponseBody) GetDatabases() []*ListOperationDatabasesResponseBodyDatabases {
	return s.Databases
}

func (s *ListOperationDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationDatabasesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOperationDatabasesResponseBody) SetDatabases(v []*ListOperationDatabasesResponseBodyDatabases) *ListOperationDatabasesResponseBody {
	s.Databases = v
	return s
}

func (s *ListOperationDatabasesResponseBody) SetRequestId(v string) *ListOperationDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationDatabasesResponseBody) SetTotalCount(v int64) *ListOperationDatabasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOperationDatabasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOperationDatabasesResponseBodyDatabases struct {
	// The address type of the database. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Private
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The remarks of the database.
	//
	// example:
	//
	// cpp
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The database ID.
	//
	// example:
	//
	// 26
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The database name.
	//
	// example:
	//
	// zDatabase
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The port of the database.
	//
	// example:
	//
	// 3306
	DatabasePort *int64 `json:"DatabasePort,omitempty" xml:"DatabasePort,omitempty"`
	// The private address of the database.
	//
	// example:
	//
	// rm-b******9b.mysql.rds.aliyuncs.com
	DatabasePrivateAddress *string `json:"DatabasePrivateAddress,omitempty" xml:"DatabasePrivateAddress,omitempty"`
	// The public address of the database.
	//
	// example:
	//
	// rm-uf******p45.mysql.rds.aliyuncs.com
	DatabasePublicAddress *string `json:"DatabasePublicAddress,omitempty" xml:"DatabasePublicAddress,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PostgreSQL**
	//
	// 	- **Oracle**
	//
	// 	- **SQLServer**
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The database type. Valid values:
	//
	// 	- **Local**: on-premises database.
	//
	// 	- **Rds**: ApsaraDB RDS instance.
	//
	// 	- **PolarDB**: PolarDB cluster.
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the ApsaraDB RDS instance.
	//
	// example:
	//
	// i-wz9225bhipya******
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The region ID of the ApsaraDB RDS instance.
	//
	// example:
	//
	// cn-shanghai
	SourceInstanceRegionId *string `json:"SourceInstanceRegionId,omitempty" xml:"SourceInstanceRegionId,omitempty"`
	// The database status. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **Release**
	//
	// example:
	//
	// Normal
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
}

func (s ListOperationDatabasesResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s ListOperationDatabasesResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *ListOperationDatabasesResponseBodyDatabases) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *ListOperationDatabasesResponseBodyDatabases) GetComment() *string {
	return s.Comment
}

func (s *ListOperationDatabasesResponseBodyDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListOperationDatabasesResponseBodyDatabases) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListOperationDatabasesResponseBodyDatabases) GetDatabasePort() *int64 {
	return s.DatabasePort
}

func (s *ListOperationDatabasesResponseBodyDatabases) GetDatabasePrivateAddress() *string {
	return s.DatabasePrivateAddress
}

func (s *ListOperationDatabasesResponseBodyDatabases) GetDatabasePublicAddress() *string {
	return s.DatabasePublicAddress
}

func (s *ListOperationDatabasesResponseBodyDatabases) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *ListOperationDatabasesResponseBodyDatabases) GetSource() *string {
	return s.Source
}

func (s *ListOperationDatabasesResponseBodyDatabases) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ListOperationDatabasesResponseBodyDatabases) GetSourceInstanceRegionId() *string {
	return s.SourceInstanceRegionId
}

func (s *ListOperationDatabasesResponseBodyDatabases) GetSourceInstanceState() *string {
	return s.SourceInstanceState
}

func (s *ListOperationDatabasesResponseBodyDatabases) SetActiveAddressType(v string) *ListOperationDatabasesResponseBodyDatabases {
	s.ActiveAddressType = &v
	return s
}

func (s *ListOperationDatabasesResponseBodyDatabases) SetComment(v string) *ListOperationDatabasesResponseBodyDatabases {
	s.Comment = &v
	return s
}

func (s *ListOperationDatabasesResponseBodyDatabases) SetDatabaseId(v string) *ListOperationDatabasesResponseBodyDatabases {
	s.DatabaseId = &v
	return s
}

func (s *ListOperationDatabasesResponseBodyDatabases) SetDatabaseName(v string) *ListOperationDatabasesResponseBodyDatabases {
	s.DatabaseName = &v
	return s
}

func (s *ListOperationDatabasesResponseBodyDatabases) SetDatabasePort(v int64) *ListOperationDatabasesResponseBodyDatabases {
	s.DatabasePort = &v
	return s
}

func (s *ListOperationDatabasesResponseBodyDatabases) SetDatabasePrivateAddress(v string) *ListOperationDatabasesResponseBodyDatabases {
	s.DatabasePrivateAddress = &v
	return s
}

func (s *ListOperationDatabasesResponseBodyDatabases) SetDatabasePublicAddress(v string) *ListOperationDatabasesResponseBodyDatabases {
	s.DatabasePublicAddress = &v
	return s
}

func (s *ListOperationDatabasesResponseBodyDatabases) SetDatabaseType(v string) *ListOperationDatabasesResponseBodyDatabases {
	s.DatabaseType = &v
	return s
}

func (s *ListOperationDatabasesResponseBodyDatabases) SetSource(v string) *ListOperationDatabasesResponseBodyDatabases {
	s.Source = &v
	return s
}

func (s *ListOperationDatabasesResponseBodyDatabases) SetSourceInstanceId(v string) *ListOperationDatabasesResponseBodyDatabases {
	s.SourceInstanceId = &v
	return s
}

func (s *ListOperationDatabasesResponseBodyDatabases) SetSourceInstanceRegionId(v string) *ListOperationDatabasesResponseBodyDatabases {
	s.SourceInstanceRegionId = &v
	return s
}

func (s *ListOperationDatabasesResponseBodyDatabases) SetSourceInstanceState(v string) *ListOperationDatabasesResponseBodyDatabases {
	s.SourceInstanceState = &v
	return s
}

func (s *ListOperationDatabasesResponseBodyDatabases) Validate() error {
	return dara.Validate(s)
}
