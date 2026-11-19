// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppTraceDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAiAppTraceDetailRequest
	GetAppId() *string
	SetEndTime(v string) *GetAiAppTraceDetailRequest
	GetEndTime() *string
	SetRegionId(v string) *GetAiAppTraceDetailRequest
	GetRegionId() *string
	SetStartTime(v string) *GetAiAppTraceDetailRequest
	GetStartTime() *string
	SetTraceId(v string) *GetAiAppTraceDetailRequest
	GetTraceId() *string
}

type GetAiAppTraceDetailRequest struct {
	// The ID of the AI application.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the query. Specify the time in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2026-01-02 16:08:38
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the query. Specify the time in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2026-01-01 16:08:38
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4d3b7c9e2f1a4b8c9d0e1f2a3b4c5d6e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s GetAiAppTraceDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppTraceDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAiAppTraceDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAiAppTraceDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetAiAppTraceDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAiAppTraceDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetAiAppTraceDetailRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *GetAiAppTraceDetailRequest) SetAppId(v string) *GetAiAppTraceDetailRequest {
	s.AppId = &v
	return s
}

func (s *GetAiAppTraceDetailRequest) SetEndTime(v string) *GetAiAppTraceDetailRequest {
	s.EndTime = &v
	return s
}

func (s *GetAiAppTraceDetailRequest) SetRegionId(v string) *GetAiAppTraceDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetAiAppTraceDetailRequest) SetStartTime(v string) *GetAiAppTraceDetailRequest {
	s.StartTime = &v
	return s
}

func (s *GetAiAppTraceDetailRequest) SetTraceId(v string) *GetAiAppTraceDetailRequest {
	s.TraceId = &v
	return s
}

func (s *GetAiAppTraceDetailRequest) Validate() error {
	return dara.Validate(s)
}
