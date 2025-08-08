// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLinkageAttributesTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAreaId(v string) *GetLinkageAttributesTemplateShrinkRequest
	GetAreaId() *string
	SetInstancesShrink(v string) *GetLinkageAttributesTemplateShrinkRequest
	GetInstancesShrink() *string
	SetMaxResults(v int64) *GetLinkageAttributesTemplateShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *GetLinkageAttributesTemplateShrinkRequest
	GetNextToken() *string
	SetTargetVariable(v string) *GetLinkageAttributesTemplateShrinkRequest
	GetTargetVariable() *string
	SetTemplateId(v string) *GetLinkageAttributesTemplateShrinkRequest
	GetTemplateId() *string
	SetVariablesShrink(v string) *GetLinkageAttributesTemplateShrinkRequest
	GetVariablesShrink() *string
}

type GetLinkageAttributesTemplateShrinkRequest struct {
	// example:
	//
	// cn-hangzhou
	AreaId          *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	InstancesShrink *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
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
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	VariablesShrink *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s GetLinkageAttributesTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLinkageAttributesTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetLinkageAttributesTemplateShrinkRequest) GetAreaId() *string {
	return s.AreaId
}

func (s *GetLinkageAttributesTemplateShrinkRequest) GetInstancesShrink() *string {
	return s.InstancesShrink
}

func (s *GetLinkageAttributesTemplateShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetLinkageAttributesTemplateShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetLinkageAttributesTemplateShrinkRequest) GetTargetVariable() *string {
	return s.TargetVariable
}

func (s *GetLinkageAttributesTemplateShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetLinkageAttributesTemplateShrinkRequest) GetVariablesShrink() *string {
	return s.VariablesShrink
}

func (s *GetLinkageAttributesTemplateShrinkRequest) SetAreaId(v string) *GetLinkageAttributesTemplateShrinkRequest {
	s.AreaId = &v
	return s
}

func (s *GetLinkageAttributesTemplateShrinkRequest) SetInstancesShrink(v string) *GetLinkageAttributesTemplateShrinkRequest {
	s.InstancesShrink = &v
	return s
}

func (s *GetLinkageAttributesTemplateShrinkRequest) SetMaxResults(v int64) *GetLinkageAttributesTemplateShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *GetLinkageAttributesTemplateShrinkRequest) SetNextToken(v string) *GetLinkageAttributesTemplateShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *GetLinkageAttributesTemplateShrinkRequest) SetTargetVariable(v string) *GetLinkageAttributesTemplateShrinkRequest {
	s.TargetVariable = &v
	return s
}

func (s *GetLinkageAttributesTemplateShrinkRequest) SetTemplateId(v string) *GetLinkageAttributesTemplateShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *GetLinkageAttributesTemplateShrinkRequest) SetVariablesShrink(v string) *GetLinkageAttributesTemplateShrinkRequest {
	s.VariablesShrink = &v
	return s
}

func (s *GetLinkageAttributesTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
