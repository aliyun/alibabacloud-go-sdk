// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *UpdateFileRequest
	GetAccessToken() *string
	SetBranchName(v string) *UpdateFileRequest
	GetBranchName() *string
	SetCommitMessage(v string) *UpdateFileRequest
	GetCommitMessage() *string
	SetContent(v string) *UpdateFileRequest
	GetContent() *string
	SetEncoding(v string) *UpdateFileRequest
	GetEncoding() *string
	SetNewPath(v string) *UpdateFileRequest
	GetNewPath() *string
	SetOldPath(v string) *UpdateFileRequest
	GetOldPath() *string
	SetOrganizationId(v string) *UpdateFileRequest
	GetOrganizationId() *string
}

type UpdateFileRequest struct {
	// example:
	//
	// 0cf2c8458ac44d9481aab2dd6ec10596v3
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
	// src/main/update.txt
	NewPath *string `json:"newPath,omitempty" xml:"newPath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// src/main/update.txt
	OldPath *string `json:"oldPath,omitempty" xml:"oldPath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdateFileRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *UpdateFileRequest) GetCommitMessage() *string {
	return s.CommitMessage
}

func (s *UpdateFileRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateFileRequest) GetEncoding() *string {
	return s.Encoding
}

func (s *UpdateFileRequest) GetNewPath() *string {
	return s.NewPath
}

func (s *UpdateFileRequest) GetOldPath() *string {
	return s.OldPath
}

func (s *UpdateFileRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateFileRequest) SetAccessToken(v string) *UpdateFileRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateFileRequest) SetBranchName(v string) *UpdateFileRequest {
	s.BranchName = &v
	return s
}

func (s *UpdateFileRequest) SetCommitMessage(v string) *UpdateFileRequest {
	s.CommitMessage = &v
	return s
}

func (s *UpdateFileRequest) SetContent(v string) *UpdateFileRequest {
	s.Content = &v
	return s
}

func (s *UpdateFileRequest) SetEncoding(v string) *UpdateFileRequest {
	s.Encoding = &v
	return s
}

func (s *UpdateFileRequest) SetNewPath(v string) *UpdateFileRequest {
	s.NewPath = &v
	return s
}

func (s *UpdateFileRequest) SetOldPath(v string) *UpdateFileRequest {
	s.OldPath = &v
	return s
}

func (s *UpdateFileRequest) SetOrganizationId(v string) *UpdateFileRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateFileRequest) Validate() error {
	return dara.Validate(s)
}
