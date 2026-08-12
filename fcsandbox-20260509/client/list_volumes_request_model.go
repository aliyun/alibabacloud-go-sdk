// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVolumesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVolumesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVolumesRequest
	GetNextToken() *string
	SetResourceGroupID(v string) *ListVolumesRequest
	GetResourceGroupID() *string
	SetStatus(v string) *ListVolumesRequest
	GetStatus() *string
	SetTeamID(v string) *ListVolumesRequest
	GetTeamID() *string
	SetUserID(v string) *ListVolumesRequest
	GetUserID() *string
	SetVolumeName(v string) *ListVolumesRequest
	GetVolumeName() *string
}

type ListVolumesRequest struct {
	// example:
	//
	// 5
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// qxGrXje86XMrYQ51aJMy
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// rg-acfmwxqyrgwabcd
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	// example:
	//
	// CREATING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 70d1c834-0383-58d8-97ac-5336eb91abcd
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// example:
	//
	// 210000000
	UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
	// example:
	//
	// workspace
	VolumeName *string `json:"volumeName,omitempty" xml:"volumeName,omitempty"`
}

func (s ListVolumesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVolumesRequest) GoString() string {
	return s.String()
}

func (s *ListVolumesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVolumesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVolumesRequest) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *ListVolumesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListVolumesRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *ListVolumesRequest) GetUserID() *string {
	return s.UserID
}

func (s *ListVolumesRequest) GetVolumeName() *string {
	return s.VolumeName
}

func (s *ListVolumesRequest) SetMaxResults(v int32) *ListVolumesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVolumesRequest) SetNextToken(v string) *ListVolumesRequest {
	s.NextToken = &v
	return s
}

func (s *ListVolumesRequest) SetResourceGroupID(v string) *ListVolumesRequest {
	s.ResourceGroupID = &v
	return s
}

func (s *ListVolumesRequest) SetStatus(v string) *ListVolumesRequest {
	s.Status = &v
	return s
}

func (s *ListVolumesRequest) SetTeamID(v string) *ListVolumesRequest {
	s.TeamID = &v
	return s
}

func (s *ListVolumesRequest) SetUserID(v string) *ListVolumesRequest {
	s.UserID = &v
	return s
}

func (s *ListVolumesRequest) SetVolumeName(v string) *ListVolumesRequest {
	s.VolumeName = &v
	return s
}

func (s *ListVolumesRequest) Validate() error {
	return dara.Validate(s)
}
