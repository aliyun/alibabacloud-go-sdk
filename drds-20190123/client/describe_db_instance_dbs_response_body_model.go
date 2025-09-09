// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbInstanceDbsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v *DescribeDbInstanceDbsResponseBodyDatabases) *DescribeDbInstanceDbsResponseBody
	GetDatabases() *DescribeDbInstanceDbsResponseBodyDatabases
	SetRequestId(v string) *DescribeDbInstanceDbsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDbInstanceDbsResponseBody
	GetSuccess() *bool
	SetTotal(v string) *DescribeDbInstanceDbsResponseBody
	GetTotal() *string
}

type DescribeDbInstanceDbsResponseBody struct {
	// Indicates the information about the storage-layer databases.
	Databases *DescribeDbInstanceDbsResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E9F3D991-08DE-4B74-BE0E-06B809******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates the total number of storage-layer databases.
	//
	// example:
	//
	// 1
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDbInstanceDbsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstanceDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponseBody) GetDatabases() *DescribeDbInstanceDbsResponseBodyDatabases {
	return s.Databases
}

func (s *DescribeDbInstanceDbsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDbInstanceDbsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDbInstanceDbsResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeDbInstanceDbsResponseBody) SetDatabases(v *DescribeDbInstanceDbsResponseBodyDatabases) *DescribeDbInstanceDbsResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDbInstanceDbsResponseBody) SetRequestId(v string) *DescribeDbInstanceDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBody) SetSuccess(v bool) *DescribeDbInstanceDbsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBody) SetTotal(v string) *DescribeDbInstanceDbsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDbInstanceDbsResponseBodyDatabases struct {
	Database []*DescribeDbInstanceDbsResponseBodyDatabasesDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s DescribeDbInstanceDbsResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstanceDbsResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponseBodyDatabases) GetDatabase() []*DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	return s.Database
}

func (s *DescribeDbInstanceDbsResponseBodyDatabases) SetDatabase(v []*DescribeDbInstanceDbsResponseBodyDatabasesDatabase) *DescribeDbInstanceDbsResponseBodyDatabases {
	s.Database = v
	return s
}

func (s *DescribeDbInstanceDbsResponseBodyDatabases) Validate() error {
	return dara.Validate(s)
}

type DescribeDbInstanceDbsResponseBodyDatabasesDatabase struct {
	// Indicates the name of a storage-layer database.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates the description of the storage-layer database.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates the state of the storage-layer database. Valid values:
	//
	// 	- **0**: The database is being created.
	//
	// 	- **1**: The database is available.
	//
	// 	- **3**: The database is being deleted.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDbInstanceDbsResponseBodyDatabasesDatabase) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstanceDbsResponseBodyDatabasesDatabase) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) GetDescription() *string {
	return s.Description
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) SetDbName(v string) *DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	s.DbName = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) SetDescription(v string) *DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	s.Description = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) SetStatus(v int32) *DescribeDbInstanceDbsResponseBodyDatabasesDatabase {
	s.Status = &v
	return s
}

func (s *DescribeDbInstanceDbsResponseBodyDatabasesDatabase) Validate() error {
	return dara.Validate(s)
}
