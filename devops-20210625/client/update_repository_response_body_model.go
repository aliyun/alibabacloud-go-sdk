// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateRepositoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateRepositoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateRepositoryResponseBody
	GetRequestId() *string
	SetResult(v *UpdateRepositoryResponseBodyResult) *UpdateRepositoryResponseBody
	GetResult() *UpdateRepositoryResponseBodyResult
	SetSuccess(v bool) *UpdateRepositoryResponseBody
	GetSuccess() *bool
}

type UpdateRepositoryResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// “”
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// EAE03103-5497-58D1-9169-E524DDE8604C
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *UpdateRepositoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateRepositoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRepositoryResponseBody) GetResult() *UpdateRepositoryResponseBodyResult {
	return s.Result
}

func (s *UpdateRepositoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRepositoryResponseBody) SetErrorCode(v string) *UpdateRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetErrorMessage(v string) *UpdateRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetRequestId(v string) *UpdateRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetResult(v *UpdateRepositoryResponseBodyResult) *UpdateRepositoryResponseBody {
	s.Result = v
	return s
}

func (s *UpdateRepositoryResponseBody) SetSuccess(v bool) *UpdateRepositoryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRepositoryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRepositoryResponseBodyResult struct {
	// example:
	//
	// false
	Archived *bool `json:"archived,omitempty" xml:"archived,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// false
	BuildsEnabled *bool `json:"buildsEnabled,omitempty" xml:"buildsEnabled,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 19238
	CreatorId *int64 `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// master
	DefaultBranch *string `json:"defaultBranch,omitempty" xml:"defaultBranch,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// https://codeup.aliyun.com/xxx/test/test.git
	HttpUrlToRepo *string `json:"httpUrlToRepo,omitempty" xml:"httpUrlToRepo,omitempty"`
	// example:
	//
	// 2825387
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IssuesEnabled *bool `json:"issuesEnabled,omitempty" xml:"issuesEnabled,omitempty"`
	// example:
	//
	// 2022-03-20 14:24:54
	LastActivityAt *string `json:"lastActivityAt,omitempty" xml:"lastActivityAt,omitempty"`
	// example:
	//
	// true
	MergeRequestsEnabled *bool `json:"mergeRequestsEnabled,omitempty" xml:"mergeRequestsEnabled,omitempty"`
	// example:
	//
	// codeup
	Name              *string                                      `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string                                      `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	Namespace         *UpdateRepositoryResponseBodyResultNamespace `json:"namespace,omitempty" xml:"namespace,omitempty" type:"Struct"`
	// example:
	//
	// codeup
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// codeup-test-org / codeup
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	// example:
	//
	// false
	SnippetsEnabled *bool `json:"snippetsEnabled,omitempty" xml:"snippetsEnabled,omitempty"`
	// example:
	//
	// git@codeup.aliyun.com:xxx/test/test.git
	SshUrlToRepo *string `json:"sshUrlToRepo,omitempty" xml:"sshUrlToRepo,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	// example:
	//
	// https://codeup.aliyun.com/xxx/test/test
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
	// example:
	//
	// true
	WikiEnabled *bool `json:"wikiEnabled,omitempty" xml:"wikiEnabled,omitempty"`
}

func (s UpdateRepositoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponseBodyResult) GetArchived() *bool {
	return s.Archived
}

func (s *UpdateRepositoryResponseBodyResult) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *UpdateRepositoryResponseBodyResult) GetBuildsEnabled() *bool {
	return s.BuildsEnabled
}

func (s *UpdateRepositoryResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateRepositoryResponseBodyResult) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *UpdateRepositoryResponseBodyResult) GetDefaultBranch() *string {
	return s.DefaultBranch
}

func (s *UpdateRepositoryResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *UpdateRepositoryResponseBodyResult) GetHttpUrlToRepo() *string {
	return s.HttpUrlToRepo
}

func (s *UpdateRepositoryResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *UpdateRepositoryResponseBodyResult) GetIssuesEnabled() *bool {
	return s.IssuesEnabled
}

func (s *UpdateRepositoryResponseBodyResult) GetLastActivityAt() *string {
	return s.LastActivityAt
}

func (s *UpdateRepositoryResponseBodyResult) GetMergeRequestsEnabled() *bool {
	return s.MergeRequestsEnabled
}

func (s *UpdateRepositoryResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *UpdateRepositoryResponseBodyResult) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *UpdateRepositoryResponseBodyResult) GetNamespace() *UpdateRepositoryResponseBodyResultNamespace {
	return s.Namespace
}

func (s *UpdateRepositoryResponseBodyResult) GetPath() *string {
	return s.Path
}

func (s *UpdateRepositoryResponseBodyResult) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *UpdateRepositoryResponseBodyResult) GetSnippetsEnabled() *bool {
	return s.SnippetsEnabled
}

func (s *UpdateRepositoryResponseBodyResult) GetSshUrlToRepo() *string {
	return s.SshUrlToRepo
}

func (s *UpdateRepositoryResponseBodyResult) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *UpdateRepositoryResponseBodyResult) GetWebUrl() *string {
	return s.WebUrl
}

func (s *UpdateRepositoryResponseBodyResult) GetWikiEnabled() *bool {
	return s.WikiEnabled
}

func (s *UpdateRepositoryResponseBodyResult) SetArchived(v bool) *UpdateRepositoryResponseBodyResult {
	s.Archived = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetAvatarUrl(v string) *UpdateRepositoryResponseBodyResult {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetBuildsEnabled(v bool) *UpdateRepositoryResponseBodyResult {
	s.BuildsEnabled = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetCreatedAt(v string) *UpdateRepositoryResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetCreatorId(v int64) *UpdateRepositoryResponseBodyResult {
	s.CreatorId = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetDefaultBranch(v string) *UpdateRepositoryResponseBodyResult {
	s.DefaultBranch = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetDescription(v string) *UpdateRepositoryResponseBodyResult {
	s.Description = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetHttpUrlToRepo(v string) *UpdateRepositoryResponseBodyResult {
	s.HttpUrlToRepo = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetId(v int64) *UpdateRepositoryResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetIssuesEnabled(v bool) *UpdateRepositoryResponseBodyResult {
	s.IssuesEnabled = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetLastActivityAt(v string) *UpdateRepositoryResponseBodyResult {
	s.LastActivityAt = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetMergeRequestsEnabled(v bool) *UpdateRepositoryResponseBodyResult {
	s.MergeRequestsEnabled = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetName(v string) *UpdateRepositoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetNameWithNamespace(v string) *UpdateRepositoryResponseBodyResult {
	s.NameWithNamespace = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetNamespace(v *UpdateRepositoryResponseBodyResultNamespace) *UpdateRepositoryResponseBodyResult {
	s.Namespace = v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetPath(v string) *UpdateRepositoryResponseBodyResult {
	s.Path = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetPathWithNamespace(v string) *UpdateRepositoryResponseBodyResult {
	s.PathWithNamespace = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetSnippetsEnabled(v bool) *UpdateRepositoryResponseBodyResult {
	s.SnippetsEnabled = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetSshUrlToRepo(v string) *UpdateRepositoryResponseBodyResult {
	s.SshUrlToRepo = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetVisibilityLevel(v int32) *UpdateRepositoryResponseBodyResult {
	s.VisibilityLevel = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetWebUrl(v string) *UpdateRepositoryResponseBodyResult {
	s.WebUrl = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) SetWikiEnabled(v bool) *UpdateRepositoryResponseBodyResult {
	s.WikiEnabled = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResult) Validate() error {
	if s.Namespace != nil {
		if err := s.Namespace.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRepositoryResponseBodyResultNamespace struct {
	// example:
	//
	// https://xxx
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// 2022-02-18 14:24:54
	CreatedAt   *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 29322
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 19238
	OwnerId *int64  `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	Path    *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 2022-03-18 14:24:54
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
}

func (s UpdateRepositoryResponseBodyResultNamespace) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryResponseBodyResultNamespace) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponseBodyResultNamespace) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdateRepositoryResponseBodyResultNamespace) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateRepositoryResponseBodyResultNamespace) GetDescription() *string {
	return s.Description
}

func (s *UpdateRepositoryResponseBodyResultNamespace) GetId() *int64 {
	return s.Id
}

func (s *UpdateRepositoryResponseBodyResultNamespace) GetName() *string {
	return s.Name
}

func (s *UpdateRepositoryResponseBodyResultNamespace) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateRepositoryResponseBodyResultNamespace) GetPath() *string {
	return s.Path
}

func (s *UpdateRepositoryResponseBodyResultNamespace) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateRepositoryResponseBodyResultNamespace) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetAvatar(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.Avatar = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetCreatedAt(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.CreatedAt = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetDescription(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.Description = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetId(v int64) *UpdateRepositoryResponseBodyResultNamespace {
	s.Id = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetName(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.Name = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetOwnerId(v int64) *UpdateRepositoryResponseBodyResultNamespace {
	s.OwnerId = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetPath(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.Path = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetUpdatedAt(v string) *UpdateRepositoryResponseBodyResultNamespace {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) SetVisibilityLevel(v int32) *UpdateRepositoryResponseBodyResultNamespace {
	s.VisibilityLevel = &v
	return s
}

func (s *UpdateRepositoryResponseBodyResultNamespace) Validate() error {
	return dara.Validate(s)
}
