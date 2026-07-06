// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateApplicationRequest
	GetAppId() *string
	SetNewAccessTokenValidity(v int32) *UpdateApplicationRequest
	GetNewAccessTokenValidity() *int32
	SetNewDisplayName(v string) *UpdateApplicationRequest
	GetNewDisplayName() *string
	SetNewIsMultiTenant(v bool) *UpdateApplicationRequest
	GetNewIsMultiTenant() *bool
	SetNewPredefinedScopes(v string) *UpdateApplicationRequest
	GetNewPredefinedScopes() *string
	SetNewRedirectUris(v string) *UpdateApplicationRequest
	GetNewRedirectUris() *string
	SetNewRefreshTokenValidity(v int32) *UpdateApplicationRequest
	GetNewRefreshTokenValidity() *int32
	SetNewRequiredScopes(v string) *UpdateApplicationRequest
	GetNewRequiredScopes() *string
	SetNewSecretRequired(v bool) *UpdateApplicationRequest
	GetNewSecretRequired() *bool
}

type UpdateApplicationRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The validity period of the access token.
	//
	// Valid values: 900 to 10800. Unit: seconds.
	//
	// example:
	//
	// 3600
	NewAccessTokenValidity *int32 `json:"NewAccessTokenValidity,omitempty" xml:"NewAccessTokenValidity,omitempty"`
	// The display name of the application.
	//
	// example:
	//
	// NewApp
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	// Specifies whether the application can be installed by using other Alibaba Cloud accounts. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	NewIsMultiTenant *bool `json:"NewIsMultiTenant,omitempty" xml:"NewIsMultiTenant,omitempty"`
	// The scope of application permissions.
	//
	// For more information about the application permission scope, see [OAuth overview](https://help.aliyun.com/document_detail/93693.html). You can also call the [ListPredefinedScopes](https://help.aliyun.com/document_detail/187206.html) operation to obtain the scopes that are supported by different application types.
	//
	// To specify multiple permissions, separate them with semicolons (;).
	//
	// The new value of this parameter overwrites the original value, and the permission specified by the new value takes effect. For example, if the original value is `/acs/ccc`, and the new value is `/acs/alidns`, `/acs/alidns` takes effect. If you want to retain the original permission and the `/acs/alidns` permission, set the value to `/acs/ccc;/acs/alidns`.
	//
	// example:
	//
	// openid
	NewPredefinedScopes *string `json:"NewPredefinedScopes,omitempty" xml:"NewPredefinedScopes,omitempty"`
	// The redirect URL.
	//
	// To specify multiple URLs, separate them with semicolons (;).
	//
	// example:
	//
	// https://www.example.com
	NewRedirectUris *string `json:"NewRedirectUris,omitempty" xml:"NewRedirectUris,omitempty"`
	// The validity period of the refresh token.
	//
	// Valid values: 7200 to 31536000. Unit: seconds.
	//
	// example:
	//
	// 7776000
	NewRefreshTokenValidity *int32 `json:"NewRefreshTokenValidity,omitempty" xml:"NewRefreshTokenValidity,omitempty"`
	// The required permission.
	//
	// You can specify one or more permissions for the `RequiredScopes` parameter. When a user grants permissions to the application, the scopes specified in this parameter are pre-selected and cannot be deselected.
	//
	// If you also specify the `NewPredefinedScopes` parameter, the `NewPredefinedScopes` parameter specifies the permissions that can be granted on the application, and this parameter specifies the required permissions.
	//
	// To enter multiple scopes, separate them with semicolons (;).
	//
	// The new value of this parameter overwrites the original value, and the required permission specified by the new value takes effect.
	//
	// > Any scope specified here must also be included in `PredefinedScopes`. Otherwise, the scope will not be set as required.
	//
	// example:
	//
	// profile;aliuid
	NewRequiredScopes *string `json:"NewRequiredScopes,omitempty" xml:"NewRequiredScopes,omitempty"`
	// Specifies whether a secret is required. Valid values:
	//
	// - true
	//
	// - false
	//
	// > 	- For applications of the WebApp and ServerApp types, this parameter is automatically set to true and cannot be changed.
	//
	// >
	//
	// > 	- For applications of the NativeApp type, this parameter can be set to true or false. If you do not set this parameter, false is used. Applications of the NativeApp type run in untrusted environments and the secrets of these applications are not protected.
	//
	// We recommend that you do not set this parameter to true unless otherwise specified. For more information, see [Access Alibaba Cloud APIs from a native application](https://help.aliyun.com/document_detail/93697.html).
	//
	// example:
	//
	// true
	NewSecretRequired *bool `json:"NewSecretRequired,omitempty" xml:"NewSecretRequired,omitempty"`
}

func (s UpdateApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateApplicationRequest) GetNewAccessTokenValidity() *int32 {
	return s.NewAccessTokenValidity
}

func (s *UpdateApplicationRequest) GetNewDisplayName() *string {
	return s.NewDisplayName
}

func (s *UpdateApplicationRequest) GetNewIsMultiTenant() *bool {
	return s.NewIsMultiTenant
}

func (s *UpdateApplicationRequest) GetNewPredefinedScopes() *string {
	return s.NewPredefinedScopes
}

func (s *UpdateApplicationRequest) GetNewRedirectUris() *string {
	return s.NewRedirectUris
}

func (s *UpdateApplicationRequest) GetNewRefreshTokenValidity() *int32 {
	return s.NewRefreshTokenValidity
}

func (s *UpdateApplicationRequest) GetNewRequiredScopes() *string {
	return s.NewRequiredScopes
}

func (s *UpdateApplicationRequest) GetNewSecretRequired() *bool {
	return s.NewSecretRequired
}

func (s *UpdateApplicationRequest) SetAppId(v string) *UpdateApplicationRequest {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewAccessTokenValidity(v int32) *UpdateApplicationRequest {
	s.NewAccessTokenValidity = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewDisplayName(v string) *UpdateApplicationRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewIsMultiTenant(v bool) *UpdateApplicationRequest {
	s.NewIsMultiTenant = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewPredefinedScopes(v string) *UpdateApplicationRequest {
	s.NewPredefinedScopes = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewRedirectUris(v string) *UpdateApplicationRequest {
	s.NewRedirectUris = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewRefreshTokenValidity(v int32) *UpdateApplicationRequest {
	s.NewRefreshTokenValidity = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewRequiredScopes(v string) *UpdateApplicationRequest {
	s.NewRequiredScopes = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewSecretRequired(v bool) *UpdateApplicationRequest {
	s.NewSecretRequired = &v
	return s
}

func (s *UpdateApplicationRequest) Validate() error {
	return dara.Validate(s)
}
