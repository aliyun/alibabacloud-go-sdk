// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppConfigMapping interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *AppConfigMapping
	GetAlias() *string
	SetModelId(v string) *AppConfigMapping
	GetModelId() *string
}

type AppConfigMapping struct {
	// example:
	//
	// list
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// page_eo4r5y6a
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
}

func (s AppConfigMapping) String() string {
	return dara.Prettify(s)
}

func (s AppConfigMapping) GoString() string {
	return s.String()
}

func (s *AppConfigMapping) GetAlias() *string {
	return s.Alias
}

func (s *AppConfigMapping) GetModelId() *string {
	return s.ModelId
}

func (s *AppConfigMapping) SetAlias(v string) *AppConfigMapping {
	s.Alias = &v
	return s
}

func (s *AppConfigMapping) SetModelId(v string) *AppConfigMapping {
	s.ModelId = &v
	return s
}

func (s *AppConfigMapping) Validate() error {
	return dara.Validate(s)
}
