// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *ListBatchTemplatesRequest
	GetEnv() *string
	SetListQuery(v *ListBatchTemplatesRequestListQuery) *ListBatchTemplatesRequest
	GetListQuery() *ListBatchTemplatesRequestListQuery
	SetOpTenantId(v int64) *ListBatchTemplatesRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *ListBatchTemplatesRequest
	GetProjectId() *int64
}

type ListBatchTemplatesRequest struct {
	// The runtime environment. Default value: PROD.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The paged query conditions.
	ListQuery *ListBatchTemplatesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListBatchTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesRequest) GetEnv() *string {
	return s.Env
}

func (s *ListBatchTemplatesRequest) GetListQuery() *ListBatchTemplatesRequestListQuery {
	return s.ListQuery
}

func (s *ListBatchTemplatesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListBatchTemplatesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListBatchTemplatesRequest) SetEnv(v string) *ListBatchTemplatesRequest {
	s.Env = &v
	return s
}

func (s *ListBatchTemplatesRequest) SetListQuery(v *ListBatchTemplatesRequestListQuery) *ListBatchTemplatesRequest {
	s.ListQuery = v
	return s
}

func (s *ListBatchTemplatesRequest) SetOpTenantId(v int64) *ListBatchTemplatesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListBatchTemplatesRequest) SetProjectId(v int64) *ListBatchTemplatesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListBatchTemplatesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBatchTemplatesRequestListQuery struct {
	// The template name keyword.
	//
	// example:
	//
	// 数据处理模板
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of records per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListBatchTemplatesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListBatchTemplatesRequestListQuery) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListBatchTemplatesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBatchTemplatesRequestListQuery) SetKeyword(v string) *ListBatchTemplatesRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListBatchTemplatesRequestListQuery) SetPageNum(v int32) *ListBatchTemplatesRequestListQuery {
	s.PageNum = &v
	return s
}

func (s *ListBatchTemplatesRequestListQuery) SetPageSize(v int32) *ListBatchTemplatesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListBatchTemplatesRequestListQuery) Validate() error {
	return dara.Validate(s)
}
