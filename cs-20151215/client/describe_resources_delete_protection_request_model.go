// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcesDeleteProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *DescribeResourcesDeleteProtectionRequest
	GetNamespace() *string
	SetResources(v string) *DescribeResourcesDeleteProtectionRequest
	GetResources() *string
}

type DescribeResourcesDeleteProtectionRequest struct {
	// The namespace of the resource to query.
	//
	// This parameter is required when resource_type is set to services. If this parameter is not specified, the namespace defaults to default.
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The name of the resource to query. Separate multiple resources with commas (,).
	//
	// - If resource_type is set to namespaces, set this parameter to namespace names. If this parameter is not specified, the deletion protection status of all namespaces in the cluster is queried.
	//
	// - If resource_type is set to services, this parameter is required. Set this parameter to service names.
	//
	// example:
	//
	// test1,test2
	Resources *string `json:"resources,omitempty" xml:"resources,omitempty"`
}

func (s DescribeResourcesDeleteProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcesDeleteProtectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcesDeleteProtectionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeResourcesDeleteProtectionRequest) GetResources() *string {
	return s.Resources
}

func (s *DescribeResourcesDeleteProtectionRequest) SetNamespace(v string) *DescribeResourcesDeleteProtectionRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeResourcesDeleteProtectionRequest) SetResources(v string) *DescribeResourcesDeleteProtectionRequest {
	s.Resources = &v
	return s
}

func (s *DescribeResourcesDeleteProtectionRequest) Validate() error {
	return dara.Validate(s)
}
