// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetJobs(v []*DatasetJob) *ListDatasetJobsResponseBody
	GetDatasetJobs() []*DatasetJob
	SetRequestId(v string) *ListDatasetJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDatasetJobsResponseBody
	GetTotalCount() *int32
}

type ListDatasetJobsResponseBody struct {
	// The jobs in the dataset.
	DatasetJobs []*DatasetJob `json:"DatasetJobs,omitempty" xml:"DatasetJobs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8D7B2E70-F770-505B-A672-09F1D8F2EC1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of jobs.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasetJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetJobsResponseBody) GetDatasetJobs() []*DatasetJob {
	return s.DatasetJobs
}

func (s *ListDatasetJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasetJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDatasetJobsResponseBody) SetDatasetJobs(v []*DatasetJob) *ListDatasetJobsResponseBody {
	s.DatasetJobs = v
	return s
}

func (s *ListDatasetJobsResponseBody) SetRequestId(v string) *ListDatasetJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetJobsResponseBody) SetTotalCount(v int32) *ListDatasetJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatasetJobsResponseBody) Validate() error {
	if s.DatasetJobs != nil {
		for _, item := range s.DatasetJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
