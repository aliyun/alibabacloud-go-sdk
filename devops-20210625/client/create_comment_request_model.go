// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateCommentRequest
	GetAccessToken() *string
	SetCommentType(v string) *CreateCommentRequest
	GetCommentType() *string
	SetContent(v string) *CreateCommentRequest
	GetContent() *string
	SetDraft(v bool) *CreateCommentRequest
	GetDraft() *bool
	SetFilePath(v string) *CreateCommentRequest
	GetFilePath() *string
	SetFromPachSetBizId(v string) *CreateCommentRequest
	GetFromPachSetBizId() *string
	SetLineNumber(v int32) *CreateCommentRequest
	GetLineNumber() *int32
	SetParentCommentBizId(v string) *CreateCommentRequest
	GetParentCommentBizId() *string
	SetPatchSetBizId(v string) *CreateCommentRequest
	GetPatchSetBizId() *string
	SetResolved(v bool) *CreateCommentRequest
	GetResolved() *bool
	SetToPatchSetBizId(v string) *CreateCommentRequest
	GetToPatchSetBizId() *string
	SetLocalId(v int64) *CreateCommentRequest
	GetLocalId() *int64
	SetOrganizationId(v string) *CreateCommentRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *CreateCommentRequest
	GetRepositoryIdentity() *string
}

type CreateCommentRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GLOBAL_COMMENT
	CommentType *string `json:"commentType,omitempty" xml:"commentType,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// false
	Draft *bool `json:"draft,omitempty" xml:"draft,omitempty"`
	// example:
	//
	// /src/main/test.java
	FilePath         *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	FromPachSetBizId *string `json:"fromPachSetBizId,omitempty" xml:"fromPachSetBizId,omitempty"`
	// example:
	//
	// 1
	LineNumber *int32 `json:"lineNumber,omitempty" xml:"lineNumber,omitempty"`
	// example:
	//
	// 2666ac1ac53841b0ba1b042e383279cc
	ParentCommentBizId *string `json:"parentCommentBizId,omitempty" xml:"parentCommentBizId,omitempty"`
	// example:
	//
	// b7d8386be17c4ca68a07140db4836257
	PatchSetBizId *string `json:"patchSetBizId,omitempty" xml:"patchSetBizId,omitempty"`
	// example:
	//
	// false
	Resolved        *bool   `json:"resolved,omitempty" xml:"resolved,omitempty"`
	ToPatchSetBizId *string `json:"toPatchSetBizId,omitempty" xml:"toPatchSetBizId,omitempty"`
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
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
}

func (s CreateCommentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCommentRequest) GoString() string {
	return s.String()
}

func (s *CreateCommentRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateCommentRequest) GetCommentType() *string {
	return s.CommentType
}

func (s *CreateCommentRequest) GetContent() *string {
	return s.Content
}

func (s *CreateCommentRequest) GetDraft() *bool {
	return s.Draft
}

func (s *CreateCommentRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateCommentRequest) GetFromPachSetBizId() *string {
	return s.FromPachSetBizId
}

func (s *CreateCommentRequest) GetLineNumber() *int32 {
	return s.LineNumber
}

func (s *CreateCommentRequest) GetParentCommentBizId() *string {
	return s.ParentCommentBizId
}

func (s *CreateCommentRequest) GetPatchSetBizId() *string {
	return s.PatchSetBizId
}

func (s *CreateCommentRequest) GetResolved() *bool {
	return s.Resolved
}

func (s *CreateCommentRequest) GetToPatchSetBizId() *string {
	return s.ToPatchSetBizId
}

func (s *CreateCommentRequest) GetLocalId() *int64 {
	return s.LocalId
}

func (s *CreateCommentRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateCommentRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *CreateCommentRequest) SetAccessToken(v string) *CreateCommentRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateCommentRequest) SetCommentType(v string) *CreateCommentRequest {
	s.CommentType = &v
	return s
}

func (s *CreateCommentRequest) SetContent(v string) *CreateCommentRequest {
	s.Content = &v
	return s
}

func (s *CreateCommentRequest) SetDraft(v bool) *CreateCommentRequest {
	s.Draft = &v
	return s
}

func (s *CreateCommentRequest) SetFilePath(v string) *CreateCommentRequest {
	s.FilePath = &v
	return s
}

func (s *CreateCommentRequest) SetFromPachSetBizId(v string) *CreateCommentRequest {
	s.FromPachSetBizId = &v
	return s
}

func (s *CreateCommentRequest) SetLineNumber(v int32) *CreateCommentRequest {
	s.LineNumber = &v
	return s
}

func (s *CreateCommentRequest) SetParentCommentBizId(v string) *CreateCommentRequest {
	s.ParentCommentBizId = &v
	return s
}

func (s *CreateCommentRequest) SetPatchSetBizId(v string) *CreateCommentRequest {
	s.PatchSetBizId = &v
	return s
}

func (s *CreateCommentRequest) SetResolved(v bool) *CreateCommentRequest {
	s.Resolved = &v
	return s
}

func (s *CreateCommentRequest) SetToPatchSetBizId(v string) *CreateCommentRequest {
	s.ToPatchSetBizId = &v
	return s
}

func (s *CreateCommentRequest) SetLocalId(v int64) *CreateCommentRequest {
	s.LocalId = &v
	return s
}

func (s *CreateCommentRequest) SetOrganizationId(v string) *CreateCommentRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateCommentRequest) SetRepositoryIdentity(v string) *CreateCommentRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *CreateCommentRequest) Validate() error {
	return dara.Validate(s)
}
