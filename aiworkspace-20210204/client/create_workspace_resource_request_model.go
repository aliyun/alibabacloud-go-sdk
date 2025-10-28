// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOption(v string) *CreateWorkspaceResourceRequest
	GetOption() *string
	SetResources(v []*CreateWorkspaceResourceRequestResources) *CreateWorkspaceResourceRequest
	GetResources() []*CreateWorkspaceResourceRequestResources
}

type CreateWorkspaceResourceRequest struct {
	// The operation to perform. Valid values:
	//
	// 	- CreateAndAttach: creates resources and associates the resources with a workspace.
	//
	// 	- Attach: associates resources with a workspace.
	//
	// >  MaxCompute supports only the Attach operation.
	//
	// example:
	//
	// CreateAndAttach
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// The resources.
	//
	// This parameter is required.
	Resources []*CreateWorkspaceResourceRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s CreateWorkspaceResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceRequest) GetOption() *string {
	return s.Option
}

func (s *CreateWorkspaceResourceRequest) GetResources() []*CreateWorkspaceResourceRequestResources {
	return s.Resources
}

func (s *CreateWorkspaceResourceRequest) SetOption(v string) *CreateWorkspaceResourceRequest {
	s.Option = &v
	return s
}

func (s *CreateWorkspaceResourceRequest) SetResources(v []*CreateWorkspaceResourceRequestResources) *CreateWorkspaceResourceRequest {
	s.Resources = v
	return s
}

func (s *CreateWorkspaceResourceRequest) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateWorkspaceResourceRequestResources struct {
	// The environment type. Valid values:
	//
	// 	- dev: development environment
	//
	// 	- prod: production environment
	//
	// This parameter is required.
	//
	// example:
	//
	// prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name of the resource group, which is unique within your Alibaba Cloud account. This parameter is required for MaxCompute, Elastic Compute Service (ECS), Lingjun, Alibaba Cloud Container Compute Service (ACS), and Realtime Compute for Apache Flink resources.
	//
	// example:
	//
	// groupName
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Specifies whether the resource is the default resource. Each type of resources has a default resource. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The labels added to the resource.
	Labels []*CreateWorkspaceResourceRequestResourcesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The resource name. The name must meet the following requirements:
	//
	// 	- The name must be 3 to 28 characters in length, and can contain only letters, digits, and underscores (_). The name must start with a letter.
	//
	// 	- The name must be unique in the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// ResourceName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// **This parameter is no longer used and will be removed. Use the ResourceType parameter instead.
	//
	// example:
	//
	// MaxCompute
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The quotas. Only MaxCompute quotas are available.
	Quotas []*CreateWorkspaceResourceRequestResourcesQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The resource types. Valid values:
	//
	// 	- MaxCompute
	//
	// 	- ECS
	//
	// 	- Lingjun
	//
	// 	- ACS
	//
	// 	- FLINK
	//
	// example:
	//
	// MaxCompute
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource specifications in the JSON format.
	Spec map[string]interface{} `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateWorkspaceResourceRequestResources) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResourceRequestResources) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceRequestResources) GetEnvType() *string {
	return s.EnvType
}

func (s *CreateWorkspaceResourceRequestResources) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateWorkspaceResourceRequestResources) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *CreateWorkspaceResourceRequestResources) GetLabels() []*CreateWorkspaceResourceRequestResourcesLabels {
	return s.Labels
}

func (s *CreateWorkspaceResourceRequestResources) GetName() *string {
	return s.Name
}

func (s *CreateWorkspaceResourceRequestResources) GetProductType() *string {
	return s.ProductType
}

func (s *CreateWorkspaceResourceRequestResources) GetQuotas() []*CreateWorkspaceResourceRequestResourcesQuotas {
	return s.Quotas
}

func (s *CreateWorkspaceResourceRequestResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateWorkspaceResourceRequestResources) GetSpec() map[string]interface{} {
	return s.Spec
}

func (s *CreateWorkspaceResourceRequestResources) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateWorkspaceResourceRequestResources) SetEnvType(v string) *CreateWorkspaceResourceRequestResources {
	s.EnvType = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetGroupName(v string) *CreateWorkspaceResourceRequestResources {
	s.GroupName = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetIsDefault(v bool) *CreateWorkspaceResourceRequestResources {
	s.IsDefault = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetLabels(v []*CreateWorkspaceResourceRequestResourcesLabels) *CreateWorkspaceResourceRequestResources {
	s.Labels = v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetName(v string) *CreateWorkspaceResourceRequestResources {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetProductType(v string) *CreateWorkspaceResourceRequestResources {
	s.ProductType = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetQuotas(v []*CreateWorkspaceResourceRequestResourcesQuotas) *CreateWorkspaceResourceRequestResources {
	s.Quotas = v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetResourceType(v string) *CreateWorkspaceResourceRequestResources {
	s.ResourceType = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetSpec(v map[string]interface{}) *CreateWorkspaceResourceRequestResources {
	s.Spec = v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetWorkspaceId(v string) *CreateWorkspaceResourceRequestResources {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Quotas != nil {
		for _, item := range s.Quotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateWorkspaceResourceRequestResourcesLabels struct {
	// The label key.
	//
	// example:
	//
	// system.support.eas
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The label value.
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateWorkspaceResourceRequestResourcesLabels) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResourceRequestResourcesLabels) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceRequestResourcesLabels) GetKey() *string {
	return s.Key
}

func (s *CreateWorkspaceResourceRequestResourcesLabels) GetValue() *string {
	return s.Value
}

func (s *CreateWorkspaceResourceRequestResourcesLabels) SetKey(v string) *CreateWorkspaceResourceRequestResourcesLabels {
	s.Key = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResourcesLabels) SetValue(v string) *CreateWorkspaceResourceRequestResourcesLabels {
	s.Value = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResourcesLabels) Validate() error {
	return dara.Validate(s)
}

type CreateWorkspaceResourceRequestResourcesQuotas struct {
	// The quota ID. You can call [ListQuotas](https://help.aliyun.com/document_detail/449144.html) to obtain the quota ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 232892******92912
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateWorkspaceResourceRequestResourcesQuotas) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResourceRequestResourcesQuotas) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceRequestResourcesQuotas) GetId() *string {
	return s.Id
}

func (s *CreateWorkspaceResourceRequestResourcesQuotas) SetId(v string) *CreateWorkspaceResourceRequestResourcesQuotas {
	s.Id = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResourcesQuotas) Validate() error {
	return dara.Validate(s)
}
