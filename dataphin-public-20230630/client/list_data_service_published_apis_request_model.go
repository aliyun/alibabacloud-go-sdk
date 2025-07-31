// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServicePublishedApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListDataServicePublishedApisRequestListQuery) *ListDataServicePublishedApisRequest
	GetListQuery() *ListDataServicePublishedApisRequestListQuery
	SetOpTenantId(v int64) *ListDataServicePublishedApisRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServicePublishedApisRequest
	GetProjectId() *int32
}

type ListDataServicePublishedApisRequest struct {
	ListQuery *ListDataServicePublishedApisRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDataServicePublishedApisRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisRequest) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisRequest) GetListQuery() *ListDataServicePublishedApisRequestListQuery {
	return s.ListQuery
}

func (s *ListDataServicePublishedApisRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServicePublishedApisRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServicePublishedApisRequest) SetListQuery(v *ListDataServicePublishedApisRequestListQuery) *ListDataServicePublishedApisRequest {
	s.ListQuery = v
	return s
}

func (s *ListDataServicePublishedApisRequest) SetOpTenantId(v int64) *ListDataServicePublishedApisRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServicePublishedApisRequest) SetProjectId(v int32) *ListDataServicePublishedApisRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServicePublishedApisRequest) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisRequestListQuery struct {
	// example:
	//
	// test
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// example:
	//
	// 102113
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDataServicePublishedApisRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisRequestListQuery) GetApiName() *string {
	return s.ApiName
}

func (s *ListDataServicePublishedApisRequestListQuery) GetGroupId() *int32 {
	return s.GroupId
}

func (s *ListDataServicePublishedApisRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListDataServicePublishedApisRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServicePublishedApisRequestListQuery) SetApiName(v string) *ListDataServicePublishedApisRequestListQuery {
	s.ApiName = &v
	return s
}

func (s *ListDataServicePublishedApisRequestListQuery) SetGroupId(v int32) *ListDataServicePublishedApisRequestListQuery {
	s.GroupId = &v
	return s
}

func (s *ListDataServicePublishedApisRequestListQuery) SetPageNo(v int32) *ListDataServicePublishedApisRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListDataServicePublishedApisRequestListQuery) SetPageSize(v int32) *ListDataServicePublishedApisRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListDataServicePublishedApisRequestListQuery) Validate() error {
	return dara.Validate(s)
}
