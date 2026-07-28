// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceExportTaskAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateResourceExportTaskAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateResourceExportTaskAttributeRequest
	GetDescription() *string
	SetExportToModule(v *UpdateResourceExportTaskAttributeRequestExportToModule) *UpdateResourceExportTaskAttributeRequest
	GetExportToModule() *UpdateResourceExportTaskAttributeRequestExportToModule
	SetIncludeRules(v []*UpdateResourceExportTaskAttributeRequestIncludeRules) *UpdateResourceExportTaskAttributeRequest
	GetIncludeRules() []*UpdateResourceExportTaskAttributeRequestIncludeRules
	SetName(v string) *UpdateResourceExportTaskAttributeRequest
	GetName() *string
	SetRamRole(v string) *UpdateResourceExportTaskAttributeRequest
	GetRamRole() *string
	SetTerraformProviderVersion(v string) *UpdateResourceExportTaskAttributeRequest
	GetTerraformProviderVersion() *string
	SetTerraformVersion(v string) *UpdateResourceExportTaskAttributeRequest
	GetTerraformVersion() *string
	SetTriggerStrategy(v string) *UpdateResourceExportTaskAttributeRequest
	GetTriggerStrategy() *string
	SetVariables(v []*UpdateResourceExportTaskAttributeRequestVariables) *UpdateResourceExportTaskAttributeRequest
	GetVariables() []*UpdateResourceExportTaskAttributeRequestVariables
}

type UpdateResourceExportTaskAttributeRequest struct {
	// The idempotency token. Format: [0-9a-zA-Z-]{1,64}. We recommend that you use a UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The description of the task.
	//
	// example:
	//
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Saves the exported template as a module. If this parameter is not set, the template is automatically saved in the registry.
	ExportToModule *UpdateResourceExportTaskAttributeRequestExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	// The list of include rules used when exporting resources.
	IncludeRules []*UpdateResourceExportTaskAttributeRequestIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	// The name of the resource export task. The name must meet the following requirements:
	//
	// - The name must be 2 to 128 characters in length.
	//
	// - The name can contain letters, digits, Chinese characters, hyphens (-), underscores (_), and periods (.). The name cannot start or end with a hyphen, underscore, or period.
	//
	// - The name must be unique among resource export tasks within the current account.
	//
	// example:
	//
	// TaskName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The RAM role.
	//
	// example:
	//
	// ramName
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// The Terraform provider version. Call **ListTerraformProviderVersions*	- to view the supported versions. Default value: the latest version.
	//
	// example:
	//
	// 1.183.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// The Terraform version. Call **ListAvailableTerraformVersions*	- to view the supported versions. Default value: 1.5.7.
	//
	// example:
	//
	// 1.5.7
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// The trigger strategy. Valid values:
	//
	// - Auto: triggered automatically when rules are modified or the trigger strategy is changed to Auto.
	//
	// - Manual: triggered manually.
	//
	// Default value: Manual.
	//
	// example:
	//
	// Manual
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// The list of variables. Sets exported resource parameters as variables.
	Variables []*UpdateResourceExportTaskAttributeRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s UpdateResourceExportTaskAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateResourceExportTaskAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateResourceExportTaskAttributeRequest) GetExportToModule() *UpdateResourceExportTaskAttributeRequestExportToModule {
	return s.ExportToModule
}

func (s *UpdateResourceExportTaskAttributeRequest) GetIncludeRules() []*UpdateResourceExportTaskAttributeRequestIncludeRules {
	return s.IncludeRules
}

func (s *UpdateResourceExportTaskAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateResourceExportTaskAttributeRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *UpdateResourceExportTaskAttributeRequest) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *UpdateResourceExportTaskAttributeRequest) GetTerraformVersion() *string {
	return s.TerraformVersion
}

func (s *UpdateResourceExportTaskAttributeRequest) GetTriggerStrategy() *string {
	return s.TriggerStrategy
}

func (s *UpdateResourceExportTaskAttributeRequest) GetVariables() []*UpdateResourceExportTaskAttributeRequestVariables {
	return s.Variables
}

func (s *UpdateResourceExportTaskAttributeRequest) SetClientToken(v string) *UpdateResourceExportTaskAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetDescription(v string) *UpdateResourceExportTaskAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetExportToModule(v *UpdateResourceExportTaskAttributeRequestExportToModule) *UpdateResourceExportTaskAttributeRequest {
	s.ExportToModule = v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetIncludeRules(v []*UpdateResourceExportTaskAttributeRequestIncludeRules) *UpdateResourceExportTaskAttributeRequest {
	s.IncludeRules = v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetName(v string) *UpdateResourceExportTaskAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetRamRole(v string) *UpdateResourceExportTaskAttributeRequest {
	s.RamRole = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetTerraformProviderVersion(v string) *UpdateResourceExportTaskAttributeRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetTerraformVersion(v string) *UpdateResourceExportTaskAttributeRequest {
	s.TerraformVersion = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetTriggerStrategy(v string) *UpdateResourceExportTaskAttributeRequest {
	s.TriggerStrategy = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetVariables(v []*UpdateResourceExportTaskAttributeRequestVariables) *UpdateResourceExportTaskAttributeRequest {
	s.Variables = v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) Validate() error {
	if s.ExportToModule != nil {
		if err := s.ExportToModule.Validate(); err != nil {
			return err
		}
	}
	if s.IncludeRules != nil {
		for _, item := range s.IncludeRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateResourceExportTaskAttributeRequestExportToModule struct {
	// The module type in which the exported template is saved. Valid values:
	//
	// - OSS: OSS
	//
	// - Registry: Terraform Registry.
	//
	// example:
	//
	// Registry
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The path for saving the template content. Set this parameter when source is set to OSS. Format: oss::https://<bucket>.oss-<region>.aliyuncs.com/<path>.zip.
	//
	// example:
	//
	// oss::https://iac-daily.oss-ap-southeast-1.aliyuncs.com/iacservice/vpc.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The path for saving the template state file. Set this parameter when source is set to OSS. Format: oss::https://<bucket>.oss-<region>.aliyuncs.com/<path>/terraform.tfstate.
	//
	// example:
	//
	// oss::https://iac-daily.oss-ap-southeast-1.aliyuncs.com/default/terraform.tfstate
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
}

func (s UpdateResourceExportTaskAttributeRequestExportToModule) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeRequestExportToModule) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeRequestExportToModule) GetSource() *string {
	return s.Source
}

func (s *UpdateResourceExportTaskAttributeRequestExportToModule) GetSourcePath() *string {
	return s.SourcePath
}

func (s *UpdateResourceExportTaskAttributeRequestExportToModule) GetStatePath() *string {
	return s.StatePath
}

func (s *UpdateResourceExportTaskAttributeRequestExportToModule) SetSource(v string) *UpdateResourceExportTaskAttributeRequestExportToModule {
	s.Source = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequestExportToModule) SetSourcePath(v string) *UpdateResourceExportTaskAttributeRequestExportToModule {
	s.SourcePath = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequestExportToModule) SetStatePath(v string) *UpdateResourceExportTaskAttributeRequestExportToModule {
	s.StatePath = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequestExportToModule) Validate() error {
	return dara.Validate(s)
}

type UpdateResourceExportTaskAttributeRequestIncludeRules struct {
	// The name of the include rule for resource export. Valid values:
	//
	// - ResourceType: required. The resource type. Call **ListResourceTypes*	- to view the supported resources. Example: ALIYUN::VPC::VPC.
	//
	// - RegionId: required. The region to which the resource belongs. Only one region is supported. Example: cn-chengdu.
	//
	// - <ResourceType>:Id: the resource ID. Example: ALIYUN::VPC::VPC:Id.
	//
	// - ResourceGroupId: the resource group ID. Example: rg-1234.
	//
	// - ZoneId: the zone to which the resource belongs. Only one zone is supported. Example: ap-southeast-1-h.
	//
	// Multiple filter conditions have an AND relationship by default. A resource must meet all filter conditions to be considered a match.
	//
	// example:
	//
	// RegionId
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The values of an include rule for resource export.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s UpdateResourceExportTaskAttributeRequestIncludeRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeRequestIncludeRules) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeRequestIncludeRules) GetKey() *string {
	return s.Key
}

func (s *UpdateResourceExportTaskAttributeRequestIncludeRules) GetValues() []*string {
	return s.Values
}

func (s *UpdateResourceExportTaskAttributeRequestIncludeRules) SetKey(v string) *UpdateResourceExportTaskAttributeRequestIncludeRules {
	s.Key = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequestIncludeRules) SetValues(v []*string) *UpdateResourceExportTaskAttributeRequestIncludeRules {
	s.Values = v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequestIncludeRules) Validate() error {
	return dara.Validate(s)
}

type UpdateResourceExportTaskAttributeRequestVariables struct {
	// The list of Terraform resource properties corresponding to the resource type.
	Properties []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	// The resource type. Call **ListResourceTypes*	- to view the supported resources.
	//
	// example:
	//
	// ALIYUN::VPC::VSwitch
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s UpdateResourceExportTaskAttributeRequestVariables) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeRequestVariables) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeRequestVariables) GetProperties() []*string {
	return s.Properties
}

func (s *UpdateResourceExportTaskAttributeRequestVariables) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateResourceExportTaskAttributeRequestVariables) SetProperties(v []*string) *UpdateResourceExportTaskAttributeRequestVariables {
	s.Properties = v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequestVariables) SetResourceType(v string) *UpdateResourceExportTaskAttributeRequestVariables {
	s.ResourceType = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequestVariables) Validate() error {
	return dara.Validate(s)
}
