// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateModuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGenerateSource(v string) *GenerateModuleRequest
	GetGenerateSource() *string
	SetParameters(v map[string]interface{}) *GenerateModuleRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *GenerateModuleRequest
	GetRegionId() *string
	SetSyntax(v string) *GenerateModuleRequest
	GetSyntax() *string
	SetTemplate(v string) *GenerateModuleRequest
	GetTemplate() *string
	SetTerraformProviderVersion(v string) *GenerateModuleRequest
	GetTerraformProviderVersion() *string
	SetTerraformResourceType(v string) *GenerateModuleRequest
	GetTerraformResourceType() *string
}

type GenerateModuleRequest struct {
	// The generation source. Valid values:
	//
	// - Resource: Generates a Terraform HCL template based on resource properties.
	//
	// - VariableToCode: Generates a final Terraform HCL template by combining variables with an existing Terraform HCL template.
	//
	// - CodeToVariable: Extracts variable information from a Terraform HCL template.
	//
	// - Module: Generates Terraform Module code based on variables.
	//
	// example:
	//
	// Resource
	GenerateSource *string `json:"generateSource,omitempty" xml:"generateSource,omitempty"`
	// The collection of parameters, passed in key:value format, such as {"vpc_name":"vpc-test"}.
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The syntax. Valid values:
	//
	// - hcl (default).
	//
	// example:
	//
	// hcl
	Syntax *string `json:"syntax,omitempty" xml:"syntax,omitempty"`
	// The existing Terraform HCL template content.
	//
	// example:
	//
	// terraform {
	//
	// }
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// The Terraform provider version.
	//
	// example:
	//
	// 1.260.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// The Terraform resource type.
	//
	// example:
	//
	// alicloud_vpc
	TerraformResourceType *string `json:"terraformResourceType,omitempty" xml:"terraformResourceType,omitempty"`
}

func (s GenerateModuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateModuleRequest) GoString() string {
	return s.String()
}

func (s *GenerateModuleRequest) GetGenerateSource() *string {
	return s.GenerateSource
}

func (s *GenerateModuleRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GenerateModuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateModuleRequest) GetSyntax() *string {
	return s.Syntax
}

func (s *GenerateModuleRequest) GetTemplate() *string {
	return s.Template
}

func (s *GenerateModuleRequest) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *GenerateModuleRequest) GetTerraformResourceType() *string {
	return s.TerraformResourceType
}

func (s *GenerateModuleRequest) SetGenerateSource(v string) *GenerateModuleRequest {
	s.GenerateSource = &v
	return s
}

func (s *GenerateModuleRequest) SetParameters(v map[string]interface{}) *GenerateModuleRequest {
	s.Parameters = v
	return s
}

func (s *GenerateModuleRequest) SetRegionId(v string) *GenerateModuleRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateModuleRequest) SetSyntax(v string) *GenerateModuleRequest {
	s.Syntax = &v
	return s
}

func (s *GenerateModuleRequest) SetTemplate(v string) *GenerateModuleRequest {
	s.Template = &v
	return s
}

func (s *GenerateModuleRequest) SetTerraformProviderVersion(v string) *GenerateModuleRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *GenerateModuleRequest) SetTerraformResourceType(v string) *GenerateModuleRequest {
	s.TerraformResourceType = &v
	return s
}

func (s *GenerateModuleRequest) Validate() error {
	return dara.Validate(s)
}
