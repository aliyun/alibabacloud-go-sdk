// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateRepositoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateRepositoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateRepositoryResponseBody
	GetRequestId() *string
	SetResult(v *CreateRepositoryResponseBodyResult) *CreateRepositoryResponseBody
	GetResult() *CreateRepositoryResponseBodyResult
	SetSuccess(v bool) *CreateRepositoryResponseBody
	GetSuccess() *bool
}

type CreateRepositoryResponseBody struct {
	// example:
	//
	// 401
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// SYSTEM_UNAUTHORIZED_ERROR
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// F590C9D8-E908-5B6C-95AC-56B7E8011FFA
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateRepositoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateRepositoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRepositoryResponseBody) GetResult() *CreateRepositoryResponseBodyResult {
	return s.Result
}

func (s *CreateRepositoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRepositoryResponseBody) SetErrorCode(v string) *CreateRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetErrorMessage(v string) *CreateRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetRequestId(v string) *CreateRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetResult(v *CreateRepositoryResponseBodyResult) *CreateRepositoryResponseBody {
	s.Result = v
	return s
}

func (s *CreateRepositoryResponseBody) SetSuccess(v bool) *CreateRepositoryResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRepositoryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRepositoryResponseBodyResult struct {
	// example:
	//
	// false
	ImportFromSvn *bool `json:"Import_from_svn,omitempty" xml:"Import_from_svn,omitempty"`
	// example:
	//
	// false
	Archived *bool `json:"archived,omitempty" xml:"archived,omitempty"`
	// example:
	//
	// ""
	AvatarUrl *string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 1233
	CreatorId *int64 `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// master
	DefaultBranch *string `json:"defaultBranch,omitempty" xml:"defaultBranch,omitempty"`
	// example:
	//
	// false
	DemoProject *bool   `json:"demoProject,omitempty" xml:"demoProject,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// ""
	HttpUrlToRepo *string `json:"httpUrlToRepo,omitempty" xml:"httpUrlToRepo,omitempty"`
	// id
	//
	// example:
	//
	// 2959
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	LastActivityAt *string `json:"lastActivityAt,omitempty" xml:"lastActivityAt,omitempty"`
	// example:
	//
	// codeupTest
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// codeup-test-org / test-codeup
	NameWithNamespace *string                                      `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	Namespace         *CreateRepositoryResponseBodyResultNamespace `json:"namespace,omitempty" xml:"namespace,omitempty" type:"Struct"`
	// example:
	//
	// test-codeup
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// codeup-test-org/test-codeup
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	// example:
	//
	// ""
	SshUrlToRepo *string `json:"sshUrlToRepo,omitempty" xml:"sshUrlToRepo,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *string `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	// web url
	//
	// example:
	//
	// ""
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s CreateRepositoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateRepositoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBodyResult) GetImportFromSvn() *bool {
	return s.ImportFromSvn
}

func (s *CreateRepositoryResponseBodyResult) GetArchived() *bool {
	return s.Archived
}

func (s *CreateRepositoryResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateRepositoryResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateRepositoryResponseBodyResult) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *CreateRepositoryResponseBodyResult) GetDefaultBranch() *string {
	return s.DefaultBranch
}

func (s *CreateRepositoryResponseBodyResult) GetDemoProject() *bool {
	return s.DemoProject
}

func (s *CreateRepositoryResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *CreateRepositoryResponseBodyResult) GetHttpUrlToRepo() *string {
	return s.HttpUrlToRepo
}

func (s *CreateRepositoryResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CreateRepositoryResponseBodyResult) GetLastActivityAt() *string {
	return s.LastActivityAt
}

func (s *CreateRepositoryResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateRepositoryResponseBodyResult) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *CreateRepositoryResponseBodyResult) GetNamespace() *CreateRepositoryResponseBodyResultNamespace {
	return s.Namespace
}

func (s *CreateRepositoryResponseBodyResult) GetPath() *string {
	return s.Path
}

func (s *CreateRepositoryResponseBodyResult) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *CreateRepositoryResponseBodyResult) GetSshUrlToRepo() *string {
	return s.SshUrlToRepo
}

func (s *CreateRepositoryResponseBodyResult) GetVisibilityLevel() *string {
	return s.VisibilityLevel
}

func (s *CreateRepositoryResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *CreateRepositoryResponseBodyResult) SetImportFromSvn(v bool) *CreateRepositoryResponseBodyResult {
	s.ImportFromSvn = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetArchived(v bool) *CreateRepositoryResponseBodyResult {
	s.Archived = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetAvatarUrl(v string) *CreateRepositoryResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetCreatedAt(v string) *CreateRepositoryResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetCreatorId(v int64) *CreateRepositoryResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetDefaultBranch(v string) *CreateRepositoryResponseBodyResult {
	s.DefaultBranch = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetDemoProject(v bool) *CreateRepositoryResponseBodyResult {
	s.DemoProject = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetDescription(v string) *CreateRepositoryResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetHttpUrlToRepo(v string) *CreateRepositoryResponseBodyResult {
	s.HttpUrlToRepo = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetId(v int64) *CreateRepositoryResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetLastActivityAt(v string) *CreateRepositoryResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetName(v string) *CreateRepositoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetNameWithNamespace(v string) *CreateRepositoryResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetNamespace(v *CreateRepositoryResponseBodyResultNamespace) *CreateRepositoryResponseBodyResult {
	s.Namespace = v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetPath(v string) *CreateRepositoryResponseBodyResult {
	s.Path = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetPathWithNamespace(v string) *CreateRepositoryResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetSshUrlToRepo(v string) *CreateRepositoryResponseBodyResult {
	s.SshUrlToRepo = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetVisibilityLevel(v string) *CreateRepositoryResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) SetWebUrl(v string) *CreateRepositoryResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *CreateRepositoryResponseBodyResult) Validate() error {
	if s.Namespace != nil {
		if err := s.Namespace.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRepositoryResponseBodyResultNamespace struct {
	// example:
	//
	// ""
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// codeup repo description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// id
	//
	// example:
	//
	// 3194
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123
	OwnerId *int64 `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// example:
	//
	// codeup
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// true
	Public *bool `json:"public,omitempty" xml:"public,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *string `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
}

func (s CreateRepositoryResponseBodyResultNamespace) String() string {
	return dara.Prettify(s)
}

func (s CreateRepositoryResponseBodyResultNamespace) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBodyResultNamespace) GetAvatar() *string {
	return s.Avatar
}

func (s *CreateRepositoryResponseBodyResultNamespace) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateRepositoryResponseBodyResultNamespace) GetDescription() *string {
	return s.Description
}

func (s *CreateRepositoryResponseBodyResultNamespace) GetId() *int64 {
	return s.Id
}

func (s *CreateRepositoryResponseBodyResultNamespace) GetName() *string {
	return s.Name
}

func (s *CreateRepositoryResponseBodyResultNamespace) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateRepositoryResponseBodyResultNamespace) GetPath() *string {
	return s.Path
}

func (s *CreateRepositoryResponseBodyResultNamespace) GetPublic() *bool {
	return s.Public
}

func (s *CreateRepositoryResponseBodyResultNamespace) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateRepositoryResponseBodyResultNamespace) GetVisibilityLevel() *string {
	return s.VisibilityLevel
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetAvatar(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Avatar = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetCreatedAt(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.CreatedAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetDescription(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Description = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetId(v int64) *CreateRepositoryResponseBodyResultNamespace {
	s.Id = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetName(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Name = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetOwnerId(v int64) *CreateRepositoryResponseBodyResultNamespace {
	s.OwnerId = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetPath(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.Path = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetPublic(v bool) *CreateRepositoryResponseBodyResultNamespace {
	s.Public = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetUpdatedAt(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.UpdatedAt = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) SetVisibilityLevel(v string) *CreateRepositoryResponseBodyResultNamespace {
	s.VisibilityLevel = &v
	return s
}

func (s *CreateRepositoryResponseBodyResultNamespace) Validate() error {
	return dara.Validate(s)
}
