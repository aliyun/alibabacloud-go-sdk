// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListConfigTemplatesRequest
	GetCurrentPage() *int64
	SetId(v int64) *ListConfigTemplatesRequest
	GetId() *int64
	SetName(v string) *ListConfigTemplatesRequest
	GetName() *string
	SetPageSize(v int64) *ListConfigTemplatesRequest
	GetPageSize() *int64
}

type ListConfigTemplatesRequest struct {
	// The number of the page to return. Pages start from Page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the configuration template.
	//
	// example:
	//
	// 3d84efaf-37d9-49fb-a3a8-b38d5c******
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the configuration template.
	//
	// example:
	//
	// config-tmpl-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListConfigTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConfigTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListConfigTemplatesRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListConfigTemplatesRequest) GetId() *int64 {
	return s.Id
}

func (s *ListConfigTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListConfigTemplatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListConfigTemplatesRequest) SetCurrentPage(v int64) *ListConfigTemplatesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListConfigTemplatesRequest) SetId(v int64) *ListConfigTemplatesRequest {
	s.Id = &v
	return s
}

func (s *ListConfigTemplatesRequest) SetName(v string) *ListConfigTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListConfigTemplatesRequest) SetPageSize(v int64) *ListConfigTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListConfigTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
