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
	Marker     *string `json:"marker,omitempty" xml:"marker,omitempty"`
	MaxItem    *int32  `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	Prefix     *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
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
