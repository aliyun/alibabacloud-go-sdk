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
	AgentSpaces []*ListAgentSpacesResponseBodyAgentSpaces `json:"agentSpaces,omitempty" xml:"agentSpaces,omitempty" type:"Repeated"`
	MaxResults  *int32                                    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string                                   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId   *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Total       *int32                                    `json:"total,omitempty" xml:"total,omitempty"`
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
	AgentSpace   *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	CmsWorkspace *string `json:"cmsWorkspace,omitempty" xml:"cmsWorkspace,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	CreateTime   *string                                             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description  *string                                             `json:"description,omitempty" xml:"description,omitempty"`
	MseNamespace *ListAgentSpacesResponseBodyAgentSpacesMseNamespace `json:"mseNamespace,omitempty" xml:"mseNamespace,omitempty" type:"Struct"`
	RegionId     *string                                             `json:"regionId,omitempty" xml:"regionId,omitempty"`
	SlsProject   *string                                             `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
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
	NamespaceId   *string `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
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
