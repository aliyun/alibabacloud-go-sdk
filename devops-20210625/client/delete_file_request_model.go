// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DeleteFileRequest
	GetAccessToken() *string
	SetBranchName(v string) *DeleteFileRequest
	GetBranchName() *string
	SetCommitMessage(v string) *DeleteFileRequest
	GetCommitMessage() *string
	SetFilePath(v string) *DeleteFileRequest
	GetFilePath() *string
	SetOrganizationId(v string) *DeleteFileRequest
	GetOrganizationId() *string
}

type DeleteFileRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// master
	BranchName *string `json:"branchName,omitempty" xml:"branchName,omitempty"`
	// This parameter is required.
	CommitMessage *string `json:"commitMessage,omitempty" xml:"commitMessage,omitempty"`
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
}

func (s DeleteFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DeleteFileRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *DeleteFileRequest) GetCommitMessage() *string {
	return s.CommitMessage
}

func (s *DeleteFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *DeleteFileRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteFileRequest) SetAccessToken(v string) *DeleteFileRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteFileRequest) SetBranchName(v string) *DeleteFileRequest {
	s.BranchName = &v
	return s
}

func (s *DeleteFileRequest) SetCommitMessage(v string) *DeleteFileRequest {
	s.CommitMessage = &v
	return s
}

func (s *DeleteFileRequest) SetFilePath(v string) *DeleteFileRequest {
	s.FilePath = &v
	return s
}

func (s *DeleteFileRequest) SetOrganizationId(v string) *DeleteFileRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteFileRequest) Validate() error {
	return dara.Validate(s)
}
