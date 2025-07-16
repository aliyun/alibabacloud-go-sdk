// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocContentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *GetDocContentShrinkRequest
	GetDentryUuid() *string
	SetTargetFormat(v string) *GetDocContentShrinkRequest
	GetTargetFormat() *string
	SetTenantContextShrink(v string) *GetDocContentShrinkRequest
	GetTenantContextShrink() *string
	SetUserToken(v string) *GetDocContentShrinkRequest
	GetUserToken() *string
}

type GetDocContentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dentry_uuid
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// example:
	//
	// markdown
	TargetFormat        *string `json:"TargetFormat,omitempty" xml:"TargetFormat,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	UserToken *string `json:"userToken,omitempty" xml:"userToken,omitempty"`
}

func (s GetDocContentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDocContentShrinkRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *GetDocContentShrinkRequest) GetTargetFormat() *string {
	return s.TargetFormat
}

func (s *GetDocContentShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetDocContentShrinkRequest) GetUserToken() *string {
	return s.UserToken
}

func (s *GetDocContentShrinkRequest) SetDentryUuid(v string) *GetDocContentShrinkRequest {
	s.DentryUuid = &v
	return s
}

func (s *GetDocContentShrinkRequest) SetTargetFormat(v string) *GetDocContentShrinkRequest {
	s.TargetFormat = &v
	return s
}

func (s *GetDocContentShrinkRequest) SetTenantContextShrink(v string) *GetDocContentShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetDocContentShrinkRequest) SetUserToken(v string) *GetDocContentShrinkRequest {
	s.UserToken = &v
	return s
}

func (s *GetDocContentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
