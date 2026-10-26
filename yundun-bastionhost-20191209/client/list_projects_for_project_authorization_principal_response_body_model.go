// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsForProjectAuthorizationPrincipalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProjectsForProjectAuthorizationPrincipalResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListProjectsForProjectAuthorizationPrincipalResponseBody
	GetNextToken() *string
	SetProjects(v []*ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) *ListProjectsForProjectAuthorizationPrincipalResponseBody
	GetProjects() []*ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects
	SetRequestId(v string) *ListProjectsForProjectAuthorizationPrincipalResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProjectsForProjectAuthorizationPrincipalResponseBody
	GetTotalCount() *int32
}

type ListProjectsForProjectAuthorizationPrincipalResponseBody struct {
	// The maximum number of entries returned per page. If a valid MaxResults request parameter is passed, the specified value is returned. Otherwise, the default value is returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next request, used for paging. You can specify this token in the NextToken request parameter for the next paged query. If this field is empty in the response, no more data is available.
	//
	// example:
	//
	// 4ieSWJCwxvW3dk3wF.BqkrZmP72nWu5zJ5NWydMqyEs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of projects.
	Projects []*ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of queried projects.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectsForProjectAuthorizationPrincipalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsForProjectAuthorizationPrincipalResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBody) GetProjects() []*ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects {
	return s.Projects
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBody) SetMaxResults(v int32) *ListProjectsForProjectAuthorizationPrincipalResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBody) SetNextToken(v string) *ListProjectsForProjectAuthorizationPrincipalResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBody) SetProjects(v []*ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) *ListProjectsForProjectAuthorizationPrincipalResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBody) SetRequestId(v string) *ListProjectsForProjectAuthorizationPrincipalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBody) SetTotalCount(v int32) *ListProjectsForProjectAuthorizationPrincipalResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBody) Validate() error {
	if s.Projects != nil {
		for _, item := range s.Projects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects struct {
	// The description of the project.
	//
	// example:
	//
	// base
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the project is the default project. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the project.
	//
	// example:
	//
	// 2
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) GetDescription() *string {
	return s.Description
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) GetName() *string {
	return s.Name
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) SetDescription(v string) *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) SetIsDefault(v bool) *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects {
	s.IsDefault = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) SetName(v string) *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) SetProjectId(v int64) *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects {
	s.ProjectId = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponseBodyProjects) Validate() error {
	return dara.Validate(s)
}
