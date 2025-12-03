// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *UpdateRepositoryRequest
	GetAccessToken() *string
	SetAdminSettingLanguage(v string) *UpdateRepositoryRequest
	GetAdminSettingLanguage() *string
	SetAvatar(v string) *UpdateRepositoryRequest
	GetAvatar() *string
	SetBuildsEnabled(v bool) *UpdateRepositoryRequest
	GetBuildsEnabled() *bool
	SetCheckEmail(v bool) *UpdateRepositoryRequest
	GetCheckEmail() *bool
	SetDefaultBranch(v string) *UpdateRepositoryRequest
	GetDefaultBranch() *string
	SetDescription(v string) *UpdateRepositoryRequest
	GetDescription() *string
	SetId(v int64) *UpdateRepositoryRequest
	GetId() *int64
	SetIssuesEnabled(v bool) *UpdateRepositoryRequest
	GetIssuesEnabled() *bool
	SetMergeRequestsEnabled(v bool) *UpdateRepositoryRequest
	GetMergeRequestsEnabled() *bool
	SetName(v string) *UpdateRepositoryRequest
	GetName() *string
	SetOpenCloneDownloadControl(v bool) *UpdateRepositoryRequest
	GetOpenCloneDownloadControl() *bool
	SetPath(v string) *UpdateRepositoryRequest
	GetPath() *string
	SetProjectCloneDownloadMethodList(v []*UpdateRepositoryRequestProjectCloneDownloadMethodList) *UpdateRepositoryRequest
	GetProjectCloneDownloadMethodList() []*UpdateRepositoryRequestProjectCloneDownloadMethodList
	SetProjectCloneDownloadRoleList(v []*UpdateRepositoryRequestProjectCloneDownloadRoleList) *UpdateRepositoryRequest
	GetProjectCloneDownloadRoleList() []*UpdateRepositoryRequestProjectCloneDownloadRoleList
	SetSnippetsEnabled(v bool) *UpdateRepositoryRequest
	GetSnippetsEnabled() *bool
	SetVisibilityLevel(v int32) *UpdateRepositoryRequest
	GetVisibilityLevel() *int32
	SetWikiEnabled(v bool) *UpdateRepositoryRequest
	GetWikiEnabled() *bool
	SetOrganizationId(v string) *UpdateRepositoryRequest
	GetOrganizationId() *string
}

type UpdateRepositoryRequest struct {
	// example:
	//
	// c3c09f1230187a879678da43c973d069
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// Java
	AdminSettingLanguage *string `json:"adminSettingLanguage,omitempty" xml:"adminSettingLanguage,omitempty"`
	// example:
	//
	// https://xxxx
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// true
	BuildsEnabled *bool `json:"buildsEnabled,omitempty" xml:"buildsEnabled,omitempty"`
	// example:
	//
	// true
	CheckEmail *bool `json:"checkEmail,omitempty" xml:"checkEmail,omitempty"`
	// example:
	//
	// master
	DefaultBranch *string `json:"defaultBranch,omitempty" xml:"defaultBranch,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2080398
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IssuesEnabled *bool `json:"issuesEnabled,omitempty" xml:"issuesEnabled,omitempty"`
	// example:
	//
	// true
	MergeRequestsEnabled *bool `json:"mergeRequestsEnabled,omitempty" xml:"mergeRequestsEnabled,omitempty"`
	// example:
	//
	// codeup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	OpenCloneDownloadControl       *bool                                                    `json:"openCloneDownloadControl,omitempty" xml:"openCloneDownloadControl,omitempty"`
	Path                           *string                                                  `json:"path,omitempty" xml:"path,omitempty"`
	ProjectCloneDownloadMethodList []*UpdateRepositoryRequestProjectCloneDownloadMethodList `json:"projectCloneDownloadMethodList,omitempty" xml:"projectCloneDownloadMethodList,omitempty" type:"Repeated"`
	ProjectCloneDownloadRoleList   []*UpdateRepositoryRequestProjectCloneDownloadRoleList   `json:"projectCloneDownloadRoleList,omitempty" xml:"projectCloneDownloadRoleList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	SnippetsEnabled *bool `json:"snippetsEnabled,omitempty" xml:"snippetsEnabled,omitempty"`
	// example:
	//
	// 0
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	// example:
	//
	// true
	WikiEnabled *bool `json:"wikiEnabled,omitempty" xml:"wikiEnabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdateRepositoryRequest) GetAdminSettingLanguage() *string {
	return s.AdminSettingLanguage
}

func (s *UpdateRepositoryRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdateRepositoryRequest) GetBuildsEnabled() *bool {
	return s.BuildsEnabled
}

func (s *UpdateRepositoryRequest) GetCheckEmail() *bool {
	return s.CheckEmail
}

func (s *UpdateRepositoryRequest) GetDefaultBranch() *string {
	return s.DefaultBranch
}

func (s *UpdateRepositoryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRepositoryRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateRepositoryRequest) GetIssuesEnabled() *bool {
	return s.IssuesEnabled
}

func (s *UpdateRepositoryRequest) GetMergeRequestsEnabled() *bool {
	return s.MergeRequestsEnabled
}

func (s *UpdateRepositoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateRepositoryRequest) GetOpenCloneDownloadControl() *bool {
	return s.OpenCloneDownloadControl
}

func (s *UpdateRepositoryRequest) GetPath() *string {
	return s.Path
}

func (s *UpdateRepositoryRequest) GetProjectCloneDownloadMethodList() []*UpdateRepositoryRequestProjectCloneDownloadMethodList {
	return s.ProjectCloneDownloadMethodList
}

func (s *UpdateRepositoryRequest) GetProjectCloneDownloadRoleList() []*UpdateRepositoryRequestProjectCloneDownloadRoleList {
	return s.ProjectCloneDownloadRoleList
}

func (s *UpdateRepositoryRequest) GetSnippetsEnabled() *bool {
	return s.SnippetsEnabled
}

func (s *UpdateRepositoryRequest) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *UpdateRepositoryRequest) GetWikiEnabled() *bool {
	return s.WikiEnabled
}

func (s *UpdateRepositoryRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateRepositoryRequest) SetAccessToken(v string) *UpdateRepositoryRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateRepositoryRequest) SetAdminSettingLanguage(v string) *UpdateRepositoryRequest {
	s.AdminSettingLanguage = &v
	return s
}

func (s *UpdateRepositoryRequest) SetAvatar(v string) *UpdateRepositoryRequest {
	s.Avatar = &v
	return s
}

func (s *UpdateRepositoryRequest) SetBuildsEnabled(v bool) *UpdateRepositoryRequest {
	s.BuildsEnabled = &v
	return s
}

func (s *UpdateRepositoryRequest) SetCheckEmail(v bool) *UpdateRepositoryRequest {
	s.CheckEmail = &v
	return s
}

func (s *UpdateRepositoryRequest) SetDefaultBranch(v string) *UpdateRepositoryRequest {
	s.DefaultBranch = &v
	return s
}

func (s *UpdateRepositoryRequest) SetDescription(v string) *UpdateRepositoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateRepositoryRequest) SetId(v int64) *UpdateRepositoryRequest {
	s.Id = &v
	return s
}

func (s *UpdateRepositoryRequest) SetIssuesEnabled(v bool) *UpdateRepositoryRequest {
	s.IssuesEnabled = &v
	return s
}

func (s *UpdateRepositoryRequest) SetMergeRequestsEnabled(v bool) *UpdateRepositoryRequest {
	s.MergeRequestsEnabled = &v
	return s
}

func (s *UpdateRepositoryRequest) SetName(v string) *UpdateRepositoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateRepositoryRequest) SetOpenCloneDownloadControl(v bool) *UpdateRepositoryRequest {
	s.OpenCloneDownloadControl = &v
	return s
}

func (s *UpdateRepositoryRequest) SetPath(v string) *UpdateRepositoryRequest {
	s.Path = &v
	return s
}

func (s *UpdateRepositoryRequest) SetProjectCloneDownloadMethodList(v []*UpdateRepositoryRequestProjectCloneDownloadMethodList) *UpdateRepositoryRequest {
	s.ProjectCloneDownloadMethodList = v
	return s
}

func (s *UpdateRepositoryRequest) SetProjectCloneDownloadRoleList(v []*UpdateRepositoryRequestProjectCloneDownloadRoleList) *UpdateRepositoryRequest {
	s.ProjectCloneDownloadRoleList = v
	return s
}

func (s *UpdateRepositoryRequest) SetSnippetsEnabled(v bool) *UpdateRepositoryRequest {
	s.SnippetsEnabled = &v
	return s
}

func (s *UpdateRepositoryRequest) SetVisibilityLevel(v int32) *UpdateRepositoryRequest {
	s.VisibilityLevel = &v
	return s
}

func (s *UpdateRepositoryRequest) SetWikiEnabled(v bool) *UpdateRepositoryRequest {
	s.WikiEnabled = &v
	return s
}

func (s *UpdateRepositoryRequest) SetOrganizationId(v string) *UpdateRepositoryRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateRepositoryRequest) Validate() error {
	if s.ProjectCloneDownloadMethodList != nil {
		for _, item := range s.ProjectCloneDownloadMethodList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProjectCloneDownloadRoleList != nil {
		for _, item := range s.ProjectCloneDownloadRoleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRepositoryRequestProjectCloneDownloadMethodList struct {
	// example:
	//
	// true
	Allowed *bool `json:"allowed,omitempty" xml:"allowed,omitempty"`
	// example:
	//
	// project:download
	PermissionCode *string `json:"permissionCode,omitempty" xml:"permissionCode,omitempty"`
}

func (s UpdateRepositoryRequestProjectCloneDownloadMethodList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryRequestProjectCloneDownloadMethodList) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryRequestProjectCloneDownloadMethodList) GetAllowed() *bool {
	return s.Allowed
}

func (s *UpdateRepositoryRequestProjectCloneDownloadMethodList) GetPermissionCode() *string {
	return s.PermissionCode
}

func (s *UpdateRepositoryRequestProjectCloneDownloadMethodList) SetAllowed(v bool) *UpdateRepositoryRequestProjectCloneDownloadMethodList {
	s.Allowed = &v
	return s
}

func (s *UpdateRepositoryRequestProjectCloneDownloadMethodList) SetPermissionCode(v string) *UpdateRepositoryRequestProjectCloneDownloadMethodList {
	s.PermissionCode = &v
	return s
}

func (s *UpdateRepositoryRequestProjectCloneDownloadMethodList) Validate() error {
	return dara.Validate(s)
}

type UpdateRepositoryRequestProjectCloneDownloadRoleList struct {
	// example:
	//
	// true
	Allowed *bool `json:"allowed,omitempty" xml:"allowed,omitempty"`
	// example:
	//
	// 15
	RoleCode *int32 `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
}

func (s UpdateRepositoryRequestProjectCloneDownloadRoleList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryRequestProjectCloneDownloadRoleList) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryRequestProjectCloneDownloadRoleList) GetAllowed() *bool {
	return s.Allowed
}

func (s *UpdateRepositoryRequestProjectCloneDownloadRoleList) GetRoleCode() *int32 {
	return s.RoleCode
}

func (s *UpdateRepositoryRequestProjectCloneDownloadRoleList) SetAllowed(v bool) *UpdateRepositoryRequestProjectCloneDownloadRoleList {
	s.Allowed = &v
	return s
}

func (s *UpdateRepositoryRequestProjectCloneDownloadRoleList) SetRoleCode(v int32) *UpdateRepositoryRequestProjectCloneDownloadRoleList {
	s.RoleCode = &v
	return s
}

func (s *UpdateRepositoryRequestProjectCloneDownloadRoleList) Validate() error {
	return dara.Validate(s)
}
