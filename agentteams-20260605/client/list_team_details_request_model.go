// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListTeamDetailsRequest
	GetEndTime() *string
	SetInstanceId(v string) *ListTeamDetailsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListTeamDetailsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTeamDetailsRequest
	GetNextToken() *string
	SetStartTime(v string) *ListTeamDetailsRequest
	GetStartTime() *string
}

type ListTeamDetailsRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListTeamDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTeamDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListTeamDetailsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTeamDetailsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTeamDetailsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTeamDetailsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTeamDetailsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTeamDetailsRequest) SetEndTime(v string) *ListTeamDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTeamDetailsRequest) SetInstanceId(v string) *ListTeamDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTeamDetailsRequest) SetMaxResults(v int32) *ListTeamDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTeamDetailsRequest) SetNextToken(v string) *ListTeamDetailsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTeamDetailsRequest) SetStartTime(v string) *ListTeamDetailsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTeamDetailsRequest) Validate() error {
	return dara.Validate(s)
}
