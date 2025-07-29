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
	// The type of cluster resource that you want to delete or retain.
	DeleteOptionsShrink *string `json:"delete_options,omitempty" xml:"delete_options,omitempty"`
	// Deprecated
	//
	// Specifies whether to retain the Server Load Balancer (SLB) resources that are created by the cluster.
	//
	// 	- `true`: retains the SLB instances that are created by the cluster.
	//
	// 	- `false`: does not retain the SLB instances that are created by the cluster.
	//
	// Default value: `false`. Set resource_type to `SLB` in the `delete_options` parameter to manage SLB instances.
	//
	// example:
	//
	// false
	KeepSlb *bool `json:"keep_slb,omitempty" xml:"keep_slb,omitempty"`
	// Specifies whether to retain all resources. If you set the parameter to `true`, the `retain_resources` parameter is ignored. The cloud resources that are created by the cluster are retained. You can call the `DescribeClusterResources` operation to query cloud resources created by the cluster. If you set the parameter to `false`, resources to be retained by default in the `delete_options` parameter are still retained. To delete these resources, set `delete_mode` to `delete` in `delete_options`.
	//
	// 	- `true`: retains all resources, including cloud resources created by the cluster.
	//
	// 	- `false`: does not retain all resources. Resources to be retained by default in the `delete_options` parameter are retained. For example, `ALB` instances are retained when this parameter is set to `false`.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	RetainAllResources *bool `json:"retain_all_resources,omitempty" xml:"retain_all_resources,omitempty"`
	// The list of resources. To retain resources when you delete a cluster, you need to specify the IDs of the resources to be retained.
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
