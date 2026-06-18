// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCasCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListCasCertificatesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCasCertificatesRequest
	GetPageSize() *int64
	SetSearchKeyword(v string) *ListCasCertificatesRequest
	GetSearchKeyword() *string
	SetSecurityToken(v string) *ListCasCertificatesRequest
	GetSecurityToken() *string
}

type ListCasCertificatesRequest struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// aliyun
	SearchKeyword *string `json:"SearchKeyword,omitempty" xml:"SearchKeyword,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListCasCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCasCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListCasCertificatesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCasCertificatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCasCertificatesRequest) GetSearchKeyword() *string {
	return s.SearchKeyword
}

func (s *ListCasCertificatesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListCasCertificatesRequest) SetPageNumber(v int64) *ListCasCertificatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCasCertificatesRequest) SetPageSize(v int64) *ListCasCertificatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCasCertificatesRequest) SetSearchKeyword(v string) *ListCasCertificatesRequest {
	s.SearchKeyword = &v
	return s
}

func (s *ListCasCertificatesRequest) SetSecurityToken(v string) *ListCasCertificatesRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListCasCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
