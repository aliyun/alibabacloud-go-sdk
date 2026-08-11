// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTeamsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTeamsRequest
	GetPageSize() *int32
	SetPlan(v string) *ListTeamsRequest
	GetPlan() *string
	SetResourceGroupID(v string) *ListTeamsRequest
	GetResourceGroupID() *string
	SetTeamName(v string) *ListTeamsRequest
	GetTeamName() *string
}

type ListTeamsRequest struct {
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of teams to display per page.
	//
	// example:
	//
	// 20
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Plan     *string `json:"plan,omitempty" xml:"plan,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmwxqyrgwabcd
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	// The team name.
	//
	// example:
	//
	// DevTeam
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
}

func (s ListTeamsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsRequest) GoString() string {
	return s.String()
}

func (s *ListTeamsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTeamsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTeamsRequest) GetPlan() *string {
	return s.Plan
}

func (s *ListTeamsRequest) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *ListTeamsRequest) GetTeamName() *string {
	return s.TeamName
}

func (s *ListTeamsRequest) SetPageNumber(v int32) *ListTeamsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTeamsRequest) SetPageSize(v int32) *ListTeamsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTeamsRequest) SetPlan(v string) *ListTeamsRequest {
	s.Plan = &v
	return s
}

func (s *ListTeamsRequest) SetResourceGroupID(v string) *ListTeamsRequest {
	s.ResourceGroupID = &v
	return s
}

func (s *ListTeamsRequest) SetTeamName(v string) *ListTeamsRequest {
	s.TeamName = &v
	return s
}

func (s *ListTeamsRequest) Validate() error {
	return dara.Validate(s)
}
