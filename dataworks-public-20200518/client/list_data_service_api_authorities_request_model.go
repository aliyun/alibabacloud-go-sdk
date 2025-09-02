// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiAuthoritiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiNameKeyword(v string) *ListDataServiceApiAuthoritiesRequest
	GetApiNameKeyword() *string
	SetPageNumber(v int32) *ListDataServiceApiAuthoritiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataServiceApiAuthoritiesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataServiceApiAuthoritiesRequest
	GetProjectId() *int64
	SetTenantId(v int64) *ListDataServiceApiAuthoritiesRequest
	GetTenantId() *int64
}

type ListDataServiceApiAuthoritiesRequest struct {
	// The keyword in the name of the API. The keyword can be used to search for the API whose name contains the keyword.
	//
	// example:
	//
	// My API name
	ApiNameKeyword *string `json:"ApiNameKeyword,omitempty" xml:"ApiNameKeyword,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID. This parameter is deprecated.
	//
	// example:
	//
	// 10001
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDataServiceApiAuthoritiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiAuthoritiesRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiAuthoritiesRequest) GetApiNameKeyword() *string {
	return s.ApiNameKeyword
}

func (s *ListDataServiceApiAuthoritiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServiceApiAuthoritiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceApiAuthoritiesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServiceApiAuthoritiesRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServiceApiAuthoritiesRequest) SetApiNameKeyword(v string) *ListDataServiceApiAuthoritiesRequest {
	s.ApiNameKeyword = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesRequest) SetPageNumber(v int32) *ListDataServiceApiAuthoritiesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesRequest) SetPageSize(v int32) *ListDataServiceApiAuthoritiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesRequest) SetProjectId(v int64) *ListDataServiceApiAuthoritiesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesRequest) SetTenantId(v int64) *ListDataServiceApiAuthoritiesRequest {
	s.TenantId = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesRequest) Validate() error {
	return dara.Validate(s)
}
