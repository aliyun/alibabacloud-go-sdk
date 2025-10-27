// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTableStatisticsResponseBody
	GetDBClusterId() *string
	SetItems(v *DescribeTableStatisticsResponseBodyItems) *DescribeTableStatisticsResponseBody
	GetItems() *DescribeTableStatisticsResponseBodyItems
	SetPageNumber(v string) *DescribeTableStatisticsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeTableStatisticsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeTableStatisticsResponseBody
	GetRequestId() *string
	SetSchemaNames(v string) *DescribeTableStatisticsResponseBody
	GetSchemaNames() *string
	SetTotalCount(v string) *DescribeTableStatisticsResponseBody
	GetTotalCount() *string
}

type DescribeTableStatisticsResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// am-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried table statistics.
	Items *DescribeTableStatisticsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4C4433FF-5D3A-4C3E-A19C-6D93B2******
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SchemaNames *string `json:"SchemaNames,omitempty" xml:"SchemaNames,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTableStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTableStatisticsResponseBody) GetItems() *DescribeTableStatisticsResponseBodyItems {
	return s.Items
}

func (s *DescribeTableStatisticsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeTableStatisticsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeTableStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTableStatisticsResponseBody) GetSchemaNames() *string {
	return s.SchemaNames
}

func (s *DescribeTableStatisticsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeTableStatisticsResponseBody) SetDBClusterId(v string) *DescribeTableStatisticsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetItems(v *DescribeTableStatisticsResponseBodyItems) *DescribeTableStatisticsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetPageNumber(v string) *DescribeTableStatisticsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetPageSize(v string) *DescribeTableStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetRequestId(v string) *DescribeTableStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetSchemaNames(v string) *DescribeTableStatisticsResponseBody {
	s.SchemaNames = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) SetTotalCount(v string) *DescribeTableStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTableStatisticsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTableStatisticsResponseBodyItems struct {
	TableStatisticRecords []*DescribeTableStatisticsResponseBodyItemsTableStatisticRecords `json:"TableStatisticRecords,omitempty" xml:"TableStatisticRecords,omitempty" type:"Repeated"`
}

func (s DescribeTableStatisticsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableStatisticsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponseBodyItems) GetTableStatisticRecords() []*DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	return s.TableStatisticRecords
}

func (s *DescribeTableStatisticsResponseBodyItems) SetTableStatisticRecords(v []*DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) *DescribeTableStatisticsResponseBodyItems {
	s.TableStatisticRecords = v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItems) Validate() error {
	if s.TableStatisticRecords != nil {
		for _, item := range s.TableStatisticRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTableStatisticsResponseBodyItemsTableStatisticRecords struct {
	// The cold data size. Unit: bytes.
	//
	// >  The parameter is returned only for AnalyticDB for MySQL clusters of V3.1.3.4 or later.
	//
	// example:
	//
	// 0
	ColdDataSize *int64 `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty"`
	// The data size of table records. Unit: bytes.
	//
	// >  The data size of table records, excluding the data size of regular index and primary key indexes.
	//
	// example:
	//
	// 15592
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The hot data size. Unit: bytes.
	//
	// example:
	//
	// 1048576
	HotDataSize *int64 `json:"HotDataSize,omitempty" xml:"HotDataSize,omitempty"`
	// The data size of regular indexes. Unit: bytes.
	//
	// example:
	//
	// 3076
	IndexSize *int64 `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	// The data size of other data. Unit: bytes.
	//
	// example:
	//
	// 1048576
	OtherSize *int64 `json:"OtherSize,omitempty" xml:"OtherSize,omitempty"`
	// The number of partitions.
	//
	// example:
	//
	// 1
	PartitionCount *int64 `json:"PartitionCount,omitempty" xml:"PartitionCount,omitempty"`
	// The data size of primary key indexes. Unit: bytes.
	//
	// example:
	//
	// 16340
	PrimaryKeyIndexSize *int64 `json:"PrimaryKeyIndexSize,omitempty" xml:"PrimaryKeyIndexSize,omitempty"`
	// The number of rows in the table.
	//
	// example:
	//
	// 3
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test_schema
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The percentage of the table size. Unit: %.
	//
	// >  Formula: Table storage percentage = Total data size of a table/Total data size of the cluster × 100%.
	//
	// example:
	//
	// 66.23
	SpaceRatio *float64 `json:"SpaceRatio,omitempty" xml:"SpaceRatio,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The total data size. Unit: bytes.
	//
	// >  The following formulas can be used to calculate the total data size: Formula 1: Total data size = Hot data size + Cold data size. Formula 2: Total data size = Data size of table records + Data size of regular indexes + Data size of primary key indexes + Data size of other data.
	//
	// example:
	//
	// 1577
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GoString() string {
	return s.String()
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GetColdDataSize() *int64 {
	return s.ColdDataSize
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GetDataSize() *int64 {
	return s.DataSize
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GetHotDataSize() *int64 {
	return s.HotDataSize
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GetIndexSize() *int64 {
	return s.IndexSize
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GetOtherSize() *int64 {
	return s.OtherSize
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GetPartitionCount() *int64 {
	return s.PartitionCount
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GetPrimaryKeyIndexSize() *int64 {
	return s.PrimaryKeyIndexSize
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GetRowCount() *int64 {
	return s.RowCount
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GetSpaceRatio() *float64 {
	return s.SpaceRatio
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GetTableName() *string {
	return s.TableName
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetColdDataSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.ColdDataSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetDataSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.DataSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetHotDataSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.HotDataSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetIndexSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.IndexSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetOtherSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.OtherSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetPartitionCount(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.PartitionCount = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetPrimaryKeyIndexSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.PrimaryKeyIndexSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetRowCount(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.RowCount = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetSchemaName(v string) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.SchemaName = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetSpaceRatio(v float64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.SpaceRatio = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetTableName(v string) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.TableName = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) SetTotalSize(v int64) *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords {
	s.TotalSize = &v
	return s
}

func (s *DescribeTableStatisticsResponseBodyItemsTableStatisticRecords) Validate() error {
	return dara.Validate(s)
}
