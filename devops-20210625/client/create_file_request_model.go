// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateFileRequest
	GetAccessToken() *string
	SetBranchName(v string) *CreateFileRequest
	GetBranchName() *string
	SetCommitMessage(v string) *CreateFileRequest
	GetCommitMessage() *string
	SetContent(v string) *CreateFileRequest
	GetContent() *string
	SetEncoding(v string) *CreateFileRequest
	GetEncoding() *string
	SetFilePath(v string) *CreateFileRequest
	GetFilePath() *string
	SetOrganizationId(v string) *CreateFileRequest
	GetOrganizationId() *string
}

type CreateFileRequest struct {
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
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// text
	Encoding *string `json:"encoding,omitempty" xml:"encoding,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /src/main/test.java
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileRequest) GoString() string {
	return s.String()
}

func (s *CreateFileRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateFileRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *CreateFileRequest) GetCommitMessage() *string {
	return s.CommitMessage
}

func (s *CreateFileRequest) GetContent() *string {
	return s.Content
}

func (s *CreateFileRequest) GetEncoding() *string {
	return s.Encoding
}

func (s *CreateFileRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateFileRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateFileRequest) SetAccessToken(v string) *CreateFileRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateFileRequest) SetBranchName(v string) *CreateFileRequest {
	s.BranchName = &v
	return s
}

func (s *CreateFileRequest) SetCommitMessage(v string) *CreateFileRequest {
	s.CommitMessage = &v
	return s
}

func (s *CreateFileRequest) SetContent(v string) *CreateFileRequest {
	s.Content = &v
	return s
}

func (s *CreateFileRequest) SetEncoding(v string) *CreateFileRequest {
	s.Encoding = &v
	return s
}

func (s *CreateFileRequest) SetFilePath(v string) *CreateFileRequest {
	s.FilePath = &v
	return s
}

func (s *CreateFileRequest) SetOrganizationId(v string) *CreateFileRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateFileRequest) Validate() error {
	return dara.Validate(s)
}
