// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ValidateTemplateResponseBody
	GetDescription() *string
	SetOutputs(v []*ValidateTemplateResponseBodyOutputs) *ValidateTemplateResponseBody
	GetOutputs() []*ValidateTemplateResponseBodyOutputs
	SetParameters(v []map[string]interface{}) *ValidateTemplateResponseBody
	GetParameters() []map[string]interface{}
	SetRequestId(v string) *ValidateTemplateResponseBody
	GetRequestId() *string
	SetResourceTypes(v *ValidateTemplateResponseBodyResourceTypes) *ValidateTemplateResponseBody
	GetResourceTypes() *ValidateTemplateResponseBodyResourceTypes
	SetResources(v []*ValidateTemplateResponseBodyResources) *ValidateTemplateResponseBody
	GetResources() []*ValidateTemplateResponseBodyResources
	SetUpdateInfo(v *ValidateTemplateResponseBodyUpdateInfo) *ValidateTemplateResponseBody
	GetUpdateInfo() *ValidateTemplateResponseBodyUpdateInfo
}

type ValidateTemplateResponseBody struct {
	// The description of the template.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The outputs of the template.
	Outputs []*ValidateTemplateResponseBodyOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// The parameters that are defined in the Parameters section of the template.
	Parameters []map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource types that are used in the template.
	ResourceTypes *ValidateTemplateResponseBodyResourceTypes `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Struct"`
	// The regular resources that are defined in the template.
	//
	// > - For a Resource Orchestration Service (ROS) template, the resource whose definition contains `Count` is not displayed as a list.
	//
	// > -  For a Terraform template, the resource whose definition contains `count` or `for_each` is not displayed as a list.
	Resources []*ValidateTemplateResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The information about the stack update. This parameter cannot be returned if the value of UpdateInfoOptions contains Disabled.
	UpdateInfo *ValidateTemplateResponseBodyUpdateInfo `json:"UpdateInfo,omitempty" xml:"UpdateInfo,omitempty" type:"Struct"`
}

func (s ValidateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ValidateTemplateResponseBody) GetOutputs() []*ValidateTemplateResponseBodyOutputs {
	return s.Outputs
}

func (s *ValidateTemplateResponseBody) GetParameters() []map[string]interface{} {
	return s.Parameters
}

func (s *ValidateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateTemplateResponseBody) GetResourceTypes() *ValidateTemplateResponseBodyResourceTypes {
	return s.ResourceTypes
}

func (s *ValidateTemplateResponseBody) GetResources() []*ValidateTemplateResponseBodyResources {
	return s.Resources
}

func (s *ValidateTemplateResponseBody) GetUpdateInfo() *ValidateTemplateResponseBodyUpdateInfo {
	return s.UpdateInfo
}

func (s *ValidateTemplateResponseBody) SetDescription(v string) *ValidateTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *ValidateTemplateResponseBody) SetOutputs(v []*ValidateTemplateResponseBodyOutputs) *ValidateTemplateResponseBody {
	s.Outputs = v
	return s
}

func (s *ValidateTemplateResponseBody) SetParameters(v []map[string]interface{}) *ValidateTemplateResponseBody {
	s.Parameters = v
	return s
}

func (s *ValidateTemplateResponseBody) SetRequestId(v string) *ValidateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateTemplateResponseBody) SetResourceTypes(v *ValidateTemplateResponseBodyResourceTypes) *ValidateTemplateResponseBody {
	s.ResourceTypes = v
	return s
}

func (s *ValidateTemplateResponseBody) SetResources(v []*ValidateTemplateResponseBodyResources) *ValidateTemplateResponseBody {
	s.Resources = v
	return s
}

func (s *ValidateTemplateResponseBody) SetUpdateInfo(v *ValidateTemplateResponseBodyUpdateInfo) *ValidateTemplateResponseBody {
	s.UpdateInfo = v
	return s
}

func (s *ValidateTemplateResponseBody) Validate() error {
	if s.Outputs != nil {
		for _, item := range s.Outputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceTypes != nil {
		if err := s.ResourceTypes.Validate(); err != nil {
			return err
		}
	}
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UpdateInfo != nil {
		if err := s.UpdateInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ValidateTemplateResponseBodyOutputs struct {
	// The description of the template output.
	//
	// example:
	//
	// The instance ID of my ECS.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The alias of the template output.
	//
	// example:
	//
	// Instance ID
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The name of the template output.
	//
	// example:
	//
	// instance_id
	OutputKey *string `json:"OutputKey,omitempty" xml:"OutputKey,omitempty"`
}

func (s ValidateTemplateResponseBodyOutputs) String() string {
	return dara.Prettify(s)
}

func (s ValidateTemplateResponseBodyOutputs) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponseBodyOutputs) GetDescription() *string {
	return s.Description
}

func (s *ValidateTemplateResponseBodyOutputs) GetLabel() *string {
	return s.Label
}

func (s *ValidateTemplateResponseBodyOutputs) GetOutputKey() *string {
	return s.OutputKey
}

func (s *ValidateTemplateResponseBodyOutputs) SetDescription(v string) *ValidateTemplateResponseBodyOutputs {
	s.Description = &v
	return s
}

func (s *ValidateTemplateResponseBodyOutputs) SetLabel(v string) *ValidateTemplateResponseBodyOutputs {
	s.Label = &v
	return s
}

func (s *ValidateTemplateResponseBodyOutputs) SetOutputKey(v string) *ValidateTemplateResponseBodyOutputs {
	s.OutputKey = &v
	return s
}

func (s *ValidateTemplateResponseBodyOutputs) Validate() error {
	return dara.Validate(s)
}

type ValidateTemplateResponseBodyResourceTypes struct {
	// The DataSource resource types that are used in the template. The value is deduplicated.
	DataSources []*string `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The regular resource types that are used in the template. The value is deduplicated.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ValidateTemplateResponseBodyResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s ValidateTemplateResponseBodyResourceTypes) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponseBodyResourceTypes) GetDataSources() []*string {
	return s.DataSources
}

func (s *ValidateTemplateResponseBodyResourceTypes) GetResources() []*string {
	return s.Resources
}

func (s *ValidateTemplateResponseBodyResourceTypes) SetDataSources(v []*string) *ValidateTemplateResponseBodyResourceTypes {
	s.DataSources = v
	return s
}

func (s *ValidateTemplateResponseBodyResourceTypes) SetResources(v []*string) *ValidateTemplateResponseBodyResourceTypes {
	s.Resources = v
	return s
}

func (s *ValidateTemplateResponseBodyResourceTypes) Validate() error {
	return dara.Validate(s)
}

type ValidateTemplateResponseBodyResources struct {
	// The pattern in which the logical IDs of regular resources are formed.
	//
	// If resources are defined in a ROS template, the following rules apply:
	//
	// 	- Resource whose definition does not contain `Count`: If the resource name defined in the template is `server`, the values of LogicalResourceIdPattern and `ResourcePath` are both `server`.``
	//
	// 	- Resource whose definition contains `Count`: If the resource name defined in the template is `server`, the value of LogicalResourceIdPattern is `server[*]`, and the value of `ResourcePath` is `server`.
	//
	// If resources and [modules](https://www.terraform.io/language/modules) are defined in a Terraform template, the following rules apply:
	//
	// 	- Resource and module whose definitions do not contain [`count`](https://www.terraform.io/language/meta-arguments/count) or [`for_each`](https://www.terraform.io/language/meta-arguments/for_each): If the resource name defined in the template is `server`, the values of LogicalResourceIdPattern and `ResourcePath` are both `server`.``
	//
	// 	- Resource and module whose definitions contain [`count`](https://www.terraform.io/language/meta-arguments/count) or [`for_each`](https://www.terraform.io/language/meta-arguments/for_each): If the resource name defined in the template is `server`, the value of LogicalResourceIdPattern is `server[*]`, and the value of `ResourcePath` is `server`.
	//
	// Examples of LogicalResourceIdPattern for resources in a Terraform template:
	//
	// 	- Valid values of LogicalResourceIdPattern if a resource belongs to the root module:
	//
	//     	- `server`: In this case, `count` and `for_each` are not contained in the resource. The value of `ResourcePath` is `server`.
	//
	//     	- `server[*]`: In this case, `count` or `for_each` is contained in the resource. The value of `ResourcePath` is `server`.
	//
	// 	- Valid values of LogicalResourceIdPattern if a resource belongs to a child module:
	//
	//     	- `app.server`: In this case, `count` and `for_each` are not contained in the `app` module and the `server` resource. The value of `ResourcePath` is `app.server`.````
	//
	//     	- `app.server[*]`: In this case, `count` or `for_each` is contained in the `server` resource, but `count` and `for_each` are not contained in the `app` module. The value of `ResourcePath` is `app.server`.
	//
	//     	- `app[*].server`: In this case, `count` or `for_each` is contained in the `app` module, but `count` and `for_each` are not contained in the `server` resource. The value of `ResourcePath` is `app.server`.
	//
	//     	- `app[*].server[*]`: In this case, `count` or `for_each` is contained in the `app` module and the `server` resource. The value of `ResourcePath` is `app.server`.````
	//
	//     	- `app.app_group[*].server`: In this case, `count` or `for_each` is contained in the `app_group` module, but `count` and `for_each` are not contained in the `app` module and the `server` resource. The value of `ResourcePath` is `app.app_group.server`. The `app_group` module is a child module of the `app` module.````
	//
	// example:
	//
	// server
	LogicalResourceIdPattern *string `json:"LogicalResourceIdPattern,omitempty" xml:"LogicalResourceIdPattern,omitempty"`
	// The path of the regular resource. In most cases, the path of a regular resource is the same as the resource name.
	//
	// example:
	//
	// server
	ResourcePath *string `json:"ResourcePath,omitempty" xml:"ResourcePath,omitempty"`
	// The regular resource type.
	//
	// example:
	//
	// ALIYUN::ECS::InstanceGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ValidateTemplateResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s ValidateTemplateResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponseBodyResources) GetLogicalResourceIdPattern() *string {
	return s.LogicalResourceIdPattern
}

func (s *ValidateTemplateResponseBodyResources) GetResourcePath() *string {
	return s.ResourcePath
}

func (s *ValidateTemplateResponseBodyResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ValidateTemplateResponseBodyResources) SetLogicalResourceIdPattern(v string) *ValidateTemplateResponseBodyResources {
	s.LogicalResourceIdPattern = &v
	return s
}

func (s *ValidateTemplateResponseBodyResources) SetResourcePath(v string) *ValidateTemplateResponseBodyResources {
	s.ResourcePath = &v
	return s
}

func (s *ValidateTemplateResponseBodyResources) SetResourceType(v string) *ValidateTemplateResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *ValidateTemplateResponseBodyResources) Validate() error {
	return dara.Validate(s)
}

type ValidateTemplateResponseBodyUpdateInfo struct {
	// The parameters that can be modified.
	ParametersAllowedToBeModified []*string `json:"ParametersAllowedToBeModified,omitempty" xml:"ParametersAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters whose changes cause service interruptions.
	//
	// > - This parameter is supported only for a small number of resource types.
	//
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersCauseInterruptionIfModified []*string `json:"ParametersCauseInterruptionIfModified,omitempty" xml:"ParametersCauseInterruptionIfModified,omitempty" type:"Repeated"`
	// The parameters whose changes trigger replacement updates for resources.
	//
	// > -  This parameter can be returned only if the value of UpdateInfoOptions contains EnableReplacement.
	//
	// > -  This parameter is valid only for updates on ROS stacks.
	ParametersCauseReplacementIfModified []*string `json:"ParametersCauseReplacementIfModified,omitempty" xml:"ParametersCauseReplacementIfModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under specific conditions.
	ParametersConditionallyAllowedToBeModified []*string `json:"ParametersConditionallyAllowedToBeModified,omitempty" xml:"ParametersConditionallyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters whose changes cause service interruptions under specific conditions.
	//
	// > - This parameter is supported only for a small number of resource types.
	//
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersConditionallyCauseInterruptionIfModified []*string `json:"ParametersConditionallyCauseInterruptionIfModified,omitempty" xml:"ParametersConditionallyCauseInterruptionIfModified,omitempty" type:"Repeated"`
	// The parameters whose changes trigger replacement updates for resources under specific conditions.
	//
	// > - This parameter can be returned only if the value of UpdateInfoOptions contains EnableReplacement.
	//
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersConditionallyCauseReplacementIfModified []*string `json:"ParametersConditionallyCauseReplacementIfModified,omitempty" xml:"ParametersConditionallyCauseReplacementIfModified,omitempty" type:"Repeated"`
	// The parameters that cannot be modified.
	ParametersNotAllowedToBeModified []*string `json:"ParametersNotAllowedToBeModified,omitempty" xml:"ParametersNotAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under uncertain conditions.
	ParametersUncertainlyAllowedToBeModified []*string `json:"ParametersUncertainlyAllowedToBeModified,omitempty" xml:"ParametersUncertainlyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters whose changes cause service interruptions under uncertain conditions.
	//
	// > - This parameter is supported only for a small number of resource types.
	//
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersUncertainlyCauseInterruptionIfModified []*string `json:"ParametersUncertainlyCauseInterruptionIfModified,omitempty" xml:"ParametersUncertainlyCauseInterruptionIfModified,omitempty" type:"Repeated"`
	// The parameters whose changes trigger replacement updates for resources under uncertain conditions.
	//
	// > -  This parameter can be returned only if the value of UpdateInfoOptions contains EnableReplacement.
	//
	// > -  This parameter is valid only for updates on ROS stacks.
	ParametersUncertainlyCauseReplacementIfModified []*string `json:"ParametersUncertainlyCauseReplacementIfModified,omitempty" xml:"ParametersUncertainlyCauseReplacementIfModified,omitempty" type:"Repeated"`
}

func (s ValidateTemplateResponseBodyUpdateInfo) String() string {
	return dara.Prettify(s)
}

func (s ValidateTemplateResponseBodyUpdateInfo) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponseBodyUpdateInfo) GetParametersAllowedToBeModified() []*string {
	return s.ParametersAllowedToBeModified
}

func (s *ValidateTemplateResponseBodyUpdateInfo) GetParametersCauseInterruptionIfModified() []*string {
	return s.ParametersCauseInterruptionIfModified
}

func (s *ValidateTemplateResponseBodyUpdateInfo) GetParametersCauseReplacementIfModified() []*string {
	return s.ParametersCauseReplacementIfModified
}

func (s *ValidateTemplateResponseBodyUpdateInfo) GetParametersConditionallyAllowedToBeModified() []*string {
	return s.ParametersConditionallyAllowedToBeModified
}

func (s *ValidateTemplateResponseBodyUpdateInfo) GetParametersConditionallyCauseInterruptionIfModified() []*string {
	return s.ParametersConditionallyCauseInterruptionIfModified
}

func (s *ValidateTemplateResponseBodyUpdateInfo) GetParametersConditionallyCauseReplacementIfModified() []*string {
	return s.ParametersConditionallyCauseReplacementIfModified
}

func (s *ValidateTemplateResponseBodyUpdateInfo) GetParametersNotAllowedToBeModified() []*string {
	return s.ParametersNotAllowedToBeModified
}

func (s *ValidateTemplateResponseBodyUpdateInfo) GetParametersUncertainlyAllowedToBeModified() []*string {
	return s.ParametersUncertainlyAllowedToBeModified
}

func (s *ValidateTemplateResponseBodyUpdateInfo) GetParametersUncertainlyCauseInterruptionIfModified() []*string {
	return s.ParametersUncertainlyCauseInterruptionIfModified
}

func (s *ValidateTemplateResponseBodyUpdateInfo) GetParametersUncertainlyCauseReplacementIfModified() []*string {
	return s.ParametersUncertainlyCauseReplacementIfModified
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersAllowedToBeModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersAllowedToBeModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersCauseInterruptionIfModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersCauseInterruptionIfModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersCauseReplacementIfModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersCauseReplacementIfModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersConditionallyAllowedToBeModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersConditionallyAllowedToBeModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersConditionallyCauseInterruptionIfModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersConditionallyCauseInterruptionIfModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersConditionallyCauseReplacementIfModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersConditionallyCauseReplacementIfModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersNotAllowedToBeModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersNotAllowedToBeModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersUncertainlyAllowedToBeModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersUncertainlyAllowedToBeModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersUncertainlyCauseInterruptionIfModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersUncertainlyCauseInterruptionIfModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersUncertainlyCauseReplacementIfModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersUncertainlyCauseReplacementIfModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) Validate() error {
	return dara.Validate(s)
}
