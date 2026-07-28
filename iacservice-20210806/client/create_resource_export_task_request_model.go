// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateResourceExportTaskRequest
	GetClientToken() *string
	SetDescription(v string) *CreateResourceExportTaskRequest
	GetDescription() *string
	SetExportToModule(v *CreateResourceExportTaskRequestExportToModule) *CreateResourceExportTaskRequest
	GetExportToModule() *CreateResourceExportTaskRequestExportToModule
	SetIncludeRules(v []*CreateResourceExportTaskRequestIncludeRules) *CreateResourceExportTaskRequest
	GetIncludeRules() []*CreateResourceExportTaskRequestIncludeRules
	SetName(v string) *CreateResourceExportTaskRequest
	GetName() *string
	SetRamRole(v string) *CreateResourceExportTaskRequest
	GetRamRole() *string
	SetTerraformProviderVersion(v string) *CreateResourceExportTaskRequest
	GetTerraformProviderVersion() *string
	SetTerraformVersion(v string) *CreateResourceExportTaskRequest
	GetTerraformVersion() *string
	SetTriggerStrategy(v string) *CreateResourceExportTaskRequest
	GetTriggerStrategy() *string
	SetVariables(v []*CreateResourceExportTaskRequestVariables) *CreateResourceExportTaskRequest
	GetVariables() []*CreateResourceExportTaskRequestVariables
}

type CreateResourceExportTaskRequest struct {
	// The idempotency token. Format: [0-9a-zA-Z-]{1,64}. We recommend that you use a UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The description of the resource export task.
	//
	// example:
	//
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Saves the exported template as a module. If this parameter is not specified, the template is automatically saved in the Registry.
	ExportToModule *CreateResourceExportTaskRequestExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	// The list of inclusion rules used when exporting resources.
	IncludeRules []*CreateResourceExportTaskRequestIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	// The name of the resource export task. The name must meet the following requirements:
	//
	// - The name must be 3 to 63 characters in length.
	//
	// - The name can contain letters, digits, Chinese characters, hyphens (-), underscores (_), and periods (.). The name cannot start or end with a hyphen, underscore, or period.
	//
	// - The name must be unique among resource export tasks within the current account.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The RAM role (1 to 128 characters). The system assumes this role to execute the template when a new job is triggered. This parameter is required when the job trigger method is not manual.
	//
	// example:
	//
	// role-name
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// The Terraform provider version. Call **ListTerraformProviderVersions*	- to view the list of supported versions. Default value: the latest version.
	//
	// example:
	//
	// 1.247.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// The Terraform version. Call **ListAvailableTerraformVersions*	- to view the list of supported versions. Default value: 1.5.7.
	//
	// example:
	//
	// 1.5.7
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// The trigger strategy. Valid values:
	//
	// - Auto: triggered when rules are modified or the trigger strategy is changed to Auto.
	//
	// - Manual: manually triggered.
	//
	// Default value: Manual.
	//
	// example:
	//
	// Auto
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// The list of variables. Exported resource parameters are set as variables.
	Variables []*CreateResourceExportTaskRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s CreateResourceExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateResourceExportTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateResourceExportTaskRequest) GetExportToModule() *CreateResourceExportTaskRequestExportToModule {
	return s.ExportToModule
}

func (s *CreateResourceExportTaskRequest) GetIncludeRules() []*CreateResourceExportTaskRequestIncludeRules {
	return s.IncludeRules
}

func (s *CreateResourceExportTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateResourceExportTaskRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *CreateResourceExportTaskRequest) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *CreateResourceExportTaskRequest) GetTerraformVersion() *string {
	return s.TerraformVersion
}

func (s *CreateResourceExportTaskRequest) GetTriggerStrategy() *string {
	return s.TriggerStrategy
}

func (s *CreateResourceExportTaskRequest) GetVariables() []*CreateResourceExportTaskRequestVariables {
	return s.Variables
}

func (s *CreateResourceExportTaskRequest) SetClientToken(v string) *CreateResourceExportTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetDescription(v string) *CreateResourceExportTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetExportToModule(v *CreateResourceExportTaskRequestExportToModule) *CreateResourceExportTaskRequest {
	s.ExportToModule = v
	return s
}

func (s *CreateResourceExportTaskRequest) SetIncludeRules(v []*CreateResourceExportTaskRequestIncludeRules) *CreateResourceExportTaskRequest {
	s.IncludeRules = v
	return s
}

func (s *CreateResourceExportTaskRequest) SetName(v string) *CreateResourceExportTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetRamRole(v string) *CreateResourceExportTaskRequest {
	s.RamRole = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetTerraformProviderVersion(v string) *CreateResourceExportTaskRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetTerraformVersion(v string) *CreateResourceExportTaskRequest {
	s.TerraformVersion = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetTriggerStrategy(v string) *CreateResourceExportTaskRequest {
	s.TriggerStrategy = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetVariables(v []*CreateResourceExportTaskRequestVariables) *CreateResourceExportTaskRequest {
	s.Variables = v
	return s
}

func (s *CreateResourceExportTaskRequest) Validate() error {
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

type CreateResourceExportTaskRequestExportToModule struct {
	// The module type in which the exported template is saved. Valid values:
	//
	// - OSS: OSS.
	//
	// - Registry: Terraform Registry.
	//
	// example:
	//
	// Registry
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The path for saving the template content. Set this parameter when source is set to OSS.
	//
	// example:
	//
	// oss::https://iac-daily.oss-ap-southeast-1.aliyuncs.com/iacservice/vpc.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The path for saving the template state file. Set this parameter when source is set to OSS.
	//
	// example:
	//
	// oss::https://iac-daily.oss-ap-southeast-1.aliyuncs.com/default/terraform.tfstate
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
}

func (s CreateResourceExportTaskRequestExportToModule) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceExportTaskRequestExportToModule) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskRequestExportToModule) GetSource() *string {
	return s.Source
}

func (s *CreateResourceExportTaskRequestExportToModule) GetSourcePath() *string {
	return s.SourcePath
}

func (s *CreateResourceExportTaskRequestExportToModule) GetStatePath() *string {
	return s.StatePath
}

func (s *CreateResourceExportTaskRequestExportToModule) SetSource(v string) *CreateResourceExportTaskRequestExportToModule {
	s.Source = &v
	return s
}

func (s *CreateResourceExportTaskRequestExportToModule) SetSourcePath(v string) *CreateResourceExportTaskRequestExportToModule {
	s.SourcePath = &v
	return s
}

func (s *CreateResourceExportTaskRequestExportToModule) SetStatePath(v string) *CreateResourceExportTaskRequestExportToModule {
	s.StatePath = &v
	return s
}

func (s *CreateResourceExportTaskRequestExportToModule) Validate() error {
	return dara.Validate(s)
}

type CreateResourceExportTaskRequestIncludeRules struct {
	// The name of the inclusion rule used when exporting resources. Valid values:
	//
	// - ResourceType: required. The resource type. Call **ListResourceTypes*	- to view the list of supported resources. Example: ALIYUN::VPC::VPC.
	//
	// - RegionId: required. The region to which the resource belongs. Only one region is supported. Example: cn-chengdu.
	//
	// - \\<ResourceType>:Id: the resource ID. Example: ALIYUN::VPC::VPC:Id.
	//
	// - ResourceGroupId: the resource group ID. Example: rg-1234.
	//
	// - ZoneId: the zone to which the resource belongs. Only one zone is supported. Example: cn-hangzhou-h.
	//
	// By default, multiple filter conditions are evaluated using the AND operator. A resource is considered a match only when all filter conditions are met.
	//
	// example:
	//
	// RegionId
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The list of values for the inclusion rule used when exporting resources.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateResourceExportTaskRequestIncludeRules) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceExportTaskRequestIncludeRules) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskRequestIncludeRules) GetKey() *string {
	return s.Key
}

func (s *CreateResourceExportTaskRequestIncludeRules) GetValues() []*string {
	return s.Values
}

func (s *CreateResourceExportTaskRequestIncludeRules) SetKey(v string) *CreateResourceExportTaskRequestIncludeRules {
	s.Key = &v
	return s
}

func (s *CreateResourceExportTaskRequestIncludeRules) SetValues(v []*string) *CreateResourceExportTaskRequestIncludeRules {
	s.Values = v
	return s
}

func (s *CreateResourceExportTaskRequestIncludeRules) Validate() error {
	return dara.Validate(s)
}

type CreateResourceExportTaskRequestVariables struct {
	// The list of properties of the Terraform resource that corresponds to the resource type.
	Properties []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	// The resource type. Call **ListResourceTypes*	- to view the list of supported resources.
	//
	// example:
	//
	// ALIYUN::VPC::VSwitch
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s CreateResourceExportTaskRequestVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceExportTaskRequestVariables) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskRequestVariables) GetProperties() []*string {
	return s.Properties
}

func (s *CreateResourceExportTaskRequestVariables) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateResourceExportTaskRequestVariables) SetProperties(v []*string) *CreateResourceExportTaskRequestVariables {
	s.Properties = v
	return s
}

func (s *CreateResourceExportTaskRequestVariables) SetResourceType(v string) *CreateResourceExportTaskRequestVariables {
	s.ResourceType = &v
	return s
}

func (s *CreateResourceExportTaskRequestVariables) Validate() error {
	return dara.Validate(s)
}
