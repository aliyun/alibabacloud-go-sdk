// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleAgentInfoAttributeShowInfoMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModuleAgentInfoAttributeShowInfoMapValue
	GetCode() *string
	SetName(v string) *ModuleAgentInfoAttributeShowInfoMapValue
	GetName() *string
	SetText(v string) *ModuleAgentInfoAttributeShowInfoMapValue
	GetText() *string
}

type ModuleAgentInfoAttributeShowInfoMapValue struct {
	// example:
	//
	// FIRST_BAGGAGE
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ModuleAgentInfoAttributeShowInfoMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleAgentInfoAttributeShowInfoMapValue) GoString() string {
	return s.String()
}

func (s *ModuleAgentInfoAttributeShowInfoMapValue) GetCode() *string {
	return s.Code
}

func (s *ModuleAgentInfoAttributeShowInfoMapValue) GetName() *string {
	return s.Name
}

func (s *ModuleAgentInfoAttributeShowInfoMapValue) GetText() *string {
	return s.Text
}

func (s *ModuleAgentInfoAttributeShowInfoMapValue) SetCode(v string) *ModuleAgentInfoAttributeShowInfoMapValue {
	s.Code = &v
	return s
}

func (s *ModuleAgentInfoAttributeShowInfoMapValue) SetName(v string) *ModuleAgentInfoAttributeShowInfoMapValue {
	s.Name = &v
	return s
}

func (s *ModuleAgentInfoAttributeShowInfoMapValue) SetText(v string) *ModuleAgentInfoAttributeShowInfoMapValue {
	s.Text = &v
	return s
}

func (s *ModuleAgentInfoAttributeShowInfoMapValue) Validate() error {
	return dara.Validate(s)
}
