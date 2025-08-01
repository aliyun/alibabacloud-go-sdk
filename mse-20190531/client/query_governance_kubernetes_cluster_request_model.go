// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGovernanceKubernetesClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryGovernanceKubernetesClusterRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *QueryGovernanceKubernetesClusterRequest
	GetClusterId() *string
	SetClusterName(v string) *QueryGovernanceKubernetesClusterRequest
	GetClusterName() *string
	SetPageNumber(v int32) *QueryGovernanceKubernetesClusterRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryGovernanceKubernetesClusterRequest
	GetPageSize() *int32
}

type QueryGovernanceKubernetesClusterRequest struct {
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
	// The ID of the Kubernetes cluster.
	//
	// example:
	//
	// c24c9354acxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the Kubernetes cluster.
	//
	// example:
	//
	// example-cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryGovernanceKubernetesClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGovernanceKubernetesClusterRequest) GoString() string {
	return s.String()
}

func (s *QueryGovernanceKubernetesClusterRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryGovernanceKubernetesClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryGovernanceKubernetesClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *QueryGovernanceKubernetesClusterRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryGovernanceKubernetesClusterRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryGovernanceKubernetesClusterRequest) SetAcceptLanguage(v string) *QueryGovernanceKubernetesClusterRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterRequest) SetClusterId(v string) *QueryGovernanceKubernetesClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterRequest) SetClusterName(v string) *QueryGovernanceKubernetesClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterRequest) SetPageNumber(v int32) *QueryGovernanceKubernetesClusterRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterRequest) SetPageSize(v int32) *QueryGovernanceKubernetesClusterRequest {
	s.PageSize = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterRequest) Validate() error {
	return dara.Validate(s)
}
