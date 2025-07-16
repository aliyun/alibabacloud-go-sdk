// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDentryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryId(v string) *DeleteDentryShrinkRequest
	GetDentryId() *string
	SetSpaceId(v string) *DeleteDentryShrinkRequest
	GetSpaceId() *string
	SetTenantContextShrink(v string) *DeleteDentryShrinkRequest
	GetTenantContextShrink() *string
	SetToRecycleBin(v bool) *DeleteDentryShrinkRequest
	GetToRecycleBin() *bool
}

type DeleteDentryShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b9XJlZ44W3NeDGyA
	DentryId *string `json:"DentryId,omitempty" xml:"DentryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// nb9XJx4EPx16QGyA
	SpaceId             *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	ToRecycleBin        *bool   `json:"ToRecycleBin,omitempty" xml:"ToRecycleBin,omitempty"`
}

func (s DeleteDentryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDentryShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDentryShrinkRequest) GetDentryId() *string {
	return s.DentryId
}

func (s *DeleteDentryShrinkRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *DeleteDentryShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *DeleteDentryShrinkRequest) GetToRecycleBin() *bool {
	return s.ToRecycleBin
}

func (s *DeleteDentryShrinkRequest) SetDentryId(v string) *DeleteDentryShrinkRequest {
	s.DentryId = &v
	return s
}

func (s *DeleteDentryShrinkRequest) SetSpaceId(v string) *DeleteDentryShrinkRequest {
	s.SpaceId = &v
	return s
}

func (s *DeleteDentryShrinkRequest) SetTenantContextShrink(v string) *DeleteDentryShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *DeleteDentryShrinkRequest) SetToRecycleBin(v bool) *DeleteDentryShrinkRequest {
	s.ToRecycleBin = &v
	return s
}

func (s *DeleteDentryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
