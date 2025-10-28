// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageAnalysisResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetStorageAnalysisResultResponseBody
	GetCode() *int64
	SetData(v *GetStorageAnalysisResultResponseBodyData) *GetStorageAnalysisResultResponseBody
	GetData() *GetStorageAnalysisResultResponseBodyData
	SetMessage(v string) *GetStorageAnalysisResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStorageAnalysisResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStorageAnalysisResultResponseBody
	GetSuccess() *bool
}

type GetStorageAnalysisResultResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetStorageAnalysisResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request is successful, **Successful*	- is returned. Otherwise, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStorageAnalysisResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStorageAnalysisResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageAnalysisResultResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetStorageAnalysisResultResponseBody) GetData() *GetStorageAnalysisResultResponseBodyData {
	return s.Data
}

func (s *GetStorageAnalysisResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStorageAnalysisResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStorageAnalysisResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStorageAnalysisResultResponseBody) SetCode(v int64) *GetStorageAnalysisResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBody) SetData(v *GetStorageAnalysisResultResponseBodyData) *GetStorageAnalysisResultResponseBody {
	s.Data = v
	return s
}

func (s *GetStorageAnalysisResultResponseBody) SetMessage(v string) *GetStorageAnalysisResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBody) SetRequestId(v string) *GetStorageAnalysisResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBody) SetSuccess(v bool) *GetStorageAnalysisResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStorageAnalysisResultResponseBodyData struct {
	// The number of databases that have been analyzed.
	//
	// example:
	//
	// 2
	AnalyzedDbCount *int64 `json:"AnalyzedDbCount,omitempty" xml:"AnalyzedDbCount,omitempty"`
	// The details of storage analysis.
	StorageAnalysisResult *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult `json:"StorageAnalysisResult,omitempty" xml:"StorageAnalysisResult,omitempty" type:"Struct"`
	// Indicates whether the task is complete.
	//
	// example:
	//
	// true
	TaskFinish *bool `json:"TaskFinish,omitempty" xml:"TaskFinish,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 910f83f4b96df0524ddc5749f615****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task progress.
	//
	// >  Valid values are integers that range from 0 to 100.
	//
	// example:
	//
	// 50
	TaskProgress *int64 `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty"`
	// The status of the storage analysis task. Valid values:
	//
	// 	- **INIT**: The task is being initialized.
	//
	// 	- **PENDING**: The task is being queued for execution.
	//
	// 	- **RECEIVED**: The task is received for execution.
	//
	// 	- **RUNNING**: The task is being executed.
	//
	// 	- **RETRY**: The task is being retried.
	//
	// 	- **SUCCESS**: The task succeeds.
	//
	// 	- **FAILURE**: The task fails.
	//
	// example:
	//
	// RUNNING
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// Indicates whether the task is successful.
	//
	// example:
	//
	// true
	TaskSuccess *bool `json:"TaskSuccess,omitempty" xml:"TaskSuccess,omitempty"`
	// The number of databases that need to be analyzed in the storage analysis task.
	//
	// example:
	//
	// 32
	TotalDbCount *int64 `json:"TotalDbCount,omitempty" xml:"TotalDbCount,omitempty"`
}

func (s GetStorageAnalysisResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStorageAnalysisResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStorageAnalysisResultResponseBodyData) GetAnalyzedDbCount() *int64 {
	return s.AnalyzedDbCount
}

func (s *GetStorageAnalysisResultResponseBodyData) GetStorageAnalysisResult() *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult {
	return s.StorageAnalysisResult
}

func (s *GetStorageAnalysisResultResponseBodyData) GetTaskFinish() *bool {
	return s.TaskFinish
}

func (s *GetStorageAnalysisResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetStorageAnalysisResultResponseBodyData) GetTaskProgress() *int64 {
	return s.TaskProgress
}

func (s *GetStorageAnalysisResultResponseBodyData) GetTaskState() *string {
	return s.TaskState
}

func (s *GetStorageAnalysisResultResponseBodyData) GetTaskSuccess() *bool {
	return s.TaskSuccess
}

func (s *GetStorageAnalysisResultResponseBodyData) GetTotalDbCount() *int64 {
	return s.TotalDbCount
}

func (s *GetStorageAnalysisResultResponseBodyData) SetAnalyzedDbCount(v int64) *GetStorageAnalysisResultResponseBodyData {
	s.AnalyzedDbCount = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyData) SetStorageAnalysisResult(v *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) *GetStorageAnalysisResultResponseBodyData {
	s.StorageAnalysisResult = v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyData) SetTaskFinish(v bool) *GetStorageAnalysisResultResponseBodyData {
	s.TaskFinish = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyData) SetTaskId(v string) *GetStorageAnalysisResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyData) SetTaskProgress(v int64) *GetStorageAnalysisResultResponseBodyData {
	s.TaskProgress = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyData) SetTaskState(v string) *GetStorageAnalysisResultResponseBodyData {
	s.TaskState = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyData) SetTaskSuccess(v bool) *GetStorageAnalysisResultResponseBodyData {
	s.TaskSuccess = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyData) SetTotalDbCount(v int64) *GetStorageAnalysisResultResponseBodyData {
	s.TotalDbCount = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyData) Validate() error {
	if s.StorageAnalysisResult != nil {
		if err := s.StorageAnalysisResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult struct {
	// The reason why the analysis on the database and table fails.
	//
	// 	- **DB_OR_TABLE_NOT_EXIST**: The specified database or table does not exist.
	//
	// 	- **DB_NOT_EXIST**: The specified database does not exist.
	//
	// example:
	//
	// DB_NOT_EXIST
	AnalysisErrorType *string `json:"AnalysisErrorType,omitempty" xml:"AnalysisErrorType,omitempty"`
	// Indicates whether the analysis on the database and table is successful.
	//
	// example:
	//
	// true
	AnalysisSuccess *bool `json:"AnalysisSuccess,omitempty" xml:"AnalysisSuccess,omitempty"`
	// The estimated average daily growth of the used storage space in the previous seven days. Unit: bytes.
	//
	// example:
	//
	// 0
	DailyIncrement *int64 `json:"DailyIncrement,omitempty" xml:"DailyIncrement,omitempty"`
	// The estimated number of days for which the remaining storage space is available.
	//
	// example:
	//
	// 99
	EstimateAvailableDays *int64 `json:"EstimateAvailableDays,omitempty" xml:"EstimateAvailableDays,omitempty"`
	// The items to be optimized, which are generated based on DAS default rules. You can ignore these items based on your business requirements, and create custom rules to generate items to be optimized based on other basic data that is returned.
	NeedOptimizeItemList []*GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList `json:"NeedOptimizeItemList,omitempty" xml:"NeedOptimizeItemList,omitempty" type:"Repeated"`
	// The information about the table.
	TableStats []*GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats `json:"TableStats,omitempty" xml:"TableStats,omitempty" type:"Repeated"`
	// The size of remaining storage.
	//
	// >  Unit: bytes.
	//
	// example:
	//
	// 146403229696
	TotalFreeStorageSize *int64 `json:"TotalFreeStorageSize,omitempty" xml:"TotalFreeStorageSize,omitempty"`
	// The total size of instance storage.
	//
	// >  Unit: bytes.
	//
	// example:
	//
	// 214748364800
	TotalStorageSize *int64 `json:"TotalStorageSize,omitempty" xml:"TotalStorageSize,omitempty"`
	// The size of used storage.
	//
	// >  Unit: bytes.
	//
	// example:
	//
	// 68345135104
	TotalUsedStorageSize *int64 `json:"TotalUsedStorageSize,omitempty" xml:"TotalUsedStorageSize,omitempty"`
}

func (s GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) String() string {
	return dara.Prettify(s)
}

func (s GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) GoString() string {
	return s.String()
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) GetAnalysisErrorType() *string {
	return s.AnalysisErrorType
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) GetAnalysisSuccess() *bool {
	return s.AnalysisSuccess
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) GetDailyIncrement() *int64 {
	return s.DailyIncrement
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) GetEstimateAvailableDays() *int64 {
	return s.EstimateAvailableDays
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) GetNeedOptimizeItemList() []*GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList {
	return s.NeedOptimizeItemList
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) GetTableStats() []*GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	return s.TableStats
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) GetTotalFreeStorageSize() *int64 {
	return s.TotalFreeStorageSize
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) GetTotalStorageSize() *int64 {
	return s.TotalStorageSize
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) GetTotalUsedStorageSize() *int64 {
	return s.TotalUsedStorageSize
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) SetAnalysisErrorType(v string) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult {
	s.AnalysisErrorType = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) SetAnalysisSuccess(v bool) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult {
	s.AnalysisSuccess = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) SetDailyIncrement(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult {
	s.DailyIncrement = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) SetEstimateAvailableDays(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult {
	s.EstimateAvailableDays = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) SetNeedOptimizeItemList(v []*GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult {
	s.NeedOptimizeItemList = v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) SetTableStats(v []*GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult {
	s.TableStats = v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) SetTotalFreeStorageSize(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult {
	s.TotalFreeStorageSize = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) SetTotalStorageSize(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult {
	s.TotalStorageSize = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) SetTotalUsedStorageSize(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult {
	s.TotalUsedStorageSize = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResult) Validate() error {
	if s.NeedOptimizeItemList != nil {
		for _, item := range s.NeedOptimizeItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TableStats != nil {
		for _, item := range s.TableStats {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList struct {
	// The data associated with the items to be optimized, which is in the JSON format.
	//
	// example:
	//
	// {
	//
	//     "autoIncrementCurrentValue": 2147483647,
	//
	//     "autoIncrementRatio": 1,
	//
	//     "dbName": "testdb01",
	//
	//     "maximumValue": 2147483647,
	//
	//     "columnName": "id",
	//
	//     "tableName": "test_table"
	//
	// }
	AssociatedData *string `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testdb01
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The optimization suggestion. Valid values:
	//
	// 	- **NEED_ANALYZE_TABLE**: You can execute the `ANALYZE TABLE` statement on the table during off-peak hours. This is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// 	- **NEED_OPTIMIZE_TABLE**: You can reclaim fragments during off-peak hours.
	//
	// 	- **CHANGE_TABLE_ENGINE_IF_NECESSARY**: Change the storage engine type of a table after risk assessment. This is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// 	- **AUTO_INCREMENT_ID_BE_TO_RUN_OUT**: Pay attention to the usage of auto-increment IDs. This is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// 	- **DUPLICATE_INDEX**: Optimize indexes of tables. This is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// 	- **TABLE_SIZE**: Pay attention to the table size. This is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// 	- **TABLE_ROWS_AND_AVG_ROW_LENGTH**: Pay attention to the number of rows in a table and the average row length. This is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// 	- **STORAGE_USED_PERCENT**: Pay attention to the space usage to prevent the instance from being locked if the instance is full.
	//
	// example:
	//
	// NEED_OPTIMIZE_TABLE
	OptimizeAdvice *string `json:"OptimizeAdvice,omitempty" xml:"OptimizeAdvice,omitempty"`
	// The item to be optimized. Valid values:
	//
	// 	- **NEED_ANALYZE_TABLE**: tables whose storage statistics obtained from `information_schema.tables` are 50 GB larger or smaller than the physical file sizes. This is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// 	- **NEED_OPTIMIZE_TABLE**: tables whose space fragments are larger than 6 GB and whose fragmentation rates are greater than 30%. The fragmentation rate of a table is generally calculated based on the following formulas:
	//
	//     	- ApsaraDB RDS for MySQL and PolarDB for MySQL: `Fragmentation rate = DataFree/(DataSize + IndexSize + DataFree)`. In this topic, PhyTotalSize = DataSize + IndexSize + DataFree. Thus, the fragmentation rate can be calculated based on the following formula: `Fragmentation rate = DataFree/PhyTotalSize`.
	//
	//     	- ApsaraDB for MongoDB: `Fragmentation rate = FragmentSize/PhyTotalSize`.
	//
	// 	- **TABLE_ENGINE**: tables whose storage engines are not InnoDB or XEngine. This is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// 	- **AUTO_INCREMENT_ID_BE_TO_RUN_OUT**: tables whose usages of auto-increment IDs exceed 80%. This is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// 	- **DUPLICATE_INDEX**: tables whose indexes are redundant or duplicate. This is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// 	- **TABLE_SIZE**: single tables whose sizes are larger than 50 GB. This is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// 	- **TABLE_ROWS_AND_AVG_ROW_LENGTH**: single tables that contain more than 5 million rows and whose average row lengths exceed 10 KB. This is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// 	- **TOTAL_DATA_FREE**: instances whose reclaimable space is larger than 60 GB and whose total fragmentation rate is larger than 5%.
	//
	// 	- **STORAGE_USED_PERCENT**: instances whose space usage is larger than 90%.
	//
	// example:
	//
	// NEED_OPTIMIZE_TABLE
	OptimizeItemName *string `json:"OptimizeItemName,omitempty" xml:"OptimizeItemName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) String() string {
	return dara.Prettify(s)
}

func (s GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) GoString() string {
	return s.String()
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) GetAssociatedData() *string {
	return s.AssociatedData
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) GetDbName() *string {
	return s.DbName
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) GetOptimizeAdvice() *string {
	return s.OptimizeAdvice
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) GetOptimizeItemName() *string {
	return s.OptimizeItemName
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) GetTableName() *string {
	return s.TableName
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) SetAssociatedData(v string) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList {
	s.AssociatedData = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) SetDbName(v string) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList {
	s.DbName = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) SetOptimizeAdvice(v string) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList {
	s.OptimizeAdvice = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) SetOptimizeItemName(v string) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList {
	s.OptimizeItemName = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) SetTableName(v string) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList {
	s.TableName = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultNeedOptimizeItemList) Validate() error {
	return dara.Validate(s)
}

type GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats struct {
	// The average length of rows. Unit: bytes.
	//
	// example:
	//
	// 154
	AvgRowLength *int64 `json:"AvgRowLength,omitempty" xml:"AvgRowLength,omitempty"`
	// The size of space fragments. Unit: bytes.
	//
	// >  This parameter is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters. The fragmentation rate of a table is generally calculated based on the following formula: `Fragmentation rate = DataFree/(DataSize + IndexSize + DataFree)`. In this topic, `Fragmentation rate = DataFree/PhyTotalSize`.
	//
	// example:
	//
	// 7340032
	DataFree *int64 `json:"DataFree,omitempty" xml:"DataFree,omitempty"`
	// 	- For ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters, this parameter indicates the amount of space occupied by data. Unit: bytes.
	//
	// 	- For ApsaraDB for MongoDB instances, this parameter indicates the size of uncompressed data, that is, the amount of data. Unit: bytes.
	//
	// example:
	//
	// 1982857216
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testdb01
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The type of the storage engine used by the table.
	//
	// >  This parameter is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// example:
	//
	// InnoDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The size of space that can be reclaimed. Unit: bytes.
	//
	// >  This parameter is applicable only to ApsaraDB for MongoDB instances. `Fragmentation rate = FragmentSize/PhyTotalSize`.
	//
	// example:
	//
	// 362221568
	FragmentSize *int64 `json:"FragmentSize,omitempty" xml:"FragmentSize,omitempty"`
	// The storage space occupied by indexes. Unit: bytes.
	//
	// example:
	//
	// 1022296064
	IndexSize *int64 `json:"IndexSize,omitempty" xml:"IndexSize,omitempty"`
	// The storage space of the table. Unit: bytes.
	//
	// >  For ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters, the value of the parameter is the sum of **DataSize**, **IndexSize**, and **DataFree**. For ApsaraDB for MongoDB instances, the value of this parameter is the sum of **DataSize*	- and **IndexSize**.
	//
	// example:
	//
	// 3012493312
	PhyTotalSize *int64 `json:"PhyTotalSize,omitempty" xml:"PhyTotalSize,omitempty"`
	// The physical file size of the table. Unit: bytes.
	//
	// >  This parameter is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters. Data of specific database instances cannot be obtained due to deployment mode.
	//
	// example:
	//
	// 3057655808
	PhysicalFileSize *int64 `json:"PhysicalFileSize,omitempty" xml:"PhysicalFileSize,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The number of rows in the table.
	//
	// example:
	//
	// 12794732
	TableRows *int64 `json:"TableRows,omitempty" xml:"TableRows,omitempty"`
	// The type of the table.
	//
	// >  This parameter is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
	//
	// example:
	//
	// BASE TABLE
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	// 	- For ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters, this parameter indicates the amount of space occupied by table data and indexes. Unit: bytes. The value is the sum of **DataSize*	- and **IndexSize**.
	//
	// 	- For ApsaraDB for MongoDB instances, this parameter indicates the actual size of space allocated by Block Manager. Unit: Bytes. The compression ratio of an ApsaraDB for MongoDB instance is calculated based on the following formula: `Compression ratio = TotalSize/DataSize`.
	//
	// example:
	//
	// 3005153280
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) String() string {
	return dara.Prettify(s)
}

func (s GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GoString() string {
	return s.String()
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetAvgRowLength() *int64 {
	return s.AvgRowLength
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetDataFree() *int64 {
	return s.DataFree
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetDataSize() *int64 {
	return s.DataSize
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetDbName() *string {
	return s.DbName
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetEngine() *string {
	return s.Engine
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetFragmentSize() *int64 {
	return s.FragmentSize
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetIndexSize() *int64 {
	return s.IndexSize
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetPhyTotalSize() *int64 {
	return s.PhyTotalSize
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetPhysicalFileSize() *int64 {
	return s.PhysicalFileSize
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetTableName() *string {
	return s.TableName
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetTableRows() *int64 {
	return s.TableRows
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetTableType() *string {
	return s.TableType
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetAvgRowLength(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.AvgRowLength = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetDataFree(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.DataFree = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetDataSize(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.DataSize = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetDbName(v string) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.DbName = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetEngine(v string) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.Engine = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetFragmentSize(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.FragmentSize = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetIndexSize(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.IndexSize = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetPhyTotalSize(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.PhyTotalSize = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetPhysicalFileSize(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.PhysicalFileSize = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetTableName(v string) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.TableName = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetTableRows(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.TableRows = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetTableType(v string) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.TableType = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) SetTotalSize(v int64) *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats {
	s.TotalSize = &v
	return s
}

func (s *GetStorageAnalysisResultResponseBodyDataStorageAnalysisResultTableStats) Validate() error {
	return dara.Validate(s)
}
