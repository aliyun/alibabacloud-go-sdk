// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSemanticViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateSemanticViewRequest
	GetDBClusterId() *string
	SetDefinition(v string) *CreateSemanticViewRequest
	GetDefinition() *string
	SetSchemaName(v string) *CreateSemanticViewRequest
	GetSchemaName() *string
	SetViewName(v string) *CreateSemanticViewRequest
	GetViewName() *string
}

type CreateSemanticViewRequest struct {
	// ADB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp198m028ih55****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// YAML definition of the semantic view.
	//
	// This parameter is required.
	//
	// example:
	//
	// YAML 内容
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// Schema name.
	//
	// This parameter is required.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// Semantic view name.
	//
	// This parameter is required.
	//
	// example:
	//
	// adb_sv_name
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s CreateSemanticViewRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSemanticViewRequest) GoString() string {
	return s.String()
}

func (s *CreateSemanticViewRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateSemanticViewRequest) GetDefinition() *string {
	return s.Definition
}

func (s *CreateSemanticViewRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *CreateSemanticViewRequest) GetViewName() *string {
	return s.ViewName
}

func (s *CreateSemanticViewRequest) SetDBClusterId(v string) *CreateSemanticViewRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateSemanticViewRequest) SetDefinition(v string) *CreateSemanticViewRequest {
	s.Definition = &v
	return s
}

func (s *CreateSemanticViewRequest) SetSchemaName(v string) *CreateSemanticViewRequest {
	s.SchemaName = &v
	return s
}

func (s *CreateSemanticViewRequest) SetViewName(v string) *CreateSemanticViewRequest {
	s.ViewName = &v
	return s
}

func (s *CreateSemanticViewRequest) Validate() error {
	return dara.Validate(s)
}
