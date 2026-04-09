// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDomainsRequest
	GetPageSize() *int32
	SetSearch(v string) *ListDomainsRequest
	GetSearch() *string
	SetWithoutMeteringData(v bool) *ListDomainsRequest
	GetWithoutMeteringData() *bool
}

type ListDomainsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Search              *string `json:"Search,omitempty" xml:"Search,omitempty"`
	WithoutMeteringData *bool   `json:"WithoutMeteringData,omitempty" xml:"WithoutMeteringData,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDomainsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListDomainsRequest) GetWithoutMeteringData() *bool {
	return s.WithoutMeteringData
}

func (s *ListDomainsRequest) SetPageNumber(v int32) *ListDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDomainsRequest) SetPageSize(v int32) *ListDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDomainsRequest) SetSearch(v string) *ListDomainsRequest {
	s.Search = &v
	return s
}

func (s *ListDomainsRequest) SetWithoutMeteringData(v bool) *ListDomainsRequest {
	s.WithoutMeteringData = &v
	return s
}

func (s *ListDomainsRequest) Validate() error {
	return dara.Validate(s)
}
