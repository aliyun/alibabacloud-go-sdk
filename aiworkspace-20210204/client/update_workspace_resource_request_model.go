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
	// Specifies whether the resource is the default resource. Currently, you can set this parameter only to true. The value false is not supported.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The collection of resource tags. If multiple tags are specified, only resources that have all the specified tags are returned.
	Labels []*UpdateWorkspaceResourceRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// **[Deprecated]*	- This field is deprecated and will be taken offline. Use the ResourceType field instead.
	//
	// example:
	//
	// MaxCompute
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The collection of resource IDs. GroupName and ResourceIds cannot be both empty, but they can be both specified. If both are specified, the GroupName of each resource ID in the collection must be the same.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The list of resource types. Valid values:
	//
	// - MaxCompute: MaxCompute big data compute EPS resources
	//
	// - ECS: general-purpose compute resources
	//
	// - Lingjun: Lingjun AI computing resources
	//
	// - ACS: ACS compute resources
	//
	// - FLINK: Flink resources
	//
	// - SelfManagedAckPro: ACK Pro unified management cluster resources
	//
	// - SelfManagedAckLingjun: ACK Lingjun unified management cluster resources
	//
	// - SelfManagedASI: ASI unified management cluster resources (third-party cloud)
	//
	// example:
	//
	// MaxCompute
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The specifications of the resource.
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
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
