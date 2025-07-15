// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWritingStyleTemplateDefine interface {
	dara.Model
	String() string
	GoString() string
	SetExample(v []*WritingStyleTemplateDefineExample) *WritingStyleTemplateDefine
	GetExample() []*WritingStyleTemplateDefineExample
	SetFields(v []*WritingStyleTemplateField) *WritingStyleTemplateDefine
	GetFields() []*WritingStyleTemplateField
}

type WritingStyleTemplateDefine struct {
	Example []*WritingStyleTemplateDefineExample `json:"Example,omitempty" xml:"Example,omitempty" type:"Repeated"`
	Fields  []*WritingStyleTemplateField         `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
}

func (s WritingStyleTemplateDefine) String() string {
	return dara.Prettify(s)
}

func (s WritingStyleTemplateDefine) GoString() string {
	return s.String()
}

func (s *WritingStyleTemplateDefine) GetExample() []*WritingStyleTemplateDefineExample {
	return s.Example
}

func (s *WritingStyleTemplateDefine) GetFields() []*WritingStyleTemplateField {
	return s.Fields
}

func (s *WritingStyleTemplateDefine) SetExample(v []*WritingStyleTemplateDefineExample) *WritingStyleTemplateDefine {
	s.Example = v
	return s
}

func (s *WritingStyleTemplateDefine) SetFields(v []*WritingStyleTemplateField) *WritingStyleTemplateDefine {
	s.Fields = v
	return s
}

func (s *WritingStyleTemplateDefine) Validate() error {
	return dara.Validate(s)
}

type WritingStyleTemplateDefineExample struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s WritingStyleTemplateDefineExample) String() string {
	return dara.Prettify(s)
}

func (s WritingStyleTemplateDefineExample) GoString() string {
	return s.String()
}

func (s *WritingStyleTemplateDefineExample) GetKey() *string {
	return s.Key
}

func (s *WritingStyleTemplateDefineExample) GetValue() *string {
	return s.Value
}

func (s *WritingStyleTemplateDefineExample) SetKey(v string) *WritingStyleTemplateDefineExample {
	s.Key = &v
	return s
}

func (s *WritingStyleTemplateDefineExample) SetValue(v string) *WritingStyleTemplateDefineExample {
	s.Value = &v
	return s
}

func (s *WritingStyleTemplateDefineExample) Validate() error {
	return dara.Validate(s)
}
