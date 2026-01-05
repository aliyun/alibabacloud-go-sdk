// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateBody(v string) *CreateTemplateRequest
	GetTemplateBody() *string
	SetTemplateType(v string) *CreateTemplateRequest
	GetTemplateType() *string
	SetTerraformVariables(v []*CreateTemplateRequestTerraformVariables) *CreateTemplateRequest
	GetTerraformVariables() []*CreateTemplateRequestTerraformVariables
}

type CreateTemplateRequest struct {
	// The content of the template.
	//
	// For more information about the template syntax, see [Structure of Terraform templates](https://help.aliyun.com/document_detail/184397.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "ROSTemplateFormatVersion": "2015-09-01",
	//
	//   "Transform": "Aliyun::Terraform-v1.1",
	//
	//   "Workspace": {
	//
	//     "main.tf": "variable  \\"name\\" {  default = \\"auto_provisioning_group\\"}"
	//
	//   },
	//
	//   "Outputs": {}
	//
	// }
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The type of the product template. Valid values:
	//
	// 	- RosTerraformTemplate: the Terraform template that is supported by Resource Orchestration Service (ROS).
	//
	// 	- RosStandardTemplate: the standard ROS template.
	//
	// This parameter is required.
	//
	// example:
	//
	// RosTerraformTemplate
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The variable settings of the Terraform template. You can configure the variables in a structured manner. Service Catalog applies the variable settings to the template.
	//
	// >  The variables must be defined in the Terraform template.
	TerraformVariables []*CreateTemplateRequestTerraformVariables `json:"TerraformVariables,omitempty" xml:"TerraformVariables,omitempty" type:"Repeated"`
}

func (s CreateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *CreateTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateTemplateRequest) GetTerraformVariables() []*CreateTemplateRequestTerraformVariables {
	return s.TerraformVariables
}

func (s *CreateTemplateRequest) SetTemplateBody(v string) *CreateTemplateRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateType(v string) *CreateTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateTemplateRequest) SetTerraformVariables(v []*CreateTemplateRequestTerraformVariables) *CreateTemplateRequest {
	s.TerraformVariables = v
	return s
}

func (s *CreateTemplateRequest) Validate() error {
	if s.TerraformVariables != nil {
		for _, item := range s.TerraformVariables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTemplateRequestTerraformVariables struct {
	// The description of the variable.
	//
	// For more information about the format of variable descriptions, see [Methods and suggestions for Terraform code development](https://help.aliyun.com/document_detail/322216.html).
	//
	// example:
	//
	// { "Label": { "en": "Instance Type" }, "AllowedValues": [ "ecs.s6-c1m1.small", "ecs.s6-c1m2.large", "ecs.s6-c1m2.xlarge" ] }
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the variable.
	//
	// example:
	//
	// instance_type
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
}

func (s CreateTemplateRequestTerraformVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequestTerraformVariables) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequestTerraformVariables) GetDescription() *string {
	return s.Description
}

func (s *CreateTemplateRequestTerraformVariables) GetVariableName() *string {
	return s.VariableName
}

func (s *CreateTemplateRequestTerraformVariables) SetDescription(v string) *CreateTemplateRequestTerraformVariables {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequestTerraformVariables) SetVariableName(v string) *CreateTemplateRequestTerraformVariables {
	s.VariableName = &v
	return s
}

func (s *CreateTemplateRequestTerraformVariables) Validate() error {
	return dara.Validate(s)
}
