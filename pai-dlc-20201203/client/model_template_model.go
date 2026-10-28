// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetCollections(v string) *ModelTemplate
	GetCollections() *string
	SetModelName(v string) *ModelTemplate
	GetModelName() *string
	SetProvider(v string) *ModelTemplate
	GetProvider() *string
}

type ModelTemplate struct {
	// The collection where the model template resides. Default value: QuickStartTemplate.
	//
	// example:
	//
	// QuickStartTemplate
	Collections *string `json:"Collections,omitempty" xml:"Collections,omitempty"`
	// The name of the model template.
	//
	// example:
	//
	// Qwen3-VL-32B-Instruct-PAI_template
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The provider of the model template. Default value: pai, which indicates that the model template is officially provided by PAI.
	//
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
}

func (s ModelTemplate) String() string {
	return dara.Prettify(s)
}

func (s ModelTemplate) GoString() string {
	return s.String()
}

func (s *ModelTemplate) GetCollections() *string {
	return s.Collections
}

func (s *ModelTemplate) GetModelName() *string {
	return s.ModelName
}

func (s *ModelTemplate) GetProvider() *string {
	return s.Provider
}

func (s *ModelTemplate) SetCollections(v string) *ModelTemplate {
	s.Collections = &v
	return s
}

func (s *ModelTemplate) SetModelName(v string) *ModelTemplate {
	s.ModelName = &v
	return s
}

func (s *ModelTemplate) SetProvider(v string) *ModelTemplate {
	s.Provider = &v
	return s
}

func (s *ModelTemplate) Validate() error {
	return dara.Validate(s)
}
