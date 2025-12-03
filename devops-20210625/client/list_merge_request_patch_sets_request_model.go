// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestPatchSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListMergeRequestPatchSetsRequest
	GetAccessToken() *string
	SetLocalId(v int64) *ListMergeRequestPatchSetsRequest
	GetLocalId() *int64
	SetOrganizationId(v string) *ListMergeRequestPatchSetsRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *ListMergeRequestPatchSetsRequest
	GetRepositoryIdentity() *string
}

type ListMergeRequestPatchSetsRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	LocalId *int64 `json:"localId,omitempty" xml:"localId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
}

func (s ListMergeRequestPatchSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestPatchSetsRequest) GoString() string {
	return s.String()
}

func (s *ListMergeRequestPatchSetsRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListMergeRequestPatchSetsRequest) GetLocalId() *int64 {
	return s.LocalId
}

func (s *ListMergeRequestPatchSetsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListMergeRequestPatchSetsRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *ListMergeRequestPatchSetsRequest) SetAccessToken(v string) *ListMergeRequestPatchSetsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListMergeRequestPatchSetsRequest) SetLocalId(v int64) *ListMergeRequestPatchSetsRequest {
	s.LocalId = &v
	return s
}

func (s *ListMergeRequestPatchSetsRequest) SetOrganizationId(v string) *ListMergeRequestPatchSetsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListMergeRequestPatchSetsRequest) SetRepositoryIdentity(v string) *ListMergeRequestPatchSetsRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *ListMergeRequestPatchSetsRequest) Validate() error {
	return dara.Validate(s)
}
