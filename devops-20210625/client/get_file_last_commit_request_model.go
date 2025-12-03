// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileLastCommitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetFileLastCommitRequest
	GetAccessToken() *string
	SetFilePath(v string) *GetFileLastCommitRequest
	GetFilePath() *string
	SetOrganizationId(v string) *GetFileLastCommitRequest
	GetOrganizationId() *string
	SetSha(v string) *GetFileLastCommitRequest
	GetSha() *string
	SetShowSignature(v bool) *GetFileLastCommitRequest
	GetShowSignature() *bool
}

type GetFileLastCommitRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// src/main/test.java
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// master
	Sha *string `json:"sha,omitempty" xml:"sha,omitempty"`
	// example:
	//
	// false
	ShowSignature *bool `json:"showSignature,omitempty" xml:"showSignature,omitempty"`
}

func (s GetFileLastCommitRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileLastCommitRequest) GoString() string {
	return s.String()
}

func (s *GetFileLastCommitRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetFileLastCommitRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *GetFileLastCommitRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetFileLastCommitRequest) GetSha() *string {
	return s.Sha
}

func (s *GetFileLastCommitRequest) GetShowSignature() *bool {
	return s.ShowSignature
}

func (s *GetFileLastCommitRequest) SetAccessToken(v string) *GetFileLastCommitRequest {
	s.AccessToken = &v
	return s
}

func (s *GetFileLastCommitRequest) SetFilePath(v string) *GetFileLastCommitRequest {
	s.FilePath = &v
	return s
}

func (s *GetFileLastCommitRequest) SetOrganizationId(v string) *GetFileLastCommitRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetFileLastCommitRequest) SetSha(v string) *GetFileLastCommitRequest {
	s.Sha = &v
	return s
}

func (s *GetFileLastCommitRequest) SetShowSignature(v bool) *GetFileLastCommitRequest {
	s.ShowSignature = &v
	return s
}

func (s *GetFileLastCommitRequest) Validate() error {
	return dara.Validate(s)
}
