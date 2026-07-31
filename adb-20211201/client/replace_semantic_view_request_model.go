// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceSemanticViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ReplaceSemanticViewRequest
	GetDBClusterId() *string
	SetDefinition(v string) *ReplaceSemanticViewRequest
	GetDefinition() *string
	SetSchemaName(v string) *ReplaceSemanticViewRequest
	GetSchemaName() *string
	SetViewName(v string) *ReplaceSemanticViewRequest
	GetViewName() *string
}

type ReplaceSemanticViewRequest struct {
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1u8c0mgfg58****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The definition of the semantic view.
	//
	// This parameter is required.
	//
	// example:
	//
	// YAML 内容
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The name of the schema.
	//
	// This parameter is required.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the semantic view.
	//
	// This parameter is required.
	//
	// example:
	//
	// sales_sv
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s ReplaceSemanticViewRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceSemanticViewRequest) GoString() string {
	return s.String()
}

func (s *ReplaceSemanticViewRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ReplaceSemanticViewRequest) GetDefinition() *string {
	return s.Definition
}

func (s *ReplaceSemanticViewRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ReplaceSemanticViewRequest) GetViewName() *string {
	return s.ViewName
}

func (s *ReplaceSemanticViewRequest) SetDBClusterId(v string) *ReplaceSemanticViewRequest {
	s.DBClusterId = &v
	return s
}

func (s *ReplaceSemanticViewRequest) SetDefinition(v string) *ReplaceSemanticViewRequest {
	s.Definition = &v
	return s
}

func (s *ReplaceSemanticViewRequest) SetSchemaName(v string) *ReplaceSemanticViewRequest {
	s.SchemaName = &v
	return s
}

func (s *ReplaceSemanticViewRequest) SetViewName(v string) *ReplaceSemanticViewRequest {
	s.ViewName = &v
	return s
}

func (s *ReplaceSemanticViewRequest) Validate() error {
	return dara.Validate(s)
}
