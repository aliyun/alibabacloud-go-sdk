// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTeamsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTeamsRequest
	GetNextToken() *string
	SetTenantContext(v *ListTeamsRequestTenantContext) *ListTeamsRequest
	GetTenantContext() *ListTeamsRequestTenantContext
}

type ListTeamsRequest struct {
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 2023-05-15T11:29Z
	NextToken     *string                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TenantContext *ListTeamsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s ListTeamsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsRequest) GoString() string {
	return s.String()
}

func (s *ListTeamsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTeamsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTeamsRequest) GetTenantContext() *ListTeamsRequestTenantContext {
	return s.TenantContext
}

func (s *ListTeamsRequest) SetMaxResults(v int32) *ListTeamsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTeamsRequest) SetNextToken(v string) *ListTeamsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTeamsRequest) SetTenantContext(v *ListTeamsRequestTenantContext) *ListTeamsRequest {
	s.TenantContext = v
	return s
}

func (s *ListTeamsRequest) Validate() error {
	return dara.Validate(s)
}

type ListTeamsRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListTeamsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ListTeamsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ListTeamsRequestTenantContext) SetTenantId(v string) *ListTeamsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ListTeamsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
