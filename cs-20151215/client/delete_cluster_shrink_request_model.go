// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteOptionsShrink(v string) *DeleteClusterShrinkRequest
	GetDeleteOptionsShrink() *string
	SetKeepSlb(v bool) *DeleteClusterShrinkRequest
	GetKeepSlb() *bool
	SetRetainAllResources(v bool) *DeleteClusterShrinkRequest
	GetRetainAllResources() *bool
	SetRetainResourcesShrink(v string) *DeleteClusterShrinkRequest
	GetRetainResourcesShrink() *string
}

type DeleteClusterShrinkRequest struct {
	// The deletion options for the resources associated with the cluster.
	DeleteOptionsShrink *string `json:"delete_options,omitempty" xml:"delete_options,omitempty"`
	// Deprecated
	//
	// Specifies whether to retain the Server Load Balancer (SLB) instances created by using Services.
	//
	// Valid values:
	//
	// - `true`: retains the created SLB resources.
	//
	// - `false`: does not retain the created SLB resources.
	//
	// Default value: `false`.
	//
	// **[Deprecated]*	- Use `SLB` in `delete_options` instead.
	//
	// example:
	//
	// false
	KeepSlb *bool `json:"keep_slb,omitempty" xml:"keep_slb,omitempty"`
	// Specifies whether to retain all resources. If you set this parameter to `true`, the `retain_resources` parameter is ignored, and all cloud resources created by the cluster that are queried by calling the `DescribeClusterResources` operation are retained. If you set this parameter to `false`, note that the resources that are retained by default in `delete_options` are still retained. To delete these resources, set `delete_mode` to `delete` in `delete_options`.
	//
	// Valid values:
	//
	// - `true`: retains all resources, including all cloud resources created by the cluster.
	//
	// - `false`: does not retain all resources, except for the resources that are retained by default in `delete_options`. For example, ALB resources are still retained even if this parameter is set to `false`.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	RetainAllResources *bool `json:"retain_all_resources,omitempty" xml:"retain_all_resources,omitempty"`
	// The list of resources. To retain resources when you delete the cluster, specify the IDs of the corresponding resources.
	RetainResourcesShrink *string `json:"retain_resources,omitempty" xml:"retain_resources,omitempty"`
}

func (s DeleteClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterShrinkRequest) GetDeleteOptionsShrink() *string {
	return s.DeleteOptionsShrink
}

func (s *DeleteClusterShrinkRequest) GetKeepSlb() *bool {
	return s.KeepSlb
}

func (s *DeleteClusterShrinkRequest) GetRetainAllResources() *bool {
	return s.RetainAllResources
}

func (s *DeleteClusterShrinkRequest) GetRetainResourcesShrink() *string {
	return s.RetainResourcesShrink
}

func (s *DeleteClusterShrinkRequest) SetDeleteOptionsShrink(v string) *DeleteClusterShrinkRequest {
	s.DeleteOptionsShrink = &v
	return s
}

func (s *DeleteClusterShrinkRequest) SetKeepSlb(v bool) *DeleteClusterShrinkRequest {
	s.KeepSlb = &v
	return s
}

func (s *DeleteClusterShrinkRequest) SetRetainAllResources(v bool) *DeleteClusterShrinkRequest {
	s.RetainAllResources = &v
	return s
}

func (s *DeleteClusterShrinkRequest) SetRetainResourcesShrink(v string) *DeleteClusterShrinkRequest {
	s.RetainResourcesShrink = &v
	return s
}

func (s *DeleteClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
