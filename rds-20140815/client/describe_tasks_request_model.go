// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeTasksRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeTasksRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTasksRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTasksRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeTasksRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeTasksRequest
	GetStatus() *string
	SetTaskAction(v string) *DescribeTasksRequest
	GetTaskAction() *string
}

type DescribeTasksRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2020-11-20T02:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Valid values: **30 to 100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2020-11-20T01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task. This parameter is invalid.
	//
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The operation that is used by the task.
	//
	// example:
	//
	// CreateInstance
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeTasksRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTasksRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeTasksRequest) GetTaskAction() *string {
	return s.TaskAction
}

func (s *DescribeTasksRequest) SetDBInstanceId(v string) *DescribeTasksRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeTasksRequest) SetEndTime(v string) *DescribeTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerAccount(v string) *DescribeTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerId(v int64) *DescribeTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetPageNumber(v int32) *DescribeTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksRequest) SetPageSize(v int32) *DescribeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerAccount(v string) *DescribeTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerId(v int64) *DescribeTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetStartTime(v string) *DescribeTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksRequest) SetStatus(v string) *DescribeTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskAction(v string) *DescribeTasksRequest {
	s.TaskAction = &v
	return s
}

func (s *DescribeTasksRequest) Validate() error {
	return dara.Validate(s)
}
