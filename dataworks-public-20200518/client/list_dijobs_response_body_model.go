// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobPaging(v *ListDIJobsResponseBodyDIJobPaging) *ListDIJobsResponseBody
	GetDIJobPaging() *ListDIJobsResponseBodyDIJobPaging
	SetRequestId(v string) *ListDIJobsResponseBody
	GetRequestId() *string
}

type ListDIJobsResponseBody struct {
	// The pagination information.
	DIJobPaging *ListDIJobsResponseBodyDIJobPaging `json:"DIJobPaging,omitempty" xml:"DIJobPaging,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7263E4AC-9D2E-5B29-B8AF-7C5012E92A41
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponseBody) GetDIJobPaging() *ListDIJobsResponseBodyDIJobPaging {
	return s.DIJobPaging
}

func (s *ListDIJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDIJobsResponseBody) SetDIJobPaging(v *ListDIJobsResponseBodyDIJobPaging) *ListDIJobsResponseBody {
	s.DIJobPaging = v
	return s
}

func (s *ListDIJobsResponseBody) SetRequestId(v string) *ListDIJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDIJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDIJobsResponseBodyDIJobPaging struct {
	// The list of tasks.
	DIJobs []*ListDIJobsResponseBodyDIJobPagingDIJobs `json:"DIJobs,omitempty" xml:"DIJobs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDIJobsResponseBodyDIJobPaging) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobsResponseBodyDIJobPaging) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponseBodyDIJobPaging) GetDIJobs() []*ListDIJobsResponseBodyDIJobPagingDIJobs {
	return s.DIJobs
}

func (s *ListDIJobsResponseBodyDIJobPaging) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDIJobsResponseBodyDIJobPaging) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDIJobsResponseBodyDIJobPaging) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDIJobsResponseBodyDIJobPaging) SetDIJobs(v []*ListDIJobsResponseBodyDIJobPagingDIJobs) *ListDIJobsResponseBodyDIJobPaging {
	s.DIJobs = v
	return s
}

func (s *ListDIJobsResponseBodyDIJobPaging) SetPageNumber(v int32) *ListDIJobsResponseBodyDIJobPaging {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobsResponseBodyDIJobPaging) SetPageSize(v int32) *ListDIJobsResponseBodyDIJobPaging {
	s.PageSize = &v
	return s
}

func (s *ListDIJobsResponseBodyDIJobPaging) SetTotalCount(v int32) *ListDIJobsResponseBodyDIJobPaging {
	s.TotalCount = &v
	return s
}

func (s *ListDIJobsResponseBodyDIJobPaging) Validate() error {
	return dara.Validate(s)
}

type ListDIJobsResponseBodyDIJobPagingDIJobs struct {
	// The task ID.
	//
	// example:
	//
	// 16626
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The type of the destination. The value Hologres is returned.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The task name.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The task status. Valid values:
	//
	// 	- Finished
	//
	// 	- Initialized
	//
	// 	- Stopped
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// example:
	//
	// Finished
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The synchronization type. Valid values:
	//
	// 	- FullAndRealtimeIncremental: one-time full synchronization and real-time incremental synchronization
	//
	// 	- RealtimeIncremental: real-time incremental synchronization
	//
	// 	- Full: one-time full synchronization
	//
	// example:
	//
	// Full
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1967
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the source. The value MySQL is returned.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
}

func (s ListDIJobsResponseBodyDIJobPagingDIJobs) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobsResponseBodyDIJobPagingDIJobs) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) GetDestinationDataSourceType() *string {
	return s.DestinationDataSourceType
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) GetJobName() *string {
	return s.JobName
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) GetJobStatus() *string {
	return s.JobStatus
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) GetMigrationType() *string {
	return s.MigrationType
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) GetSourceDataSourceType() *string {
	return s.SourceDataSourceType
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) SetDIJobId(v int64) *ListDIJobsResponseBodyDIJobPagingDIJobs {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) SetDestinationDataSourceType(v string) *ListDIJobsResponseBodyDIJobPagingDIJobs {
	s.DestinationDataSourceType = &v
	return s
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) SetJobName(v string) *ListDIJobsResponseBodyDIJobPagingDIJobs {
	s.JobName = &v
	return s
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) SetJobStatus(v string) *ListDIJobsResponseBodyDIJobPagingDIJobs {
	s.JobStatus = &v
	return s
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) SetMigrationType(v string) *ListDIJobsResponseBodyDIJobPagingDIJobs {
	s.MigrationType = &v
	return s
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) SetProjectId(v int64) *ListDIJobsResponseBodyDIJobPagingDIJobs {
	s.ProjectId = &v
	return s
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) SetSourceDataSourceType(v string) *ListDIJobsResponseBodyDIJobPagingDIJobs {
	s.SourceDataSourceType = &v
	return s
}

func (s *ListDIJobsResponseBodyDIJobPagingDIJobs) Validate() error {
	return dara.Validate(s)
}
