// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserHonorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *QueryUserHonorsRequestTenantContext) *QueryUserHonorsRequest
	GetTenantContext() *QueryUserHonorsRequestTenantContext
	SetMaxResults(v int32) *QueryUserHonorsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryUserHonorsRequest
	GetNextToken() *string
	SetOrgId(v int64) *QueryUserHonorsRequest
	GetOrgId() *int64
	SetUserId(v string) *QueryUserHonorsRequest
	GetUserId() *string
}

type QueryUserHonorsRequest struct {
	TenantContext *QueryUserHonorsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s QueryUserHonorsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserHonorsRequest) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsRequest) GetTenantContext() *QueryUserHonorsRequestTenantContext {
	return s.TenantContext
}

func (s *QueryUserHonorsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryUserHonorsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryUserHonorsRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *QueryUserHonorsRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryUserHonorsRequest) SetTenantContext(v *QueryUserHonorsRequestTenantContext) *QueryUserHonorsRequest {
	s.TenantContext = v
	return s
}

func (s *QueryUserHonorsRequest) SetMaxResults(v int32) *QueryUserHonorsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryUserHonorsRequest) SetNextToken(v string) *QueryUserHonorsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryUserHonorsRequest) SetOrgId(v int64) *QueryUserHonorsRequest {
	s.OrgId = &v
	return s
}

func (s *QueryUserHonorsRequest) SetUserId(v string) *QueryUserHonorsRequest {
	s.UserId = &v
	return s
}

func (s *QueryUserHonorsRequest) Validate() error {
	return dara.Validate(s)
}

type QueryUserHonorsRequestTenantContext struct {
	// example:
	//
	// 189477710813728
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryUserHonorsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryUserHonorsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryUserHonorsRequestTenantContext) SetTenantId(v string) *QueryUserHonorsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryUserHonorsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
