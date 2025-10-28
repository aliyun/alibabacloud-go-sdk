// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudBenchTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCloudBenchTasksResponseBody
	GetCode() *string
	SetData(v *DescribeCloudBenchTasksResponseBodyData) *DescribeCloudBenchTasksResponseBody
	GetData() *DescribeCloudBenchTasksResponseBodyData
	SetMessage(v string) *DescribeCloudBenchTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCloudBenchTasksResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeCloudBenchTasksResponseBody
	GetSuccess() *string
}

type DescribeCloudBenchTasksResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	Data *DescribeCloudBenchTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
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
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudBenchTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudBenchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCloudBenchTasksResponseBody) GetData() *DescribeCloudBenchTasksResponseBodyData {
	return s.Data
}

func (s *DescribeCloudBenchTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCloudBenchTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudBenchTasksResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeCloudBenchTasksResponseBody) SetCode(v string) *DescribeCloudBenchTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) SetData(v *DescribeCloudBenchTasksResponseBodyData) *DescribeCloudBenchTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) SetMessage(v string) *DescribeCloudBenchTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) SetRequestId(v string) *DescribeCloudBenchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) SetSuccess(v string) *DescribeCloudBenchTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudBenchTasksResponseBodyData struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The detailed information of the stress testing task.
	List *DescribeCloudBenchTasksResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCloudBenchTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudBenchTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *DescribeCloudBenchTasksResponseBodyData) GetList() *DescribeCloudBenchTasksResponseBodyDataList {
	return s.List
}

func (s *DescribeCloudBenchTasksResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeCloudBenchTasksResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudBenchTasksResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetExtra(v string) *DescribeCloudBenchTasksResponseBodyData {
	s.Extra = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetList(v *DescribeCloudBenchTasksResponseBodyDataList) *DescribeCloudBenchTasksResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetPageNo(v int32) *DescribeCloudBenchTasksResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetPageSize(v int32) *DescribeCloudBenchTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) SetTotal(v int64) *DescribeCloudBenchTasksResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyData) Validate() error {
	if s.List != nil {
		if err := s.List.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudBenchTasksResponseBodyDataList struct {
	CloudbenchTasks []*DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks `json:"cloudbenchTasks,omitempty" xml:"cloudbenchTasks,omitempty" type:"Repeated"`
}

func (s DescribeCloudBenchTasksResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudBenchTasksResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponseBodyDataList) GetCloudbenchTasks() []*DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	return s.CloudbenchTasks
}

func (s *DescribeCloudBenchTasksResponseBodyDataList) SetCloudbenchTasks(v []*DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) *DescribeCloudBenchTasksResponseBodyDataList {
	s.CloudbenchTasks = v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataList) Validate() error {
	if s.CloudbenchTasks != nil {
		for _, item := range s.CloudbenchTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks struct {
	// The archiving task ID.
	//
	// example:
	//
	// \\"202105211430070112231480820340758****
	ArchiveJobId *string `json:"ArchiveJobId,omitempty" xml:"ArchiveJobId,omitempty"`
	// The name of the table that was archived to Object Storage Service (OSS).
	//
	// example:
	//
	// custins15546355_161604665****
	ArchiveOssTableName *string `json:"ArchiveOssTableName,omitempty" xml:"ArchiveOssTableName,omitempty"`
	// The archiving status of the file that stores the analysis result of full SQL statistics. Valid values:
	//
	// 	- **0**: The file archiving is not started.
	//
	// 	- **1**: The file is archived.
	//
	// 	- **2**: An error occurred.
	//
	// 	- **3**: The file is being archived.
	//
	// 	- **4**: The archived file does not need to be downloaded.
	//
	// example:
	//
	// 1
	ArchiveState *int32 `json:"ArchiveState,omitempty" xml:"ArchiveState,omitempty"`
	// The ID of the backup set. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/26273.html) operation to query the ID of the backup set.
	//
	// example:
	//
	// 229132
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup type. Valid values:
	//
	// 	- **TIMESTAMP**: Data is restored to the state at a specific point in time.
	//
	// 	- **BACKUPID**: Data is restored from a backup set that is identified by an ID.
	//
	// example:
	//
	// TIMESTAMP
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The substep in the stress testing task. Valid values:
	//
	// 	- **NEW**: Initialize the stress testing task.
	//
	// 	- **WAIT_BUY_ECS**: Purchase an ECS instance.
	//
	// 	- **WAIT_START_ECS**: Start the ECS instance.
	//
	// 	- **WAIT_INSTALL_JDK**: Install the Java Development Kit (JDK).
	//
	// 	- **WAIT_INSTALL_DBGATEWAY**: Install the database gateway (DBGateway).
	//
	// 	- **ADD_SECURITY_IPS_STEP**: Configure the whitelist of the security group.
	//
	// 	- **ARCHIVE**: Archive the file that stores the analysis results of full SQL statistics.
	//
	// 	- **DOWNLOAD**: Download the file that stores the analysis result of full SQL statistics.
	//
	// 	- **PROCEED**: Preprocess the file that stores the analysis result of full SQL statistics.
	//
	// 	- **PRE_LOAD**: Preload the file that stores the analysis result of full SQL statistics.
	//
	// 	- **VALIDATE**: Verify the functionality of stress testing.
	//
	// 	- **PRESSURE**: Start the stress testing task.
	//
	// example:
	//
	// PROCEED
	BenchStep *string `json:"BenchStep,omitempty" xml:"BenchStep,omitempty"`
	// The status that indicates the substep performed for the stress testing task. Valid values:
	//
	// 	- **NEW**: The task is being initialized.
	//
	// 	- **RUNNING**: The task is running.
	//
	// 	- **FAILED**: The task failed.
	//
	// 	- **FINISHED**: The task is complete.
	//
	// 	- **Terminated**: The task is terminated.
	//
	// 	- **Deleted**: The task is deleted.
	//
	// example:
	//
	// FINISHED
	BenchStepStatus *string `json:"BenchStepStatus,omitempty" xml:"BenchStepStatus,omitempty"`
	// The DBGateway ID of the stress testing client.
	//
	// example:
	//
	// 58598b2af48a0193dfc16fc6964ef****
	ClientGatewayId *string `json:"ClientGatewayId,omitempty" xml:"ClientGatewayId,omitempty"`
	// The type of the stress testing client. Valid values:
	//
	// 	- **ECS**: indicates that you must prepare the DBGateway.
	//
	// 	- **DAS_ECS**: indicates that DAS automatically purchases and deploys an ECS instance for stress testing.
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
	// The UUID of the destination instance.
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
	// The port number of the destination instance.
	//
	// example:
	//
	// 3306
	DstPort *int32 `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The type of the identifier that is used to indicate the destination instance. Valid values:
	//
	// 	- **Instance*	- (default): the instance ID.
	//
	// 	- **ConnectionString**: the endpoint of the instance.
	//
	// example:
	//
	// Instance
	DstType *string `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The specification of the DTS instance.
	//
	// > For more information about the specifications of DTS instances and the test performance of each instance, see [Specifications of data migration instances](https://help.aliyun.com/document_detail/26606.html).
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
	// The name of the Data Transmission Service (DTS) migration task.
	//
	// example:
	//
	// RDS_TO_RDS_MIGRATION
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The status of the DTS migration task. Valid values:
	//
	// 	- **NOT_STARTED**: The task is not started.
	//
	// 	- **PRE_CHECKING**: The task is in precheck.
	//
	// 	- **PRE_CHECK_FAILED**: The precheck failed.
	//
	// 	- **CHECKING**: The task is being checked.
	//
	// 	- **MIGRATING**: The data is being migrated.
	//
	// 	- **CATCHED**: The data is migrated from the source instance to the destination instance.
	//
	// 	- **SUSPENDING**: The task is suspended.
	//
	// 	- **MIGRATION_FAILED**: The data failed to be migrated.
	//
	// 	- **FINISHED**: The task is complete.
	//
	// 	- **INITIALIZING**: The synchronization is being initialized.
	//
	// 	- **INITIALIZE_FAILED**: The synchronization failed to be initialized.
	//
	// 	- **SYNCHRONIZING**: The data is being synchronized.
	//
	// 	- **MODIFYING**: The roles of the instances are being changed.
	//
	// 	- **SWITCHING**: The roles of the instances are being switched.
	//
	// 	- **FAILED**: The task failed.
	//
	// example:
	//
	// CHECKING
	DtsJobState *int32 `json:"DtsJobState,omitempty" xml:"DtsJobState,omitempty"`
	// The status of the DTS migration task. Valid values:
	//
	// 	- **NOT_STARTED**: The task is not started.
	//
	// 	- **PRE_CHECKING**: The task is in precheck.
	//
	// 	- **PRE_CHECK_FAILED**: The precheck failed.
	//
	// 	- **CHECKING**: The task is being checked.
	//
	// 	- **MIGRATING**: The data is being migrated.
	//
	// 	- **CATCHED**: The data is migrated from the source instance to the destination instance.
	//
	// 	- **SUSPENDING**: The task is suspended.
	//
	// 	- **MIGRATION_FAILED**: The data failed to be migrated.
	//
	// 	- **FINISHED**: The task is complete.
	//
	// 	- **INITIALIZING**: The synchronization is being initialized.
	//
	// 	- **INITIALIZE_FAILED**: The synchronization failed to be initialized.
	//
	// 	- **SYNCHRONIZING**: The data is being synchronized.
	//
	// 	- **MODIFYING**: The roles of the instances are being changed.
	//
	// 	- **SWITCHING**: The roles of the instances are being switched.
	//
	// 	- **FAILED**: The task failed.
	//
	// example:
	//
	// PRE_CHECK_FAILED
	DtsJobStatus *string `json:"DtsJobStatus,omitempty" xml:"DtsJobStatus,omitempty"`
	// The ID of the Elastic Compute Service (ECS) instance.
	//
	// example:
	//
	// i-bp1ecr5go2go1****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The state that indicates the last operation performed for the stress testing task. Valid values:
	//
	// 	- **WAIT_TARGET**: prepares the destination instance.
	//
	// 	- **WAIT_DBGATEWAY**: prepares the DBGateway.
	//
	// 	- **WAIT_SQL**: prepares the full SQL statistics.
	//
	// 	- **WAIT_LOGIC**: prepares to replay the traffic.
	//
	// > When the state of a stress testing task changes to the state that is specified by the EndState parameter, the stress testing task is complete.
	//
	// example:
	//
	// WAIT_TARGET
	EndState *string `json:"EndState,omitempty" xml:"EndState,omitempty"`
	// The error code returned for the substep of the stress testing task.
	//
	// example:
	//
	// 10109
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the task failed.
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
	// The rate at which the stress testing task replayed the traffic. This value is a positive integer. Valid values: **0*	- to **30**. Default value: **1**.
	//
	// example:
	//
	// 1
	Rate *int32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The duration of the stress testing task. Unit: millisecond.
	//
	// example:
	//
	// 86400000
	RequestDuration *int64 `json:"RequestDuration,omitempty" xml:"RequestDuration,omitempty"`
	// The duration of the stress testing task of the smart pressure test type. Unit: millisecond.
	//
	// example:
	//
	// 86400000
	SmartPressureTime *int32 `json:"SmartPressureTime,omitempty" xml:"SmartPressureTime,omitempty"`
	// The source of the task. Valid values:
	//
	// 	- **DAS**
	//
	// 	- **OPEN_API**
	//
	// example:
	//
	// DAS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The reused information about the analysis result of full SQL statistics.
	//
	// example:
	//
	// {"sqlUuid":"task_a37d2f07-45cb-4413-a2a6-c66c68****","metaUuid":"task_211e2561-5c0c-486b-864c-56b511****","sqlFile":"cl-1620057600000-1800626.sc","metaFile":"cl-1620057600000-1800626.meta"}
	SqlCompleteReuse *string `json:"SqlCompleteReuse,omitempty" xml:"SqlCompleteReuse,omitempty"`
	// The database engine of the source instance. Valid values:
	//
	// example:
	//
	// RDS
	SrcInstanceArea *string `json:"SrcInstanceArea,omitempty" xml:"SrcInstanceArea,omitempty"`
	// The UUID of the source instance.
	//
	// example:
	//
	// hdm_3063db6792965c080a4bcb6e6304****
	SrcInstanceUuid *string `json:"SrcInstanceUuid,omitempty" xml:"SrcInstanceUuid,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	SrcPublicIp *string `json:"SrcPublicIp,omitempty" xml:"SrcPublicIp,omitempty"`
	// The state that indicates the operation performed for the stress testing task. Valid values:
	//
	// 	- **WAIT_TARGET**: prepares the destination instance.
	//
	// 	- **WAIT_DBGATEWAY**: prepares the DBGateway.
	//
	// 	- **WAIT_SQL**: prepares the full SQL statistics.
	//
	// 	- **WAIT_LOGIC**: prepares to replay the traffic.
	//
	// example:
	//
	// WAIT_TARGET
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the stress testing task. Valid values:
	//
	// 	- **SUCCESS**: The task was successful.
	//
	// 	- **IGNORED**: The task was ignored.
	//
	// 	- **RUNNING**: The task is running.
	//
	// 	- **EXCEPTION**: The task is abnormal.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the table that is used for stress testing.
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
	// 	- **pressure test*	- (default): A task of this type replays the traffic that is captured from the source instance on the destination instance at the maximum playback rate that is supported by the destination instance.
	//
	// 	- **smart pressure test**: A task of this type analyzes the traffic that is captured from the source instance over a short period of time and generates traffic on the destination instance for continuous stress testing. The business model based on which the traffic is generated on the destination instance and the traffic distribution are consistent with those on the source instance. Stress testing tasks of this type can help you reduce the amount of time that is consumed to collect data from the source instance and reduce storage costs and performance overheads.
	//
	// example:
	//
	// pressure test
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The topic that contains the consumed data. This topic is a topic in Message Queue for Apache Kafka.
	//
	// example:
	//
	// das
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 1091411816252****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The version of the stress testing task. Valid values:
	//
	// 	- **V2.0**
	//
	// 	- **V3.0**
	//
	// example:
	//
	// V3.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The path of the temporary directory that is generated for stress testing.
	//
	// example:
	//
	// /tmp/bench/
	WorkDir *string `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
}

func (s DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GoString() string {
	return s.String()
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetArchiveJobId() *string {
	return s.ArchiveJobId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetArchiveOssTableName() *string {
	return s.ArchiveOssTableName
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetArchiveState() *int32 {
	return s.ArchiveState
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetBenchStep() *string {
	return s.BenchStep
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetBenchStepStatus() *string {
	return s.BenchStepStatus
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetClientGatewayId() *string {
	return s.ClientGatewayId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDescription() *string {
	return s.Description
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDstInstanceUuid() *string {
	return s.DstInstanceUuid
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDstIp() *string {
	return s.DstIp
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDstPort() *int32 {
	return s.DstPort
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDstType() *string {
	return s.DstType
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDtsJobState() *int32 {
	return s.DtsJobState
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetDtsJobStatus() *string {
	return s.DtsJobStatus
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetEndState() *string {
	return s.EndState
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetExternal() *string {
	return s.External
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetRate() *int32 {
	return s.Rate
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetRequestDuration() *int64 {
	return s.RequestDuration
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetSmartPressureTime() *int32 {
	return s.SmartPressureTime
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetSource() *string {
	return s.Source
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetSqlCompleteReuse() *string {
	return s.SqlCompleteReuse
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetSrcInstanceArea() *string {
	return s.SrcInstanceArea
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetSrcInstanceUuid() *string {
	return s.SrcInstanceUuid
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetSrcPublicIp() *string {
	return s.SrcPublicIp
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetState() *string {
	return s.State
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetTableSchema() *string {
	return s.TableSchema
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetTopic() *string {
	return s.Topic
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetUserId() *string {
	return s.UserId
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetVersion() *string {
	return s.Version
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) GetWorkDir() *string {
	return s.WorkDir
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetArchiveJobId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ArchiveJobId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetArchiveOssTableName(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ArchiveOssTableName = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetArchiveState(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ArchiveState = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetBackupId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.BackupId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetBackupType(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.BackupType = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetBenchStep(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.BenchStep = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetBenchStepStatus(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.BenchStepStatus = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetClientGatewayId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ClientGatewayId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetClientType(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ClientType = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDescription(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Description = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDstInstanceUuid(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DstInstanceUuid = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDstIp(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DstIp = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDstPort(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DstPort = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDstType(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DstType = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobClass(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobName(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobName = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobState(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobState = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetDtsJobStatus(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.DtsJobStatus = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetEcsInstanceId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetEndState(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.EndState = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetErrorCode(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ErrorCode = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetErrorMessage(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetExternal(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.External = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetRate(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Rate = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetRequestDuration(v int64) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.RequestDuration = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSmartPressureTime(v int32) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SmartPressureTime = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSource(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Source = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSqlCompleteReuse(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SqlCompleteReuse = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSrcInstanceArea(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SrcInstanceArea = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSrcInstanceUuid(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SrcInstanceUuid = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetSrcPublicIp(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.SrcPublicIp = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetState(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.State = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetStatus(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Status = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetTableSchema(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.TableSchema = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetTaskId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetTaskType(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.TaskType = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetTopic(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Topic = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetUserId(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.UserId = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetVersion(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.Version = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) SetWorkDir(v string) *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks {
	s.WorkDir = &v
	return s
}

func (s *DescribeCloudBenchTasksResponseBodyDataListCloudbenchTasks) Validate() error {
	return dara.Validate(s)
}
