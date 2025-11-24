// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddClusterIntoServiceMeshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AddClusterIntoServiceMeshRequest
	GetClusterId() *string
	SetDiscoveryOnly(v bool) *AddClusterIntoServiceMeshRequest
	GetDiscoveryOnly() *bool
	SetIgnoreNamespaceCheck(v bool) *AddClusterIntoServiceMeshRequest
	GetIgnoreNamespaceCheck() *bool
	SetKubeconfig(v string) *AddClusterIntoServiceMeshRequest
	GetKubeconfig() *string
	SetServiceMeshId(v string) *AddClusterIntoServiceMeshRequest
	GetServiceMeshId() *string
}

type AddClusterIntoServiceMeshRequest struct {
	// The ID of the cluster to be added.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce3c25e247da24f3aab9b7edfae83****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Indicates whether to add the cluster to an ASM instance in only service discovery mode.
	//
	// example:
	//
	// false
	DiscoveryOnly *bool `json:"DiscoveryOnly,omitempty" xml:"DiscoveryOnly,omitempty"`
	// Specifies whether to check that the cluster you want to add to the ASM instance belongs to the istio-system namespace. This parameter is often used in scenarios where you migrate traffic from self-managed open source Istio to ASM. Valid values: true and false.
	//
	// example:
	//
	// false
	IgnoreNamespaceCheck *bool `json:"IgnoreNamespaceCheck,omitempty" xml:"IgnoreNamespaceCheck,omitempty"`
	// The cluster certificate.
	//
	// example:
	//
	// apiVersion: v1 clusters: - cluster:     server: https://47.110.xx.xx:6443     certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURUakNDQWphZ0F3SUJBZ0lVYzBQVy82ejR1aHlxYkRRdnNsV1htSmpJeFdNd0RRWUpLb1pJaHZjTkFRRUwKQlFBd1BqRW5NQThHQTFVRUNoTUlhR0Z1WjNwb2IzVXdGQVlEVlFRS0V3MWhiR2xpWVdKaElHTnNiM1ZrTVJNdwpFUVlEVlFRREV3cHJkV0psY201bGRHVnpNQ0FYRFRJeU1EUXdOekExTVRnd01Gb1lEekl3TlRJd016TXdNRFV4Ck9EQXdXakErTVNjd0R3WURWUVFLRXdob1lXNW5lbWh2ZFRBVUJnTlZCQW9URFdGc2FXSmhZbUVnWTJ4dmRXUXgKRXpBUkJnTlZCQU1UQ210MVltVnlibVYwWlhNd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJE****
	Kubeconfig *string `json:"Kubeconfig,omitempty" xml:"Kubeconfig,omitempty"`
	// The ID of the Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s AddClusterIntoServiceMeshRequest) String() string {
	return dara.Prettify(s)
}

func (s AddClusterIntoServiceMeshRequest) GoString() string {
	return s.String()
}

func (s *AddClusterIntoServiceMeshRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddClusterIntoServiceMeshRequest) GetDiscoveryOnly() *bool {
	return s.DiscoveryOnly
}

func (s *AddClusterIntoServiceMeshRequest) GetIgnoreNamespaceCheck() *bool {
	return s.IgnoreNamespaceCheck
}

func (s *AddClusterIntoServiceMeshRequest) GetKubeconfig() *string {
	return s.Kubeconfig
}

func (s *AddClusterIntoServiceMeshRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *AddClusterIntoServiceMeshRequest) SetClusterId(v string) *AddClusterIntoServiceMeshRequest {
	s.ClusterId = &v
	return s
}

func (s *AddClusterIntoServiceMeshRequest) SetDiscoveryOnly(v bool) *AddClusterIntoServiceMeshRequest {
	s.DiscoveryOnly = &v
	return s
}

func (s *AddClusterIntoServiceMeshRequest) SetIgnoreNamespaceCheck(v bool) *AddClusterIntoServiceMeshRequest {
	s.IgnoreNamespaceCheck = &v
	return s
}

func (s *AddClusterIntoServiceMeshRequest) SetKubeconfig(v string) *AddClusterIntoServiceMeshRequest {
	s.Kubeconfig = &v
	return s
}

func (s *AddClusterIntoServiceMeshRequest) SetServiceMeshId(v string) *AddClusterIntoServiceMeshRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *AddClusterIntoServiceMeshRequest) Validate() error {
	return dara.Validate(s)
}
