// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *GetAgentSpaceResponseBody
	GetAgentSpace() *string
	SetCmsWorkspace(v string) *GetAgentSpaceResponseBody
	GetCmsWorkspace() *string
	SetCmsWorkspaceBindType(v string) *GetAgentSpaceResponseBody
	GetCmsWorkspaceBindType() *string
	SetCreateTime(v string) *GetAgentSpaceResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetAgentSpaceResponseBody
	GetDescription() *string
	SetMseNamespace(v *GetAgentSpaceResponseBodyMseNamespace) *GetAgentSpaceResponseBody
	GetMseNamespace() *GetAgentSpaceResponseBodyMseNamespace
	SetRegionId(v string) *GetAgentSpaceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetAgentSpaceResponseBody
	GetRequestId() *string
	SetSlsProject(v string) *GetAgentSpaceResponseBody
	GetSlsProject() *string
	SetUpdateTime(v string) *GetAgentSpaceResponseBody
	GetUpdateTime() *string
}

type GetAgentSpaceResponseBody struct {
	AgentSpace           *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	CmsWorkspace         *string `json:"cmsWorkspace,omitempty" xml:"cmsWorkspace,omitempty"`
	CmsWorkspaceBindType *string `json:"cmsWorkspaceBindType,omitempty" xml:"cmsWorkspaceBindType,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	CreateTime   *string                                `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description  *string                                `json:"description,omitempty" xml:"description,omitempty"`
	MseNamespace *GetAgentSpaceResponseBodyMseNamespace `json:"mseNamespace,omitempty" xml:"mseNamespace,omitempty" type:"Struct"`
	RegionId     *string                                `json:"regionId,omitempty" xml:"regionId,omitempty"`
	RequestId    *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SlsProject   *string                                `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetAgentSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentSpaceResponseBody) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetAgentSpaceResponseBody) GetCmsWorkspace() *string {
	return s.CmsWorkspace
}

func (s *GetAgentSpaceResponseBody) GetCmsWorkspaceBindType() *string {
	return s.CmsWorkspaceBindType
}

func (s *GetAgentSpaceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAgentSpaceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetAgentSpaceResponseBody) GetMseNamespace() *GetAgentSpaceResponseBodyMseNamespace {
	return s.MseNamespace
}

func (s *GetAgentSpaceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAgentSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentSpaceResponseBody) GetSlsProject() *string {
	return s.SlsProject
}

func (s *GetAgentSpaceResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAgentSpaceResponseBody) SetAgentSpace(v string) *GetAgentSpaceResponseBody {
	s.AgentSpace = &v
	return s
}

func (s *GetAgentSpaceResponseBody) SetCmsWorkspace(v string) *GetAgentSpaceResponseBody {
	s.CmsWorkspace = &v
	return s
}

func (s *GetAgentSpaceResponseBody) SetCmsWorkspaceBindType(v string) *GetAgentSpaceResponseBody {
	s.CmsWorkspaceBindType = &v
	return s
}

func (s *GetAgentSpaceResponseBody) SetCreateTime(v string) *GetAgentSpaceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetAgentSpaceResponseBody) SetDescription(v string) *GetAgentSpaceResponseBody {
	s.Description = &v
	return s
}

func (s *GetAgentSpaceResponseBody) SetMseNamespace(v *GetAgentSpaceResponseBodyMseNamespace) *GetAgentSpaceResponseBody {
	s.MseNamespace = v
	return s
}

func (s *GetAgentSpaceResponseBody) SetRegionId(v string) *GetAgentSpaceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetAgentSpaceResponseBody) SetRequestId(v string) *GetAgentSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentSpaceResponseBody) SetSlsProject(v string) *GetAgentSpaceResponseBody {
	s.SlsProject = &v
	return s
}

func (s *GetAgentSpaceResponseBody) SetUpdateTime(v string) *GetAgentSpaceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetAgentSpaceResponseBody) Validate() error {
	if s.MseNamespace != nil {
		if err := s.MseNamespace.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentSpaceResponseBodyMseNamespace struct {
	NamespaceId   *string `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
}

func (s GetAgentSpaceResponseBodyMseNamespace) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpaceResponseBodyMseNamespace) GoString() string {
	return s.String()
}

func (s *GetAgentSpaceResponseBodyMseNamespace) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetAgentSpaceResponseBodyMseNamespace) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *GetAgentSpaceResponseBodyMseNamespace) SetNamespaceId(v string) *GetAgentSpaceResponseBodyMseNamespace {
	s.NamespaceId = &v
	return s
}

func (s *GetAgentSpaceResponseBodyMseNamespace) SetNamespaceName(v string) *GetAgentSpaceResponseBodyMseNamespace {
	s.NamespaceName = &v
	return s
}

func (s *GetAgentSpaceResponseBodyMseNamespace) Validate() error {
	return dara.Validate(s)
}
