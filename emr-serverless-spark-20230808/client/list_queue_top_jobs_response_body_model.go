// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueTopJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueueTopJobs(v *ListQueueTopJobsResponseBodyQueueTopJobs) *ListQueueTopJobsResponseBody
	GetQueueTopJobs() *ListQueueTopJobsResponseBodyQueueTopJobs
	SetRequestId(v string) *ListQueueTopJobsResponseBody
	GetRequestId() *string
}

type ListQueueTopJobsResponseBody struct {
	// The job details.
	QueueTopJobs *ListQueueTopJobsResponseBodyQueueTopJobs `json:"queueTopJobs,omitempty" xml:"queueTopJobs,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListQueueTopJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQueueTopJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueTopJobsResponseBody) GetQueueTopJobs() *ListQueueTopJobsResponseBodyQueueTopJobs {
	return s.QueueTopJobs
}

func (s *ListQueueTopJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQueueTopJobsResponseBody) SetQueueTopJobs(v *ListQueueTopJobsResponseBodyQueueTopJobs) *ListQueueTopJobsResponseBody {
	s.QueueTopJobs = v
	return s
}

func (s *ListQueueTopJobsResponseBody) SetRequestId(v string) *ListQueueTopJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueueTopJobsResponseBody) Validate() error {
	if s.QueueTopJobs != nil {
		if err := s.QueueTopJobs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQueueTopJobsResponseBodyQueueTopJobs struct {
	// The list of jobs.
	Jobs []*ListQueueTopJobsResponseBodyQueueTopJobsJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	// The maximum number of records returned in a single request.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// dev_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// The total number of CUs consumed by the queue.
	//
	// example:
	//
	// 128
	TotalCu *string `json:"totalCu,omitempty" xml:"totalCu,omitempty"`
	// The total number of jobs in the queue.
	//
	// example:
	//
	// 25
	TotalJobCount *int32 `json:"totalJobCount,omitempty" xml:"totalJobCount,omitempty"`
}

func (s ListQueueTopJobsResponseBodyQueueTopJobs) String() string {
	return dara.Prettify(s)
}

func (s ListQueueTopJobsResponseBodyQueueTopJobs) GoString() string {
	return s.String()
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobs) GetJobs() []*ListQueueTopJobsResponseBodyQueueTopJobsJobs {
	return s.Jobs
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobs) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobs) GetQueueName() *string {
	return s.QueueName
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobs) GetTotalCu() *string {
	return s.TotalCu
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobs) GetTotalJobCount() *int32 {
	return s.TotalJobCount
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobs) SetJobs(v []*ListQueueTopJobsResponseBodyQueueTopJobsJobs) *ListQueueTopJobsResponseBodyQueueTopJobs {
	s.Jobs = v
	return s
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobs) SetMaxResults(v string) *ListQueueTopJobsResponseBodyQueueTopJobs {
	s.MaxResults = &v
	return s
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobs) SetQueueName(v string) *ListQueueTopJobsResponseBodyQueueTopJobs {
	s.QueueName = &v
	return s
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobs) SetTotalCu(v string) *ListQueueTopJobsResponseBodyQueueTopJobs {
	s.TotalCu = &v
	return s
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobs) SetTotalJobCount(v int32) *ListQueueTopJobsResponseBodyQueueTopJobs {
	s.TotalJobCount = &v
	return s
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobs) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQueueTopJobsResponseBodyQueueTopJobsJobs struct {
	// The number of CPU cores requested by the job.
	//
	// example:
	//
	// 8
	CpuCores *string `json:"cpuCores,omitempty" xml:"cpuCores,omitempty"`
	// The number of CUs consumed by the job.
	//
	// example:
	//
	// 8
	Cu *string `json:"cu,omitempty" xml:"cu,omitempty"`
	// The job ID.
	//
	// example:
	//
	// jr-1234abcd
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The job type.
	//
	// example:
	//
	// spark-batch
	JobType *string `json:"jobType,omitempty" xml:"jobType,omitempty"`
	// The amount of memory requested by the job.
	//
	// example:
	//
	// 32
	MemoryGiB *string `json:"memoryGiB,omitempty" xml:"memoryGiB,omitempty"`
}

func (s ListQueueTopJobsResponseBodyQueueTopJobsJobs) String() string {
	return dara.Prettify(s)
}

func (s ListQueueTopJobsResponseBodyQueueTopJobsJobs) GoString() string {
	return s.String()
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobsJobs) GetCpuCores() *string {
	return s.CpuCores
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobsJobs) GetCu() *string {
	return s.Cu
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobsJobs) GetJobId() *string {
	return s.JobId
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobsJobs) GetJobType() *string {
	return s.JobType
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobsJobs) GetMemoryGiB() *string {
	return s.MemoryGiB
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobsJobs) SetCpuCores(v string) *ListQueueTopJobsResponseBodyQueueTopJobsJobs {
	s.CpuCores = &v
	return s
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobsJobs) SetCu(v string) *ListQueueTopJobsResponseBodyQueueTopJobsJobs {
	s.Cu = &v
	return s
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobsJobs) SetJobId(v string) *ListQueueTopJobsResponseBodyQueueTopJobsJobs {
	s.JobId = &v
	return s
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobsJobs) SetJobType(v string) *ListQueueTopJobsResponseBodyQueueTopJobsJobs {
	s.JobType = &v
	return s
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobsJobs) SetMemoryGiB(v string) *ListQueueTopJobsResponseBodyQueueTopJobsJobs {
	s.MemoryGiB = &v
	return s
}

func (s *ListQueueTopJobsResponseBodyQueueTopJobsJobs) Validate() error {
	return dara.Validate(s)
}
