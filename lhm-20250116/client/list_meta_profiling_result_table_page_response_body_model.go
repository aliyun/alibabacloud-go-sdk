// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaProfilingResultTablePageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListMetaProfilingResultTablePageResponseBodyData) *ListMetaProfilingResultTablePageResponseBody
	GetData() []*ListMetaProfilingResultTablePageResponseBodyData
	SetErrCode(v string) *ListMetaProfilingResultTablePageResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMetaProfilingResultTablePageResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *ListMetaProfilingResultTablePageResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListMetaProfilingResultTablePageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListMetaProfilingResultTablePageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMetaProfilingResultTablePageResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListMetaProfilingResultTablePageResponseBody
	GetTotalCount() *int32
}

type ListMetaProfilingResultTablePageResponseBody struct {
	// The list of table profiling results. For the element structure, see the descriptions of the sub-fields.
	Data []*ListMetaProfilingResultTablePageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code. This parameter is an empty string if the call is successful.
	//
	// example:
	//
	// InvalidParameter
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. This parameter is an empty string if the call is successful.
	//
	// example:
	//
	// The parameter is invalid or missing.
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The current page number, starting from 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of records to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues of this call.
	//
	// example:
	//
	// 3F45B8C1-2A6E-4C9B-9E1F-7D2A1B0C5E66
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success, and a value of false indicates failure. If the call fails, troubleshoot the issue based on errCode and errMessage.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that meet the conditions.
	//
	// example:
	//
	// 128
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMetaProfilingResultTablePageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingResultTablePageResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingResultTablePageResponseBody) GetData() []*ListMetaProfilingResultTablePageResponseBodyData {
	return s.Data
}

func (s *ListMetaProfilingResultTablePageResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMetaProfilingResultTablePageResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMetaProfilingResultTablePageResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListMetaProfilingResultTablePageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaProfilingResultTablePageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetaProfilingResultTablePageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMetaProfilingResultTablePageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMetaProfilingResultTablePageResponseBody) SetData(v []*ListMetaProfilingResultTablePageResponseBodyData) *ListMetaProfilingResultTablePageResponseBody {
	s.Data = v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBody) SetErrCode(v string) *ListMetaProfilingResultTablePageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBody) SetErrMessage(v string) *ListMetaProfilingResultTablePageResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBody) SetPageIndex(v int32) *ListMetaProfilingResultTablePageResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBody) SetPageSize(v int32) *ListMetaProfilingResultTablePageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBody) SetRequestId(v string) *ListMetaProfilingResultTablePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBody) SetSuccess(v bool) *ListMetaProfilingResultTablePageResponseBody {
	s.Success = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBody) SetTotalCount(v int32) *ListMetaProfilingResultTablePageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMetaProfilingResultTablePageResponseBodyData struct {
	// The average file size.
	//
	// example:
	//
	// 2048
	AverageFileSize *string `json:"averageFileSize,omitempty" xml:"averageFileSize,omitempty"`
	// The average partition size.
	//
	// example:
	//
	// 1048576
	AveragePartitionSize *string `json:"averagePartitionSize,omitempty" xml:"averagePartitionSize,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-07-19 22:00:02
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The name of the database, which is used to filter results by database.
	//
	// example:
	//
	// ods
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// The value of the data partition to which the table belongs, which identifies the partition to which this row of data belongs.
	//
	// example:
	//
	// 20240719
	Ds *string `json:"ds,omitempty" xml:"ds,omitempty"`
	// The type of storage engine used by the table.
	//
	// example:
	//
	// PARQUET
	Engine *string `json:"engine,omitempty" xml:"engine,omitempty"`
	// The number of files.
	//
	// example:
	//
	// 10
	FileNum *int32 `json:"fileNum,omitempty" xml:"fileNum,omitempty"`
	// Indicates whether the table is an external table.
	//
	// example:
	//
	// 0
	IsExternal *int32 `json:"isExternal,omitempty" xml:"isExternal,omitempty"`
	// Indicates whether the table is a partitioned table.
	//
	// example:
	//
	// 1
	IsPtTable *int32 `json:"isPtTable,omitempty" xml:"isPtTable,omitempty"`
	// Indicates whether the table is a transactional table.
	//
	// example:
	//
	// 0
	IsTransactional *int32 `json:"isTransactional,omitempty" xml:"isTransactional,omitempty"`
	// The last access time.
	//
	// example:
	//
	// 2024-07-19 22:00:02
	LastAccessed *string `json:"lastAccessed,omitempty" xml:"lastAccessed,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2024-07-19 22:00:02
	LastModified *string `json:"lastModified,omitempty" xml:"lastModified,omitempty"`
	// The creation time of the largest partition.
	//
	// example:
	//
	// 2024-07-19 22:00:02
	MaxPtCrtTime *string `json:"maxPtCrtTime,omitempty" xml:"maxPtCrtTime,omitempty"`
	// The list of largest partitions.
	MaximumPartitionList []*string `json:"maximumPartitionList,omitempty" xml:"maximumPartitionList,omitempty" type:"Repeated"`
	// The migration progress percentage.
	//
	// example:
	//
	// 80
	MigrationProcess *float64 `json:"migrationProcess,omitempty" xml:"migrationProcess,omitempty"`
	// The migration status, which is used to filter by migration progress or indicate the migration status of an object.
	//
	// example:
	//
	// DONE
	MigrationStatus *string `json:"migrationStatus,omitempty" xml:"migrationStatus,omitempty"`
	// The creation time of the smallest partition.
	//
	// example:
	//
	// 2024-01-01 08:00:00
	MinPtCrtTime *string `json:"minPtCrtTime,omitempty" xml:"minPtCrtTime,omitempty"`
	// The list of smallest partitions.
	MinimumPartitionList []*string `json:"minimumPartitionList,omitempty" xml:"minimumPartitionList,omitempty" type:"Repeated"`
	// Indicates whether the object is newly added.
	//
	// example:
	//
	// 1
	NewAdded *int32 `json:"newAdded,omitempty" xml:"newAdded,omitempty"`
	// The list of partition field names.
	PartitionFieldNameList []*string `json:"partitionFieldNameList,omitempty" xml:"partitionFieldNameList,omitempty" type:"Repeated"`
	// The number of partitions.
	//
	// example:
	//
	// 2048
	PartitionNum *int32 `json:"partitionNum,omitempty" xml:"partitionNum,omitempty"`
	// The number of data rows.
	//
	// example:
	//
	// 1000
	RecordCount *int64 `json:"recordCount,omitempty" xml:"recordCount,omitempty"`
	// The data storage size.
	//
	// example:
	//
	// 1.5
	StorageSize *string `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
	// The table comment or remark.
	//
	// example:
	//
	// Order details table
	TableComment *string `json:"tableComment,omitempty" xml:"tableComment,omitempty"`
	// The table ID.
	//
	// example:
	//
	// 5001
	TableId *string `json:"tableId,omitempty" xml:"tableId,omitempty"`
	// The last modification time of the table schema.
	//
	// example:
	//
	// 2024-07-19 22:00:02
	TableModTime *string `json:"tableModTime,omitempty" xml:"tableModTime,omitempty"`
	// The name of the table, which is used to filter results by table name.
	//
	// example:
	//
	// ods_order_detail
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// The identifier of the table type, such as managed table (MANAGED_TABLE) or external table (EXTERNAL_TABLE).
	//
	// example:
	//
	// MANAGED_TABLE
	TableType *string `json:"tableType,omitempty" xml:"tableType,omitempty"`
}

func (s ListMetaProfilingResultTablePageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingResultTablePageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetAverageFileSize() *string {
	return s.AverageFileSize
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetAveragePartitionSize() *string {
	return s.AveragePartitionSize
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetCreated() *string {
	return s.Created
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetDbName() *string {
	return s.DbName
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetDs() *string {
	return s.Ds
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetEngine() *string {
	return s.Engine
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetFileNum() *int32 {
	return s.FileNum
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetIsExternal() *int32 {
	return s.IsExternal
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetIsPtTable() *int32 {
	return s.IsPtTable
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetIsTransactional() *int32 {
	return s.IsTransactional
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetLastAccessed() *string {
	return s.LastAccessed
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetLastModified() *string {
	return s.LastModified
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetMaxPtCrtTime() *string {
	return s.MaxPtCrtTime
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetMaximumPartitionList() []*string {
	return s.MaximumPartitionList
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetMigrationProcess() *float64 {
	return s.MigrationProcess
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetMigrationStatus() *string {
	return s.MigrationStatus
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetMinPtCrtTime() *string {
	return s.MinPtCrtTime
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetMinimumPartitionList() []*string {
	return s.MinimumPartitionList
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetNewAdded() *int32 {
	return s.NewAdded
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetPartitionFieldNameList() []*string {
	return s.PartitionFieldNameList
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetRecordCount() *int64 {
	return s.RecordCount
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetStorageSize() *string {
	return s.StorageSize
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetTableComment() *string {
	return s.TableComment
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetTableId() *string {
	return s.TableId
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetTableModTime() *string {
	return s.TableModTime
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) GetTableType() *string {
	return s.TableType
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetAverageFileSize(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.AverageFileSize = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetAveragePartitionSize(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.AveragePartitionSize = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetCreated(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.Created = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetDbName(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.DbName = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetDs(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.Ds = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetEngine(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.Engine = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetFileNum(v int32) *ListMetaProfilingResultTablePageResponseBodyData {
	s.FileNum = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetIsExternal(v int32) *ListMetaProfilingResultTablePageResponseBodyData {
	s.IsExternal = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetIsPtTable(v int32) *ListMetaProfilingResultTablePageResponseBodyData {
	s.IsPtTable = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetIsTransactional(v int32) *ListMetaProfilingResultTablePageResponseBodyData {
	s.IsTransactional = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetLastAccessed(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.LastAccessed = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetLastModified(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.LastModified = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetMaxPtCrtTime(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.MaxPtCrtTime = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetMaximumPartitionList(v []*string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.MaximumPartitionList = v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetMigrationProcess(v float64) *ListMetaProfilingResultTablePageResponseBodyData {
	s.MigrationProcess = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetMigrationStatus(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.MigrationStatus = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetMinPtCrtTime(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.MinPtCrtTime = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetMinimumPartitionList(v []*string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.MinimumPartitionList = v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetNewAdded(v int32) *ListMetaProfilingResultTablePageResponseBodyData {
	s.NewAdded = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetPartitionFieldNameList(v []*string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.PartitionFieldNameList = v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetPartitionNum(v int32) *ListMetaProfilingResultTablePageResponseBodyData {
	s.PartitionNum = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetRecordCount(v int64) *ListMetaProfilingResultTablePageResponseBodyData {
	s.RecordCount = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetStorageSize(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.StorageSize = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetTableComment(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.TableComment = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetTableId(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.TableId = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetTableModTime(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.TableModTime = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetTableName(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.TableName = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) SetTableType(v string) *ListMetaProfilingResultTablePageResponseBodyData {
	s.TableType = &v
	return s
}

func (s *ListMetaProfilingResultTablePageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
