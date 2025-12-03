// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMergeRequestCommentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListMergeRequestCommentsRequest
	GetAccessToken() *string
	SetCommentType(v string) *ListMergeRequestCommentsRequest
	GetCommentType() *string
	SetFilePath(v string) *ListMergeRequestCommentsRequest
	GetFilePath() *string
	SetPatchSetBizIds(v []*string) *ListMergeRequestCommentsRequest
	GetPatchSetBizIds() []*string
	SetResolved(v bool) *ListMergeRequestCommentsRequest
	GetResolved() *bool
	SetState(v string) *ListMergeRequestCommentsRequest
	GetState() *string
	SetLocalId(v int64) *ListMergeRequestCommentsRequest
	GetLocalId() *int64
	SetOrganizationId(v string) *ListMergeRequestCommentsRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *ListMergeRequestCommentsRequest
	GetRepositoryIdentity() *string
}

type ListMergeRequestCommentsRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// GLOBAL_COMMENT
	CommentType *string `json:"commentType,omitempty" xml:"commentType,omitempty"`
	// example:
	//
	// /src/main/test.java
	FilePath       *string   `json:"filePath,omitempty" xml:"filePath,omitempty"`
	PatchSetBizIds []*string `json:"patchSetBizIds,omitempty" xml:"patchSetBizIds,omitempty" type:"Repeated"`
	// example:
	//
	// false
	Resolved *bool `json:"resolved,omitempty" xml:"resolved,omitempty"`
	// example:
	//
	// OPENED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
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

func (s ListMergeRequestCommentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMergeRequestCommentsRequest) GoString() string {
	return s.String()
}

func (s *ListMergeRequestCommentsRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListMergeRequestCommentsRequest) GetCommentType() *string {
	return s.CommentType
}

func (s *ListMergeRequestCommentsRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *ListMergeRequestCommentsRequest) GetPatchSetBizIds() []*string {
	return s.PatchSetBizIds
}

func (s *ListMergeRequestCommentsRequest) GetResolved() *bool {
	return s.Resolved
}

func (s *ListMergeRequestCommentsRequest) GetState() *string {
	return s.State
}

func (s *ListMergeRequestCommentsRequest) GetLocalId() *int64 {
	return s.LocalId
}

func (s *ListMergeRequestCommentsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListMergeRequestCommentsRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *ListMergeRequestCommentsRequest) SetAccessToken(v string) *ListMergeRequestCommentsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListMergeRequestCommentsRequest) SetCommentType(v string) *ListMergeRequestCommentsRequest {
	s.CommentType = &v
	return s
}

func (s *ListMergeRequestCommentsRequest) SetFilePath(v string) *ListMergeRequestCommentsRequest {
	s.FilePath = &v
	return s
}

func (s *ListMergeRequestCommentsRequest) SetPatchSetBizIds(v []*string) *ListMergeRequestCommentsRequest {
	s.PatchSetBizIds = v
	return s
}

func (s *ListMergeRequestCommentsRequest) SetResolved(v bool) *ListMergeRequestCommentsRequest {
	s.Resolved = &v
	return s
}

func (s *ListMergeRequestCommentsRequest) SetState(v string) *ListMergeRequestCommentsRequest {
	s.State = &v
	return s
}

func (s *ListMergeRequestCommentsRequest) SetLocalId(v int64) *ListMergeRequestCommentsRequest {
	s.LocalId = &v
	return s
}

func (s *ListMergeRequestCommentsRequest) SetOrganizationId(v string) *ListMergeRequestCommentsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListMergeRequestCommentsRequest) SetRepositoryIdentity(v string) *ListMergeRequestCommentsRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *ListMergeRequestCommentsRequest) Validate() error {
	return dara.Validate(s)
}
