// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValuateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaId(v string) *ValuateTemplateRequest
	GetAreaId() *string
	SetClientToken(v string) *ValuateTemplateRequest
	GetClientToken() *string
	SetInstances(v []*ValuateTemplateRequestInstances) *ValuateTemplateRequest
	GetInstances() []*ValuateTemplateRequestInstances
	SetResourceGroupId(v string) *ValuateTemplateRequest
	GetResourceGroupId() *string
	SetTemplateId(v string) *ValuateTemplateRequest
	GetTemplateId() *string
	SetVariables(v map[string]interface{}) *ValuateTemplateRequest
	GetVariables() map[string]interface{}
}

type ValuateTemplateRequest struct {
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
	Instances []*ValuateTemplateRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
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
	Variables map[string]interface{} `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s ValuateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ValuateTemplateRequest) GoString() string {
	return s.String()
}

func (s *ValuateTemplateRequest) GetAreaId() *string {
	return s.AreaId
}

func (s *ValuateTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ValuateTemplateRequest) GetInstances() []*ValuateTemplateRequestInstances {
	return s.Instances
}

func (s *ValuateTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ValuateTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ValuateTemplateRequest) GetVariables() map[string]interface{} {
	return s.Variables
}

func (s *ValuateTemplateRequest) SetAreaId(v string) *ValuateTemplateRequest {
	s.AreaId = &v
	return s
}

func (s *ValuateTemplateRequest) SetClientToken(v string) *ValuateTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *ValuateTemplateRequest) SetInstances(v []*ValuateTemplateRequestInstances) *ValuateTemplateRequest {
	s.Instances = v
	return s
}

func (s *ValuateTemplateRequest) SetResourceGroupId(v string) *ValuateTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ValuateTemplateRequest) SetTemplateId(v string) *ValuateTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ValuateTemplateRequest) SetVariables(v map[string]interface{}) *ValuateTemplateRequest {
	s.Variables = v
	return s
}

func (s *ValuateTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type ValuateTemplateRequestInstances struct {
	// The instance ID.
	//
	// example:
	//
	// vpc-bp1q56trhtaq40vlq5oj
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the application instance that is displayed on the diagram.
	//
	// example:
	//
	// ecs
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s ValuateTemplateRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s ValuateTemplateRequestInstances) GoString() string {
	return s.String()
}

func (s *ValuateTemplateRequestInstances) GetId() *string {
	return s.Id
}

func (s *ValuateTemplateRequestInstances) GetNodeName() *string {
	return s.NodeName
}

func (s *ValuateTemplateRequestInstances) GetNodeType() *string {
	return s.NodeType
}

func (s *ValuateTemplateRequestInstances) SetId(v string) *ValuateTemplateRequestInstances {
	s.Id = &v
	return s
}

func (s *ValuateTemplateRequestInstances) SetNodeName(v string) *ValuateTemplateRequestInstances {
	s.NodeName = &v
	return s
}

func (s *ValuateTemplateRequestInstances) SetNodeType(v string) *ValuateTemplateRequestInstances {
	s.NodeType = &v
	return s
}

func (s *ValuateTemplateRequestInstances) Validate() error {
	return dara.Validate(s)
}
