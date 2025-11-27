// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudMigrationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeCloudMigrationResultResponseBodyItems) *DescribeCloudMigrationResultResponseBody
	GetItems() []*DescribeCloudMigrationResultResponseBodyItems
	SetPageNumber(v int64) *DescribeCloudMigrationResultResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCloudMigrationResultResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeCloudMigrationResultResponseBody
	GetRequestId() *string
	SetTotalSize(v int32) *DescribeCloudMigrationResultResponseBody
	GetTotalSize() *int32
}

type DescribeCloudMigrationResultResponseBody struct {
	// The details about the cloud migration task.
	Items []*DescribeCloudMigrationResultResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1B983C48-9793-5EAA-8F7F-00EAEC517675
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeCloudMigrationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudMigrationResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudMigrationResultResponseBody) GetItems() []*DescribeCloudMigrationResultResponseBodyItems {
	return s.Items
}

func (s *DescribeCloudMigrationResultResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCloudMigrationResultResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCloudMigrationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudMigrationResultResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *DescribeCloudMigrationResultResponseBody) SetItems(v []*DescribeCloudMigrationResultResponseBodyItems) *DescribeCloudMigrationResultResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCloudMigrationResultResponseBody) SetPageNumber(v int64) *DescribeCloudMigrationResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBody) SetPageSize(v int64) *DescribeCloudMigrationResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBody) SetRequestId(v string) *DescribeCloudMigrationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBody) SetTotalSize(v int32) *DescribeCloudMigrationResultResponseBody {
	s.TotalSize = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudMigrationResultResponseBodyItems struct {
	// The details about the migration task.
	//
	// example:
	//
	// [Check rds empty]\\nCheck rds databases: success\\n[Check source connectivity]\\nCheck ip connectable: success\\nCheck port connectable: success\\nCheck database connectable: success\\nCheck account replication privilege: success\\nCheck account createrole privilege: success\\nCheck account monitor privilege: success\\n[Check source version]\\nCheck major version consistent: success\\n[Check source glibc version]\\nCheck source glibc version compatible: warning(warning:source glibc version is not compatible with rds pg)\\n[Check disk size]\\nCheck disk size enough: success\\n[Check wal keep size]\\nCheck wal keep size large enough: success\\n[Check spec params]\\nCheck if spec params too large: success\\n[Start RDS instance]\\n2022-02-25 17:00:29 --- Start RDS instance as slave for data replication\\n[Synchronize data]\\n2022-02-25 17:01:05 --- Synchronize data from source to RDS by streaming replication \\n
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2022-02-25T08:53:13Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the task was modified.
	//
	// example:
	//
	// 2022-03-01T06:39:51Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The migration phase of the migration task.
	//
	// 	- **precheck**: precheck
	//
	// 	- **basebackup**: full data backup
	//
	// 	- **startup**: link establishment
	//
	// 	- **increment**: incremental data synchronization
	//
	// 	- **switch**: cloud migration-triggered switchover
	//
	// 	- **success**: cloud migration completed
	//
	// example:
	//
	// switch
	MigrateStage *string `json:"MigrateStage,omitempty" xml:"MigrateStage,omitempty"`
	// The information about the replication link.
	//
	// example:
	//
	// {\\"Status\\":\\"streaming\\",\\"ReceiveStartLsn\\":\\"0/3000000\\",\\"ReceivedTli\\":\\"1\\",\\"LatestEndTime\\":\\"2022-02-25 17:03:59.3344+08\\",\\"Synced\\":\\"true\\",\\"IsSlave\\":\\"true\\",\\"ReplayTimestamp\\":\\"null\\",\\"LastMsgSendTime\\":\\"2022-03-01 14:42:57.967537+08\\",\\"Conninfo\\":\\"user=migratetest password=*******	- channel_binding=prefer dbname=replication host=172.16.254.203 port=5432 application_name=rds_db_instance fallback_application_name=walreceiver sslmode=prefer sslcompression=1 sslsni=1 ssl_min_protocol_version=TLSv1.2 gssencmode=prefer krbsrvname=postgres target_session_attrs=any\\",\\"LastMsgReceiptTime\\":\\"2022-03-01 14:42:57.96727+08\\",\\"LatestEndLsn\\":\\"0/3000148\\",\\"ReceivedLsn\\":\\"0/3000148\\",\\"ReplayLsn\\":\\"0/3000148\\",\\"ReceiveStartTli\\":\\"1\\",\\"ReplayLag\\":\\"0\\"}
	ReplicationInfo *string `json:"ReplicationInfo,omitempty" xml:"ReplicationInfo,omitempty"`
	// The status of data replication.
	//
	// 	- **unstarted**
	//
	// 	- **catchup**
	//
	// 	- **streaming**
	//
	// 	- **disconnect**
	//
	// 	- **finish**
	//
	// example:
	//
	// streaming
	ReplicationState *string `json:"ReplicationState,omitempty" xml:"ReplicationState,omitempty"`
	// The username.
	//
	// example:
	//
	// migratetest
	SourceAccount *string `json:"SourceAccount,omitempty" xml:"SourceAccount,omitempty"`
	// The environment in which the self-managed PostgreSQL instance runs.
	//
	// 	- **idcOnVpc**: The self-managed PostgreSQL instance resides in a data center. The data center can communicate with the VPC to which the ApsaraDB RDS for PostgreSQL instance belongs.
	//
	// 	- **ecsOnVpc**: The self-managed PostgreSQL instance resides on an ECS instance.
	//
	// example:
	//
	// ecsonvpc
	SourceCategory *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	// The private IP address that is used to connect to the self-managed PostgreSQL instance.
	//
	// example:
	//
	// 172.16.XX.XX
	SourceIpAddress *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	// The password.
	//
	// example:
	//
	// 123456
	SourcePassword *string `json:"SourcePassword,omitempty" xml:"SourcePassword,omitempty"`
	// The port number that is used to connect to the self-managed PostgreSQL instance.
	//
	// example:
	//
	// 5432
	SourcePort *int64 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The time when the switchover was performed.
	//
	// example:
	//
	// 2022-03-01T06:40:51Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// A reserved parameter. The return value of this parameter is empty.
	//
	// example:
	//
	// null
	TargetEip *string `json:"TargetEip,omitempty" xml:"TargetEip,omitempty"`
	// The ID of the destination instance.
	//
	// example:
	//
	// pgm-bp102g323jd4****
	TargetInstanceName *string `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 440437220
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// 362c6c7a-4d20-4eac-898c-1495ceab374c
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DescribeCloudMigrationResultResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudMigrationResultResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetDetail() *string {
	return s.Detail
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetMigrateStage() *string {
	return s.MigrateStage
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetReplicationInfo() *string {
	return s.ReplicationInfo
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetReplicationState() *string {
	return s.ReplicationState
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetSourceAccount() *string {
	return s.SourceAccount
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetSourceCategory() *string {
	return s.SourceCategory
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetSourceIpAddress() *string {
	return s.SourceIpAddress
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetSourcePassword() *string {
	return s.SourcePassword
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetSourcePort() *int64 {
	return s.SourcePort
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetTargetEip() *string {
	return s.TargetEip
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetTargetInstanceName() *string {
	return s.TargetInstanceName
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeCloudMigrationResultResponseBodyItems) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetDetail(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.Detail = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetGmtCreated(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetGmtModified(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetMigrateStage(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.MigrateStage = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetReplicationInfo(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.ReplicationInfo = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetReplicationState(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.ReplicationState = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetSourceAccount(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.SourceAccount = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetSourceCategory(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.SourceCategory = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetSourceIpAddress(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.SourceIpAddress = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetSourcePassword(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.SourcePassword = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetSourcePort(v int64) *DescribeCloudMigrationResultResponseBodyItems {
	s.SourcePort = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetSwitchTime(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.SwitchTime = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetTargetEip(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.TargetEip = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetTargetInstanceName(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.TargetInstanceName = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetTaskId(v int64) *DescribeCloudMigrationResultResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) SetTaskName(v string) *DescribeCloudMigrationResultResponseBodyItems {
	s.TaskName = &v
	return s
}

func (s *DescribeCloudMigrationResultResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
