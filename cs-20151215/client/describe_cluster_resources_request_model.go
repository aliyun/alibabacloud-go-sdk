// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWithAddonResources(v bool) *DescribeClusterResourcesRequest
	GetWithAddonResources() *bool
}

type DescribeClusterResourcesRequest struct {
	// Specifies whether to query the resources created by cluster components.
	//
	// example:
	//
	// false
	WithAddonResources *bool `json:"with_addon_resources,omitempty" xml:"with_addon_resources,omitempty"`
}

func (s DescribeClusterResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourcesRequest) GetWithAddonResources() *bool {
	return s.WithAddonResources
}

func (s *DescribeClusterResourcesRequest) SetWithAddonResources(v bool) *DescribeClusterResourcesRequest {
	s.WithAddonResources = &v
	return s
}

func (s *DescribeClusterResourcesRequest) Validate() error {
	return dara.Validate(s)
}
