// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOversizeNonPartitionTableInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeOversizeNonPartitionTableInfosResponseBody
	GetDBClusterId() *string
	SetDetectionItems(v []*DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems) *DescribeOversizeNonPartitionTableInfosResponseBody
	GetDetectionItems() []*DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems
	SetPageNumber(v int32) *DescribeOversizeNonPartitionTableInfosResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeOversizeNonPartitionTableInfosResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeOversizeNonPartitionTableInfosResponseBody
	GetRequestId() *string
	SetTables(v []*DescribeOversizeNonPartitionTableInfosResponseBodyTables) *DescribeOversizeNonPartitionTableInfosResponseBody
	GetTables() []*DescribeOversizeNonPartitionTableInfosResponseBodyTables
	SetTotalCount(v string) *DescribeOversizeNonPartitionTableInfosResponseBody
	GetTotalCount() *string
}

type DescribeOversizeNonPartitionTableInfosResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// am-bp16t5ci7r74s****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried detection items and detection results.
	DetectionItems []*DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems `json:"DetectionItems,omitempty" xml:"DetectionItems,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried oversized non-partitioned tables.
	Tables []*DescribeOversizeNonPartitionTableInfosResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOversizeNonPartitionTableInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOversizeNonPartitionTableInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) GetDetectionItems() []*DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems {
	return s.DetectionItems
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) GetTables() []*DescribeOversizeNonPartitionTableInfosResponseBodyTables {
	return s.Tables
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) SetDBClusterId(v string) *DescribeOversizeNonPartitionTableInfosResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) SetDetectionItems(v []*DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems) *DescribeOversizeNonPartitionTableInfosResponseBody {
	s.DetectionItems = v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) SetPageNumber(v int32) *DescribeOversizeNonPartitionTableInfosResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) SetPageSize(v int32) *DescribeOversizeNonPartitionTableInfosResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) SetRequestId(v string) *DescribeOversizeNonPartitionTableInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) SetTables(v []*DescribeOversizeNonPartitionTableInfosResponseBodyTables) *DescribeOversizeNonPartitionTableInfosResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) SetTotalCount(v string) *DescribeOversizeNonPartitionTableInfosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems struct {
	// The information about the detection result.
	//
	// example:
	//
	// Multiple oversized non-partitioned tables are detected.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the detection item.
	//
	// example:
	//
	// Oversized non-partitioned tables
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The severity level of the detection result.
	//
	// example:
	//
	// NORMAL
	//
	// WARNING
	//
	// CRITICAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems) GoString() string {
	return s.String()
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems) GetMessage() *string {
	return s.Message
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems) GetName() *string {
	return s.Name
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems) SetMessage(v string) *DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems {
	s.Message = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems) SetName(v string) *DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems {
	s.Name = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems) SetStatus(v string) *DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems {
	s.Status = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyDetectionItems) Validate() error {
	return dara.Validate(s)
}

type DescribeOversizeNonPartitionTableInfosResponseBodyTables struct {
	// The data size of the table. Unit: bytes.
	//
	// example:
	//
	// 2921132457984
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The data size of regular indexes. Unit: bytes.
	//
	// example:
	//
	// 111
	IndexSize *int64 `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	// The size of hot data. Unit: bytes.
	//
	// example:
	//
	// 1223
	LocalDataSize *int64 `json:"LocalDataSize,omitempty" xml:"LocalDataSize,omitempty"`
	// The data size of the primary key index. Unit: bytes.
	//
	// example:
	//
	// 123
	PrimaryKeySize *int64 `json:"PrimaryKeySize,omitempty" xml:"PrimaryKeySize,omitempty"`
	// The size of cold data. Unit: bytes.
	//
	// example:
	//
	// 122
	RemoteDataSize *int64 `json:"RemoteDataSize,omitempty" xml:"RemoteDataSize,omitempty"`
	// The number of rows in the table.
	//
	// example:
	//
	// 1222
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The percentage of the table size. Unit: %.
	//
	// >  Formula: Table storage percentage = Total data size of a table/Total data size of the cluster × 100%.
	//
	// example:
	//
	// 0.3
	SpaceRatio *float64 `json:"SpaceRatio,omitempty" xml:"SpaceRatio,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeOversizeNonPartitionTableInfosResponseBodyTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeOversizeNonPartitionTableInfosResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) GetDataSize() *int64 {
	return s.DataSize
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) GetIndexSize() *int64 {
	return s.IndexSize
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) GetLocalDataSize() *int64 {
	return s.LocalDataSize
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) GetPrimaryKeySize() *int64 {
	return s.PrimaryKeySize
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) GetRemoteDataSize() *int64 {
	return s.RemoteDataSize
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) GetRowCount() *int64 {
	return s.RowCount
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) GetSpaceRatio() *float64 {
	return s.SpaceRatio
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) GetTableName() *string {
	return s.TableName
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) SetDataSize(v int64) *DescribeOversizeNonPartitionTableInfosResponseBodyTables {
	s.DataSize = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) SetIndexSize(v int64) *DescribeOversizeNonPartitionTableInfosResponseBodyTables {
	s.IndexSize = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) SetLocalDataSize(v int64) *DescribeOversizeNonPartitionTableInfosResponseBodyTables {
	s.LocalDataSize = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) SetPrimaryKeySize(v int64) *DescribeOversizeNonPartitionTableInfosResponseBodyTables {
	s.PrimaryKeySize = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) SetRemoteDataSize(v int64) *DescribeOversizeNonPartitionTableInfosResponseBodyTables {
	s.RemoteDataSize = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) SetRowCount(v int64) *DescribeOversizeNonPartitionTableInfosResponseBodyTables {
	s.RowCount = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) SetSchemaName(v string) *DescribeOversizeNonPartitionTableInfosResponseBodyTables {
	s.SchemaName = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) SetSpaceRatio(v float64) *DescribeOversizeNonPartitionTableInfosResponseBodyTables {
	s.SpaceRatio = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) SetTableName(v string) *DescribeOversizeNonPartitionTableInfosResponseBodyTables {
	s.TableName = &v
	return s
}

func (s *DescribeOversizeNonPartitionTableInfosResponseBodyTables) Validate() error {
	return dara.Validate(s)
}
