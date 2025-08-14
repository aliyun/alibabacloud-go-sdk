// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveSnapshotJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobList(v []*ListLiveSnapshotJobsResponseBodyJobList) *ListLiveSnapshotJobsResponseBody
	GetJobList() []*ListLiveSnapshotJobsResponseBodyJobList
	SetPageNo(v int32) *ListLiveSnapshotJobsResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListLiveSnapshotJobsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListLiveSnapshotJobsResponseBody
	GetRequestId() *string
	SetSortBy(v string) *ListLiveSnapshotJobsResponseBody
	GetSortBy() *string
	SetTotalCount(v int64) *ListLiveSnapshotJobsResponseBody
	GetTotalCount() *int64
}

type ListLiveSnapshotJobsResponseBody struct {
	// The list of jobs.
	JobList []*ListLiveSnapshotJobsResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sorting order of the jobs by creation time.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLiveSnapshotJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotJobsResponseBody) GetJobList() []*ListLiveSnapshotJobsResponseBodyJobList {
	return s.JobList
}

func (s *ListLiveSnapshotJobsResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListLiveSnapshotJobsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveSnapshotJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveSnapshotJobsResponseBody) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveSnapshotJobsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLiveSnapshotJobsResponseBody) SetJobList(v []*ListLiveSnapshotJobsResponseBodyJobList) *ListLiveSnapshotJobsResponseBody {
	s.JobList = v
	return s
}

func (s *ListLiveSnapshotJobsResponseBody) SetPageNo(v int32) *ListLiveSnapshotJobsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBody) SetPageSize(v int32) *ListLiveSnapshotJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBody) SetRequestId(v string) *ListLiveSnapshotJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBody) SetSortBy(v string) *ListLiveSnapshotJobsResponseBody {
	s.SortBy = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBody) SetTotalCount(v int64) *ListLiveSnapshotJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLiveSnapshotJobsResponseBodyJobList struct {
	// The time when the template was created.
	//
	// example:
	//
	// 2022-07-20T02:48:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The job ID.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the job.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The output information.
	SnapshotOutput *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput `json:"SnapshotOutput,omitempty" xml:"SnapshotOutput,omitempty" type:"Struct"`
	// The state of the job.
	//
	// Valid values:
	//
	// 	- init: The job is not started.
	//
	// 	- paused: The job is paused.
	//
	// 	- started: The job is in progress.
	//
	// example:
	//
	// started
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287666****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The interval between two adjacent snapshots. Unit: seconds.
	//
	// example:
	//
	// 5
	TimeInterval *int32 `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
}

func (s ListLiveSnapshotJobsResponseBodyJobList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotJobsResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) GetJobId() *string {
	return s.JobId
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) GetJobName() *string {
	return s.JobName
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) GetSnapshotOutput() *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput {
	return s.SnapshotOutput
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) GetStatus() *string {
	return s.Status
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) GetTimeInterval() *int32 {
	return s.TimeInterval
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) SetCreateTime(v string) *ListLiveSnapshotJobsResponseBodyJobList {
	s.CreateTime = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) SetJobId(v string) *ListLiveSnapshotJobsResponseBodyJobList {
	s.JobId = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) SetJobName(v string) *ListLiveSnapshotJobsResponseBodyJobList {
	s.JobName = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) SetSnapshotOutput(v *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput) *ListLiveSnapshotJobsResponseBodyJobList {
	s.SnapshotOutput = v
	return s
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) SetStatus(v string) *ListLiveSnapshotJobsResponseBodyJobList {
	s.Status = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) SetTemplateId(v string) *ListLiveSnapshotJobsResponseBodyJobList {
	s.TemplateId = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) SetTemplateName(v string) *ListLiveSnapshotJobsResponseBodyJobList {
	s.TemplateName = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) SetTimeInterval(v int32) *ListLiveSnapshotJobsResponseBodyJobList {
	s.TimeInterval = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBodyJobList) Validate() error {
	return dara.Validate(s)
}

type ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput struct {
	// The bucket of the output endpoint. If the storage type is set to oss, the OSS bucket is returned.
	//
	// example:
	//
	// testbucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The output endpoint. If the storage type is set to oss, the Object Storage Service (OSS) domain name is returned.
	//
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The storage type. The value can only be oss.
	//
	// example:
	//
	// oss
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput) GetBucket() *string {
	return s.Bucket
}

func (s *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput) GetStorageType() *string {
	return s.StorageType
}

func (s *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput) SetBucket(v string) *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput {
	s.Bucket = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput) SetEndpoint(v string) *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput {
	s.Endpoint = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput) SetStorageType(v string) *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput {
	s.StorageType = &v
	return s
}

func (s *ListLiveSnapshotJobsResponseBodyJobListSnapshotOutput) Validate() error {
	return dara.Validate(s)
}
