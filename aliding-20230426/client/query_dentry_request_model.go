// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDentryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryId(v string) *QueryDentryRequest
	GetDentryId() *string
	SetIncludeSpace(v bool) *QueryDentryRequest
	GetIncludeSpace() *bool
	SetSpaceId(v string) *QueryDentryRequest
	GetSpaceId() *string
	SetTenantContext(v *QueryDentryRequestTenantContext) *QueryDentryRequest
	GetTenantContext() *QueryDentryRequestTenantContext
}

type QueryDentryRequest struct {
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
	SpaceId       *string                          `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContext *QueryDentryRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s QueryDentryRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryRequest) GoString() string {
	return s.String()
}

func (s *QueryDentryRequest) GetDentryId() *string {
	return s.DentryId
}

func (s *QueryDentryRequest) GetIncludeSpace() *bool {
	return s.IncludeSpace
}

func (s *QueryDentryRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *QueryDentryRequest) GetTenantContext() *QueryDentryRequestTenantContext {
	return s.TenantContext
}

func (s *QueryDentryRequest) SetDentryId(v string) *QueryDentryRequest {
	s.DentryId = &v
	return s
}

func (s *QueryDentryRequest) SetIncludeSpace(v bool) *QueryDentryRequest {
	s.IncludeSpace = &v
	return s
}

func (s *QueryDentryRequest) SetSpaceId(v string) *QueryDentryRequest {
	s.SpaceId = &v
	return s
}

func (s *QueryDentryRequest) SetTenantContext(v *QueryDentryRequestTenantContext) *QueryDentryRequest {
	s.TenantContext = v
	return s
}

func (s *QueryDentryRequest) Validate() error {
	return dara.Validate(s)
}

type QueryDentryRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryDentryRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryDentryRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryDentryRequestTenantContext) SetTenantId(v string) *QueryDentryRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryDentryRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
