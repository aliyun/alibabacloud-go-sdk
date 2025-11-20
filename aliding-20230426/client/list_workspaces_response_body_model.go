// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListWorkspacesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkspacesResponseBody
	GetRequestId() *string
	SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody
	GetWorkspaces() []*ListWorkspacesResponseBodyWorkspaces
}

type ListWorkspacesResponseBody struct {
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId  *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Workspaces []*ListWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspacesResponseBody) GetWorkspaces() []*ListWorkspacesResponseBodyWorkspaces {
	return s.Workspaces
}

func (s *ListWorkspacesResponseBody) SetNextToken(v string) *ListWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

func (s *ListWorkspacesResponseBody) Validate() error {
	if s.Workspaces != nil {
		for _, item := range s.Workspaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspacesResponseBodyWorkspaces struct {
	// example:
	//
	// ding16b241fd05********288
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// https://example/file-manage-files/zh-CN/202***13/ldet/XXXX.jpg
	Cover *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	// example:
	//
	// 2023-05-15T11:29Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// workspace_creator_id
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// workspace_description
	Description *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon        *ListWorkspacesResponseBodyWorkspacesIcon `json:"Icon,omitempty" xml:"Icon,omitempty" type:"Struct"`
	// example:
	//
	// 2023-05-15T11:29Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// workspace_modifier_id
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// workspace_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// READER
	PermissionRole *string `json:"PermissionRole,omitempty" xml:"PermissionRole,omitempty"`
	// example:
	//
	// ZgpG2NdyVXXjrKKzIBqMp5zkVMwvDqPk
	RootNodeId *string `json:"RootNodeId,omitempty" xml:"RootNodeId,omitempty"`
	// example:
	//
	// iPes3KGYA9DxYSdz2iPuI8ZwiEiE
	TeamId *string `json:"TeamId,omitempty" xml:"TeamId,omitempty"`
	// example:
	//
	// TEAM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// workspace_url
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// By8jQS1ZYjGn5b0M
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaces) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetCorpId() *string {
	return s.CorpId
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetCover() *string {
	return s.Cover
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetDescription() *string {
	return s.Description
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetIcon() *ListWorkspacesResponseBodyWorkspacesIcon {
	return s.Icon
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetModifierId() *string {
	return s.ModifierId
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetName() *string {
	return s.Name
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetPermissionRole() *string {
	return s.PermissionRole
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetRootNodeId() *string {
	return s.RootNodeId
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetTeamId() *string {
	return s.TeamId
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetType() *string {
	return s.Type
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetUrl() *string {
	return s.Url
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCorpId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CorpId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCover(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Cover = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreateTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreatorId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CreatorId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDescription(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Description = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetIcon(v *ListWorkspacesResponseBodyWorkspacesIcon) *ListWorkspacesResponseBodyWorkspaces {
	s.Icon = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetModifiedTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.ModifiedTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetModifierId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.ModifierId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Name = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetPermissionRole(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.PermissionRole = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetRootNodeId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.RootNodeId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetTeamId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.TeamId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetType(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Type = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetUrl(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Url = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) Validate() error {
	if s.Icon != nil {
		if err := s.Icon.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkspacesResponseBodyWorkspacesIcon struct {
	// example:
	//
	// URL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// icon_url
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspacesIcon) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspacesIcon) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspacesIcon) GetType() *string {
	return s.Type
}

func (s *ListWorkspacesResponseBodyWorkspacesIcon) GetValue() *string {
	return s.Value
}

func (s *ListWorkspacesResponseBodyWorkspacesIcon) SetType(v string) *ListWorkspacesResponseBodyWorkspacesIcon {
	s.Type = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesIcon) SetValue(v string) *ListWorkspacesResponseBodyWorkspacesIcon {
	s.Value = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspacesIcon) Validate() error {
	return dara.Validate(s)
}
