// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobSanityCheckResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestID(v string) *ListJobSanityCheckResultsResponseBody
	GetRequestID() *string
	SetSanityCheckResults(v [][]*SanityCheckResultItem) *ListJobSanityCheckResultsResponseBody
	GetSanityCheckResults() [][]*SanityCheckResultItem
	SetTotalCount(v int32) *ListJobSanityCheckResultsResponseBody
	GetTotalCount() *int32
}

type ListJobSanityCheckResultsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AC9xxx-3xxx-5xxx2-xxxx-FA5
	RequestID *string `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	// The sanity check results.
	SanityCheckResults [][]*SanityCheckResultItem `json:"SanityCheckResults,omitempty" xml:"SanityCheckResults,omitempty" type:"Repeated"`
	// The total number of results that meet the filter conditions.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobSanityCheckResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobSanityCheckResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobSanityCheckResultsResponseBody) GetRequestID() *string {
	return s.RequestID
}

func (s *ListJobSanityCheckResultsResponseBody) GetSanityCheckResults() [][]*SanityCheckResultItem {
	return s.SanityCheckResults
}

func (s *ListJobSanityCheckResultsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListJobSanityCheckResultsResponseBody) SetRequestID(v string) *ListJobSanityCheckResultsResponseBody {
	s.RequestID = &v
	return s
}

func (s *ListJobSanityCheckResultsResponseBody) SetSanityCheckResults(v [][]*SanityCheckResultItem) *ListJobSanityCheckResultsResponseBody {
	s.SanityCheckResults = v
	return s
}

func (s *ListJobSanityCheckResultsResponseBody) SetTotalCount(v int32) *ListJobSanityCheckResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobSanityCheckResultsResponseBody) Validate() error {
	return dara.Validate(s)
}
