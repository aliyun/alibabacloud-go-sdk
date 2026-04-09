// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResolveCountSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGranularity(v string) *GetResolveCountSummaryRequest
	GetGranularity() *string
	SetTimeSpan(v int32) *GetResolveCountSummaryRequest
	GetTimeSpan() *int32
}

type GetResolveCountSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// day
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7
	TimeSpan *int32 `json:"TimeSpan,omitempty" xml:"TimeSpan,omitempty"`
}

func (s GetResolveCountSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResolveCountSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetResolveCountSummaryRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *GetResolveCountSummaryRequest) GetTimeSpan() *int32 {
	return s.TimeSpan
}

func (s *GetResolveCountSummaryRequest) SetGranularity(v string) *GetResolveCountSummaryRequest {
	s.Granularity = &v
	return s
}

func (s *GetResolveCountSummaryRequest) SetTimeSpan(v int32) *GetResolveCountSummaryRequest {
	s.TimeSpan = &v
	return s
}

func (s *GetResolveCountSummaryRequest) Validate() error {
	return dara.Validate(s)
}
