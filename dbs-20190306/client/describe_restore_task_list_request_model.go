// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DescribeRestoreTaskListRequest
	GetBackupPlanId() *string
	SetClientToken(v string) *DescribeRestoreTaskListRequest
	GetClientToken() *string
	SetEndTimestamp(v int64) *DescribeRestoreTaskListRequest
	GetEndTimestamp() *int64
	SetOwnerId(v string) *DescribeRestoreTaskListRequest
	GetOwnerId() *string
	SetPageNum(v int32) *DescribeRestoreTaskListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeRestoreTaskListRequest
	GetPageSize() *int32
	SetRestoreTaskId(v string) *DescribeRestoreTaskListRequest
	GetRestoreTaskId() *string
	SetStartTimestamp(v int64) *DescribeRestoreTaskListRequest
	GetStartTimestamp() *int64
}

type DescribeRestoreTaskListRequest struct {
	// The ID of the backup plan.
	//
	// > Specify either BackupPlanId or RestoreTaskId.
	//
	// example:
	//
	// dbs1hvb0wwwXXXXX
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// A client token that is used to ensure the idempotence of the request. This prevents duplicate requests.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The end time of the backup.
	//
	// example:
	//
	// 1570701361528
	EndTimestamp *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be greater than or equal to 0 and not exceed the maximum value of the integer data type. Default value: 0.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// > Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the restore job. You can specify multiple IDs. Separate them with commas (,). Call the [CreateRestoreTask](https://help.aliyun.com/document_detail/2869836.html) operation to obtain this parameter.
	//
	// > Specify either RestoreTaskId or BackupPlanId. An error occurs if you specify both parameters.
	//
	// example:
	//
	// s102h7rf5anq
	RestoreTaskId *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	// The start time of the backup.
	//
	// example:
	//
	// 1570701361528
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRestoreTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTaskListRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeRestoreTaskListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeRestoreTaskListRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeRestoreTaskListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeRestoreTaskListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeRestoreTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreTaskListRequest) GetRestoreTaskId() *string {
	return s.RestoreTaskId
}

func (s *DescribeRestoreTaskListRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeRestoreTaskListRequest) SetBackupPlanId(v string) *DescribeRestoreTaskListRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) SetClientToken(v string) *DescribeRestoreTaskListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) SetEndTimestamp(v int64) *DescribeRestoreTaskListRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) SetOwnerId(v string) *DescribeRestoreTaskListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) SetPageNum(v int32) *DescribeRestoreTaskListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) SetPageSize(v int32) *DescribeRestoreTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) SetRestoreTaskId(v string) *DescribeRestoreTaskListRequest {
	s.RestoreTaskId = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) SetStartTimestamp(v int64) *DescribeRestoreTaskListRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeRestoreTaskListRequest) Validate() error {
	return dara.Validate(s)
}
