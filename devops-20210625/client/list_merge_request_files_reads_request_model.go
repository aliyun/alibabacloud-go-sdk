// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestFilesReadsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListMergeRequestFilesReadsRequest
	GetAccessToken() *string
	SetFromPatchSetBizId(v string) *ListMergeRequestFilesReadsRequest
	GetFromPatchSetBizId() *string
	SetLocalId(v int64) *ListMergeRequestFilesReadsRequest
	GetLocalId() *int64
	SetOrganizationId(v string) *ListMergeRequestFilesReadsRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *ListMergeRequestFilesReadsRequest
	GetRepositoryIdentity() *string
	SetToPatchSetBizId(v string) *ListMergeRequestFilesReadsRequest
	GetToPatchSetBizId() *string
}

type ListMergeRequestFilesReadsRequest struct {
	// example:
	//
	// agp_4d57a6796b3626f52064ab1fba5384a5
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5e733626d53f4b04a6aa0e23d4ff72b8
	FromPatchSetBizId *string `json:"fromPatchSetBizId,omitempty" xml:"fromPatchSetBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	LocalId *int64 `json:"localId,omitempty" xml:"localId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 513fcfd81a9142d2bb0db4f72c0aa15b
	ToPatchSetBizId *string `json:"toPatchSetBizId,omitempty" xml:"toPatchSetBizId,omitempty"`
}

func (s ListMergeRequestFilesReadsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestFilesReadsRequest) GoString() string {
	return s.String()
}

func (s *ListMergeRequestFilesReadsRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListMergeRequestFilesReadsRequest) GetFromPatchSetBizId() *string {
	return s.FromPatchSetBizId
}

func (s *ListMergeRequestFilesReadsRequest) GetLocalId() *int64 {
	return s.LocalId
}

func (s *ListMergeRequestFilesReadsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListMergeRequestFilesReadsRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *ListMergeRequestFilesReadsRequest) GetToPatchSetBizId() *string {
	return s.ToPatchSetBizId
}

func (s *ListMergeRequestFilesReadsRequest) SetAccessToken(v string) *ListMergeRequestFilesReadsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListMergeRequestFilesReadsRequest) SetFromPatchSetBizId(v string) *ListMergeRequestFilesReadsRequest {
	s.FromPatchSetBizId = &v
	return s
}

func (s *ListMergeRequestFilesReadsRequest) SetLocalId(v int64) *ListMergeRequestFilesReadsRequest {
	s.LocalId = &v
	return s
}

func (s *ListMergeRequestFilesReadsRequest) SetOrganizationId(v string) *ListMergeRequestFilesReadsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListMergeRequestFilesReadsRequest) SetRepositoryIdentity(v string) *ListMergeRequestFilesReadsRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *ListMergeRequestFilesReadsRequest) SetToPatchSetBizId(v string) *ListMergeRequestFilesReadsRequest {
	s.ToPatchSetBizId = &v
	return s
}

func (s *ListMergeRequestFilesReadsRequest) Validate() error {
	return dara.Validate(s)
}
