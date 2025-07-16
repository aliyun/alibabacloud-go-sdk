// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMineWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMineWorkspaceResponseBody
	GetRequestId() *string
	SetWorkspace(v *GetMineWorkspaceResponseBodyWorkspace) *GetMineWorkspaceResponseBody
	GetWorkspace() *GetMineWorkspaceResponseBodyWorkspace
}

type GetMineWorkspaceResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Workspace *GetMineWorkspaceResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
}

func (s GetMineWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMineWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMineWorkspaceResponseBody) GetWorkspace() *GetMineWorkspaceResponseBodyWorkspace {
	return s.Workspace
}

func (s *GetMineWorkspaceResponseBody) SetRequestId(v string) *GetMineWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMineWorkspaceResponseBody) SetWorkspace(v *GetMineWorkspaceResponseBodyWorkspace) *GetMineWorkspaceResponseBody {
	s.Workspace = v
	return s
}

func (s *GetMineWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMineWorkspaceResponseBodyWorkspace struct {
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
	Description *string                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon        *GetMineWorkspaceResponseBodyWorkspaceIcon `json:"Icon,omitempty" xml:"Icon,omitempty" type:"Struct"`
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
	// MNDoBb60VLBPraakI1Ywxyyn8lemrZQ3
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
	// 9r09jSO3WARyxd8A
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetMineWorkspaceResponseBodyWorkspace) String() string {
	return dara.Prettify(s)
}

func (s GetMineWorkspaceResponseBodyWorkspace) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetCorpId() *string {
	return s.CorpId
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetCover() *string {
	return s.Cover
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetDescription() *string {
	return s.Description
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetIcon() *GetMineWorkspaceResponseBodyWorkspaceIcon {
	return s.Icon
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetName() *string {
	return s.Name
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetPermissionRole() *string {
	return s.PermissionRole
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetRootNodeId() *string {
	return s.RootNodeId
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetTeamId() *string {
	return s.TeamId
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetType() *string {
	return s.Type
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetUrl() *string {
	return s.Url
}

func (s *GetMineWorkspaceResponseBodyWorkspace) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetCorpId(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.CorpId = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetCover(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.Cover = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetCreateTime(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.CreateTime = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetCreatorId(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.CreatorId = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetDescription(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.Description = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetIcon(v *GetMineWorkspaceResponseBodyWorkspaceIcon) *GetMineWorkspaceResponseBodyWorkspace {
	s.Icon = v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetModifiedTime(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.ModifiedTime = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetModifierId(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.ModifierId = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetName(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.Name = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetPermissionRole(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.PermissionRole = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetRootNodeId(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.RootNodeId = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetTeamId(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.TeamId = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetType(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.Type = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetUrl(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.Url = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) SetWorkspaceId(v string) *GetMineWorkspaceResponseBodyWorkspace {
	s.WorkspaceId = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspace) Validate() error {
	return dara.Validate(s)
}

type GetMineWorkspaceResponseBodyWorkspaceIcon struct {
	// example:
	//
	// URL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://example/file-manage-files/zh-CN/202***13/ldet/avatar3.jpg
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMineWorkspaceResponseBodyWorkspaceIcon) String() string {
	return dara.Prettify(s)
}

func (s GetMineWorkspaceResponseBodyWorkspaceIcon) GoString() string {
	return s.String()
}

func (s *GetMineWorkspaceResponseBodyWorkspaceIcon) GetType() *string {
	return s.Type
}

func (s *GetMineWorkspaceResponseBodyWorkspaceIcon) GetValue() *string {
	return s.Value
}

func (s *GetMineWorkspaceResponseBodyWorkspaceIcon) SetType(v string) *GetMineWorkspaceResponseBodyWorkspaceIcon {
	s.Type = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspaceIcon) SetValue(v string) *GetMineWorkspaceResponseBodyWorkspaceIcon {
	s.Value = &v
	return s
}

func (s *GetMineWorkspaceResponseBodyWorkspaceIcon) Validate() error {
	return dara.Validate(s)
}
