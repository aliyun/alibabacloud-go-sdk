// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGovernanceKubernetesClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ModifyGovernanceKubernetesClusterRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *ModifyGovernanceKubernetesClusterRequest
	GetClusterId() *string
	SetNamespaceInfos(v []*ModifyGovernanceKubernetesClusterRequestNamespaceInfos) *ModifyGovernanceKubernetesClusterRequest
	GetNamespaceInfos() []*ModifyGovernanceKubernetesClusterRequestNamespaceInfos
	SetRegionId(v string) *ModifyGovernanceKubernetesClusterRequest
	GetRegionId() *string
}

type ModifyGovernanceKubernetesClusterRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cd23228b3c80c4d4f9ad87cc3****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the namespace for which Microservices Engine(MSE) Microservices Governance is enabled.
	NamespaceInfos []*ModifyGovernanceKubernetesClusterRequestNamespaceInfos `json:"NamespaceInfos,omitempty" xml:"NamespaceInfos,omitempty" type:"Repeated"`
	// The ID of the region in which the instance resides. The region is supported by MSE.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyGovernanceKubernetesClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGovernanceKubernetesClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyGovernanceKubernetesClusterRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ModifyGovernanceKubernetesClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyGovernanceKubernetesClusterRequest) GetNamespaceInfos() []*ModifyGovernanceKubernetesClusterRequestNamespaceInfos {
	return s.NamespaceInfos
}

func (s *ModifyGovernanceKubernetesClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGovernanceKubernetesClusterRequest) SetAcceptLanguage(v string) *ModifyGovernanceKubernetesClusterRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterRequest) SetClusterId(v string) *ModifyGovernanceKubernetesClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterRequest) SetNamespaceInfos(v []*ModifyGovernanceKubernetesClusterRequestNamespaceInfos) *ModifyGovernanceKubernetesClusterRequest {
	s.NamespaceInfos = v
	return s
}

func (s *ModifyGovernanceKubernetesClusterRequest) SetRegionId(v string) *ModifyGovernanceKubernetesClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyGovernanceKubernetesClusterRequestNamespaceInfos struct {
	// The microservice namespace. If you do not specify this parameter, Microservice Governance is not enabled for the namespace.
	//
	// example:
	//
	// default
	MseNamespace *string `json:"MseNamespace,omitempty" xml:"MseNamespace,omitempty"`
	// The name of the Kubernetes namespace.
	//
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyGovernanceKubernetesClusterRequestNamespaceInfos) String() string {
	return dara.Prettify(s)
}

func (s ModifyGovernanceKubernetesClusterRequestNamespaceInfos) GoString() string {
	return s.String()
}

func (s *ModifyGovernanceKubernetesClusterRequestNamespaceInfos) GetMseNamespace() *string {
	return s.MseNamespace
}

func (s *ModifyGovernanceKubernetesClusterRequestNamespaceInfos) GetName() *string {
	return s.Name
}

func (s *ModifyGovernanceKubernetesClusterRequestNamespaceInfos) SetMseNamespace(v string) *ModifyGovernanceKubernetesClusterRequestNamespaceInfos {
	s.MseNamespace = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterRequestNamespaceInfos) SetName(v string) *ModifyGovernanceKubernetesClusterRequestNamespaceInfos {
	s.Name = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterRequestNamespaceInfos) Validate() error {
	return dara.Validate(s)
}
