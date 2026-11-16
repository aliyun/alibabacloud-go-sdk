// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWorkspacesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListWorkspacesResponseBody
	GetTotalCount() *int32
	SetWorkspaceList(v []*ListWorkspacesResponseBodyWorkspaceList) *ListWorkspacesResponseBody
	GetWorkspaceList() []*ListWorkspacesResponseBodyWorkspaceList
}

type ListWorkspacesResponseBody struct {
	// The request ID, which is used to locate and troubleshoot requests.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of workspaces that meet the query conditions.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of workspaces on the current page.
	WorkspaceList []*ListWorkspacesResponseBodyWorkspaceList `json:"WorkspaceList,omitempty" xml:"WorkspaceList,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspacesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWorkspacesResponseBody) GetWorkspaceList() []*ListWorkspacesResponseBodyWorkspaceList {
	return s.WorkspaceList
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalCount(v int32) *ListWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetWorkspaceList(v []*ListWorkspacesResponseBodyWorkspaceList) *ListWorkspacesResponseBody {
	s.WorkspaceList = v
	return s
}

func (s *ListWorkspacesResponseBody) Validate() error {
	if s.WorkspaceList != nil {
		for _, item := range s.WorkspaceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspacesResponseBodyWorkspaceList struct {
	// The workspace code.
	//
	// example:
	//
	// c40***
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The creation time. The time is in UTC and follows the ISO 8601 standard, ending with Z.
	//
	// example:
	//
	// 2026-01-07T02:21:36Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the default project for the workspace.
	//
	// example:
	//
	// pd_061716****
	DefaultProjectId *string `json:"DefaultProjectId,omitempty" xml:"DefaultProjectId,omitempty"`
	// The workspace status. Valid values: Online (enabled) and Offline (disabled).
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// Video Production Workspace
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The current number of users in the workspace, returned as a string.
	//
	// example:
	//
	// 2
	UserCount *string `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws_zna577pdximv***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaceList) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaceList) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaceList) GetCode() *string {
	return s.Code
}

func (s *ListWorkspacesResponseBodyWorkspaceList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWorkspacesResponseBodyWorkspaceList) GetDefaultProjectId() *string {
	return s.DefaultProjectId
}

func (s *ListWorkspacesResponseBodyWorkspaceList) GetStatus() *string {
	return s.Status
}

func (s *ListWorkspacesResponseBodyWorkspaceList) GetTitle() *string {
	return s.Title
}

func (s *ListWorkspacesResponseBodyWorkspaceList) GetUserCount() *string {
	return s.UserCount
}

func (s *ListWorkspacesResponseBodyWorkspaceList) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspacesResponseBodyWorkspaceList) SetCode(v string) *ListWorkspacesResponseBodyWorkspaceList {
	s.Code = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaceList) SetCreateTime(v string) *ListWorkspacesResponseBodyWorkspaceList {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaceList) SetDefaultProjectId(v string) *ListWorkspacesResponseBodyWorkspaceList {
	s.DefaultProjectId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaceList) SetStatus(v string) *ListWorkspacesResponseBodyWorkspaceList {
	s.Status = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaceList) SetTitle(v string) *ListWorkspacesResponseBodyWorkspaceList {
	s.Title = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaceList) SetUserCount(v string) *ListWorkspacesResponseBodyWorkspaceList {
	s.UserCount = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaceList) SetWorkspaceId(v string) *ListWorkspacesResponseBodyWorkspaceList {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaceList) Validate() error {
	return dara.Validate(s)
}
