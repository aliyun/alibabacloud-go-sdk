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
	// The name can be up to 64 characters in length. The name can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The type of the application. Valid values:
	//
	// 	- WebApp: a web application that interacts with a browser.
	//
	// 	- NativeApp: a native application that runs on an operating system, such as a desktop operating system or a mobile operating system.
	//
	// 	- ServerApp: an application that accesses Alibaba Cloud services without the need of manual user logon. User provisioning is automated based on the System for Cross-Domain Identity Management (SCIM) protocol.
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
	// Indicates whether the application can be installed by using other Alibaba Cloud accounts. Valid values:
	//
	// 	- true: If you do not set this parameter for applications of the NativeApp and ServerApp types, true is used.
	//
	// 	- false: If you do not set this parameter for applications of the WebApp type, false is used.
	//
	// example:
	//
	// false
	IsMultiTenant *bool `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty"`
	// The scope of application permissions.
	//
	// For more information about the application permission scope, see [Open authorization scope](https://help.aliyun.com/document_detail/93693.html). You can also call the [ListPredefinedScopes](https://help.aliyun.com/document_detail/187206.html) operation to query the permissions that are supported by different types of applications.
	//
	// If you enter multiple permission scopes, separate them with semicolons (;).
	//
	// example:
	//
	// aliuid
	PredefinedScopes *string `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty"`
	ProtocolVersion  *string `json:"ProtocolVersion,omitempty" xml:"ProtocolVersion,omitempty"`
	// The callback URL.
	//
	// If you enter multiple callback URLs, separate them with semicolons (;).
	//
	// example:
	//
	// https://www.example.com
	RedirectUris *string `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty"`
	// The validity period of the refreshed token.
	//
	// Valid values: 7200 to 31536000. Unit: seconds.
	//
	// Default value:
	//
	// 	- For applications of the WebApp and ServerApp types, if this parameter is left empty, the value 2592000 is used. The value 2592000 indicates that the validity period of the refreshed token is 30 days.
	//
	// 	- For applications of the NativeApp type, if this parameter is left empty, the value 7776000 is used. The value 7776000 indicates that the validity period of the refreshed token is 90 days.
	//
	// example:
	//
	// 2592000
	RefreshTokenValidity *int32 `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// The required permission.
	//
	// You can specify one or more permissions for the `RequiredScopes` parameter. After you specify this parameter, the required permissions are automatically selected and cannot be revoked when a user grants permissions on the application.
	//
	// If you enter multiple permission scopes, separate them with semicolons (;).
	//
	// >  If the permission that you specify for the `RequiredScopes` parameter is not included in the value of the `PredefinedScopes` parameter, the permission does not take effect.
	//
	// example:
	//
	// aliuid;profile
	RequiredScopes *string `json:"RequiredScopes,omitempty" xml:"RequiredScopes,omitempty"`
	// Indicates whether a secret is required. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >- For applications of the WebApp and ServerApp types, this parameter is automatically set to true and cannot be changed.
	//
	// >- For applications of the NativeApp type, this parameter can be set to true or false. If you do not set this parameter, false is used. Applications of the NativeApp type run in untrusted environments and the secrets of these applications are not protected. Therefore, we recommend that you do not set this parameter to true unless otherwise specified. For more information, see [Use an application of the NativeApp type to log on to Alibaba Cloud](https://help.aliyun.com/document_detail/93697.html).
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
