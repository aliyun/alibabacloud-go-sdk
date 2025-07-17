// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCorrectOrderDetail(v *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) *GetDataCorrectOrderDetailResponseBody
	GetDataCorrectOrderDetail() *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail
	SetErrorCode(v string) *GetDataCorrectOrderDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataCorrectOrderDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataCorrectOrderDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCorrectOrderDetailResponseBody
	GetSuccess() *bool
}

type GetDataCorrectOrderDetailResponseBody struct {
	// The information about the data change ticket.
	DataCorrectOrderDetail *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail `json:"DataCorrectOrderDetail,omitempty" xml:"DataCorrectOrderDetail,omitempty" type:"Struct"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
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

func (s GetDataCorrectOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBody) GetDataCorrectOrderDetail() *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail {
	return s.DataCorrectOrderDetail
}

func (s *GetDataCorrectOrderDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataCorrectOrderDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataCorrectOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCorrectOrderDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCorrectOrderDetailResponseBody) SetDataCorrectOrderDetail(v *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) *GetDataCorrectOrderDetailResponseBody {
	s.DataCorrectOrderDetail = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBody) SetErrorCode(v string) *GetDataCorrectOrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBody) SetErrorMessage(v string) *GetDataCorrectOrderDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBody) SetRequestId(v string) *GetDataCorrectOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBody) SetSuccess(v bool) *GetDataCorrectOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail struct {
	// The configurations of the ticket. This parameter is used to store the configuration information specific to a data change ticket type.
	ConfigDetail *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail `json:"ConfigDetail,omitempty" xml:"ConfigDetail,omitempty" type:"Struct"`
	// The information about the database in which data is changed.
	DatabaseList *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList `json:"DatabaseList,omitempty" xml:"DatabaseList,omitempty" type:"Struct"`
	// The execution mode of the ticket after the ticket is approved. Valid values:
	//
	// - **COMMITOR**: The data change is performed by the user who submits the ticket.
	//
	// - **AUTO**: The data change is automatically performed after the ticket is approved.
	//
	// - **LAST_AUDITOR**: The data change is performed by the last approver of the ticket.
	//
	// example:
	//
	// COMMITOR
	ExecMode *string `json:"ExecMode,omitempty" xml:"ExecMode,omitempty"`
	// The details of the ticket.
	OrderDetail *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail `json:"OrderDetail,omitempty" xml:"OrderDetail,omitempty" type:"Struct"`
	// The precheck details of the ticket.
	PreCheckDetail *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail `json:"PreCheckDetail,omitempty" xml:"PreCheckDetail,omitempty" type:"Struct"`
	// The specific state of the data change ticket. Valid values:
	//
	// >  The state of the ticket is not exactly equivalent to the status code for the ticket. To query the status code of the ticket, you can call the [GetOrderBaseInfo](https://help.aliyun.com/document_detail/465868.html) operation and check the value of StatusCode in the response.
	//
	// 	- **new**: The ticket is created.
	//
	// 	- **precheck**: The ticket is in the pre-check phase.
	//
	// 	- **precheckFailed**: The ticket failed to pass the precheck.
	//
	// 	- **precheck_success**: The ticket passes the precheck and waits to be submitted for approval.
	//
	// 	- **toaudit**: The ticket is being reviewed.
	//
	// 	- **Approved**: The ticket is approved.
	//
	// 	- **reject**: The ticket is rejected.
	//
	// 	- **waiting**: The task is submitted and waits to be scheduled.
	//
	// 	- **processing**: The task is being executed.
	//
	// 	- **Success**: The task is successful.
	//
	// example:
	//
	// approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) GetConfigDetail() *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	return s.ConfigDetail
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) GetDatabaseList() *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList {
	return s.DatabaseList
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) GetExecMode() *string {
	return s.ExecMode
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) GetOrderDetail() *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	return s.OrderDetail
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) GetPreCheckDetail() *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail {
	return s.PreCheckDetail
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) GetStatus() *string {
	return s.Status
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) SetConfigDetail(v *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail {
	s.ConfigDetail = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) SetDatabaseList(v *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail {
	s.DatabaseList = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) SetExecMode(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail {
	s.ExecMode = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) SetOrderDetail(v *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail {
	s.OrderDetail = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) SetPreCheckDetail(v *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail {
	s.PreCheckDetail = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) SetStatus(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail {
	s.Status = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetail) Validate() error {
	return dara.Validate(s)
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail struct {
	// Indicates whether the task is a scheduled task for historical data cleanup. This parameter is a reserved parameter and is valid only if the value of DetailType is CRON_CLEAR_DATA.
	//
	// example:
	//
	// true
	Cron *bool `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The number of times the scheduled task is run. This parameter is valid only if the value of DetailType is CRON_CLEAR_DATA.
	//
	// example:
	//
	// 0
	CronCallTimes *int32 `json:"CronCallTimes,omitempty" xml:"CronCallTimes,omitempty"`
	// The additional configuration information about historical data cleanup. This parameter is valid only if the value of DetailType is CRON_CLEAR_DATA.
	CronExtConfig *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig `json:"CronExtConfig,omitempty" xml:"CronExtConfig,omitempty" type:"Struct"`
	// The CRON expression of the scheduled task. This parameter is valid only if the value of DetailType is CRON_CLEAR_DATA.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	CronFormat *string `json:"CronFormat,omitempty" xml:"CronFormat,omitempty"`
	// The time when the task was last run.
	//
	// example:
	//
	// 2024-04-19 02:00:00.0
	CronLastCallStartTime *string `json:"CronLastCallStartTime,omitempty" xml:"CronLastCallStartTime,omitempty"`
	// The time when the task is run next time. This parameter is returned only if the value of CronStatus is SUCCESS.
	//
	// example:
	//
	// 2024-04-19 02:00:00
	CronNextCallTime *string `json:"CronNextCallTime,omitempty" xml:"CronNextCallTime,omitempty"`
	// The state of the scheduled task. If this parameter is empty, the task is not run. Valid values:
	//
	// 	- PAUSE: The task is suspended.
	//
	// 	- WAITING: The task is waiting to be run.
	//
	// 	- SUCCESS: The task is run.
	//
	// example:
	//
	// SUCCESS
	CronStatus *string `json:"CronStatus,omitempty" xml:"CronStatus,omitempty"`
	// The name of the table to which data is to be imported. This parameter is valid only if the value of DetailType is BIG_FILE. If the value of FileType is SQL, this parameter is empty.
	//
	// example:
	//
	// tb_import_tb_name
	CsvTableName *string `json:"CsvTableName,omitempty" xml:"CsvTableName,omitempty"`
	// The ID of the current data change task. This is a reserved parameter and can be ignored.
	//
	// example:
	//
	// 13***
	CurrentTaskId *int64 `json:"CurrentTaskId,omitempty" xml:"CurrentTaskId,omitempty"`
	// The type of the ticket. Valid values:
	//
	// 	- COMMON: regular data change.
	//
	// 	- CHUNK_DML: lock-free data change.
	//
	// 	- BIG_FILE: large data import.
	//
	// 	- CRON_CLEAR_DATA: historical data cleanup.
	//
	// 	- PROCEDURE: programmable object change.
	//
	// example:
	//
	// BIG_FILE
	DetailType *string `json:"DetailType,omitempty" xml:"DetailType,omitempty"`
	// The execution duration of the scheduled task. Unit: hour. This parameter is valid only if the value of DetailType is CRON_CLEAR_DATA. If the value is greater than 0, an execution duration is set.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The encoding method of the file. This parameter may be empty, which indicates the value of AUTO. Valid values:
	//
	// 	- **AUTO**: automatic identification.
	//
	// 	- **UTF-8**: UTF-8 encoding.
	//
	// 	- **GBK**: GBK encoding.
	//
	// 	- **ISO-8859-1**: ISO-8859-1 encoding.
	//
	// example:
	//
	// UTF-8
	FileEncoding *string `json:"FileEncoding,omitempty" xml:"FileEncoding,omitempty"`
	// The type of the file to be imported. This parameter is valid if the value of DetailType is BIG_FILE. Valid values:
	//
	// 	- **SQL**: an SQL file.
	//
	// 	- **CSV**: a CSV file.
	//
	// 	- **EXCEL**: an Excel file.
	//
	// 	- **JSON**: a JSON file, which is supported only by MongoDB databases.
	//
	// example:
	//
	// CSV
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The additional configuration information about data import. This parameter is valid if the value of DetailType is BIG_FILE.
	ImportExtConfig *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig `json:"ImportExtConfig,omitempty" xml:"ImportExtConfig,omitempty" type:"Struct"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetCron() *bool {
	return s.Cron
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetCronCallTimes() *int32 {
	return s.CronCallTimes
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetCronExtConfig() *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig {
	return s.CronExtConfig
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetCronFormat() *string {
	return s.CronFormat
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetCronLastCallStartTime() *string {
	return s.CronLastCallStartTime
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetCronNextCallTime() *string {
	return s.CronNextCallTime
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetCronStatus() *string {
	return s.CronStatus
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetCsvTableName() *string {
	return s.CsvTableName
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetCurrentTaskId() *int64 {
	return s.CurrentTaskId
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetDetailType() *string {
	return s.DetailType
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetDuration() *int32 {
	return s.Duration
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetFileEncoding() *string {
	return s.FileEncoding
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetFileType() *string {
	return s.FileType
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) GetImportExtConfig() *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig {
	return s.ImportExtConfig
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetCron(v bool) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.Cron = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetCronCallTimes(v int32) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.CronCallTimes = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetCronExtConfig(v *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.CronExtConfig = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetCronFormat(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.CronFormat = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetCronLastCallStartTime(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.CronLastCallStartTime = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetCronNextCallTime(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.CronNextCallTime = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetCronStatus(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.CronStatus = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetCsvTableName(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.CsvTableName = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetCurrentTaskId(v int64) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.CurrentTaskId = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetDetailType(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.DetailType = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetDuration(v int32) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.Duration = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetFileEncoding(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.FileEncoding = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetFileType(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.FileType = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) SetImportExtConfig(v *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail {
	s.ImportExtConfig = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetail) Validate() error {
	return dara.Validate(s)
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig struct {
	// The number of times defragmentation is performed. This parameter is valid only if the value of OptimizeTableAfterEveryClearTimes is greater than 0.
	//
	// example:
	//
	// 0
	CurrentClearTaskCount *int32 `json:"CurrentClearTaskCount,omitempty" xml:"CurrentClearTaskCount,omitempty"`
	// Indicates whether the Periodically Optimize Table feature is enabled. Valid values:
	//
	// 	- **0*	- (default): The feature is disabled.
	//
	// 	- **A value greater than 0**: The feature is enabled. The value indicates the number of cleanups after which the system performs defragmentation.
	//
	// example:
	//
	// 0
	OptimizeTableAfterEveryClearTimes *int32 `json:"OptimizeTableAfterEveryClearTimes,omitempty" xml:"OptimizeTableAfterEveryClearTimes,omitempty"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig) GetCurrentClearTaskCount() *int32 {
	return s.CurrentClearTaskCount
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig) GetOptimizeTableAfterEveryClearTimes() *int32 {
	return s.OptimizeTableAfterEveryClearTimes
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig) SetCurrentClearTaskCount(v int32) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig {
	s.CurrentClearTaskCount = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig) SetOptimizeTableAfterEveryClearTimes(v int32) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig {
	s.OptimizeTableAfterEveryClearTimes = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailCronExtConfig) Validate() error {
	return dara.Validate(s)
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig struct {
	// Indicates whether the first row of the CSV file contains field names. Valid values:
	//
	// 	- **true**: The first row in the CSV file contains field names.
	//
	// 	- **false**: The first row in the CSV file contains data.
	//
	// >  This parameter is valid if the value of **FileType*	- is **CSV*	- or **EXCEL**.
	//
	// example:
	//
	// true
	CsvFirstRowIsColumnDef *bool `json:"CsvFirstRowIsColumnDef,omitempty" xml:"CsvFirstRowIsColumnDef,omitempty"`
	// Indicates whether an error that occurs is ignored. Valid values:
	//
	// 	- **true**: If an error occurs when SQL statements are being executed, DMS skips the current SQL statement and continues to execute subsequent SQL statements.
	//
	// 	- **false**: If an error occurs when SQL statements are being executed, DMS stops executing subsequent SQL statements.
	//
	// example:
	//
	// false
	IgnoreError *bool `json:"IgnoreError,omitempty" xml:"IgnoreError,omitempty"`
	// The import mode. Valid values:
	//
	// 	- **FAST_MODE**: fast mode. In the Execute step, the uploaded file is read and SQL statements are executed to import data to the specified destination database. Compared with the security mode, this mode can be used to import data in a less secure but more efficient manner.
	//
	// 	- **SAFE_MODE**: security mode. In the Precheck step, the uploaded file is parsed, and SQL statements or CSV file data is cached. In the Execute step, the cached SQL statements are read and executed to import data, or the cached CSV file data is read and imported to the specified destination database. Compared with the fast mode, this mode can be used to import data in a more secure but less efficient manner.
	//
	// example:
	//
	// FAST_MODE
	ImportMode *string `json:"ImportMode,omitempty" xml:"ImportMode,omitempty"`
	// The mode in which data is to be imported to the destination table. Valid values:
	//
	// 	- **INSERT**: The database checks the primary key during data insertion. If the primary key is duplicated, an error is reported.
	//
	// 	- **INSERT_IGNORE**: If the imported data contains data records that are the same as those in the destination table, the new data records are ignored.
	//
	// 	- **REPLACE_INTO**: If the imported data contains a row that has the same value for the primary key or unique index as an existing row in the destination table, the system deletes the existing row and inserts the new row into the destination table.
	//
	// >  This parameter is valid if the value of FileType is CSV or EXCEL.
	//
	// example:
	//
	// INSERT
	InsertType *string `json:"InsertType,omitempty" xml:"InsertType,omitempty"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig) GetCsvFirstRowIsColumnDef() *bool {
	return s.CsvFirstRowIsColumnDef
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig) GetIgnoreError() *bool {
	return s.IgnoreError
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig) GetImportMode() *string {
	return s.ImportMode
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig) GetInsertType() *string {
	return s.InsertType
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig) SetCsvFirstRowIsColumnDef(v bool) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig {
	s.CsvFirstRowIsColumnDef = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig) SetIgnoreError(v bool) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig {
	s.IgnoreError = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig) SetImportMode(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig {
	s.ImportMode = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig) SetInsertType(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig {
	s.InsertType = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailConfigDetailImportExtConfig) Validate() error {
	return dara.Validate(s)
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList struct {
	Database []*GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList) GetDatabase() []*GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase {
	return s.Database
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList) SetDatabase(v []*GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList {
	s.Database = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseList) Validate() error {
	return dara.Validate(s)
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase struct {
	// The database ID.
	//
	// example:
	//
	// 1860****
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The engine of the database.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the database belongs. Valid values:
	//
	// 	- product: production environment.
	//
	// 	- dev: development environment.
	//
	// 	- pre: pre-release environment.
	//
	// 	- test: test environment.
	//
	// 	- sit: system integration testing (SIT) environment
	//
	// 	- uat: user acceptance testing (UAT) environment.
	//
	// 	- pet: stress testing environment.
	//
	// 	- stag: staging environment.
	//
	// example:
	//
	// product
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true.**: The database is a logical database.
	//
	// 	- **false**: The database is a physical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// xxx@xxx:3306
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) GetDbId() *int32 {
	return s.DbId
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) GetDbType() *string {
	return s.DbType
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) GetEnvType() *string {
	return s.EnvType
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) GetLogic() *bool {
	return s.Logic
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) GetSearchName() *string {
	return s.SearchName
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) SetDbId(v int32) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase {
	s.DbId = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) SetDbType(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase {
	s.DbType = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) SetEnvType(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase {
	s.EnvType = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) SetLogic(v bool) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase {
	s.Logic = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) SetSearchName(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase {
	s.SearchName = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailDatabaseListDatabase) Validate() error {
	return dara.Validate(s)
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail struct {
	// The number of affected rows that is obtained by the precheck.
	//
	// example:
	//
	// 100
	ActualAffectRows *int64 `json:"ActualAffectRows,omitempty" xml:"ActualAffectRows,omitempty"`
	// The name of the attachment that contains the SQL statements used to change data.
	//
	// example:
	//
	// xxx
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// The category of the reason for the data change.
	//
	// example:
	//
	// test
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The estimated number of affected rows.
	//
	// example:
	//
	// 100
	EstimateAffectRows *int64 `json:"EstimateAffectRows,omitempty" xml:"EstimateAffectRows,omitempty"`
	// The executed SQL statements.
	//
	// example:
	//
	// update t1 set name = \\"xxx\\" where id <= 100
	ExeSQL *string `json:"ExeSQL,omitempty" xml:"ExeSQL,omitempty"`
	// Indicates whether the precheck result is ignored. Valid values:
	//
	// - **true**: The precheck result is ignored.
	//
	// - **false**: The precheck result is not ignored.
	//
	// example:
	//
	// false
	IgnoreAffectRows *bool `json:"IgnoreAffectRows,omitempty" xml:"IgnoreAffectRows,omitempty"`
	// The reason why the precheck result is ignored.
	//
	// example:
	//
	// test
	IgnoreAffectRowsReason *string `json:"IgnoreAffectRowsReason,omitempty" xml:"IgnoreAffectRowsReason,omitempty"`
	// The name of the attachment that contains the SQL statements used to roll back the data change.
	//
	// example:
	//
	// test
	RbAttachmentName *string `json:"RbAttachmentName,omitempty" xml:"RbAttachmentName,omitempty"`
	// The SQL statements used to roll back the data change.
	//
	// example:
	//
	// empty
	RbSQL *string `json:"RbSQL,omitempty" xml:"RbSQL,omitempty"`
	// The format of the SQL statements used to roll back the data change. Valid values:
	//
	// - **TEXT**: text
	//
	// - **ATTACHMENT**: attachment
	//
	// example:
	//
	// text
	RbSQLType *string `json:"RbSQLType,omitempty" xml:"RbSQLType,omitempty"`
	// The format of the SQL statements used to change data. Valid values:
	//
	// - **TEXT**: text
	//
	// - **ATTACHMENT**: attachment
	//
	// example:
	//
	// text
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GetActualAffectRows() *int64 {
	return s.ActualAffectRows
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GetClassify() *string {
	return s.Classify
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GetEstimateAffectRows() *int64 {
	return s.EstimateAffectRows
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GetExeSQL() *string {
	return s.ExeSQL
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GetIgnoreAffectRows() *bool {
	return s.IgnoreAffectRows
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GetIgnoreAffectRowsReason() *string {
	return s.IgnoreAffectRowsReason
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GetRbAttachmentName() *string {
	return s.RbAttachmentName
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GetRbSQL() *string {
	return s.RbSQL
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GetRbSQLType() *string {
	return s.RbSQLType
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) GetSqlType() *string {
	return s.SqlType
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetActualAffectRows(v int64) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.ActualAffectRows = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetAttachmentName(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.AttachmentName = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetClassify(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.Classify = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetEstimateAffectRows(v int64) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.EstimateAffectRows = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetExeSQL(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.ExeSQL = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetIgnoreAffectRows(v bool) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.IgnoreAffectRows = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetIgnoreAffectRowsReason(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.IgnoreAffectRowsReason = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetRbAttachmentName(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.RbAttachmentName = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetRbSQL(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.RbSQL = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetRbSQLType(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.RbSQLType = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) SetSqlType(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail {
	s.SqlType = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailOrderDetail) Validate() error {
	return dara.Validate(s)
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail struct {
	TaskCheckDO []*GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO `json:"TaskCheckDO,omitempty" xml:"TaskCheckDO,omitempty" type:"Repeated"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail) GetTaskCheckDO() []*GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO {
	return s.TaskCheckDO
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail) SetTaskCheckDO(v []*GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail {
	s.TaskCheckDO = v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetail) Validate() error {
	return dara.Validate(s)
}

type GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO struct {
	// The state of the precheck. Valid values:
	//
	// 	- **WAITING**: The ticket is pending precheck.
	//
	// 	- **RUNNING**: The ticket is being prechecked.
	//
	// 	- **SUCCESS**: The ticket passes the precheck.
	//
	// 	- **FAIL**: The ticket fails the precheck.
	//
	// example:
	//
	// SUCCESS
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The check step of the precheck. Valid values:
	//
	// 	- **SQL_PARSE**: The system checks the syntax of the SQL statement.
	//
	// 	- **SQL_TYPE_CHECK**: The system checks the type of the SQL statement.
	//
	// 	- **PERMISSION_CHECK**: The system checks the permissions required for the data change.
	//
	// 	- **ROW_CHECK**: The system checks the number of affected rows.
	//
	// example:
	//
	// PERMISSION_CHECK
	CheckStep *string `json:"CheckStep,omitempty" xml:"CheckStep,omitempty"`
	// The message that appears when a check step is executed.
	//
	// example:
	//
	// tip messsage
	UserTip *string `json:"UserTip,omitempty" xml:"UserTip,omitempty"`
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) GoString() string {
	return s.String()
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) GetCheckStatus() *string {
	return s.CheckStatus
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) GetCheckStep() *string {
	return s.CheckStep
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) GetUserTip() *string {
	return s.UserTip
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) SetCheckStatus(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO {
	s.CheckStatus = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) SetCheckStep(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO {
	s.CheckStep = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) SetUserTip(v string) *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO {
	s.UserTip = &v
	return s
}

func (s *GetDataCorrectOrderDetailResponseBodyDataCorrectOrderDetailPreCheckDetailTaskCheckDO) Validate() error {
	return dara.Validate(s)
}
