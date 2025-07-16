// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDentryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryId(v string) *CopyDentryShrinkRequest
	GetDentryId() *string
	SetName(v string) *CopyDentryShrinkRequest
	GetName() *string
	SetSpaceId(v string) *CopyDentryShrinkRequest
	GetSpaceId() *string
	SetTargetSpaceId(v string) *CopyDentryShrinkRequest
	GetTargetSpaceId() *string
	SetTenantContextShrink(v string) *CopyDentryShrinkRequest
	GetTenantContextShrink() *string
	SetToNextDentryId(v string) *CopyDentryShrinkRequest
	GetToNextDentryId() *string
	SetToParentDentryId(v string) *CopyDentryShrinkRequest
	GetToParentDentryId() *string
	SetToPrevDentryId(v string) *CopyDentryShrinkRequest
	GetToPrevDentryId() *string
}

type CopyDentryShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b9XJljElJv6RPGyA
	DentryId *string `json:"DentryId,omitempty" xml:"DentryId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 24458420428
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b9XJljElJv6RPG
	TargetSpaceId       *string `json:"TargetSpaceId,omitempty" xml:"TargetSpaceId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// b9XJljElJv6RPGyA2
	ToNextDentryId *string `json:"ToNextDentryId,omitempty" xml:"ToNextDentryId,omitempty"`
	// example:
	//
	// b9XJljElJv6RPGyA4
	ToParentDentryId *string `json:"ToParentDentryId,omitempty" xml:"ToParentDentryId,omitempty"`
	// example:
	//
	// b9XJljElJv6RPGyA3
	ToPrevDentryId *string `json:"ToPrevDentryId,omitempty" xml:"ToPrevDentryId,omitempty"`
}

func (s CopyDentryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CopyDentryShrinkRequest) GetDentryId() *string {
	return s.DentryId
}

func (s *CopyDentryShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CopyDentryShrinkRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *CopyDentryShrinkRequest) GetTargetSpaceId() *string {
	return s.TargetSpaceId
}

func (s *CopyDentryShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CopyDentryShrinkRequest) GetToNextDentryId() *string {
	return s.ToNextDentryId
}

func (s *CopyDentryShrinkRequest) GetToParentDentryId() *string {
	return s.ToParentDentryId
}

func (s *CopyDentryShrinkRequest) GetToPrevDentryId() *string {
	return s.ToPrevDentryId
}

func (s *CopyDentryShrinkRequest) SetDentryId(v string) *CopyDentryShrinkRequest {
	s.DentryId = &v
	return s
}

func (s *CopyDentryShrinkRequest) SetName(v string) *CopyDentryShrinkRequest {
	s.Name = &v
	return s
}

func (s *CopyDentryShrinkRequest) SetSpaceId(v string) *CopyDentryShrinkRequest {
	s.SpaceId = &v
	return s
}

func (s *CopyDentryShrinkRequest) SetTargetSpaceId(v string) *CopyDentryShrinkRequest {
	s.TargetSpaceId = &v
	return s
}

func (s *CopyDentryShrinkRequest) SetTenantContextShrink(v string) *CopyDentryShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CopyDentryShrinkRequest) SetToNextDentryId(v string) *CopyDentryShrinkRequest {
	s.ToNextDentryId = &v
	return s
}

func (s *CopyDentryShrinkRequest) SetToParentDentryId(v string) *CopyDentryShrinkRequest {
	s.ToParentDentryId = &v
	return s
}

func (s *CopyDentryShrinkRequest) SetToPrevDentryId(v string) *CopyDentryShrinkRequest {
	s.ToPrevDentryId = &v
	return s
}

func (s *CopyDentryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
