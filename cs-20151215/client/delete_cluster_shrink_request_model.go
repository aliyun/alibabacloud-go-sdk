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
	// The options for deleting the resources that are associated with the cluster.
	DeleteOptionsShrink *string `json:"delete_options,omitempty" xml:"delete_options,omitempty"`
	// Deprecated
	//
	// Whether to retain SLB resources. Valid values:
	//
	// - `true`: Retains the SLB resources that are created for the cluster.
	//
	// - `false`: Does not retain the SLB resources that are created for the cluster.
	//
	// Default value: `false`.
	//
	// Use the `delete_options` parameter to manage `SLB` resources instead.
	//
	// example:
	//
	// false
	KeepSlb *bool `json:"keep_slb,omitempty" xml:"keep_slb,omitempty"`
	// Whether to retain all associated resources. If you set this parameter to `true`, the `retain_resources` parameter is ignored, and all cloud resources that are created with the cluster and can be queried by calling `DescribeClusterResources` are retained. If you set this parameter to `false`, note that resources that are configured to be retained by default in the `delete_options` parameter are still retained. To delete these resources, you must explicitly set the `delete_mode` parameter to `delete` for them in `delete_options`.
	//
	// - `true`: Retains all associated cloud resources that are created with the cluster.
	//
	// - `false`: Does not retain all associated cloud resources. Resources that are configured to be retained by default in the `delete_options` parameter, such as `ALB`, are still retained when this parameter is set to `false`.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	RetainAllResources *bool `json:"retain_all_resources,omitempty" xml:"retain_all_resources,omitempty"`
	// The IDs of resources to retain when the cluster is deleted.
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
