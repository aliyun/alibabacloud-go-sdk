// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v int64) *ListRunsResponseBody
	GetNextPageToken() *int64
	SetRuns(v []*Run) *ListRunsResponseBody
	GetRuns() []*Run
	SetTotalCount(v int64) *ListRunsResponseBody
	GetTotalCount() *int64
	SetRequestId(v string) *ListRunsResponseBody
	GetRequestId() *string
}

type ListRunsResponseBody struct {
	// The pagination token that is used to retrieve the next page. You do not need to specify this parameter for the first request. You must specify the pagination token in the result of the previous query. If the pagination token is 0, no next page exists. You can obtain the pagination token that is used to retrieve the next page in the value of the **NextPageToken*	- field.
	//
	// example:
	//
	// 0
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The runs.
	Runs []*Run `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
	// The total number of entries returned. By default, this parameter is not returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRunsResponseBody) GetNextPageToken() *int64 {
	return s.NextPageToken
}

func (s *ListRunsResponseBody) GetRuns() []*Run {
	return s.Runs
}

func (s *ListRunsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRunsResponseBody) SetNextPageToken(v int64) *ListRunsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListRunsResponseBody) SetRuns(v []*Run) *ListRunsResponseBody {
	s.Runs = v
	return s
}

func (s *ListRunsResponseBody) SetTotalCount(v int64) *ListRunsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRunsResponseBody) SetRequestId(v string) *ListRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRunsResponseBody) Validate() error {
	if s.Runs != nil {
		for _, item := range s.Runs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
