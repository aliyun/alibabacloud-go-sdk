// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListDataServiceApplicationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataServiceApplicationsRequest
	GetPageSize() *int32
	SetProjectIdList(v string) *ListDataServiceApplicationsRequest
	GetProjectIdList() *string
	SetTenantId(v int64) *ListDataServiceApplicationsRequest
	GetTenantId() *int64
}

type ListDataServiceApplicationsRequest struct {
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
	// The ID of the workspace based on which you want to query the basic information of applications. You can specify multiple IDs. Separate them with commas (,). You must specify at least one workspace ID. You can specify a maximum of 50 workspace IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000,100001
	ProjectIdList *string `json:"ProjectIdList,omitempty" xml:"ProjectIdList,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// example:
	//
	// 100002
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDataServiceApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceApplicationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServiceApplicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceApplicationsRequest) GetProjectIdList() *string {
	return s.ProjectIdList
}

func (s *ListDataServiceApplicationsRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServiceApplicationsRequest) SetPageNumber(v int32) *ListDataServiceApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataServiceApplicationsRequest) SetPageSize(v int32) *ListDataServiceApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceApplicationsRequest) SetProjectIdList(v string) *ListDataServiceApplicationsRequest {
	s.ProjectIdList = &v
	return s
}

func (s *ListDataServiceApplicationsRequest) SetTenantId(v int64) *ListDataServiceApplicationsRequest {
	s.TenantId = &v
	return s
}

func (s *ListDataServiceApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
