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
	// The type of cluster resource that you want to delete or retain.
	DeleteOptions []*DeleteClusterRequestDeleteOptions `json:"delete_options,omitempty" xml:"delete_options,omitempty" type:"Repeated"`
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
	// The deletion policy for the specified type of resource. Valid values:
	//
	// 	- delete: deletes the specified type of resource.
	//
	// 	- retain: retains the specified type of resource.
	//
	// example:
	//
	// delete
	DeleteMode *string `json:"delete_mode,omitempty" xml:"delete_mode,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- SLB: SLB resources created for Services. By default, the SLB resources are automatically deleted.
	//
	// 	- ALB: Application Load Balancer (ALB) resources created by the ALB Ingress controller. By default, the ALB resources are retained.
	//
	// 	- SLS_Data: Simple Log Service projects used by the cluster logging feature. By default, the Simple Log Service projects are retained.
	//
	// 	- SLS_ControlPlane: Simple Log Service projects used to store the logs of control planes in ACK managed clusters. By default, the Simple Log Service projects are retained.
	//
	// 	- PrivateZone: PrivateZone resources created by ACK Serverless clusters. By default, the PrivateZone resources are retained.
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
