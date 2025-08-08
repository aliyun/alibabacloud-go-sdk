// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLinkageAttributesTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaId(v string) *GetLinkageAttributesTemplateRequest
	GetAreaId() *string
	SetInstances(v []*GetLinkageAttributesTemplateRequestInstances) *GetLinkageAttributesTemplateRequest
	GetInstances() []*GetLinkageAttributesTemplateRequestInstances
	SetMaxResults(v int64) *GetLinkageAttributesTemplateRequest
	GetMaxResults() *int64
	SetNextToken(v string) *GetLinkageAttributesTemplateRequest
	GetNextToken() *string
	SetTargetVariable(v string) *GetLinkageAttributesTemplateRequest
	GetTargetVariable() *string
	SetTemplateId(v string) *GetLinkageAttributesTemplateRequest
	GetTemplateId() *string
	SetVariables(v map[string]interface{}) *GetLinkageAttributesTemplateRequest
	GetVariables() map[string]interface{}
}

type GetLinkageAttributesTemplateRequest struct {
	// example:
	//
	// cn-hangzhou
	AreaId    *string                                         `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	Instances []*GetLinkageAttributesTemplateRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ${instance_type}
	TargetVariable *string `json:"TargetVariable,omitempty" xml:"TargetVariable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// XFKR6WYRVS24S07R
	TemplateId *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Variables  map[string]interface{} `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s GetLinkageAttributesTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLinkageAttributesTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetLinkageAttributesTemplateRequest) GetAreaId() *string {
	return s.AreaId
}

func (s *GetLinkageAttributesTemplateRequest) GetInstances() []*GetLinkageAttributesTemplateRequestInstances {
	return s.Instances
}

func (s *GetLinkageAttributesTemplateRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetLinkageAttributesTemplateRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetLinkageAttributesTemplateRequest) GetTargetVariable() *string {
	return s.TargetVariable
}

func (s *GetLinkageAttributesTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLinkageAttributesTemplateRequest) GetVariables() map[string]interface{} {
	return s.Variables
}

func (s *GetLinkageAttributesTemplateRequest) SetAreaId(v string) *GetLinkageAttributesTemplateRequest {
	s.AreaId = &v
	return s
}

func (s *GetLinkageAttributesTemplateRequest) SetInstances(v []*GetLinkageAttributesTemplateRequestInstances) *GetLinkageAttributesTemplateRequest {
	s.Instances = v
	return s
}

func (s *GetLinkageAttributesTemplateRequest) SetMaxResults(v int64) *GetLinkageAttributesTemplateRequest {
	s.MaxResults = &v
	return s
}

func (s *GetLinkageAttributesTemplateRequest) SetNextToken(v string) *GetLinkageAttributesTemplateRequest {
	s.NextToken = &v
	return s
}

func (s *GetLinkageAttributesTemplateRequest) SetTargetVariable(v string) *GetLinkageAttributesTemplateRequest {
	s.TargetVariable = &v
	return s
}

func (s *GetLinkageAttributesTemplateRequest) SetTemplateId(v string) *GetLinkageAttributesTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetLinkageAttributesTemplateRequest) SetVariables(v map[string]interface{}) *GetLinkageAttributesTemplateRequest {
	s.Variables = v
	return s
}

func (s *GetLinkageAttributesTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type GetLinkageAttributesTemplateRequestInstances struct {
	// example:
	//
	// vpc-**1q56**taq40*****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// vpc
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// vpc
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s GetLinkageAttributesTemplateRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s GetLinkageAttributesTemplateRequestInstances) GoString() string {
	return s.String()
}

func (s *GetLinkageAttributesTemplateRequestInstances) GetId() *string {
	return s.Id
}

func (s *GetLinkageAttributesTemplateRequestInstances) GetNodeName() *string {
	return s.NodeName
}

func (s *GetLinkageAttributesTemplateRequestInstances) GetNodeType() *string {
	return s.NodeType
}

func (s *GetLinkageAttributesTemplateRequestInstances) SetId(v string) *GetLinkageAttributesTemplateRequestInstances {
	s.Id = &v
	return s
}

func (s *GetLinkageAttributesTemplateRequestInstances) SetNodeName(v string) *GetLinkageAttributesTemplateRequestInstances {
	s.NodeName = &v
	return s
}

func (s *GetLinkageAttributesTemplateRequestInstances) SetNodeType(v string) *GetLinkageAttributesTemplateRequestInstances {
	s.NodeType = &v
	return s
}

func (s *GetLinkageAttributesTemplateRequestInstances) Validate() error {
	return dara.Validate(s)
}
