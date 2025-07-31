// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupDBsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v *DescribeBackupDBsResponseBodyDatabases) *DescribeBackupDBsResponseBody
	GetDatabases() *DescribeBackupDBsResponseBodyDatabases
	SetPageNumber(v int32) *DescribeBackupDBsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupDBsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBackupDBsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeBackupDBsResponseBody
	GetTotalCount() *int32
}

type DescribeBackupDBsResponseBody struct {
	// The details of the databases.
	Databases *DescribeBackupDBsResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Struct"`
	// The page number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AF0AD89-ED4F-44AD-B65F-BFC1D5CD9455
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of returned databases.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupDBsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDBsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupDBsResponseBody) GetDatabases() *DescribeBackupDBsResponseBodyDatabases {
	return s.Databases
}

func (s *DescribeBackupDBsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupDBsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupDBsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupDBsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBackupDBsResponseBody) SetDatabases(v *DescribeBackupDBsResponseBodyDatabases) *DescribeBackupDBsResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeBackupDBsResponseBody) SetPageNumber(v int32) *DescribeBackupDBsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupDBsResponseBody) SetPageSize(v int32) *DescribeBackupDBsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupDBsResponseBody) SetRequestId(v string) *DescribeBackupDBsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupDBsResponseBody) SetTotalCount(v int32) *DescribeBackupDBsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupDBsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupDBsResponseBodyDatabases struct {
	Database []*DescribeBackupDBsResponseBodyDatabasesDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s DescribeBackupDBsResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDBsResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *DescribeBackupDBsResponseBodyDatabases) GetDatabase() []*DescribeBackupDBsResponseBodyDatabasesDatabase {
	return s.Database
}

func (s *DescribeBackupDBsResponseBodyDatabases) SetDatabase(v []*DescribeBackupDBsResponseBodyDatabasesDatabase) *DescribeBackupDBsResponseBodyDatabases {
	s.Database = v
	return s
}

func (s *DescribeBackupDBsResponseBodyDatabases) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupDBsResponseBodyDatabasesDatabase struct {
	// The name of the database.
	//
	// example:
	//
	// mongodbtest
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DescribeBackupDBsResponseBodyDatabasesDatabase) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDBsResponseBodyDatabasesDatabase) GoString() string {
	return s.String()
}

func (s *DescribeBackupDBsResponseBodyDatabasesDatabase) GetDBName() *string {
	return s.DBName
}

func (s *DescribeBackupDBsResponseBodyDatabasesDatabase) SetDBName(v string) *DescribeBackupDBsResponseBodyDatabasesDatabase {
	s.DBName = &v
	return s
}

func (s *DescribeBackupDBsResponseBodyDatabasesDatabase) Validate() error {
	return dara.Validate(s)
}
