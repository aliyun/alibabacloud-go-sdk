// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResourceSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSchemaName(v string) *ListInstanceResourceSchemasRequest
	GetSchemaName() *string
}

type ListInstanceResourceSchemasRequest struct {
	// example:
	//
	// default
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s ListInstanceResourceSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResourceSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceResourceSchemasRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListInstanceResourceSchemasRequest) SetSchemaName(v string) *ListInstanceResourceSchemasRequest {
	s.SchemaName = &v
	return s
}

func (s *ListInstanceResourceSchemasRequest) Validate() error {
	return dara.Validate(s)
}
