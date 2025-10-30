// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterResourcesResponse
	GetStatusCode() *int32
	SetBody(v []*DescribeClusterResourcesResponseBody) *DescribeClusterResourcesResponse
	GetBody() []*DescribeClusterResourcesResponseBody
}

type DescribeClusterResourcesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*DescribeClusterResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s DescribeClusterResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterResourcesResponse) GetBody() []*DescribeClusterResourcesResponseBody {
	return s.Body
}

func (s *DescribeClusterResourcesResponse) SetHeaders(v map[string]*string) *DescribeClusterResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResourcesResponse) SetStatusCode(v int32) *DescribeClusterResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterResourcesResponse) SetBody(v []*DescribeClusterResourcesResponseBody) *DescribeClusterResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterResourcesResponse) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterResourcesResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// cb95aa626a47740afbf6aa099b65****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2020-09-11T10:11:54+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// lb-wz9poz4r0ymh8u0uf****
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// The resource information. For more information about how to query the source information about the resource, see [ListStackResources](https://help.aliyun.com/document_detail/133836.html).
	//
	// example:
	//
	// {\\"Id\\":\\"k8s_master_slb\\",\\"Name\\":\\"k8s_master_slb\\",\\"Type\\":\\"ALIYUN::SLB::LoadBalancer\\",\\"Status\\":\\"CREATE_COMPLETE\\",\\"StatusReason\\":\\"state changed\\",\\"Updated\\":\\"2020-05-21T13:25:02\\",\\"PhysicalId\\":\\"lb-wz9poz4r0ymh8u0uf****\\"}
	ResourceInfo *string `json:"resource_info,omitempty" xml:"resource_info,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::SLB::LoadBalancer
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The resource status. Valid values:
	//
	// 	- `CREATE_COMPLETE`: the resource is created.
	//
	// 	- `CREATE_FAILED`: the resource failed to be created.
	//
	// 	- `CREATE_IN_PROGRESS`: the resource is being created.
	//
	// 	- `DELETE_FAILED`: the resource failed to be deleted.
	//
	// 	- `DELETE_IN_PROGRESS`: the resource is being deleted.
	//
	// 	- `ROLLBACK_COMPLETE`: the resource is rolled back.
	//
	// 	- `ROLLBACK_FAILED`: the resource failed to be rolled back.
	//
	// 	- `ROLLBACK_IN_PROGRESS`: the resource is being rolled back.
	//
	// example:
	//
	// CREATE_COMPLETE
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// Specifies whether the resource is created by Container Service for Kubernetes (ACK). Valid values:
	//
	// 	- 1: the resource is created by ACK.
	//
	// 	- 0: the resource is an existing resource.
	//
	// example:
	//
	// 1
	AutoCreate *int64 `json:"auto_create,omitempty" xml:"auto_create,omitempty"`
	// The dependent resources.
	Dependencies []*DescribeClusterResourcesResponseBodyDependencies `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Repeated"`
	// The Kubernetes object with which the resource is associated.
	AssociatedObject *DescribeClusterResourcesResponseBodyAssociatedObject `json:"associated_object,omitempty" xml:"associated_object,omitempty" type:"Struct"`
	// The deletion behavior of the resource when the cluster is deleted.
	DeleteBehavior *DescribeClusterResourcesResponseBodyDeleteBehavior `json:"delete_behavior,omitempty" xml:"delete_behavior,omitempty" type:"Struct"`
	// The type of the resource creator. Valid values:
	//
	// 	- user: The resource is created by the user.
	//
	// 	- system: The resource is created by the ACK management system.
	//
	// 	- addon: The resource is created by a cluster component.
	//
	// example:
	//
	// addon
	CreatorType *string `json:"creator_type,omitempty" xml:"creator_type,omitempty"`
	// The additional information about the resource.
	//
	// example:
	//
	// {"IP": "xx.xx.xx.xx"}
	ExtraInfo map[string]interface{} `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
}

func (s DescribeClusterResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourcesResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterResourcesResponseBody) GetCreated() *string {
	return s.Created
}

func (s *DescribeClusterResourcesResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeClusterResourcesResponseBody) GetResourceInfo() *string {
	return s.ResourceInfo
}

func (s *DescribeClusterResourcesResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeClusterResourcesResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeClusterResourcesResponseBody) GetAutoCreate() *int64 {
	return s.AutoCreate
}

func (s *DescribeClusterResourcesResponseBody) GetDependencies() []*DescribeClusterResourcesResponseBodyDependencies {
	return s.Dependencies
}

func (s *DescribeClusterResourcesResponseBody) GetAssociatedObject() *DescribeClusterResourcesResponseBodyAssociatedObject {
	return s.AssociatedObject
}

func (s *DescribeClusterResourcesResponseBody) GetDeleteBehavior() *DescribeClusterResourcesResponseBodyDeleteBehavior {
	return s.DeleteBehavior
}

func (s *DescribeClusterResourcesResponseBody) GetCreatorType() *string {
	return s.CreatorType
}

func (s *DescribeClusterResourcesResponseBody) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *DescribeClusterResourcesResponseBody) SetClusterId(v string) *DescribeClusterResourcesResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetCreated(v string) *DescribeClusterResourcesResponseBody {
	s.Created = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetInstanceId(v string) *DescribeClusterResourcesResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetResourceInfo(v string) *DescribeClusterResourcesResponseBody {
	s.ResourceInfo = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetResourceType(v string) *DescribeClusterResourcesResponseBody {
	s.ResourceType = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetState(v string) *DescribeClusterResourcesResponseBody {
	s.State = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetAutoCreate(v int64) *DescribeClusterResourcesResponseBody {
	s.AutoCreate = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetDependencies(v []*DescribeClusterResourcesResponseBodyDependencies) *DescribeClusterResourcesResponseBody {
	s.Dependencies = v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetAssociatedObject(v *DescribeClusterResourcesResponseBodyAssociatedObject) *DescribeClusterResourcesResponseBody {
	s.AssociatedObject = v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetDeleteBehavior(v *DescribeClusterResourcesResponseBodyDeleteBehavior) *DescribeClusterResourcesResponseBody {
	s.DeleteBehavior = v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetCreatorType(v string) *DescribeClusterResourcesResponseBody {
	s.CreatorType = &v
	return s
}

func (s *DescribeClusterResourcesResponseBody) SetExtraInfo(v map[string]interface{}) *DescribeClusterResourcesResponseBody {
	s.ExtraInfo = v
	return s
}

func (s *DescribeClusterResourcesResponseBody) Validate() error {
	if s.Dependencies != nil {
		for _, item := range s.Dependencies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AssociatedObject != nil {
		if err := s.AssociatedObject.Validate(); err != nil {
			return err
		}
	}
	if s.DeleteBehavior != nil {
		if err := s.DeleteBehavior.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterResourcesResponseBodyDependencies struct {
	// The ID of the cluster to which the dependent resource is related.
	//
	// example:
	//
	// cc5ee03f63e43425cb6f71f1a1756****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The dependent resource type.
	//
	// example:
	//
	// ALIYUN::VPC::NatGateway
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The dependent resource ID.
	//
	// example:
	//
	// ngw-wz9sphwk42sdtjixo****
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
}

func (s DescribeClusterResourcesResponseBodyDependencies) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourcesResponseBodyDependencies) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourcesResponseBodyDependencies) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterResourcesResponseBodyDependencies) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeClusterResourcesResponseBodyDependencies) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeClusterResourcesResponseBodyDependencies) SetClusterId(v string) *DescribeClusterResourcesResponseBodyDependencies {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterResourcesResponseBodyDependencies) SetResourceType(v string) *DescribeClusterResourcesResponseBodyDependencies {
	s.ResourceType = &v
	return s
}

func (s *DescribeClusterResourcesResponseBodyDependencies) SetInstanceId(v string) *DescribeClusterResourcesResponseBodyDependencies {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterResourcesResponseBodyDependencies) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterResourcesResponseBodyAssociatedObject struct {
	// The Kubernetes object type.
	//
	// example:
	//
	// Service
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The namespace in which the Kubernetes object resides.
	//
	// example:
	//
	// kube-system
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The Kubernetes object name.
	//
	// example:
	//
	// nginx-ingress-lb
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeClusterResourcesResponseBodyAssociatedObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourcesResponseBodyAssociatedObject) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourcesResponseBodyAssociatedObject) GetKind() *string {
	return s.Kind
}

func (s *DescribeClusterResourcesResponseBodyAssociatedObject) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeClusterResourcesResponseBodyAssociatedObject) GetName() *string {
	return s.Name
}

func (s *DescribeClusterResourcesResponseBodyAssociatedObject) SetKind(v string) *DescribeClusterResourcesResponseBodyAssociatedObject {
	s.Kind = &v
	return s
}

func (s *DescribeClusterResourcesResponseBodyAssociatedObject) SetNamespace(v string) *DescribeClusterResourcesResponseBodyAssociatedObject {
	s.Namespace = &v
	return s
}

func (s *DescribeClusterResourcesResponseBodyAssociatedObject) SetName(v string) *DescribeClusterResourcesResponseBodyAssociatedObject {
	s.Name = &v
	return s
}

func (s *DescribeClusterResourcesResponseBodyAssociatedObject) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterResourcesResponseBodyDeleteBehavior struct {
	// Specifies whether to delete the resource by default when the cluster is deleted.
	//
	// example:
	//
	// false
	DeleteByDefault *bool `json:"delete_by_default,omitempty" xml:"delete_by_default,omitempty"`
	// Specifies whether the default behavior returned in delete_by_default can be changed.
	//
	// example:
	//
	// false
	Changeable *bool `json:"changeable,omitempty" xml:"changeable,omitempty"`
}

func (s DescribeClusterResourcesResponseBodyDeleteBehavior) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourcesResponseBodyDeleteBehavior) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourcesResponseBodyDeleteBehavior) GetDeleteByDefault() *bool {
	return s.DeleteByDefault
}

func (s *DescribeClusterResourcesResponseBodyDeleteBehavior) GetChangeable() *bool {
	return s.Changeable
}

func (s *DescribeClusterResourcesResponseBodyDeleteBehavior) SetDeleteByDefault(v bool) *DescribeClusterResourcesResponseBodyDeleteBehavior {
	s.DeleteByDefault = &v
	return s
}

func (s *DescribeClusterResourcesResponseBodyDeleteBehavior) SetChangeable(v bool) *DescribeClusterResourcesResponseBodyDeleteBehavior {
	s.Changeable = &v
	return s
}

func (s *DescribeClusterResourcesResponseBodyDeleteBehavior) Validate() error {
	return dara.Validate(s)
}
