// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSslCertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertNameLike(v string) *ListSslCertsRequest
	GetCertNameLike() *string
	SetDomainName(v string) *ListSslCertsRequest
	GetDomainName() *string
	SetPageNumber(v int32) *ListSslCertsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSslCertsRequest
	GetPageSize() *int32
}

type ListSslCertsRequest struct {
	// Name matching keyword.
	//
	// example:
	//
	// ali
	CertNameLike *string `json:"certNameLike,omitempty" xml:"certNameLike,omitempty"`
	// Domain name.
	//
	// example:
	//
	// fun.iot.evideocloud.com.cn
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// Page number, default is 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size, default is 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListSslCertsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSslCertsRequest) GoString() string {
	return s.String()
}

func (s *ListSslCertsRequest) GetCertNameLike() *string {
	return s.CertNameLike
}

func (s *ListSslCertsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ListSslCertsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSslCertsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSslCertsRequest) SetCertNameLike(v string) *ListSslCertsRequest {
	s.CertNameLike = &v
	return s
}

func (s *ListSslCertsRequest) SetDomainName(v string) *ListSslCertsRequest {
	s.DomainName = &v
	return s
}

func (s *ListSslCertsRequest) SetPageNumber(v int32) *ListSslCertsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSslCertsRequest) SetPageSize(v int32) *ListSslCertsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSslCertsRequest) Validate() error {
	return dara.Validate(s)
}
