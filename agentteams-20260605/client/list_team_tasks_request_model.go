// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListTeamTasksRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListTeamTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTeamTasksRequest
	GetNextToken() *string
	SetTeam(v string) *ListTeamTasksRequest
	GetTeam() *string
}

type ListTeamTasksRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	Team *string `json:"Team,omitempty" xml:"Team,omitempty"`
}

func (s ListTeamTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTeamTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTeamTasksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTeamTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTeamTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTeamTasksRequest) GetTeam() *string {
	return s.Team
}

func (s *ListTeamTasksRequest) SetInstanceId(v string) *ListTeamTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTeamTasksRequest) SetMaxResults(v int32) *ListTeamTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTeamTasksRequest) SetNextToken(v string) *ListTeamTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListTeamTasksRequest) SetTeam(v string) *ListTeamTasksRequest {
	s.Team = &v
	return s
}

func (s *ListTeamTasksRequest) Validate() error {
	return dara.Validate(s)
}
