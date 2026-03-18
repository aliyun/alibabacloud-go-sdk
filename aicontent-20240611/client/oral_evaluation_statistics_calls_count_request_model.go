// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOralEvaluationStatisticsCallsCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationAccessId(v string) *OralEvaluationStatisticsCallsCountRequest
	GetApplicationAccessId() *string
	SetEndTime(v string) *OralEvaluationStatisticsCallsCountRequest
	GetEndTime() *string
	SetGranularity(v string) *OralEvaluationStatisticsCallsCountRequest
	GetGranularity() *string
	SetProjectId(v string) *OralEvaluationStatisticsCallsCountRequest
	GetProjectId() *string
	SetStartTime(v string) *OralEvaluationStatisticsCallsCountRequest
	GetStartTime() *string
}

type OralEvaluationStatisticsCallsCountRequest struct {
	// appId,appkey
	//
	// example:
	//
	// a0007g7
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// example:
	//
	// 2024-10-15 07:40:01
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// DAY
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// example:
	//
	// 123
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 2024-10-14 07:40:01
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s OralEvaluationStatisticsCallsCountRequest) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsCallsCountRequest) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsCallsCountRequest) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *OralEvaluationStatisticsCallsCountRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *OralEvaluationStatisticsCallsCountRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *OralEvaluationStatisticsCallsCountRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *OralEvaluationStatisticsCallsCountRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *OralEvaluationStatisticsCallsCountRequest) SetApplicationAccessId(v string) *OralEvaluationStatisticsCallsCountRequest {
	s.ApplicationAccessId = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountRequest) SetEndTime(v string) *OralEvaluationStatisticsCallsCountRequest {
	s.EndTime = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountRequest) SetGranularity(v string) *OralEvaluationStatisticsCallsCountRequest {
	s.Granularity = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountRequest) SetProjectId(v string) *OralEvaluationStatisticsCallsCountRequest {
	s.ProjectId = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountRequest) SetStartTime(v string) *OralEvaluationStatisticsCallsCountRequest {
	s.StartTime = &v
	return s
}

func (s *OralEvaluationStatisticsCallsCountRequest) Validate() error {
	return dara.Validate(s)
}
