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
	SetStorageClass(v string) *ListVolumesRequest
	GetStorageClass() *string
	SetTeamID(v string) *ListVolumesRequest
	GetTeamID() *string
	SetUserID(v string) *ListVolumesRequest
	GetUserID() *string
	SetVolumeName(v string) *ListVolumesRequest
	GetVolumeName() *string
}

type ListVolumesRequest struct {
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 5
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Leave this parameter empty for the first request. To query the next page, set this parameter to the nextToken value returned in the previous response.
	//
	// example:
	//
	// qxGrXje86XMrYQ51aJMy
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmwxqyrgwabcd
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	// Filters by storage volume status. Valid values: CREATING (creating), AVAILABLE (available), ERROR (error), and DELETING (deleting).
	//
	// example:
	//
	// CREATING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Filters by storage class. Valid values: OSS, AGENTIC_BUCKET, AGENTIC_FS, and JUICE_FS.
	//
	// example:
	//
	// OSS
	StorageClass *string `json:"storageClass,omitempty" xml:"storageClass,omitempty"`
	// The ID of the team to which the storage volume belongs.
	//
	// example:
	//
	// 70d1c834-0383-58d8-97ac-5336eb91abcd
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The user ID of the creator.
	//
	// example:
	//
	// 210000000
	UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
	// The name of the storage volume, which is unique within the team.
	//
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

func (s *ListVolumesRequest) GetStorageClass() *string {
	return s.StorageClass
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

func (s *ListVolumesRequest) SetStorageClass(v string) *ListVolumesRequest {
	s.StorageClass = &v
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
