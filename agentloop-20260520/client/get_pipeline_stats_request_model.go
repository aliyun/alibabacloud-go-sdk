// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetPipelineStatsRequest
	GetEndTime() *int64
	SetGranularity(v string) *GetPipelineStatsRequest
	GetGranularity() *string
	SetStartTime(v int64) *GetPipelineStatsRequest
	GetStartTime() *int64
}

type GetPipelineStatsRequest struct {
	// example:
	//
	// 1735660800
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// Hour
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// example:
	//
	// 1735574400
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetPipelineStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineStatsRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineStatsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetPipelineStatsRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *GetPipelineStatsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetPipelineStatsRequest) SetEndTime(v int64) *GetPipelineStatsRequest {
	s.EndTime = &v
	return s
}

func (s *GetPipelineStatsRequest) SetGranularity(v string) *GetPipelineStatsRequest {
	s.Granularity = &v
	return s
}

func (s *GetPipelineStatsRequest) SetStartTime(v int64) *GetPipelineStatsRequest {
	s.StartTime = &v
	return s
}

func (s *GetPipelineStatsRequest) Validate() error {
	return dara.Validate(s)
}
