// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationDataSourceType(v string) *ListDIJobsRequest
	GetDestinationDataSourceType() *string
	SetJobName(v string) *ListDIJobsRequest
	GetJobName() *string
	SetPageNumber(v int32) *ListDIJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDIJobsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDIJobsRequest
	GetProjectId() *int64
	SetSourceDataSourceType(v string) *ListDIJobsRequest
	GetSourceDataSourceType() *string
}

type ListDIJobsRequest struct {
	// The destination type. If you do not configure this parameter, no limits are imposed on the tasks.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The name of the task. Fuzzy match is supported. If you do not configure this parameter, no limits are imposed on the tasks.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1967
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The source type. If you do not configure this parameter, no limits are imposed on the tasks.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
}

func (s ListDIJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobsRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobsRequest) GetDestinationDataSourceType() *string {
	return s.DestinationDataSourceType
}

func (s *ListDIJobsRequest) GetJobName() *string {
	return s.JobName
}

func (s *ListDIJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDIJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDIJobsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDIJobsRequest) GetSourceDataSourceType() *string {
	return s.SourceDataSourceType
}

func (s *ListDIJobsRequest) SetDestinationDataSourceType(v string) *ListDIJobsRequest {
	s.DestinationDataSourceType = &v
	return s
}

func (s *ListDIJobsRequest) SetJobName(v string) *ListDIJobsRequest {
	s.JobName = &v
	return s
}

func (s *ListDIJobsRequest) SetPageNumber(v int32) *ListDIJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobsRequest) SetPageSize(v int32) *ListDIJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDIJobsRequest) SetProjectId(v int64) *ListDIJobsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDIJobsRequest) SetSourceDataSourceType(v string) *ListDIJobsRequest {
	s.SourceDataSourceType = &v
	return s
}

func (s *ListDIJobsRequest) Validate() error {
	return dara.Validate(s)
}
