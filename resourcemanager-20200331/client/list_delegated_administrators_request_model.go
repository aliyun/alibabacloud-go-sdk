// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDelegatedAdministratorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListDelegatedAdministratorsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDelegatedAdministratorsRequest
	GetPageSize() *int64
	SetServicePrincipal(v string) *ListDelegatedAdministratorsRequest
	GetServicePrincipal() *string
}

type ListDelegatedAdministratorsRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The identifier of the trusted service.
	//
	// For more information, see the `Trusted service identifier` column in [Supported trusted services](https://help.aliyun.com/document_detail/208133.html).
	//
	// example:
	//
	// cloudfw.aliyuncs.com
	ServicePrincipal *string `json:"ServicePrincipal,omitempty" xml:"ServicePrincipal,omitempty"`
}

func (s ListDelegatedAdministratorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDelegatedAdministratorsRequest) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDelegatedAdministratorsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDelegatedAdministratorsRequest) GetServicePrincipal() *string {
	return s.ServicePrincipal
}

func (s *ListDelegatedAdministratorsRequest) SetPageNumber(v int64) *ListDelegatedAdministratorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDelegatedAdministratorsRequest) SetPageSize(v int64) *ListDelegatedAdministratorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDelegatedAdministratorsRequest) SetServicePrincipal(v string) *ListDelegatedAdministratorsRequest {
	s.ServicePrincipal = &v
	return s
}

func (s *ListDelegatedAdministratorsRequest) Validate() error {
	return dara.Validate(s)
}
