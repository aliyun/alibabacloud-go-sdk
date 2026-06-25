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
	// The paging token. For the first query, leave this parameter empty. The token for the next page is returned in the response. If the returned token is 0, all results have been returned. You can obtain the token for the next page from the **NextPageToken*	- field in the response.
	//
	// example:
	//
	// 0
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The list of runs.
	Runs []*Run `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
	// The total number of entries that meet the filter criteria. This parameter is optional and may not be returned.
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
