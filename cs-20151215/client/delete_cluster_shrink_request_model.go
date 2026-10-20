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
	// The deletion options for cluster-associated resources.
	DeleteOptionsShrink *string `json:"delete_options,omitempty" xml:"delete_options,omitempty"`
	// Deprecated
	//
	// Specifies whether to retain SLB instances. Valid values:
	//
	// - `true`: Retains the created SLB resources.
	//
	// - `false`: Does not retain the created SLB resources.
	//
	// Default value: `false`.
	//
	// Use `SLB` in `delete_options` to manage this setting.
	//
	// example:
	//
	// false
	KeepSlb *bool `json:"keep_slb,omitempty" xml:"keep_slb,omitempty"`
	// Specifies whether to retain all resources. If this parameter is set to `true`, `retain_resources` is ignored, and cloud resources created through the cluster that are queried by the `DescribeClusterResources` operation are retained. If this parameter is set to `false`, note that resources that are retained by default in `delete_options` are still retained. To delete these resources, set `delete_mode` to `delete` separately in `delete_options`.
	//
	// - `true`: Retains all resources, including all cloud resources created through the cluster.
	//
	// - `false`: Does not retain all resources (except resources defined as retained by default in `delete_options`. For example, `ALB` is still retained when this parameter is set to `false`).
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	RetainAllResources *bool `json:"retain_all_resources,omitempty" xml:"retain_all_resources,omitempty"`
	// The list of resources. To retain resources when deleting a cluster, specify the IDs of the corresponding resources.
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
