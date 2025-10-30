// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceIndexUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDBInstanceIndexUsageResponseBodyItems) *DescribeDBInstanceIndexUsageResponseBody
	GetItems() []*DescribeDBInstanceIndexUsageResponseBodyItems
	SetPageNumber(v int32) *DescribeDBInstanceIndexUsageResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeDBInstanceIndexUsageResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDBInstanceIndexUsageResponseBody
	GetTotalCount() *int32
}

type DescribeDBInstanceIndexUsageResponseBody struct {
	// The queried index usage.
	Items []*DescribeDBInstanceIndexUsageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceIndexUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIndexUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIndexUsageResponseBody) GetItems() []*DescribeDBInstanceIndexUsageResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstanceIndexUsageResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceIndexUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceIndexUsageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDBInstanceIndexUsageResponseBody) SetItems(v []*DescribeDBInstanceIndexUsageResponseBodyItems) *DescribeDBInstanceIndexUsageResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBody) SetPageNumber(v int32) *DescribeDBInstanceIndexUsageResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBody) SetRequestId(v string) *DescribeDBInstanceIndexUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBody) SetTotalCount(v int32) *DescribeDBInstanceIndexUsageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceIndexUsageResponseBodyItems struct {
	// The name of the database.
	//
	// example:
	//
	// adbtest
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The definition of the index.
	//
	// example:
	//
	// CREATE INDEX idx1 ON schema1.tab1_ptr_2010
	IndexDef *string `json:"IndexDef,omitempty" xml:"IndexDef,omitempty"`
	// The name of the index.
	//
	// example:
	//
	// idx1
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	// The number of index scans.
	//
	// example:
	//
	// 50000
	IndexScanTimes *int64 `json:"IndexScanTimes,omitempty" xml:"IndexScanTimes,omitempty"`
	// The size of the index. Unit: bytes.
	//
	// example:
	//
	// 10000
	IndexSize *string `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	// Indicates whether the table is a partitioned table. Valid values:
	//
	// 	- **true**: The table is a partitioned table.
	//
	// 	- **false**: The table is not a partitioned table.
	//
	// example:
	//
	// true
	IsPartitionTable *bool `json:"IsPartitionTable,omitempty" xml:"IsPartitionTable,omitempty"`
	// The name of the parent table.
	//
	// example:
	//
	// tab1
	ParentTableName *string `json:"ParentTableName,omitempty" xml:"ParentTableName,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// schema1
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// tab1_ptr_2010
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The time when the table was last deleted, inserted, or updated.
	//
	// example:
	//
	// 2022-08-08T20:00:00Z
	TimeLastUpdated *string `json:"TimeLastUpdated,omitempty" xml:"TimeLastUpdated,omitempty"`
}

func (s DescribeDBInstanceIndexUsageResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceIndexUsageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) GetIndexDef() *string {
	return s.IndexDef
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) GetIndexName() *string {
	return s.IndexName
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) GetIndexScanTimes() *int64 {
	return s.IndexScanTimes
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) GetIndexSize() *string {
	return s.IndexSize
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) GetIsPartitionTable() *bool {
	return s.IsPartitionTable
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) GetParentTableName() *string {
	return s.ParentTableName
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) GetTimeLastUpdated() *string {
	return s.TimeLastUpdated
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetDatabaseName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIndexDef(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IndexDef = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIndexName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IndexName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIndexScanTimes(v int64) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IndexScanTimes = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIndexSize(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IndexSize = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetIsPartitionTable(v bool) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.IsPartitionTable = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetParentTableName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.ParentTableName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetSchemaName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetTableName(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) SetTimeLastUpdated(v string) *DescribeDBInstanceIndexUsageResponseBodyItems {
	s.TimeLastUpdated = &v
	return s
}

func (s *DescribeDBInstanceIndexUsageResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
