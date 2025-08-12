// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRuleTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteMetricRuleTemplateRequest
	GetRegionId() *string
	SetTemplateId(v string) *DeleteMetricRuleTemplateRequest
	GetTemplateId() *string
}

type DeleteMetricRuleTemplateRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteMetricRuleTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMetricRuleTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteMetricRuleTemplateRequest) SetRegionId(v string) *DeleteMetricRuleTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMetricRuleTemplateRequest) SetTemplateId(v string) *DeleteMetricRuleTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteMetricRuleTemplateRequest) Validate() error {
	return dara.Validate(s)
}
