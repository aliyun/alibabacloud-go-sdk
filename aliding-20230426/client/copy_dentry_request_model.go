// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDentryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryId(v string) *CopyDentryRequest
	GetDentryId() *string
	SetName(v string) *CopyDentryRequest
	GetName() *string
	SetSpaceId(v string) *CopyDentryRequest
	GetSpaceId() *string
	SetTargetSpaceId(v string) *CopyDentryRequest
	GetTargetSpaceId() *string
	SetTenantContext(v *CopyDentryRequestTenantContext) *CopyDentryRequest
	GetTenantContext() *CopyDentryRequestTenantContext
	SetToNextDentryId(v string) *CopyDentryRequest
	GetToNextDentryId() *string
	SetToParentDentryId(v string) *CopyDentryRequest
	GetToParentDentryId() *string
	SetToPrevDentryId(v string) *CopyDentryRequest
	GetToPrevDentryId() *string
}

type CopyDentryRequest struct {
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
	TargetSpaceId *string                         `json:"TargetSpaceId,omitempty" xml:"TargetSpaceId,omitempty"`
	TenantContext *CopyDentryRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s CopyDentryRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryRequest) GoString() string {
	return s.String()
}

func (s *CopyDentryRequest) GetDentryId() *string {
	return s.DentryId
}

func (s *CopyDentryRequest) GetName() *string {
	return s.Name
}

func (s *CopyDentryRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *CopyDentryRequest) GetTargetSpaceId() *string {
	return s.TargetSpaceId
}

func (s *CopyDentryRequest) GetTenantContext() *CopyDentryRequestTenantContext {
	return s.TenantContext
}

func (s *CopyDentryRequest) GetToNextDentryId() *string {
	return s.ToNextDentryId
}

func (s *CopyDentryRequest) GetToParentDentryId() *string {
	return s.ToParentDentryId
}

func (s *CopyDentryRequest) GetToPrevDentryId() *string {
	return s.ToPrevDentryId
}

func (s *CopyDentryRequest) SetDentryId(v string) *CopyDentryRequest {
	s.DentryId = &v
	return s
}

func (s *CopyDentryRequest) SetName(v string) *CopyDentryRequest {
	s.Name = &v
	return s
}

func (s *CopyDentryRequest) SetSpaceId(v string) *CopyDentryRequest {
	s.SpaceId = &v
	return s
}

func (s *CopyDentryRequest) SetTargetSpaceId(v string) *CopyDentryRequest {
	s.TargetSpaceId = &v
	return s
}

func (s *CopyDentryRequest) SetTenantContext(v *CopyDentryRequestTenantContext) *CopyDentryRequest {
	s.TenantContext = v
	return s
}

func (s *CopyDentryRequest) SetToNextDentryId(v string) *CopyDentryRequest {
	s.ToNextDentryId = &v
	return s
}

func (s *CopyDentryRequest) SetToParentDentryId(v string) *CopyDentryRequest {
	s.ToParentDentryId = &v
	return s
}

func (s *CopyDentryRequest) SetToPrevDentryId(v string) *CopyDentryRequest {
	s.ToPrevDentryId = &v
	return s
}

func (s *CopyDentryRequest) Validate() error {
	return dara.Validate(s)
}

type CopyDentryRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CopyDentryRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CopyDentryRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CopyDentryRequestTenantContext) SetTenantId(v string) *CopyDentryRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CopyDentryRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
