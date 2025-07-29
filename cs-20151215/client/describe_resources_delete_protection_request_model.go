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
	// The namespace in which the resources that you want to query reside.
	//
	// This parameter is required when you set resource_type to services. Default value: default.
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The names of the resources that you want to query. Separate multiple resource names with commas (,).
	//
	// 	- When you set resource_type to namespaces, you must specify namespace names. If you leave this parameter empty, all namespaces in the cluster are queried.
	//
	// 	- If you set resource_type to services, you must specify Service names.
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
