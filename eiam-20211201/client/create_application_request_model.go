// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *CreateApplicationRequest
	GetApplicationName() *string
	SetApplicationSourceType(v string) *CreateApplicationRequest
	GetApplicationSourceType() *string
	SetApplicationTemplateId(v string) *CreateApplicationRequest
	GetApplicationTemplateId() *string
	SetDescription(v string) *CreateApplicationRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateApplicationRequest
	GetInstanceId() *string
	SetLogoUrl(v string) *CreateApplicationRequest
	GetLogoUrl() *string
	SetSsoType(v string) *CreateApplicationRequest
	GetSsoType() *string
}

type CreateApplicationRequest struct {
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ram Account SSO
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The type of the application source. Valid values:
	//
	// 	- urn:alibaba:idaas:app:source:template: application template
	//
	// 	- urn:alibaba:idaas:app:source:standard: standard protocol
	//
	// This parameter is required.
	//
	// example:
	//
	// urn:alibaba:idaas:app:source:standard
	ApplicationSourceType *string `json:"ApplicationSourceType,omitempty" xml:"ApplicationSourceType,omitempty"`
	// The ID of the application template. This parameter is required if you set the ApplicationSourceType parameter to urn:alibaba:idaas:app:source:template.
	//
	// example:
	//
	// template_cloud_ram
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// RAM user SSO application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk2676xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The URL of the application logo.
	//
	// example:
	//
	// https://oss.cn-hangzhou.aliyuncs.com/logo.png
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The SSO protocol. Valid values:
	//
	// 	- saml2: the SAML 2.0 protocol.
	//
	// 	- oidc: the OpenID Connect protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// saml2
	SsoType *string `json:"SsoType,omitempty" xml:"SsoType,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *CreateApplicationRequest) GetApplicationSourceType() *string {
	return s.ApplicationSourceType
}

func (s *CreateApplicationRequest) GetApplicationTemplateId() *string {
	return s.ApplicationTemplateId
}

func (s *CreateApplicationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApplicationRequest) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *CreateApplicationRequest) GetSsoType() *string {
	return s.SsoType
}

func (s *CreateApplicationRequest) SetApplicationName(v string) *CreateApplicationRequest {
	s.ApplicationName = &v
	return s
}

func (s *CreateApplicationRequest) SetApplicationSourceType(v string) *CreateApplicationRequest {
	s.ApplicationSourceType = &v
	return s
}

func (s *CreateApplicationRequest) SetApplicationTemplateId(v string) *CreateApplicationRequest {
	s.ApplicationTemplateId = &v
	return s
}

func (s *CreateApplicationRequest) SetDescription(v string) *CreateApplicationRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationRequest) SetInstanceId(v string) *CreateApplicationRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateApplicationRequest) SetLogoUrl(v string) *CreateApplicationRequest {
	s.LogoUrl = &v
	return s
}

func (s *CreateApplicationRequest) SetSsoType(v string) *CreateApplicationRequest {
	s.SsoType = &v
	return s
}

func (s *CreateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
