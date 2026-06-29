// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTemplatesRequest
	GetPageSize() *int32
	SetSearchKey(v string) *ListTemplatesRequest
	GetSearchKey() *string
	SetTypes(v []*string) *ListTemplatesRequest
	GetTypes() []*string
}

type ListTemplatesRequest struct {
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Search content
	//
	// example:
	//
	// demo
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// List of application types.
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplatesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *ListTemplatesRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListTemplatesRequest) SetPageNumber(v int32) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int32) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetSearchKey(v string) *ListTemplatesRequest {
	s.SearchKey = &v
	return s
}

func (s *ListTemplatesRequest) SetTypes(v []*string) *ListTemplatesRequest {
	s.Types = v
	return s
}

func (s *ListTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
