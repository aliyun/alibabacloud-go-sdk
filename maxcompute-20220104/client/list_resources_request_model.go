// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMarker(v string) *ListResourcesRequest
	GetMarker() *string
	SetMaxItem(v int32) *ListResourcesRequest
	GetMaxItem() *int32
	SetName(v string) *ListResourcesRequest
	GetName() *string
	SetSchemaName(v string) *ListResourcesRequest
	GetSchemaName() *string
}

type ListResourcesRequest struct {
	// Specifies the marker after which the returned list begins.
	//
	// example:
	//
	// cHlvZHBzX3VkZl8xMDExNV8xNDU3NDI4NDkzKg==
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// res
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// default
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListResourcesRequest) GetMaxItem() *int32 {
	return s.MaxItem
}

func (s *ListResourcesRequest) GetName() *string {
	return s.Name
}

func (s *ListResourcesRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListResourcesRequest) SetMarker(v string) *ListResourcesRequest {
	s.Marker = &v
	return s
}

func (s *ListResourcesRequest) SetMaxItem(v int32) *ListResourcesRequest {
	s.MaxItem = &v
	return s
}

func (s *ListResourcesRequest) SetName(v string) *ListResourcesRequest {
	s.Name = &v
	return s
}

func (s *ListResourcesRequest) SetSchemaName(v string) *ListResourcesRequest {
	s.SchemaName = &v
	return s
}

func (s *ListResourcesRequest) Validate() error {
	return dara.Validate(s)
}
