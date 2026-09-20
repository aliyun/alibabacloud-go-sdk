// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServicePublishedApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiNameKeyword(v string) *ListDataServicePublishedApisRequest
	GetApiNameKeyword() *string
	SetApiPathKeyword(v string) *ListDataServicePublishedApisRequest
	GetApiPathKeyword() *string
	SetCreatorId(v string) *ListDataServicePublishedApisRequest
	GetCreatorId() *string
	SetPageNumber(v int32) *ListDataServicePublishedApisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataServicePublishedApisRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataServicePublishedApisRequest
	GetProjectId() *int64
	SetTenantId(v int64) *ListDataServicePublishedApisRequest
	GetTenantId() *int64
}

type ListDataServicePublishedApisRequest struct {
	// The keyword used to filter APIs by name. APIs whose names contain this keyword are returned.
	//
	// example:
	//
	// My API name
	ApiNameKeyword *string `json:"ApiNameKeyword,omitempty" xml:"ApiNameKeyword,omitempty"`
	// The keyword used to filter APIs by path. APIs whose paths contain this keyword are returned.
	//
	// example:
	//
	// /test/
	ApiPathKeyword *string `json:"ApiPathKeyword,omitempty" xml:"ApiPathKeyword,omitempty"`
	// The Alibaba Cloud ID of the API creator. Only APIs created by this user are returned.
	//
	// example:
	//
	// 12345
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
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
	// The tenant ID. To obtain the tenant ID, log on to the [DataWorks console](https://workbench.data.aliyun.com/console), go to the DataStudio page, click your username in the upper-right corner, and select Menu > User Info.
	//
	// example:
	//
	// 10001
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDataServicePublishedApisRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisRequest) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisRequest) GetApiNameKeyword() *string {
	return s.ApiNameKeyword
}

func (s *ListDataServicePublishedApisRequest) GetApiPathKeyword() *string {
	return s.ApiPathKeyword
}

func (s *ListDataServicePublishedApisRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDataServicePublishedApisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServicePublishedApisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServicePublishedApisRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServicePublishedApisRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServicePublishedApisRequest) SetApiNameKeyword(v string) *ListDataServicePublishedApisRequest {
	s.ApiNameKeyword = &v
	return s
}

func (s *ListDataServicePublishedApisRequest) SetApiPathKeyword(v string) *ListDataServicePublishedApisRequest {
	s.ApiPathKeyword = &v
	return s
}

func (s *ListDataServicePublishedApisRequest) SetCreatorId(v string) *ListDataServicePublishedApisRequest {
	s.CreatorId = &v
	return s
}

func (s *ListDataServicePublishedApisRequest) SetPageNumber(v int32) *ListDataServicePublishedApisRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataServicePublishedApisRequest) SetPageSize(v int32) *ListDataServicePublishedApisRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataServicePublishedApisRequest) SetProjectId(v int64) *ListDataServicePublishedApisRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServicePublishedApisRequest) SetTenantId(v int64) *ListDataServicePublishedApisRequest {
	s.TenantId = &v
	return s
}

func (s *ListDataServicePublishedApisRequest) Validate() error {
	return dara.Validate(s)
}
