// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdpDepartmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListIdpDepartmentsRequest
	GetCurrentPage() *int64
	SetIdpConfigId(v string) *ListIdpDepartmentsRequest
	GetIdpConfigId() *string
	SetPageSize(v int64) *ListIdpDepartmentsRequest
	GetPageSize() *int64
}

type ListIdpDepartmentsRequest struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1440
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIdpDepartmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIdpDepartmentsRequest) GoString() string {
	return s.String()
}

func (s *ListIdpDepartmentsRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListIdpDepartmentsRequest) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *ListIdpDepartmentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListIdpDepartmentsRequest) SetCurrentPage(v int64) *ListIdpDepartmentsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListIdpDepartmentsRequest) SetIdpConfigId(v string) *ListIdpDepartmentsRequest {
	s.IdpConfigId = &v
	return s
}

func (s *ListIdpDepartmentsRequest) SetPageSize(v int64) *ListIdpDepartmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIdpDepartmentsRequest) Validate() error {
	return dara.Validate(s)
}
