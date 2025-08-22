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
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description    *string                                                 `json:"description,omitempty" xml:"description,omitempty"`
	ExportToModule *UpdateResourceExportTaskAttributeRequestExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	IncludeRules   []*UpdateResourceExportTaskAttributeRequestIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ramName
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// example:
	//
	// 1.183.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// example:
	//
	// 1.2.6
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// example:
	//
	// Manual
	TriggerStrategy *string                                              `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	Variables       []*UpdateResourceExportTaskAttributeRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type UpdateResourceExportTaskAttributeRequestExportToModule struct {
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
	// example:
	//
	// ZoneId
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
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
	Properties []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	// example:
	//
	// AliCloud::VPC::VPC
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
