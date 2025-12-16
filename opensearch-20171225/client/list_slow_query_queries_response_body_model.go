// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSlowQueryQueriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSlowQueryQueriesResponseBody
	GetRequestId() *string
	SetResult(v *ListSlowQueryQueriesResponseBodyResult) *ListSlowQueryQueriesResponseBody
	GetResult() *ListSlowQueryQueriesResponseBodyResult
}

type ListSlowQueryQueriesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EB250CA0-ACFD-C5DE-17CD-01445BFE8AE5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result *ListSlowQueryQueriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListSlowQueryQueriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSlowQueryQueriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSlowQueryQueriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSlowQueryQueriesResponseBody) GetResult() *ListSlowQueryQueriesResponseBodyResult {
	return s.Result
}

func (s *ListSlowQueryQueriesResponseBody) SetRequestId(v string) *ListSlowQueryQueriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBody) SetResult(v *ListSlowQueryQueriesResponseBodyResult) *ListSlowQueryQueriesResponseBody {
	s.Result = v
	return s
}

func (s *ListSlowQueryQueriesResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSlowQueryQueriesResponseBodyResult struct {
	// The content of the optimization suggestion for the query.
	//
	// example:
	//
	// no data
	AppQuery *string `json:"appQuery,omitempty" xml:"appQuery,omitempty"`
	// The end of the time range that was queried.
	//
	// example:
	//
	// 1589990340
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// The ID of the optimization suggestion.
	//
	// example:
	//
	// 0
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// The beginning of the time range that was queried.
	//
	// example:
	//
	// 1589986800
	Start *int32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListSlowQueryQueriesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSlowQueryQueriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSlowQueryQueriesResponseBodyResult) GetAppQuery() *string {
	return s.AppQuery
}

func (s *ListSlowQueryQueriesResponseBodyResult) GetEnd() *int32 {
	return s.End
}

func (s *ListSlowQueryQueriesResponseBodyResult) GetIndex() *int32 {
	return s.Index
}

func (s *ListSlowQueryQueriesResponseBodyResult) GetStart() *int32 {
	return s.Start
}

func (s *ListSlowQueryQueriesResponseBodyResult) SetAppQuery(v string) *ListSlowQueryQueriesResponseBodyResult {
	s.AppQuery = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBodyResult) SetEnd(v int32) *ListSlowQueryQueriesResponseBodyResult {
	s.End = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBodyResult) SetIndex(v int32) *ListSlowQueryQueriesResponseBodyResult {
	s.Index = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBodyResult) SetStart(v int32) *ListSlowQueryQueriesResponseBodyResult {
	s.Start = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
