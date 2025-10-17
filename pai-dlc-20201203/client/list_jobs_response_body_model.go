// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*JobItem) *ListJobsResponseBody
	GetJobs() []*JobItem
	SetRequestId(v string) *ListJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListJobsResponseBody
	GetTotalCount() *int64
}

type ListJobsResponseBody struct {
	// The jobs.
	Jobs []*JobItem `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The request ID used to troubleshoot issues.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of jobs that meet the filter conditions.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) GetJobs() []*JobItem {
	return s.Jobs
}

func (s *ListJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListJobsResponseBody) SetJobs(v []*JobItem) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int64) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobsResponseBody) Validate() error {
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
