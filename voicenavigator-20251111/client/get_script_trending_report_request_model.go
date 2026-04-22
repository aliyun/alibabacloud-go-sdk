// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScriptTrendingReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetScriptTrendingReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetScriptTrendingReportRequest
	GetInstanceId() *string
	SetScriptId(v string) *GetScriptTrendingReportRequest
	GetScriptId() *string
	SetStartTime(v int64) *GetScriptTrendingReportRequest
	GetStartTime() *int64
	SetTimeInterval(v string) *GetScriptTrendingReportRequest
	GetTimeInterval() *string
}

type GetScriptTrendingReportRequest struct {
	// example:
	//
	// 1582103299434
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 12f407b22cbe4890ac595f09985848d5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 891264b9-5883-4068-94a6-241ceb8d51e4
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// 1582267398628
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1d
	TimeInterval *string `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
}

func (s GetScriptTrendingReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScriptTrendingReportRequest) GoString() string {
	return s.String()
}

func (s *GetScriptTrendingReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetScriptTrendingReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetScriptTrendingReportRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *GetScriptTrendingReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetScriptTrendingReportRequest) GetTimeInterval() *string {
	return s.TimeInterval
}

func (s *GetScriptTrendingReportRequest) SetEndTime(v int64) *GetScriptTrendingReportRequest {
	s.EndTime = &v
	return s
}

func (s *GetScriptTrendingReportRequest) SetInstanceId(v string) *GetScriptTrendingReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetScriptTrendingReportRequest) SetScriptId(v string) *GetScriptTrendingReportRequest {
	s.ScriptId = &v
	return s
}

func (s *GetScriptTrendingReportRequest) SetStartTime(v int64) *GetScriptTrendingReportRequest {
	s.StartTime = &v
	return s
}

func (s *GetScriptTrendingReportRequest) SetTimeInterval(v string) *GetScriptTrendingReportRequest {
	s.TimeInterval = &v
	return s
}

func (s *GetScriptTrendingReportRequest) Validate() error {
	return dara.Validate(s)
}
