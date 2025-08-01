// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGovernanceKubernetesClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ModifyGovernanceKubernetesClusterShrinkRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *ModifyGovernanceKubernetesClusterShrinkRequest
	GetClusterId() *string
	SetNamespaceInfosShrink(v string) *ModifyGovernanceKubernetesClusterShrinkRequest
	GetNamespaceInfosShrink() *string
	SetRegionId(v string) *ModifyGovernanceKubernetesClusterShrinkRequest
	GetRegionId() *string
}

type ModifyGovernanceKubernetesClusterShrinkRequest struct {
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
	NamespaceInfosShrink *string `json:"NamespaceInfos,omitempty" xml:"NamespaceInfos,omitempty"`
	// The ID of the region in which the instance resides. The region is supported by MSE.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyGovernanceKubernetesClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGovernanceKubernetesClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyGovernanceKubernetesClusterShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ModifyGovernanceKubernetesClusterShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyGovernanceKubernetesClusterShrinkRequest) GetNamespaceInfosShrink() *string {
	return s.NamespaceInfosShrink
}

func (s *ModifyGovernanceKubernetesClusterShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGovernanceKubernetesClusterShrinkRequest) SetAcceptLanguage(v string) *ModifyGovernanceKubernetesClusterShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterShrinkRequest) SetClusterId(v string) *ModifyGovernanceKubernetesClusterShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterShrinkRequest) SetNamespaceInfosShrink(v string) *ModifyGovernanceKubernetesClusterShrinkRequest {
	s.NamespaceInfosShrink = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterShrinkRequest) SetRegionId(v string) *ModifyGovernanceKubernetesClusterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
