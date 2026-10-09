// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupabaseBackupJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMode(v string) *ListSupabaseBackupJobsRequest
	GetBackupMode() *string
	SetMaxResults(v int32) *ListSupabaseBackupJobsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSupabaseBackupJobsRequest
	GetNextToken() *string
	SetProjectId(v string) *ListSupabaseBackupJobsRequest
	GetProjectId() *string
	SetRegionId(v string) *ListSupabaseBackupJobsRequest
	GetRegionId() *string
}

type ListSupabaseBackupJobsRequest struct {
	// The backup mode. Valid values:
	//
	// - Automated: automatic backup
	//
	// - Manual: manual backup
	//
	// If this parameter is not specified, all backup tasks are returned.
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The maximum number of entries to return for this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The paging token. Do not specify this parameter for the first query. For subsequent queries, specify the NextToken value returned in the previous response.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
}

func (s ListSupabaseBackupJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseBackupJobsRequest) GoString() string {
	return s.String()
}

func (s *ListSupabaseBackupJobsRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *ListSupabaseBackupJobsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSupabaseBackupJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupabaseBackupJobsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListSupabaseBackupJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSupabaseBackupJobsRequest) SetBackupMode(v string) *ListSupabaseBackupJobsRequest {
	s.BackupMode = &v
	return s
}

func (s *ListSupabaseBackupJobsRequest) SetMaxResults(v int32) *ListSupabaseBackupJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSupabaseBackupJobsRequest) SetNextToken(v string) *ListSupabaseBackupJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSupabaseBackupJobsRequest) SetProjectId(v string) *ListSupabaseBackupJobsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListSupabaseBackupJobsRequest) SetRegionId(v string) *ListSupabaseBackupJobsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSupabaseBackupJobsRequest) Validate() error {
	return dara.Validate(s)
}
