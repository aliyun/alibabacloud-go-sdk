// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataI18nValue interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DataI18nValue
	GetDescription() *string
	SetName(v string) *DataI18nValue
	GetName() *string
	SetReadme(v string) *DataI18nValue
	GetReadme() *string
}

type DataI18nValue struct {
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

func (s DataI18nValue) String() string {
	return dara.Prettify(s)
}

func (s DataI18nValue) GoString() string {
	return s.String()
}

func (s *DataI18nValue) GetDescription() *string {
	return s.Description
}

func (s *DataI18nValue) GetName() *string {
	return s.Name
}

func (s *DataI18nValue) GetReadme() *string {
	return s.Readme
}

func (s *DataI18nValue) SetDescription(v string) *DataI18nValue {
	s.Description = &v
	return s
}

func (s *DataI18nValue) SetName(v string) *DataI18nValue {
	s.Name = &v
	return s
}

func (s *DataI18nValue) SetReadme(v string) *DataI18nValue {
	s.Readme = &v
	return s
}

func (s *DataI18nValue) Validate() error {
	return dara.Validate(s)
}
