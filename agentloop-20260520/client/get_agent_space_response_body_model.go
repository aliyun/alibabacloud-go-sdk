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
	// The name of the AgentSpace.
	//
	// example:
	//
	// test-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// The CloudMonitor workspace.
	//
	// example:
	//
	// test-cms-workspace
	CmsWorkspace *string `json:"cmsWorkspace,omitempty" xml:"cmsWorkspace,omitempty"`
	// The binding type of the CloudMonitor 2.0 workspace.
	//
	// example:
	//
	// AutoCreated
	CmsWorkspaceBindType *string `json:"cmsWorkspaceBindType,omitempty" xml:"cmsWorkspaceBindType,omitempty"`
	// The time when the AgentSpace was created.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-05-05T12:39:36Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The Microservices Engine (MSE) namespace.
	MseNamespace *GetAgentSpaceResponseBodyMseNamespace `json:"mseNamespace,omitempty" xml:"mseNamespace,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D0173835-9E0F-508F-8BFA-9F556E59C302
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// default-cms-1837787111545040-cn-beijing
	SlsProject *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
	// The time when the AgentSpace was last updated.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2025-10-20T02:28:14Z
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
	// The ID of the MSE namespace.
	//
	// example:
	//
	// emr-dataware
	NamespaceId *string `json:"namespaceId,omitempty" xml:"namespaceId,omitempty"`
	// The name of the MSE namespace.
	//
	// example:
	//
	// terraform-alicloud-modules
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
