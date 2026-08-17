// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListDomainItemsRequest
	GetCurrentPage() *int32
	SetItemValue(v string) *ListDomainItemsRequest
	GetItemValue() *string
	SetListId(v string) *ListDomainItemsRequest
	GetListId() *string
	SetListType(v string) *ListDomainItemsRequest
	GetListType() *string
	SetPageSize(v int32) *ListDomainItemsRequest
	GetPageSize() *int32
}

type ListDomainItemsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// example.com
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ladl-6f1exxxxx6ab59
	ListId *string `json:"ListId,omitempty" xml:"ListId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// la_domain_white_list
	ListType *string `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDomainItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDomainItemsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainItemsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDomainItemsRequest) GetItemValue() *string {
	return s.ItemValue
}

func (s *ListDomainItemsRequest) GetListId() *string {
	return s.ListId
}

func (s *ListDomainItemsRequest) GetListType() *string {
	return s.ListType
}

func (s *ListDomainItemsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDomainItemsRequest) SetCurrentPage(v int32) *ListDomainItemsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDomainItemsRequest) SetItemValue(v string) *ListDomainItemsRequest {
	s.ItemValue = &v
	return s
}

func (s *ListDomainItemsRequest) SetListId(v string) *ListDomainItemsRequest {
	s.ListId = &v
	return s
}

func (s *ListDomainItemsRequest) SetListType(v string) *ListDomainItemsRequest {
	s.ListType = &v
	return s
}

func (s *ListDomainItemsRequest) SetPageSize(v int32) *ListDomainItemsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDomainItemsRequest) Validate() error {
	return dara.Validate(s)
}
