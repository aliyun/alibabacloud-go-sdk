// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateRepositoryRequest
	GetAccessToken() *string
	SetAvatarUrl(v string) *CreateRepositoryRequest
	GetAvatarUrl() *string
	SetDescription(v string) *CreateRepositoryRequest
	GetDescription() *string
	SetGitignoreType(v string) *CreateRepositoryRequest
	GetGitignoreType() *string
	SetImportAccount(v string) *CreateRepositoryRequest
	GetImportAccount() *string
	SetImportDemoProject(v bool) *CreateRepositoryRequest
	GetImportDemoProject() *bool
	SetImportRepoType(v string) *CreateRepositoryRequest
	GetImportRepoType() *string
	SetImportToken(v string) *CreateRepositoryRequest
	GetImportToken() *string
	SetImportTokenEncrypted(v string) *CreateRepositoryRequest
	GetImportTokenEncrypted() *string
	SetImportUrl(v string) *CreateRepositoryRequest
	GetImportUrl() *string
	SetInitStandardService(v bool) *CreateRepositoryRequest
	GetInitStandardService() *bool
	SetIsCryptoEnabled(v bool) *CreateRepositoryRequest
	GetIsCryptoEnabled() *bool
	SetLocalImportUrl(v string) *CreateRepositoryRequest
	GetLocalImportUrl() *string
	SetName(v string) *CreateRepositoryRequest
	GetName() *string
	SetNamespaceId(v int64) *CreateRepositoryRequest
	GetNamespaceId() *int64
	SetPath(v string) *CreateRepositoryRequest
	GetPath() *string
	SetReadmeType(v string) *CreateRepositoryRequest
	GetReadmeType() *string
	SetVisibilityLevel(v int32) *CreateRepositoryRequest
	GetVisibilityLevel() *int32
	SetCreateParentPath(v bool) *CreateRepositoryRequest
	GetCreateParentPath() *bool
	SetOrganizationId(v string) *CreateRepositoryRequest
	GetOrganizationId() *string
	SetSync(v bool) *CreateRepositoryRequest
	GetSync() *bool
}

type CreateRepositoryRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl   *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// Java
	GitignoreType *string `json:"gitignoreType,omitempty" xml:"gitignoreType,omitempty"`
	// example:
	//
	// codeup-test
	ImportAccount *string `json:"importAccount,omitempty" xml:"importAccount,omitempty"`
	// example:
	//
	// true
	ImportDemoProject *bool `json:"importDemoProject,omitempty" xml:"importDemoProject,omitempty"`
	// example:
	//
	// GIT
	ImportRepoType *string `json:"importRepoType,omitempty" xml:"importRepoType,omitempty"`
	// example:
	//
	// xxxxx
	ImportToken *string `json:"importToken,omitempty" xml:"importToken,omitempty"`
	// example:
	//
	// text
	ImportTokenEncrypted *string `json:"importTokenEncrypted,omitempty" xml:"importTokenEncrypted,omitempty"`
	// example:
	//
	// https://github.com/a/b.git
	ImportUrl *string `json:"importUrl,omitempty" xml:"importUrl,omitempty"`
	// example:
	//
	// true
	InitStandardService *bool `json:"initStandardService,omitempty" xml:"initStandardService,omitempty"`
	// example:
	//
	// false
	IsCryptoEnabled *bool `json:"isCryptoEnabled,omitempty" xml:"isCryptoEnabled,omitempty"`
	// example:
	//
	// ""
	LocalImportUrl *string `json:"localImportUrl,omitempty" xml:"localImportUrl,omitempty"`
	// This parameter is required.
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	NamespaceId *int64  `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	Path        *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// USER_GUIDE
	ReadmeType *string `json:"readmeType,omitempty" xml:"readmeType,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	// example:
	//
	// true
	CreateParentPath *bool `json:"createParentPath,omitempty" xml:"createParentPath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// false
	Sync *bool `json:"sync,omitempty" xml:"sync,omitempty"`
}

func (s CreateRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRepositoryRequest) GoString() string {
	return s.String()
}

func (s *CreateRepositoryRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateRepositoryRequest) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateRepositoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRepositoryRequest) GetGitignoreType() *string {
	return s.GitignoreType
}

func (s *CreateRepositoryRequest) GetImportAccount() *string {
	return s.ImportAccount
}

func (s *CreateRepositoryRequest) GetImportDemoProject() *bool {
	return s.ImportDemoProject
}

func (s *CreateRepositoryRequest) GetImportRepoType() *string {
	return s.ImportRepoType
}

func (s *CreateRepositoryRequest) GetImportToken() *string {
	return s.ImportToken
}

func (s *CreateRepositoryRequest) GetImportTokenEncrypted() *string {
	return s.ImportTokenEncrypted
}

func (s *CreateRepositoryRequest) GetImportUrl() *string {
	return s.ImportUrl
}

func (s *CreateRepositoryRequest) GetInitStandardService() *bool {
	return s.InitStandardService
}

func (s *CreateRepositoryRequest) GetIsCryptoEnabled() *bool {
	return s.IsCryptoEnabled
}

func (s *CreateRepositoryRequest) GetLocalImportUrl() *string {
	return s.LocalImportUrl
}

func (s *CreateRepositoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateRepositoryRequest) GetNamespaceId() *int64 {
	return s.NamespaceId
}

func (s *CreateRepositoryRequest) GetPath() *string {
	return s.Path
}

func (s *CreateRepositoryRequest) GetReadmeType() *string {
	return s.ReadmeType
}

func (s *CreateRepositoryRequest) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *CreateRepositoryRequest) GetCreateParentPath() *bool {
	return s.CreateParentPath
}

func (s *CreateRepositoryRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateRepositoryRequest) GetSync() *bool {
	return s.Sync
}

func (s *CreateRepositoryRequest) SetAccessToken(v string) *CreateRepositoryRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateRepositoryRequest) SetAvatarUrl(v string) *CreateRepositoryRequest {
	s.AvatarUrl = &v
	return s
}

func (s *CreateRepositoryRequest) SetDescription(v string) *CreateRepositoryRequest {
	s.Description = &v
	return s
}

func (s *CreateRepositoryRequest) SetGitignoreType(v string) *CreateRepositoryRequest {
	s.GitignoreType = &v
	return s
}

func (s *CreateRepositoryRequest) SetImportAccount(v string) *CreateRepositoryRequest {
	s.ImportAccount = &v
	return s
}

func (s *CreateRepositoryRequest) SetImportDemoProject(v bool) *CreateRepositoryRequest {
	s.ImportDemoProject = &v
	return s
}

func (s *CreateRepositoryRequest) SetImportRepoType(v string) *CreateRepositoryRequest {
	s.ImportRepoType = &v
	return s
}

func (s *CreateRepositoryRequest) SetImportToken(v string) *CreateRepositoryRequest {
	s.ImportToken = &v
	return s
}

func (s *CreateRepositoryRequest) SetImportTokenEncrypted(v string) *CreateRepositoryRequest {
	s.ImportTokenEncrypted = &v
	return s
}

func (s *CreateRepositoryRequest) SetImportUrl(v string) *CreateRepositoryRequest {
	s.ImportUrl = &v
	return s
}

func (s *CreateRepositoryRequest) SetInitStandardService(v bool) *CreateRepositoryRequest {
	s.InitStandardService = &v
	return s
}

func (s *CreateRepositoryRequest) SetIsCryptoEnabled(v bool) *CreateRepositoryRequest {
	s.IsCryptoEnabled = &v
	return s
}

func (s *CreateRepositoryRequest) SetLocalImportUrl(v string) *CreateRepositoryRequest {
	s.LocalImportUrl = &v
	return s
}

func (s *CreateRepositoryRequest) SetName(v string) *CreateRepositoryRequest {
	s.Name = &v
	return s
}

func (s *CreateRepositoryRequest) SetNamespaceId(v int64) *CreateRepositoryRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateRepositoryRequest) SetPath(v string) *CreateRepositoryRequest {
	s.Path = &v
	return s
}

func (s *CreateRepositoryRequest) SetReadmeType(v string) *CreateRepositoryRequest {
	s.ReadmeType = &v
	return s
}

func (s *CreateRepositoryRequest) SetVisibilityLevel(v int32) *CreateRepositoryRequest {
	s.VisibilityLevel = &v
	return s
}

func (s *CreateRepositoryRequest) SetCreateParentPath(v bool) *CreateRepositoryRequest {
	s.CreateParentPath = &v
	return s
}

func (s *CreateRepositoryRequest) SetOrganizationId(v string) *CreateRepositoryRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateRepositoryRequest) SetSync(v bool) *CreateRepositoryRequest {
	s.Sync = &v
	return s
}

func (s *CreateRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
