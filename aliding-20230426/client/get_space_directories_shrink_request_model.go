// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpaceDirectoriesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryId(v string) *GetSpaceDirectoriesShrinkRequest
	GetDentryId() *string
	SetMaxResults(v int32) *GetSpaceDirectoriesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetSpaceDirectoriesShrinkRequest
	GetNextToken() *string
	SetSpaceId(v string) *GetSpaceDirectoriesShrinkRequest
	GetSpaceId() *string
	SetTenantContextShrink(v string) *GetSpaceDirectoriesShrinkRequest
	GetTenantContextShrink() *string
}

type GetSpaceDirectoriesShrinkRequest struct {
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
	SpaceId             *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetSpaceDirectoriesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesShrinkRequest) GetDentryId() *string {
	return s.DentryId
}

func (s *GetSpaceDirectoriesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetSpaceDirectoriesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetSpaceDirectoriesShrinkRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *GetSpaceDirectoriesShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetSpaceDirectoriesShrinkRequest) SetDentryId(v string) *GetSpaceDirectoriesShrinkRequest {
	s.DentryId = &v
	return s
}

func (s *GetSpaceDirectoriesShrinkRequest) SetMaxResults(v int32) *GetSpaceDirectoriesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *GetSpaceDirectoriesShrinkRequest) SetNextToken(v string) *GetSpaceDirectoriesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *GetSpaceDirectoriesShrinkRequest) SetSpaceId(v string) *GetSpaceDirectoriesShrinkRequest {
	s.SpaceId = &v
	return s
}

func (s *GetSpaceDirectoriesShrinkRequest) SetTenantContextShrink(v string) *GetSpaceDirectoriesShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetSpaceDirectoriesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
