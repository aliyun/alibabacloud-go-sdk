// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeInstanceStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *ListRealtimeInstanceStatsRequest
	GetInstanceIds() []*string
	SetPageNumber(v int32) *ListRealtimeInstanceStatsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRealtimeInstanceStatsRequest
	GetPageSize() *int32
}

type ListRealtimeInstanceStatsRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRealtimeInstanceStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeInstanceStatsRequest) GoString() string {
	return s.String()
}

func (s *ListRealtimeInstanceStatsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListRealtimeInstanceStatsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRealtimeInstanceStatsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRealtimeInstanceStatsRequest) SetInstanceIds(v []*string) *ListRealtimeInstanceStatsRequest {
	s.InstanceIds = v
	return s
}

func (s *ListRealtimeInstanceStatsRequest) SetPageNumber(v int32) *ListRealtimeInstanceStatsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRealtimeInstanceStatsRequest) SetPageSize(v int32) *ListRealtimeInstanceStatsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRealtimeInstanceStatsRequest) Validate() error {
	return dara.Validate(s)
}
