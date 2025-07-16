// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpaceDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryId(v string) *GetSpaceDirectoriesRequest
	GetDentryId() *string
	SetMaxResults(v int32) *GetSpaceDirectoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetSpaceDirectoriesRequest
	GetNextToken() *string
	SetSpaceId(v string) *GetSpaceDirectoriesRequest
	GetSpaceId() *string
	SetTenantContext(v *GetSpaceDirectoriesRequestTenantContext) *GetSpaceDirectoriesRequest
	GetTenantContext() *GetSpaceDirectoriesRequestTenantContext
}

type GetSpaceDirectoriesRequest struct {
	// example:
	//
	// asdasd
	DentryId *string `json:"DentryId,omitempty" xml:"DentryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 123123
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qweqwe
	SpaceId       *string                                  `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContext *GetSpaceDirectoriesRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetSpaceDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesRequest) GetDentryId() *string {
	return s.DentryId
}

func (s *GetSpaceDirectoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetSpaceDirectoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetSpaceDirectoriesRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *GetSpaceDirectoriesRequest) GetTenantContext() *GetSpaceDirectoriesRequestTenantContext {
	return s.TenantContext
}

func (s *GetSpaceDirectoriesRequest) SetDentryId(v string) *GetSpaceDirectoriesRequest {
	s.DentryId = &v
	return s
}

func (s *GetSpaceDirectoriesRequest) SetMaxResults(v int32) *GetSpaceDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *GetSpaceDirectoriesRequest) SetNextToken(v string) *GetSpaceDirectoriesRequest {
	s.NextToken = &v
	return s
}

func (s *GetSpaceDirectoriesRequest) SetSpaceId(v string) *GetSpaceDirectoriesRequest {
	s.SpaceId = &v
	return s
}

func (s *GetSpaceDirectoriesRequest) SetTenantContext(v *GetSpaceDirectoriesRequestTenantContext) *GetSpaceDirectoriesRequest {
	s.TenantContext = v
	return s
}

func (s *GetSpaceDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}

type GetSpaceDirectoriesRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetSpaceDirectoriesRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSpaceDirectoriesRequestTenantContext) SetTenantId(v string) *GetSpaceDirectoriesRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetSpaceDirectoriesRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
