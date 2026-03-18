// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSchemaName(v string) *GetTableInfoRequest
	GetSchemaName() *string
	SetType(v string) *GetTableInfoRequest
	GetType() *string
}

type GetTableInfoRequest struct {
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTableInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableInfoRequest) GoString() string {
	return s.String()
}

func (s *GetTableInfoRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetTableInfoRequest) GetType() *string {
	return s.Type
}

func (s *GetTableInfoRequest) SetSchemaName(v string) *GetTableInfoRequest {
	s.SchemaName = &v
	return s
}

func (s *GetTableInfoRequest) SetType(v string) *GetTableInfoRequest {
	s.Type = &v
	return s
}

func (s *GetTableInfoRequest) Validate() error {
	return dara.Validate(s)
}
