// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteOptions(v []*DeleteClusterRequestDeleteOptions) *DeleteClusterRequest
	GetDeleteOptions() []*DeleteClusterRequestDeleteOptions
	SetKeepSlb(v bool) *DeleteClusterRequest
	GetKeepSlb() *bool
	SetRetainAllResources(v bool) *DeleteClusterRequest
	GetRetainAllResources() *bool
	SetRetainResources(v []*string) *DeleteClusterRequest
	GetRetainResources() []*string
}

type DeleteClusterRequest struct {
	// The deletion options for the resources associated with the cluster.
	DeleteOptions []*DeleteClusterRequestDeleteOptions `json:"delete_options,omitempty" xml:"delete_options,omitempty" type:"Repeated"`
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
	RetainResources []*string `json:"retain_resources,omitempty" xml:"retain_resources,omitempty" type:"Repeated"`
}

func (s DeleteClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) GetDeleteOptions() []*DeleteClusterRequestDeleteOptions {
	return s.DeleteOptions
}

func (s *DeleteClusterRequest) GetKeepSlb() *bool {
	return s.KeepSlb
}

func (s *DeleteClusterRequest) GetRetainAllResources() *bool {
	return s.RetainAllResources
}

func (s *DeleteClusterRequest) GetRetainResources() []*string {
	return s.RetainResources
}

func (s *DeleteClusterRequest) SetDeleteOptions(v []*DeleteClusterRequestDeleteOptions) *DeleteClusterRequest {
	s.DeleteOptions = v
	return s
}

func (s *DeleteClusterRequest) SetKeepSlb(v bool) *DeleteClusterRequest {
	s.KeepSlb = &v
	return s
}

func (s *DeleteClusterRequest) SetRetainAllResources(v bool) *DeleteClusterRequest {
	s.RetainAllResources = &v
	return s
}

func (s *DeleteClusterRequest) SetRetainResources(v []*string) *DeleteClusterRequest {
	s.RetainResources = v
	return s
}

func (s *DeleteClusterRequest) Validate() error {
	if s.DeleteOptions != nil {
		for _, item := range s.DeleteOptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteClusterRequestDeleteOptions struct {
	// The deletion mode for this type of resource. Valid values:
	//
	// - `delete`: deletes this type of resource.
	//
	// - `retain`: retains this type of resource.
	//
	// example:
	//
	// delete
	DeleteMode *string `json:"delete_mode,omitempty" xml:"delete_mode,omitempty"`
	// The type of the resource. Valid values:
	//
	// - `SLB`: the SLB resources created by using Services. By default, these resources are deleted. You can choose to retain them.
	//
	// - `ALB`: the ALB resources created by the ALB Ingress controller. By default, these resources are retained. You can choose to delete them.
	//
	// - `SLS_Data`: the SLS projects used for the cluster log feature. By default, these resources are retained. You can choose to delete them.
	//
	// - `SLS_ControlPlane`: the SLS projects used for the control plane logs of managed clusters. By default, these resources are retained. You can choose to delete them.
	//
	// - `PrivateZone`: the PrivateZone resources created by ACK Serverless clusters. By default, these resources are retained. You can choose to delete them.
	//
	// example:
	//
	// SLS_Data
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
}

func (s DeleteClusterRequestDeleteOptions) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterRequestDeleteOptions) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequestDeleteOptions) GetDeleteMode() *string {
	return s.DeleteMode
}

func (s *DeleteClusterRequestDeleteOptions) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteClusterRequestDeleteOptions) SetDeleteMode(v string) *DeleteClusterRequestDeleteOptions {
	s.DeleteMode = &v
	return s
}

func (s *DeleteClusterRequestDeleteOptions) SetResourceType(v string) *DeleteClusterRequestDeleteOptions {
	s.ResourceType = &v
	return s
}

func (s *DeleteClusterRequestDeleteOptions) Validate() error {
	return dara.Validate(s)
}
