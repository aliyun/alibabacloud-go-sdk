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
	// The namespace to which the resource to query belongs.
	//
	// This parameter is required when ResourceType is set to services. If this parameter is not specified, the default namespace is used.
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The names of the resources to query. Separate multiple resource names with commas (,).
	//
	// - If ResourceType is set to namespaces, specify the namespace names. If this parameter is left empty, the deletion protection status of all namespaces in the cluster is queried.
	//
	// - If ResourceType is set to services, this parameter is required. Specify the service names.
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
