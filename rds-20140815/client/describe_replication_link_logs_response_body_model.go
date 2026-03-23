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
	DBInstanceId *string                                         `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Items        []*DescribeReplicationLinkLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalSize    *int32                                          `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
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
	Detail             *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	GmtCreated         *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified        *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ReplicationInfo    *string `json:"ReplicationInfo,omitempty" xml:"ReplicationInfo,omitempty"`
	ReplicationState   *string `json:"ReplicationState,omitempty" xml:"ReplicationState,omitempty"`
	ReplicatorAccount  *string `json:"ReplicatorAccount,omitempty" xml:"ReplicatorAccount,omitempty"`
	ReplicatorPassword *string `json:"ReplicatorPassword,omitempty" xml:"ReplicatorPassword,omitempty"`
	SourceAddress      *string `json:"SourceAddress,omitempty" xml:"SourceAddress,omitempty"`
	SourceCategory     *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	SourcePort         *int64  `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	TargetInstanceId   *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	TaskId             *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName           *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskStage          *string `json:"TaskStage,omitempty" xml:"TaskStage,omitempty"`
	TaskStatus         *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType           *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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
