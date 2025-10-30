// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetWorkspaceResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetWorkspaceResponseBody
	GetDescription() *string
	SetDisplayName(v string) *GetWorkspaceResponseBody
	GetDisplayName() *string
	SetLastModifyTime(v string) *GetWorkspaceResponseBody
	GetLastModifyTime() *string
	SetRegionId(v string) *GetWorkspaceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetWorkspaceResponseBody
	GetRequestId() *string
	SetSlsProject(v string) *GetWorkspaceResponseBody
	GetSlsProject() *string
	SetWorkspaceName(v string) *GetWorkspaceResponseBody
	GetWorkspaceName() *string
}

type GetWorkspaceResponseBody struct {
	// Creation Time
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Workspace Description
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Workspace Display Name
	//
	// example:
	//
	// workspace-test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Last Modified Time
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	LastModifyTime *string `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Request ID
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Log Service Project Name
	//
	// example:
	//
	// sls-project-test-001
	SlsProject *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
	// Workspace Name
	//
	// This parameter is required.
	//
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetWorkspaceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetWorkspaceResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetWorkspaceResponseBody) GetLastModifyTime() *string {
	return s.LastModifyTime
}

func (s *GetWorkspaceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceResponseBody) GetSlsProject() *string {
	return s.SlsProject
}

func (s *GetWorkspaceResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *GetWorkspaceResponseBody) SetCreateTime(v string) *GetWorkspaceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetDescription(v string) *GetWorkspaceResponseBody {
	s.Description = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetDisplayName(v string) *GetWorkspaceResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetLastModifyTime(v string) *GetWorkspaceResponseBody {
	s.LastModifyTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRegionId(v string) *GetWorkspaceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetSlsProject(v string) *GetWorkspaceResponseBody {
	s.SlsProject = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspaceName(v string) *GetWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *GetWorkspaceResponseBody) Validate() error {
	return dara.Validate(s)
}
