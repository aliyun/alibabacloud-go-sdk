// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeClustersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQueryShrink(v string) *ListComputeClustersShrinkRequest
	GetListQueryShrink() *string
	SetMaxResults(v int32) *ListComputeClustersShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListComputeClustersShrinkRequest
	GetNextToken() *string
	SetOpTenantId(v int64) *ListComputeClustersShrinkRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListComputeClustersShrinkRequest
	GetOpUserId() *string
}

type ListComputeClustersShrinkRequest struct {
	// The query conditions.
	//
	// This parameter is required.
	ListQueryShrink *string `json:"ListQuery,omitempty" xml:"ListQuery,omitempty"`
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

func (s ListComputeClustersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeClustersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListComputeClustersShrinkRequest) GetListQueryShrink() *string {
	return s.ListQueryShrink
}

func (s *ListComputeClustersShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComputeClustersShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComputeClustersShrinkRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListComputeClustersShrinkRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListComputeClustersShrinkRequest) SetListQueryShrink(v string) *ListComputeClustersShrinkRequest {
	s.ListQueryShrink = &v
	return s
}

func (s *ListComputeClustersShrinkRequest) SetMaxResults(v int32) *ListComputeClustersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListComputeClustersShrinkRequest) SetNextToken(v string) *ListComputeClustersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListComputeClustersShrinkRequest) SetOpTenantId(v int64) *ListComputeClustersShrinkRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListComputeClustersShrinkRequest) SetOpUserId(v string) *ListComputeClustersShrinkRequest {
	s.OpUserId = &v
	return s
}

func (s *ListComputeClustersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
