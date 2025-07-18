// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdpConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListIdpConfigsRequest
	GetCurrentPage() *int64
	SetInclude(v string) *ListIdpConfigsRequest
	GetInclude() *string
	SetPageSize(v int64) *ListIdpConfigsRequest
	GetPageSize() *int64
}

type ListIdpConfigsRequest struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// CSAS,DingTalk,LDAP
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIdpConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIdpConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListIdpConfigsRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListIdpConfigsRequest) GetInclude() *string {
	return s.Include
}

func (s *ListIdpConfigsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListIdpConfigsRequest) SetCurrentPage(v int64) *ListIdpConfigsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListIdpConfigsRequest) SetInclude(v string) *ListIdpConfigsRequest {
	s.Include = &v
	return s
}

func (s *ListIdpConfigsRequest) SetPageSize(v int64) *ListIdpConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIdpConfigsRequest) Validate() error {
	return dara.Validate(s)
}
