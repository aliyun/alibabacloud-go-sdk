// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommitWithMultipleFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateCommitWithMultipleFilesRequest
	GetAccessToken() *string
	SetActions(v []*CreateCommitWithMultipleFilesRequestActions) *CreateCommitWithMultipleFilesRequest
	GetActions() []*CreateCommitWithMultipleFilesRequestActions
	SetBranch(v string) *CreateCommitWithMultipleFilesRequest
	GetBranch() *string
	SetCommitMessage(v string) *CreateCommitWithMultipleFilesRequest
	GetCommitMessage() *string
	SetOrganizationId(v string) *CreateCommitWithMultipleFilesRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *CreateCommitWithMultipleFilesRequest
	GetRepositoryIdentity() *string
}

type CreateCommitWithMultipleFilesRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	Actions []*CreateCommitWithMultipleFilesRequestActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// master
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// auto
	CommitMessage *string `json:"commitMessage,omitempty" xml:"commitMessage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
}

func (s CreateCommitWithMultipleFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCommitWithMultipleFilesRequest) GoString() string {
	return s.String()
}

func (s *CreateCommitWithMultipleFilesRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateCommitWithMultipleFilesRequest) GetActions() []*CreateCommitWithMultipleFilesRequestActions {
	return s.Actions
}

func (s *CreateCommitWithMultipleFilesRequest) GetBranch() *string {
	return s.Branch
}

func (s *CreateCommitWithMultipleFilesRequest) GetCommitMessage() *string {
	return s.CommitMessage
}

func (s *CreateCommitWithMultipleFilesRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateCommitWithMultipleFilesRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *CreateCommitWithMultipleFilesRequest) SetAccessToken(v string) *CreateCommitWithMultipleFilesRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateCommitWithMultipleFilesRequest) SetActions(v []*CreateCommitWithMultipleFilesRequestActions) *CreateCommitWithMultipleFilesRequest {
	s.Actions = v
	return s
}

func (s *CreateCommitWithMultipleFilesRequest) SetBranch(v string) *CreateCommitWithMultipleFilesRequest {
	s.Branch = &v
	return s
}

func (s *CreateCommitWithMultipleFilesRequest) SetCommitMessage(v string) *CreateCommitWithMultipleFilesRequest {
	s.CommitMessage = &v
	return s
}

func (s *CreateCommitWithMultipleFilesRequest) SetOrganizationId(v string) *CreateCommitWithMultipleFilesRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateCommitWithMultipleFilesRequest) SetRepositoryIdentity(v string) *CreateCommitWithMultipleFilesRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *CreateCommitWithMultipleFilesRequest) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCommitWithMultipleFilesRequestActions struct {
	// example:
	//
	// create
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// xxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// src/test.java
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// example:
	//
	// src/main/test.java
	PreviousPath *string `json:"previousPath,omitempty" xml:"previousPath,omitempty"`
}

func (s CreateCommitWithMultipleFilesRequestActions) String() string {
	return dara.Prettify(s)
}

func (s CreateCommitWithMultipleFilesRequestActions) GoString() string {
	return s.String()
}

func (s *CreateCommitWithMultipleFilesRequestActions) GetAction() *string {
	return s.Action
}

func (s *CreateCommitWithMultipleFilesRequestActions) GetContent() *string {
	return s.Content
}

func (s *CreateCommitWithMultipleFilesRequestActions) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateCommitWithMultipleFilesRequestActions) GetPreviousPath() *string {
	return s.PreviousPath
}

func (s *CreateCommitWithMultipleFilesRequestActions) SetAction(v string) *CreateCommitWithMultipleFilesRequestActions {
	s.Action = &v
	return s
}

func (s *CreateCommitWithMultipleFilesRequestActions) SetContent(v string) *CreateCommitWithMultipleFilesRequestActions {
	s.Content = &v
	return s
}

func (s *CreateCommitWithMultipleFilesRequestActions) SetFilePath(v string) *CreateCommitWithMultipleFilesRequestActions {
	s.FilePath = &v
	return s
}

func (s *CreateCommitWithMultipleFilesRequestActions) SetPreviousPath(v string) *CreateCommitWithMultipleFilesRequestActions {
	s.PreviousPath = &v
	return s
}

func (s *CreateCommitWithMultipleFilesRequestActions) Validate() error {
	return dara.Validate(s)
}
