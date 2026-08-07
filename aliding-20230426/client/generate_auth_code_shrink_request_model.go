// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAuthCodeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucAppName(v string) *GenerateAuthCodeShrinkRequest
	GetBucAppName() *string
	SetSsoTicket(v string) *GenerateAuthCodeShrinkRequest
	GetSsoTicket() *string
	SetTenantContextShrink(v string) *GenerateAuthCodeShrinkRequest
	GetTenantContextShrink() *string
	SetValidRedirectUri(v string) *GenerateAuthCodeShrinkRequest
	GetValidRedirectUri() *string
}

type GenerateAuthCodeShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ali-qwenwork
	BucAppName *string `json:"BucAppName,omitempty" xml:"BucAppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// placeholder-sso-ticket
	SsoTicket           *string `json:"SsoTicket,omitempty" xml:"SsoTicket,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/oauth/callback
	ValidRedirectUri *string `json:"ValidRedirectUri,omitempty" xml:"ValidRedirectUri,omitempty"`
}

func (s GenerateAuthCodeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAuthCodeShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateAuthCodeShrinkRequest) GetBucAppName() *string {
	return s.BucAppName
}

func (s *GenerateAuthCodeShrinkRequest) GetSsoTicket() *string {
	return s.SsoTicket
}

func (s *GenerateAuthCodeShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GenerateAuthCodeShrinkRequest) GetValidRedirectUri() *string {
	return s.ValidRedirectUri
}

func (s *GenerateAuthCodeShrinkRequest) SetBucAppName(v string) *GenerateAuthCodeShrinkRequest {
	s.BucAppName = &v
	return s
}

func (s *GenerateAuthCodeShrinkRequest) SetSsoTicket(v string) *GenerateAuthCodeShrinkRequest {
	s.SsoTicket = &v
	return s
}

func (s *GenerateAuthCodeShrinkRequest) SetTenantContextShrink(v string) *GenerateAuthCodeShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GenerateAuthCodeShrinkRequest) SetValidRedirectUri(v string) *GenerateAuthCodeShrinkRequest {
	s.ValidRedirectUri = &v
	return s
}

func (s *GenerateAuthCodeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
