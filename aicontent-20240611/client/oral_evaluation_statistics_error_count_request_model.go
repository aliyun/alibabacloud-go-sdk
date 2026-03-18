// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOralEvaluationStatisticsErrorCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationAccessId(v string) *OralEvaluationStatisticsErrorCountRequest
	GetApplicationAccessId() *string
	SetEndTime(v string) *OralEvaluationStatisticsErrorCountRequest
	GetEndTime() *string
	SetErrorCode(v []*string) *OralEvaluationStatisticsErrorCountRequest
	GetErrorCode() []*string
	SetGranularity(v string) *OralEvaluationStatisticsErrorCountRequest
	GetGranularity() *string
	SetProjectId(v string) *OralEvaluationStatisticsErrorCountRequest
	GetProjectId() *string
	SetStartTime(v string) *OralEvaluationStatisticsErrorCountRequest
	GetStartTime() *string
}

type OralEvaluationStatisticsErrorCountRequest struct {
	// appId,appkey
	//
	// example:
	//
	// a0007g7
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// example:
	//
	// 2024-08-22 06:24:53
	EndTime   *string   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ErrorCode []*string `json:"errorCode,omitempty" xml:"errorCode,omitempty" type:"Repeated"`
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
	// 2024-09-27 09:32:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s OralEvaluationStatisticsErrorCountRequest) String() string {
	return dara.Prettify(s)
}

func (s OralEvaluationStatisticsErrorCountRequest) GoString() string {
	return s.String()
}

func (s *OralEvaluationStatisticsErrorCountRequest) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *OralEvaluationStatisticsErrorCountRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *OralEvaluationStatisticsErrorCountRequest) GetErrorCode() []*string {
	return s.ErrorCode
}

func (s *OralEvaluationStatisticsErrorCountRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *OralEvaluationStatisticsErrorCountRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *OralEvaluationStatisticsErrorCountRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *OralEvaluationStatisticsErrorCountRequest) SetApplicationAccessId(v string) *OralEvaluationStatisticsErrorCountRequest {
	s.ApplicationAccessId = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountRequest) SetEndTime(v string) *OralEvaluationStatisticsErrorCountRequest {
	s.EndTime = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountRequest) SetErrorCode(v []*string) *OralEvaluationStatisticsErrorCountRequest {
	s.ErrorCode = v
	return s
}

func (s *OralEvaluationStatisticsErrorCountRequest) SetGranularity(v string) *OralEvaluationStatisticsErrorCountRequest {
	s.Granularity = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountRequest) SetProjectId(v string) *OralEvaluationStatisticsErrorCountRequest {
	s.ProjectId = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountRequest) SetStartTime(v string) *OralEvaluationStatisticsErrorCountRequest {
	s.StartTime = &v
	return s
}

func (s *OralEvaluationStatisticsErrorCountRequest) Validate() error {
	return dara.Validate(s)
}
