// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *UpdateWorkspaceResourceRequest
	GetGroupName() *string
	SetIsDefault(v bool) *UpdateWorkspaceResourceRequest
	GetIsDefault() *bool
	SetLabels(v []*UpdateWorkspaceResourceRequestLabels) *UpdateWorkspaceResourceRequest
	GetLabels() []*UpdateWorkspaceResourceRequestLabels
	SetProductType(v string) *UpdateWorkspaceResourceRequest
	GetProductType() *string
	SetResourceIds(v []*string) *UpdateWorkspaceResourceRequest
	GetResourceIds() []*string
	SetResourceType(v string) *UpdateWorkspaceResourceRequest
	GetResourceType() *string
	SetSpec(v map[string]interface{}) *UpdateWorkspaceResourceRequest
	GetSpec() map[string]interface{}
}

type UpdateWorkspaceResourceRequest struct {
	// The group name.
	//
	// example:
	//
	// group-kjds******sd
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Specifies whether the resource is the default resource. This parameter can only be set to true and cannot be set to false.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The resource tags. If you specify multiple tags, only resources that meet all the specified tag-based filter conditions are returned.
	Labels []*UpdateWorkspaceResourceRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// **This field is no longer used and will be removed. Use the ResourceType field.
	//
	// example:
	//
	// MaxCompute
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The resource IDs.
	//
	// You cannot leave both GroupName and ResourceIds empty. If you specify both the parameters, the value of GroupName of each resource ID in the dataset must be the same.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The resource type. Valid values:
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
	// The specification of the resource.
	//
	// example:
	//
	// {
	//
	//       "clusterType": "share"
	//
	// }
	Spec map[string]interface{} `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateWorkspaceResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResourceRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateWorkspaceResourceRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *UpdateWorkspaceResourceRequest) GetLabels() []*UpdateWorkspaceResourceRequestLabels {
	return s.Labels
}

func (s *UpdateWorkspaceResourceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *UpdateWorkspaceResourceRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *UpdateWorkspaceResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateWorkspaceResourceRequest) GetSpec() map[string]interface{} {
	return s.Spec
}

func (s *UpdateWorkspaceResourceRequest) SetGroupName(v string) *UpdateWorkspaceResourceRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateWorkspaceResourceRequest) SetIsDefault(v bool) *UpdateWorkspaceResourceRequest {
	s.IsDefault = &v
	return s
}

func (s *UpdateWorkspaceResourceRequest) SetLabels(v []*UpdateWorkspaceResourceRequestLabels) *UpdateWorkspaceResourceRequest {
	s.Labels = v
	return s
}

func (s *UpdateWorkspaceResourceRequest) SetProductType(v string) *UpdateWorkspaceResourceRequest {
	s.ProductType = &v
	return s
}

func (s *UpdateWorkspaceResourceRequest) SetResourceIds(v []*string) *UpdateWorkspaceResourceRequest {
	s.ResourceIds = v
	return s
}

func (s *UpdateWorkspaceResourceRequest) SetResourceType(v string) *UpdateWorkspaceResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateWorkspaceResourceRequest) SetSpec(v map[string]interface{}) *UpdateWorkspaceResourceRequest {
	s.Spec = v
	return s
}

func (s *UpdateWorkspaceResourceRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkspaceResourceRequestLabels struct {
	// The tag key.
	//
	// example:
	//
	// system.******
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// True
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateWorkspaceResourceRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResourceRequestLabels) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResourceRequestLabels) GetKey() *string {
	return s.Key
}

func (s *UpdateWorkspaceResourceRequestLabels) GetValue() *string {
	return s.Value
}

func (s *UpdateWorkspaceResourceRequestLabels) SetKey(v string) *UpdateWorkspaceResourceRequestLabels {
	s.Key = &v
	return s
}

func (s *UpdateWorkspaceResourceRequestLabels) SetValue(v string) *UpdateWorkspaceResourceRequestLabels {
	s.Value = &v
	return s
}

func (s *UpdateWorkspaceResourceRequestLabels) Validate() error {
	return dara.Validate(s)
}
