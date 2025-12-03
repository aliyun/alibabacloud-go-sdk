// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepositoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetRepositoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetRepositoryResponseBody
	GetErrorMessage() *string
	SetRepository(v *GetRepositoryResponseBodyRepository) *GetRepositoryResponseBody
	GetRepository() *GetRepositoryResponseBodyRepository
	SetRequestId(v string) *GetRepositoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRepositoryResponseBody
	GetSuccess() *bool
}

type GetRepositoryResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string                              `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Repository   *GetRepositoryResponseBodyRepository `json:"repository,omitempty" xml:"repository,omitempty" type:"Struct"`
	// example:
	//
	// 37294673-00CA-5B8B-914F-A8B35511E90A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetRepositoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepositoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetRepositoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetRepositoryResponseBody) GetRepository() *GetRepositoryResponseBodyRepository {
	return s.Repository
}

func (s *GetRepositoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRepositoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRepositoryResponseBody) SetErrorCode(v string) *GetRepositoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRepositoryResponseBody) SetErrorMessage(v string) *GetRepositoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepository(v *GetRepositoryResponseBodyRepository) *GetRepositoryResponseBody {
	s.Repository = v
	return s
}

func (s *GetRepositoryResponseBody) SetRequestId(v string) *GetRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepositoryResponseBody) SetSuccess(v bool) *GetRepositoryResponseBody {
	s.Success = &v
	return s
}

func (s *GetRepositoryResponseBody) Validate() error {
	if s.Repository != nil {
		if err := s.Repository.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRepositoryResponseBodyRepository struct {
	// example:
	//
	// false
	Archive *bool `json:"archive,omitempty" xml:"archive,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// 2020-06-19T04:02:00.744Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 19258
	CreatorId *int64 `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// master
	DefaultBranch *string `json:"defaultBranch,omitempty" xml:"defaultBranch,omitempty"`
	// example:
	//
	// false
	DemoProjectStatus *bool `json:"demoProjectStatus,omitempty" xml:"demoProjectStatus,omitempty"`
	// example:
	//
	// repo desc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// https://codeup.aliyun.com/xxx/test/test.git
	HttpUrlToRepository *string `json:"httpUrlToRepository,omitempty" xml:"httpUrlToRepository,omitempty"`
	// example:
	//
	// 100
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2020-06-19T04:02:00.744Z
	LastActivityAt *string `json:"lastActivityAt,omitempty" xml:"lastActivityAt,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// test / test
	NameWithNamespace *string                                       `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	Namespace         *GetRepositoryResponseBodyRepositoryNamespace `json:"namespace,omitempty" xml:"namespace,omitempty" type:"Struct"`
	// example:
	//
	// test
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// test/test
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	// example:
	//
	// git@codeup.aliyun.com:xxx/test/test.git
	SshUrlToRepository *string `json:"sshUrlToRepository,omitempty" xml:"sshUrlToRepository,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	// example:
	//
	// https://codeup.aliyun.com/xxx/test/test
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s GetRepositoryResponseBodyRepository) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryResponseBodyRepository) GoString() string {
	return s.String()
}

func (s *GetRepositoryResponseBodyRepository) GetArchive() *bool {
	return s.Archive
}

func (s *GetRepositoryResponseBodyRepository) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetRepositoryResponseBodyRepository) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetRepositoryResponseBodyRepository) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *GetRepositoryResponseBodyRepository) GetDefaultBranch() *string {
	return s.DefaultBranch
}

func (s *GetRepositoryResponseBodyRepository) GetDemoProjectStatus() *bool {
	return s.DemoProjectStatus
}

func (s *GetRepositoryResponseBodyRepository) GetDescription() *string {
	return s.Description
}

func (s *GetRepositoryResponseBodyRepository) GetHttpUrlToRepository() *string {
	return s.HttpUrlToRepository
}

func (s *GetRepositoryResponseBodyRepository) GetId() *int64 {
	return s.Id
}

func (s *GetRepositoryResponseBodyRepository) GetLastActivityAt() *string {
	return s.LastActivityAt
}

func (s *GetRepositoryResponseBodyRepository) GetName() *string {
	return s.Name
}

func (s *GetRepositoryResponseBodyRepository) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *GetRepositoryResponseBodyRepository) GetNamespace() *GetRepositoryResponseBodyRepositoryNamespace {
	return s.Namespace
}

func (s *GetRepositoryResponseBodyRepository) GetPath() *string {
	return s.Path
}

func (s *GetRepositoryResponseBodyRepository) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *GetRepositoryResponseBodyRepository) GetSshUrlToRepository() *string {
	return s.SshUrlToRepository
}

func (s *GetRepositoryResponseBodyRepository) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *GetRepositoryResponseBodyRepository) GetWebUrl() *string {
	return s.WebUrl
}

func (s *GetRepositoryResponseBodyRepository) SetArchive(v bool) *GetRepositoryResponseBodyRepository {
	s.Archive = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetAvatarUrl(v string) *GetRepositoryResponseBodyRepository {
	s.AvatarUrl = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetCreatedAt(v string) *GetRepositoryResponseBodyRepository {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetCreatorId(v int64) *GetRepositoryResponseBodyRepository {
	s.CreatorId = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetDefaultBranch(v string) *GetRepositoryResponseBodyRepository {
	s.DefaultBranch = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetDemoProjectStatus(v bool) *GetRepositoryResponseBodyRepository {
	s.DemoProjectStatus = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetDescription(v string) *GetRepositoryResponseBodyRepository {
	s.Description = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetHttpUrlToRepository(v string) *GetRepositoryResponseBodyRepository {
	s.HttpUrlToRepository = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetId(v int64) *GetRepositoryResponseBodyRepository {
	s.Id = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetLastActivityAt(v string) *GetRepositoryResponseBodyRepository {
	s.LastActivityAt = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetName(v string) *GetRepositoryResponseBodyRepository {
	s.Name = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetNameWithNamespace(v string) *GetRepositoryResponseBodyRepository {
	s.NameWithNamespace = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetNamespace(v *GetRepositoryResponseBodyRepositoryNamespace) *GetRepositoryResponseBodyRepository {
	s.Namespace = v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetPath(v string) *GetRepositoryResponseBodyRepository {
	s.Path = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetPathWithNamespace(v string) *GetRepositoryResponseBodyRepository {
	s.PathWithNamespace = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetSshUrlToRepository(v string) *GetRepositoryResponseBodyRepository {
	s.SshUrlToRepository = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetVisibilityLevel(v int32) *GetRepositoryResponseBodyRepository {
	s.VisibilityLevel = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) SetWebUrl(v string) *GetRepositoryResponseBodyRepository {
	s.WebUrl = &v
	return s
}

func (s *GetRepositoryResponseBodyRepository) Validate() error {
	if s.Namespace != nil {
		if err := s.Namespace.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRepositoryResponseBodyRepositoryNamespace struct {
	// example:
	//
	// https://xxx.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// 2020-06-19T04:02:00.744Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// repo desc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// id
	//
	// example:
	//
	// 100
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// xxx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 100
	OwnerId *int64 `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// example:
	//
	// test
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 2020-06-19T04:02:00.744Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
}

func (s GetRepositoryResponseBodyRepositoryNamespace) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryResponseBodyRepositoryNamespace) GoString() string {
	return s.String()
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) GetAvatar() *string {
	return s.Avatar
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) GetDescription() *string {
	return s.Description
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) GetId() *int64 {
	return s.Id
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) GetName() *string {
	return s.Name
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) GetPath() *string {
	return s.Path
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetAvatar(v string) *GetRepositoryResponseBodyRepositoryNamespace {
	s.Avatar = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetCreatedAt(v string) *GetRepositoryResponseBodyRepositoryNamespace {
	s.CreatedAt = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetDescription(v string) *GetRepositoryResponseBodyRepositoryNamespace {
	s.Description = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetId(v int64) *GetRepositoryResponseBodyRepositoryNamespace {
	s.Id = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetName(v string) *GetRepositoryResponseBodyRepositoryNamespace {
	s.Name = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetOwnerId(v int64) *GetRepositoryResponseBodyRepositoryNamespace {
	s.OwnerId = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetPath(v string) *GetRepositoryResponseBodyRepositoryNamespace {
	s.Path = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetUpdatedAt(v string) *GetRepositoryResponseBodyRepositoryNamespace {
	s.UpdatedAt = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) SetVisibilityLevel(v int32) *GetRepositoryResponseBodyRepositoryNamespace {
	s.VisibilityLevel = &v
	return s
}

func (s *GetRepositoryResponseBodyRepositoryNamespace) Validate() error {
	return dara.Validate(s)
}
