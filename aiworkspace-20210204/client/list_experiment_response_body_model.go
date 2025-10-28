// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExperiments(v []*Experiment) *ListExperimentResponseBody
	GetExperiments() []*Experiment
	SetNextPageToken(v int64) *ListExperimentResponseBody
	GetNextPageToken() *int64
	SetTotalCount(v int64) *ListExperimentResponseBody
	GetTotalCount() *int64
	SetRequestId(v string) *ListExperimentResponseBody
	GetRequestId() *string
}

type ListExperimentResponseBody struct {
	// The list of experiments.
	Experiments []*Experiment `json:"Experiments,omitempty" xml:"Experiments,omitempty" type:"Repeated"`
	// The pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 0
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0C6835C5-A424-5AFB-ACC2-F1E3CA1ABF7C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentResponseBody) GetExperiments() []*Experiment {
	return s.Experiments
}

func (s *ListExperimentResponseBody) GetNextPageToken() *int64 {
	return s.NextPageToken
}

func (s *ListExperimentResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExperimentResponseBody) SetExperiments(v []*Experiment) *ListExperimentResponseBody {
	s.Experiments = v
	return s
}

func (s *ListExperimentResponseBody) SetNextPageToken(v int64) *ListExperimentResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListExperimentResponseBody) SetTotalCount(v int64) *ListExperimentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExperimentResponseBody) SetRequestId(v string) *ListExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentResponseBody) Validate() error {
	if s.Experiments != nil {
		for _, item := range s.Experiments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
