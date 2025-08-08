// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelProviderSpec interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorization(v *ModelProviderAuthorization) *ModelProviderSpec
	GetAuthorization() *ModelProviderAuthorization
	SetSchema(v *ModelProviderSchema) *ModelProviderSpec
	GetSchema() *ModelProviderSchema
}

type ModelProviderSpec struct {
	Authorization *ModelProviderAuthorization `json:"authorization,omitempty" xml:"authorization,omitempty"`
	Schema        *ModelProviderSchema        `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s ModelProviderSpec) String() string {
	return dara.Prettify(s)
}

func (s ModelProviderSpec) GoString() string {
	return s.String()
}

func (s *ModelProviderSpec) GetAuthorization() *ModelProviderAuthorization {
	return s.Authorization
}

func (s *ModelProviderSpec) GetSchema() *ModelProviderSchema {
	return s.Schema
}

func (s *ModelProviderSpec) SetAuthorization(v *ModelProviderAuthorization) *ModelProviderSpec {
	s.Authorization = v
	return s
}

func (s *ModelProviderSpec) SetSchema(v *ModelProviderSchema) *ModelProviderSpec {
	s.Schema = v
	return s
}

func (s *ModelProviderSpec) Validate() error {
	return dara.Validate(s)
}
