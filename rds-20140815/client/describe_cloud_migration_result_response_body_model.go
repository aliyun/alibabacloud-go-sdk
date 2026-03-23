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
	Items      []*DescribeCloudMigrationResultResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int64                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalSize  *int32                                           `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
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
	Detail             *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	GmtCreated         *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified        *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MigrateStage       *string `json:"MigrateStage,omitempty" xml:"MigrateStage,omitempty"`
	ReplicationInfo    *string `json:"ReplicationInfo,omitempty" xml:"ReplicationInfo,omitempty"`
	ReplicationState   *string `json:"ReplicationState,omitempty" xml:"ReplicationState,omitempty"`
	SourceAccount      *string `json:"SourceAccount,omitempty" xml:"SourceAccount,omitempty"`
	SourceCategory     *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	SourceIpAddress    *string `json:"SourceIpAddress,omitempty" xml:"SourceIpAddress,omitempty"`
	SourcePassword     *string `json:"SourcePassword,omitempty" xml:"SourcePassword,omitempty"`
	SourcePort         *int64  `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
	SwitchTime         *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	TargetEip          *string `json:"TargetEip,omitempty" xml:"TargetEip,omitempty"`
	TargetInstanceName *string `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	TaskId             *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName           *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
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
