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
	// example:
	//
	// Resource
	GenerateSource *string                `json:"generateSource,omitempty" xml:"generateSource,omitempty"`
	Parameters     map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// terraform
	Syntax *string `json:"syntax,omitempty" xml:"syntax,omitempty"`
	// example:
	//
	// generateSource ==
	Template *string `json:"template,omitempty" xml:"template,omitempty"`
	// example:
	//
	// 1.189.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// example:
	//
	// alicloud_db_instance
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
