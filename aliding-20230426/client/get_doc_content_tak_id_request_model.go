// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocContentTakIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *GetDocContentTakIdRequest
	GetDentryUuid() *string
	SetGenerateCp(v bool) *GetDocContentTakIdRequest
	GetGenerateCp() *bool
	SetTargetFormat(v string) *GetDocContentTakIdRequest
	GetTargetFormat() *string
	SetTenantContext(v *GetDocContentTakIdRequestTenantContext) *GetDocContentTakIdRequest
	GetTenantContext() *GetDocContentTakIdRequestTenantContext
}

type GetDocContentTakIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// qnYMoO1rWxrkmoj2I5L2PYkoJ47Z3je9
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	GenerateCp *bool   `json:"GenerateCp,omitempty" xml:"GenerateCp,omitempty"`
	// example:
	//
	// markdown
	TargetFormat  *string                                 `json:"TargetFormat,omitempty" xml:"TargetFormat,omitempty"`
	TenantContext *GetDocContentTakIdRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetDocContentTakIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentTakIdRequest) GoString() string {
	return s.String()
}

func (s *GetDocContentTakIdRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *GetDocContentTakIdRequest) GetGenerateCp() *bool {
	return s.GenerateCp
}

func (s *GetDocContentTakIdRequest) GetTargetFormat() *string {
	return s.TargetFormat
}

func (s *GetDocContentTakIdRequest) GetTenantContext() *GetDocContentTakIdRequestTenantContext {
	return s.TenantContext
}

func (s *GetDocContentTakIdRequest) SetDentryUuid(v string) *GetDocContentTakIdRequest {
	s.DentryUuid = &v
	return s
}

func (s *GetDocContentTakIdRequest) SetGenerateCp(v bool) *GetDocContentTakIdRequest {
	s.GenerateCp = &v
	return s
}

func (s *GetDocContentTakIdRequest) SetTargetFormat(v string) *GetDocContentTakIdRequest {
	s.TargetFormat = &v
	return s
}

func (s *GetDocContentTakIdRequest) SetTenantContext(v *GetDocContentTakIdRequestTenantContext) *GetDocContentTakIdRequest {
	s.TenantContext = v
	return s
}

func (s *GetDocContentTakIdRequest) Validate() error {
	return dara.Validate(s)
}

type GetDocContentTakIdRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetDocContentTakIdRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentTakIdRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetDocContentTakIdRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDocContentTakIdRequestTenantContext) SetTenantId(v string) *GetDocContentTakIdRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetDocContentTakIdRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
