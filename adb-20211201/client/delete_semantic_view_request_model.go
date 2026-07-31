// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSemanticViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteSemanticViewRequest
	GetDBClusterId() *string
	SetSchemaName(v string) *DeleteSemanticViewRequest
	GetSchemaName() *string
	SetViewName(v string) *DeleteSemanticViewRequest
	GetViewName() *string
}

type DeleteSemanticViewRequest struct {
	// The ID of the ADB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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
	// adb_sv_name
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s DeleteSemanticViewRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSemanticViewRequest) GoString() string {
	return s.String()
}

func (s *DeleteSemanticViewRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteSemanticViewRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DeleteSemanticViewRequest) GetViewName() *string {
	return s.ViewName
}

func (s *DeleteSemanticViewRequest) SetDBClusterId(v string) *DeleteSemanticViewRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteSemanticViewRequest) SetSchemaName(v string) *DeleteSemanticViewRequest {
	s.SchemaName = &v
	return s
}

func (s *DeleteSemanticViewRequest) SetViewName(v string) *DeleteSemanticViewRequest {
	s.ViewName = &v
	return s
}

func (s *DeleteSemanticViewRequest) Validate() error {
	return dara.Validate(s)
}
