// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValuateTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaId(v string) *ValuateTemplateShrinkRequest
	GetAreaId() *string
	SetClientToken(v string) *ValuateTemplateShrinkRequest
	GetClientToken() *string
	SetInstancesShrink(v string) *ValuateTemplateShrinkRequest
	GetInstancesShrink() *string
	SetResourceGroupId(v string) *ValuateTemplateShrinkRequest
	GetResourceGroupId() *string
	SetTemplateId(v string) *ValuateTemplateShrinkRequest
	GetTemplateId() *string
	SetVariablesShrink(v string) *ValuateTemplateShrinkRequest
	GetVariablesShrink() *string
}

type ValuateTemplateShrinkRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instances to be replaced.
	InstancesShrink *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// The ID of the resource group to which the application belongs.
	//
	// example:
	//
	// rg-acfmyjt3c5om3fi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0KSHPM6SJU03TNZP
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The parameter values that are contained in the template. If the template contains no parameter values, the default values are used.
	VariablesShrink *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s ValuateTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ValuateTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ValuateTemplateShrinkRequest) GetAreaId() *string {
	return s.AreaId
}

func (s *ValuateTemplateShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ValuateTemplateShrinkRequest) GetInstancesShrink() *string {
	return s.InstancesShrink
}

func (s *ValuateTemplateShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ValuateTemplateShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ValuateTemplateShrinkRequest) GetVariablesShrink() *string {
	return s.VariablesShrink
}

func (s *ValuateTemplateShrinkRequest) SetAreaId(v string) *ValuateTemplateShrinkRequest {
	s.AreaId = &v
	return s
}

func (s *ValuateTemplateShrinkRequest) SetClientToken(v string) *ValuateTemplateShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ValuateTemplateShrinkRequest) SetInstancesShrink(v string) *ValuateTemplateShrinkRequest {
	s.InstancesShrink = &v
	return s
}

func (s *ValuateTemplateShrinkRequest) SetResourceGroupId(v string) *ValuateTemplateShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ValuateTemplateShrinkRequest) SetTemplateId(v string) *ValuateTemplateShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *ValuateTemplateShrinkRequest) SetVariablesShrink(v string) *ValuateTemplateShrinkRequest {
	s.VariablesShrink = &v
	return s
}

func (s *ValuateTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
