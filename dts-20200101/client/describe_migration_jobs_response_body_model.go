// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeMigrationJobsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeMigrationJobsResponseBody
	GetErrMessage() *string
	SetMigrationJobs(v *DescribeMigrationJobsResponseBodyMigrationJobs) *DescribeMigrationJobsResponseBody
	GetMigrationJobs() *DescribeMigrationJobsResponseBodyMigrationJobs
	SetPageNumber(v int32) *DescribeMigrationJobsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeMigrationJobsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeMigrationJobsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeMigrationJobsResponseBody
	GetSuccess() *string
	SetTotalRecordCount(v int64) *DescribeMigrationJobsResponseBody
	GetTotalRecordCount() *int64
}

type DescribeMigrationJobsResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The list of data migration instances and the details of each instance.
	MigrationJobs *DescribeMigrationJobsResponseBodyMigrationJobs `json:"MigrationJobs,omitempty" xml:"MigrationJobs,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries that can be displayed on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0ED4846D-ED60-435D-88C0-7EC0CE4D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of data migration instances that belong to your Alibaba Cloud account.
	//
	// example:
	//
	// 300
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeMigrationJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeMigrationJobsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeMigrationJobsResponseBody) GetMigrationJobs() *DescribeMigrationJobsResponseBodyMigrationJobs {
	return s.MigrationJobs
}

func (s *DescribeMigrationJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMigrationJobsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeMigrationJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMigrationJobsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeMigrationJobsResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeMigrationJobsResponseBody) SetErrCode(v string) *DescribeMigrationJobsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetErrMessage(v string) *DescribeMigrationJobsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetMigrationJobs(v *DescribeMigrationJobsResponseBodyMigrationJobs) *DescribeMigrationJobsResponseBody {
	s.MigrationJobs = v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetPageNumber(v int32) *DescribeMigrationJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetPageRecordCount(v int32) *DescribeMigrationJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetRequestId(v string) *DescribeMigrationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetSuccess(v string) *DescribeMigrationJobsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetTotalRecordCount(v int64) *DescribeMigrationJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobs struct {
	MigrationJob []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob `json:"MigrationJob,omitempty" xml:"MigrationJob,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobs) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobs) GetMigrationJob() []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	return s.MigrationJob
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobs) SetMigrationJob(v []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) *DescribeMigrationJobsResponseBodyMigrationJobs {
	s.MigrationJob = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobs) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob struct {
	// The details of full data migration.
	DataInitialization *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty" type:"Struct"`
	// The details of incremental data migration.
	DataSynchronization *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty" type:"Struct"`
	// The connection settings of the destination instance.
	DestinationEndpoint *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The time when the data migration instance was created. The time is displayed in the *yyyy-MM-dd*T*HH:mm:ss*Z format in UTC.
	//
	// example:
	//
	// 2021-06-22T09:02:13Z
	InstanceCreateTime *string `json:"InstanceCreateTime,omitempty" xml:"InstanceCreateTime,omitempty"`
	// The time when the data migration task was created. The time is displayed in the *yyyy-MM-dd*T*HH:mm:ss*Z format in UTC.
	//
	// example:
	//
	// 2021-06-22T08:53:55Z
	JobCreateTime *string `json:"JobCreateTime,omitempty" xml:"JobCreateTime,omitempty"`
	// The specification of the data migration instance. Valid values: **small**, **medium**, **large**, **xlarge**, and **2xlarge**. For more information, see [Specifications of data migration instances](https://help.aliyun.com/document_detail/26606.html).
	//
	// example:
	//
	// 2xlarge
	MigrationJobClass *string `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	// The ID of the data migration instance.
	//
	// example:
	//
	// dtson2z28evm33****
	MigrationJobID *string `json:"MigrationJobID,omitempty" xml:"MigrationJobID,omitempty"`
	// The name of the data migration task.
	//
	// example:
	//
	// dtstest
	MigrationJobName *string `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	// The status of the data migration task. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is being prechecked.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **Migrating**: The task is migrating data.
	//
	// 	- **Suspending**: The task is paused.
	//
	// 	- **MigrationFailed**: The task failed to migrate data.
	//
	// 	- **Finished**: The task is completed.
	//
	// example:
	//
	// Migrating
	MigrationJobStatus *string `json:"MigrationJobStatus,omitempty" xml:"MigrationJobStatus,omitempty"`
	// The migration types.
	MigrationMode *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// The objects that are migrated by the task.
	MigrationObject *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty" type:"Struct"`
	// The billing method of the data migration instance. The value is **PostPaid*	- (pay-as-you-go).
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The precheck details.
	Precheck *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck `json:"Precheck,omitempty" xml:"Precheck,omitempty" type:"Struct"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The details of schema migration.
	StructureInitialization *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty" type:"Struct"`
	// The collection of tags.
	Tags *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetDataInitialization() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	return s.DataInitialization
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetDataSynchronization() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	return s.DataSynchronization
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetDestinationEndpoint() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetInstanceCreateTime() *string {
	return s.InstanceCreateTime
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetJobCreateTime() *string {
	return s.JobCreateTime
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetMigrationJobClass() *string {
	return s.MigrationJobClass
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetMigrationJobID() *string {
	return s.MigrationJobID
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetMigrationJobName() *string {
	return s.MigrationJobName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetMigrationJobStatus() *string {
	return s.MigrationJobStatus
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetMigrationMode() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	return s.MigrationMode
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetMigrationObject() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject {
	return s.MigrationObject
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetPayType() *string {
	return s.PayType
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetPrecheck() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck {
	return s.Precheck
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetSourceEndpoint() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetStructureInitialization() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	return s.StructureInitialization
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GetTags() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags {
	return s.Tags
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetDataInitialization(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.DataInitialization = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetDataSynchronization(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.DataSynchronization = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetDestinationEndpoint(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetInstanceCreateTime(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.InstanceCreateTime = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetJobCreateTime(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.JobCreateTime = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobClass(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobClass = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobStatus = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationMode(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationMode = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationObject(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationObject = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetPayType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.PayType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetPrecheck(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.Precheck = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetSourceEndpoint(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetStructureInitialization(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.StructureInitialization = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetTags(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.Tags = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization struct {
	// The error message returned if full data migration failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The migration progress. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of records that have been migrated during full data migration.
	//
	// example:
	//
	// 200001
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of full data migration. Valid values:
	//
	// 	- **NotStarted**: Full data migration is not started.
	//
	// 	- **Migrating**: Full data migration is in progress.
	//
	// 	- **Failed**: Full data migration failed.
	//
	// 	- **Finished**: Full data migration is completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GetPercent() *string {
	return s.Percent
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GetProgress() *string {
	return s.Progress
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetProgress(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Progress = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization struct {
	// The latency of incremental data migration. Unit: seconds.
	//
	// example:
	//
	// 0
	Delay *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The error message returned if incremental data migration failed.
	//
	// example:
	//
	// Open: open \\\\\\\\?\\\\F:\\\\KINGDEE BACK\\\\AIS20221025151008_Data.mdf: The process cannot access the file because it is being used by another process.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of incremental data migration. Unit: %.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The status of incremental data migration. Valid values:
	//
	// 	- **NotStarted**: Incremental data migration is not started.
	//
	// 	- **Migrating**: Incremental data migration is in progress.
	//
	// 	- **Failed**: Incremental data migration failed.
	//
	// 	- **Finished**: Incremental data migration is completed.
	//
	// 	- **Catched**: Incremental data migration is not delayed.
	//
	// example:
	//
	// Catched
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) GetDelay() *string {
	return s.Delay
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) GetPercent() *string {
	return s.Percent
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetDelay(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Delay = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint struct {
	// The name of the database to which the migration object in the destination instance belongs.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database type of the destination instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The endpoint of the destination instance.
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the destination instance.
	//
	// example:
	//
	// rm-bp1zc3iyqe3qw****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the destination instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is returned only if the database type of the destination instance is **Oracle**.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the destination instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The database account of the destination instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetIP() *string {
	return s.IP
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetDatabaseName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetEngineName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetIP(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetInstanceID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetInstanceType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetOracleSID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetPort(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetUserName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode struct {
	// Indicates whether full data migration is performed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Indicates whether incremental data migration is performed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Indicates whether schema migration is performed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) SetDataInitialization(v bool) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) SetDataSynchronization(v bool) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) SetStructureInitialization(v bool) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject struct {
	SynchronousObject []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject `json:"SynchronousObject,omitempty" xml:"SynchronousObject,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) GetSynchronousObject() []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	return s.SynchronousObject
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) SetSynchronousObject(v []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject {
	s.SynchronousObject = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject struct {
	// The name of the database to which the migration object in the source instance belongs.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The names of the migrated tables.
	TableList *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
	// Indicates whether an entire database is migrated. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	WholeDatabase *string `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) GetTableList() *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList {
	return s.TableList
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) GetWholeDatabase() *string {
	return s.WholeDatabase
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetDatabaseName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetTableList(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.TableList = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetWholeDatabase(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.WholeDatabase = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) GetTable() []*string {
	return s.Table
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) SetTable(v []*string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList {
	s.Table = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck struct {
	// The precheck progress. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck result. Valid values:
	//
	// 	- **Success**: The task passed the precheck.
	//
	// 	- **Failed**: The task failed to pass the precheck.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) GetPercent() *string {
	return s.Percent
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint struct {
	// The name of the database to which the migration object in the source instance belongs.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database type of the source instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The endpoint of the source instance.
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the source instance.
	//
	// example:
	//
	// rm-bp1i99e8l7913****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the source instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is returned only if the database type of the source instance is **Oracle**.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the source instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The database account of the source instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetIP() *string {
	return s.IP
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetDatabaseName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetEngineName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetIP(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetInstanceID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetInstanceType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetOracleSID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetPort(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetUserName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization struct {
	// The error message returned if schema migration failed.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: ERROR: type "geometry" does not exist;
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of schema migration. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables whose schemas have been migrated.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of schema migration. Valid values:
	//
	// 	- **NotStarted**: Schema migration is not started.
	//
	// 	- **Migrating**: Schema migration is in progress.
	//
	// 	- **Failed**: Schema migration failed.
	//
	// 	- **Finished**: Schema migration is completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GetPercent() *string {
	return s.Percent
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GetProgress() *string {
	return s.Progress
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetProgress(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Progress = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags struct {
	Tag []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) GetTag() []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag {
	return s.Tag
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) SetTag(v []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags {
	s.Tag = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that corresponds to the tag key.
	//
	// example:
	//
	// testvalue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) SetKey(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) SetValue(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) Validate() error {
	return dara.Validate(s)
}
