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
	// The deletion options for cluster-associated resources.
	DeleteOptions []*DeleteClusterRequestDeleteOptions `json:"delete_options,omitempty" xml:"delete_options,omitempty" type:"Repeated"`
	// Deprecated
	//
	// Specifies whether to retain SLB resources. Valid values:
	//
	// - `true`: retains the created SLB resources.
	//
	// - `false`: does not retain the created SLB resources.
	//
	// Default value: `false`.
	//
	// Use `SLB` in `delete_options` to manage this setting.
	//
	// example:
	//
	// false
	KeepSlb *bool `json:"keep_slb,omitempty" xml:"keep_slb,omitempty"`
	// Specifies whether to retain all resources. If this parameter is set to `true`, `retain_resources` is ignored, and cloud resources created through the cluster that are queried by the `DescribeClusterResources` operation are retained. If this parameter is set to `false`, resources that are retained by default in `delete_options` are still retained. To delete these resources, set `delete_mode` to `delete` in `delete_options`.
	//
	// - `true`: retains all resources, including all cloud resources created through the cluster.
	//
	// - `false`: does not retain all resources, except resources defined as retained by default in `delete_options`. For example, `ALB` resources are still retained when this parameter is set to `false`.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	RetainAllResources *bool `json:"retain_all_resources,omitempty" xml:"retain_all_resources,omitempty"`
	// The resource list. To retain resources when you delete a cluster, specify the IDs of the resources to retain.
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
	// The deletion policy for this resource type. Valid values:
	//
	// - delete: deletes resources of this type.
	//
	// - retain: retains resources of this type.
	//
	// example:
	//
	// delete
	DeleteMode *string `json:"delete_mode,omitempty" xml:"delete_mode,omitempty"`
	// The resource type. Valid values:
	//
	// - SLB: SLB resources created through services. Deleted by default. You can choose to retain them.
	//
	// - ALB: ALB resources created by ALB Ingress Controller. Retained by default. You can choose to delete them.
	//
	// - SLS_Data: the SLS project used by the cluster logging feature. Retained by default. You can choose to delete it.
	//
	// - SLS_ControlPlane: the SLS project used by the control plane logs of managed clusters. Retained by default. You can choose to delete it.
	//
	// - PrivateZone: PrivateZone resources created by ACK Serverless clusters. Retained by default. You can choose to delete them.
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
