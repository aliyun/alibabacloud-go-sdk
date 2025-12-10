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
	// Specifies the number of running records of the migration task to be returned.\\
	//
	// Valid values: 0 - 1000.\\
	//
	// Default value: 1000.
	//
	// example:
	//
	// 100
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The marker after which the running history of the task is listed.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// test_marker
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The execution ID of the task. If you specify an execution ID, only the running history related to the execution ID is listed.
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
