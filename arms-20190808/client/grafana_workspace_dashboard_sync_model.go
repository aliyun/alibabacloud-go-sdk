// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceDashboardSync interface {
	dara.Model
	String() string
	GoString() string
	SetDashboardTitle(v string) *GrafanaWorkspaceDashboardSync
	GetDashboardTitle() *string
	SetDashboardURL(v string) *GrafanaWorkspaceDashboardSync
	GetDashboardURL() *string
	SetDashboardUid(v string) *GrafanaWorkspaceDashboardSync
	GetDashboardUid() *string
	SetFolderId(v string) *GrafanaWorkspaceDashboardSync
	GetFolderId() *string
	SetFolderTitle(v string) *GrafanaWorkspaceDashboardSync
	GetFolderTitle() *string
	SetFolderURL(v string) *GrafanaWorkspaceDashboardSync
	GetFolderURL() *string
	SetFolderUid(v string) *GrafanaWorkspaceDashboardSync
	GetFolderUid() *string
	SetOrgId(v string) *GrafanaWorkspaceDashboardSync
	GetOrgId() *string
	SetOrgName(v string) *GrafanaWorkspaceDashboardSync
	GetOrgName() *string
	SetType(v string) *GrafanaWorkspaceDashboardSync
	GetType() *string
}

type GrafanaWorkspaceDashboardSync struct {
	// example:
	//
	// testTitle
	DashboardTitle *string `json:"dashboardTitle,omitempty" xml:"dashboardTitle,omitempty"`
	// example:
	//
	// https://g.console.aliyun.com/d/1098370038733503-14960236-422-3/ack-pro-apiserver
	DashboardURL *string `json:"dashboardURL,omitempty" xml:"dashboardURL,omitempty"`
	// example:
	//
	// xxxdvxsea
	DashboardUid *string `json:"dashboardUid,omitempty" xml:"dashboardUid,omitempty"`
	// example:
	//
	// 123456
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// example:
	//
	// test
	FolderTitle *string `json:"folderTitle,omitempty" xml:"folderTitle,omitempty"`
	// example:
	//
	// https://g.console.aliyun.com/d/1098370038733503-14960236-422-3/ack-pro-apiserver
	FolderURL *string `json:"folderURL,omitempty" xml:"folderURL,omitempty"`
	// example:
	//
	// vxeupqn
	FolderUid *string `json:"folderUid,omitempty" xml:"folderUid,omitempty"`
	// example:
	//
	// 1
	OrgId *string `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// example:
	//
	// user123
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// example:
	//
	// normal
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GrafanaWorkspaceDashboardSync) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceDashboardSync) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceDashboardSync) GetDashboardTitle() *string {
	return s.DashboardTitle
}

func (s *GrafanaWorkspaceDashboardSync) GetDashboardURL() *string {
	return s.DashboardURL
}

func (s *GrafanaWorkspaceDashboardSync) GetDashboardUid() *string {
	return s.DashboardUid
}

func (s *GrafanaWorkspaceDashboardSync) GetFolderId() *string {
	return s.FolderId
}

func (s *GrafanaWorkspaceDashboardSync) GetFolderTitle() *string {
	return s.FolderTitle
}

func (s *GrafanaWorkspaceDashboardSync) GetFolderURL() *string {
	return s.FolderURL
}

func (s *GrafanaWorkspaceDashboardSync) GetFolderUid() *string {
	return s.FolderUid
}

func (s *GrafanaWorkspaceDashboardSync) GetOrgId() *string {
	return s.OrgId
}

func (s *GrafanaWorkspaceDashboardSync) GetOrgName() *string {
	return s.OrgName
}

func (s *GrafanaWorkspaceDashboardSync) GetType() *string {
	return s.Type
}

func (s *GrafanaWorkspaceDashboardSync) SetDashboardTitle(v string) *GrafanaWorkspaceDashboardSync {
	s.DashboardTitle = &v
	return s
}

func (s *GrafanaWorkspaceDashboardSync) SetDashboardURL(v string) *GrafanaWorkspaceDashboardSync {
	s.DashboardURL = &v
	return s
}

func (s *GrafanaWorkspaceDashboardSync) SetDashboardUid(v string) *GrafanaWorkspaceDashboardSync {
	s.DashboardUid = &v
	return s
}

func (s *GrafanaWorkspaceDashboardSync) SetFolderId(v string) *GrafanaWorkspaceDashboardSync {
	s.FolderId = &v
	return s
}

func (s *GrafanaWorkspaceDashboardSync) SetFolderTitle(v string) *GrafanaWorkspaceDashboardSync {
	s.FolderTitle = &v
	return s
}

func (s *GrafanaWorkspaceDashboardSync) SetFolderURL(v string) *GrafanaWorkspaceDashboardSync {
	s.FolderURL = &v
	return s
}

func (s *GrafanaWorkspaceDashboardSync) SetFolderUid(v string) *GrafanaWorkspaceDashboardSync {
	s.FolderUid = &v
	return s
}

func (s *GrafanaWorkspaceDashboardSync) SetOrgId(v string) *GrafanaWorkspaceDashboardSync {
	s.OrgId = &v
	return s
}

func (s *GrafanaWorkspaceDashboardSync) SetOrgName(v string) *GrafanaWorkspaceDashboardSync {
	s.OrgName = &v
	return s
}

func (s *GrafanaWorkspaceDashboardSync) SetType(v string) *GrafanaWorkspaceDashboardSync {
	s.Type = &v
	return s
}

func (s *GrafanaWorkspaceDashboardSync) Validate() error {
	return dara.Validate(s)
}
