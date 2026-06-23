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
	// 集群ID。
	//
	// example:
	//
	// cb95aa626a47740afbf6aa099b65****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 资源创建时间。
	//
	// example:
	//
	// 2023-08-15T14:34:42+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// 资源ID。
	//
	// example:
	//
	// ngw-wz9sphwk42sdtjixo****
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// 资源信息。关于资源的源信息，请参见[ListStackResources](https://help.aliyun.com/document_detail/133836.html)。
	//
	// example:
	//
	// {\\"Id\\":\\"KubernetesWorkerRole\\",\\"Name\\":\\"KubernetesWorkerRole\\",\\"Type\\":\\"ALIYUN::RAM::Role\\",\\"Status\\":\\"CREATE_COMPLETE\\",\\"StatusReason\\":\\"state changed\\",\\"Updated\\":\\"2025-04-10T06:21:17\\",\\"PhysicalId\\":\\"KubernetesWorkerRole-7e611193-225f-40f6-bc3c-ea8633******\\"}
	ResourceInfo *string `json:"resource_info,omitempty" xml:"resource_info,omitempty"`
	// 资源类型。
	//
	// example:
	//
	// ALIYUN::VPC::NatGateway
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// 资源状态。可选值：
	//
	// - `CREATE_COMPLETE`：成功创建资源。
	//
	// - `CREATE_FAILED`：创建资源失败。
	//
	// - `CREATE_IN_PROGRESS`：创建资源中。
	//
	// - `DELETE_FAILED`：删除资源失败。
	//
	// - `DELETE_IN_PROGRESS`：删除资源中。
	//
	// - `ROLLBACK_COMPLETE`：成功回滚。
	//
	// - `ROLLBACK_FAILED`：回滚失败。
	//
	// - `ROLLBACK_IN_PROGRESS`：回滚中。
	//
	// example:
	//
	// CREATE_COMPLETE
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// 资源是否由ACK创建：
	//
	// - 1：表示由ACK创建。
	//
	// - 0：表示该资源为已有资源。
	//
	// example:
	//
	// 1
	AutoCreate *int64 `json:"auto_create,omitempty" xml:"auto_create,omitempty"`
	// 依赖资源列表。
	Dependencies []*DescribeClusterResourcesResponseBodyDependencies `json:"dependencies,omitempty" xml:"dependencies,omitempty" type:"Repeated"`
	// 资源关联的k8s对象。
	AssociatedObject *DescribeClusterResourcesResponseBodyAssociatedObject `json:"associated_object,omitempty" xml:"associated_object,omitempty" type:"Struct"`
	// 删除集群时该资源的删除行为。
	DeleteBehavior *DescribeClusterResourcesResponseBodyDeleteBehavior `json:"delete_behavior,omitempty" xml:"delete_behavior,omitempty" type:"Struct"`
	// 该资源创建者的类型。可能的取值：
	//
	// - user：由用户自行创建；
	//
	// - system：由ACK管控系统创建；
	//
	// - addon：由集群组件创建。
	//
	// example:
	//
	// addon
	CreatorType *string `json:"creator_type,omitempty" xml:"creator_type,omitempty"`
	// 资源的其他信息。
	//
	// example:
	//
	// { "type": "SLS_Data" }
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
	// 依赖资源的集群ID。
	//
	// example:
	//
	// cc5ee03f63e43425cb6f71f1a1756****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// 依赖资源类型。
	//
	// example:
	//
	// ALIYUN::VPC::NatGateway
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// 依赖资源实例ID。
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
	// k8s对象类型。
	//
	// example:
	//
	// Service
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// k8s对象命名空间。
	//
	// example:
	//
	// kube-system
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// k8s对象名称。
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
	// 删除集群时是否默认删除该资源。
	//
	// - true：默认删除该资源。
	//
	// - fasle：不删除该资源。
	//
	// example:
	//
	// false
	DeleteByDefault *bool `json:"delete_by_default,omitempty" xml:"delete_by_default,omitempty"`
	// `delete_by_default`的默认行为是否可以更改。
	//
	// - true：可以更改。
	//
	// - false：不支持更改。
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
