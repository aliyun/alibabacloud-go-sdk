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
	// The options for deleting the resources that are associated with the cluster.
	DeleteOptions []*DeleteClusterRequestDeleteOptions `json:"delete_options,omitempty" xml:"delete_options,omitempty" type:"Repeated"`
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
	// The deletion policy for the specified resource type. Valid values:
	//
	// - delete: Deletes the resources.
	//
	// - retain: Retains the resources.
	//
	// example:
	//
	// delete
	DeleteMode *string `json:"delete_mode,omitempty" xml:"delete_mode,omitempty"`
	// The type of resource. Valid values:
	//
	// - SLB: the SLB resources created for Services. These resources are deleted by default, but you can choose to retain them.
	//
	// - ALB: the ALB resources created by the ALB Ingress controller. These resources are retained by default, but you can choose to delete them.
	//
	// - SLS_Data: the SLS project used for cluster logs. This resource is retained by default, but you can choose to delete it.
	//
	// - SLS_ControlPlane: the SLS project used for control plane logs in a managed cluster. This resource is retained by default, but you can choose to delete it.
	//
	// - PrivateZone: the PrivateZone resource created by an ACK Serverless cluster. This resource is retained by default, but you can choose to delete it.
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
