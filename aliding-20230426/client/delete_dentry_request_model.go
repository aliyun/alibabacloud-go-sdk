// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDentryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryId(v string) *DeleteDentryRequest
	GetDentryId() *string
	SetSpaceId(v string) *DeleteDentryRequest
	GetSpaceId() *string
	SetTenantContext(v *DeleteDentryRequestTenantContext) *DeleteDentryRequest
	GetTenantContext() *DeleteDentryRequestTenantContext
	SetToRecycleBin(v bool) *DeleteDentryRequest
	GetToRecycleBin() *bool
}

type DeleteDentryRequest struct {
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
	SpaceId       *string                           `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContext *DeleteDentryRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	ToRecycleBin  *bool                             `json:"ToRecycleBin,omitempty" xml:"ToRecycleBin,omitempty"`
}

func (s DeleteDentryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDentryRequest) GoString() string {
	return s.String()
}

func (s *DeleteDentryRequest) GetDentryId() *string {
	return s.DentryId
}

func (s *DeleteDentryRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *DeleteDentryRequest) GetTenantContext() *DeleteDentryRequestTenantContext {
	return s.TenantContext
}

func (s *DeleteDentryRequest) GetToRecycleBin() *bool {
	return s.ToRecycleBin
}

func (s *DeleteDentryRequest) SetDentryId(v string) *DeleteDentryRequest {
	s.DentryId = &v
	return s
}

func (s *DeleteDentryRequest) SetSpaceId(v string) *DeleteDentryRequest {
	s.SpaceId = &v
	return s
}

func (s *DeleteDentryRequest) SetTenantContext(v *DeleteDentryRequestTenantContext) *DeleteDentryRequest {
	s.TenantContext = v
	return s
}

func (s *DeleteDentryRequest) SetToRecycleBin(v bool) *DeleteDentryRequest {
	s.ToRecycleBin = &v
	return s
}

func (s *DeleteDentryRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDentryRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteDentryRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteDentryRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeleteDentryRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteDentryRequestTenantContext) SetTenantId(v string) *DeleteDentryRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeleteDentryRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
