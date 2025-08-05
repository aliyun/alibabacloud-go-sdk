// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStorageTablesInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListStorageTablesInfoResponseBodyData) *ListStorageTablesInfoResponseBody
	GetData() *ListStorageTablesInfoResponseBodyData
	SetErrorCode(v string) *ListStorageTablesInfoResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListStorageTablesInfoResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *ListStorageTablesInfoResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListStorageTablesInfoResponseBody
	GetRequestId() *string
}

type ListStorageTablesInfoResponseBody struct {
	// The data returned.
	Data *ListStorageTablesInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 	- 1xx: informational response. The request is received and is being processed.
	//
	// 	- 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// 	- 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// 	- 4xx: client error. The request contains invalid request parameters and syntaxes, or specific request conditions cannot be met.
	//
	// 	- 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc12e6a16679892465424670db3eb
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListStorageTablesInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStorageTablesInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoResponseBody) GetData() *ListStorageTablesInfoResponseBodyData {
	return s.Data
}

func (s *ListStorageTablesInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListStorageTablesInfoResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListStorageTablesInfoResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListStorageTablesInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStorageTablesInfoResponseBody) SetData(v *ListStorageTablesInfoResponseBodyData) *ListStorageTablesInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListStorageTablesInfoResponseBody) SetErrorCode(v string) *ListStorageTablesInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListStorageTablesInfoResponseBody) SetErrorMsg(v string) *ListStorageTablesInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListStorageTablesInfoResponseBody) SetHttpCode(v int32) *ListStorageTablesInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListStorageTablesInfoResponseBody) SetRequestId(v string) *ListStorageTablesInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStorageTablesInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListStorageTablesInfoResponseBodyData struct {
	// The date on which the statistics are collected.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The page number.
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
	// The table storage information.
	StorageTableInfoList []*ListStorageTablesInfoResponseBodyDataStorageTableInfoList `json:"storageTableInfoList,omitempty" xml:"storageTableInfoList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStorageTablesInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListStorageTablesInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoResponseBodyData) GetDate() *string {
	return s.Date
}

func (s *ListStorageTablesInfoResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStorageTablesInfoResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStorageTablesInfoResponseBodyData) GetStorageTableInfoList() []*ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	return s.StorageTableInfoList
}

func (s *ListStorageTablesInfoResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListStorageTablesInfoResponseBodyData) SetDate(v string) *ListStorageTablesInfoResponseBodyData {
	s.Date = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) SetPageNumber(v int64) *ListStorageTablesInfoResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) SetPageSize(v int64) *ListStorageTablesInfoResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) SetStorageTableInfoList(v []*ListStorageTablesInfoResponseBodyDataStorageTableInfoList) *ListStorageTablesInfoResponseBodyData {
	s.StorageTableInfoList = v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) SetTotalCount(v int64) *ListStorageTablesInfoResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListStorageTablesInfoResponseBodyDataStorageTableInfoList struct {
	// The date on which the statistics are collected. This value is not returned.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// Indicates whether the table is a partitioned table.
	//
	// example:
	//
	// false
	IsPartitioned *bool `json:"isPartitioned,omitempty" xml:"isPartitioned,omitempty"`
	// The time when the table was last accessed. This value is returned when the table is a non-partitioned table.
	//
	// >  The data collection method is upgraded from July 2023. If the data is not accessed after the upgrade or is accessed by using ALGO jobs or the direct read method of Hologres, the last access time cannot be collected.
	//
	// example:
	//
	// 1694589365
	LastAccessTime *int64 `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	// The storage usage at the long-term storage tier.
	//
	// example:
	//
	// 0
	LongTermStorage *float64 `json:"longTermStorage,omitempty" xml:"longTermStorage,omitempty"`
	// The number of long-term storage files.
	//
	// example:
	//
	// 0
	LongTermStorageFileCount *int64 `json:"longTermStorageFileCount,omitempty" xml:"longTermStorageFileCount,omitempty"`
	// The unit of the storage usage at the long-term storage tier.
	//
	// example:
	//
	// B
	LongTermStorageUnit *string `json:"longTermStorageUnit,omitempty" xml:"longTermStorageUnit,omitempty"`
	// The storage usage at the low-frequency tier.
	//
	// example:
	//
	// 0
	LowFreqStorage *float64 `json:"lowFreqStorage,omitempty" xml:"lowFreqStorage,omitempty"`
	// The number of low-frequency storage files.
	//
	// example:
	//
	// 0
	LowFreqStorageFileCount *int64 `json:"lowFreqStorageFileCount,omitempty" xml:"lowFreqStorageFileCount,omitempty"`
	// The unit of the storage usage at the low-frequency storage tier.
	//
	// example:
	//
	// B
	LowFreqStorageUnit *string `json:"lowFreqStorageUnit,omitempty" xml:"lowFreqStorageUnit,omitempty"`
	// The project name.
	//
	// example:
	//
	// odps_project
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// The change rate of the total storage usage compared with that of the recent {$recentDays} days.
	//
	// example:
	//
	// 0
	Rate *float64 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The schema name.
	//
	// example:
	//
	// schema
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// The storage usage at the standard storage tier.
	//
	// example:
	//
	// 600
	StandardStorage *float64 `json:"standardStorage,omitempty" xml:"standardStorage,omitempty"`
	// The number of standard storage files.
	//
	// example:
	//
	// 2
	StandardStorageFileCount *int64 `json:"standardStorageFileCount,omitempty" xml:"standardStorageFileCount,omitempty"`
	// The unit of the storage usage at the standard storage tier.
	//
	// example:
	//
	// KB
	StandardStorageUnit *string `json:"standardStorageUnit,omitempty" xml:"standardStorageUnit,omitempty"`
	// The table storage type.
	//
	// 	- standard
	//
	// 	- lowfrequency
	//
	// 	- longterm
	//
	// 	- unknown: This value is returned when the table is a partitioned table. You can call the ListStoragePartitionsInfo operation to query the storage type of each partition.
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
	// >
	//
	// 	- Access behaviors include:
	//
	// 	- The table is used as the input table of an SQL task.
	//
	// 	- The table is downloaded by Tunnel.
	//
	// 	- The table is read by calling the Storage API. The partition granularity of the partitioned table is not available. Each time an access operation is performed, the access frequency is incremented by 1.
	//
	// 	- The data collection method is upgraded from July 2023. If the data is not accessed after the upgrade or is accessed by using ALGO jobs or the direct read method of Hologres, the access frequency cannot be collected.
	//
	// example:
	//
	// 10
	TotalFrequency *int64 `json:"totalFrequency,omitempty" xml:"totalFrequency,omitempty"`
	// The total amount of accessed data.
	//
	// >  The amount of data that is read by all access behaviors.
	//
	// example:
	//
	// 1
	TotalInputAmount *float64 `json:"totalInputAmount,omitempty" xml:"totalInputAmount,omitempty"`
	// The unit of the total amount of accessed data.
	//
	// example:
	//
	// GB
	TotalInputAmountUnit *string `json:"totalInputAmountUnit,omitempty" xml:"totalInputAmountUnit,omitempty"`
	// The total storage usage. For a partitioned table, this parameter indicates the sum of the storage usage of all partitions. If the storage types of partitions are different, the value is the sum of the storage usage of each storage type.
	//
	// example:
	//
	// 600
	TotalStorage *float64 `json:"totalStorage,omitempty" xml:"totalStorage,omitempty"`
	// The total number of files.
	//
	// example:
	//
	// 2
	TotalStorageFileCount *int64 `json:"totalStorageFileCount,omitempty" xml:"totalStorageFileCount,omitempty"`
	// The unit of storage usage.
	//
	// example:
	//
	// KB
	TotalStorageUnit *string `json:"totalStorageUnit,omitempty" xml:"totalStorageUnit,omitempty"`
}

func (s ListStorageTablesInfoResponseBodyDataStorageTableInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GoString() string {
	return s.String()
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetDate() *string {
	return s.Date
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetIsPartitioned() *bool {
	return s.IsPartitioned
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLastAccessTime() *int64 {
	return s.LastAccessTime
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLongTermStorage() *float64 {
	return s.LongTermStorage
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLongTermStorageFileCount() *int64 {
	return s.LongTermStorageFileCount
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLongTermStorageUnit() *string {
	return s.LongTermStorageUnit
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLowFreqStorage() *float64 {
	return s.LowFreqStorage
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLowFreqStorageFileCount() *int64 {
	return s.LowFreqStorageFileCount
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetLowFreqStorageUnit() *string {
	return s.LowFreqStorageUnit
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetRate() *float64 {
	return s.Rate
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetStandardStorage() *float64 {
	return s.StandardStorage
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetStandardStorageFileCount() *int64 {
	return s.StandardStorageFileCount
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetStandardStorageUnit() *string {
	return s.StandardStorageUnit
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetStorageType() *string {
	return s.StorageType
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTableName() *string {
	return s.TableName
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTotalFrequency() *int64 {
	return s.TotalFrequency
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTotalInputAmount() *float64 {
	return s.TotalInputAmount
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTotalInputAmountUnit() *string {
	return s.TotalInputAmountUnit
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTotalStorage() *float64 {
	return s.TotalStorage
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTotalStorageFileCount() *int64 {
	return s.TotalStorageFileCount
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) GetTotalStorageUnit() *string {
	return s.TotalStorageUnit
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetDate(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.Date = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetIsPartitioned(v bool) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.IsPartitioned = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLastAccessTime(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LastAccessTime = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLongTermStorage(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LongTermStorage = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLongTermStorageFileCount(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LongTermStorageFileCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLongTermStorageUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LongTermStorageUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLowFreqStorage(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LowFreqStorage = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLowFreqStorageFileCount(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LowFreqStorageFileCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetLowFreqStorageUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.LowFreqStorageUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetProjectName(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.ProjectName = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetRate(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.Rate = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetSchemaName(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.SchemaName = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetStandardStorage(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.StandardStorage = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetStandardStorageFileCount(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.StandardStorageFileCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetStandardStorageUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.StandardStorageUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetStorageType(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.StorageType = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTableName(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TableName = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalFrequency(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalFrequency = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalInputAmount(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalInputAmount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalInputAmountUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalInputAmountUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalStorage(v float64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalStorage = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalStorageFileCount(v int64) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalStorageFileCount = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) SetTotalStorageUnit(v string) *ListStorageTablesInfoResponseBodyDataStorageTableInfoList {
	s.TotalStorageUnit = &v
	return s
}

func (s *ListStorageTablesInfoResponseBodyDataStorageTableInfoList) Validate() error {
	return dara.Validate(s)
}
