// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataDesensPlanTemplateValue interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DataDesensPlanTemplateValue
	GetName() *string
	SetSupportWaterMark(v bool) *DataDesensPlanTemplateValue
	GetSupportWaterMark() *bool
	SetExtParamTemplate(v []interface{}) *DataDesensPlanTemplateValue
	GetExtParamTemplate() []interface{}
}

type DataDesensPlanTemplateValue struct {
	// The name of the data masking method.
	//
	// example:
	//
	// hash
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether a watermark is added. Valid values:
	//
	// 	- true: allow
	//
	// 	- false: disallow
	//
	// example:
	//
	// true
	SupportWaterMark *bool `json:"SupportWaterMark,omitempty" xml:"SupportWaterMark,omitempty"`
	// The data masking parameters and their descriptions.
	ExtParamTemplate []interface{} `json:"ExtParamTemplate,omitempty" xml:"ExtParamTemplate,omitempty" type:"Repeated"`
}

func (s DataDesensPlanTemplateValue) String() string {
	return dara.Prettify(s)
}

func (s DataDesensPlanTemplateValue) GoString() string {
	return s.String()
}

func (s *DataDesensPlanTemplateValue) GetName() *string {
	return s.Name
}

func (s *DataDesensPlanTemplateValue) GetSupportWaterMark() *bool {
	return s.SupportWaterMark
}

func (s *DataDesensPlanTemplateValue) GetExtParamTemplate() []interface{} {
	return s.ExtParamTemplate
}

func (s *DataDesensPlanTemplateValue) SetName(v string) *DataDesensPlanTemplateValue {
	s.Name = &v
	return s
}

func (s *DataDesensPlanTemplateValue) SetSupportWaterMark(v bool) *DataDesensPlanTemplateValue {
	s.SupportWaterMark = &v
	return s
}

func (s *DataDesensPlanTemplateValue) SetExtParamTemplate(v []interface{}) *DataDesensPlanTemplateValue {
	s.ExtParamTemplate = v
	return s
}

func (s *DataDesensPlanTemplateValue) Validate() error {
	return dara.Validate(s)
}
