// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDentriesInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIdsForAppPropertiesShrink(v string) *QueryDentriesInfoShrinkRequest
	GetAppIdsForAppPropertiesShrink() *string
	SetDentryId(v string) *QueryDentriesInfoShrinkRequest
	GetDentryId() *string
	SetSpaceId(v string) *QueryDentriesInfoShrinkRequest
	GetSpaceId() *string
	SetTenantContextShrink(v string) *QueryDentriesInfoShrinkRequest
	GetTenantContextShrink() *string
	SetUnionId(v string) *QueryDentriesInfoShrinkRequest
	GetUnionId() *string
	SetWithThumbnail(v bool) *QueryDentriesInfoShrinkRequest
	GetWithThumbnail() *bool
}

type QueryDentriesInfoShrinkRequest struct {
	AppIdsForAppPropertiesShrink *string `json:"AppIdsForAppProperties,omitempty" xml:"AppIdsForAppProperties,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 140901622636
	DentryId *string `json:"DentryId,omitempty" xml:"DentryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 22443475065
	SpaceId             *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// urv3ZIAtcmmIgQzHq08YcAiEiE
	UnionId *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
	// example:
	//
	// false
	WithThumbnail *bool `json:"WithThumbnail,omitempty" xml:"WithThumbnail,omitempty"`
}

func (s QueryDentriesInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDentriesInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryDentriesInfoShrinkRequest) GetAppIdsForAppPropertiesShrink() *string {
	return s.AppIdsForAppPropertiesShrink
}

func (s *QueryDentriesInfoShrinkRequest) GetDentryId() *string {
	return s.DentryId
}

func (s *QueryDentriesInfoShrinkRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *QueryDentriesInfoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryDentriesInfoShrinkRequest) GetUnionId() *string {
	return s.UnionId
}

func (s *QueryDentriesInfoShrinkRequest) GetWithThumbnail() *bool {
	return s.WithThumbnail
}

func (s *QueryDentriesInfoShrinkRequest) SetAppIdsForAppPropertiesShrink(v string) *QueryDentriesInfoShrinkRequest {
	s.AppIdsForAppPropertiesShrink = &v
	return s
}

func (s *QueryDentriesInfoShrinkRequest) SetDentryId(v string) *QueryDentriesInfoShrinkRequest {
	s.DentryId = &v
	return s
}

func (s *QueryDentriesInfoShrinkRequest) SetSpaceId(v string) *QueryDentriesInfoShrinkRequest {
	s.SpaceId = &v
	return s
}

func (s *QueryDentriesInfoShrinkRequest) SetTenantContextShrink(v string) *QueryDentriesInfoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryDentriesInfoShrinkRequest) SetUnionId(v string) *QueryDentriesInfoShrinkRequest {
	s.UnionId = &v
	return s
}

func (s *QueryDentriesInfoShrinkRequest) SetWithThumbnail(v bool) *QueryDentriesInfoShrinkRequest {
	s.WithThumbnail = &v
	return s
}

func (s *QueryDentriesInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
