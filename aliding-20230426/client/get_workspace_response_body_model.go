// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkspaceResponseBody
	GetRequestId() *string
	SetWorkspace(v *GetWorkspaceResponseBodyWorkspace) *GetWorkspaceResponseBody
	GetWorkspace() *GetWorkspaceResponseBodyWorkspace
}

type GetWorkspaceResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Workspace *GetWorkspaceResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
}

func (s GetWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceResponseBody) GetWorkspace() *GetWorkspaceResponseBodyWorkspace {
	return s.Workspace
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspace(v *GetWorkspaceResponseBodyWorkspace) *GetWorkspaceResponseBody {
	s.Workspace = v
	return s
}

func (s *GetWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWorkspaceResponseBodyWorkspace struct {
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
	// 01472825524039877041
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// workspace_description
	Description *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon        *GetWorkspaceResponseBodyWorkspaceIcon `json:"Icon,omitempty" xml:"Icon,omitempty" type:"Struct"`
	// example:
	//
	// 2023-05-15T11:29Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 01472825524039877041
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
	// ydxXB52LJqqK7xxNTXyo390kJqjMp697
	RootNodeId *string `json:"RootNodeId,omitempty" xml:"RootNodeId,omitempty"`
	// example:
	//
	// lHiicjNFM2iSFYSdz2iPuI8ZwiEiE
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
	// MJ0pDSKMV9dO20E4
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetWorkspaceResponseBodyWorkspace) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyWorkspace) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyWorkspace) GetCorpId() *string {
	return s.CorpId
}

func (s *GetWorkspaceResponseBodyWorkspace) GetCover() *string {
	return s.Cover
}

func (s *GetWorkspaceResponseBodyWorkspace) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetWorkspaceResponseBodyWorkspace) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetWorkspaceResponseBodyWorkspace) GetDescription() *string {
	return s.Description
}

func (s *GetWorkspaceResponseBodyWorkspace) GetIcon() *GetWorkspaceResponseBodyWorkspaceIcon {
	return s.Icon
}

func (s *GetWorkspaceResponseBodyWorkspace) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetWorkspaceResponseBodyWorkspace) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetWorkspaceResponseBodyWorkspace) GetName() *string {
	return s.Name
}

func (s *GetWorkspaceResponseBodyWorkspace) GetPermissionRole() *string {
	return s.PermissionRole
}

func (s *GetWorkspaceResponseBodyWorkspace) GetRootNodeId() *string {
	return s.RootNodeId
}

func (s *GetWorkspaceResponseBodyWorkspace) GetTeamId() *string {
	return s.TeamId
}

func (s *GetWorkspaceResponseBodyWorkspace) GetType() *string {
	return s.Type
}

func (s *GetWorkspaceResponseBodyWorkspace) GetUrl() *string {
	return s.Url
}

func (s *GetWorkspaceResponseBodyWorkspace) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceResponseBodyWorkspace) SetCorpId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.CorpId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetCover(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Cover = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetCreateTime(v string) *GetWorkspaceResponseBodyWorkspace {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetCreatorId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.CreatorId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetDescription(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Description = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetIcon(v *GetWorkspaceResponseBodyWorkspaceIcon) *GetWorkspaceResponseBodyWorkspace {
	s.Icon = v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetModifiedTime(v string) *GetWorkspaceResponseBodyWorkspace {
	s.ModifiedTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetModifierId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.ModifierId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetName(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Name = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetPermissionRole(v string) *GetWorkspaceResponseBodyWorkspace {
	s.PermissionRole = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetRootNodeId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.RootNodeId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetTeamId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.TeamId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetType(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Type = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetUrl(v string) *GetWorkspaceResponseBodyWorkspace {
	s.Url = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) SetWorkspaceId(v string) *GetWorkspaceResponseBodyWorkspace {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspace) Validate() error {
	return dara.Validate(s)
}

type GetWorkspaceResponseBodyWorkspaceIcon struct {
	// example:
	//
	// URL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://example/file-manage-files/zh-CN/202***13/ldet/avatar3.jpg
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetWorkspaceResponseBodyWorkspaceIcon) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyWorkspaceIcon) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyWorkspaceIcon) GetType() *string {
	return s.Type
}

func (s *GetWorkspaceResponseBodyWorkspaceIcon) GetValue() *string {
	return s.Value
}

func (s *GetWorkspaceResponseBodyWorkspaceIcon) SetType(v string) *GetWorkspaceResponseBodyWorkspaceIcon {
	s.Type = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspaceIcon) SetValue(v string) *GetWorkspaceResponseBodyWorkspaceIcon {
	s.Value = &v
	return s
}

func (s *GetWorkspaceResponseBodyWorkspaceIcon) Validate() error {
	return dara.Validate(s)
}
