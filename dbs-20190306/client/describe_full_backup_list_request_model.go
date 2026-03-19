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
	SetBackupSetStatus(v string) *DescribeFullBackupListRequest
	GetBackupSetStatus() *string
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
	SetShowProgress(v string) *DescribeFullBackupListRequest
	GetShowProgress() *string
	SetShowStorageType(v bool) *DescribeFullBackupListRequest
	GetShowStorageType() *bool
	SetStartTimestamp(v int64) *DescribeFullBackupListRequest
	GetStartTimestamp() *int64
}

type DescribeFullBackupListRequest struct {
	// The ID of the backup plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbsr179qz******
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the backup set.
	//
	// example:
	//
	// 1iukx5h******
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// example:
	//
	// finish
	BackupSetStatus *string `json:"BackupSetStatus,omitempty" xml:"BackupSetStatus,omitempty"`
	// A token that ensures idempotence and prevents duplicate requests.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The end time of the backup, in UNIX timestamp format.
	//
	// example:
	//
	// 1676887128
	EndTimestamp *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: integers from 0 to the maximum integer value. Default value: 0.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - 30
	//
	// - 50
	//
	// - 100
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	ShowProgress *string `json:"ShowProgress,omitempty" xml:"ShowProgress,omitempty"`
	// Specifies whether to return the storage class.
	//
	// example:
	//
	// true
	ShowStorageType *bool `json:"ShowStorageType,omitempty" xml:"ShowStorageType,omitempty"`
	// The start time of the backup.
	//
	// example:
	//
	// 1676887100
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

func (s *DescribeFullBackupListRequest) GetBackupSetStatus() *string {
	return s.BackupSetStatus
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

func (s *DescribeFullBackupListRequest) GetShowProgress() *string {
	return s.ShowProgress
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

func (s *DescribeFullBackupListRequest) SetBackupSetStatus(v string) *DescribeFullBackupListRequest {
	s.BackupSetStatus = &v
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

func (s *DescribeFullBackupListRequest) SetShowProgress(v string) *DescribeFullBackupListRequest {
	s.ShowProgress = &v
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
