// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDcdnKvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *ListDcdnKvRequest
	GetNamespace() *string
	SetPageNumber(v int32) *ListDcdnKvRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDcdnKvRequest
	GetPageSize() *int32
	SetPrefix(v string) *ListDcdnKvRequest
	GetPrefix() *string
}

type ListDcdnKvRequest struct {
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The number of the page to return. The product of PageNumber and PageSize cannot exceed 50,000.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 50. Maximum value: 100.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The prefix to query.
	//
	// example:
	//
	// prefix-
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListDcdnKvRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDcdnKvRequest) GoString() string {
	return s.String()
}

func (s *ListDcdnKvRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListDcdnKvRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDcdnKvRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDcdnKvRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListDcdnKvRequest) SetNamespace(v string) *ListDcdnKvRequest {
	s.Namespace = &v
	return s
}

func (s *ListDcdnKvRequest) SetPageNumber(v int32) *ListDcdnKvRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDcdnKvRequest) SetPageSize(v int32) *ListDcdnKvRequest {
	s.PageSize = &v
	return s
}

func (s *ListDcdnKvRequest) SetPrefix(v string) *ListDcdnKvRequest {
	s.Prefix = &v
	return s
}

func (s *ListDcdnKvRequest) Validate() error {
	return dara.Validate(s)
}
