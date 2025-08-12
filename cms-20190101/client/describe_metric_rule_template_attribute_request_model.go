// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleTemplateAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeMetricRuleTemplateAttributeRequest
	GetName() *string
	SetRegionId(v string) *DescribeMetricRuleTemplateAttributeRequest
	GetRegionId() *string
	SetTemplateId(v string) *DescribeMetricRuleTemplateAttributeRequest
	GetTemplateId() *string
}

type DescribeMetricRuleTemplateAttributeRequest struct {
	// The name of the alert template. You must specify at least one of the `Name` and `TemplateId` parameters.
	//
	// For information about how to obtain the name of an alert template, see [DescribeMetricRuleTemplateList](https://help.aliyun.com/document_detail/114982.html).
	//
	// example:
	//
	// ECS_Template1
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert template. You must specify at least one of the `Name` and `TemplateId` parameters.
	//
	// For information about how to obtain the ID of an alert template, see [DescribeMetricRuleTemplateList](https://help.aliyun.com/document_detail/114982.html).
	//
	// example:
	//
	// 70****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeRequest) GetName() *string {
	return s.Name
}

func (s *DescribeMetricRuleTemplateAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMetricRuleTemplateAttributeRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeMetricRuleTemplateAttributeRequest) SetName(v string) *DescribeMetricRuleTemplateAttributeRequest {
	s.Name = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeRequest) SetRegionId(v string) *DescribeMetricRuleTemplateAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeRequest) SetTemplateId(v string) *DescribeMetricRuleTemplateAttributeRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeRequest) Validate() error {
	return dara.Validate(s)
}
