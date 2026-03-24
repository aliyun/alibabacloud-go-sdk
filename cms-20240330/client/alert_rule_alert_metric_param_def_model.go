// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleAlertMetricParamDef interface {
	dara.Model
	String() string
	GoString() string
	SetMaxWidth(v int32) *AlertRuleAlertMetricParamDef
	GetMaxWidth() *int32
	SetMinWidth(v int32) *AlertRuleAlertMetricParamDef
	GetMinWidth() *int32
	SetName(v string) *AlertRuleAlertMetricParamDef
	GetName() *string
	SetPlaceholderCn(v string) *AlertRuleAlertMetricParamDef
	GetPlaceholderCn() *string
	SetPlaceholderEn(v string) *AlertRuleAlertMetricParamDef
	GetPlaceholderEn() *string
	SetType(v string) *AlertRuleAlertMetricParamDef
	GetType() *string
	SetValue(v string) *AlertRuleAlertMetricParamDef
	GetValue() *string
	SetValues(v []*AlertRuleAlertMetricParamDefValues) *AlertRuleAlertMetricParamDef
	GetValues() []*AlertRuleAlertMetricParamDefValues
}

type AlertRuleAlertMetricParamDef struct {
	// 输入框的最大宽度，仅对SELECT_PARAM、INPUT_PARAM生效
	//
	// example:
	//
	// 200
	MaxWidth *int32 `json:"maxWidth,omitempty" xml:"maxWidth,omitempty"`
	// 输入框的最小宽度，仅对SELECT_PARAM、INPUT_PARAM生效
	//
	// example:
	//
	// 100
	MinWidth *int32 `json:"minWidth,omitempty" xml:"minWidth,omitempty"`
	// 名称
	//
	// example:
	//
	// env
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 仅对INPUT_PARAM有效。用于前端展示的中文展位符
	//
	// example:
	//
	// 请输入值
	PlaceholderCn *string `json:"placeholderCn,omitempty" xml:"placeholderCn,omitempty"`
	// 仅对INPUT_PARAM有效。用于前端展示的英文展位符
	//
	// example:
	//
	// Enter value
	PlaceholderEn *string `json:"placeholderEn,omitempty" xml:"placeholderEn,omitempty"`
	// ● TEXT_PARAM: 只读文本参数，由后台定义，前端不显示用户输入控件
	//
	// ● INPUT_PARAM：输入框参数
	//
	// ● SELECT_PARAM：选择框参数
	//
	// example:
	//
	// TEXT_PARAM
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 值
	//
	// example:
	//
	// staging
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 仅对SELECT_PARAM有效。  下拉列表的可选值列表。
	Values []*AlertRuleAlertMetricParamDefValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s AlertRuleAlertMetricParamDef) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleAlertMetricParamDef) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricParamDef) GetMaxWidth() *int32 {
	return s.MaxWidth
}

func (s *AlertRuleAlertMetricParamDef) GetMinWidth() *int32 {
	return s.MinWidth
}

func (s *AlertRuleAlertMetricParamDef) GetName() *string {
	return s.Name
}

func (s *AlertRuleAlertMetricParamDef) GetPlaceholderCn() *string {
	return s.PlaceholderCn
}

func (s *AlertRuleAlertMetricParamDef) GetPlaceholderEn() *string {
	return s.PlaceholderEn
}

func (s *AlertRuleAlertMetricParamDef) GetType() *string {
	return s.Type
}

func (s *AlertRuleAlertMetricParamDef) GetValue() *string {
	return s.Value
}

func (s *AlertRuleAlertMetricParamDef) GetValues() []*AlertRuleAlertMetricParamDefValues {
	return s.Values
}

func (s *AlertRuleAlertMetricParamDef) SetMaxWidth(v int32) *AlertRuleAlertMetricParamDef {
	s.MaxWidth = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetMinWidth(v int32) *AlertRuleAlertMetricParamDef {
	s.MinWidth = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetName(v string) *AlertRuleAlertMetricParamDef {
	s.Name = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetPlaceholderCn(v string) *AlertRuleAlertMetricParamDef {
	s.PlaceholderCn = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetPlaceholderEn(v string) *AlertRuleAlertMetricParamDef {
	s.PlaceholderEn = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetType(v string) *AlertRuleAlertMetricParamDef {
	s.Type = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetValue(v string) *AlertRuleAlertMetricParamDef {
	s.Value = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetValues(v []*AlertRuleAlertMetricParamDefValues) *AlertRuleAlertMetricParamDef {
	s.Values = v
	return s
}

func (s *AlertRuleAlertMetricParamDef) Validate() error {
	if s.Values != nil {
		for _, item := range s.Values {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertRuleAlertMetricParamDefValues struct {
	// 选项的中文显示名称
	//
	// example:
	//
	// 测试环境
	LabelCn *string `json:"labelCn,omitempty" xml:"labelCn,omitempty"`
	// 选项的英文显示名称
	//
	// example:
	//
	// Staging
	LabelEn *string `json:"labelEn,omitempty" xml:"labelEn,omitempty"`
	// 值
	//
	// example:
	//
	// staging
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleAlertMetricParamDefValues) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleAlertMetricParamDefValues) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricParamDefValues) GetLabelCn() *string {
	return s.LabelCn
}

func (s *AlertRuleAlertMetricParamDefValues) GetLabelEn() *string {
	return s.LabelEn
}

func (s *AlertRuleAlertMetricParamDefValues) GetValue() *string {
	return s.Value
}

func (s *AlertRuleAlertMetricParamDefValues) SetLabelCn(v string) *AlertRuleAlertMetricParamDefValues {
	s.LabelCn = &v
	return s
}

func (s *AlertRuleAlertMetricParamDefValues) SetLabelEn(v string) *AlertRuleAlertMetricParamDefValues {
	s.LabelEn = &v
	return s
}

func (s *AlertRuleAlertMetricParamDefValues) SetValue(v string) *AlertRuleAlertMetricParamDefValues {
	s.Value = &v
	return s
}

func (s *AlertRuleAlertMetricParamDefValues) Validate() error {
	return dara.Validate(s)
}
