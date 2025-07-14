// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIngressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListIngressesRequest
	GetAppId() *string
	SetCurrentPage(v int32) *ListIngressesRequest
	GetCurrentPage() *int32
	SetNamespaceId(v string) *ListIngressesRequest
	GetNamespaceId() *string
	SetPageSize(v int32) *ListIngressesRequest
	GetPageSize() *int32
}

type ListIngressesRequest struct {
	// The ID of an application.
	//
	// example:
	//
	// bbf3a590-6d13-46fe-8ca9-c947a20b****
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of a namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIngressesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIngressesRequest) GoString() string {
	return s.String()
}

func (s *ListIngressesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListIngressesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListIngressesRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListIngressesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIngressesRequest) SetAppId(v string) *ListIngressesRequest {
	s.AppId = &v
	return s
}

func (s *ListIngressesRequest) SetCurrentPage(v int32) *ListIngressesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListIngressesRequest) SetNamespaceId(v string) *ListIngressesRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListIngressesRequest) SetPageSize(v int32) *ListIngressesRequest {
	s.PageSize = &v
	return s
}

func (s *ListIngressesRequest) Validate() error {
	return dara.Validate(s)
}
