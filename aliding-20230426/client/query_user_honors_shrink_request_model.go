// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserHonorsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *QueryUserHonorsShrinkRequest
	GetTenantContextShrink() *string
	SetMaxResults(v int32) *QueryUserHonorsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryUserHonorsShrinkRequest
	GetNextToken() *string
	SetOrgId(v int64) *QueryUserHonorsShrinkRequest
	GetOrgId() *int64
	SetUserId(v string) *QueryUserHonorsShrinkRequest
	GetUserId() *string
}

type QueryUserHonorsShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// 200
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
	// This parameter is required.
	//
	// example:
	//
	// 123123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryUserHonorsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserHonorsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryUserHonorsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryUserHonorsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryUserHonorsShrinkRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *QueryUserHonorsShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryUserHonorsShrinkRequest) SetTenantContextShrink(v string) *QueryUserHonorsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryUserHonorsShrinkRequest) SetMaxResults(v int32) *QueryUserHonorsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryUserHonorsShrinkRequest) SetNextToken(v string) *QueryUserHonorsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryUserHonorsShrinkRequest) SetOrgId(v int64) *QueryUserHonorsShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *QueryUserHonorsShrinkRequest) SetUserId(v string) *QueryUserHonorsShrinkRequest {
	s.UserId = &v
	return s
}

func (s *QueryUserHonorsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
