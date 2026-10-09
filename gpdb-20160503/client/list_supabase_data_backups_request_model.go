// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupabaseDataBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *ListSupabaseDataBackupsRequest
	GetBackupId() *string
	SetBackupMode(v string) *ListSupabaseDataBackupsRequest
	GetBackupMode() *string
	SetBackupStatus(v string) *ListSupabaseDataBackupsRequest
	GetBackupStatus() *string
	SetDataType(v string) *ListSupabaseDataBackupsRequest
	GetDataType() *string
	SetEndTime(v string) *ListSupabaseDataBackupsRequest
	GetEndTime() *string
	SetMaxResults(v int32) *ListSupabaseDataBackupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSupabaseDataBackupsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListSupabaseDataBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSupabaseDataBackupsRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListSupabaseDataBackupsRequest
	GetProjectId() *string
	SetRegionId(v string) *ListSupabaseDataBackupsRequest
	GetRegionId() *string
	SetStartTime(v string) *ListSupabaseDataBackupsRequest
	GetStartTime() *string
}

type ListSupabaseDataBackupsRequest struct {
	// The ID of the backup set. You can obtain the ID from the BackupSetId parameter returned by the ListSupabaseDataBackups operation.
	//
	// example:
	//
	// 327329803
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup pattern. Valid values: Automated: automatic backup. Manual: manual backup.
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The status of the backup set. Valid values: Success: the backup is successful; Failed: the backup fails.
	//
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The backup type. Valid values: DATA: full backup; RESTOREPOI: restorable point.
	//
	// example:
	//
	// DATA
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The end time of the query. The end time must be later than the start time. Format: yyyy-MM-ddTHH:mmZ (UTC).
	//
	// example:
	//
	// 2011-06-01T16:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries to return for the current request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The paging token for paged query. Do not specify this parameter for the first request. For subsequent requests, specify the NextToken value returned by the previous response.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. The value must be greater than 0 and cannot exceed the maximum value of an integer. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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
	// Instance ID of the Supabase instance. You can obtain instance ID from the Supabase page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// sbp-263****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the query. Format: yyyy-MM-ddTHH:mmZ (UTC).
	//
	// example:
	//
	// 2011-06-01T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListSupabaseDataBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseDataBackupsRequest) GoString() string {
	return s.String()
}

func (s *ListSupabaseDataBackupsRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *ListSupabaseDataBackupsRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *ListSupabaseDataBackupsRequest) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *ListSupabaseDataBackupsRequest) GetDataType() *string {
	return s.DataType
}

func (s *ListSupabaseDataBackupsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListSupabaseDataBackupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSupabaseDataBackupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupabaseDataBackupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSupabaseDataBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSupabaseDataBackupsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListSupabaseDataBackupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSupabaseDataBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListSupabaseDataBackupsRequest) SetBackupId(v string) *ListSupabaseDataBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *ListSupabaseDataBackupsRequest) SetBackupMode(v string) *ListSupabaseDataBackupsRequest {
	s.BackupMode = &v
	return s
}

func (s *ListSupabaseDataBackupsRequest) SetBackupStatus(v string) *ListSupabaseDataBackupsRequest {
	s.BackupStatus = &v
	return s
}

func (s *ListSupabaseDataBackupsRequest) SetDataType(v string) *ListSupabaseDataBackupsRequest {
	s.DataType = &v
	return s
}

func (s *ListSupabaseDataBackupsRequest) SetEndTime(v string) *ListSupabaseDataBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *ListSupabaseDataBackupsRequest) SetMaxResults(v int32) *ListSupabaseDataBackupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSupabaseDataBackupsRequest) SetNextToken(v string) *ListSupabaseDataBackupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSupabaseDataBackupsRequest) SetPageNumber(v int32) *ListSupabaseDataBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSupabaseDataBackupsRequest) SetPageSize(v int32) *ListSupabaseDataBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSupabaseDataBackupsRequest) SetProjectId(v string) *ListSupabaseDataBackupsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListSupabaseDataBackupsRequest) SetRegionId(v string) *ListSupabaseDataBackupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSupabaseDataBackupsRequest) SetStartTime(v string) *ListSupabaseDataBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *ListSupabaseDataBackupsRequest) Validate() error {
	return dara.Validate(s)
}
