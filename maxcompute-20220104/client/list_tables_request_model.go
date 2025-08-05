// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListTablesRequest
	GetMarker() *string
	SetMaxItem(v int32) *ListTablesRequest
	GetMaxItem() *int32
	SetPrefix(v string) *ListTablesRequest
	GetPrefix() *string
	SetSchemaName(v string) *ListTablesRequest
	GetSchemaName() *string
	SetType(v string) *ListTablesRequest
	GetType() *string
}

type ListTablesRequest struct {
	// Specifies the marker after which the returned list begins.
	//
	// example:
	//
	// Y29tbWlzc2lvbl9leHRlcm5hbF91cmdlXzFfd3Ih
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The names of the returned resources. The names must start with the value specified by the prefix parameter. If the prefix parameter is set to a, the names of the returned resources must start with a.
	//
	// example:
	//
	// a
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// The type of the table.
	//
	// example:
	//
	// internal
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListTablesRequest) GetMaxItem() *int32 {
	return s.MaxItem
}

func (s *ListTablesRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListTablesRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListTablesRequest) GetType() *string {
	return s.Type
}

func (s *ListTablesRequest) SetMarker(v string) *ListTablesRequest {
	s.Marker = &v
	return s
}

func (s *ListTablesRequest) SetMaxItem(v int32) *ListTablesRequest {
	s.MaxItem = &v
	return s
}

func (s *ListTablesRequest) SetPrefix(v string) *ListTablesRequest {
	s.Prefix = &v
	return s
}

func (s *ListTablesRequest) SetSchemaName(v string) *ListTablesRequest {
	s.SchemaName = &v
	return s
}

func (s *ListTablesRequest) SetType(v string) *ListTablesRequest {
	s.Type = &v
	return s
}

func (s *ListTablesRequest) Validate() error {
	return dara.Validate(s)
}
