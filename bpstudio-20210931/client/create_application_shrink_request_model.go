// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaId(v string) *CreateApplicationShrinkRequest
	GetAreaId() *string
	SetClientToken(v string) *CreateApplicationShrinkRequest
	GetClientToken() *string
	SetConfigurationShrink(v string) *CreateApplicationShrinkRequest
	GetConfigurationShrink() *string
	SetInstancesShrink(v string) *CreateApplicationShrinkRequest
	GetInstancesShrink() *string
	SetName(v string) *CreateApplicationShrinkRequest
	GetName() *string
	SetProcessVariablesShrink(v string) *CreateApplicationShrinkRequest
	GetProcessVariablesShrink() *string
	SetResourceGroupId(v string) *CreateApplicationShrinkRequest
	GetResourceGroupId() *string
	SetTemplateId(v string) *CreateApplicationShrinkRequest
	GetTemplateId() *string
	SetVariablesShrink(v string) *CreateApplicationShrinkRequest
	GetVariablesShrink() *string
}

type CreateApplicationShrinkRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The parameters that are used to configure the application you want to create. For example, enableMonitor specifies whether to automatically create a CloudMonitor task for the application, and enableReport specifies whether to generate reports.
	//
	// example:
	//
	// {"enableMonitor":"0", "enableReport":"1"}
	ConfigurationShrink *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The instances in which you want to create the application. You can create applications in an existing virtual private cloud (VPC).
	InstancesShrink *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// The name of the application.
	//
	// 	- The application name must be unique. You can call the [ListApplication](https://www.alibabacloud.com/help/en/bp-studio/latest/api-bpstudio-2021-09-31-listapplication) operation to query the existing applications.
	//
	// 	- The application name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http:// or https://`. The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// cadt-application
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProcessVariablesShrink *string `json:"ProcessVariables,omitempty" xml:"ProcessVariables,omitempty"`
	// The ID of the resource group to which the application you want to create belongs.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0KSHPM6SJU03TNZP
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The parameter values that are contained in the template. If the template contains no parameter values, the default values are used.
	//
	// example:
	//
	// {"variable1":"1"}
	VariablesShrink *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s CreateApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationShrinkRequest) GetAreaId() *string {
	return s.AreaId
}

func (s *CreateApplicationShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateApplicationShrinkRequest) GetConfigurationShrink() *string {
	return s.ConfigurationShrink
}

func (s *CreateApplicationShrinkRequest) GetInstancesShrink() *string {
	return s.InstancesShrink
}

func (s *CreateApplicationShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateApplicationShrinkRequest) GetProcessVariablesShrink() *string {
	return s.ProcessVariablesShrink
}

func (s *CreateApplicationShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateApplicationShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateApplicationShrinkRequest) GetVariablesShrink() *string {
	return s.VariablesShrink
}

func (s *CreateApplicationShrinkRequest) SetAreaId(v string) *CreateApplicationShrinkRequest {
	s.AreaId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetClientToken(v string) *CreateApplicationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetConfigurationShrink(v string) *CreateApplicationShrinkRequest {
	s.ConfigurationShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetInstancesShrink(v string) *CreateApplicationShrinkRequest {
	s.InstancesShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetName(v string) *CreateApplicationShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetProcessVariablesShrink(v string) *CreateApplicationShrinkRequest {
	s.ProcessVariablesShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetResourceGroupId(v string) *CreateApplicationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetTemplateId(v string) *CreateApplicationShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetVariablesShrink(v string) *CreateApplicationShrinkRequest {
	s.VariablesShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
