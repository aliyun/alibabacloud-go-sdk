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
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// OK
	Description    *string                                        `json:"description,omitempty" xml:"description,omitempty"`
	ExportToModule *CreateResourceExportTaskRequestExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	IncludeRules   []*CreateResourceExportTaskRequestIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// {}
	RamRole                  *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// example:
	//
	// 1.2.6
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// example:
	//
	// Auto
	TriggerStrategy *string                                     `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	Variables       []*CreateResourceExportTaskRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
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
	// example:
	//
	// Registry
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// alibaba/security-group/alicloud
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// example:
	//
	// /
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
	// example:
	//
	// ZoneId
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
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
	Properties []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	// example:
	//
	// AliCloud::VPC::VPC
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
