// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFullBackupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *DescribeFullBackupListRequest
	GetBackupPlanId() *string
	SetBackupSetId(v string) *DescribeFullBackupListRequest
	GetBackupSetId() *string
	SetClientToken(v string) *DescribeFullBackupListRequest
	GetClientToken() *string
	SetEndTimestamp(v int64) *DescribeFullBackupListRequest
	GetEndTimestamp() *int64
	SetOwnerId(v string) *DescribeFullBackupListRequest
	GetOwnerId() *string
	SetPageNum(v int32) *DescribeFullBackupListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeFullBackupListRequest
	GetPageSize() *int32
	SetShowStorageType(v bool) *DescribeFullBackupListRequest
	GetShowStorageType() *bool
	SetStartTimestamp(v int64) *DescribeFullBackupListRequest
	GetStartTimestamp() *int64
}

type DescribeFullBackupListRequest struct {
	// The error code.
	//
	// This parameter is required.
	//
	// example:
	//
	// The total number of full backup tasks.
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The end time of the backup task, such as 1554560477000.
	//
	// example:
	//
	// The point in time when the backup set expires, such as 1554560477000.
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndTimestamp *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The error message.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// true
	ShowStorageType *bool `json:"ShowStorageType,omitempty" xml:"ShowStorageType,omitempty"`
	// Queries full backup tasks.
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeFullBackupListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFullBackupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeFullBackupListRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *DescribeFullBackupListRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeFullBackupListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeFullBackupListRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeFullBackupListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeFullBackupListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeFullBackupListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFullBackupListRequest) GetShowStorageType() *bool {
	return s.ShowStorageType
}

func (s *DescribeFullBackupListRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeFullBackupListRequest) SetBackupPlanId(v string) *DescribeFullBackupListRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetBackupSetId(v string) *DescribeFullBackupListRequest {
	s.BackupSetId = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetClientToken(v string) *DescribeFullBackupListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetEndTimestamp(v int64) *DescribeFullBackupListRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetOwnerId(v string) *DescribeFullBackupListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetPageNum(v int32) *DescribeFullBackupListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetPageSize(v int32) *DescribeFullBackupListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetShowStorageType(v bool) *DescribeFullBackupListRequest {
	s.ShowStorageType = &v
	return s
}

func (s *DescribeFullBackupListRequest) SetStartTimestamp(v int64) *DescribeFullBackupListRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeFullBackupListRequest) Validate() error {
	return dara.Validate(s)
}
