// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspace interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *Workspace
	GetCreator() *string
	SetGmtCreateTime(v string) *Workspace
	GetGmtCreateTime() *string
	SetGmtModifyTime(v string) *Workspace
	GetGmtModifyTime() *string
	SetMembers(v []*Member) *Workspace
	GetMembers() []*Member
	SetQuotas(v []*Quota) *Workspace
	GetQuotas() []*Quota
	SetTotalResources(v *Resources) *Workspace
	GetTotalResources() *Resources
	SetWorkspaceAdmins(v []*Member) *Workspace
	GetWorkspaceAdmins() []*Member
	SetWorkspaceId(v string) *Workspace
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *Workspace
	GetWorkspaceName() *string
}

type Workspace struct {
	// example:
	//
	// ken
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtModifyTime   *string    `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	Members         []*Member  `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Quotas          []*Quota   `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	TotalResources  *Resources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty"`
	WorkspaceAdmins []*Member  `json:"WorkspaceAdmins,omitempty" xml:"WorkspaceAdmins,omitempty" type:"Repeated"`
	// example:
	//
	// ws-20210126170216-mtl37ge7gkvdz
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// dlc-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s Workspace) String() string {
	return dara.Prettify(s)
}

func (s Workspace) GoString() string {
	return s.String()
}

func (s *Workspace) GetCreator() *string {
	return s.Creator
}

func (s *Workspace) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Workspace) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *Workspace) GetMembers() []*Member {
	return s.Members
}

func (s *Workspace) GetQuotas() []*Quota {
	return s.Quotas
}

func (s *Workspace) GetTotalResources() *Resources {
	return s.TotalResources
}

func (s *Workspace) GetWorkspaceAdmins() []*Member {
	return s.WorkspaceAdmins
}

func (s *Workspace) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Workspace) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *Workspace) SetCreator(v string) *Workspace {
	s.Creator = &v
	return s
}

func (s *Workspace) SetGmtCreateTime(v string) *Workspace {
	s.GmtCreateTime = &v
	return s
}

func (s *Workspace) SetGmtModifyTime(v string) *Workspace {
	s.GmtModifyTime = &v
	return s
}

func (s *Workspace) SetMembers(v []*Member) *Workspace {
	s.Members = v
	return s
}

func (s *Workspace) SetQuotas(v []*Quota) *Workspace {
	s.Quotas = v
	return s
}

func (s *Workspace) SetTotalResources(v *Resources) *Workspace {
	s.TotalResources = v
	return s
}

func (s *Workspace) SetWorkspaceAdmins(v []*Member) *Workspace {
	s.WorkspaceAdmins = v
	return s
}

func (s *Workspace) SetWorkspaceId(v string) *Workspace {
	s.WorkspaceId = &v
	return s
}

func (s *Workspace) SetWorkspaceName(v string) *Workspace {
	s.WorkspaceName = &v
	return s
}

func (s *Workspace) Validate() error {
	return dara.Validate(s)
}
