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
	// > You must specify either this parameter or `RestoreTaskId`.
	//
	// example:
	//
	// dbs1hvb0wwwXXXXX
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// A client-generated token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The end of the time range to query restore tasks. Specify the time in the UNIX timestamp format. The timestamp is a millisecond-level timestamp.
	//
	// example:
	//
	// 1570701361528
	EndTimestamp *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be an integer that is greater than or equal to 0. Default value: 0.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// > The default value is 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the restore task. You can specify multiple IDs, separated by commas. You can call the [CreateRestoreTask](https://help.aliyun.com/document_detail/2869836.html) operation to obtain the restore task ID.
	//
	// > You must specify either this parameter or `BackupPlanId`. Specifying both parameters returns an error.
	//
	// example:
	//
	// s102h7rf5anq
	RestoreTaskId *string `json:"RestoreTaskId,omitempty" xml:"RestoreTaskId,omitempty"`
	// The beginning of the time range to query restore tasks. Specify the time in the UNIX timestamp format. The timestamp is a millisecond-level timestamp.
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
