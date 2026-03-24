// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformAction interface {
	dara.Model
	String() string
	GoString() string
	SetFilterSetting(v *FilterSetting) *TransformAction
	GetFilterSetting() *FilterSetting
	SetLabelKey(v string) *TransformAction
	GetLabelKey() *string
	SetMapping(v map[string]*string) *TransformAction
	GetMapping() map[string]*string
	SetRegExp(v string) *TransformAction
	GetRegExp() *string
	SetSource(v string) *TransformAction
	GetSource() *string
	SetTarget(v string) *TransformAction
	GetTarget() *string
	SetType(v string) *TransformAction
	GetType() *string
	SetValue(v string) *TransformAction
	GetValue() *string
	SetVariable(v string) *TransformAction
	GetVariable() *string
}

type TransformAction struct {
	// 筛选配置
	FilterSetting *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	// 标签名
	//
	// example:
	//
	// labelkey1
	LabelKey *string `json:"labelKey,omitempty" xml:"labelKey,omitempty"`
	// Mapping配置。
	Mapping map[string]*string `json:"mapping,omitempty" xml:"mapping,omitempty"`
	// 正则表达式
	//
	// example:
	//
	// (.*):(.*)
	RegExp *string `json:"regExp,omitempty" xml:"regExp,omitempty"`
	// 引用路径
	//
	// example:
	//
	// data.subject
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 目标位置
	//
	// example:
	//
	// SUBJECT
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
	// 转换类型
	//
	// example:
	//
	// SET_FIELD
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 设置的值
	//
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 变量名
	//
	// example:
	//
	// var1
	Variable *string `json:"variable,omitempty" xml:"variable,omitempty"`
}

func (s TransformAction) String() string {
	return dara.Prettify(s)
}

func (s TransformAction) GoString() string {
	return s.String()
}

func (s *TransformAction) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *TransformAction) GetLabelKey() *string {
	return s.LabelKey
}

func (s *TransformAction) GetMapping() map[string]*string {
	return s.Mapping
}

func (s *TransformAction) GetRegExp() *string {
	return s.RegExp
}

func (s *TransformAction) GetSource() *string {
	return s.Source
}

func (s *TransformAction) GetTarget() *string {
	return s.Target
}

func (s *TransformAction) GetType() *string {
	return s.Type
}

func (s *TransformAction) GetValue() *string {
	return s.Value
}

func (s *TransformAction) GetVariable() *string {
	return s.Variable
}

func (s *TransformAction) SetFilterSetting(v *FilterSetting) *TransformAction {
	s.FilterSetting = v
	return s
}

func (s *TransformAction) SetLabelKey(v string) *TransformAction {
	s.LabelKey = &v
	return s
}

func (s *TransformAction) SetMapping(v map[string]*string) *TransformAction {
	s.Mapping = v
	return s
}

func (s *TransformAction) SetRegExp(v string) *TransformAction {
	s.RegExp = &v
	return s
}

func (s *TransformAction) SetSource(v string) *TransformAction {
	s.Source = &v
	return s
}

func (s *TransformAction) SetTarget(v string) *TransformAction {
	s.Target = &v
	return s
}

func (s *TransformAction) SetType(v string) *TransformAction {
	s.Type = &v
	return s
}

func (s *TransformAction) SetValue(v string) *TransformAction {
	s.Value = &v
	return s
}

func (s *TransformAction) SetVariable(v string) *TransformAction {
	s.Variable = &v
	return s
}

func (s *TransformAction) Validate() error {
	if s.FilterSetting != nil {
		if err := s.FilterSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}
