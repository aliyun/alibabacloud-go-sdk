// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProjectsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListProjectsResponseBody
	GetNextToken() *string
	SetProjectList(v []*ListProjectsResponseBodyProjectList) *ListProjectsResponseBody
	GetProjectList() []*ListProjectsResponseBodyProjectList
	SetRequestId(v string) *ListProjectsResponseBody
	GetRequestId() *string
}

type ListProjectsResponseBody struct {
	// A response field retained for compatibility. It is not used for current pagination and can be ignored.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A response field retained for compatibility. It is not used for current pagination and can be ignored.
	//
	// example:
	//
	// Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of projects on the current page.
	ProjectList []*ListProjectsResponseBodyProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Repeated"`
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProjectsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsResponseBody) GetProjectList() []*ListProjectsResponseBodyProjectList {
	return s.ProjectList
}

func (s *ListProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsResponseBody) SetMaxResults(v int32) *ListProjectsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsResponseBody) SetNextToken(v string) *ListProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBody) SetProjectList(v []*ListProjectsResponseBodyProjectList) *ListProjectsResponseBody {
	s.ProjectList = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) Validate() error {
	if s.ProjectList != nil {
		for _, item := range s.ProjectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsResponseBodyProjectList struct {
	// The permissions of the current caller in the project, such as Manage for management permissions.
	//
	// example:
	//
	// Manage
	Auth *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
	// The URL of the cover image.
	//
	// example:
	//
	// https://example.com/cover.png
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// The creation time in UTC format following the ISO 8601 standard, ending with Z.
	//
	// example:
	//
	// 2026-01-07T02:21:36Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The username of the user who created the project.
	//
	// example:
	//
	// aliyun_183320****
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// The project description.
	//
	// example:
	//
	// Product demo project
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The project ID.
	//
	// example:
	//
	// pd_061716****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// Product Demo
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws_zna577pdximv***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListProjectsResponseBodyProjectList) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyProjectList) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjectList) GetAuth() *string {
	return s.Auth
}

func (s *ListProjectsResponseBodyProjectList) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *ListProjectsResponseBodyProjectList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProjectsResponseBodyProjectList) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListProjectsResponseBodyProjectList) GetDescription() *string {
	return s.Description
}

func (s *ListProjectsResponseBodyProjectList) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListProjectsResponseBodyProjectList) GetTitle() *string {
	return s.Title
}

func (s *ListProjectsResponseBodyProjectList) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListProjectsResponseBodyProjectList) SetAuth(v string) *ListProjectsResponseBodyProjectList {
	s.Auth = &v
	return s
}

func (s *ListProjectsResponseBodyProjectList) SetCoverUrl(v string) *ListProjectsResponseBodyProjectList {
	s.CoverUrl = &v
	return s
}

func (s *ListProjectsResponseBodyProjectList) SetCreateTime(v string) *ListProjectsResponseBodyProjectList {
	s.CreateTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjectList) SetCreateUserName(v string) *ListProjectsResponseBodyProjectList {
	s.CreateUserName = &v
	return s
}

func (s *ListProjectsResponseBodyProjectList) SetDescription(v string) *ListProjectsResponseBodyProjectList {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyProjectList) SetProjectId(v string) *ListProjectsResponseBodyProjectList {
	s.ProjectId = &v
	return s
}

func (s *ListProjectsResponseBodyProjectList) SetTitle(v string) *ListProjectsResponseBodyProjectList {
	s.Title = &v
	return s
}

func (s *ListProjectsResponseBodyProjectList) SetWorkspaceId(v string) *ListProjectsResponseBodyProjectList {
	s.WorkspaceId = &v
	return s
}

func (s *ListProjectsResponseBodyProjectList) Validate() error {
	return dara.Validate(s)
}
