// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody
	GetApplication() *GetApplicationResponseBodyApplication
	SetRequestId(v string) *GetApplicationResponseBody
	GetRequestId() *string
}

type GetApplicationResponseBody struct {
	// The information about the application.
	Application *GetApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6616F09B-2768-4C11-8866-A8EE4C4A583E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) GetApplication() *GetApplicationResponseBodyApplication {
	return s.Application
}

func (s *GetApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationResponseBody) SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody {
	s.Application = v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyApplication struct {
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
	// The name of the application.
	//
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The type of the application. Valid values:
	//
	// 	- WebApp: a web application.
	//
	// 	- NativeApp: a native application that runs on an operating system, such as a desktop or mobile operating system.
	//
	// 	- ServerApp: an application that can access Alibaba Cloud services without the need for user logon. Only applications that synchronize user information based on the System for Cross-domain Identity Management (SCIM) protocol are supported.
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
	DelegatedScope *GetApplicationResponseBodyApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
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
	// The callback URL.
	RedirectUris *GetApplicationResponseBodyApplicationRedirectUris `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Struct"`
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

func (s GetApplicationResponseBodyApplication) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplication) GetAccessTokenValidity() *int32 {
	return s.AccessTokenValidity
}

func (s *GetApplicationResponseBodyApplication) GetAccountId() *string {
	return s.AccountId
}

func (s *GetApplicationResponseBodyApplication) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationResponseBodyApplication) GetAppName() *string {
	return s.AppName
}

func (s *GetApplicationResponseBodyApplication) GetAppType() *string {
	return s.AppType
}

func (s *GetApplicationResponseBodyApplication) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetApplicationResponseBodyApplication) GetDelegatedScope() *GetApplicationResponseBodyApplicationDelegatedScope {
	return s.DelegatedScope
}

func (s *GetApplicationResponseBodyApplication) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetApplicationResponseBodyApplication) GetIsMultiTenant() *bool {
	return s.IsMultiTenant
}

func (s *GetApplicationResponseBodyApplication) GetProtocolVersion() *string {
	return s.ProtocolVersion
}

func (s *GetApplicationResponseBodyApplication) GetRedirectUris() *GetApplicationResponseBodyApplicationRedirectUris {
	return s.RedirectUris
}

func (s *GetApplicationResponseBodyApplication) GetRefreshTokenValidity() *int32 {
	return s.RefreshTokenValidity
}

func (s *GetApplicationResponseBodyApplication) GetSecretRequired() *bool {
	return s.SecretRequired
}

func (s *GetApplicationResponseBodyApplication) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetApplicationResponseBodyApplication) SetAccessTokenValidity(v int32) *GetApplicationResponseBodyApplication {
	s.AccessTokenValidity = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAccountId(v string) *GetApplicationResponseBodyApplication {
	s.AccountId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAppId(v string) *GetApplicationResponseBodyApplication {
	s.AppId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAppName(v string) *GetApplicationResponseBodyApplication {
	s.AppName = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAppType(v string) *GetApplicationResponseBodyApplication {
	s.AppType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCreateDate(v string) *GetApplicationResponseBodyApplication {
	s.CreateDate = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetDelegatedScope(v *GetApplicationResponseBodyApplicationDelegatedScope) *GetApplicationResponseBodyApplication {
	s.DelegatedScope = v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetDisplayName(v string) *GetApplicationResponseBodyApplication {
	s.DisplayName = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetIsMultiTenant(v bool) *GetApplicationResponseBodyApplication {
	s.IsMultiTenant = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetProtocolVersion(v string) *GetApplicationResponseBodyApplication {
	s.ProtocolVersion = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetRedirectUris(v *GetApplicationResponseBodyApplicationRedirectUris) *GetApplicationResponseBodyApplication {
	s.RedirectUris = v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetRefreshTokenValidity(v int32) *GetApplicationResponseBodyApplication {
	s.RefreshTokenValidity = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetSecretRequired(v bool) *GetApplicationResponseBodyApplication {
	s.SecretRequired = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetUpdateDate(v string) *GetApplicationResponseBodyApplication {
	s.UpdateDate = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyApplicationDelegatedScope struct {
	// The information about the permissions that are granted on the application.
	PredefinedScopes *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s GetApplicationResponseBodyApplicationDelegatedScope) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationDelegatedScope) GetPredefinedScopes() *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes {
	return s.PredefinedScopes
}

func (s *GetApplicationResponseBodyApplicationDelegatedScope) SetPredefinedScopes(v *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) *GetApplicationResponseBodyApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

func (s *GetApplicationResponseBodyApplicationDelegatedScope) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) GetPredefinedScope() []*GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	return s.PredefinedScope
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
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

func (s GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GetName() *string {
	return s.Name
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GetRequired() *bool {
	return s.Required
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetRequired(v bool) *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Required = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) Validate() error {
	return dara.Validate(s)
}

type GetApplicationResponseBodyApplicationRedirectUris struct {
	RedirectUri []*string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyApplicationRedirectUris) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationRedirectUris) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationRedirectUris) GetRedirectUri() []*string {
	return s.RedirectUri
}

func (s *GetApplicationResponseBodyApplicationRedirectUris) SetRedirectUri(v []*string) *GetApplicationResponseBodyApplicationRedirectUris {
	s.RedirectUri = v
	return s
}

func (s *GetApplicationResponseBodyApplicationRedirectUris) Validate() error {
	return dara.Validate(s)
}
