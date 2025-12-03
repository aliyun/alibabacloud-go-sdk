// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListUserResourcesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListUserResourcesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListUserResourcesResponseBody
	GetRequestId() *string
	SetResult(v []*ListUserResourcesResponseBodyResult) *ListUserResourcesResponseBody
	GetResult() []*ListUserResourcesResponseBodyResult
	SetSuccess(v bool) *ListUserResourcesResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListUserResourcesResponseBody
	GetTotal() *int64
}

type ListUserResourcesResponseBody struct {
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// A35869D5-BB29-5F84-A4DD-B09985EA2AFA
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListUserResourcesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListUserResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListUserResourcesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListUserResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserResourcesResponseBody) GetResult() []*ListUserResourcesResponseBodyResult {
	return s.Result
}

func (s *ListUserResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserResourcesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListUserResourcesResponseBody) SetErrorCode(v string) *ListUserResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetErrorMessage(v string) *ListUserResourcesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetRequestId(v string) *ListUserResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetResult(v []*ListUserResourcesResponseBodyResult) *ListUserResourcesResponseBody {
	s.Result = v
	return s
}

func (s *ListUserResourcesResponseBody) SetSuccess(v bool) *ListUserResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetTotal(v int64) *ListUserResourcesResponseBody {
	s.Total = &v
	return s
}

func (s *ListUserResourcesResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserResourcesResponseBodyResult struct {
	GroupInfos      []*ListUserResourcesResponseBodyResultGroupInfos      `json:"groupInfos,omitempty" xml:"groupInfos,omitempty" type:"Repeated"`
	RepositoryInfos []*ListUserResourcesResponseBodyResultRepositoryInfos `json:"repositoryInfos,omitempty" xml:"repositoryInfos,omitempty" type:"Repeated"`
	UserInfo        *ListUserResourcesResponseBodyResultUserInfo          `json:"userInfo,omitempty" xml:"userInfo,omitempty" type:"Struct"`
}

func (s ListUserResourcesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBodyResult) GetGroupInfos() []*ListUserResourcesResponseBodyResultGroupInfos {
	return s.GroupInfos
}

func (s *ListUserResourcesResponseBodyResult) GetRepositoryInfos() []*ListUserResourcesResponseBodyResultRepositoryInfos {
	return s.RepositoryInfos
}

func (s *ListUserResourcesResponseBodyResult) GetUserInfo() *ListUserResourcesResponseBodyResultUserInfo {
	return s.UserInfo
}

func (s *ListUserResourcesResponseBodyResult) SetGroupInfos(v []*ListUserResourcesResponseBodyResultGroupInfos) *ListUserResourcesResponseBodyResult {
	s.GroupInfos = v
	return s
}

func (s *ListUserResourcesResponseBodyResult) SetRepositoryInfos(v []*ListUserResourcesResponseBodyResultRepositoryInfos) *ListUserResourcesResponseBodyResult {
	s.RepositoryInfos = v
	return s
}

func (s *ListUserResourcesResponseBodyResult) SetUserInfo(v *ListUserResourcesResponseBodyResultUserInfo) *ListUserResourcesResponseBodyResult {
	s.UserInfo = v
	return s
}

func (s *ListUserResourcesResponseBodyResult) Validate() error {
	if s.GroupInfos != nil {
		for _, item := range s.GroupInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RepositoryInfos != nil {
		for _, item := range s.RepositoryInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserResourcesResponseBodyResultGroupInfos struct {
	GroupInfo *ListUserResourcesResponseBodyResultGroupInfosGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	GroupRole *ListUserResourcesResponseBodyResultGroupInfosGroupRole `json:"groupRole,omitempty" xml:"groupRole,omitempty" type:"Struct"`
}

func (s ListUserResourcesResponseBodyResultGroupInfos) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesResponseBodyResultGroupInfos) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBodyResultGroupInfos) GetGroupInfo() *ListUserResourcesResponseBodyResultGroupInfosGroupInfo {
	return s.GroupInfo
}

func (s *ListUserResourcesResponseBodyResultGroupInfos) GetGroupRole() *ListUserResourcesResponseBodyResultGroupInfosGroupRole {
	return s.GroupRole
}

func (s *ListUserResourcesResponseBodyResultGroupInfos) SetGroupInfo(v *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) *ListUserResourcesResponseBodyResultGroupInfos {
	s.GroupInfo = v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfos) SetGroupRole(v *ListUserResourcesResponseBodyResultGroupInfosGroupRole) *ListUserResourcesResponseBodyResultGroupInfos {
	s.GroupRole = v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfos) Validate() error {
	if s.GroupInfo != nil {
		if err := s.GroupInfo.Validate(); err != nil {
			return err
		}
	}
	if s.GroupRole != nil {
		if err := s.GroupRole.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserResourcesResponseBodyResultGroupInfosGroupInfo struct {
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// test-group
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 35268
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test-group
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	// example:
	//
	// 1234
	OwnerId *int64 `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// example:
	//
	// 1183319
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// test-group
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// test-org/test-group
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
}

func (s ListUserResourcesResponseBodyResultGroupInfosGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesResponseBodyResultGroupInfosGroupInfo) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) GetDescription() *string {
	return s.Description
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) GetId() *int64 {
	return s.Id
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) GetName() *string {
	return s.Name
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) GetParentId() *int64 {
	return s.ParentId
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) GetPath() *string {
	return s.Path
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) SetCreatedAt(v string) *ListUserResourcesResponseBodyResultGroupInfosGroupInfo {
	s.CreatedAt = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) SetDescription(v string) *ListUserResourcesResponseBodyResultGroupInfosGroupInfo {
	s.Description = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) SetId(v int64) *ListUserResourcesResponseBodyResultGroupInfosGroupInfo {
	s.Id = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) SetName(v string) *ListUserResourcesResponseBodyResultGroupInfosGroupInfo {
	s.Name = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) SetNameWithNamespace(v string) *ListUserResourcesResponseBodyResultGroupInfosGroupInfo {
	s.NameWithNamespace = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) SetOwnerId(v int64) *ListUserResourcesResponseBodyResultGroupInfosGroupInfo {
	s.OwnerId = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) SetParentId(v int64) *ListUserResourcesResponseBodyResultGroupInfosGroupInfo {
	s.ParentId = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) SetPath(v string) *ListUserResourcesResponseBodyResultGroupInfosGroupInfo {
	s.Path = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) SetPathWithNamespace(v string) *ListUserResourcesResponseBodyResultGroupInfosGroupInfo {
	s.PathWithNamespace = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) SetUpdatedAt(v string) *ListUserResourcesResponseBodyResultGroupInfosGroupInfo {
	s.UpdatedAt = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) SetVisibilityLevel(v int32) *ListUserResourcesResponseBodyResultGroupInfosGroupInfo {
	s.VisibilityLevel = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupInfo) Validate() error {
	return dara.Validate(s)
}

type ListUserResourcesResponseBodyResultGroupInfosGroupRole struct {
	// example:
	//
	// 40
	AccessLevel *int32  `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	CnRoleName  *string `json:"cnRoleName,omitempty" xml:"cnRoleName,omitempty"`
	// example:
	//
	// Admin
	EnRoleName *string `json:"enRoleName,omitempty" xml:"enRoleName,omitempty"`
	// example:
	//
	// 35268
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// Namespace
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListUserResourcesResponseBodyResultGroupInfosGroupRole) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesResponseBodyResultGroupInfosGroupRole) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupRole) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupRole) GetCnRoleName() *string {
	return s.CnRoleName
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupRole) GetEnRoleName() *string {
	return s.EnRoleName
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupRole) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupRole) GetSourceType() *string {
	return s.SourceType
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupRole) SetAccessLevel(v int32) *ListUserResourcesResponseBodyResultGroupInfosGroupRole {
	s.AccessLevel = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupRole) SetCnRoleName(v string) *ListUserResourcesResponseBodyResultGroupInfosGroupRole {
	s.CnRoleName = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupRole) SetEnRoleName(v string) *ListUserResourcesResponseBodyResultGroupInfosGroupRole {
	s.EnRoleName = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupRole) SetSourceId(v int64) *ListUserResourcesResponseBodyResultGroupInfosGroupRole {
	s.SourceId = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupRole) SetSourceType(v string) *ListUserResourcesResponseBodyResultGroupInfosGroupRole {
	s.SourceType = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultGroupInfosGroupRole) Validate() error {
	return dara.Validate(s)
}

type ListUserResourcesResponseBodyResultRepositoryInfos struct {
	RepositoryInfo *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo `json:"repositoryInfo,omitempty" xml:"repositoryInfo,omitempty" type:"Struct"`
	RepositoryRole *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole `json:"repositoryRole,omitempty" xml:"repositoryRole,omitempty" type:"Struct"`
}

func (s ListUserResourcesResponseBodyResultRepositoryInfos) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesResponseBodyResultRepositoryInfos) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfos) GetRepositoryInfo() *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	return s.RepositoryInfo
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfos) GetRepositoryRole() *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole {
	return s.RepositoryRole
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfos) SetRepositoryInfo(v *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) *ListUserResourcesResponseBodyResultRepositoryInfos {
	s.RepositoryInfo = v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfos) SetRepositoryRole(v *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) *ListUserResourcesResponseBodyResultRepositoryInfos {
	s.RepositoryRole = v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfos) Validate() error {
	if s.RepositoryInfo != nil {
		if err := s.RepositoryInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RepositoryRole != nil {
		if err := s.RepositoryRole.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo struct {
	// example:
	//
	// 40
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	// example:
	//
	// false
	Archived *bool `json:"archived,omitempty" xml:"archived,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 12679
	CreatorId   *int64  `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// false
	Encrypted *bool `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// example:
	//
	// 37229
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	LastActivityAt *string `json:"lastActivityAt,omitempty" xml:"lastActivityAt,omitempty"`
	// example:
	//
	// test-repo
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	NameWithNamespace *string `json:"nameWithNamespace,omitempty" xml:"nameWithNamespace,omitempty"`
	// example:
	//
	// 600002
	NamespaceId *int64 `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	// example:
	//
	// test-repo
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// test-org/test-group/test-repo
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	// example:
	//
	// 2022-01-14T21:08:26+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
}

func (s ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetArchived() *bool {
	return s.Archived
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetDescription() *string {
	return s.Description
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetId() *int64 {
	return s.Id
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetLastActivityAt() *string {
	return s.LastActivityAt
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetName() *string {
	return s.Name
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetNameWithNamespace() *string {
	return s.NameWithNamespace
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetNamespaceId() *int64 {
	return s.NamespaceId
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetPath() *string {
	return s.Path
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetAccessLevel(v int32) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.AccessLevel = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetArchived(v bool) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.Archived = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetCreatedAt(v string) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.CreatedAt = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetCreatorId(v int64) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.CreatorId = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetDescription(v string) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.Description = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetEncrypted(v bool) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.Encrypted = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetId(v int64) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.Id = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetLastActivityAt(v string) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.LastActivityAt = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetName(v string) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.Name = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetNameWithNamespace(v string) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.NameWithNamespace = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetNamespaceId(v int64) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.NamespaceId = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetPath(v string) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.Path = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetPathWithNamespace(v string) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.PathWithNamespace = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetUpdatedAt(v string) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.UpdatedAt = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) SetVisibilityLevel(v int32) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo {
	s.VisibilityLevel = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryInfo) Validate() error {
	return dara.Validate(s)
}

type ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole struct {
	// example:
	//
	// 40
	AccessLevel *int32  `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	CnRoleName  *string `json:"cnRoleName,omitempty" xml:"cnRoleName,omitempty"`
	// example:
	//
	// Admin
	EnRoleName *string `json:"enRoleName,omitempty" xml:"enRoleName,omitempty"`
	// example:
	//
	// 37229
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// Project
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) GetCnRoleName() *string {
	return s.CnRoleName
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) GetEnRoleName() *string {
	return s.EnRoleName
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) GetSourceType() *string {
	return s.SourceType
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) SetAccessLevel(v int32) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole {
	s.AccessLevel = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) SetCnRoleName(v string) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole {
	s.CnRoleName = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) SetEnRoleName(v string) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole {
	s.EnRoleName = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) SetSourceId(v int64) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole {
	s.SourceId = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) SetSourceType(v string) *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole {
	s.SourceType = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultRepositoryInfosRepositoryRole) Validate() error {
	return dara.Validate(s)
}

type ListUserResourcesResponseBodyResultUserInfo struct {
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 19230
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test-codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// test-codeup
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListUserResourcesResponseBodyResultUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesResponseBodyResultUserInfo) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBodyResultUserInfo) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *ListUserResourcesResponseBodyResultUserInfo) GetEmail() *string {
	return s.Email
}

func (s *ListUserResourcesResponseBodyResultUserInfo) GetId() *int64 {
	return s.Id
}

func (s *ListUserResourcesResponseBodyResultUserInfo) GetName() *string {
	return s.Name
}

func (s *ListUserResourcesResponseBodyResultUserInfo) GetState() *string {
	return s.State
}

func (s *ListUserResourcesResponseBodyResultUserInfo) GetUsername() *string {
	return s.Username
}

func (s *ListUserResourcesResponseBodyResultUserInfo) SetAvatarUrl(v string) *ListUserResourcesResponseBodyResultUserInfo {
	s.AvatarUrl = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultUserInfo) SetEmail(v string) *ListUserResourcesResponseBodyResultUserInfo {
	s.Email = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultUserInfo) SetId(v int64) *ListUserResourcesResponseBodyResultUserInfo {
	s.Id = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultUserInfo) SetName(v string) *ListUserResourcesResponseBodyResultUserInfo {
	s.Name = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultUserInfo) SetState(v string) *ListUserResourcesResponseBodyResultUserInfo {
	s.State = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultUserInfo) SetUsername(v string) *ListUserResourcesResponseBodyResultUserInfo {
	s.Username = &v
	return s
}

func (s *ListUserResourcesResponseBodyResultUserInfo) Validate() error {
	return dara.Validate(s)
}
