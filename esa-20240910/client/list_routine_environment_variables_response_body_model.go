// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineEnvironmentVariablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListRoutineEnvironmentVariablesResponseBody
	GetCount() *int64
	SetEnvironmentVariables(v map[string]*EnvironmentVariablesValue) *ListRoutineEnvironmentVariablesResponseBody
	GetEnvironmentVariables() map[string]*EnvironmentVariablesValue
	SetPageNumber(v int64) *ListRoutineEnvironmentVariablesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRoutineEnvironmentVariablesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListRoutineEnvironmentVariablesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRoutineEnvironmentVariablesResponseBody
	GetTotalCount() *int64
}

type ListRoutineEnvironmentVariablesResponseBody struct {
	// The number of environment variables.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The environment variable dictionary.
	EnvironmentVariables map[string]*EnvironmentVariablesValue `json:"EnvironmentVariables,omitempty" xml:"EnvironmentVariables,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of environment variables.
	//
	// example:
	//
	// 16
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRoutineEnvironmentVariablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineEnvironmentVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoutineEnvironmentVariablesResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListRoutineEnvironmentVariablesResponseBody) GetEnvironmentVariables() map[string]*EnvironmentVariablesValue {
	return s.EnvironmentVariables
}

func (s *ListRoutineEnvironmentVariablesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRoutineEnvironmentVariablesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRoutineEnvironmentVariablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRoutineEnvironmentVariablesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRoutineEnvironmentVariablesResponseBody) SetCount(v int64) *ListRoutineEnvironmentVariablesResponseBody {
	s.Count = &v
	return s
}

func (s *ListRoutineEnvironmentVariablesResponseBody) SetEnvironmentVariables(v map[string]*EnvironmentVariablesValue) *ListRoutineEnvironmentVariablesResponseBody {
	s.EnvironmentVariables = v
	return s
}

func (s *ListRoutineEnvironmentVariablesResponseBody) SetPageNumber(v int64) *ListRoutineEnvironmentVariablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRoutineEnvironmentVariablesResponseBody) SetPageSize(v int64) *ListRoutineEnvironmentVariablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRoutineEnvironmentVariablesResponseBody) SetRequestId(v string) *ListRoutineEnvironmentVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoutineEnvironmentVariablesResponseBody) SetTotalCount(v int64) *ListRoutineEnvironmentVariablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRoutineEnvironmentVariablesResponseBody) Validate() error {
	return dara.Validate(s)
}
