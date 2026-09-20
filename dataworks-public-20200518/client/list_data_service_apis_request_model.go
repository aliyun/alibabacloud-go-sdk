// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiNameKeyword(v string) *ListDataServiceApisRequest
	GetApiNameKeyword() *string
	SetApiPathKeyword(v string) *ListDataServiceApisRequest
	GetApiPathKeyword() *string
	SetCreatorId(v string) *ListDataServiceApisRequest
	GetCreatorId() *string
	SetPageNumber(v int32) *ListDataServiceApisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataServiceApisRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataServiceApisRequest
	GetProjectId() *int64
	SetTenantId(v int64) *ListDataServiceApisRequest
	GetTenantId() *int64
}

type ListDataServiceApisRequest struct {
	// The keyword in API names. The keyword is used to search for the APIs whose names contain the keyword.
	//
	// example:
	//
	// My API name
	ApiNameKeyword *string `json:"ApiNameKeyword,omitempty" xml:"ApiNameKeyword,omitempty"`
	// The keyword in API paths. The keyword is used to search for the APIs whose paths contain the keyword.
	//
	// example:
	//
	// /test/
	ApiPathKeyword *string `json:"ApiPathKeyword,omitempty" xml:"ApiPathKeyword,omitempty"`
	// The ID of the Alibaba Cloud account used by the creator of the APIs. The ID is used to search for the APIs created by the creator.
	//
	// example:
	//
	// 12345
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// example:
	//
	// 10001
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDataServiceApisRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisRequest) GetApiNameKeyword() *string {
	return s.ApiNameKeyword
}

func (s *ListDataServiceApisRequest) GetApiPathKeyword() *string {
	return s.ApiPathKeyword
}

func (s *ListDataServiceApisRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDataServiceApisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServiceApisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceApisRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServiceApisRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServiceApisRequest) SetApiNameKeyword(v string) *ListDataServiceApisRequest {
	s.ApiNameKeyword = &v
	return s
}

func (s *ListDataServiceApisRequest) SetApiPathKeyword(v string) *ListDataServiceApisRequest {
	s.ApiPathKeyword = &v
	return s
}

func (s *ListDataServiceApisRequest) SetCreatorId(v string) *ListDataServiceApisRequest {
	s.CreatorId = &v
	return s
}

func (s *ListDataServiceApisRequest) SetPageNumber(v int32) *ListDataServiceApisRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataServiceApisRequest) SetPageSize(v int32) *ListDataServiceApisRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceApisRequest) SetProjectId(v int64) *ListDataServiceApisRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApisRequest) SetTenantId(v int64) *ListDataServiceApisRequest {
	s.TenantId = &v
	return s
}

func (s *ListDataServiceApisRequest) Validate() error {
	return dara.Validate(s)
}
