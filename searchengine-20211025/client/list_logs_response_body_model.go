// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListLogsResponseBody
	GetRequestId() *string
	SetResult(v *ListLogsResponseBodyResult) *ListLogsResponseBody
	GetResult() *ListLogsResponseBodyResult
}

type ListLogsResponseBody struct {
	// id of request
	//
	// example:
	//
	// 022F36C7-9FB4-5D67-BEBC-3D14B0984463
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// ListResult
	Result *ListLogsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogsResponseBody) GetResult() *ListLogsResponseBodyResult {
	return s.Result
}

func (s *ListLogsResponseBody) SetRequestId(v string) *ListLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogsResponseBody) SetResult(v *ListLogsResponseBodyResult) *ListLogsResponseBody {
	s.Result = v
	return s
}

func (s *ListLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLogsResponseBodyResult struct {
	// The result.
	Result []interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListLogsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListLogsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListLogsResponseBodyResult) GetResult() []interface{} {
	return s.Result
}

func (s *ListLogsResponseBodyResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLogsResponseBodyResult) SetResult(v []interface{}) *ListLogsResponseBodyResult {
	s.Result = v
	return s
}

func (s *ListLogsResponseBodyResult) SetTotalCount(v int32) *ListLogsResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListLogsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
