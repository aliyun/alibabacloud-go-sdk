// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListMergeRequestLabelsRequest
	GetAccessToken() *string
	SetLocalId(v int64) *ListMergeRequestLabelsRequest
	GetLocalId() *int64
	SetOrganizationId(v string) *ListMergeRequestLabelsRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *ListMergeRequestLabelsRequest
	GetRepositoryIdentity() *string
}

type ListMergeRequestLabelsRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	LocalId *int64 `json:"localId,omitempty" xml:"localId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
}

func (s ListMergeRequestLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestLabelsRequest) GoString() string {
	return s.String()
}

func (s *ListMergeRequestLabelsRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListMergeRequestLabelsRequest) GetLocalId() *int64 {
	return s.LocalId
}

func (s *ListMergeRequestLabelsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListMergeRequestLabelsRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *ListMergeRequestLabelsRequest) SetAccessToken(v string) *ListMergeRequestLabelsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListMergeRequestLabelsRequest) SetLocalId(v int64) *ListMergeRequestLabelsRequest {
	s.LocalId = &v
	return s
}

func (s *ListMergeRequestLabelsRequest) SetOrganizationId(v string) *ListMergeRequestLabelsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListMergeRequestLabelsRequest) SetRepositoryIdentity(v string) *ListMergeRequestLabelsRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *ListMergeRequestLabelsRequest) Validate() error {
	return dara.Validate(s)
}
