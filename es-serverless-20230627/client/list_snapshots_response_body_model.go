// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListSnapshotsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSnapshotsResponseBody
	GetRequestId() *string
	SetResult(v []map[string]interface{}) *ListSnapshotsResponseBody
	GetResult() []map[string]interface{}
	SetTotalCount(v int32) *ListSnapshotsResponseBody
	GetTotalCount() *int32
}

type ListSnapshotsResponseBody struct {
	// example:
	//
	// ODgyObrnP3
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 22E9EE78-F567-550A-8F7C-20E9CD3DE489
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSnapshotsResponseBody) GetResult() []map[string]interface{} {
	return s.Result
}

func (s *ListSnapshotsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSnapshotsResponseBody) SetNextToken(v string) *ListSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetRequestId(v string) *ListSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetResult(v []map[string]interface{}) *ListSnapshotsResponseBody {
	s.Result = v
	return s
}

func (s *ListSnapshotsResponseBody) SetTotalCount(v int32) *ListSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSnapshotsResponseBody) Validate() error {
	return dara.Validate(s)
}
