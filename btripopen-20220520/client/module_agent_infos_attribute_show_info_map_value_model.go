// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleAgentInfosAttributeShowInfoMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModuleAgentInfosAttributeShowInfoMapValue
	GetCode() *string
	SetName(v string) *ModuleAgentInfosAttributeShowInfoMapValue
	GetName() *string
	SetText(v string) *ModuleAgentInfosAttributeShowInfoMapValue
	GetText() *string
}

type ModuleAgentInfosAttributeShowInfoMapValue struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ModuleAgentInfosAttributeShowInfoMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleAgentInfosAttributeShowInfoMapValue) GoString() string {
	return s.String()
}

func (s *ModuleAgentInfosAttributeShowInfoMapValue) GetCode() *string {
	return s.Code
}

func (s *ModuleAgentInfosAttributeShowInfoMapValue) GetName() *string {
	return s.Name
}

func (s *ModuleAgentInfosAttributeShowInfoMapValue) GetText() *string {
	return s.Text
}

func (s *ModuleAgentInfosAttributeShowInfoMapValue) SetCode(v string) *ModuleAgentInfosAttributeShowInfoMapValue {
	s.Code = &v
	return s
}

func (s *ModuleAgentInfosAttributeShowInfoMapValue) SetName(v string) *ModuleAgentInfosAttributeShowInfoMapValue {
	s.Name = &v
	return s
}

func (s *ModuleAgentInfosAttributeShowInfoMapValue) SetText(v string) *ModuleAgentInfosAttributeShowInfoMapValue {
	s.Text = &v
	return s
}

func (s *ModuleAgentInfosAttributeShowInfoMapValue) Validate() error {
	return dara.Validate(s)
}
