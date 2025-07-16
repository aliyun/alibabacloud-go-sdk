// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocContentTakIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *GetDocContentTakIdShrinkRequest
	GetDentryUuid() *string
	SetGenerateCp(v bool) *GetDocContentTakIdShrinkRequest
	GetGenerateCp() *bool
	SetTargetFormat(v string) *GetDocContentTakIdShrinkRequest
	GetTargetFormat() *string
	SetTenantContextShrink(v string) *GetDocContentTakIdShrinkRequest
	GetTenantContextShrink() *string
}

type GetDocContentTakIdShrinkRequest struct {
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
	TargetFormat        *string `json:"TargetFormat,omitempty" xml:"TargetFormat,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetDocContentTakIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentTakIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDocContentTakIdShrinkRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *GetDocContentTakIdShrinkRequest) GetGenerateCp() *bool {
	return s.GenerateCp
}

func (s *GetDocContentTakIdShrinkRequest) GetTargetFormat() *string {
	return s.TargetFormat
}

func (s *GetDocContentTakIdShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetDocContentTakIdShrinkRequest) SetDentryUuid(v string) *GetDocContentTakIdShrinkRequest {
	s.DentryUuid = &v
	return s
}

func (s *GetDocContentTakIdShrinkRequest) SetGenerateCp(v bool) *GetDocContentTakIdShrinkRequest {
	s.GenerateCp = &v
	return s
}

func (s *GetDocContentTakIdShrinkRequest) SetTargetFormat(v string) *GetDocContentTakIdShrinkRequest {
	s.TargetFormat = &v
	return s
}

func (s *GetDocContentTakIdShrinkRequest) SetTenantContextShrink(v string) *GetDocContentTakIdShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetDocContentTakIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
