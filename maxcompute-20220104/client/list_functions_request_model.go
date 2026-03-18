// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListFunctionsRequest
	GetMarker() *string
	SetMaxItem(v int32) *ListFunctionsRequest
	GetMaxItem() *int32
	SetPrefix(v string) *ListFunctionsRequest
	GetPrefix() *string
	SetSchemaName(v string) *ListFunctionsRequest
	GetSchemaName() *string
}

type ListFunctionsRequest struct {
	Marker     *string `json:"marker,omitempty" xml:"marker,omitempty"`
	MaxItem    *int32  `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	Prefix     *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s ListFunctionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListFunctionsRequest) GetMaxItem() *int32 {
	return s.MaxItem
}

func (s *ListFunctionsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListFunctionsRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListFunctionsRequest) SetMarker(v string) *ListFunctionsRequest {
	s.Marker = &v
	return s
}

func (s *ListFunctionsRequest) SetMaxItem(v int32) *ListFunctionsRequest {
	s.MaxItem = &v
	return s
}

func (s *ListFunctionsRequest) SetPrefix(v string) *ListFunctionsRequest {
	s.Prefix = &v
	return s
}

func (s *ListFunctionsRequest) SetSchemaName(v string) *ListFunctionsRequest {
	s.SchemaName = &v
	return s
}

func (s *ListFunctionsRequest) Validate() error {
	return dara.Validate(s)
}
