// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMergeRequestChangeTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetMergeRequestChangeTreeRequest
	GetAccessToken() *string
	SetFromPatchSetBizId(v string) *GetMergeRequestChangeTreeRequest
	GetFromPatchSetBizId() *string
	SetLocalId(v int64) *GetMergeRequestChangeTreeRequest
	GetLocalId() *int64
	SetOrganizationId(v string) *GetMergeRequestChangeTreeRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *GetMergeRequestChangeTreeRequest
	GetRepositoryIdentity() *string
	SetToPatchSetBizId(v string) *GetMergeRequestChangeTreeRequest
	GetToPatchSetBizId() *string
}

type GetMergeRequestChangeTreeRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
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
	// 7
	LocalId *int64 `json:"localId,omitempty" xml:"localId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
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

func (s GetMergeRequestChangeTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestChangeTreeRequest) GoString() string {
	return s.String()
}

func (s *GetMergeRequestChangeTreeRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetMergeRequestChangeTreeRequest) GetFromPatchSetBizId() *string {
	return s.FromPatchSetBizId
}

func (s *GetMergeRequestChangeTreeRequest) GetLocalId() *int64 {
	return s.LocalId
}

func (s *GetMergeRequestChangeTreeRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetMergeRequestChangeTreeRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *GetMergeRequestChangeTreeRequest) GetToPatchSetBizId() *string {
	return s.ToPatchSetBizId
}

func (s *GetMergeRequestChangeTreeRequest) SetAccessToken(v string) *GetMergeRequestChangeTreeRequest {
	s.AccessToken = &v
	return s
}

func (s *GetMergeRequestChangeTreeRequest) SetFromPatchSetBizId(v string) *GetMergeRequestChangeTreeRequest {
	s.FromPatchSetBizId = &v
	return s
}

func (s *GetMergeRequestChangeTreeRequest) SetLocalId(v int64) *GetMergeRequestChangeTreeRequest {
	s.LocalId = &v
	return s
}

func (s *GetMergeRequestChangeTreeRequest) SetOrganizationId(v string) *GetMergeRequestChangeTreeRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetMergeRequestChangeTreeRequest) SetRepositoryIdentity(v string) *GetMergeRequestChangeTreeRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *GetMergeRequestChangeTreeRequest) SetToPatchSetBizId(v string) *GetMergeRequestChangeTreeRequest {
	s.ToPatchSetBizId = &v
	return s
}

func (s *GetMergeRequestChangeTreeRequest) Validate() error {
	return dara.Validate(s)
}
