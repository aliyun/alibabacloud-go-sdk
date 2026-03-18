// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOralEvaluationStatisticsConcurrentCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationAccessId(v string) *OralEvaluationStatisticsConcurrentCountRequest
	GetApplicationAccessId() *string
	SetEndTime(v string) *OralEvaluationStatisticsConcurrentCountRequest
	GetEndTime() *string
	SetGranularity(v string) *OralEvaluationStatisticsConcurrentCountRequest
	GetGranularity() *string
	SetProjectId(v string) *OralEvaluationStatisticsConcurrentCountRequest
	GetProjectId() *string
	SetStartTime(v string) *OralEvaluationStatisticsConcurrentCountRequest
	GetStartTime() *string
}

type OralEvaluationStatisticsConcurrentCountRequest struct {
	// appId,appkey
	//
	// example:
	//
	// a0007g7
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// example:
	//
	// 2024-09-29 14:22:48
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
	// 2024-09-29 05:00:01
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s OralEvaluationStatisticsConcurrentCountRequest) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsConcurrentCountRequest) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) SetApplicationAccessId(v string) *OralEvaluationStatisticsConcurrentCountRequest {
	s.ApplicationAccessId = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) SetEndTime(v string) *OralEvaluationStatisticsConcurrentCountRequest {
	s.EndTime = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) SetGranularity(v string) *OralEvaluationStatisticsConcurrentCountRequest {
	s.Granularity = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) SetProjectId(v string) *OralEvaluationStatisticsConcurrentCountRequest {
	s.ProjectId = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) SetStartTime(v string) *OralEvaluationStatisticsConcurrentCountRequest {
	s.StartTime = &v
	return s
}

func (s *OralEvaluationStatisticsConcurrentCountRequest) Validate() error {
	return dara.Validate(s)
}
