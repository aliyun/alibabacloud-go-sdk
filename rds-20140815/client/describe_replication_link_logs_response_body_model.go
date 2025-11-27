// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicationLinkLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeReplicationLinkLogsResponseBody
	GetDBInstanceId() *string
	SetItems(v []*DescribeReplicationLinkLogsResponseBodyItems) *DescribeReplicationLinkLogsResponseBody
	GetItems() []*DescribeReplicationLinkLogsResponseBodyItems
	SetRequestId(v string) *DescribeReplicationLinkLogsResponseBody
	GetRequestId() *string
	SetTotalSize(v int32) *DescribeReplicationLinkLogsResponseBody
	GetTotalSize() *int32
}

type DescribeReplicationLinkLogsResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// pgm-bp1trqb4p1xd****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The items.
	Items []*DescribeReplicationLinkLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 16C62438-491B-5C02-9B49-BA924A1372A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeReplicationLinkLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationLinkLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReplicationLinkLogsResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeReplicationLinkLogsResponseBody) GetItems() []*DescribeReplicationLinkLogsResponseBodyItems {
	return s.Items
}

func (s *DescribeReplicationLinkLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeReplicationLinkLogsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *DescribeReplicationLinkLogsResponseBody) SetDBInstanceId(v string) *DescribeReplicationLinkLogsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBody) SetItems(v []*DescribeReplicationLinkLogsResponseBodyItems) *DescribeReplicationLinkLogsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBody) SetRequestId(v string) *DescribeReplicationLinkLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBody) SetTotalSize(v int32) *DescribeReplicationLinkLogsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBody) Validate() error {
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

type DescribeReplicationLinkLogsResponseBodyItems struct {
	// The details of the task.
	//
	// example:
	//
	// [Check rds empty]\\nCheck rds databases: success\\n[Check source connectivity]\\nCheck ip connectable: success\\nCheck port connectable: success\\nCheck database connectable: success\\nCheck account replication privilege: success\\nCheck account createrole privilege: success\\nCheck account monitor privilege: success\\n[Check source version]\\nCheck major version consistent: success\\n[Check source glibc version]\\nCheck source glibc version compatible: warning(warning:source glibc version is not compatible with rds pg)\\n[Check disk size]\\nCheck disk size enough: success\\n[Check wal keep size]\\nCheck wal keep size large enough: success\\n[Check spec params]\\nCheck if spec params too large: success\\n [Check triggers]\\nCheck triggers compatible: success\\n[Check user functions]\\nCheck user functions compatible: success\\n*Migrate check success*
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The creation time. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-02-25T06:57:41Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The modification time. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-01T06:39:51Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The synchronization information. This parameter is a reserved parameter.
	//
	// example:
	//
	// None
	ReplicationInfo *string `json:"ReplicationInfo,omitempty" xml:"ReplicationInfo,omitempty"`
	// The status of the synchronization. Valid values:
	//
	// 	- **steaming**: The synchronization is in progress.
	//
	// 	- **finish**: The synchronization is complete.
	//
	// 	- **disconnect**: The synchronization is disconnected.
	//
	// example:
	//
	// finish
	ReplicationState *string `json:"ReplicationState,omitempty" xml:"ReplicationState,omitempty"`
	// The account of the database that is used for data synchronization.
	//
	// example:
	//
	// testdbuser
	ReplicatorAccount *string `json:"ReplicatorAccount,omitempty" xml:"ReplicatorAccount,omitempty"`
	// The password of the account.
	//
	// example:
	//
	// testpassword
	ReplicatorPassword *string `json:"ReplicatorPassword,omitempty" xml:"ReplicatorPassword,omitempty"`
	// The endpoint of the source instance.
	//
	// example:
	//
	// pgm-****.pg.rds.aliyuncs.com
	SourceAddress *string `json:"SourceAddress,omitempty" xml:"SourceAddress,omitempty"`
	// The type of the source instance. Valid values:
	//
	// 	- other: other instances
	//
	// 	- aliyunRDS: an ApsaraDB RDS instance
	//
	// example:
	//
	// aliyunRDS
	SourceCategory *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	// The port number of the source instance.
	//
	// example:
	//
	// 5432
	SourcePort *int64 `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	// The destination instance ID.
	//
	// example:
	//
	// pgm-bp1l4dutw453****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 8413252
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// test01
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The stage of the task. Valid values:
	//
	// 	- **precheck**: the precheck stage.
	//
	// 	- **basebackup**: the basic backup stage.
	//
	// 	- **startup**: the startup stage.
	//
	// 	- **increment**: the incremental synchronization stage.
	//
	// example:
	//
	// increment
	TaskStage *string `json:"TaskStage,omitempty" xml:"TaskStage,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **success**
	//
	// 	- **failure**
	//
	// 	- **running**
	//
	// example:
	//
	// success
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **create**: creates a synchronization link.
	//
	// 	- **create-dryrun**: performs a precheck before a synchronization link is created.
	//
	// example:
	//
	// create
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeReplicationLinkLogsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationLinkLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetDetail() *string {
	return s.Detail
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetReplicationInfo() *string {
	return s.ReplicationInfo
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetReplicationState() *string {
	return s.ReplicationState
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetReplicatorAccount() *string {
	return s.ReplicatorAccount
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetReplicatorPassword() *string {
	return s.ReplicatorPassword
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetSourceAddress() *string {
	return s.SourceAddress
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetSourceCategory() *string {
	return s.SourceCategory
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetSourcePort() *int64 {
	return s.SourcePort
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetTaskStage() *string {
	return s.TaskStage
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetDetail(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.Detail = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetGmtCreated(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetGmtModified(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetReplicationInfo(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.ReplicationInfo = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetReplicationState(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.ReplicationState = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetReplicatorAccount(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.ReplicatorAccount = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetReplicatorPassword(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.ReplicatorPassword = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetSourceAddress(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.SourceAddress = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetSourceCategory(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.SourceCategory = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetSourcePort(v int64) *DescribeReplicationLinkLogsResponseBodyItems {
	s.SourcePort = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetTargetInstanceId(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.TargetInstanceId = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetTaskId(v int64) *DescribeReplicationLinkLogsResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetTaskName(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.TaskName = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetTaskStage(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.TaskStage = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetTaskStatus(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.TaskStatus = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) SetTaskType(v string) *DescribeReplicationLinkLogsResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *DescribeReplicationLinkLogsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
