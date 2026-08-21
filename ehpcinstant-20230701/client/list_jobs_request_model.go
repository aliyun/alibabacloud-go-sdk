// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *ListJobsRequestFilter) *ListJobsRequest
	GetFilter() *ListJobsRequestFilter
	SetPageNumber(v int32) *ListJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobsRequest
	GetPageSize() *int32
	SetSortBy(v *ListJobsRequestSortBy) *ListJobsRequest
	GetSortBy() *ListJobsRequestSortBy
}

type ListJobsRequest struct {
	// The filter conditions for querying jobs.
	Filter *ListJobsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The current page number.
	//
	// Start value: 1
	//
	// Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The default value is 50. The maximum value is 100.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting method.
	SortBy *ListJobsRequestSortBy `json:"SortBy,omitempty" xml:"SortBy,omitempty" type:"Struct"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetFilter() *ListJobsRequestFilter {
	return s.Filter
}

func (s *ListJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsRequest) GetSortBy() *ListJobsRequestSortBy {
	return s.SortBy
}

func (s *ListJobsRequest) SetFilter(v *ListJobsRequestFilter) *ListJobsRequest {
	s.Filter = v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetSortBy(v *ListJobsRequestSortBy) *ListJobsRequest {
	s.SortBy = v
	return s
}

func (s *ListJobsRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	if s.SortBy != nil {
		if err := s.SortBy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsRequestFilter struct {
	// The ID of the job.
	//
	// example:
	//
	// job-xxxx
	JobId  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobIds []*string `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
	// The name of the job. Fuzzy search is supported.
	//
	// example:
	//
	// testJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// jt-xxxx
	JobTemplateId *string `json:"JobTemplateId,omitempty" xml:"JobTemplateId,omitempty"`
	// example:
	//
	// MyPool
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The status of the job. Valid values:
	//
	// - Pending: The job is in the queue.
	//
	// - Initing: The job is initializing.
	//
	// - Succeeded: The job was successful.
	//
	// - Failed: The job failed.
	//
	// - Running: The job is running.
	//
	// - Exception: A scheduling exception occurred.
	//
	// - Retrying: The job is being retried.
	//
	// - Expired: The job timed out.
	//
	// - Suspended: The job is in hibernation.
	//
	// - Restarting: The job is restarting.
	//
	// - Deleted: The job is deleted.
	//
	// example:
	//
	// Running
	Status *string                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag    []*ListJobsRequestFilterTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The time after which the jobs were submitted. This is a UNIX timestamp based on the local time of the region. For sites in the Chinese mainland, the time zone is UTC+8.
	//
	// example:
	//
	// 1703819914
	TimeCreatedAfter *int32 `json:"TimeCreatedAfter,omitempty" xml:"TimeCreatedAfter,omitempty"`
	// The time before which the jobs were submitted. This is a UNIX timestamp based on the local time of the region. For sites in the Chinese mainland, the time zone is UTC+8.
	//
	// example:
	//
	// 1703820113
	TimeCreatedBefore *int32 `json:"TimeCreatedBefore,omitempty" xml:"TimeCreatedBefore,omitempty"`
}

func (s ListJobsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListJobsRequestFilter) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsRequestFilter) GetJobIds() []*string {
	return s.JobIds
}

func (s *ListJobsRequestFilter) GetJobName() *string {
	return s.JobName
}

func (s *ListJobsRequestFilter) GetJobTemplateId() *string {
	return s.JobTemplateId
}

func (s *ListJobsRequestFilter) GetPoolName() *string {
	return s.PoolName
}

func (s *ListJobsRequestFilter) GetStatus() *string {
	return s.Status
}

func (s *ListJobsRequestFilter) GetTag() []*ListJobsRequestFilterTag {
	return s.Tag
}

func (s *ListJobsRequestFilter) GetTimeCreatedAfter() *int32 {
	return s.TimeCreatedAfter
}

func (s *ListJobsRequestFilter) GetTimeCreatedBefore() *int32 {
	return s.TimeCreatedBefore
}

func (s *ListJobsRequestFilter) SetJobId(v string) *ListJobsRequestFilter {
	s.JobId = &v
	return s
}

func (s *ListJobsRequestFilter) SetJobIds(v []*string) *ListJobsRequestFilter {
	s.JobIds = v
	return s
}

func (s *ListJobsRequestFilter) SetJobName(v string) *ListJobsRequestFilter {
	s.JobName = &v
	return s
}

func (s *ListJobsRequestFilter) SetJobTemplateId(v string) *ListJobsRequestFilter {
	s.JobTemplateId = &v
	return s
}

func (s *ListJobsRequestFilter) SetPoolName(v string) *ListJobsRequestFilter {
	s.PoolName = &v
	return s
}

func (s *ListJobsRequestFilter) SetStatus(v string) *ListJobsRequestFilter {
	s.Status = &v
	return s
}

func (s *ListJobsRequestFilter) SetTag(v []*ListJobsRequestFilterTag) *ListJobsRequestFilter {
	s.Tag = v
	return s
}

func (s *ListJobsRequestFilter) SetTimeCreatedAfter(v int32) *ListJobsRequestFilter {
	s.TimeCreatedAfter = &v
	return s
}

func (s *ListJobsRequestFilter) SetTimeCreatedBefore(v int32) *ListJobsRequestFilter {
	s.TimeCreatedBefore = &v
	return s
}

func (s *ListJobsRequestFilter) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobsRequestFilterTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListJobsRequestFilterTag) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequestFilterTag) GoString() string {
	return s.String()
}

func (s *ListJobsRequestFilterTag) GetKey() *string {
	return s.Key
}

func (s *ListJobsRequestFilterTag) GetValue() *string {
	return s.Value
}

func (s *ListJobsRequestFilterTag) SetKey(v string) *ListJobsRequestFilterTag {
	s.Key = &v
	return s
}

func (s *ListJobsRequestFilterTag) SetValue(v string) *ListJobsRequestFilterTag {
	s.Value = &v
	return s
}

func (s *ListJobsRequestFilterTag) Validate() error {
	return dara.Validate(s)
}

type ListJobsRequestSortBy struct {
	// The field to sort by. Valid values:
	//
	// - time_start
	//
	// - job_name
	//
	// example:
	//
	// time_start
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The sort order. Valid values:
	//
	// - ASC (default): Ascending
	//
	// - DESC: Descending
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ListJobsRequestSortBy) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequestSortBy) GoString() string {
	return s.String()
}

func (s *ListJobsRequestSortBy) GetLabel() *string {
	return s.Label
}

func (s *ListJobsRequestSortBy) GetOrder() *string {
	return s.Order
}

func (s *ListJobsRequestSortBy) SetLabel(v string) *ListJobsRequestSortBy {
	s.Label = &v
	return s
}

func (s *ListJobsRequestSortBy) SetOrder(v string) *ListJobsRequestSortBy {
	s.Order = &v
	return s
}

func (s *ListJobsRequestSortBy) Validate() error {
	return dara.Validate(s)
}
