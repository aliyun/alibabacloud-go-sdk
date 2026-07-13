// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkerStatsDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListWorkerStatsDetailsRequest
	GetEndTime() *string
	SetInstanceId(v string) *ListWorkerStatsDetailsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListWorkerStatsDetailsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkerStatsDetailsRequest
	GetNextToken() *string
	SetStartTime(v string) *ListWorkerStatsDetailsRequest
	GetStartTime() *string
}

type ListWorkerStatsDetailsRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListWorkerStatsDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkerStatsDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkerStatsDetailsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListWorkerStatsDetailsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWorkerStatsDetailsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkerStatsDetailsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkerStatsDetailsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListWorkerStatsDetailsRequest) SetEndTime(v string) *ListWorkerStatsDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListWorkerStatsDetailsRequest) SetInstanceId(v string) *ListWorkerStatsDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListWorkerStatsDetailsRequest) SetMaxResults(v int32) *ListWorkerStatsDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkerStatsDetailsRequest) SetNextToken(v string) *ListWorkerStatsDetailsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkerStatsDetailsRequest) SetStartTime(v string) *ListWorkerStatsDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *ListWorkerStatsDetailsRequest) Validate() error {
	return dara.Validate(s)
}
