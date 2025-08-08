// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToolsetSpec interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v *Authorization) *ToolsetSpec
	GetAuthConfig() *Authorization
	SetSchema(v *ToolsetSchema) *ToolsetSpec
	GetSchema() *ToolsetSchema
}

type ToolsetSpec struct {
	AuthConfig *Authorization `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	Schema     *ToolsetSchema `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s ToolsetSpec) String() string {
	return dara.Prettify(s)
}

func (s ToolsetSpec) GoString() string {
	return s.String()
}

func (s *ToolsetSpec) GetAuthConfig() *Authorization {
	return s.AuthConfig
}

func (s *ToolsetSpec) GetSchema() *ToolsetSchema {
	return s.Schema
}

func (s *ToolsetSpec) SetAuthConfig(v *Authorization) *ToolsetSpec {
	s.AuthConfig = v
	return s
}

func (s *ToolsetSpec) SetSchema(v *ToolsetSchema) *ToolsetSpec {
	s.Schema = v
	return s
}

func (s *ToolsetSpec) Validate() error {
	return dara.Validate(s)
}
