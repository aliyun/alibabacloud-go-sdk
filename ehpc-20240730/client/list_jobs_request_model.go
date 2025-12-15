// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListJobsRequest
	GetClusterId() *string
	SetJobFilter(v *ListJobsRequestJobFilter) *ListJobsRequest
	GetJobFilter() *ListJobsRequestJobFilter
	SetPageNumber(v string) *ListJobsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListJobsRequest
	GetPageSize() *string
}

type ListJobsRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-csbua72***
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job filter information.
	JobFilter *ListJobsRequestJobFilter `json:"JobFilter,omitempty" xml:"JobFilter,omitempty" type:"Struct"`
	// The page number of the page to return.
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// 	- Maximum value: 50.
	//
	// 	- Default value: 10
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListJobsRequest) GetJobFilter() *ListJobsRequestJobFilter {
	return s.JobFilter
}

func (s *ListJobsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListJobsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListJobsRequest) SetClusterId(v string) *ListJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobsRequest) SetJobFilter(v *ListJobsRequestJobFilter) *ListJobsRequest {
	s.JobFilter = v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v string) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v string) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) Validate() error {
	if s.JobFilter != nil {
		if err := s.JobFilter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsRequestJobFilter struct {
	// The time when the job was last updated. The value is a UNIX timestamp representing the number of seconds that have elapsed since 1970-01-01T00:00:00Z.
	//
	// example:
	//
	// 1724123085
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The time when the job started. The value is a UNIX timestamp representing the number of seconds that have elapsed since 1970-01-01T00:00:00Z.
	//
	// example:
	//
	// 1724122486
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// Job diagnosis and analysis list.
	Diagnosis []*ListJobsRequestJobFilterDiagnosis `json:"Diagnosis,omitempty" xml:"Diagnosis,omitempty" type:"Repeated"`
	// The job name. Fuzzy match is supported.
	//
	// example:
	//
	// testjob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The job status. Valid values:
	//
	// 	- all: returns all jobs.
	//
	// 	- finished: returns completed jobs.
	//
	// 	- notfinish: returns uncompleted jobs.
	//
	// Default value: all.
	//
	// example:
	//
	// all
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The compute nodes that run the jobs.
	Nodes []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The queues to which the jobs belong.
	Queues []*string `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	// The result sorting configurations.
	SortBy *ListJobsRequestJobFilterSortBy `json:"SortBy,omitempty" xml:"SortBy,omitempty" type:"Struct"`
	// The users that run the jobs.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListJobsRequestJobFilter) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequestJobFilter) GoString() string {
	return s.String()
}

func (s *ListJobsRequestJobFilter) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *ListJobsRequestJobFilter) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *ListJobsRequestJobFilter) GetDiagnosis() []*ListJobsRequestJobFilterDiagnosis {
	return s.Diagnosis
}

func (s *ListJobsRequestJobFilter) GetJobName() *string {
	return s.JobName
}

func (s *ListJobsRequestJobFilter) GetJobStatus() *string {
	return s.JobStatus
}

func (s *ListJobsRequestJobFilter) GetNodes() []*string {
	return s.Nodes
}

func (s *ListJobsRequestJobFilter) GetQueues() []*string {
	return s.Queues
}

func (s *ListJobsRequestJobFilter) GetSortBy() *ListJobsRequestJobFilterSortBy {
	return s.SortBy
}

func (s *ListJobsRequestJobFilter) GetUsers() []*string {
	return s.Users
}

func (s *ListJobsRequestJobFilter) SetCreateTimeEnd(v string) *ListJobsRequestJobFilter {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListJobsRequestJobFilter) SetCreateTimeStart(v string) *ListJobsRequestJobFilter {
	s.CreateTimeStart = &v
	return s
}

func (s *ListJobsRequestJobFilter) SetDiagnosis(v []*ListJobsRequestJobFilterDiagnosis) *ListJobsRequestJobFilter {
	s.Diagnosis = v
	return s
}

func (s *ListJobsRequestJobFilter) SetJobName(v string) *ListJobsRequestJobFilter {
	s.JobName = &v
	return s
}

func (s *ListJobsRequestJobFilter) SetJobStatus(v string) *ListJobsRequestJobFilter {
	s.JobStatus = &v
	return s
}

func (s *ListJobsRequestJobFilter) SetNodes(v []*string) *ListJobsRequestJobFilter {
	s.Nodes = v
	return s
}

func (s *ListJobsRequestJobFilter) SetQueues(v []*string) *ListJobsRequestJobFilter {
	s.Queues = v
	return s
}

func (s *ListJobsRequestJobFilter) SetSortBy(v *ListJobsRequestJobFilterSortBy) *ListJobsRequestJobFilter {
	s.SortBy = v
	return s
}

func (s *ListJobsRequestJobFilter) SetUsers(v []*string) *ListJobsRequestJobFilter {
	s.Users = v
	return s
}

func (s *ListJobsRequestJobFilter) Validate() error {
	if s.Diagnosis != nil {
		for _, item := range s.Diagnosis {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SortBy != nil {
		if err := s.SortBy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsRequestJobFilterDiagnosis struct {
	// Job diagnosis threshold comparator.
	//
	// example:
	//
	// greater
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Job diagnosis and analysis metrics
	//
	// example:
	//
	// run_duration
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// Job diagnosis threshold.
	//
	// example:
	//
	// 24
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListJobsRequestJobFilterDiagnosis) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequestJobFilterDiagnosis) GoString() string {
	return s.String()
}

func (s *ListJobsRequestJobFilterDiagnosis) GetOperator() *string {
	return s.Operator
}

func (s *ListJobsRequestJobFilterDiagnosis) GetOption() *string {
	return s.Option
}

func (s *ListJobsRequestJobFilterDiagnosis) GetThreshold() *string {
	return s.Threshold
}

func (s *ListJobsRequestJobFilterDiagnosis) SetOperator(v string) *ListJobsRequestJobFilterDiagnosis {
	s.Operator = &v
	return s
}

func (s *ListJobsRequestJobFilterDiagnosis) SetOption(v string) *ListJobsRequestJobFilterDiagnosis {
	s.Option = &v
	return s
}

func (s *ListJobsRequestJobFilterDiagnosis) SetThreshold(v string) *ListJobsRequestJobFilterDiagnosis {
	s.Threshold = &v
	return s
}

func (s *ListJobsRequestJobFilterDiagnosis) Validate() error {
	return dara.Validate(s)
}

type ListJobsRequestJobFilterSortBy struct {
	// The order in which jobs are sorted based on their execution time. Valid values:
	//
	// 	- asc: in ascending order.
	//
	// 	- desc: in descending order.
	//
	// Default value: desc.
	//
	// example:
	//
	// asc
	ExecuteOrder *string `json:"ExecuteOrder,omitempty" xml:"ExecuteOrder,omitempty"`
	// The order in which jobs are sorted based on their queuing time. Valid values:
	//
	// 	- asc: in ascending order.
	//
	// 	- desc: in descending order.
	//
	// Default value: desc.
	//
	// example:
	//
	// desc
	PendOrder *string `json:"PendOrder,omitempty" xml:"PendOrder,omitempty"`
	// The order in which jobs are sorted based on their submitting time. Valid values:
	//
	// 	- asc: in ascending order.
	//
	// 	- desc: in descending order.
	//
	// Default value: desc.
	//
	// example:
	//
	// asc
	SubmitOrder *string `json:"SubmitOrder,omitempty" xml:"SubmitOrder,omitempty"`
}

func (s ListJobsRequestJobFilterSortBy) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequestJobFilterSortBy) GoString() string {
	return s.String()
}

func (s *ListJobsRequestJobFilterSortBy) GetExecuteOrder() *string {
	return s.ExecuteOrder
}

func (s *ListJobsRequestJobFilterSortBy) GetPendOrder() *string {
	return s.PendOrder
}

func (s *ListJobsRequestJobFilterSortBy) GetSubmitOrder() *string {
	return s.SubmitOrder
}

func (s *ListJobsRequestJobFilterSortBy) SetExecuteOrder(v string) *ListJobsRequestJobFilterSortBy {
	s.ExecuteOrder = &v
	return s
}

func (s *ListJobsRequestJobFilterSortBy) SetPendOrder(v string) *ListJobsRequestJobFilterSortBy {
	s.PendOrder = &v
	return s
}

func (s *ListJobsRequestJobFilterSortBy) SetSubmitOrder(v string) *ListJobsRequestJobFilterSortBy {
	s.SubmitOrder = &v
	return s
}

func (s *ListJobsRequestJobFilterSortBy) Validate() error {
	return dara.Validate(s)
}
