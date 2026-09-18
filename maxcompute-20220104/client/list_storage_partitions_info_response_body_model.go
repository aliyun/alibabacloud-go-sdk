// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStoragePartitionsInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListStoragePartitionsInfoResponseBodyData) *ListStoragePartitionsInfoResponseBody
	GetData() *ListStoragePartitionsInfoResponseBodyData
	SetErrorCode(v string) *ListStoragePartitionsInfoResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListStoragePartitionsInfoResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *ListStoragePartitionsInfoResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListStoragePartitionsInfoResponseBody
	GetRequestId() *string
}

type ListStoragePartitionsInfoResponseBody struct {
	// The returned data.
	Data *ListStoragePartitionsInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: Informational response - The request has been received and is being processed.
	//
	// - 2xx: Success - The request was successfully received, understood, and accepted.
	//
	// - 3xx: Redirection - Further action must be taken to complete the request.
	//
	// - 4xx: Client error - The request contains invalid parameters or syntax, or cannot be fulfilled.
	//
	// - 5xx: Server error - The server failed to fulfill a valid request.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0be3e0bd16661643917136451ebf55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListStoragePartitionsInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStoragePartitionsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoResponseBody) GetData() *ListStoragePartitionsInfoResponseBodyData {
	return s.Data
}

func (s *ListStoragePartitionsInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListStoragePartitionsInfoResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListStoragePartitionsInfoResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListStoragePartitionsInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStoragePartitionsInfoResponseBody) SetData(v *ListStoragePartitionsInfoResponseBodyData) *ListStoragePartitionsInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) SetErrorCode(v string) *ListStoragePartitionsInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) SetErrorMsg(v string) *ListStoragePartitionsInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) SetHttpCode(v int32) *ListStoragePartitionsInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) SetRequestId(v string) *ListStoragePartitionsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStoragePartitionsInfoResponseBodyData struct {
	// The date to which the statistics apply.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The page number of the returned data.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The storage information for the partitions.
	StoragePartitionInfoList []*ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList `json:"storagePartitionInfoList,omitempty" xml:"storagePartitionInfoList,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 57
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStoragePartitionsInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListStoragePartitionsInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoResponseBodyData) GetDate() *string {
	return s.Date
}

func (s *ListStoragePartitionsInfoResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStoragePartitionsInfoResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStoragePartitionsInfoResponseBodyData) GetStoragePartitionInfoList() []*ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	return s.StoragePartitionInfoList
}

func (s *ListStoragePartitionsInfoResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetDate(v string) *ListStoragePartitionsInfoResponseBodyData {
	s.Date = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetPageNumber(v int64) *ListStoragePartitionsInfoResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetPageSize(v int64) *ListStoragePartitionsInfoResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetStoragePartitionInfoList(v []*ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) *ListStoragePartitionsInfoResponseBodyData {
	s.StoragePartitionInfoList = v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) SetTotalCount(v int64) *ListStoragePartitionsInfoResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyData) Validate() error {
	if s.StoragePartitionInfoList != nil {
		for _, item := range s.StoragePartitionInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList struct {
	// The number of files.
	//
	// example:
	//
	// 2
	FileCount *int64 `json:"fileCount,omitempty" xml:"fileCount,omitempty"`
	// The storage size.
	//
	// example:
	//
	// 1
	FileSize *float64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The unit of the storage size.
	//
	// example:
	//
	// GB
	FileSizeUnit *string `json:"fileSizeUnit,omitempty" xml:"fileSizeUnit,omitempty"`
	// Indicates whether the table is a partitioned table. You can ignore this parameter because this operation returns data only for partitions.
	//
	// example:
	//
	// false
	IsPartitioned *bool `json:"isPartitioned,omitempty" xml:"isPartitioned,omitempty"`
	// The last access time of the partition.
	//
	// > Data collection for this metric began a gradual rollout in July 2023. Consequently, the lastAccessTime may not be recorded for a partition that has not been accessed since then or is accessed only by ALGO jobs or direct reads from Hologres.
	//
	// example:
	//
	// 1694589365
	LastAccessTime *int64 `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	// The partition name.
	//
	// example:
	//
	// ds=20241201
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// The project name.
	//
	// example:
	//
	// odps_project
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// The period-over-period change in the total storage usage over the last {$recentDays} days. This API operation does not return this parameter.
	//
	// example:
	//
	// 1%
	Rate *float64 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The schema name.
	//
	// example:
	//
	// schema
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// The storage type. Valid values:
	//
	// - `standard`: Standard storage
	//
	// - `lowfrequency`: Infrequent-access storage
	//
	// - `longterm`: Archive storage
	//
	// example:
	//
	// standard
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
	// The table name.
	//
	// example:
	//
	// bank_data
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// The access frequency.
	//
	// > - Access activities include:
	//
	// >
	//
	// > > 	- The table is used as input in a SQL compute task.
	//
	// > >
	//
	// > > 	- The table is downloaded via Tunnel.
	//
	// > >
	//
	// > > 	- The table data is read by calling the `Read` operation of the StorageAPI. Partition-level data for partitioned tables is not available. Each access activity increases the access frequency by 1.
	//
	// >
	//
	// > - Data collection for this metric began a gradual rollout in July 2023. Consequently, the access frequency may not be recorded for tables that have not been accessed since then or are accessed only by ALGO jobs or direct reads from Hologres.
	//
	// example:
	//
	// 10
	TotalFrequency *int64 `json:"totalFrequency,omitempty" xml:"totalFrequency,omitempty"`
	// The total data accessed.
	//
	// > The cumulative amount of data read from all access operations.
	//
	// example:
	//
	// 1
	TotalInputAmount *float64 `json:"totalInputAmount,omitempty" xml:"totalInputAmount,omitempty"`
	// The unit of the total data accessed.
	//
	// example:
	//
	// GB
	TotalInputAmountUnit *string `json:"totalInputAmountUnit,omitempty" xml:"totalInputAmountUnit,omitempty"`
	// The type of the object. The value is always PARTITION.
	//
	// example:
	//
	// PARTITION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GoString() string {
	return s.String()
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetFileCount() *int64 {
	return s.FileCount
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetFileSize() *float64 {
	return s.FileSize
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetFileSizeUnit() *string {
	return s.FileSizeUnit
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetIsPartitioned() *bool {
	return s.IsPartitioned
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetLastAccessTime() *int64 {
	return s.LastAccessTime
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetPartition() *string {
	return s.Partition
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetRate() *float64 {
	return s.Rate
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetStorageType() *string {
	return s.StorageType
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetTableName() *string {
	return s.TableName
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetTotalFrequency() *int64 {
	return s.TotalFrequency
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetTotalInputAmount() *float64 {
	return s.TotalInputAmount
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetTotalInputAmountUnit() *string {
	return s.TotalInputAmountUnit
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) GetType() *string {
	return s.Type
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetFileCount(v int64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.FileCount = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetFileSize(v float64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.FileSize = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetFileSizeUnit(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.FileSizeUnit = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetIsPartitioned(v bool) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.IsPartitioned = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetLastAccessTime(v int64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.LastAccessTime = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetPartition(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.Partition = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetProjectName(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.ProjectName = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetRate(v float64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.Rate = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetSchemaName(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.SchemaName = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetStorageType(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.StorageType = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetTableName(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.TableName = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetTotalFrequency(v int64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.TotalFrequency = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetTotalInputAmount(v float64) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.TotalInputAmount = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetTotalInputAmountUnit(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.TotalInputAmountUnit = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) SetType(v string) *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList {
	s.Type = &v
	return s
}

func (s *ListStoragePartitionsInfoResponseBodyDataStoragePartitionInfoList) Validate() error {
	return dara.Validate(s)
}
