// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSlowQueryCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSlowQueryCategoriesResponseBody
	GetRequestId() *string
	SetResult(v *ListSlowQueryCategoriesResponseBodyResult) *ListSlowQueryCategoriesResponseBody
	GetResult() *ListSlowQueryCategoriesResponseBodyResult
}

type ListSlowQueryCategoriesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4406F40B-A0A2-9D5D-531F-3B6936567584
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The data returned.
	Result *ListSlowQueryCategoriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListSlowQueryCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSlowQueryCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSlowQueryCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSlowQueryCategoriesResponseBody) GetResult() *ListSlowQueryCategoriesResponseBodyResult {
	return s.Result
}

func (s *ListSlowQueryCategoriesResponseBody) SetRequestId(v string) *ListSlowQueryCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSlowQueryCategoriesResponseBody) SetResult(v *ListSlowQueryCategoriesResponseBodyResult) *ListSlowQueryCategoriesResponseBody {
	s.Result = v
	return s
}

func (s *ListSlowQueryCategoriesResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSlowQueryCategoriesResponseBodyResult struct {
	// The status of the analysis. Valid values:
	//
	// 	- PENDING: preparing
	//
	// 	- SUCCESS: succeeded
	//
	// 	- RUNNING: running
	//
	// 	- FAILED: failed
	//
	// 	- N/A: unknown
	//
	// example:
	//
	// "PENDING"
	AnalyzeStatus *string `json:"analyzeStatus,omitempty" xml:"analyzeStatus,omitempty"`
	// The timestamp that indicates the end of the time range to query.
	//
	// example:
	//
	// 1589990340
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// The timestamp that indicates the beginning of the time range to query.
	//
	// example:
	//
	// 1589986800
	Start *int32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListSlowQueryCategoriesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSlowQueryCategoriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSlowQueryCategoriesResponseBodyResult) GetAnalyzeStatus() *string {
	return s.AnalyzeStatus
}

func (s *ListSlowQueryCategoriesResponseBodyResult) GetEnd() *int32 {
	return s.End
}

func (s *ListSlowQueryCategoriesResponseBodyResult) GetStart() *int32 {
	return s.Start
}

func (s *ListSlowQueryCategoriesResponseBodyResult) SetAnalyzeStatus(v string) *ListSlowQueryCategoriesResponseBodyResult {
	s.AnalyzeStatus = &v
	return s
}

func (s *ListSlowQueryCategoriesResponseBodyResult) SetEnd(v int32) *ListSlowQueryCategoriesResponseBodyResult {
	s.End = &v
	return s
}

func (s *ListSlowQueryCategoriesResponseBodyResult) SetStart(v int32) *ListSlowQueryCategoriesResponseBodyResult {
	s.Start = &v
	return s
}

func (s *ListSlowQueryCategoriesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
