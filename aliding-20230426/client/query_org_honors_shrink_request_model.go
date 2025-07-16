// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgHonorsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *QueryOrgHonorsShrinkRequest
	GetTenantContextShrink() *string
	SetMaxResults(v int32) *QueryOrgHonorsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryOrgHonorsShrinkRequest
	GetNextToken() *string
	SetOrgId(v int64) *QueryOrgHonorsShrinkRequest
	GetOrgId() *int64
}

type QueryOrgHonorsShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// 48
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
}

func (s QueryOrgHonorsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgHonorsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryOrgHonorsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryOrgHonorsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryOrgHonorsShrinkRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *QueryOrgHonorsShrinkRequest) SetTenantContextShrink(v string) *QueryOrgHonorsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryOrgHonorsShrinkRequest) SetMaxResults(v int32) *QueryOrgHonorsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryOrgHonorsShrinkRequest) SetNextToken(v string) *QueryOrgHonorsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryOrgHonorsShrinkRequest) SetOrgId(v int64) *QueryOrgHonorsShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *QueryOrgHonorsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
