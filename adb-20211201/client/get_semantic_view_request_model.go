// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSemanticViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetSemanticViewRequest
	GetDBClusterId() *string
	SetSchemaName(v string) *GetSemanticViewRequest
	GetSchemaName() *string
	SetViewName(v string) *GetSemanticViewRequest
	GetViewName() *string
}

type GetSemanticViewRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1d8lbdj22rx****
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
	// sv_name
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
}

func (s GetSemanticViewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSemanticViewRequest) GoString() string {
	return s.String()
}

func (s *GetSemanticViewRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSemanticViewRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetSemanticViewRequest) GetViewName() *string {
	return s.ViewName
}

func (s *GetSemanticViewRequest) SetDBClusterId(v string) *GetSemanticViewRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetSemanticViewRequest) SetSchemaName(v string) *GetSemanticViewRequest {
	s.SchemaName = &v
	return s
}

func (s *GetSemanticViewRequest) SetViewName(v string) *GetSemanticViewRequest {
	s.ViewName = &v
	return s
}

func (s *GetSemanticViewRequest) Validate() error {
	return dara.Validate(s)
}
