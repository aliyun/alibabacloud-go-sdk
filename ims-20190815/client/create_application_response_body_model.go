// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplication(v *CreateApplicationResponseBodyApplication) *CreateApplicationResponseBody
	GetApplication() *CreateApplicationResponseBodyApplication
	SetRequestId(v string) *CreateApplicationResponseBody
	GetRequestId() *string
}

type CreateApplicationResponseBody struct {
	// The information about the application.
	Application *CreateApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6616F09B-2768-4C11-8866-A8EE4C4A583E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) GetApplication() *CreateApplicationResponseBodyApplication {
	return s.Application
}

func (s *CreateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationResponseBody) SetApplication(v *CreateApplicationResponseBodyApplication) *CreateApplicationResponseBody {
	s.Application = v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationResponseBodyApplication struct {
	// The validity period of the access token. Unit: seconds.
	//
	// example:
	//
	// 3600
	AccessTokenValidity *int32 `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// The ID of the Alibaba Cloud account to which the application belongs.
	//
	// example:
	//
	// 177242285274****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type.
	//
	// example:
	//
	// WebApp
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-23T08:06:57Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted on the application.
	DelegatedScope *CreateApplicationResponseBodyApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// myapp
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Indicates whether the application can be installed by using other Alibaba Cloud accounts.
	//
	// example:
	//
	// true
	IsMultiTenant   *bool   `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty"`
	ProtocolVersion *string `json:"ProtocolVersion,omitempty" xml:"ProtocolVersion,omitempty"`
	// The callback URLs.
	RedirectUris *CreateApplicationResponseBodyApplicationRedirectUris `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Struct"`
	// The validity period of the refresh token. Unit: seconds.
	//
	// example:
	//
	// 7776000
	RefreshTokenValidity *int32 `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// Indicates whether a secret is required.
	//
	// example:
	//
	// true
	SecretRequired *bool `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-23T08:06:57Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateApplicationResponseBodyApplication) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplication) GetAccessTokenValidity() *int32 {
	return s.AccessTokenValidity
}

func (s *CreateApplicationResponseBodyApplication) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateApplicationResponseBodyApplication) GetAppId() *string {
	return s.AppId
}

func (s *CreateApplicationResponseBodyApplication) GetAppName() *string {
	return s.AppName
}

func (s *CreateApplicationResponseBodyApplication) GetAppType() *string {
	return s.AppType
}

func (s *CreateApplicationResponseBodyApplication) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateApplicationResponseBodyApplication) GetDelegatedScope() *CreateApplicationResponseBodyApplicationDelegatedScope {
	return s.DelegatedScope
}

func (s *CreateApplicationResponseBodyApplication) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateApplicationResponseBodyApplication) GetIsMultiTenant() *bool {
	return s.IsMultiTenant
}

func (s *CreateApplicationResponseBodyApplication) GetProtocolVersion() *string {
	return s.ProtocolVersion
}

func (s *CreateApplicationResponseBodyApplication) GetRedirectUris() *CreateApplicationResponseBodyApplicationRedirectUris {
	return s.RedirectUris
}

func (s *CreateApplicationResponseBodyApplication) GetRefreshTokenValidity() *int32 {
	return s.RefreshTokenValidity
}

func (s *CreateApplicationResponseBodyApplication) GetSecretRequired() *bool {
	return s.SecretRequired
}

func (s *CreateApplicationResponseBodyApplication) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *CreateApplicationResponseBodyApplication) SetAccessTokenValidity(v int32) *CreateApplicationResponseBodyApplication {
	s.AccessTokenValidity = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetAccountId(v string) *CreateApplicationResponseBodyApplication {
	s.AccountId = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetAppId(v string) *CreateApplicationResponseBodyApplication {
	s.AppId = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetAppName(v string) *CreateApplicationResponseBodyApplication {
	s.AppName = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetAppType(v string) *CreateApplicationResponseBodyApplication {
	s.AppType = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetCreateDate(v string) *CreateApplicationResponseBodyApplication {
	s.CreateDate = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetDelegatedScope(v *CreateApplicationResponseBodyApplicationDelegatedScope) *CreateApplicationResponseBodyApplication {
	s.DelegatedScope = v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetDisplayName(v string) *CreateApplicationResponseBodyApplication {
	s.DisplayName = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetIsMultiTenant(v bool) *CreateApplicationResponseBodyApplication {
	s.IsMultiTenant = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetProtocolVersion(v string) *CreateApplicationResponseBodyApplication {
	s.ProtocolVersion = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetRedirectUris(v *CreateApplicationResponseBodyApplicationRedirectUris) *CreateApplicationResponseBodyApplication {
	s.RedirectUris = v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetRefreshTokenValidity(v int32) *CreateApplicationResponseBodyApplication {
	s.RefreshTokenValidity = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetSecretRequired(v bool) *CreateApplicationResponseBodyApplication {
	s.SecretRequired = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetUpdateDate(v string) *CreateApplicationResponseBodyApplication {
	s.UpdateDate = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationResponseBodyApplicationDelegatedScope struct {
	// The information about the permissions that are granted on the application.
	PredefinedScopes *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s CreateApplicationResponseBodyApplicationDelegatedScope) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBodyApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScope) GetPredefinedScopes() *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes {
	return s.PredefinedScopes
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScope) SetPredefinedScopes(v *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) *CreateApplicationResponseBodyApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScope) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) GetPredefinedScope() []*CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	return s.PredefinedScope
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission.
	//
	// example:
	//
	// Obtain the OpenID of the user. This is the default permission that you cannot remove.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the permission is automatically selected by default when you install the application. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// `openid` is required by default.
	//
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GetName() *string {
	return s.Name
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GetRequired() *bool {
	return s.Required
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetRequired(v bool) *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Required = &v
	return s
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationResponseBodyApplicationRedirectUris struct {
	RedirectUri []*string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty" type:"Repeated"`
}

func (s CreateApplicationResponseBodyApplicationRedirectUris) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBodyApplicationRedirectUris) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplicationRedirectUris) GetRedirectUri() []*string {
	return s.RedirectUri
}

func (s *CreateApplicationResponseBodyApplicationRedirectUris) SetRedirectUri(v []*string) *CreateApplicationResponseBodyApplicationRedirectUris {
	s.RedirectUri = v
	return s
}

func (s *CreateApplicationResponseBodyApplicationRedirectUris) Validate() error {
	return dara.Validate(s)
}
