// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListComputeSourcesRequestListQuery) *ListComputeSourcesRequest
	GetListQuery() *ListComputeSourcesRequestListQuery
	SetOpTenantId(v int64) *ListComputeSourcesRequest
	GetOpTenantId() *int64
}

type ListComputeSourcesRequest struct {
	// This parameter is required.
	ListQuery *ListComputeSourcesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListComputeSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListComputeSourcesRequest) GetListQuery() *ListComputeSourcesRequestListQuery {
	return s.ListQuery
}

func (s *ListComputeSourcesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListComputeSourcesRequest) SetListQuery(v *ListComputeSourcesRequestListQuery) *ListComputeSourcesRequest {
	s.ListQuery = v
	return s
}

func (s *ListComputeSourcesRequest) SetOpTenantId(v int64) *ListComputeSourcesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListComputeSourcesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComputeSourcesRequestListQuery struct {
	BindProject *bool `json:"BindProject,omitempty" xml:"BindProject,omitempty"`
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListComputeSourcesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListComputeSourcesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListComputeSourcesRequestListQuery) GetBindProject() *bool {
	return s.BindProject
}

func (s *ListComputeSourcesRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListComputeSourcesRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListComputeSourcesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeSourcesRequestListQuery) GetType() *string {
	return s.Type
}

func (s *ListComputeSourcesRequestListQuery) SetBindProject(v bool) *ListComputeSourcesRequestListQuery {
	s.BindProject = &v
	return s
}

func (s *ListComputeSourcesRequestListQuery) SetKeyword(v string) *ListComputeSourcesRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListComputeSourcesRequestListQuery) SetPageNo(v int32) *ListComputeSourcesRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListComputeSourcesRequestListQuery) SetPageSize(v int32) *ListComputeSourcesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListComputeSourcesRequestListQuery) SetType(v string) *ListComputeSourcesRequestListQuery {
	s.Type = &v
	return s
}

func (s *ListComputeSourcesRequestListQuery) Validate() error {
	return dara.Validate(s)
}
