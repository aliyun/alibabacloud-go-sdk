// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSpacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpaces(v []*ListAgentSpacesResponseBodyAgentSpaces) *ListAgentSpacesResponseBody
	GetAgentSpaces() []*ListAgentSpacesResponseBodyAgentSpaces
	SetMaxResults(v int32) *ListAgentSpacesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentSpacesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAgentSpacesResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListAgentSpacesResponseBody
	GetTotal() *int32
}

type ListAgentSpacesResponseBody struct {
	// The AgentSpaces information.
	AgentSpaces []*ListAgentSpacesResponseBodyAgentSpaces `json:"agentSpaces,omitempty" xml:"agentSpaces,omitempty" type:"Repeated"`
	// The maximum number of results returned.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. This parameter is empty if no more pages are available.
	//
	// example:
	//
	// b5754ef15c784abc8696d82790d2985c
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4E4AC775-2358-5B52-B6FB-171459D7B14B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 13
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAgentSpacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentSpacesResponseBody) GetAgentSpaces() []*ListAgentSpacesResponseBodyAgentSpaces {
	return s.AgentSpaces
}

func (s *ListAgentSpacesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentSpacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentSpacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentSpacesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAgentSpacesResponseBody) SetAgentSpaces(v []*ListAgentSpacesResponseBodyAgentSpaces) *ListAgentSpacesResponseBody {
	s.AgentSpaces = v
	return s
}

func (s *ListAgentSpacesResponseBody) SetMaxResults(v int32) *ListAgentSpacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAgentSpacesResponseBody) SetNextToken(v string) *ListAgentSpacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgentSpacesResponseBody) SetRequestId(v string) *ListAgentSpacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentSpacesResponseBody) SetTotal(v int32) *ListAgentSpacesResponseBody {
	s.Total = &v
	return s
}

func (s *ListAgentSpacesResponseBody) Validate() error {
	if s.AgentSpaces != nil {
		for _, item := range s.AgentSpaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentSpacesResponseBodyAgentSpaces struct {
	// The AgentSpace name.
	//
	// example:
	//
	// test-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The CMS workspace.
	//
	// example:
	//
	// test-cms-workspace
	CmsWorkspace *string `json:"cmsWorkspace,omitempty" xml:"cmsWorkspace,omitempty"`
	// The creation time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2023-08-23T04:06:06Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The MSE namespace.
	MseNamespace *ListAgentSpacesResponseBodyAgentSpacesMseNamespace `json:"mseNamespace,omitempty" xml:"mseNamespace,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The Simple Log Service project name.
	//
	// example:
	//
	// default-cms-1152309027070167-cn-beijing
	SlsProject *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
	// The update time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-02-11T08:40:23Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListAgentSpacesResponseBodyAgentSpaces) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpacesResponseBodyAgentSpaces) GoString() string {
	return s.String()
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) GetCmsWorkspace() *string {
	return s.CmsWorkspace
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) GetDescription() *string {
	return s.Description
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) GetMseNamespace() *ListAgentSpacesResponseBodyAgentSpacesMseNamespace {
	return s.MseNamespace
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) GetSlsProject() *string {
	return s.SlsProject
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) SetAgentSpace(v string) *ListAgentSpacesResponseBodyAgentSpaces {
	s.AgentSpace = &v
	return s
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) SetCmsWorkspace(v string) *ListAgentSpacesResponseBodyAgentSpaces {
	s.CmsWorkspace = &v
	return s
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) SetCreateTime(v string) *ListAgentSpacesResponseBodyAgentSpaces {
	s.CreateTime = &v
	return s
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) SetDescription(v string) *ListAgentSpacesResponseBodyAgentSpaces {
	s.Description = &v
	return s
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) SetMseNamespace(v *ListAgentSpacesResponseBodyAgentSpacesMseNamespace) *ListAgentSpacesResponseBodyAgentSpaces {
	s.MseNamespace = v
	return s
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) SetRegionId(v string) *ListAgentSpacesResponseBodyAgentSpaces {
	s.RegionId = &v
	return s
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) SetSlsProject(v string) *ListAgentSpacesResponseBodyAgentSpaces {
	s.SlsProject = &v
	return s
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) SetUpdateTime(v string) *ListAgentSpacesResponseBodyAgentSpaces {
	s.UpdateTime = &v
	return s
}

func (s *ListAgentSpacesResponseBodyAgentSpaces) Validate() error {
	if s.MseNamespace != nil {
		if err := s.MseNamespace.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentSpacesResponseBodyAgentSpacesMseNamespace struct {
	// The MSE namespace ID.
	//
	// example:
	//
	// phoenixcloud-raw-logs
	NamespaceId *string `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	// The MSE namespace name.
	//
	// example:
	//
	// terraform-alicloud-modules
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
}

func (s ListAgentSpacesResponseBodyAgentSpacesMseNamespace) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpacesResponseBodyAgentSpacesMseNamespace) GoString() string {
	return s.String()
}

func (s *ListAgentSpacesResponseBodyAgentSpacesMseNamespace) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListAgentSpacesResponseBodyAgentSpacesMseNamespace) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListAgentSpacesResponseBodyAgentSpacesMseNamespace) SetNamespaceId(v string) *ListAgentSpacesResponseBodyAgentSpacesMseNamespace {
	s.NamespaceId = &v
	return s
}

func (s *ListAgentSpacesResponseBodyAgentSpacesMseNamespace) SetNamespaceName(v string) *ListAgentSpacesResponseBodyAgentSpacesMseNamespace {
	s.NamespaceName = &v
	return s
}

func (s *ListAgentSpacesResponseBodyAgentSpacesMseNamespace) Validate() error {
	return dara.Validate(s)
}
