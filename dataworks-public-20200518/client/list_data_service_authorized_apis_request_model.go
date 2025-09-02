// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceAuthorizedApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiNameKeyword(v string) *ListDataServiceAuthorizedApisRequest
	GetApiNameKeyword() *string
	SetPageNumber(v int32) *ListDataServiceAuthorizedApisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataServiceAuthorizedApisRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataServiceAuthorizedApisRequest
	GetProjectId() *int64
	SetTenantId(v int64) *ListDataServiceAuthorizedApisRequest
	GetTenantId() *int64
}

type ListDataServiceAuthorizedApisRequest struct {
	// The keyword in API names. The keyword is used to search for the APIs whose names contain the keyword.
	//
	// example:
	//
	// My API Name
	ApiNameKeyword *string `json:"ApiNameKeyword,omitempty" xml:"ApiNameKeyword,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspace ID.
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

func (s ListDataServiceAuthorizedApisRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedApisRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedApisRequest) GetApiNameKeyword() *string {
	return s.ApiNameKeyword
}

func (s *ListDataServiceAuthorizedApisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServiceAuthorizedApisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceAuthorizedApisRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServiceAuthorizedApisRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServiceAuthorizedApisRequest) SetApiNameKeyword(v string) *ListDataServiceAuthorizedApisRequest {
	s.ApiNameKeyword = &v
	return s
}

func (s *ListDataServiceAuthorizedApisRequest) SetPageNumber(v int32) *ListDataServiceAuthorizedApisRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataServiceAuthorizedApisRequest) SetPageSize(v int32) *ListDataServiceAuthorizedApisRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceAuthorizedApisRequest) SetProjectId(v int64) *ListDataServiceAuthorizedApisRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceAuthorizedApisRequest) SetTenantId(v int64) *ListDataServiceAuthorizedApisRequest {
	s.TenantId = &v
	return s
}

func (s *ListDataServiceAuthorizedApisRequest) Validate() error {
	return dara.Validate(s)
}
