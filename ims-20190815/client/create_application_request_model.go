// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTokenValidity(v int32) *CreateApplicationRequest
	GetAccessTokenValidity() *int32
	SetAppName(v string) *CreateApplicationRequest
	GetAppName() *string
	SetAppType(v string) *CreateApplicationRequest
	GetAppType() *string
	SetDisplayName(v string) *CreateApplicationRequest
	GetDisplayName() *string
	SetIsMultiTenant(v bool) *CreateApplicationRequest
	GetIsMultiTenant() *bool
	SetPredefinedScopes(v string) *CreateApplicationRequest
	GetPredefinedScopes() *string
	SetProtocolVersion(v string) *CreateApplicationRequest
	GetProtocolVersion() *string
	SetRedirectUris(v string) *CreateApplicationRequest
	GetRedirectUris() *string
	SetRefreshTokenValidity(v int32) *CreateApplicationRequest
	GetRefreshTokenValidity() *int32
	SetRequiredScopes(v string) *CreateApplicationRequest
	GetRequiredScopes() *string
	SetSecretRequired(v bool) *CreateApplicationRequest
	GetSecretRequired() *bool
}

type CreateApplicationRequest struct {
	// The validity period of the access token.
	//
	// Valid values: 900 to 10800. Unit: seconds.
	//
	// Default value: 3600.
	//
	// example:
	//
	// 3600
	AccessTokenValidity *int32 `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// The application name.
	//
	// The name can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The type of the application. Valid values:
	//
	// - WebApp: a web application that is based on browser interaction.
	//
	// - NativeApp: a native application that runs on an operating system, such as a desktop or mobile operating system.
	//
	// - ServerApp: an application that directly accesses Alibaba Cloud services without user logon. Currently, only applications that use the System for Cross-domain Identity Management (SCIM) protocol for user synchronization are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// WebApp
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The display name of the application.
	//
	// The name can be up to 24 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// myapp
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Specifies whether the application can be installed by other Alibaba Cloud accounts. Valid values:
	//
	// - true: For NativeApp and ServerApp applications, the default value is \\`true\\` if you leave this parameter empty.
	//
	// - false: For WebApp applications, the default value is \\`false\\` if you leave this parameter empty.
	//
	// example:
	//
	// false
	IsMultiTenant *bool `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty"`
	// The scopes of the application.
	//
	// For information about the valid values and descriptions of scopes, see [OAuth scopes](https://help.aliyun.com/document_detail/93693.html). You can also call the [ListPredefinedScopes](https://help.aliyun.com/document_detail/187206.html) operation to obtain the scopes that are supported by different application types.
	//
	// To enter multiple scopes, separate them with semicolons (;).
	//
	// example:
	//
	// aliuid;profile
	PredefinedScopes *string `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty"`
	// The OAuth protocol version of the application. Valid values:
	//
	// - `2.0`: OAuth 2.0.
	//
	// - `2.1`: OAuth 2.1.
	//
	// Default value: `2.0`.
	//
	// example:
	//
	// 2.0
	ProtocolVersion *string `json:"ProtocolVersion,omitempty" xml:"ProtocolVersion,omitempty"`
	// The webhook address.
	//
	// To enter multiple webhook addresses, separate them with semicolons (;).
	//
	// example:
	//
	// https://www.example.com
	RedirectUris *string `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty"`
	// The validity period of the refresh token.
	//
	// Valid values: 7200 to 31536000. Unit: seconds.
	//
	// Default value:
	//
	// - For NativeApp and ServerApp applications, the default value is 2,592,000 seconds (30 days) if you leave this parameter empty.
	//
	// - For WebApp applications, the default value is 7,776,000 seconds (90 days) if you leave this parameter empty.
	//
	// example:
	//
	// 2592000
	RefreshTokenValidity *int32 `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// The required scopes.
	//
	// You can specify one or more scopes in `RequiredScopes` as required. When a user grants permissions to the application, the required scopes are selected by default and cannot be deselected.
	//
	// To enter multiple scopes, separate them with semicolons (;).
	//
	// > If a scope that you specify in `RequiredScopes` is not within the range of `PredefinedScopes`, the required setting for that scope does not take effect.
	//
	// example:
	//
	// aliuid
	RequiredScopes *string `json:"RequiredScopes,omitempty" xml:"RequiredScopes,omitempty"`
	// Specifies whether an application key is required. Valid values:
	//
	// - true
	//
	// - false
	//
	// > 	- For WebApp and ServerApp applications, this parameter is forcibly set to \\`true\\` and cannot be changed.
	//
	// - For NativeApp applications, you can set this parameter to \\`true\\` or \\`false\\`. If you do not set this parameter, the default value is \\`false\\`. Because these applications often run in untrusted environments where application keys cannot be effectively protected, do not set this parameter to \\`true\\` unless necessary. For more information, see [Log on to Alibaba Cloud from a native application](https://help.aliyun.com/document_detail/93697.html).
	//
	// example:
	//
	// true
	SecretRequired *bool `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) GetAccessTokenValidity() *int32 {
	return s.AccessTokenValidity
}

func (s *CreateApplicationRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateApplicationRequest) GetAppType() *string {
	return s.AppType
}

func (s *CreateApplicationRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateApplicationRequest) GetIsMultiTenant() *bool {
	return s.IsMultiTenant
}

func (s *CreateApplicationRequest) GetPredefinedScopes() *string {
	return s.PredefinedScopes
}

func (s *CreateApplicationRequest) GetProtocolVersion() *string {
	return s.ProtocolVersion
}

func (s *CreateApplicationRequest) GetRedirectUris() *string {
	return s.RedirectUris
}

func (s *CreateApplicationRequest) GetRefreshTokenValidity() *int32 {
	return s.RefreshTokenValidity
}

func (s *CreateApplicationRequest) GetRequiredScopes() *string {
	return s.RequiredScopes
}

func (s *CreateApplicationRequest) GetSecretRequired() *bool {
	return s.SecretRequired
}

func (s *CreateApplicationRequest) SetAccessTokenValidity(v int32) *CreateApplicationRequest {
	s.AccessTokenValidity = &v
	return s
}

func (s *CreateApplicationRequest) SetAppName(v string) *CreateApplicationRequest {
	s.AppName = &v
	return s
}

func (s *CreateApplicationRequest) SetAppType(v string) *CreateApplicationRequest {
	s.AppType = &v
	return s
}

func (s *CreateApplicationRequest) SetDisplayName(v string) *CreateApplicationRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateApplicationRequest) SetIsMultiTenant(v bool) *CreateApplicationRequest {
	s.IsMultiTenant = &v
	return s
}

func (s *CreateApplicationRequest) SetPredefinedScopes(v string) *CreateApplicationRequest {
	s.PredefinedScopes = &v
	return s
}

func (s *CreateApplicationRequest) SetProtocolVersion(v string) *CreateApplicationRequest {
	s.ProtocolVersion = &v
	return s
}

func (s *CreateApplicationRequest) SetRedirectUris(v string) *CreateApplicationRequest {
	s.RedirectUris = &v
	return s
}

func (s *CreateApplicationRequest) SetRefreshTokenValidity(v int32) *CreateApplicationRequest {
	s.RefreshTokenValidity = &v
	return s
}

func (s *CreateApplicationRequest) SetRequiredScopes(v string) *CreateApplicationRequest {
	s.RequiredScopes = &v
	return s
}

func (s *CreateApplicationRequest) SetSecretRequired(v bool) *CreateApplicationRequest {
	s.SecretRequired = &v
	return s
}

func (s *CreateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
