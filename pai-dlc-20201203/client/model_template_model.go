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
	Collections *string `json:"Collections,omitempty" xml:"Collections,omitempty"`
	ModelName   *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Provider    *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
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
