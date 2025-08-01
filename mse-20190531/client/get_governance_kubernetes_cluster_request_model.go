// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGovernanceKubernetesClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetGovernanceKubernetesClusterRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *GetGovernanceKubernetesClusterRequest
	GetClusterId() *string
	SetRegionId(v string) *GetGovernanceKubernetesClusterRequest
	GetRegionId() *string
}

type GetGovernanceKubernetesClusterRequest struct {
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
	// cd23228b3c80c4d4f9ad7af1d87cc****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the region in which the instance resides. The region is supported by MSE.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetGovernanceKubernetesClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceKubernetesClusterRequest) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetGovernanceKubernetesClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetGovernanceKubernetesClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetGovernanceKubernetesClusterRequest) SetAcceptLanguage(v string) *GetGovernanceKubernetesClusterRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetGovernanceKubernetesClusterRequest) SetClusterId(v string) *GetGovernanceKubernetesClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *GetGovernanceKubernetesClusterRequest) SetRegionId(v string) *GetGovernanceKubernetesClusterRequest {
	s.RegionId = &v
	return s
}

func (s *GetGovernanceKubernetesClusterRequest) Validate() error {
	return dara.Validate(s)
}
