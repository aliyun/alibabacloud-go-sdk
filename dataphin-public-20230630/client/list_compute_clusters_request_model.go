// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListComputeClustersRequestListQuery) *ListComputeClustersRequest
	GetListQuery() *ListComputeClustersRequestListQuery
	SetMaxResults(v int32) *ListComputeClustersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListComputeClustersRequest
	GetNextToken() *string
	SetOpTenantId(v int64) *ListComputeClustersRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListComputeClustersRequest
	GetOpUserId() *string
}

type ListComputeClustersRequest struct {
	// The query conditions.
	//
	// This parameter is required.
	ListQuery *ListComputeClustersRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The maximum number of records to return in this response.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. An empty value indicates that no more results are available.
	//
	// example:
	//
	// fdccfa4f825bf506c591e285f1123403
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListComputeClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeClustersRequest) GoString() string {
	return s.String()
}

func (s *ListComputeClustersRequest) GetListQuery() *ListComputeClustersRequestListQuery {
	return s.ListQuery
}

func (s *ListComputeClustersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComputeClustersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComputeClustersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListComputeClustersRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListComputeClustersRequest) SetListQuery(v *ListComputeClustersRequestListQuery) *ListComputeClustersRequest {
	s.ListQuery = v
	return s
}

func (s *ListComputeClustersRequest) SetMaxResults(v int32) *ListComputeClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListComputeClustersRequest) SetNextToken(v string) *ListComputeClustersRequest {
	s.NextToken = &v
	return s
}

func (s *ListComputeClustersRequest) SetOpTenantId(v int64) *ListComputeClustersRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListComputeClustersRequest) SetOpUserId(v string) *ListComputeClustersRequest {
	s.OpUserId = &v
	return s
}

func (s *ListComputeClustersRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComputeClustersRequestListQuery struct {
	// The keyword for filtering.
	//
	// example:
	//
	// cluster
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. The value must be greater than 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of records per page. The value must be greater than 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of cluster versions.
	TypeVersionList []*string `json:"TypeVersionList,omitempty" xml:"TypeVersionList,omitempty" type:"Repeated"`
}

func (s ListComputeClustersRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListComputeClustersRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListComputeClustersRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListComputeClustersRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListComputeClustersRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListComputeClustersRequestListQuery) GetTypeVersionList() []*string {
	return s.TypeVersionList
}

func (s *ListComputeClustersRequestListQuery) SetKeyword(v string) *ListComputeClustersRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListComputeClustersRequestListQuery) SetPageNo(v int32) *ListComputeClustersRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListComputeClustersRequestListQuery) SetPageSize(v int32) *ListComputeClustersRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListComputeClustersRequestListQuery) SetTypeVersionList(v []*string) *ListComputeClustersRequestListQuery {
	s.TypeVersionList = v
	return s
}

func (s *ListComputeClustersRequestListQuery) Validate() error {
	return dara.Validate(s)
}
