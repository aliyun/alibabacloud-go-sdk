// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iItemsI18nValue interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ItemsI18nValue
	GetDescription() *string
	SetName(v string) *ItemsI18nValue
	GetName() *string
	SetReadme(v string) *ItemsI18nValue
	GetReadme() *string
}

type ItemsI18nValue struct {
	// The MCP service description in the corresponding language.
	//
	// example:
	//
	// An MCP service for querying knowledge bases
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The MCP marketplace template name in the corresponding language.
	//
	// example:
	//
	// Knowledge Base
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The MCP marketplace template usage instructions in the corresponding language.
	//
	// example:
	//
	// # Knowledge Base\\nKnowledge base query service
	Readme *string `json:"readme,omitempty" xml:"readme,omitempty"`
}

func (s ItemsI18nValue) String() string {
	return dara.Prettify(s)
}

func (s ItemsI18nValue) GoString() string {
	return s.String()
}

func (s *ItemsI18nValue) GetDescription() *string {
	return s.Description
}

func (s *ItemsI18nValue) GetName() *string {
	return s.Name
}

func (s *ItemsI18nValue) GetReadme() *string {
	return s.Readme
}

func (s *ItemsI18nValue) SetDescription(v string) *ItemsI18nValue {
	s.Description = &v
	return s
}

func (s *ItemsI18nValue) SetName(v string) *ItemsI18nValue {
	s.Name = &v
	return s
}

func (s *ItemsI18nValue) SetReadme(v string) *ItemsI18nValue {
	s.Readme = &v
	return s
}

func (s *ItemsI18nValue) Validate() error {
	return dara.Validate(s)
}
