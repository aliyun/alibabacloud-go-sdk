// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDentryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryId(v string) *QueryDentryShrinkRequest
	GetDentryId() *string
	SetIncludeSpace(v bool) *QueryDentryShrinkRequest
	GetIncludeSpace() *bool
	SetSpaceId(v string) *QueryDentryShrinkRequest
	GetSpaceId() *string
	SetTenantContextShrink(v string) *QueryDentryShrinkRequest
	GetTenantContextShrink() *string
}

type QueryDentryShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx
	DentryId *string `json:"DentryId,omitempty" xml:"DentryId,omitempty"`
	// example:
	//
	// true
	IncludeSpace *bool `json:"IncludeSpace,omitempty" xml:"IncludeSpace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// space-fxhb96vuddz8htqt
	SpaceId             *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryDentryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryDentryShrinkRequest) GetDentryId() *string {
	return s.DentryId
}

func (s *QueryDentryShrinkRequest) GetIncludeSpace() *bool {
	return s.IncludeSpace
}

func (s *QueryDentryShrinkRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *QueryDentryShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryDentryShrinkRequest) SetDentryId(v string) *QueryDentryShrinkRequest {
	s.DentryId = &v
	return s
}

func (s *QueryDentryShrinkRequest) SetIncludeSpace(v bool) *QueryDentryShrinkRequest {
	s.IncludeSpace = &v
	return s
}

func (s *QueryDentryShrinkRequest) SetSpaceId(v string) *QueryDentryShrinkRequest {
	s.SpaceId = &v
	return s
}

func (s *QueryDentryShrinkRequest) SetTenantContextShrink(v string) *QueryDentryShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryDentryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
