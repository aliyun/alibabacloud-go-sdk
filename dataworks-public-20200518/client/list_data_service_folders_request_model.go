// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceFoldersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderNameKeyword(v string) *ListDataServiceFoldersRequest
	GetFolderNameKeyword() *string
	SetGroupId(v string) *ListDataServiceFoldersRequest
	GetGroupId() *string
	SetPageNumber(v int32) *ListDataServiceFoldersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataServiceFoldersRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataServiceFoldersRequest
	GetProjectId() *int64
	SetTenantId(v int64) *ListDataServiceFoldersRequest
	GetTenantId() *int64
}

type ListDataServiceFoldersRequest struct {
	// The keyword in folder names. The keyword is used to search for folders whose names contain this keyword.
	//
	// example:
	//
	// Keyword in folder names
	FolderNameKeyword *string `json:"FolderNameKeyword,omitempty" xml:"FolderNameKeyword,omitempty"`
	// The ID of the business process to which the folders belong.
	//
	// example:
	//
	// ds_123abc
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the page that appears, click the username for the logon in the upper-right corner and click User Info in the Menu section.
	//
	// example:
	//
	// 10002
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDataServiceFoldersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceFoldersRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceFoldersRequest) GetFolderNameKeyword() *string {
	return s.FolderNameKeyword
}

func (s *ListDataServiceFoldersRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListDataServiceFoldersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServiceFoldersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceFoldersRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServiceFoldersRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServiceFoldersRequest) SetFolderNameKeyword(v string) *ListDataServiceFoldersRequest {
	s.FolderNameKeyword = &v
	return s
}

func (s *ListDataServiceFoldersRequest) SetGroupId(v string) *ListDataServiceFoldersRequest {
	s.GroupId = &v
	return s
}

func (s *ListDataServiceFoldersRequest) SetPageNumber(v int32) *ListDataServiceFoldersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataServiceFoldersRequest) SetPageSize(v int32) *ListDataServiceFoldersRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceFoldersRequest) SetProjectId(v int64) *ListDataServiceFoldersRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceFoldersRequest) SetTenantId(v int64) *ListDataServiceFoldersRequest {
	s.TenantId = &v
	return s
}

func (s *ListDataServiceFoldersRequest) Validate() error {
	return dara.Validate(s)
}
