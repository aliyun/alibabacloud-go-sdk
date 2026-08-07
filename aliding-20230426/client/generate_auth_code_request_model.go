// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAuthCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucAppName(v string) *GenerateAuthCodeRequest
	GetBucAppName() *string
	SetSsoTicket(v string) *GenerateAuthCodeRequest
	GetSsoTicket() *string
	SetTenantContext(v *GenerateAuthCodeRequestTenantContext) *GenerateAuthCodeRequest
	GetTenantContext() *GenerateAuthCodeRequestTenantContext
	SetValidRedirectUri(v string) *GenerateAuthCodeRequest
	GetValidRedirectUri() *string
}

type GenerateAuthCodeRequest struct {
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
	SsoTicket     *string                               `json:"SsoTicket,omitempty" xml:"SsoTicket,omitempty"`
	TenantContext *GenerateAuthCodeRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/oauth/callback
	ValidRedirectUri *string `json:"ValidRedirectUri,omitempty" xml:"ValidRedirectUri,omitempty"`
}

func (s GenerateAuthCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAuthCodeRequest) GoString() string {
	return s.String()
}

func (s *GenerateAuthCodeRequest) GetBucAppName() *string {
	return s.BucAppName
}

func (s *GenerateAuthCodeRequest) GetSsoTicket() *string {
	return s.SsoTicket
}

func (s *GenerateAuthCodeRequest) GetTenantContext() *GenerateAuthCodeRequestTenantContext {
	return s.TenantContext
}

func (s *GenerateAuthCodeRequest) GetValidRedirectUri() *string {
	return s.ValidRedirectUri
}

func (s *GenerateAuthCodeRequest) SetBucAppName(v string) *GenerateAuthCodeRequest {
	s.BucAppName = &v
	return s
}

func (s *GenerateAuthCodeRequest) SetSsoTicket(v string) *GenerateAuthCodeRequest {
	s.SsoTicket = &v
	return s
}

func (s *GenerateAuthCodeRequest) SetTenantContext(v *GenerateAuthCodeRequestTenantContext) *GenerateAuthCodeRequest {
	s.TenantContext = v
	return s
}

func (s *GenerateAuthCodeRequest) SetValidRedirectUri(v string) *GenerateAuthCodeRequest {
	s.ValidRedirectUri = &v
	return s
}

func (s *GenerateAuthCodeRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateAuthCodeRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GenerateAuthCodeRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GenerateAuthCodeRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GenerateAuthCodeRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GenerateAuthCodeRequestTenantContext) SetTenantId(v string) *GenerateAuthCodeRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GenerateAuthCodeRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
