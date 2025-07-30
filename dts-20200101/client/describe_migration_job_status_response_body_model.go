// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrationJobStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInitializationStatus(v *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) *DescribeMigrationJobStatusResponseBody
	GetDataInitializationStatus() *DescribeMigrationJobStatusResponseBodyDataInitializationStatus
	SetDataSynchronizationStatus(v *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) *DescribeMigrationJobStatusResponseBody
	GetDataSynchronizationStatus() *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus
	SetDestinationEndpoint(v *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) *DescribeMigrationJobStatusResponseBody
	GetDestinationEndpoint() *DescribeMigrationJobStatusResponseBodyDestinationEndpoint
	SetErrCode(v string) *DescribeMigrationJobStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeMigrationJobStatusResponseBody
	GetErrMessage() *string
	SetMigrationJobClass(v string) *DescribeMigrationJobStatusResponseBody
	GetMigrationJobClass() *string
	SetMigrationJobId(v string) *DescribeMigrationJobStatusResponseBody
	GetMigrationJobId() *string
	SetMigrationJobName(v string) *DescribeMigrationJobStatusResponseBody
	GetMigrationJobName() *string
	SetMigrationJobStatus(v string) *DescribeMigrationJobStatusResponseBody
	GetMigrationJobStatus() *string
	SetMigrationMode(v *DescribeMigrationJobStatusResponseBodyMigrationMode) *DescribeMigrationJobStatusResponseBody
	GetMigrationMode() *DescribeMigrationJobStatusResponseBodyMigrationMode
	SetMigrationObject(v string) *DescribeMigrationJobStatusResponseBody
	GetMigrationObject() *string
	SetPayType(v string) *DescribeMigrationJobStatusResponseBody
	GetPayType() *string
	SetPrecheckStatus(v *DescribeMigrationJobStatusResponseBodyPrecheckStatus) *DescribeMigrationJobStatusResponseBody
	GetPrecheckStatus() *DescribeMigrationJobStatusResponseBodyPrecheckStatus
	SetRequestId(v string) *DescribeMigrationJobStatusResponseBody
	GetRequestId() *string
	SetSourceEndpoint(v *DescribeMigrationJobStatusResponseBodySourceEndpoint) *DescribeMigrationJobStatusResponseBody
	GetSourceEndpoint() *DescribeMigrationJobStatusResponseBodySourceEndpoint
	SetStructureInitializationStatus(v *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) *DescribeMigrationJobStatusResponseBody
	GetStructureInitializationStatus() *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus
	SetSuccess(v string) *DescribeMigrationJobStatusResponseBody
	GetSuccess() *string
	SetTaskId(v string) *DescribeMigrationJobStatusResponseBody
	GetTaskId() *string
}

type DescribeMigrationJobStatusResponseBody struct {
	// The status of full data migration.
	DataInitializationStatus *DescribeMigrationJobStatusResponseBodyDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The status of incremental data migration.
	DataSynchronizationStatus *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The connection settings of the destination instance.
	DestinationEndpoint *DescribeMigrationJobStatusResponseBodyDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
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
	// dtsz2v12jfo309****
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	// The name of the data migration task.
	//
	// example:
	//
	// MySQL migration
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
	MigrationMode *DescribeMigrationJobStatusResponseBodyMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// The objects that are migrated by the task.
	//
	// example:
	//
	// [{\\"DBName\\":\\"dtstestdata\\",\\"TableIncludes\\":[{\\"TableName\\":\\"customer\\"}]}]
	MigrationObject *string `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty"`
	// The billing method of the data migration instance. The value is **PostPaid*	- (pay-as-you-go).
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The precheck details.
	PrecheckStatus *DescribeMigrationJobStatusResponseBodyPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A032E3B4-929B-48E9-97B9-37587CBF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeMigrationJobStatusResponseBodySourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The status of schema migration.
	StructureInitializationStatus *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// z2v12jfo309****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBody) GetDataInitializationStatus() *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeMigrationJobStatusResponseBody) GetDataSynchronizationStatus() *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeMigrationJobStatusResponseBody) GetDestinationEndpoint() *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeMigrationJobStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeMigrationJobStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeMigrationJobStatusResponseBody) GetMigrationJobClass() *string {
	return s.MigrationJobClass
}

func (s *DescribeMigrationJobStatusResponseBody) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *DescribeMigrationJobStatusResponseBody) GetMigrationJobName() *string {
	return s.MigrationJobName
}

func (s *DescribeMigrationJobStatusResponseBody) GetMigrationJobStatus() *string {
	return s.MigrationJobStatus
}

func (s *DescribeMigrationJobStatusResponseBody) GetMigrationMode() *DescribeMigrationJobStatusResponseBodyMigrationMode {
	return s.MigrationMode
}

func (s *DescribeMigrationJobStatusResponseBody) GetMigrationObject() *string {
	return s.MigrationObject
}

func (s *DescribeMigrationJobStatusResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeMigrationJobStatusResponseBody) GetPrecheckStatus() *DescribeMigrationJobStatusResponseBodyPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeMigrationJobStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMigrationJobStatusResponseBody) GetSourceEndpoint() *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeMigrationJobStatusResponseBody) GetStructureInitializationStatus() *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeMigrationJobStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeMigrationJobStatusResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeMigrationJobStatusResponseBody) SetDataInitializationStatus(v *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) *DescribeMigrationJobStatusResponseBody {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetDataSynchronizationStatus(v *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) *DescribeMigrationJobStatusResponseBody {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetDestinationEndpoint(v *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) *DescribeMigrationJobStatusResponseBody {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetErrCode(v string) *DescribeMigrationJobStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetErrMessage(v string) *DescribeMigrationJobStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobClass(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobClass = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobId(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobName(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobStatus(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobStatus = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationMode(v *DescribeMigrationJobStatusResponseBodyMigrationMode) *DescribeMigrationJobStatusResponseBody {
	s.MigrationMode = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationObject(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationObject = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetPayType(v string) *DescribeMigrationJobStatusResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetPrecheckStatus(v *DescribeMigrationJobStatusResponseBodyPrecheckStatus) *DescribeMigrationJobStatusResponseBody {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetRequestId(v string) *DescribeMigrationJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetSourceEndpoint(v *DescribeMigrationJobStatusResponseBodySourceEndpoint) *DescribeMigrationJobStatusResponseBody {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetStructureInitializationStatus(v *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) *DescribeMigrationJobStatusResponseBody {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetSuccess(v string) *DescribeMigrationJobStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetTaskId(v string) *DescribeMigrationJobStatusResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobStatusResponseBodyDataInitializationStatus struct {
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
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) SetErrorMessage(v string) *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) SetPercent(v string) *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) SetProgress(v string) *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) SetStatus(v string) *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus struct {
	// The UNIX timestamp generated when the latest incremental data is migrated. Unit: seconds.
	//
	// example:
	//
	// 1612507847
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
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
	// The task has failed for too long and cannot be restored
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
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) GetDelay() *string {
	return s.Delay
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetCheckpoint(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.Checkpoint = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetDelay(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.Delay = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetErrorMessage(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetPercent(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetStatus(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobStatusResponseBodyDestinationEndpoint struct {
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the destination instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
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
	// The system ID (SID) of the Oracle database.
	//
	// >  This parameter is returned only if the database type of the destination instance is **Oracle**.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"oracleSID,omitempty" xml:"oracleSID,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) GetIP() *string {
	return s.IP
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetDatabaseName(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetEngineName(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetIP(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetInstanceId(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetInstanceType(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetPort(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetUserName(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetOracleSID(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobStatusResponseBodyMigrationMode struct {
	// Indicates whether full data migration is performed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"dataInitialization,omitempty" xml:"dataInitialization,omitempty"`
	// Indicates whether incremental data migration is performed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"dataSynchronization,omitempty" xml:"dataSynchronization,omitempty"`
	// Indicates whether schema migration is performed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"structureInitialization,omitempty" xml:"structureInitialization,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeMigrationJobStatusResponseBodyMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeMigrationJobStatusResponseBodyMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeMigrationJobStatusResponseBodyMigrationMode) SetDataInitialization(v bool) *DescribeMigrationJobStatusResponseBodyMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyMigrationMode) SetDataSynchronization(v bool) *DescribeMigrationJobStatusResponseBodyMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyMigrationMode) SetStructureInitialization(v bool) *DescribeMigrationJobStatusResponseBodyMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobStatusResponseBodyPrecheckStatus struct {
	// The result of each precheck item.
	Detail *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	// The precheck progress. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck status. Valid values:
	//
	// 	- **NotStarted**
	//
	// 	- **Suspending**:
	//
	// 	- **Checking**
	//
	// 	- **Failed**
	//
	// 	- **Finished**
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) GetDetail() *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) SetDetail(v *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail) *DescribeMigrationJobStatusResponseBodyPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) SetPercent(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) SetStatus(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail struct {
	CheckItem []*DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem `json:"CheckItem,omitempty" xml:"CheckItem,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail) GetCheckItem() []*DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem {
	return s.CheckItem
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail) SetCheckItem(v []*DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail {
	s.CheckItem = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem struct {
	// The precheck result. Valid values:
	//
	// 	- **Success**: The task passed the precheck.
	//
	// 	- **Failed**: The task failed to pass the precheck.
	//
	// example:
	//
	// Success
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The error message returned if the task failed to pass the precheck.
	//
	// >  This parameter is returned only if the return value of the **CheckStatus*	- parameter is **Failed**.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the precheck item.
	//
	// example:
	//
	// CHECK_CONN_SRC
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The method to fix the precheck failure.
	//
	// >  This parameter is returned only if the return value of the **CheckStatus*	- parameter is **Failed**.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) GetCheckStatus() *string {
	return s.CheckStatus
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) GetItemName() *string {
	return s.ItemName
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) SetCheckStatus(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem {
	s.CheckStatus = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) SetErrorMessage(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) SetItemName(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem {
	s.ItemName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) SetRepairMethod(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem {
	s.RepairMethod = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobStatusResponseBodySourceEndpoint struct {
	// The name of the database to which the migration object in the source instance belongs.
	//
	// example:
	//
	// dtstestdatabase
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the source instance.
	//
	// 	- **RDS**: ApsaraDB RDS instance
	//
	// 	- **ECS**: self-managed database that is hosted on Elastic Compute Service (ECS)
	//
	// 	- **LocalInstance**: self-managed database with a public IP address
	//
	// 	- **Express**: self-managed database that is connected over Express Connect, VPN Gateway, or Smart Access Gateway
	//
	// 	- **MongoDB**: ApsaraDB for MongoDB instance
	//
	// 	- **POLARDB**: PolarDB for MySQL cluster (available only for the China site)
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
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
	// The SID of the Oracle database.
	//
	// >  This parameter is returned only if the database type of the source instance is **Oracle**.
	//
	// example:
	//
	// dtstestdatabase
	OracleSID *string `json:"oracleSID,omitempty" xml:"oracleSID,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodySourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodySourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) GetIP() *string {
	return s.IP
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetDatabaseName(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetEngineName(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetIP(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetInstanceId(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetInstanceType(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetPort(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetUserName(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetOracleSID(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeMigrationJobStatusResponseBodyStructureInitializationStatus struct {
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
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) SetErrorMessage(v string) *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) SetPercent(v string) *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) SetProgress(v string) *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) SetStatus(v string) *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}
