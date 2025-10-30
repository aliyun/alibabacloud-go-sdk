// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListProjectsRequestListQuery) *ListProjectsRequest
	GetListQuery() *ListProjectsRequestListQuery
	SetOpTenantId(v int64) *ListProjectsRequest
	GetOpTenantId() *int64
}

type ListProjectsRequest struct {
	// This parameter is required.
	ListQuery *ListProjectsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetListQuery() *ListProjectsRequestListQuery {
	return s.ListQuery
}

func (s *ListProjectsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListProjectsRequest) SetListQuery(v *ListProjectsRequestListQuery) *ListProjectsRequest {
	s.ListQuery = v
	return s
}

func (s *ListProjectsRequest) SetOpTenantId(v int64) *ListProjectsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsRequestListQuery struct {
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// BASIC
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TagList  []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
}

func (s ListProjectsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListProjectsRequestListQuery) GetEnv() *string {
	return s.Env
}

func (s *ListProjectsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProjectsRequestListQuery) GetMode() *string {
	return s.Mode
}

func (s *ListProjectsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListProjectsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectsRequestListQuery) GetTagList() []*string {
	return s.TagList
}

func (s *ListProjectsRequestListQuery) SetEnv(v string) *ListProjectsRequestListQuery {
	s.Env = &v
	return s
}

func (s *ListProjectsRequestListQuery) SetKeyword(v string) *ListProjectsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListProjectsRequestListQuery) SetMode(v string) *ListProjectsRequestListQuery {
	s.Mode = &v
	return s
}

func (s *ListProjectsRequestListQuery) SetPageNo(v int32) *ListProjectsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListProjectsRequestListQuery) SetPageSize(v int32) *ListProjectsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListProjectsRequestListQuery) SetTagList(v []*string) *ListProjectsRequestListQuery {
	s.TagList = v
	return s
}

func (s *ListProjectsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
