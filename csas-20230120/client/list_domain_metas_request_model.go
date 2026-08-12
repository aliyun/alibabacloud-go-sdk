// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainMetasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListDomainMetasRequest
	GetCurrentPage() *int32
	SetDefaultTemplate(v bool) *ListDomainMetasRequest
	GetDefaultTemplate() *bool
	SetListType(v string) *ListDomainMetasRequest
	GetListType() *string
	SetName(v string) *ListDomainMetasRequest
	GetName() *string
	SetPageSize(v int32) *ListDomainMetasRequest
	GetPageSize() *int32
}

type ListDomainMetasRequest struct {
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether to include system default template lists.
	//
	// example:
	//
	// false
	DefaultTemplate *bool `json:"DefaultTemplate,omitempty" xml:"DefaultTemplate,omitempty"`
	// The list type (blacklist/whitelist).
	//
	// This parameter is required.
	//
	// example:
	//
	// la_domain_white_list
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// The list name. Fuzzy match is supported.
	//
	// example:
	//
	// OfficeDomains
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page in a paged query. Settings: 1 to 1000. Paging is used to return results.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDomainMetasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDomainMetasRequest) GoString() string {
	return s.String()
}

func (s *ListDomainMetasRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDomainMetasRequest) GetDefaultTemplate() *bool {
	return s.DefaultTemplate
}

func (s *ListDomainMetasRequest) GetListType() *string {
	return s.ListType
}

func (s *ListDomainMetasRequest) GetName() *string {
	return s.Name
}

func (s *ListDomainMetasRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDomainMetasRequest) SetCurrentPage(v int32) *ListDomainMetasRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDomainMetasRequest) SetDefaultTemplate(v bool) *ListDomainMetasRequest {
	s.DefaultTemplate = &v
	return s
}

func (s *ListDomainMetasRequest) SetListType(v string) *ListDomainMetasRequest {
	s.ListType = &v
	return s
}

func (s *ListDomainMetasRequest) SetName(v string) *ListDomainMetasRequest {
	s.Name = &v
	return s
}

func (s *ListDomainMetasRequest) SetPageSize(v int32) *ListDomainMetasRequest {
	s.PageSize = &v
	return s
}

func (s *ListDomainMetasRequest) Validate() error {
	return dara.Validate(s)
}
