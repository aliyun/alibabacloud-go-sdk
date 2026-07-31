// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSemanticViewNamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListSemanticViewNamesRequest
	GetDBClusterId() *string
	SetSchemaName(v string) *ListSemanticViewNamesRequest
	GetSchemaName() *string
}

type ListSemanticViewNamesRequest struct {
	// The ADB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp*****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the schema to which the semantic view belongs.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s ListSemanticViewNamesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticViewNamesRequest) GoString() string {
	return s.String()
}

func (s *ListSemanticViewNamesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListSemanticViewNamesRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListSemanticViewNamesRequest) SetDBClusterId(v string) *ListSemanticViewNamesRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListSemanticViewNamesRequest) SetSchemaName(v string) *ListSemanticViewNamesRequest {
	s.SchemaName = &v
	return s
}

func (s *ListSemanticViewNamesRequest) Validate() error {
	return dara.Validate(s)
}
