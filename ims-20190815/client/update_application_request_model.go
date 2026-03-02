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
	// The new display name.
	//
	// example:
	//
	// NewApp
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	// Indicates whether the application can be installed by other Alibaba Cloud accounts. Valid values:
	//
	// - true: The application can be installed.
	//
	// - false: The application cannot be installed.
	//
	// example:
	//
	// true
	NewIsMultiTenant *bool `json:"NewIsMultiTenant,omitempty" xml:"NewIsMultiTenant,omitempty"`
	// The permission scopes of the application.
	//
	// For more information about the valid values and descriptions of permission scopes, see [OAuth scopes](https://help.aliyun.com/document_detail/93693.html). You can also call the [ListPredefinedScopes](https://help.aliyun.com/document_detail/187206.html) operation to obtain the permission scopes that are supported by different types of applications.
	//
	// If you enter multiple permission scopes, separate them with semicolons (;).
	//
	// The new permission scopes overwrite the original ones. For example, if the original permission scope is `/acs/ccc` and you set the new permission scope to `/acs/alidns`, the permission scope that takes effect is `/acs/alidns`. If you want to add `/acs/alidns` to the original scope, set the new permission scope to `/acs/ccc;/acs/alidns`.
	//
	// example:
	//
	// openid
	NewPredefinedScopes *string `json:"NewPredefinedScopes,omitempty" xml:"NewPredefinedScopes,omitempty"`
	// The webhook address.
	//
	// If you enter multiple webhook addresses, separate them with semicolons (;).
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
	// The required permission scopes of the application.
	//
	// You can set one or more scopes specified in `RequiredScopes` as required. After a scope is set as required, it is selected by default and cannot be deselected when a user grants permissions to the application.
	//
	// If you also specify `NewPredefinedScopes`, the list of application scopes is reset by `NewPredefinedScopes` first. Then, this parameter is used to configure whether the scopes are required.
	//
	// If you enter multiple permission scopes, separate them with semicolons (;).
	//
	// The new required scopes overwrite the original ones.
	//
	// > If a scope that you specify for `RequiredScopes` is not within the range of `PredefinedScopes`, the required setting for that scope does not take effect.
	//
	// example:
	//
	// profile;aliuid
	NewRequiredScopes *string `json:"NewRequiredScopes,omitempty" xml:"NewRequiredScopes,omitempty"`
	// Indicates whether an application key is required. Valid values:
	//
	// - true
	//
	// - false
	//
	// > 	- For applications of the WebApp and ServerApp types, this parameter is forcibly set to true and cannot be modified.
	//
	// - For applications of the NativeApp type, you can set this parameter to true or false. The default value is false. These applications often run in untrusted environments and cannot effectively protect application keys. Do not set this parameter to true unless necessary. For more information, see [Log on to Alibaba Cloud using a native application](https://help.aliyun.com/document_detail/93697.html).
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
