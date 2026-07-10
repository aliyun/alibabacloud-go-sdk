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
	SetResourceGroupID(v string) *ListTeamsRequest
	GetResourceGroupID() *string
	SetTeamName(v string) *ListTeamsRequest
	GetTeamName() *string
}

type ListTeamsRequest struct {
	PageNumber      *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize        *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	TeamName        *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
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
