// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgHonorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *QueryOrgHonorsRequestTenantContext) *QueryOrgHonorsRequest
	GetTenantContext() *QueryOrgHonorsRequestTenantContext
	SetMaxResults(v int32) *QueryOrgHonorsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryOrgHonorsRequest
	GetNextToken() *string
	SetOrgId(v int64) *QueryOrgHonorsRequest
	GetOrgId() *int64
}

type QueryOrgHonorsRequest struct {
	TenantContext *QueryOrgHonorsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s QueryOrgHonorsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgHonorsRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsRequest) GetTenantContext() *QueryOrgHonorsRequestTenantContext {
	return s.TenantContext
}

func (s *QueryOrgHonorsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryOrgHonorsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryOrgHonorsRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *QueryOrgHonorsRequest) SetTenantContext(v *QueryOrgHonorsRequestTenantContext) *QueryOrgHonorsRequest {
	s.TenantContext = v
	return s
}

func (s *QueryOrgHonorsRequest) SetMaxResults(v int32) *QueryOrgHonorsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryOrgHonorsRequest) SetNextToken(v string) *QueryOrgHonorsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryOrgHonorsRequest) SetOrgId(v int64) *QueryOrgHonorsRequest {
	s.OrgId = &v
	return s
}

func (s *QueryOrgHonorsRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryOrgHonorsRequestTenantContext struct {
	// example:
	//
	// 487986704507650
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryOrgHonorsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgHonorsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryOrgHonorsRequestTenantContext) SetTenantId(v string) *QueryOrgHonorsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryOrgHonorsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
