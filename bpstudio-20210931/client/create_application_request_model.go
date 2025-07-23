// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaId(v string) *CreateApplicationRequest
	GetAreaId() *string
	SetClientToken(v string) *CreateApplicationRequest
	GetClientToken() *string
	SetConfiguration(v map[string]*string) *CreateApplicationRequest
	GetConfiguration() map[string]*string
	SetInstances(v []*CreateApplicationRequestInstances) *CreateApplicationRequest
	GetInstances() []*CreateApplicationRequestInstances
	SetName(v string) *CreateApplicationRequest
	GetName() *string
	SetProcessVariables(v map[string]interface{}) *CreateApplicationRequest
	GetProcessVariables() map[string]interface{}
	SetResourceGroupId(v string) *CreateApplicationRequest
	GetResourceGroupId() *string
	SetTemplateId(v string) *CreateApplicationRequest
	GetTemplateId() *string
	SetVariables(v map[string]interface{}) *CreateApplicationRequest
	GetVariables() map[string]interface{}
}

type CreateApplicationRequest struct {
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
	Configuration map[string]*string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The instances in which you want to create the application. You can create applications in an existing virtual private cloud (VPC).
	Instances []*CreateApplicationRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
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
	Name             *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	ProcessVariables map[string]interface{} `json:"ProcessVariables,omitempty" xml:"ProcessVariables,omitempty"`
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
	Variables map[string]interface{} `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) GetAreaId() *string {
	return s.AreaId
}

func (s *CreateApplicationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateApplicationRequest) GetConfiguration() map[string]*string {
	return s.Configuration
}

func (s *CreateApplicationRequest) GetInstances() []*CreateApplicationRequestInstances {
	return s.Instances
}

func (s *CreateApplicationRequest) GetName() *string {
	return s.Name
}

func (s *CreateApplicationRequest) GetProcessVariables() map[string]interface{} {
	return s.ProcessVariables
}

func (s *CreateApplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateApplicationRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateApplicationRequest) GetVariables() map[string]interface{} {
	return s.Variables
}

func (s *CreateApplicationRequest) SetAreaId(v string) *CreateApplicationRequest {
	s.AreaId = &v
	return s
}

func (s *CreateApplicationRequest) SetClientToken(v string) *CreateApplicationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationRequest) SetConfiguration(v map[string]*string) *CreateApplicationRequest {
	s.Configuration = v
	return s
}

func (s *CreateApplicationRequest) SetInstances(v []*CreateApplicationRequestInstances) *CreateApplicationRequest {
	s.Instances = v
	return s
}

func (s *CreateApplicationRequest) SetName(v string) *CreateApplicationRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationRequest) SetProcessVariables(v map[string]interface{}) *CreateApplicationRequest {
	s.ProcessVariables = v
	return s
}

func (s *CreateApplicationRequest) SetResourceGroupId(v string) *CreateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationRequest) SetTemplateId(v string) *CreateApplicationRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateApplicationRequest) SetVariables(v map[string]interface{}) *CreateApplicationRequest {
	s.Variables = v
	return s
}

func (s *CreateApplicationRequest) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationRequestInstances struct {
	// The ID of the instance.
	//
	// example:
	//
	// vpc-bp1q56trhtaq40vlq5ojm
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// vpc
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// vpc
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s CreateApplicationRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequestInstances) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestInstances) GetId() *string {
	return s.Id
}

func (s *CreateApplicationRequestInstances) GetNodeName() *string {
	return s.NodeName
}

func (s *CreateApplicationRequestInstances) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateApplicationRequestInstances) SetId(v string) *CreateApplicationRequestInstances {
	s.Id = &v
	return s
}

func (s *CreateApplicationRequestInstances) SetNodeName(v string) *CreateApplicationRequestInstances {
	s.NodeName = &v
	return s
}

func (s *CreateApplicationRequestInstances) SetNodeType(v string) *CreateApplicationRequestInstances {
	s.NodeType = &v
	return s
}

func (s *CreateApplicationRequestInstances) Validate() error {
	return dara.Validate(s)
}
