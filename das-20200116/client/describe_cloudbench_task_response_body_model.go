// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudbenchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCloudbenchTaskResponseBody
	GetCode() *string
	SetData(v *DescribeCloudbenchTaskResponseBodyData) *DescribeCloudbenchTaskResponseBody
	GetData() *DescribeCloudbenchTaskResponseBodyData
	SetMessage(v string) *DescribeCloudbenchTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCloudbenchTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeCloudbenchTaskResponseBody
	GetSuccess() *string
}

type DescribeCloudbenchTaskResponseBody struct {
	// The returned status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the total number of entries and error codes.
	Data *DescribeCloudbenchTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >If the request is successful, **Successful*	- is returned. If the request fails, an error message is returned, such as an error code.
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
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudbenchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudbenchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCloudbenchTaskResponseBody) GetData() *DescribeCloudbenchTaskResponseBodyData {
	return s.Data
}

func (s *DescribeCloudbenchTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCloudbenchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudbenchTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeCloudbenchTaskResponseBody) SetCode(v string) *DescribeCloudbenchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBody) SetData(v *DescribeCloudbenchTaskResponseBodyData) *DescribeCloudbenchTaskResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudbenchTaskResponseBody) SetMessage(v string) *DescribeCloudbenchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBody) SetRequestId(v string) *DescribeCloudbenchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBody) SetSuccess(v string) *DescribeCloudbenchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudbenchTaskResponseBodyData struct {
	// The archiving task ID.
	//
	// example:
	//
	// \\"202105211430070112231480820340758****
	ArchiveJobId *string `json:"ArchiveJobId,omitempty" xml:"ArchiveJobId,omitempty"`
	// The name of the archived OSS table.
	//
	// example:
	//
	// custins15546355_161604665****
	ArchiveOssTableName *string `json:"ArchiveOssTableName,omitempty" xml:"ArchiveOssTableName,omitempty"`
	// The SQL archiving state. Valid values:
	//
	// - **0**: not started.
	//
	// - **1**: completed.
	//
	// - **2**: error.
	//
	// - **3**: running.
	//
	// - **4**: no download required.
	//
	// example:
	//
	// 1
	ArchiveState *int32 `json:"ArchiveState,omitempty" xml:"ArchiveState,omitempty"`
	// The backup set ID. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/26273.html) operation to obtain the backup set ID.
	//
	// example:
	//
	// 229132
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup type. Valid values:
	//
	// - **TIMESTAMP**: by backup time.
	//
	// - **BACKUPID**: by backup set ID.
	//
	// example:
	//
	// TIMESTAMP
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The substep of the stress testing task. Valid values:
	//
	// - **NEW**: task initialization.
	//
	// - **WAIT_BUY_ECS**: purchasing an ECS instance.
	//
	// - **WAIT_START_ECS**: starting the ECS instance.
	//
	// - **WAIT_INSTALL_JDK**: installing JDK.
	//
	// - **WAIT_INSTALL_DBGATEWAY**: installing DBGateway.
	//
	// - **ADD_SECURITY_IPS_STEP**: configuring the security group whitelist.
	//
	// - **ARCHIVE**: archiving full SQL.
	//
	// - **DOWNLOAD**: downloading the full SQL file.
	//
	// - **PROCEED**: preprocessing the full SQL file.
	//
	// - **PRE_LOAD**: preloading the full SQL file.
	//
	// - **VALIDATE**: functional verification.
	//
	// - **PRESSURE**: performance stress testing.
	//
	// example:
	//
	// PROCEED
	BenchStep *string `json:"BenchStep,omitempty" xml:"BenchStep,omitempty"`
	// The status of the stress testing substep. Valid values:
	//
	// - **NEW**: task initialization.
	//
	// - **RUNNING**: running.
	//
	// - **FAILED**: failed.
	//
	// - **FINISHED**: completed.
	//
	// - **Terminated**: terminated.
	//
	// - **Deleted**: deleted.
	//
	// example:
	//
	// FINISHED
	BenchStepStatus *string `json:"BenchStepStatus,omitempty" xml:"BenchStepStatus,omitempty"`
	// The database gateway ID of the stress testing machine.
	//
	// example:
	//
	// 58598b2af48a0193dfc16fc6964ef****
	ClientGatewayId *string `json:"ClientGatewayId,omitempty" xml:"ClientGatewayId,omitempty"`
	// The type of the stress testing machine. Valid values:
	//
	// - **ECS**: You need to prepare the [Database Gateway](https://help.aliyun.com/document_detail/64905.html) on your own.
	//
	// - **DAS_ECS**: an ECS instance automatically purchased and deployed by DAS.
	//
	// example:
	//
	// ECS
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The description of the stress testing task.
	//
	// example:
	//
	// test-das-bench-0501
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identity of the target instance.
	//
	// example:
	//
	// hdm_d887b5ccf99fa0dc9a1e5aaac368****
	DstInstanceUuid *string `json:"DstInstanceUuid,omitempty" xml:"DstInstanceUuid,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	DstIp *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	// The port of the target instance.
	//
	// example:
	//
	// 3306
	DstPort *int32 `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The type of the target instance. Valid values:
	//
	// - **Instance*	- (default): instance ID.
	//
	// - **ConnectionString**: endpoint of the instance.
	//
	// example:
	//
	// Instance
	DstType *string `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The DTS task specification.
	//
	// example:
	//
	// medium
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The ID of the DTS migration task.
	//
	// example:
	//
	// i03e3zty16i****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The name of the Data Transmission Service (DTS) task.
	//
	// example:
	//
	// Migration between RDS instances
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The DTS task state. Valid values:
	//
	// - **NOT_STARTED**: not started.
	//
	// - **PRE_CHECKING**: precheck in progress.
	//
	// - **PRE_CHECK_FAILED**: precheck failed.
	//
	// - **CHECKING**: checking.
	//
	// - **MIGRATING**: migrating.
	//
	// - **CATCHED**: caught up.
	//
	// - **SUSPENDING**: suspending.
	//
	// - **MIGRATION_FAILED**: migration failed.
	//
	// - **FINISHED**: completed.
	//
	// - **INITIALIZING**: synchronization initializing.
	//
	// - **INITIALIZE_FAILED**: synchronization initialization failed.
	//
	// - **SYNCHRONIZING**: synchronizing.
	//
	// - **MODIFYING**: modifying synchronization objects.
	//
	// - **SWITCHING**: switching.
	//
	// - **FAILED**: failed.
	//
	// example:
	//
	// CHECKING
	DtsJobState *int32 `json:"DtsJobState,omitempty" xml:"DtsJobState,omitempty"`
	// The DTS task state. Valid values:
	//
	// - **NOT_STARTED**: not started.
	//
	// - **PRE_CHECKING**: precheck in progress.
	//
	// - **PRE_CHECK_FAILED**: precheck failed.
	//
	// - **CHECKING**: checking.
	//
	// - **MIGRATING**: migrating.
	//
	// - **CATCHED**: caught up.
	//
	// - **SUSPENDING**: suspending.
	//
	// - **MIGRATION_FAILED**: migration failed.
	//
	// - **FINISHED**: completed.
	//
	// - **INITIALIZING**: synchronization initializing.
	//
	// - **INITIALIZE_FAILED**: synchronization initialization failed.
	//
	// - **SYNCHRONIZING**: synchronizing.
	//
	// - **MODIFYING**: modifying synchronization objects.
	//
	// - **SWITCHING**: switching.
	//
	// - **FAILED**: failed.
	//
	// example:
	//
	// PRE_CHECKING
	DtsJobStatus *string `json:"DtsJobStatus,omitempty" xml:"DtsJobStatus,omitempty"`
	// The ECS instance ID.
	//
	// example:
	//
	// i-bp1ecr5go2go1****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The state after the stress testing task ends. Valid values:
	//
	// - **WAIT_TARGET**: preparing the target instance for stress testing.
	//
	// - **WAIT_DBGATEWAY**: preparing the stress testing deployment.
	//
	// - **WAIT_SQL**: preparing full SQL.
	//
	// - **WAIT_LOGIC**: preparing to start traffic replay.
	//
	// >After the stress testing task executes the state specified by EndState, the task directly reaches the completed state.
	//
	// example:
	//
	// WAIT_LOGIC
	EndState *string `json:"EndState,omitempty" xml:"EndState,omitempty"`
	// The error code returned by the internal stress testing task.
	//
	// example:
	//
	// 10910
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the task fails.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by [com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException:Could not create connection to database server. Attempted reconnect 3 times. Giving up.][com.mysql.jdbc.exceptions.jdbc4.CommunicationsException:Communications link failure\\n\\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.][java.net.ConnectException:Connection timed out (Connection timed out)] About more information in [https://yq.aliyun.com/articles/499178].
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The additional information.
	//
	// example:
	//
	// Null
	External *string `json:"External,omitempty" xml:"External,omitempty"`
	// The stress testing rate multiplier. The replay rate must be a positive integer. Valid values: **0*	- to **30**. Default value: **1**.
	//
	// example:
	//
	// 1
	Rate *int64 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The duration of the stress testing task.
	//
	// example:
	//
	// 864000
	RequestDuration *int64 `json:"RequestDuration,omitempty" xml:"RequestDuration,omitempty"`
	// The generated stress testing duration. Unit: milliseconds.
	//
	// example:
	//
	// 86400000
	SmartPressureTime *int32 `json:"SmartPressureTime,omitempty" xml:"SmartPressureTime,omitempty"`
	// The task source. Valid values:
	//
	// - **DAS**.
	//
	// - **OPEN_API**.
	//
	// example:
	//
	// DAS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The reuse information of the full SQL.
	//
	// example:
	//
	// {"sqlUuid":"task_a37d2f07-45cb-****-a2a6-c66c62****","metaUuid":"task_211e2561-5c0c-486b-864c-56b511****","sqlFile":"cl-1620057600000-1800626.sc","metaFile":"cl-1620057600000-180****.meta"}
	SqlCompleteReuse *string `json:"SqlCompleteReuse,omitempty" xml:"SqlCompleteReuse,omitempty"`
	// The database type of the source instance.
	//
	// example:
	//
	// RDS
	SrcInstanceArea *string `json:"SrcInstanceArea,omitempty" xml:"SrcInstanceArea,omitempty"`
	// The UUID of the source instance.
	//
	// example:
	//
	// a364e414-e68b-4e5c-9166-65b3a153****
	SrcInstanceUuid *string `json:"SrcInstanceUuid,omitempty" xml:"SrcInstanceUuid,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	SrcPublicIp *string `json:"SrcPublicIp,omitempty" xml:"SrcPublicIp,omitempty"`
	// The current state of the stress testing task. Valid values:
	//
	// - **WAIT_TARGET**: preparing the target instance for stress testing.
	//
	// - **WAIT_DBGATEWAY**: preparing the stress testing deployment.
	//
	// - **WAIT_SQL**: preparing full SQL.
	//
	// - **WAIT_LOGIC**: preparing to start traffic replay.
	//
	// example:
	//
	// WAIT_TARGET
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The running status of the stress testing task. Valid values:
	//
	// - **SUCCESS**: successful.
	//
	// - **IGNORED**: ignored.
	//
	// - **RUNNING**: running.
	//
	// - **EXCEPTION**: abnormal.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The table names involved in the stress testing task.
	//
	// example:
	//
	// [{"TABLE_NAME":"customer1","TABLE_SCHEMA":"tpcc"}]
	TableSchema *string `json:"TableSchema,omitempty" xml:"TableSchema,omitempty"`
	// The task ID.
	//
	// example:
	//
	// e5cec704-0518-430f-8263-76f4dcds****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the stress testing task. Valid values:
	//
	// - **pressure test*	- (default): intelligent stress testing. The traffic captured from the target instance is replayed on the destination instance at the maximum speed supported by the destination instance specifications.
	//
	// - **smart pressure test**: generated stress testing. By analyzing and learning the traffic captured from the target instance within a short period of time, traffic that is consistent with the business model and traffic distribution of the original traffic is generated for continuous stress testing. This reduces the time required to collect data from the target instance and lowers storage costs and performance overhead.
	//
	// example:
	//
	// pressure test
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The Kafka consumption topic.
	//
	// example:
	//
	// das
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 109141182625****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The version of the stress testing task. Valid values:
	//
	// - **V2.0**
	//
	// - **V3.0**
	//
	// example:
	//
	// V3.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The temporary directory generated by the stress testing task.
	//
	// example:
	//
	// /tmp/bench/
	WorkDir *string `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
}

func (s DescribeCloudbenchTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudbenchTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetArchiveJobId() *string {
	return s.ArchiveJobId
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetArchiveOssTableName() *string {
	return s.ArchiveOssTableName
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetArchiveState() *int32 {
	return s.ArchiveState
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetBenchStep() *string {
	return s.BenchStep
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetBenchStepStatus() *string {
	return s.BenchStepStatus
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetClientGatewayId() *string {
	return s.ClientGatewayId
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetDstInstanceUuid() *string {
	return s.DstInstanceUuid
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetDstIp() *string {
	return s.DstIp
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetDstPort() *int32 {
	return s.DstPort
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetDstType() *string {
	return s.DstType
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetDtsJobState() *int32 {
	return s.DtsJobState
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetDtsJobStatus() *string {
	return s.DtsJobStatus
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetEndState() *string {
	return s.EndState
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetExternal() *string {
	return s.External
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetRate() *int64 {
	return s.Rate
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetRequestDuration() *int64 {
	return s.RequestDuration
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetSmartPressureTime() *int32 {
	return s.SmartPressureTime
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetSqlCompleteReuse() *string {
	return s.SqlCompleteReuse
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetSrcInstanceArea() *string {
	return s.SrcInstanceArea
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetSrcInstanceUuid() *string {
	return s.SrcInstanceUuid
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetSrcPublicIp() *string {
	return s.SrcPublicIp
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetState() *string {
	return s.State
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetTableSchema() *string {
	return s.TableSchema
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetTopic() *string {
	return s.Topic
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *DescribeCloudbenchTaskResponseBodyData) GetWorkDir() *string {
	return s.WorkDir
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetArchiveJobId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.ArchiveJobId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetArchiveOssTableName(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.ArchiveOssTableName = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetArchiveState(v int32) *DescribeCloudbenchTaskResponseBodyData {
	s.ArchiveState = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetBackupId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.BackupId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetBackupType(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetBenchStep(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.BenchStep = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetBenchStepStatus(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.BenchStepStatus = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetClientGatewayId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.ClientGatewayId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetClientType(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDescription(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDstInstanceUuid(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DstInstanceUuid = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDstIp(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DstIp = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDstPort(v int32) *DescribeCloudbenchTaskResponseBodyData {
	s.DstPort = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDstType(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DstType = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDtsJobClass(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDtsJobId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DtsJobId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDtsJobName(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DtsJobName = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDtsJobState(v int32) *DescribeCloudbenchTaskResponseBodyData {
	s.DtsJobState = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetDtsJobStatus(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.DtsJobStatus = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetEcsInstanceId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetEndState(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.EndState = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetErrorCode(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetErrorMessage(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetExternal(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.External = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetRate(v int64) *DescribeCloudbenchTaskResponseBodyData {
	s.Rate = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetRequestDuration(v int64) *DescribeCloudbenchTaskResponseBodyData {
	s.RequestDuration = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetSmartPressureTime(v int32) *DescribeCloudbenchTaskResponseBodyData {
	s.SmartPressureTime = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetSource(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.Source = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetSqlCompleteReuse(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.SqlCompleteReuse = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetSrcInstanceArea(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.SrcInstanceArea = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetSrcInstanceUuid(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.SrcInstanceUuid = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetSrcPublicIp(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.SrcPublicIp = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetState(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.State = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetStatus(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetTableSchema(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.TableSchema = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetTaskId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetTaskType(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetTopic(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.Topic = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetUserId(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.UserId = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetVersion(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) SetWorkDir(v string) *DescribeCloudbenchTaskResponseBodyData {
	s.WorkDir = &v
	return s
}

func (s *DescribeCloudbenchTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
