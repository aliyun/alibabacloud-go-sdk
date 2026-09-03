// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListJobHistoryRequest
	GetCount() *int32
	SetMarker(v string) *ListJobHistoryRequest
	GetMarker() *string
	SetRuntimeId(v int32) *ListJobHistoryRequest
	GetRuntimeId() *int32
}

type ListJobHistoryRequest struct {
	// The maximum number of history entries to return.<br> Valid values: 1 to 1000.<br> Default value: 1000.<br><br>
	//
	// example:
	//
	// 100
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The pagination token. Set this parameter to the marker value returned in the previous response to retrieve the next page of results. If not specified, results are returned from the beginning.
	//
	// example:
	//
	// test_marker
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The execution ID of a specific run. Specify this parameter to retrieve the run history for only that execution.
	//
	// example:
	//
	// 1
	RuntimeId *int32 `json:"runtimeId,omitempty" xml:"runtimeId,omitempty"`
}

func (s ListJobHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListJobHistoryRequest) GetCount() *int32 {
	return s.Count
}

func (s *ListJobHistoryRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListJobHistoryRequest) GetRuntimeId() *int32 {
	return s.RuntimeId
}

func (s *ListJobHistoryRequest) SetCount(v int32) *ListJobHistoryRequest {
	s.Count = &v
	return s
}

func (s *ListJobHistoryRequest) SetMarker(v string) *ListJobHistoryRequest {
	s.Marker = &v
	return s
}

func (s *ListJobHistoryRequest) SetRuntimeId(v int32) *ListJobHistoryRequest {
	s.RuntimeId = &v
	return s
}

func (s *ListJobHistoryRequest) Validate() error {
	return dara.Validate(s)
}
