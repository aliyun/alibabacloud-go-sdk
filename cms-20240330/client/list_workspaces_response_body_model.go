// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkspacesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkspacesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkspacesResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListWorkspacesResponseBody
	GetTotal() *int32
	SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody
	GetWorkspaces() []*ListWorkspacesResponseBodyWorkspaces
}

type ListWorkspacesResponseBody struct {
	// The number of entries returned per page. Default value: 50. Maximum value: 50.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page of results.
	//
	// example:
	//
	// xxxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
	// The list of workspaces.
	Workspaces []*ListWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspacesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListWorkspacesResponseBody) GetWorkspaces() []*ListWorkspacesResponseBodyWorkspaces {
	return s.Workspaces
}

func (s *ListWorkspacesResponseBody) SetMaxResults(v int32) *ListWorkspacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetNextToken(v string) *ListWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotal(v int32) *ListWorkspacesResponseBody {
	s.Total = &v
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
	// The time when the workspace was created.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the workspace.
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the workspace.
	//
	// example:
	//
	// workspace-test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The time when the workspace was last modified.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2025-03-11T08:21:58Z
	LastModifyTime *string `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// sls-project-test-001
	SlsProject *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
	// The name of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaces) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetDescription() *string {
	return s.Description
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetLastModifyTime() *string {
	return s.LastModifyTime
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetSlsProject() *string {
	return s.SlsProject
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreateTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDescription(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Description = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDisplayName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.DisplayName = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetLastModifyTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.LastModifyTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetRegionId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.RegionId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetSlsProject(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.SlsProject = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) Validate() error {
	return dara.Validate(s)
}
